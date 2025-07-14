    import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';

import Icons from './components/icons'
import "@/assets/main.less"

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Antd)

app.use(Icons)

app.mount('#app')
