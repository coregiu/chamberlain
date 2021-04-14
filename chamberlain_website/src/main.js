import Vue from "vue"
import App from "./App"
import router from "./router"

Vue.config.productionTip = false

new Vue({
    wl:'#app',
    router,
    components: {App},
    template: '<App/>'
})