import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/login",
      name: "登录",
      component: () => import("@/views/login/index.vue")
    },
    {
      path: "/",
      name: "首页",
      component: () => import("@/views/home/index.vue")
    }
  ]
});

router.beforeEach((to, from, next) => {
  const token_permit =
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsImlzcyI6Inptb3MifQ._9cV76YTSsMPw9yIJvjC6Dfdeaw4S1v0elIM9PczaBo";
  const token = localStorage.getItem("token");
  if (token === token_permit) {
    next();
  } else {
    if (to.path !== "/login") {
      next({ path: "/login" });
    } else {
      next();
    }
  }
});

export default router;
