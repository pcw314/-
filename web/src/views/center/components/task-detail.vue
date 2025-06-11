<template>
  <div class="task-detail flex flex-row flex-justify-center mt-10px">
    <div
      class="p-20px position-relative"
      style="
        background-color: #fff;
        width: 1200px;
        border-top-right-radius: 10px;
        border-top-left-radius: 10px;
      "
    >
      <el-button
        plain
        v-if="isSelf && userStore.userInfo.role == 2"
        :type="bgC"
        class="judge"
        >{{ judgeText }}</el-button
      >

      <div class="info flex flex-items-center flex-justify-between">
        <el-icon class="close" @click="close"><Close /></el-icon>
        <div class="flex flex-1 flex-items-center">
          <img
            :src="item!.enterprise_info.avatar"
            alt=""
            style="width: 80px; height: 80px; border-radius: 50%"
          />
          <div class="ml-2 ellips-one">
            {{ item!.post }}
            <span class="ml-4px" style="font-size: 14px; color: #999">{{
              item!.enterprise_info.name
            }}</span>
          </div>
        </div>
        <span
          :class="isSelf ? 'mr-80px' : ''"
          style="color: #999; font-size: 14px"
          >{{ new Date(item!.created_at).toLocaleString() }}</span
        >
      </div>
      <div class="block">
        <div class="block-item">{{ item!.salary }}</div>
        <div class="block-item">{{ item!.working_time }}</div>
        <div class="block-item">{{ item!.contact_number }}</div>
        <div class="block-item">{{ item!.place }}</div>
      </div>
      <div class="px-10px py-10px mt-20px">
        <span>岗位描述：</span>
        <div class="text-align-justify mt-10px desc" v-html="desc"></div>
      </div>
      <div class="px-10px py-10px">
        <span>任务详情：</span>
        <div class="text-align-justify mt-10px desc" v-html="require"></div>
      </div>
      <div class="flex flex-item-center mt-10px  "  v-if="isSelf && userStore.userInfo.role == 2">
        驳回原因:<el-tag type="danger" plain class="ml-10px">{{ item!.reply }}</el-tag>
      </div>
      <div
        class="flex flex-justify-end pb-10px mr-10px"
        v-if="!isSelf && userStore.userInfo?.role == 1"
      >
        <el-button type="primary" :icon="Star" @click="collectOrNo">{{
          item!.is_collect == 1 ? "取消收藏" : "收藏"
        }}</el-button>
        <el-button type="success" :icon=" Message" @click="goToChat">立马沟通</el-button>
        <el-button type="danger" :icon=" Finished" @click="showComplain">举报</el-button>
      </div>
      <div
        class="flex flex-justify-end pb-10px mr-10px"
        v-if="isSelf && userStore.userInfo.role == 2"
      >
        <el-button type="warning" @click="editTask">修改</el-button>
        <el-button type="danger" @click="getOffTaskApi">下架</el-button>
        <el-button type="danger" @click="getOffTaskApi">上架</el-button>
        <el-button
          style="background-color: #cf0c0c; color: #fff"
          @click="deleteApi"
          >删除</el-button
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineComponent, reactive, computed, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { useUserStore } from "@/store/module/user";
import { createConversation, deleteTaskById, getOffTask } from "@/api/center";
import { Message, Star,Finished } from "@element-plus/icons-vue";
import { cancelCollectTaskById, collectTaskById } from "@/api/common";
const emit = defineEmits(["editTask",'showComplian','goToChat']);
const userStore = useUserStore();

const props = defineProps({
  item: Object,
  taskId: String,
});
const showComplain = ()=>{
  emit('showComplian',props.item!.id)
}
const desc = computed(() => props.item!.description.replace(/\n/g, "<br>"));
const require = computed(() => props.item!.require.replace(/\n/g, "<br>"));
const isSelf = computed(() => {
  return (
    userStore.getUserInfo?.id == props.item!.enterprise_info.id &&
    userStore.userInfo.role == 2
  );
});
const judgeText = computed(() => {
  return props.item!.state == 0
    ? "待审核"
    : props.item!.state == 1
      ? "已发布"
      : props.item!.state == -1
        ? "已下架"
        : "被驳回";
});
const bgC = computed(() => {
  return props.item!.state == 0
    ? "primary"
    : props.item!.state == 1
      ? "success"
      : props.item!.state == -1
        ? "warning"
        : "danger";
});
const editTask = () => {
  emit("editTask", props.item!.id as number);
};
const close = () => {
  let e = localStorage.getItem('activeTab')==2?2:localStorage.getItem('activeTab')==4?4:1
  let isFresh = e==2
  emit("close", e,isFresh);
};
const deleteApi = () => {
  deleteTaskById(props.item!.id as number, "").then((res) => {
    if (res.code == 200) {
      close();
    }
  });
};
const getOffTaskApi = () => {
  getOffTask(props.item!.id as number).then((res) => {
    if (res.code == 200) {
      close();
    }
  });
};
const collectOrNo = () => {
  let funcName = props.item!.is_collect
    ?  cancelCollectTaskById
    : collectTaskById;
  funcName(props.item!.id as number).then((res) => {
    props.item!.is_collect = !props.item!.is_collect;
    ElMessage.success(res.Message);
  });
};

const goToChat = ()=>{
  createConversation({receiver_id:props.item?.enterprise_info.user_id,job_id:props.item?.id}).then(res=>{
    if(res.code==200){
      emit('goTochat',res.data,)

    }
  })
}
</script>

<style scoped>
.task-detail {
  position: fixed;
  top: 60px;
  left: 50%;
  transform: translateX(-50%);
  background-color: #d8e2f5;
  padding-top: 10px;
  overflow-y: scroll;
}

.task-detail {
  overflow: hidden;
}
.desc {
  color: rgb(190, 186, 186);
  font-size: 14px;
}
.block {
  margin-top: 20px;
  display: flex;
  align-content: center;
  .block-item {
    border-radius: 2px;
    margin-left: 10px;
    padding: 6px 8px;
    background-color: #c6c8ca;
    color: #fff;
    transition: all 0.6s ease;
    &:hover {
      cursor: pointer;
      background-color: #409eff;
      color: #fff;
    }
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
.judge {
  width: 140px;
  position: absolute;
  transform: rotate(40deg);
  top: 40px;
  right: 0;
}
</style>
