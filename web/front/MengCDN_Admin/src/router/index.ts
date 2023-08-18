import {createRouter, createWebHistory} from 'vue-router'
import AdminView from "@/views/AdminView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "Admin",
            component: AdminView,
            children: [
                {
                    path: "",
                    name: "Index",
                    component: () => import("@/views/Index.vue")
                },
                {
                    path: "/Github",
                    name: "Github",
                    component: () => import("@/views/Github.vue")
                },
                {
                    path: "/NPM",
                    name: "NPM",
                    component: () => import("@/views/NPM.vue")
                },
                {
                    path: "/WordPressPlugin",
                    name: "WordPressPlugin",
                    component: () => import("@/views/WordPress.vue")
                },
                {
                    path: "/WordPressTheme",
                    name: "WordPressTheme",
                    component: () => import("@/views/WordPressTh.vue")
                }
            ]
        },
        {
            path: '/Login',
            name:
                'Login',
            component:
                () => import('../views/LoginView.vue')
        }
    ]
})

router.beforeEach((to, from, next) => {

    //model 2
    const token = window.sessionStorage.getItem("token");
    if ((to.path === "/Login" || to.path === "/Login/") && token) {
        return next("/");
    }

    if (!token && (to.path != "/Login")) {
        return next("/Login");
    }

    return next();
});

export default router
