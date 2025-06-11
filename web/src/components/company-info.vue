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

      <el-form-item label="公司名称" required prop="name">
        <el-input v-model="form.name" placeholder="输入公司名称"> </el-input>
      </el-form-item>

      <el-form-item label="公司法人" required prop="lega_person">
        <el-input v-model="form.lega_person" placeholder="输入法人名称"> </el-input>
      </el-form-item>

      <el-form-item label="电话：" required prop="phone">
        <el-input v-model="form.phone" placeholder="请输入电话"> </el-input>
      </el-form-item>

      <template v-if="isAdd">
        <el-form-item label="账号" required prop="username">
          <el-input v-model="form.username" placeholder="设置账号"> </el-input>
        </el-form-item>
        <el-form-item label="密码" required prop="password">
          <el-input v-model="form.password" placeholder="设置密码"> </el-input>
        </el-form-item>
      </template>

      <el-form-item label="校区" required prop="school_id">
        <el-select
          transfer="true"
          :popper-append-to-body="false" 
          style="cursor: pointer; width: 100%"
          v-model="form.school_id"
          filterable               
          remote                   
          reserve-keyword         
          placeholder="输入您检索的校区" 
          
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
        <el-button type="primary" @click="editCompany">
          {{ isAdd ? "保存" : "确认修改" }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import type { FormInstance, FormRules } from "element-plus"; // 引入 Element Plus 表单类型
import { 
  uploadFile, // 假设这是正确的上传文件 API 函数名
  editCompanyInfo, 
  adminCreateCompany, 
  getSchoolList 
} from "@/api/common"; // 引入 API 请求函数
import {
  ElMessage,
  genFileId,
  UploadInstance,
  UploadProps,
  UploadRawFile,
  UploadUserFile,
} from "element-plus"; // 引入 Element Plus 相关类型和组件
import { defineComponent, onMounted, reactive, ref, defineEmits, defineProps } from "vue"; // 引入 Vue 相关 API

// 定义组件触发的事件
const emit = defineEmits(["close"]);
// 定义组件接收的属性
const props = defineProps({
  isAdd: Boolean, // 是否是新增模式
  item: Object,   // 编辑模式下传入的项数据
});

// 表单实例引用
const ruleFormRef = ref<FormInstance>();
// 表单数据模型
const form = reactive({
  avatar: "",       // 头像 URL 或路径
  name: "",         // 公司名称
  phone: "",        // 电话
  school_id: "",    // 校区 ID
  lega_person: "",  // 法人
  username: "",     // 账号 (仅新增时)
  password: "",     // 密码 (仅新增时)
});

// 控制对话框显示
const showDialog = ref(true);
// 上传文件列表
const fileList = ref<UploadUserFile[]>([]);
// 上传组件实例引用
const uploadRef = ref<UploadInstance>();
// 所有校区列表
const allSchoolList = reactive<any[]>([]);

/**
 * 文件状态改变时的钩子函数 (选择文件后)
 * @param rawFile 原始文件信息
 */
const onchange = (rawFile: UploadUserFile) => {
  // 检查文件类型
  if (rawFile.raw?.type !== "image/jpeg" && rawFile.raw?.type !== "image/png") {
    ElMessage.error("头像图片必须是 JPG 或 PNG 格式!");
    // 清除错误的文件选择
    fileList.value = fileList.value.filter(file => file.uid !== rawFile.uid); 
    return false;
  }
  // 生成本地预览 URL
  if(rawFile.raw) {
    form.avatar = URL.createObjectURL(rawFile.raw);
  }
  
  // 延迟执行上传操作 (注意: 这个延迟并非必需, 除非有特殊原因)
  setTimeout(() => {
    // 确保文件仍在列表中再上传 (防止用户快速更换文件)
    if (fileList.value.some(file => file.uid === rawFile.uid) && rawFile.status === 'ready') {
       upload();
    }
  }, 1000); 
};

/**
 * 上传文件之前的钩子 (主要用于文件校验)
 * @param rawFile 原始文件信息
 */
const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
  if (rawFile.type !== "image/jpeg" && rawFile.type !== "image/png") {
    ElMessage.error("头像图片必须是 JPG 或 PNG 格式!");
    return false;
  }
  // 可以在这里添加文件大小限制等校验
  // const isLt2M = rawFile.size / 1024 / 1024 < 2;
  // if (!isLt2M) {
  //   ElMessage.error('Avatar picture size can not exceed 2MB!');
  //   return false;
  // }
  return true;
};

/**
 * 文件超出个数限制时的钩子
 * @param files 新选择的文件列表
 */
const handleExceed: UploadProps["onExceed"] = (files) => {
  uploadRef.value!.clearFiles(); // 清除旧文件
  const file = files[0] as UploadRawFile;
  file.uid = genFileId(); // 为新文件生成唯一 ID
  uploadRef.value!.handleStart(file); // 手动开始处理新文件 (会触发 on-change)
};

// 组件挂载后执行
onMounted(() => {
  // 如果是编辑模式, 填充表单数据
  if (!props.isAdd && props.item) {
    form.phone = props.item.phone;
    form.school_id = props.item.school_id;
    form.name = props.item.name;
    form.lega_person = props.item.lega_person;
    form.avatar = props.item.avatar;
  }
  // 获取校区列表
  getSchoolList(0).then((res) => {
    if (res.code !== 200) {
      ElMessage.error(res.Message ? res.Message : "获取校区数据异常");
    } else {
      // 清空并填充校区列表 (避免重复添加)
      allSchoolList.splice(0, allSchoolList.length, ...res.data.list);
    }
  }).catch(error => {
      ElMessage.error("请求校区列表失败");
      console.error("Error fetching school list:", error);
  });
});

/**
 * 取消操作, 关闭对话框
 */
const cancel = () => {
  emit("close");
};

/**
 * 执行文件上传操作
 */
const upload = () => {
  if (fileList.value.length === 0 || !fileList.value[0].raw) {
      ElMessage.warning("请先选择要上传的头像文件");
      return;
  }
  
  let formData = new FormData();
  let img = fileList.value[0].raw;
  formData.append("files", img); // 'files' 是后端接收文件的字段名

  // 调用上传 API
  uploadFile(formData)
    .then((res) => {
      if (res.code === 200 && res.data && res.data.length > 0) {
        // 上传成功, 更新 form.avatar 为服务器返回的路径
        form.avatar = res.data[0].path;
        ElMessage.success("头像上传成功");
        // 上传成功后，可以考虑从 fileList 移除，避免重复上传
        // fileList.value = []; 
      } else {
        // 处理上传成功但返回数据格式不符或 code 不为 200 的情况
        ElMessage.error(res.Message || "头像上传失败");
        // 上传失败，可能需要重置头像预览为之前的状态或默认
        // if (!props.isAdd && props.item) {
        //   form.avatar = props.item.avatar; 
        // } else {
        //   form.avatar = ''; 
        // }
        // fileList.value = []; // 清空文件列表
      }
    })
    .catch((err) => {
      // 网络错误或其他异常
      ElMessage.error("头像上传请求失败");
      console.error("Upload error:", err);
       // 上传失败，可能需要重置头像预览
       // ... (同上)
       // fileList.value = []; // 清空文件列表
    });
};

/**
 * 提交表单 (新增或编辑)
 */
const editCompany = async () => {
  // 校验表单
  if (!ruleFormRef.value) return;
  await ruleFormRef.value.validate((valid) => {
    if (valid) {
      // 根据 isAdd 标志调用不同的 API
      const apiCall = props.isAdd
        ? adminCreateCompany(form)
        : editCompanyInfo(props.item!.id, form); // 在编辑时确保 props.item.id 存在

      apiCall.then((res) => {
        if (res.code === 200) {
          ElMessage.success(props.isAdd ? "保存成功" : "修改成功");
          emit("close"); // 操作成功后关闭对话框
        } else {
          // API 返回错误信息
          ElMessage.error(res.Message || (props.isAdd ? "保存失败" : "修改失败"));
        }
      }).catch(error => {
          // 网络或请求错误
          ElMessage.error("操作请求失败");
          console.error("Submit error:", error);
      });
    } else {
      ElMessage.warning("请检查表单填写是否完整且正确");
      return false; // 校验失败，阻止提交
    }
  });
};
</script>

<style scoped>
/* 移除 .form 样式, 因为它看起来是多余的并且可能导致布局问题 */
/* .form { ... } */ 

/* 头像上传器样式 */
.avatar-uploader {
  display: flex;
  justify-content: center;
}

.avatar-uploader .avatar {
  width: 140px;
  height: 140px;
  display: block;
  border: 1px dashed var(--el-border-color); /* 给图片也加上边框，保持一致性 */
  border-radius: 6px;
}

/* Element Plus 上传组件内部样式微调 (如果需要) */
:deep(.avatar-uploader .el-upload) {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

:deep(.avatar-uploader .el-upload:hover) {
  border-color: var(--el-color-primary);
}

/* 上传图标样式 */
.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 140px;  /* 保持和图片区域大小一致 */
  height: 140px; /* 保持和图片区域大小一致 */
  text-align: center;
  display: flex; /* 使用 flex 居中图标 */
  justify-content: center; /* 水平居中 */
  align-items: center;    /* 垂直居中 */
  border: 1px dashed var(--el-border-color); /* 给图标区域加上边框 */
  border-radius: 6px;
}
</style>