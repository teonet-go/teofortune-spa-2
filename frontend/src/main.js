import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

import teoweb from 'teoweb'

const app = createApp(App)
app.use(VueAxios, axios)

app.config.globalProperties.axios = axios;
app.config.globalProperties.teoweb = teoweb();

app.mount('#app')
