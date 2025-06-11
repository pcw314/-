<template>
  <div class="contain flex flex-col">
    <div class="nav flex flex-justify-center p-10px">
      <div
        class="flex flex-row flex-items-center flex-justify-between"
        style="width: 1200px"
      >
        <div class="flex flex-row flex-1 flex-items-center">
          <div class="item flex flex-items-center" @click="showCity = true">
            <el-icon size="26" color="#409eff">
              <Location />
            </el-icon>
            <span class="active-color">{{ schoolName }}</span>
            <span>[切换]</span>
          </div>
          <div
            class="item"
            :class="{ active: activeTab == 1 }"
            @click="change(1)"
          >
            任务大厅
          </div>
          <div
            class="item"
            v-if="userStore.userInfo.role === 1"
            :class="{ active: activeTab == 2 }"
            @click="change(2)"
          >
            我的收藏
          </div>
          <div
            class="item"
            v-if="userStore.userInfo.role === 2"
            :class="{ active: activeTab == 3 }"
            @click="change(3)"
          >
            发布任务
          </div>
          <div
            class="item"
            v-if="userStore.userInfo.role === 2"
            :class="{ active: activeTab == 4 }"
            @click="change(4)"
          >
            任务管理
          </div>
          <div
            class="item"
            :class="{ active: activeTab == 5 }"
            @click="change(5)"
          >
            举报管理
          </div>
        </div>
        <div
          class="flex flex-row flex-items-center mr-20px"
          style="min-width: 180px"
        >
          <el-badge
            :value="unreadNumToal"
            :max="99"
            :show-zero="false"
            class="item"
            :offset="[-10, 0]"
          >
            <div
              class="item flex flex-items-center"
              :class="{ active: showChat }"
              @click="showChatEvent"
            >
              消息
            </div>
          </el-badge>

          <el-dropdown style="border: none; outline: none">
            <div class="flex flex-items-center item">
              <img
                class="ml-20px item"
                :src="userStore.userInfo.avatar"
                alt=""
                style="width: 40px; height: 40px; border-radius: 50%"
              />
              <div class="nickname">{{ userStore.userInfo.name }}</div>
            </div>

            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="logOut">退出登录</el-dropdown-item>
                <el-dropdown-item @click="showEditInfo=true">修改资料</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </div>
    <div
      class="flex flex-justify-center mt-20px mb-40px"
      v-if="activeTab == 1 && !showCity && !showEdit"
    >
      <div style="min-width: 1200px" class="flex flex-justify-center">
        <el-input
          v-model="pageSize.search"
          placeholder="搜索岗位名称，发布商户"
          clearable
          size="large"
          style="width: 664px"
        >
          <template #append>
            <el-button type="primary" :icon="Search" @click="search" />
          </template>
        </el-input>
      </div>
    </div>
    <!-- 任务列表 -->
    <el-scrollbar
      @scroll="scrollOne"
      v-if="
        (activeTab != 5 &&
          !showChat &&
          !showDetail &&
          !showEdit &&
          activeTab != 3 &&
          !showCity &&
          activeTab != 2) ||
        activeTab == 4
      "
    >
      <div
        class="flex-1 flex flex-row flex-justify-center mt-10px"
        style="min-width: 1200px"
      >
        <div class="flex flex-row flex-wrap" style="width: 1200px">
          <task-item
            :item="item"
            v-for="item in taskList"
            :key="item.id"
            @selectItem="showTaskDetail"
          ></task-item>
        </div>
      </div>
    </el-scrollbar>
    <el-scrollbar @scroll="scrollOne" v-if="activeTab == 2 && !showChat">
      <div
        class="flex-1 flex flex-row flex-justify-center mt-10px"
        style="min-width: 1200px"
      >
        <div class="flex flex-row flex-wrap" style="width: 1200px">
          <task-item
            :item="item"
            v-for="item in taskList"
            :key="item.id"
            @selectItem="showTaskDetail"
          ></task-item>
        </div>
      </div>
    </el-scrollbar>
    <div
      v-if="(activeTab == 3 || showEdit) && !showChat"
      class="flex flex-row pt-20px pb-20px flex-justify-center"
    >
      <div class="flex flex-row flex-wrap" style="width: 1200px">
        <post-task
          @post-after="change"
          :taskId="editTaskId"
          @close="closeEdit"
        ></post-task>
      </div>
    </div>
    <el-scrollbar
      class="flex-1"
      v-if="activeTab == 5 && !showChat"
      style="background-color:"
    >
      <div class="flex flex-row flex-justify-center pt-20px">
        <div class="flex flex-row flex-wrap" style="width: 1200px">
          <div
            class="com-plain p-2 position-relative"
            v-for="(item,index) in complainList"
            :key="item.id"
          >
          <span class="position-absolute" style="right: 10px;">
            <el-tag type="primary" v-if="item.state==0"> 待审核</el-tag>
            <el-tag type="danger" v-if="item.state==-1">被驳回</el-tag>
            <el-tag type="success" v-if="item.state==1">已通过</el-tag>
          </span>
            <div  class="flex flex-items-center mb-10px" v-if="item.job_id==0">
              <span style="color:gray;font-size: 14px;" class="mr-4px">被举报用户:</span>
              <img v-if="item.avatar" :src="item.avatar" alt=""  style="width: 40px;height: 40px;border-radius: 50%;"/>
              <span class="ml-4px">{{ item.name }}</span>
            </div>
            <div class="mb-10px
            " v-else>
            <span style="color:gray;font-size: 14px;">被举报岗位:</span>
              {{ item.name }}
            </div>
            <div v-if="item.state==-1">
              <span style="color:gray;font-size: 14px;">驳回原因：</span>
              <span>{{ item.replay }}</span>
            </div>
            <div>
              <span style="color:gray;font-size: 14px;">举报原因：</span>
              <span>{{ item.reason }}</span>
            </div>
            <div class="mt-10px">
              <span style="color:gray;font-size: 14px;">举报图片：</span>
              <el-scrollbar class="mt-4px" style="height: 120px">
                <template class="flex flex-row flex-wrap">
                  <div v-for="iitem in item.image.split(',')">
                    <img
                      class="mr-10px mb-10px"
                      :src="iitem"
                      alt=""
                      style="width: 100px; height: 100px"
                    />
                  </div>
                </template>
              </el-scrollbar>
            </div>
            <div v-if="item.state===0" class="pb-4px">
              <el-button type="danger" size="small" @click="deleteUserCom(index,item.id)">取消举报</el-button>

            </div>
          </div>
        </div>
      </div>
    </el-scrollbar>
    <!-- 消息列表 -->
    <chat
      :chatId="chatId"
      v-if="showChat"
      ref="chatRef"
      @readNum="readNum"
      @setUnread="setUnread"
      @showComplian="showDialogCom"
    ></chat>
    <!-- 任务详情 -->
    <task-detail
      v-if="showDetail && selectItem && !showChat"
      :item="selectItem"
      @edit-task="editTask"
      @close="close"
      @showComplian="showDialogCom"
      @goTochat="goTochat"
    ></task-detail>
  </div>
  <choise-city
    v-if="showCity"
    @selectSchool="selectSchool"
    @close="showCity = false"
  ></choise-city>

  <el-dialog v-model="showComplain" title="举报" width="500" destroy-on-close>
    <el-form :model="form">
      <el-form-item label="图片上传">
        <el-upload
          v-model:file-list="fileList"
          list-type="picture-card"
          multiple
          accept=".jpeg,.png"
          :on-remove="handleRemove"
          :auto-upload="false"
        >
          <el-icon>
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="原因">
        <el-input type="textarea" v-model="form.reason" :row="4"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="showComplain = false">取消</el-button>
        <el-button type="primary" @click="comfirmComplain"> 确定 </el-button>
      </div>
    </template>
  </el-dialog>
  <InfoEdit v-if="showEditInfo" @close="showEditInfo=false"></InfoEdit>
</template>

<script lang="ts" setup>
import taskItem from "./components/task-item.vue";
import {
  ElLoading,
  ElMessage,
  UploadProps,
  UploadUserFile,
} from "element-plus";

import { Search, Message, Star } from "@element-plus/icons-vue";
import { defineComponent, reactive, ref, onMounted } from "vue";
import choiseCity from "./components/choise-city.vue";
import chat from "./components/chat.vue";
import PostTask from "./components/post-task.vue";
import taskDetail from "./components/task-detail.vue";
import { useUserStore } from "@/store/module/user";
import {
  getCollectList,
  getCompanyInfo,
  getUserInfoById,
  uploadFile,
} from "@/api/common";
import {
  getTaskList,
  getSelfPost,
  complainJob,
  getChatList,
  complainUser,
  getUserComplainList,
  deletComplainUser
} from "@/api/center";
import { useUserSo } from "@/store/module/socket";
import InfoEdit from "./components/info-edit.vue";
const form = reactive({
  image: "",
  reason: "",
});
const showEditInfo =ref(false)
const complainType = ref<string>(null);
const showChat = ref(false);
const complainId = ref(null);
const showDialogCom = (id, type = "") => {

  showComplain.value = true;
  complainId.value = id;
  complainType.value = type;

};
const chatId = ref(-1);
const fileList = ref<UploadUserFile[]>([]);
const showComplain = ref(false);
const showDetail = ref(false);
const taskDetails = localStorage.getItem("taskDetail")
  ? localStorage.getItem("taskDetail")
  : null;
const selectItem = ref(taskDetails);
const userStore = useUserStore();
const schoolName = ref(userStore.userInfo.school_name);
const schoolId = ref(userStore.userInfo.school_id);
const scrollOne = (e) => {};
const tab = localStorage.getItem("activeTab")
  ? localStorage.getItem("activeTab")
  : 1;
const keywords = ref("");
const activeTab = ref(tab);
const taskList = reactive([]);
const editTaskId = ref("");
const showEdit = ref(false);
let loadingRef = ref(null);
const complainList = ref<any[]>([]);
const pageSize = reactive({
  page: 1,
  size: 50,
  search: "",
});
const showCity = ref(false);
const change = (e) => {

  complainList.value = [];
  showChat.value = false;
  pageSize.search = "";
  pageSize.size = 99;
  pageSize.page = 1;
  activeTab.value = e;
  localStorage.setItem("activeTab", e);
  editTaskId.value = "";
  showEdit.value = false;
  showDetail.value = false;
  taskList.length = 0;

  if (e == 1) {
    getTaskListApi();
  } else if (e == 4) {
    getSelfPostApi();
  } else if (e == 2) {
    getCollect();
  } else if (e == 5) {
    getComplain();
  }else{
    loadingRef.value.close();
  }
};
const readNum = (num) => {
  unreadNumToal.value = unreadNumToal.value - num;
};
const setUnread = (num) => {
  unreadNumToal.value = num;
};
const getCollect = () => {
  getCollectList()
    .then((res) => {
      if (res.code == 200) {
        taskList.push(...res.data.list);
      }
      loadingRef.value.close();
      ElMessage.success(res.Message ? res.Message : "数据异常");
    })
    .catch((res) => {
      loadingRef.value.close();
    });
};
onMounted(() => {
  connect();
  loadingRef.value = ElLoading.service({
    lock: true,
    text: "Loading",
    background: "rgba(0, 0, 0, 0.7)",
  });
  let funcName =
    userStore.userInfo.role == 1 ? getUserInfoById : getCompanyInfo;
  funcName(0).then(async (res) => {
    if (res.code == 200) {
      userStore.setUserInfo(res.data);
    }
    let ab = localStorage.getItem("activeTab")
      ? localStorage.getItem("activeTab")
      : 1;
    change(ab);
  });
});
const getComplain = () => {
  getUserComplainList(pageSize)
    .then((res) => {
      if (res.code == 200) {
        complainList.value.push(...res.data.list);
      }
      loadingRef.value.close();
      ElMessage.success(res.Message ? res.Message : "数据异常");
    })
    .catch(() => {
      loadingRef.value.close();
    });
};
const search = () => {
  pageSize.page = 1;
  pageSize.size = 99;
  taskList.length = 0;
  getTaskListApi();
};
const getTaskListApi = () => {
  getTaskList(schoolId.value, pageSize)
    .then((res) => {
      if (res.code == 200) {
        taskList.push(...res.data.list);
      }
      loadingRef.value.close();
      ElMessage.success(res.Message ? res.Message : "数据异常");
    })
    .catch(() => {
      loadingRef.value.close();
    });
};
const getSelfPostApi = () => {
  taskList.length = 0;
  getSelfPost()
    .then((res) => {
      if (res.code == 200) {
        taskList.push(...res.data.list);
      }
      loadingRef.value.close();
    })
    .catch(() => {
      loadingRef.value.close();
    });
};
const showTaskDetail = (e) => {
  selectItem.value = e;
  localStorage.setItem("taskDetail", e);
  showDetail.value = true;
};
const editTask = (e) => {
  editTaskId.value = e + "";
  console.log("editTaskId.value ", editTaskId.value);
  showEdit.value = true;
  showDetail.value = false;
};
const closeEdit = () => {
  showEdit.value = false;
  editTaskId.value = "";
};
const close = (e = 1, isFresh = false) => {
  showDetail.value = false;
  if (isFresh) {
    change(e);
  } else {
    loadingRef.value.close();
  }
};
const logOut = () => {
  userStore.loginOut();
};
const selectSchool = (e) => {
  schoolName.value = e.name;
  schoolId.value = e.id;
  taskList.length = 0;
  showCity.value = false;
  getTaskListApi();
};
const unreadNumToal = ref(0);
const handleRemove: UploadProps["onRemove"] = (uploadFile, uploadFiles) => {
  console.log(uploadFile, uploadFiles);
};
const total = ref(0);
const comfirmComplain = () => {
  fileList.value.forEach((item) => {
    upload(item);
  });
};
const upload = async (item) => {
  let formData = new FormData();
  let img = item.raw;
  formData.append("files", img as Blob);
  uploadFile(formData)
    .then((res) => {
      if (res.code === 200) {
        total.value = total.value + 1;
        form.image = form.image + res.data[0].path + ",";
        if (fileList.value.length === total.value) {
          if (!complainType.value) {
            complainJob(complainId.value, form).then((res) => {
              if (res.code == 200) {
                showComplain.value = false;
                ElMessage.success(res.Message);
              } else {
                ElMessage.error(res.Message);
              }
            });
          } else {
            complainUser(complainId.value, form).then((res) => {
              if (res.code == 200) {
                showComplain.value = false;
                ElMessage.success(res.Message);
              } else {
                ElMessage.error(res.Message);
              }
            });
          }
        }
        ElMessage.success("上传成功");
      } else {
        ElMessage.success("上传失败");
      }
    })
    .catch((res) => {
      ElMessage.success("上传失败");
    });
};
const showChatEvent = async () => {
  activeTab.value = -1;
  showChat.value = true;
    showChat.value = true;
};
const goTochat = async (id) => {
  chatId.value = id;
  activeTab.value = -1;

getChatListApi()
setTimeout(() => {
  showChat.value = true;
}, 300);
};
const chatRef = ref<any>(null);
const getChatListApi = () => {
  getChatList().then((res) => {
    if (res.code == 200) {
      if (res.data) {
        let defaultId = res.data && res.data.length ? res.data[0].id : -1;
        res.data.forEach((item) => {
          unreadNumToal.value = unreadNumToal.value + item.unread_num;
        });
        localStorage.setItem("chatList", JSON.stringify(res.data));
        localStorage.setItem("defaultId", defaultId);
      }
    }
  });
};
const connect = () => {
  let url = "http://127.0.0.1:8080/message/ws?token=" + userStore.token;
  let socket = new WebSocket(url);
  socket.onopen = function (event) {
    getChatListApi();
  };

  // 收到消息时触发
  socket.onmessage = function (event) {
    // chatRef.value?.emitChangeChatList();
    chatRef.value?.emitPutMessage(event.data);
  };

  // 连接关闭时触发
  socket.onclose = function (event) {
    console.log("Connection closed");
  };
  // 发生错误时触发
  socket.onerror = function (error) {
    connect();
    console.log(error);
  };
};
const deleteUserCom = (index,id)=>{
  deletComplainUser(id).then(res=>{
    if(res.code===200){
      ElMessage.success(res.Message)
      complainList.value.splice(index,1)
    }else{
      ElMessage.error(res.Message)

    }
  })
}
</script>

<style scoped>
* {
  box-sizing: border-box !important;
}

body {
  background-color: #4a4c50;
}

.contain {
  height: 100vh;
  background-color: #d8e2f5;
  overflow-y: scroll;

  .nav {
    width: 100%;
    background-color: black;
    /* padding: 16px; */
    color: #fff;

    .item {
      cursor: pointer;
      margin-right: 20px;

      &:hover {
        color: #409eff;
      }
    }

    .active-color {
      color: #409eff;
    }
  }
}

.task,
.task-detail {
  background-color: #d8e2f5;
  padding-top: 10px;
  overflow-y: scroll;
}

.task-detail {
  overflow: hidden;
}

.info {
  border-top-right-radius: 10px;
  border-top-left-radius: 10px;
  padding: 12px;
  border-bottom: 1px solid #d8e2f5;
  background: linear-gradient(90deg, #f5fcfc 0, #fcfbfa 100%);
}

.ellips-one {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* .contain::before {
  z-index: -1;
  content: "";
  position: absolute;
  height: 614px;
  top: 60px;
  left: 0;
  right: 0;
  background: linear-gradient(180deg, #00bebd 0, #88fffe 50%, #fff 100%);;
} */
.active {
  color: #409eff;
}
:deep(.el-scrollbar__view) {
  display: flex !important;
  flex-direction: column;
  flex: 1;
}
.com-plain {
  border-radius: 6px;
  background-color: #fff;
  width: 360px;
  min-height: 100px;
  max-height: 260px;
  margin-right: 10px;
  margin-bottom: 10px;
  &:last-child {
    margin-right: auto !important;
  }
}
</style>
