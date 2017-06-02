var path = require('path');
var webpack = require('webpack');
var HtmlwebpackPlugin = require('html-webpack-plugin');
var merge = require('webpack-merge');

const ENV = process.env.npm_lifecycle_event;
process.env.BABEL_ENV = ENV;

const common = {
    entry: {
        'bundle': path.join(__dirname, 'public/js/app.js'),
    },
    output: {
        path: path.join(__dirname, 'public/dist'),
        filename: "[name].js",
    },
    cache: true,
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
    plugins: [
        new webpack.ProvidePlugin({
            jQuery: 'jquery',
            $: 'jquery',
            Tether: 'tether',
            React: 'react',
            ReactDOM: 'react-dom',
        }),
        new HtmlwebpackPlugin({
            title: 'revel-base',
        }),
    ],
};

if (ENV === 'start' || !ENV) {
    module.exports = merge(common, {
        devtool: '#source-map',
        devServer: {
            contentBase: path.join(__dirname, 'public/dist'),
            publicPath: '/public/',
            historyApiFallback: false,
            hot: true,
            inline: true,
            host: "0.0.0.0",
            port: 3000,
            proxy: {
                '**': {
                    target: 'http://localhost:8080',
                    secure: false,
                },
            },
        },
        plugins: [
            new webpack.HotModuleReplacementPlugin()
        ]
    });
} else if (ENV === 'build') {
    module.exports = common;
}
