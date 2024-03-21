import {createApp} from 'vue'
import router from './routers'
import App from './App.vue'
const app = createApp(App)

import "./style/index.less"

app.use(router).mount('#app')
