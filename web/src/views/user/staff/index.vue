<template>
  <div class="main h-full pt-10px flex-col flex">
    <div class="flex justify-between top">
      <el-input
        clearable
        @clear="search"
        v-model="keyword"
        style="max-width: 600px"
        placeholder="请输入检索的账号"
        class="input-with-select"
      >
        <template #append>
          <el-button :icon="Search" @click="search" />
        </template>
      </el-input>
      <el-button size="large" type="primary" class="ml-20px" @click="addStudent"
        >新增员工</el-button
      >
    </div>
    <div class="flex-1 flex flex-col">
      <div class="flex-1">
        <el-auto-resizer>
          <template #default="{ height, width }">
            <el-table-v2
              :columns="columns"
              :data="tableData"
              :width="width"
              :height="height"
            />
          </template>
        </el-auto-resizer>
      </div>
      <div class="flex flex-justify-end pb-20px pr-30px">
        <el-pagination
          @change="changePageSize"
          :page-size="param.size"
          background
          layout="prev, pager, next"
          :total="total"
        />
      </div>
    </div>
    <el-drawer
      v-model="showStaffInfo"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <template #header="{ close }">
        <h4>{{ isAdd ? "创建管理员" : "修改资料" }}</h4>
        <el-icon style="cursor: pointer" @click="close"
          ><CloseBold :class="close"
        /></el-icon>
      </template>
      <div>
        <el-form ref="ruleFormRef" :model="form" label-width="auto">
      <el-form-item label="头像" :required="true" prop="avatar">
        <div
          class="flex flex-row mb-20px flex-content-center flex-wrap ml-10px"
        >
          <el-upload
            ref="uploadRef"
            class="avatar-uploader"
            :auto-upload="false"
            v-model:file-list="fileList"
            :show-file-list="false"
            :on-change="onchange"
            :before-upload="beforeAvatarUpload"
            accept=".jpeg,.png"
            :limit="1"
            :on-exceed="handleExceed"
          >
            <template #trigger>
              <img v-if="form.avatar" :src="form.avatar" class="avatar" />
              <el-icon v-else class="avatar-uploader-icon">
                <Plus />
              </el-icon>
            </template>
          </el-upload>
        </div>
      </el-form-item>
      <el-form-item label="昵称" required prop="name">
        <el-input v-model="form.name" placeholder="输入昵称"> </el-input>
      </el-form-item>
    
      <el-form-item label="联系方式" required prop="phone">
        <el-input v-model="form.phone" placeholder="请输入联系方式"> </el-input>
      </el-form-item>
   
      <template v-if="isAdd">
        <el-form-item label="账号" required prop="username">
          <el-input v-model="form.username" placeholder="设置账号"> </el-input>
        </el-form-item>
        <el-form-item label="密码" required prop="password">
          <el-input v-model="form.password" placeholder="设置密码"> </el-input>
        </el-form-item>
      </template>
    
    </el-form>
      </div>
      <template #footer>
      <div style="flex: auto">
        <el-button type="primary" @click="confirmClick">{{ isAdd?"保存" :"确认修改"}}</el-button>
      </div>
    </template>
    </el-drawer>
  </div>
</template>

<script lang="tsx" setup>
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import type { ComponentSize, FormInstance, FormRules } from "element-plus";
import {
    genFileId,
  UploadInstance,
  UploadProps,
  UploadRawFile,
  UploadUserFile,
  ElButton,
  ElIcon,
  ElMessage,
  ElTag,
  ElTooltip,
  TableV2FixedDir,
} from "element-plus";
import { getUserList } from "@/api/studen";
import companyInfo from "@/components/company-info.vue";
import {
  deleteStaffById,
  banUserById,
  initUserPassword,
  getComapnyStaffList,
  uploadFile,
  adminCreateStaff,
  editStaffById,
} from "@/api/common";
import { useUserStore } from "@/store/module/user";
const form = reactive({
  avatar: "",
  name: "",
  phone: "",
  username: "",
  password: "",
});
const showStaffInfo = ref(false);
const total = ref(0);
const tableData = ref([]);
const param = reactive({
  page: 1,
  size: 10,
});
const fileList = ref<UploadUserFile[]>([]);
const uploadRef = ref<UploadInstance>();
const userStore = useUserStore()

const onchange = (rawFile) => {
  if (rawFile.raw.type !== "image/jpeg" && rawFile.raw.type !== "image/png") {
    ElMessage.error("Avatar picture must be JPG or PNG format!");
    return false;
  }
  form.avatar = URL.createObjectURL(rawFile.raw);
  setTimeout(() => {
    upload();
  }, 1000);
};
const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
  if (rawFile.type !== "image/jpeg" && rawFile.type !== "image/png") {
    ElMessage.error("Avatar picture must be JPG or PNG format!");
    return false;
  }

  return true;
};
const handleExceed: UploadProps["onExceed"] = (files) => {
  uploadRef.value!.clearFiles();
  const file = files[0] as UploadRawFile;
  file.uid = genFileId();
  uploadRef.value!.handleStart(file);
};
const upload = () => {
  let formData = new FormData();
  let img = fileList.value[0].raw;
  formData.append("files", img as Blob);
  uploadFile(formData)
    .then((res) => {
      if (res.code === 200) {
        form.avatar = res.data[0].path;
        ElMessage.success("上传成功");
      } else {
        ElMessage.success("上传失败");
      }
    })
    .catch((res) => {
      ElMessage.success("上传失败");
    });
};
const item = ref({});
const columns: Column<any>[] = [
  {
    key: "name",
    title: "管理员",
    dataKey: "name",
    flexGrow:1,
    width: 100,
    align: "center",
  },

  {
    key: "phone",
    title: "联系方式",
    dataKey: "phone",
    flexGrow:1,
    width: 100,
    align: "center",

  },
  {
    key: "username",
    title: "账号",
    dataKey: "username",
    width: 100,
    align: "center",
    flexGrow:1,
  },
  {
    key: "options",
    title: "操作",
    dataKey: "options",
    flexGrow:1,
    width: 100,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return (
        <div class="flex flex-items-center">
          <el-button
          disabled={userStore.userInfo.user_id!==row.user_id &&  userStore.userInfo.user_id!==2}

            type="primary"
            onClick={() => {
              editStudentInfo(row);
            }}
          >
            编辑资料
          </el-button>
          <el-button
            type="danger"
            disabled={userStore.userInfo.user_id!==2 || userStore.userInfo.user_id===row.user_id}
            onClick={() => {
              deletStaff(row.user_id);
            }}
          >
            删除
          </el-button>
          <el-tooltip
            class="box-item"
            effect="light"
            content="初始化后密码为123456"
            placement="top-start"
          >
            <el-button
          disabled={userStore.userInfo.user_id!==row.user_id }
              style="background-color:#6E65E6;color:#fff;"
              onClick={() => {
                initPassword(row.user_id);
              }}
            >
              修改密码
            </el-button>
          </el-tooltip>
          <el-button
          disabled={userStore.userInfo.user_id!==2 }
            type="warning"
            onClick={() => {
              ban(row.user_id);
            }}
          >
            {row.state === 1 ? "禁用" : "解除"}
          </el-button>
        </div>
      );
    },
  },
];
const ruleFormRef = ref<FormInstance>();

const keyword = ref("");
onMounted(() => {
  loadTable();
});
const editStudentInfo = (row) => {
form.name= row.name
form.avatar = row.avatar
form.phone = row.phone
form.id = row.user_id
  showStaffInfo.value = true;
  isAdd.value = false;
};
const loadTable = () => {
  getComapnyStaffList({ search: keyword.value, ...param }).then((res) => {
    if ((res.code = 200)) {
      tableData.value = res.data.list;
      total.value = res.data.total;
    }
  });
};
const changePageSize = (e) => {
  param.page = e;
  loadTable();
};
const confirmClick = async()=>{
    await ruleFormRef.value!.validate((valid) => {

    if (valid) {
      if (isAdd.value) {
        adminCreateStaff(form).then((res) => {
          if (res.code === 200) {
            showStaffInfo.value =false
            search()
          }
          ElMessage.success(res.Message);
        });
      } else {
        editStaffById(form.id , form).then((res) => {
          if (res.code === 200) {
            showStaffInfo.value =false
    search()
          }
          ElMessage.success(res.Message);
        });
      }
    }
  });
    
}
const search = () => {
  param.page = 1;
  loadTable();
};
const isAdd = ref(true);
const close = (isFresh = false) => {
  showStaffInfo.value = false;
  if (isAdd.value) {
    param.page = 1;
  }
  loadTable();
};
const addStudent = () => {
  isAdd.value = true;
  showStaffInfo.value = true;
};
const deletStaff = (id) => {
  deleteStaffById(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
  });
};
const ban = (id) => {
  banUserById(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
  });
};
const initPassword = (id) => {
  initUserPassword(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
  });
};
</script>

<style scoped>
.main {
  padding-left: 8px;

  .top {
    padding: 10px;
    border-radius: 4px;
    background-color: aliceblue;
  }

  :deep(.el-table-v2__main) {
    top: 0;
    right: 0;
    left: 0;
    bottom: 0;
  }

  :deep(.el-table-v2__empty) {
    top: 50% !important;
    left: 50% !important;
    transform: translate(-50%, -50%);
  }
}
.close:hover {
  color: #409eff;
}
.form {
  position: absolute;
  top: 60px;
  background-color: #fff;
  z-index: 999999;
  left: 50%;
  transform: translateX(-50%);
}
</style>
<style>
.avatar-uploader {
  display: flex;
  justify-content: center;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}

.avatar-uploader .avatar {
  width: 140px;
  height: 140px;
  display: block;
}
</style>
