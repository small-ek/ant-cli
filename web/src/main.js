import {createApp} from 'vue'
import router from './routers'
import App from './App.vue'
import i18n from './lang'
const app = createApp(App)

import "./style/index.less"

app.use(router).use(i18n).mount('#app')
