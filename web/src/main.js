import Vue from 'vue'
import Router from "vue-router";
import App from './App.vue'
import vuetify from './plugins/vuetify'

Vue.use(Router)
Vue.config.productionTip = false
let router = new Router({
  mode: 'hash',
  root:  '/'
});


new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
