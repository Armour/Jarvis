import path from 'path';
import merge from 'webpack-merge';

import CopyWebpackPlugin from 'copy-webpack-plugin';
import OfflinePlugin from 'offline-plugin';

import BaseWebpackConfig from './webpack.config.base.babel';

export default merge(BaseWebpackConfig, {
  // The point or points to enter the application.
  entry: {
    app: './frontend/src/index',
  },

  // Affecting the output of the compilation
  output: {
    // path: the output directory as an absolute path (required)
    path: path.resolve(__dirname, 'frontend/dist/prod'),
    // filename: specifies the name of output file on disk (required)
    filename: '[name].[chunkhash:10].js',
    // publicPath: specifies the server-relative URL of the output resource directory
    // https://webpack.js.org/configuration/output/#output-publicpath
    publicPath: '/',
  },

  // A list of used webpack plugins
  plugins: [
    // Copy static files to build dir
    new CopyWebpackPlugin([
      {
        from: 'frontend/public/*',
        to: '[name].[ext]',
        ignore: [
          'favicon.ico',
          'index.ejs',
          'logo.png',
        ],
      },
    ]),
    // It's always better if OfflinePlugin is the last plugin added
    new OfflinePlugin(),
  ],

  // Source map mode
  // https://webpack.js.org/configuration/devtool
  devtool: 'source-map',
});
