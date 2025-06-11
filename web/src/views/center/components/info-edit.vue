<template>
  <el-drawer v-model="showDrawer" :with-header="false" :close-on-click-modal="false">
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
      <el-form-item
        :label="userStore.userInfo.role === 1 ? '昵称' : '公司名称'"
        required
        prop="name"
      >
        <el-input v-model="form.name" placeholder="请输入"> </el-input>
      </el-form-item>

      <el-form-item label="联系方式" required prop="phone">
        <el-input v-model="form.phone" placeholder="请输入联系方式"> </el-input>
      </el-form-item>
      <el-form-item
        v-if="userStore.userInfo.role === 1"
        label="年龄"
        required
        prop="age"
      >
        <el-input
          class="input"
          type="number"
          v-model="form.age"
          placeholder="Please input"
          size="large"
        />
      </el-form-item>
      <el-form-item
        v-if="userStore.userInfo.role === 1"
        label="性别"
        required
        prop="sex"
      >
        <el-radio-group v-model="form.sex">
          <el-radio :value=1 size="large">男</el-radio>
          <el-radio :value=2 size="large">女</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item
        v-if="userStore.userInfo.role === 2"
        label="法人"
        required
        prop="lega_person"
      >
        <el-input
          class="input"
          v-model="form.lega_person"
          placeholder="Please input"
          size="large"
        />
      </el-form-item>
      <el-form-item label="所在校区" required prop="school_id">
        <el-select
          style="cursor: pointer; width: 240px"
          v-model="form.school_id"
          filterable
          remote
          reserve-keyword
          placeholder="收入您检索的校区"
        >
          <el-option
            v-for="item in allSchoolList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        v-if="userStore.userInfo.role === 1"
        label="专业"
        required
        prop="major"
      >
        <el-input
          class="input"
          v-model="form.major"
          placeholder="Please input"
          size="large"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancelClick">取消</el-button>
        <el-button type="primary" @click="confirmClick">确认修改</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script lang="ts" setup>
import { editCompanyInfo } from "@/api/center";
import { editStudentInfo, getSchoolList, uploadFile } from "@/api/common";
import { useUserStore } from "@/store/module/user";
import { formToJSON } from "axios";
import { log } from "console";

import {
  ElMessage,
  FormInstance,
  genFileId,
  UploadInstance,
  UploadProps,
  UploadRawFile,
  UploadUserFile,
} from "element-plus";
import { onMounted, reactive, ref } from "vue";
let userStore = useUserStore();

const form = reactive({
  avatar: userStore.userInfo?.avatar,
  age: userStore.userInfo?.age,
  phone: userStore.userInfo?.phone,
  name: userStore.userInfo?.name,
  sex: Number(userStore.userInfo?.sex),
  school_id: userStore.userInfo?.school_id,
  major: userStore.userInfo?.major,
  lega_person: userStore.userInfo?.lega_person,
});
const ruleFormRef = ref<FormInstance>();

const allSchoolList = reactive<any[]>([]);

const fileList = ref<UploadUserFile[]>([]);
const uploadRef = ref<UploadInstance>();
const showDrawer = ref(true);
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
const emit = defineEmits(["close"]);
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
onMounted(() => {
  getSchoolList(0).then((res) => {
    if (res.code !== 200) {
      ElMessage.success(res.Message ? res.Message : "数据异常");
    } else {
      allSchoolList.push(...res.data.list);
    }
  });
});
const cancelClick = () => {
  emit("close");
};
const confirmClick = async () => {
  console.log(form)
  await ruleFormRef.value!.validate((valid) => {
    if (valid) {
      if (userStore.userInfo.role === 1) {
        editStudentInfo(userStore.userInfo.id, form).then((res) => {
          if (res.code == 200) {
            userStore.setUserInfo(res.data);
            ElMessage.success(res.Message);
          } else {
            ElMessage.error(res.Message);
          }
          emit('close')

        });
      } else {
        editCompanyInfo(userStore.userInfo.id, form).then((res) => {
          if (res.code == 200) {
            userStore.setUserInfo(res.data);
            ElMessage.success(res.Message);
          } else {
            ElMessage.error(res.Message);
          }
        })
        emit('close')

      }
    }
  });
};
</script>

<style scoped lang="scss">
.border {
  border-bottom: 1px solid #e4e7ed;
}
.user-info {
  width: 60px;
  height: 36px;
  border: none !important;
  outline: none !important;
  &:hover {
    background-color: #eaebeb;
  }
}

.icon-user {
  font-size: 24px;
}

.context-menu {
  position: absolute;
  background: white;
  border: 1px solid #ccc;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  z-index: 999;
}

.menu-item {
  padding: 8px 12px;
  cursor: pointer;
  font-size: 14px;

  &:hover {
    background-color: #f3f3f5;
    color: #409eff;
  }
}

.menu-item.disabled {
  color: #ccc;
  cursor: not-allowed;
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
