<template>
  <el-dialog
    v-model="showDialog"
    width="500"
    :show-close="false"
    :close-on-click-modal="false"
  >
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
      <el-form-item label="性别" required prop="sex">
        <el-radio-group v-model="form.sex">
          <el-radio :value=1 size="large">男</el-radio>
          <el-radio :value=2 size="large">女</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="电话：" required prop="phone">
        <el-input v-model="form.phone" placeholder="请输入电话"> </el-input>
      </el-form-item>
      <el-form-item label="年龄：" required prop="age">
        <el-input type="number" v-model="form.age" placeholder="输入年龄：">
        </el-input>
      </el-form-item>
      <el-form-item label="专业：" required prop="major">
        <el-input v-model="form.major" placeholder="输入专业名称"> </el-input>
      </el-form-item>
      <template v-if="isAdd">
        <el-form-item label="账号" required prop="username">
          <el-input v-model="form.username" placeholder="设置账号"> </el-input>
        </el-form-item>
        <el-form-item label="密码" required prop="password">
          <el-input v-model="form.password" placeholder="设置密码"> </el-input>
        </el-form-item>
      </template>
      <el-form-item label="校区" required class="" prop="school_id">
        <el-select
          transfer="true"
          :popper-append-to-body="false"
          style="cursor: pointer; width: 100%"
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
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="editStudent">
          {{ isAdd ? "保存" : "确认修改" }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import type { ComponentSize, FormInstance, FormRules } from "element-plus";
import { getSchoolList, uploadFile, adminCreateStudent,editStudentInfo  } from "@/api/common";
import {
  ElMessage,
  genFileId,
  UploadInstance,
  UploadProps,
  UploadRawFile,
  UploadUserFile,
} from "element-plus";
import { defineComponent, onMounted, reactive, ref } from "vue";
const emit = defineEmits(["close"]);
const props = defineProps({
  isAdd: Boolean,
  item: Object,
});
const ruleFormRef = ref<FormInstance>();
const form = reactive({
  avatar: "",
  name: "",
  phone: "",
  sex: 1,
  school_id: "",
  major: "",
  age: null,
  username: "",
  password: "",
});
const showDialog = ref(true);
const fileList = ref<UploadUserFile[]>([]);
const uploadRef = ref<UploadInstance>();

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
const allSchoolList = reactive<any[]>([]);
onMounted(() => {
  if (!props.isAdd) {
    form.age = props.item.age;
    form.phone = props.item.phone;
    form.major = props.item.major;
    form.school_id = props.item.school_id;
    form.name = props.item.name;
    form.avatar = props.item.avatar;
    form.sex = props.item.sex;
    console.log(form, "form");
  }
  getSchoolList(0).then((res) => {
    if (res.code !== 200) {
      ElMessage.success(res.Message ? res.Message : "数据异常");
    } else {
      allSchoolList.push(...res.data.list);
    }
  });
});
const cancel = () => {
  emit("close");
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
const editStudent = async () => {
  await ruleFormRef.value.validate((valid) => {
    form.age = form.age * 1;
    if (valid) {
      if (props.isAdd) {
        adminCreateStudent(form).then((res) => {
          if (res.code === 200) {
            emit("close");
          }
          ElMessage.success(res.Message);
        });
      } else {
        editStudentInfo(props.item.id, form).then((res) => {
          if (res.code === 200) {
            emit("close");
          }
          ElMessage.success(res.Message);
        });
      }
    }
  });
};
</script>

<style scoped>
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
