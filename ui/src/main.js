import g from 'guark'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import VueHotkey from "v-hotkey";

Vue.use(VueHotkey);

Vue.config.productionTip = false

new Vue({
	store,
	render: h => h(App),
	created: () => g.hook("created"),
	mounted: () => g.hook("mounted"),
}).$mount('#app')