<template>
  <el-steps
    style="max-width: 500px;margin: 16px 0"
    :active="props.active"
    align-center
    finish-status="success"
  >
    <el-step title="开始" />
    <el-step title="数据库配置" />
    <el-step title="备份任务配置" />
    <el-step title="存储配置" />
  </el-steps>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps({
  active: Number
});
</script>
