import { createApp } from 'vue'

import App from './App.vue'

import axios from 'axios' // to use axios (get http request)
import VueAxios from 'vue-axios' // to use axios
import teoweb from 'teoweb' // to use teonet webrtc connection
import 'bootstrap/dist/css/bootstrap.css' // to use bootstrap CSS
import 'bootstrap/dist/js/bootstrap.js' // to use bootstrap JS
import './bootstrap_theme_tolge.js' // to use bootstrap color theme tolge
import './telergam_init.js' // to init telegram bot functionality

const app = createApp(App)
app.use(VueAxios, axios)
app.config.globalProperties.axios = axios;
app.config.globalProperties.teoweb = teoweb();
app.mount('#app')
