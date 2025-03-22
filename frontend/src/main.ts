import {createApp} from 'vue'
import {createPinia} from 'pinia'
import App from './App.vue'
import './style.css';
import i18n from './i18n'

const app = createApp(App)
const pinia = createPinia()

app.use(i18n)
app.use(pinia)
app.mount('#app')
