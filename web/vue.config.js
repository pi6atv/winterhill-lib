module.exports = {
  publicPath: '/winterhill/',
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "^/winterhill/api/status": {
        // target: "https://webcontrol.pi6atv.ampr.org",
        target: "http://localhost:8080",
        // pathRewrite: { '^/drx': '/receivers' },
        changeOrigin: true,
      },
    },
  }
}
