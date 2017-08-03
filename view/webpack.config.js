module.exports = {
    entry:'./src/index.js',
    output:{
        path:__dirname,
        filename:'bundle.js',
    },
    module:{
        loaders:[
            {test:/\.(js)$/, use:'babel-loader'},
            {test:/\.(css)$/, use:['style-loader', 'css-loader']}
        ]
    }
}
