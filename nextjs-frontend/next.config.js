const { parsed: myEnv } = require('dotenv').config({
  path: `.bank-${process.env.BANK_CODE}.env`,
});
const webpack = require('webpack');

/** @type {import('next').NextConfig} */
const nextConfig = {
  distDir:
    process.env.NODE_ENV !== 'production'
      ? `.${myEnv.NEXT_PUBLIC_BANK_CODE}-next`
      : '.next',
  webpack: (config) => {
    config.plugins.push(new webpack.EnvironmentPlugin(myEnv));
    return config;
  },
  async redirects() {
    return [
      {
        source: '/',
        destination: '/bank-accounts',
        permanent: false,
      },
    ];
  },
};

module.exports = nextConfig;
