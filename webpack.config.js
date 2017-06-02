var path = require('path');
var webpack = require('webpack');

module.exports = {
    entry: {
        'bundle': path.join(__dirname, 'public/js/app.js'),
    },
    output: {
        path: path.join(__dirname, 'public/dist'),
        filename: "[name].js",
    },
    cache: true,
    devtool: '#source-map',
    devServer: {
        contentBase: path.join(__dirname, 'public/dist'),
        inline: true,
        host: "0.0.0.0",
        port: 3000,
        proxy: {
            '/api/*': {
               target: 'http://localhost:8080',
               secure: false,
            }
        },
    },
    module: {
        loaders: [
            {
                test: /.js$/,
                exclude: /node_modules/,
                loader: "babel-loader",
                query: {
                    presets: ['react', 'es2015']
                }
            },
            {
                test: /\.html$/,
                loaders: "file?name=[name].[ext]",
            },
            {
                test: /\.css$/,
                loaders: "file?name=[name].[ext]",
            },
        ],
    },
}
