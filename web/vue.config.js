module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "^/api/status": {
        target: "http://localhost:8080",
        // target: "https://drx.pi6atv.ampr.org/",
        // pathRewrite: { '^/drx': '/receivers' },
        changeOrigin: true,
      },
    },
  }
}
