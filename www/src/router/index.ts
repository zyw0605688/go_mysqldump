import { createRouter, createWebHashHistory, type RouteRecordRaw } from "vue-router";

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/login",
      name: "登录",
      component: () => import("@/views/login/index.vue"),
    },
    {
      path: "/",
      name: "首页",
      component: () => import("@/views/home/index.vue"),
    }
  ]
});

// 防止路由匹配失败
router.addRoute({
  path: "/:pathMatch(.*)*",
  redirect: "/",
  name: "NotFound"
});

router.beforeEach((to, from, next) => {
  // console.log("路由守卫", to, from);
  const token_permit =
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsImlzcyI6Inptb3MifQ._9cV76YTSsMPw9yIJvjC6Dfdeaw4S1v0elIM9PczaBo";
  const token = localStorage.getItem("token");
  if (token === token_permit) {
    if (to.path !== "/back") {
      next();
    } else {
      next({ path: "/" });
    }
  } else {
    if (to.path !== "/login") {
      next({ path: "/login" });
    } else {
      next();
    }
  }
});

export default router;
