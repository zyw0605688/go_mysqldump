<template>
  <div>
    <Header />
    <div class="container">
      <DB class="db"></DB>
      <S3 class="s3"></S3>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Header from "./layout/header.vue";
import DB from "./components/DB.vue";
import S3 from "./components/S3.vue";
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  margin-top: 16px;

  .db {
    width: 60%;
  }

  .s3 {
    width: 39%;
    margin-left: 8px;
  }
}
</style>
