<template>
  <div class="flex-1 card flex justify-between mt-60px position-relative">
    <el-icon v-if="taskId" class="close" @click="close"><Close /></el-icon>
    <img class="back-img" src="@/static/img/post-task.svg" alt="" />
    <div class="flex-1 pl-15px">
      <el-form :model="form" label-width="auto" style="max-width: 600px">
        <el-form-item required label="岗位类型">
          <el-radio-group v-model="form.type">
            <el-radio :value="1">全职</el-radio>
            <el-radio :value="0">兼职</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item required label="联系人">
          <el-input type v-model="form.contact_name" />
        </el-form-item>
        <el-form-item required label="发布校区">
          <el-select
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

        <el-form-item required label="联系电话">
          <el-input type v-model="form.contact_number" />
        </el-form-item>
        <el-form-item required label="工作地点">
          <el-input type v-model="form.place" />
        </el-form-item>
        <el-form-item required label="工作时间">
          <el-input type v-model="form.working_time" />
        </el-form-item>
        <el-form-item required label="岗位名称">
          <el-input v-model="form.post" />
        </el-form-item>
        <el-form-item required label="岗位薪资">
          <el-input v-model="form.salary" />
        </el-form-item>
        <el-form-item required label="岗位描述">
          <el-input type="textarea" v-model="form.description" :rows="3" />
        </el-form-item>
        <el-form-item required label="任职要求">
          <el-input type="textarea" v-model="form.require" :rows="3" />
        </el-form-item>
        <div class="flex-1 flex flex-items-center flex-justify-center">
          <el-button
            type="primary"
            size="large"
            v-if="!taskId"
            @click="saveTask"
            >保存</el-button
          >
          <el-button    v-if="taskId" type="warning" size="large" @click="edit">修改</el-button>
          <el-button   v-if="taskId" type="danger" size="large" @click="deleteApi"
            >下架</el-button
          >
          <el-button
          v-if="taskId" 

            style="background-color: #cf0c0c; color: #fff;"
            size="large"
          >删除</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineComponent, reactive, onMounted, ref } from "vue";
import { getSchoolList } from "@/api/common";
import {
  postTask,
  getTaskById,
  editTaskById,
  deleteTaskById,
  getOffTask,
} from "@/api/center";

import { ElMessage } from "element-plus";

const props = defineProps({
  taskId: String,
});
const emit = defineEmits(["postAfter", "close"]);
const close = () => {
  emit("close");
};
let activeTab = ref(localStorage.getItem('activeTab'))
let form = reactive({
  post: "",
  salary: "",
  contact_name: "",
  contact_number: "",
  description: "",
  require: "",
  type: 0,
  place: "",
  working_time: "",
  school: "",
});
const allSchoolList = reactive([]);
onMounted(() => {
  getSchoolList(0).then((res) => {
    if (res.code !== 200) {
      ElMessage.success(res.Message ? res.Message : "数据异常");
    } else {
      allSchoolList.push(...res.data.list);
    }
  });
});
const saveTask = () => {
  postTask(form).then((res) => {
    if (res.code == 200) {
      emit("postAfter", 1);
      localStorage.setItem("activeTab", 1);
    }
    localStorage.setItem("activeTab", 3);
    ElMessage.success(res.Message ? res.Message : "数据异常");
  });
};
onMounted(() => {
  if (props.taskId) {
    getTaskByIdApi();
  }
});

const getTaskByIdApi = () => {
  console.log(props.taskId, "props.taskId");
  getTaskById(props.taskId).then((res) => {
    if (res.code == 200) {
      form = Object.assign(form, res.data);
    }
  });
};
const edit = () => {
  editTaskById(props.taskId, form).then((res) => {
    if (res.code == 200) {
      emit("postAfter", 1);
    }
    ElMessage.success(res.Message ? res.Message : "数据异常");
  });
};
const deleteApi = () => {
  deleteTaskById(props.item.id).then((res) => {
    if (res.code == 200) {
      emit("postAfter", 1);
    }
    ElMessage.success(res.Message ? res.Message : "数据异常");
  });
};
const getOffTaskApi = () => {
  getOffTask(props.item.id).then((res) => {
    if (res.code == 200) {
      emit("postAfter", 1);
    }
    ElMessage.success(res.Message ? res.Message : "数据异常");
  });
};
</script>

<style scoped>
.card {
  border-radius: 12px;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 16px 40px 0 rgba(138, 149, 158, 0.2);

  .back-img {
    height: 100%;
  }
}
.close {
  cursor: pointer;
  position: absolute;
  right: 14px;
  top: 10px;
  &:hover {
    color: #409eff;
  }
}
</style>
