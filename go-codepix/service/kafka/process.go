package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/moura1001/codepix/domain/model"
	"github.com/moura1001/codepix/service/dto"
	"github.com/moura1001/codepix/service/factory"
	"github.com/moura1001/codepix/service/usecase"
	"gorm.io/gorm"
)

type KafkaProcessor struct {
	database                     *gorm.DB
	producer                     *ckafka.Producer
	deliveryChan                 chan ckafka.Event
	transactionsTopic            string
	transactionConfirmationTopic string
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		database:                     database,
		producer:                     producer,
		deliveryChan:                 deliveryChan,
		transactionsTopic:            os.Getenv("TRANSACTIONS_TOPIC"),
		transactionConfirmationTopic: os.Getenv("TRANSACTION_CONFIRMATION_TOPIC"),
	}
}

func (kp *KafkaProcessor) Consume() {
	msgChan := make(chan *ckafka.Message)
	consumer := NewKafkaConsumer(msgChan)
	go consumer.Consume([]string{kp.transactionsTopic, kp.transactionConfirmationTopic})

	for message := range msgChan {
		log.Printf("message received: %s\n", string(message.Value))
		go kp.processMessage(message)
	}
}

func (kp *KafkaProcessor) processMessage(message *ckafka.Message) {
	switch topic := *message.TopicPartition.Topic; topic {
	case kp.transactionsTopic:
		transactionId, transactionDTOInput, err := kp.processTransaction(message)
		if err != nil {
			log.Println(err)
			if transactionId != "" {
				err = kp.cancelTransactionNew(transactionId, err.Error(), transactionDTOInput, factory.NewTransactionUseCase(kp.database))
				if err != nil {
					log.Println(err)
				}

			} else if transactionDTOInput != nil {
				sourceBankTopic := "bank" + transactionDTOInput.BankCodeFrom
				transactionJsonOutput, err := dto.NewTransactionDTOOutputJsonNew(transactionDTOInput.RelatedTransactionIdFrom, model.TRANSACTION_STATUS_CANCELLED, err.Error(), *transactionDTOInput)
				if err != nil {
					log.Printf("error to inform new transaction cancel cause to source bank. Details: %s\n", err)
					return
				}

				err = kp.publishIntoBankTopic(sourceBankTopic, transactionJsonOutput)
				if err != nil {
					log.Printf("error to inform new transaction cancel cause to source bank. Details: %s\n", err)
				}

			}
		}
	case kp.transactionConfirmationTopic:
		err := kp.processTransactionConfirmation(message)
		if err != nil {
			log.Println(err)
		}
	default:
		log.Printf("unable to process messages from topic %s. Message consumed: %s\n", topic, string(message.Value))
	}
}

func (kp *KafkaProcessor) processTransaction(message *ckafka.Message) (string, *dto.TransactionDTOInputNew, error) {
	errMessageTemplate := "error during new transaction processing. Details: %s"

	transactionDTOInput, err := dto.NewTransactionDTOInputNew(message.Value)
	if err != nil {
		return "", transactionDTOInput, fmt.Errorf(errMessageTemplate, err)
	}

	transactionUseCase := factory.NewTransactionUseCase(kp.database)

	createdTransaction, err := transactionUseCase.Register(
		transactionDTOInput.BankCodeTo,
		transactionDTOInput.AccountNumberTo,
		transactionDTOInput.Amount,
		transactionDTOInput.PixKeyFrom,
		transactionDTOInput.PixKeyFromKind,
		transactionDTOInput.Description,
	)
	if err != nil {
		return "", transactionDTOInput, fmt.Errorf(errMessageTemplate, err)
	}

	destinationBankTopic := "bank" + createdTransaction.AccountTo.Bank.Code
	transactionJsonOutput, err := dto.NewTransactionDTOOutputJsonNew(createdTransaction.Id, createdTransaction.Status, "", *transactionDTOInput)
	if err != nil {
		return createdTransaction.Id, transactionDTOInput, fmt.Errorf(errMessageTemplate, err)
	}

	err = kp.publishIntoBankTopic(destinationBankTopic, transactionJsonOutput)
	if err != nil {
		return createdTransaction.Id, transactionDTOInput, fmt.Errorf(errMessageTemplate, err)
	}

	return createdTransaction.Id, transactionDTOInput, nil
}

func (kp *KafkaProcessor) processTransactionConfirmation(message *ckafka.Message) error {
	errMessageTemplate := "error during transaction confirmation processing. Details: %s"

	transactionDTOInput, err := dto.NewTransactionDTOInputExistent(message.Value)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	transactionUseCase := factory.NewTransactionUseCase(kp.database)

	switch transactionDTOInput.Status {
	case model.TRANSACTION_STATUS_CONFIRMED:
		err = kp.confirmTransaction(transactionDTOInput, transactionUseCase)
		if err != nil {
			return fmt.Errorf(errMessageTemplate, err)
		}
	case model.TRANSACTION_STATUS_COMPLETED:
		_, err = transactionUseCase.Complete(transactionDTOInput.Id)
		if err != nil {
			return fmt.Errorf(errMessageTemplate, err)
		}
	default:
		err = kp.cancelTransactionExistent(transactionDTOInput, transactionUseCase)
		if err != nil {
			return fmt.Errorf(errMessageTemplate, err)
		}
	}

	return nil
}

func (kp *KafkaProcessor) confirmTransaction(transactionInput *dto.TransactionDTOInputExistent, transactionUseCase usecase.TransactionUseCase) error {
	errMessageTemplate := "error to confirm transaction. Details: %s"

	confirmedTransaction, err := transactionUseCase.Confirm(transactionInput.Id)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	sourceBankTopic := "bank" + confirmedTransaction.PixKeyFrom.Account.Bank.Code
	transactionJsonOutput, err := dto.NewTransactionDTOOutputJsonExistent(*transactionInput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	err = kp.publishIntoBankTopic(sourceBankTopic, transactionJsonOutput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	return nil
}

func (kp *KafkaProcessor) cancelTransactionNew(id string, cancelDescription string, transactionInput *dto.TransactionDTOInputNew, transactionUseCase usecase.TransactionUseCase) error {
	errMessageTemplate := "error to cancel new transaction. Details: %s"

	cancelledTransaction, err := transactionUseCase.Cancel(id, cancelDescription)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	sourceBankTopic := "bank" + cancelledTransaction.PixKeyFrom.Account.Bank.Code
	transactionJsonOutput, err := dto.NewTransactionDTOOutputJsonNew(id, model.TRANSACTION_STATUS_CANCELLED, cancelDescription, *transactionInput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	err = kp.publishIntoBankTopic(sourceBankTopic, transactionJsonOutput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	return nil
}

func (kp *KafkaProcessor) cancelTransactionExistent(transactionInput *dto.TransactionDTOInputExistent, transactionUseCase usecase.TransactionUseCase) error {
	errMessageTemplate := "error to cancel existent transaction. Details: %s"

	cancelledTransaction, err := transactionUseCase.Cancel(transactionInput.Id, transactionInput.CancelDescription)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	sourceBankTopic := "bank" + cancelledTransaction.PixKeyFrom.Account.Bank.Code
	transactionJsonOutput, err := dto.NewTransactionDTOOutputJsonExistent(*transactionInput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	err = kp.publishIntoBankTopic(sourceBankTopic, transactionJsonOutput)
	if err != nil {
		return fmt.Errorf(errMessageTemplate, err)
	}

	return nil
}

func (kp *KafkaProcessor) publishIntoBankTopic(bankTopic string, message []byte) error {
	return Publish(string(message), bankTopic, kp.producer, kp.deliveryChan)

}
