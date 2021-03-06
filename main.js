import '@babel/polyfill'
import 'mutationobserver-shim'
import Vue from 'vue'
import './plugins/bootstrap-vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

axios.defaults.baseURL = 'http://localhost:11000/api/'
// for deploy - uncomment next line
//axios.defaults.baseURL = '/api/' 


Vue.use(VueAxios, axios)
Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
