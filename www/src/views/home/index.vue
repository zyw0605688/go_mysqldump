<template>
  <div>
    <Header />
    <div class="container">
      <S3 class="s3"></S3>
      <DB class="db"></DB>
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

  .s3 {
    width: 39%;
  }

  .db {
    margin-left: 8px;
    width: 60%;
  }


}
</style>
