<template>
  <div>
    <el-button @click="showAddDialog">新增S3存储</el-button>
    <el-table
      ref="multipleTable"
      style="width: 100%; height: 760px; margin-top: 16px"
      border
      :data="data.tableData"
    >
      <el-table-column
        prop="endpoint"
        label="访问域名"
        :show-overflow-tooltip="true"
      ></el-table-column>
      <el-table-column
        prop="bucketName"
        label="存储桶"
        :show-overflow-tooltip="true"
      ></el-table-column>
      <el-table-column prop="region" label="区域" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="accessKey" label="密钥" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="secretKey" label="密匙" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column align="left" label="操作">
        <template #default="scope">
          <el-button
            type="primary"
            link
            class="table-button"
            @click="getDetailAndShowUpdateFormDialog(scope.row)"
          >
            编辑
          </el-button>
          <el-button type="danger" link @click="deleteAccount(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      v-model="data.formDialogVisible"
      :title="data.type"
      destroy-on-close
      style="width: 500px"
    >
      <el-form
        :model="data.formData"
        label-position="right"
        ref="elFormRef"
        label-width="80px"
        style="margin-top: 16px"
      >
        <el-form-item label="名称" prop="accessKey">
          <el-input v-model="data.formData.name" placeholder="仅用于区分" clearable></el-input>
        </el-form-item>
        <el-form-item label="密钥" prop="accessKey">
          <el-input v-model="data.formData.accessKey" placeholder="访问密钥Access Key,secretId" clearable></el-input>
        </el-form-item>
        <el-form-item label="密匙" prop="secretKey">
          <el-input v-model="data.formData.secretKey" placeholder="秘密访问密钥Secret Access Key，secretKey" clearable></el-input>
        </el-form-item>
        <el-form-item label="访问域名" prop="endpoint">
          <el-input v-model="data.formData.endpoint" clearable></el-input>
        </el-form-item>
        <el-form-item label="存储桶" prop="bucketName">
          <el-input v-model="data.formData.bucketName" clearable></el-input>
        </el-form-item>
        <el-form-item label="区域" prop="region">
          <el-input v-model="data.formData.region" clearable></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeFormDialog">取 消</el-button>
          <el-button type="primary" @click="onSubmit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { reactive } from "_vue@3.5.13@vue";
import http from "@/service/http";

const data = reactive({
  activeName: "1",
  dbList: [],
  tableData: [] as any,
  formDialogVisible: false,
  formData: {},
  type: ""
});
const getTableData = async () => {
  const res = await http.get("/s3/list");
  data.tableData = res.data;
};
getTableData();
const deleteAccount = async (item: any) => {
  await http.delete(`/s3/delete?ID=${item.ID}`);
  await getTableData();
};

const showAddDialog = () => {
  data.formDialogVisible = true;
  data.type = "添加";
};
// 打开更新弹窗
const getDetailAndShowUpdateFormDialog = async (row) => {
  data.type = "编辑";
  data.formData = row;
  data.formDialogVisible = true;
};
const onSubmit = async () => {
  const params = JSON.parse(JSON.stringify(data.formData));
  await http.post("/s3/update", params);
  window.location.reload()
};

// 关闭弹窗
const closeFormDialog = () => {
  data.type = "";
  data.formDialogVisible = false;
  data.formData = {};
  data.dbList = [];
};
</script>

<style scoped lang="scss">
</style>
