import {createRouter, createWebHashHistory} from 'vue-router'
import Blog from "../views/blog.vue"
import Notebook from "../views/notebook.vue"
import Dashboard from "../views/dashboard.vue"
import InputDetails from "../views/input_details.vue"
import InputAnalysis from "../views/input_analysis.vue"
import Login from "../views/login.vue"
import Logout from "../views/logout.vue"
import UserMgmt from "../views/user_mgmt.vue"
import SysBack from "../views/sys_back.vue"
import SysLog from "../views/sys_log.vue"
import UserInfo from "../views/user_info.vue"
import UpdatePassword from "../views/update_password.vue"

const routes: Array<any> = [
    {
        path: "/",
        name: "Home",
        component: Dashboard,
    },{
        path: "/blog",
        name: "Blog",
        component: Blog,
    },
    {
        path: "/notebook",
        name: "Notebook",
        component: Notebook,
    },
    {
        path: "/input_details",
        name: "InputDetails",
        component: InputDetails,
    },
    {
        path: "/input_analysis",
        name: "InputAnalysis",
        component: InputAnalysis,
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
    {
        path: "/logout",
        name: "Logout",
        component: Logout,
    },
    {
        path: "/user_mgmt",
        name: "UserMgmt",
        component: UserMgmt,
    },
    {
        path: "/sys_back",
        name: "SysBack",
        component: SysBack,
    },
    {
        path: "/sys_log",
        name: "SysLog",
        component: SysLog,
    },
    {
        path: "/user_info",
        name: "UserInfo",
        component: UserInfo,
    },
    {
        path: "/update_password",
        name: "UpdatePassword",
        component: UpdatePassword,
    },
    {
        path: "/logout",
        name: "Logout",
        component: Logout,
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;