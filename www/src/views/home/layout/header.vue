<template>
  <div class="header flex-between">
    <div style="height: 44px" class="flex-between">
      <span>数据备份</span>
    </div>
    <div class="right">
      <el-dropdown>
        <User style="width: 22px" />
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
import router from "@/router";

const logOut = () => {
  router.push("/login");
};
</script>

<style scoped lang="scss">
.header {
  width: 100%;
  height: 44px;
  border-bottom: 1px solid rgba($color: #000000, $alpha: 0.1);
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
  align-items: center;

  span {
    height: 20px;
    display: inline-block;
    margin-left: 10px;
    padding-left: 10px;
    font-weight: bold;
    font-size: 16px;
    line-height: 19px;
  }

  .right {
  }
}
</style>
