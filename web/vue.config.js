module.exports = {
  publicPath: '/winterhill/',
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "^/winterhill/api/": {
        target: "https://webcontrol.pi6atv.ampr.org",
        // target: "http://localhost:8080",
        changeOrigin: true,
      },
      "^/winterhill/ws/": {
        target: "https://webcontrol.pi6atv.ampr.org",
        changeOrigin: true,
      },
    },
  }
}
