const path = require('path');
const fs = require('fs');

const HtmlWebPackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');
const PAGE_DIR = path.join('src', 'pages', path.sep);

const htmlPlugins = getFilesFromDir(PAGE_DIR, ['.html']).map((filePath) => {
    const fileName = filePath.replace(PAGE_DIR, '');
    return new HtmlWebPackPlugin({
        chunks: [fileName.replace(path.extname(fileName), ''), 'vendor'],
        template: filePath,
        filename: fileName,
    });
});

const entry = getFilesFromDir(PAGE_DIR, ['.jsx']).reduce((obj, filePath) => {
    const entryChunkName = filePath.replace(path.extname(filePath), '').replace(PAGE_DIR, '');
    obj[entryChunkName] = `./${filePath}`;
    return obj;
}, {});

module.exports = {
    entry: entry,
    plugins: [
        ...htmlPlugins,
        new CopyPlugin({
            patterns: [{ from: './src/favicon.ico' }],
        }),
    ],
    resolve: {
        alias: {
            src: path.resolve(__dirname, 'src'),
            components: path.resolve(__dirname, 'src', 'components'),
        },
    },
    module: {
        rules: [
            {
                test: /\.(js)|(jsx)$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env', '@babel/preset-react'],
                        plugins: ['transform-class-properties'],
                    },
                },
            },
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
        ],
    },
    optimization: {
        splitChunks: {
            cacheGroups: {
                vendor: {
                    test: /node_modules/,
                    chunks: 'initial',
                    name: 'vendor',
                    enforce: true,
                },
            },
        },
    },
};

function getFilesFromDir(dir, fileTypes) {
    const filesToReturn = [];
    function walkDir(currentPath) {
        const files = fs.readdirSync(currentPath);
        for (let i in files) {
            const curFile = path.join(currentPath, files[i]);
            if (fs.statSync(curFile).isFile() && fileTypes.indexOf(path.extname(curFile)) != -1) {
                filesToReturn.push(curFile);
            } else if (fs.statSync(curFile).isDirectory()) {
                walkDir(curFile);
            }
        }
    }
    walkDir(dir);
    return filesToReturn;
}
