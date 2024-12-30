<template>
  <div class="common-layout">
    <el-container>
      <el-main width="1402px" padding="0px">
        <el-header
          class="drag-area"
          style="height: 48px; display: flex; padding: 0 24px; justify-content: space-between"
        >
          <div style="margin-top: 24px"></div>
        </el-header>
        <div class="login_main">
          <div class="login_box">
            <div class="login_top"></div>
            <div class="login_loginIn">
              <div class="login_logoinIn_middle" style="margin-top: 24px">
                <el-form
                  ref="loginForm"
                  label-width="auto"
                  label-position="left"
                  :model="formData"
                  style="margin: 0 !important; padding: 0 !important"
                >
                  <el-form-item prop="user" label="用户名" style="margin-top: 16px">
                    <el-input
                      v-model="formData.user"
                      placeholder="用户名"
                      clearable
                      style="width: 310px"
                    />
                  </el-form-item>
                  <el-form-item prop="password" label="密码" style="margin-top: 16px">
                    <el-input
                      style="width: 310px"
                      class="mimaInput"
                      v-model="formData.password"
                      type="password"
                      clearable
                      placeholder="密码"
                      @keydown.enter.native="loginIn"
                    />
                  </el-form-item>
                </el-form>
                <el-button style="margin-top: 16px" @click="loginIn"> 登录 </el-button>
              </div>
            </div>
          </div>
        </div>
      </el-main>
      <!-- el-aside -->
      <el-aside width="480px">
        <div class="right_title"></div>
        <div class="right_img"></div>
      </el-aside>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import router from "@/router";

const formData = reactive({
  user: "admin",
  password: "123456"
});

const loginIn = (): void => {
  if (formData.user === "admin" && formData.password === "123456") {
    // 存储数据

    const token =
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsImlzcyI6Inptb3MifQ._9cV76YTSsMPw9yIJvjC6Dfdeaw4S1v0elIM9PczaBo";
    localStorage.setItem("token", token);
    router.push("/");
  }
};
</script>

<style lang="scss" scoped>
.el-main {
  padding: 0;
}

.title {
  line-height: 48px;
  width: 276px;
}

.el-aside {
  height: 100vh;
  background: linear-gradient(180deg, #00a870 0%, #52b292 100%);
}

.right_title {
  width: 360px;
  height: 104px;
  font-size: 16px;
  font-family: Medium;
  font-weight: 400;
  line-height: 35px;
  color: #ffffff;
  line-height: 40px;
  opacity: 1;
  position: relative;
  top: 210px;
  left: 81px;
}

.right_img {
  width: 90%;
  height: 362.86px;
  position: relative;
  top: 326px;
  left: 38px;
  overflow: hidden;

  img {
    width: 100%;
  }
}

.login_box {
  width: 324px;
  margin: 0 auto;
  margin-top: 187px;

  .login_loginIn {
    width: 323.75px;
    height: 430px;
    margin: 0 auto;
    transform: translateY(40px);
    overflow: hidden;

    .login_logoinIn_middle {
      margin-top: 24 !important;

      .el-input {
        height: 40px;
      }

      .login_button {
        margin-top: 32px;
        width: 325px;
        height: 48px;
        background: var(--primary-button-green);
        border-radius: 8px 8px 8px 8px;
        border: 0;
        opacity: 1;

        P {
          margin: 0 auto;
          width: 64px;
          height: 21px;
          font-size: 16px;
          font-weight: 400;
          color: #ffffff;
          line-height: 22px;
        }
      }

      .login_button:hover {
        background: var(--primary-button-green);
      }
    }
  }
}

.login_top {
  text-align: center;
}
</style>
