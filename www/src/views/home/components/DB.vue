<template>
  <div>
    <el-button @click="showAddDialog">新增数据库</el-button>
    <el-table
      ref="multipleTable"
      style="width: 100%; height: 760px; margin-top: 16px"
      border
      :data="data.tableData"
    >
      <el-table-column type="index" label="No." width="55" align="center" />
      <el-table-column prop="host" label="主机" width="130" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="port" label="端口" width="90" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="username" label="用户名" width="90" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="is_local_store" label="本地存储" width="90" :show-overflow-tooltip="true">
        <template #default="scope">
          {{ scope.row.is_local_store ? "是" : "否"}}
        </template>
      </el-table-column>
      <el-table-column prop="dbs" label="数据库" :show-overflow-tooltip="true">
        <template #default="scope">
          <el-tag v-for="(item,index) in scope.row.dbs" :key="index" style="margin: 2px">{{item}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="cron" label="定时任务" :show-overflow-tooltip="true" width="120"></el-table-column>
      <el-table-column align="left" label="操作" width="140">
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
        <el-form-item label="主机" prop="host">
          <el-input v-model="data.formData.host" placeholder="主机ip地址" clearable></el-input>
        </el-form-item>
        <el-form-item label="端口号" prop="port">
          <el-input v-model="data.formData.port" placeholder="端口号" clearable></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="data.formData.username"
            placeholder="数据库用户名"
            clearable
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <div style="display: flex; width: 100%">
            <el-input
              v-model="data.formData.password"
              placeholder="数据库密码"
              clearable
              style="flex: 1"
            ></el-input>
            <el-button @click="getDbsByDsn">测试连接</el-button>
          </div>
        </el-form-item>
        <el-form-item label="数据库" prop="dbs">
          <el-select
            v-model="data.formData.dbs"
            multiple
            clearable
            placeholder="请选择要备份的数据库(多选)"
          >
            <el-option
              v-for="(item, index) in data.dbList"
              :key="index"
              :label="item"
              :value="item"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="定时任务" prop="cron">
          <el-input
            v-model="data.formData.cron"
            placeholder="请输入cron定时任务表达式"
            clearable
          ></el-input>
        </el-form-item>
        <el-form-item label="本地存储" prop="cron">
          <el-switch v-model="data.formData.is_local_store"></el-switch>
        </el-form-item>
        <el-form-item label="选择S3" prop="dbList">
          <el-select
            v-model="data.formData.s3s"
            multiple
            clearable
            placeholder="请选择要上传的s3存储"
          >
            <el-option
              v-for="(item, index) in data.s3List"
              :key="index"
              :label="item.name"
              :value="item.ID"
            ></el-option>
          </el-select>
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
import { onMounted } from "vue";

const data = reactive({
  dbList: [],
  s3List: [],
  tableData: [] as any,
  formDialogVisible: false,
  formData: {
    is_local_store: true
  },
  type: ""
});
const getTableData = async () => {
  const res = await http.get("/db/list");
  data.tableData = res.data;
};
const getS3Data = async () => {
  const res = await http.get("/s3/list");
  data.s3List = res.data;
};
onMounted(async () => {
  await getS3Data();
  await getTableData();
});

const getDbsByDsn = async()=>{
  const { host, port, username, password } = data.formData;
  const dsn = `${username}:${password}@tcp(${host}:${port})/`;
  const res = await http.post("/other/getDbsByDsn",{dsn});
  data.dbList = res.data
}

const deleteAccount = async (item: any) => {
  await http.delete(`/db/delete?ID=${item.ID}`);
  await getTableData();
};

const showAddDialog = async () => {
  data.formDialogVisible = true;
  data.type = "添加";
  await getS3Data()
};
// 打开更新弹窗
const getDetailAndShowUpdateFormDialog = async (row) => {
  data.type = "编辑";
  data.formData = row;
  data.formDialogVisible = true;
  await getDbsByDsn()
  await getS3Data()
};
const onSubmit = async () => {
  const params = JSON.parse(JSON.stringify(data.formData));
  await http.post("/db/update", params);
  await getTableData();
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

<style scoped lang="scss"></style>
