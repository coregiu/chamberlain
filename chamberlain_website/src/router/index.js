import Vue from "vue"
import Router from "vue-router"
import Blog from "@/components/blog"
import Dashboard from "@/components/dashboard"
import Input from "@/components/input"
import Login from "@/components/login"
import System from "@/components/system"
import User from "@/components/user"

Vue.use(Router)

export default new Router({
    mode: "history",
    router:[{
        path: '/',
        name: 'blog',
        component: Blog
    },{
        path: '/blog',
        name: 'blog',
        component: Blog
    },{
        path: '/dashboard',
        name: 'dashboard',
        component: Dashboard
    },{
        path: '/input',
        name: 'input',
        component: Input
    },{
        path: '/login',
        name: 'login',
        component: Login
    },{
        path: '/system',
        name: 'system',
        component: System
    },{
        path: '/user',
        name: 'user',
        component: User
    }]
})