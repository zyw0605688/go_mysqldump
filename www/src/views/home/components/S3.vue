<template>
  <div>
    <div class="title">
      <div class="host">
        <div>本地备份位置</div>
        <el-popover
          placement="top-start"
          :width="200"
          trigger="hover"
          content="将映射到宿主机，请填写宿主机位置"
        >
          <template #reference>
            <el-icon style="top: 3px">
              <InfoFilled />
            </el-icon>
          </template>
        </el-popover>

        <el-input
          v-model="data.formData.hostPath"
          placeholder="如：/data/mysql_backup"
          clearable
          style="width: 260px; margin-left: 8px"
        ></el-input>
        <el-button type="primary" @click="showAddDialog" style="margin-left: 8px">确定</el-button>
      </div>
      <el-button @click="showAddDialog" style="margin-left: 8px">新增S3存储</el-button>
    </div>
    <el-table
      ref="multipleTable"
      style="width: 100%; height: 760px; margin-top: 16px"
      border
      :data="data.tableData"
    >
      <el-table-column prop="accessKey" label="AK" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="secretKey" label="SK" :show-overflow-tooltip="true"></el-table-column>
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
        <el-form-item label="AK" prop="accessKey">
          <el-input v-model="data.formData.accessKey" clearable></el-input>
        </el-form-item>
        <el-form-item label="SK" prop="secretKey">
          <el-input v-model="data.formData.secretKey" clearable></el-input>
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
    const params = JSON.parse(JSON.stringify(data.formData));
    console.log(params);
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

<style scoped lang="scss">
.title {
  display: flex;
  align-items: baseline;
  justify-content: space-between;

  .host {
    display: flex;
    align-items: baseline;
  }
}
</style>
