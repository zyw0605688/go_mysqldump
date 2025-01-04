<template>
  <el-container>
    <el-main>
      <el-form ref="loginForm" label-width="auto" label-position="left" :model="formData">
        <el-form-item prop="user" label="用户名">
          <el-input v-model="formData.user" placeholder="用户名" clearable style="width: 310px" />
        </el-form-item>
        <el-form-item prop="password" label="密码">
          <el-input
            style="width: 310px"
            v-model="formData.password"
            type="password"
            clearable
            placeholder="密码"
            @keydown.enter.native="loginIn"
          />
        </el-form-item>
        <el-form-item>
          <el-button @click="loginIn">登录</el-button>
        </el-form-item>
      </el-form>
    </el-main>
    <el-aside width="480px"> </el-aside>
  </el-container>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const formData = reactive({
  user: "admin",
  password: "123456"
});

const loginIn = () => {
  if (formData.user === "admin" && formData.password === "123456") {
    const token =
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsImlzcyI6Inptb3MifQ._9cV76YTSsMPw9yIJvjC6Dfdeaw4S1v0elIM9PczaBo";
    localStorage.setItem("token", token);
    router.push("/");
  }
};
</script>

<style lang="scss" scoped>
.el-container {
  display: flex;

  .el-main {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .el-aside {
    height: 98vh;
    background: linear-gradient(180deg, #00a870 0%, #52b292 100%);
  }
}
</style>
