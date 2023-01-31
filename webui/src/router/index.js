import {createRouter, createWebHashHistory} from 'vue-router'
import Search from '../views/Search.vue'
import Login from '../views/doLogin.vue'
import Profile from '../views/Profile.vue'
import MyProfile from '../views/MyProfile.vue'
import Stream from '../views/Stream.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/ricerca', component: Search},
		{path: '/session', component: Login},
		{path: '/ricerca/profile/:username', component: Profile},
		{path: '/profile/', component: MyProfile},
		{path: '/stream', component: Stream},
	]
})

export default router
