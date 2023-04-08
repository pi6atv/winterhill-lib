module.exports = {
  publicPath: '.',
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "^/api/": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
      "^/ws/": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  }
}
