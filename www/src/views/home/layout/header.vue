<template>
  <div class="header">
    <div class="left">
      MySql 数据备份
      <el-popover placement="top-start" :width="480" trigger="hover">
        <template #reference>
          <el-icon style="margin-left: 4px">
            <InfoFilled />
          </el-icon>
        </template>
        <template #default>
          <ul>
            <li>一个定时备份mysql数据库的小工具。</li>
            <li>支持备份到本机，也可以上传s3云存储，或两者共同使用。</li>
            <li>如果开启保存到本地，将映射/mysql_backup到宿主机。</li>
            <li>如需上传s3,请先添加s3配置。</li>
          </ul>
        </template>
      </el-popover>
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

const content = `


`;

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
