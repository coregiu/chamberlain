import VueRouter  from 'vue-router'
import Blog from "../views/blog.vue"
import Dashboard from "../views/dashboard.vue"
import Input from "../views/input.vue"
import Login from "../views/login.vue"
import System from "../views/system.vue"
import User from "../views/user.vue"

const routes: Array<any> = [
    {
        path: "/",
        name: "Home",
        component: Blog,
    },{
        path: "/blog",
        name: "Blog",
        component: Blog,
    },
    {
        path: "/dashboard",
        name: "Dashboard",
        component: Dashboard,
    },
    {
        path: "/input",
        name: "Input",
        component: Input,
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
    {
        path: "/system",
        name: "System",
        component: System,
    },
    {
        path: "/user",
        name: "User",
        component: User,
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
});

export default router;