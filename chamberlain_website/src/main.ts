import {createApp} from 'vue'
import router from './router'
import App from './App.vue'
import './index.scss'

// @ts-ignore
createApp(App).use(router).mount('#app')