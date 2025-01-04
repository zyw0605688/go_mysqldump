<template>
  <div>
    <Header />
    <div class="db">
      <el-button type="primary" @click="showAddDialog" style="margin-top: 16px">新增</el-button>
      <el-table
        ref="multipleTable"
        style="width: 100%; height: 760px; margin-top: 16px"
        border
        :data="data.tableData"
      >
        <el-table-column type="index" label="No." width="55" align="center" />
        <el-table-column
          prop="userName"
          label="账号"
          :show-overflow-tooltip="true"
        ></el-table-column>
        <el-table-column
          prop="nickName"
          label="昵称"
          :show-overflow-tooltip="true"
        ></el-table-column>
        <el-table-column
          prop="phone"
          label="手机号"
          :show-overflow-tooltip="true"
        ></el-table-column>
        <el-table-column align="left" label="操作" width="220">
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
    </div>
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
        <el-collapse v-model="data.activeName" accordion>
          <el-collapse-item title="数据库配置" name="1">
            <el-form-item label="主机ip" prop="host">
              <el-input v-model="data.formData.host" clearable></el-input>
            </el-form-item>
            <el-form-item label="端口号" prop="port">
              <el-input v-model="data.formData.port" clearable></el-input>
            </el-form-item>
            <el-form-item label="用户名" prop="username">
              <el-input v-model="data.formData.username" clearable></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <div style="display: flex; width: 100%">
                <el-input v-model="data.formData.password" clearable style="flex: 1"></el-input>
                <el-button>测试连接</el-button>
              </div>
            </el-form-item>
            <el-form-item label="数据库" prop="dbList">
              <el-select v-model="data.formData.dbList" multiple clearable>
                <el-option
                  v-for="(item, index) in data.dbList"
                  :key="index"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="任务配置" name="2">
            <el-form-item label="定时任务" prop="cron">
              <el-input v-model="data.formData.cron" placeholder="请输入cron定时任务表达式" clearable></el-input>
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="存储配置" name="3">
            <el-form-item label="本机路径" prop="hostPath">
              <el-input v-model="data.formData.hostPath" clearable></el-input>
            </el-form-item>
            <div>S3配置</div>
            <el-form-item label="secretID" prop="secretID">
              <el-input v-model="data.formData.secretID" clearable></el-input>
            </el-form-item>
            <el-form-item label="secretKey" prop="secretKey">
              <el-input v-model="data.formData.secretKey" clearable></el-input>
            </el-form-item>
            <el-form-item label="endpoint" prop="endpoint">
              <el-input v-model="data.formData.endpoint" clearable></el-input>
            </el-form-item>
            <el-form-item label="存储桶" prop="bucketName">
              <el-input v-model="data.formData.bucketName" clearable></el-input>
            </el-form-item>
            <el-form-item label="区域" prop="region">
              <el-input v-model="data.formData.region" clearable></el-input>
            </el-form-item>
          </el-collapse-item>
        </el-collapse>
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

<script lang="ts" setup>
import { reactive } from "vue";
import http from "@/service/http";
import Header from "./layout/header.vue";

const data = reactive({
  activeName: "1",
  dbList: [],
  tableData: [] as any,
  formDialogVisible: false,
  formData: {},
  type: ""
});
const getTableData = async () => {
  const res = await http.get("/db");
  data.tableData = res.data.list;
};
getTableData();
const deleteAccount = async (item: any) => {
  await http.delete(`/db?id=${item.id}`);
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
};
const onSubmit = async () => {
  if (data.type != "info") {
    const params = JSON.parse(JSON.stringify(data.formData))
    console.log(params)
    await http.post("/db", params);
    await getTableData();
  }
  closeFormDialog();
};

// 关闭弹窗
const closeFormDialog = () => {
  data.type = "";
  data.formDialogVisible = false;
  data.formData = {};
  data.dbList = [];
};
</script>

<style lang="scss" scoped></style>
