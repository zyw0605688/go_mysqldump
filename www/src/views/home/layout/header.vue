<template>
  <div class="header">
    <div class="left">
      MySql 数据备份
    </div>
    <div class="right">
      <el-dropdown>
        <User style="width: 24px" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item class="menu-item" @click="logOut">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useRouter } from "vue-router";

const router = useRouter();

const logOut = () => {
  router.push("/login");
};
</script>

<style scoped lang="scss">
.header {
  width: 100%;
  height: 48px;
  border-bottom: 1px solid rgba($color: #000000, $alpha: 0.1);
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
  align-items: center;

  .left {
    font-size: 18px;
    display: flex;
    align-items: center;
  }
}
</style>
