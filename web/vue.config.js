module.exports = {
  publicPath: '/winterhill/',
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "^/winterhill/api/": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
      "^/winterhill/ws/": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  }
}
