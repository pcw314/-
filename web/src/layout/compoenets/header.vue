<template>
  <div @click="menuVisible = false">
    <div class="flex p-3 h-14 flex-items-center border">
      <div></div>
      <i
        v-if="configStore.getIsScrollMenue"
        class="iconfont icon-mianbaoxie"
        style="color: #333; cursor: pointer"
        @click="openClose"
      ></i>
      <div class="flex-1 ml-16px mr-10px">
        <breadcrumb></breadcrumb>
      </div>
      <!-- <div
        class="user-info flex flex-justify-center flex-items-center cursor-pointer"
        @click="changeScrollMenue"
      >
        <i class="iconfont font-size-6 icon-majesticons_2home"></i>
      </div> -->

      <el-dropdown>
        <div
          class="user-info flex flex-justify-center flex-items-center cursor-pointer"
        >
        <i v-if="!form.avatar" class="iconfont icon-user" style="cursor: pointer"></i>
        <img v-else  class="" :src="form.avatar" style="width: 40px;height: 40px;border-radius: 50%;"></img>

          
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <div class="flex align-center" @click="showStaffInfo=true">
                <i class="iconfont icon-user" style="font-size: 20px"></i>
                <span>修改资料</span>
              </div>
            </el-dropdown-item>
            <el-dropdown-item>
              <div class="flex align-center" @click="loginOut">
                <i class="iconfont icon-exit" style="font-size: 18px"></i>
                <span>退出登录</span>
              </div>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
  <el-drawer
      v-model="showStaffInfo"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <template #header="{ close }">
        <h4>修改资料</h4>
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
</template>

<script lang="ts" setup>
import {
  ElMessage,
  ElLoading
} from "element-plus";
import { useConfigStore } from "@/store/module/config";
import breadcrumb from "./breadcrumb.vue";
import { ref, watch,reactive } from "vue";
import { useUserStore } from "@/store/module/user";
import companyInfo from "@/components/company-info.vue";
import {
  editStaffById,
  getStaffInfo,
  uploadFile,
} from "@/api/common";
import { useRouter } from "vue-router";
import { onMounted } from "vue";
import { nextTick } from "vue";

let userStore = useUserStore();
const ruleFormRef = ref<FormInstance>();

const form = reactive({
  avatar:'' ,
  name: '',
  phone:'',
  username: "",
  password: "",
});
const configStore = useConfigStore();
const showStaffInfo = ref(false);
const fileList = ref<UploadUserFile[]>([]);
const uploadRef = ref<UploadInstance>();
const { currentRoute } = useRouter();
const toRouter = useRouter();
const menuVisible = ref(false);

const goToOverview = () => {
  toRouter.push("/home/overview");
};

const activeRoute = ref<string>("");

const openClose = () => {
  userStore.setIsCollopse(!userStore.getIsCollopse);
};

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
watch(
  () => currentRoute.value,
  (route: any) => {
    if (route.path.startsWith("/redirect/")) {
      return;
    }
    activeRoute.value = route.path;
  },
  {
    immediate: true,
  }
);

const loginOut = () => {
  userStore.loginOut();
};
const confirmClick = async()=>{
    await ruleFormRef.value!.validate((valid) => {

    if (valid) {
      editStaffById(userStore.userInfo.id , form).then((res) => {
          if (res.code === 200) {
            showStaffInfo.value =false
            userStore.setUserInfo(res.data);
          }
          ElMessage.success(res.Message);
        });
    }
  });
    
}
onMounted(()=>{
 let loadingRef = ElLoading.service({
    lock: true,
    text: "Loading",
    background: "rgba(0, 0, 0, 0.7)",
  });
  getStaffInfo().then((res) => {
    if (res.code == 200) {
      form.avatar= res.data.avatar
  form.name=res.data.name
  form.phone=res.data.phone
      userStore.setUserInfo(res.data)
    } else {
      ElMessage.error(res.message)
    }
    loadingRef.close();

  }).catch(()=>{
    loadingRef.close();

  })
  form.avatar= userStore.userInfo?.avatar
  form.name=userStore.userInfo?.name
  form.phone=userStore.userInfo?.phone
})
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