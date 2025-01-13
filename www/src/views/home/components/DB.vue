<template>
  <div>
    <el-button @click="showAddDialog">新增数据库</el-button>
    <el-popover
      placement="top-start"
      :width="200"
      trigger="hover"
      content="若更新备份配置，请点击完成手动重载"
    >
      <template #reference>
        <el-button @click="Reload">重载备份</el-button>
      </template>
    </el-popover>
    <el-table
      ref="multipleTable"
      style="width: 100%; height: 760px; margin-top: 16px"
      border
      :data="data.tableData"
    >
      <el-table-column type="index" label="No." width="55" align="center" />
      <el-table-column
        prop="host"
        label="主机"
        width="130"
        :show-overflow-tooltip="true"
      ></el-table-column>
      <el-table-column
        prop="port"
        label="端口"
        width="90"
        :show-overflow-tooltip="true"
      ></el-table-column>
      <el-table-column
        prop="username"
        label="用户名"
        width="90"
        :show-overflow-tooltip="true"
      ></el-table-column>
      <el-table-column
        prop="is_local_store"
        label="本地存储"
        width="90"
        :show-overflow-tooltip="true"
      >
        <template #default="scope">
          {{ scope.row.is_local_store ? "是" : "否" }}
        </template>
      </el-table-column>
      <el-table-column prop="dbs" label="数据库" :show-overflow-tooltip="true">
        <template #default="scope">
          <el-tag v-for="(item, index) in scope.row.dbs" :key="index" style="margin: 2px"
            >{{ item }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="cron"
        label="定时任务"
        :show-overflow-tooltip="true"
        width="120"
      ></el-table-column>
      <el-table-column align="left" label="操作" width="180">
        <template #default="scope">
          <el-button
            type="primary"
            link
            class="table-button"
            @click="getDetailAndShowUpdateFormDialog(scope.row)"
          >
            编辑
          </el-button>
          <el-popconfirm title="确定要删除吗？" @confirm="deleteAccount(scope.row)">
            <template #reference>
              <el-button type="danger" link>删除</el-button>
            </template>
          </el-popconfirm>
          <el-button link @click="GetBackupList(scope.row)">查看备份</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      v-model="data.formDialogVisible"
      :title="data.type"
      destroy-on-close
      :before-close="closeFormDialog"
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
        <el-form-item label="启用备份" prop="is_backup">
          <el-switch v-model="data.formData.is_backup"></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeFormDialog">取 消</el-button>
          <el-button type="primary" @click="onSubmit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-drawer
      v-model="data.drawerVisible"
      title="本地备份文件"
      direction="rtl"
      :before-close="beforeCloseDrawer"
    >
      <el-table :data="data.backUpFileList" stripe border>
        <el-table-column type="index" label="No." width="60" align="center" />
        <el-table-column prop="file" label="文件" :show-overflow-tooltip="true"></el-table-column>
      </el-table>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import { reactive, onMounted } from "vue";
// @ts-ignore
import { reload, s3list, DbsByDsn, dbList, dbUpdate, dbDelete, getBackupList } from "@/service/api";
import { ElMessage } from "element-plus";

const data = reactive({
  dbList: [],
  s3List: [],
  tableData: [] as any,
  formDialogVisible: false,
  drawerVisible: false,
  formData: {
    is_local_store: true,
    is_backup: false,
    host: "",
    port: "",
    username: "",
    password: "",
    dbs: [],
    cron: "",
    s3s: []
  },
  type: "",
  backUpFileList: []
});
const getTableData = async () => {
  const res = await dbList();
  data.tableData = res.data;
};
const getS3Data = async () => {
  const res = await s3list();
  data.s3List = res.data;
};
onMounted(async () => {
  await getS3Data();
  await getTableData();
});

const getDbsByDsn = async () => {
  const { host, port, username, password } = data.formData;
  const dsn = `${username}:${password}@tcp(${host}:${port})/`;
  const res = await DbsByDsn({ dsn });
  data.dbList = res.data;
};

const deleteAccount = async (item: any) => {
  await dbDelete(item.ID);
  await getTableData();
};

const showAddDialog = async () => {
  data.formDialogVisible = true;
  data.type = "添加";
  await getS3Data();
};
// 打开更新弹窗
const getDetailAndShowUpdateFormDialog = async (row) => {
  data.type = "编辑";
  data.formData = row;
  data.formDialogVisible = true;
  await getDbsByDsn();
  await getS3Data();
};
const onSubmit = async () => {
  const params = JSON.parse(JSON.stringify(data.formData));
  await dbUpdate(params);
  await getTableData();
  closeFormDialog();
};

// 关闭弹窗
const closeFormDialog = () => {
  data.type = "";
  data.formDialogVisible = false;
  data.formData = {
    is_local_store: true,
    is_backup: false,
    host: "",
    port: "",
    username: "",
    password: "",
    dbs: [],
    cron: "",
    s3s: []
  };
  data.dbList = [];
};

const Reload = async () => {
  await reload();
  ElMessage.success("重新加载定时任务成功！");
};

const GetBackupList = async (val) => {
  const res = await getBackupList(val.ID);
  data.drawerVisible = true;
  data.backUpFileList = res.data;
};
const beforeCloseDrawer = () => {
  data.drawerVisible = false;
  data.backUpFileList = [];
};
</script>

<style scoped lang="scss"></style>
