<template>
  <div class="flex flex-1 flex-col flex-items-center py-20px">
    <div class="flex flex-row" style="width: 1200px; height: 100%">
      <div
        class="message mr-20px flex flex-col"
        style="min-width: 400px; overflow-y: scroll; height: 800px"
      >
        <!-- <el-input
        clearable
          class="mb-20px"
          @change="searchContact"
          v-model="keyword"
          placeholder="搜索30天内的联系人,点击enter进行搜索"
          :suffix-icon="Search"
        /> -->
        <div
          @click="changeTalk(item)"
          class="item-message mb-10px position-relative"
          v-for="item in chatList"
          :key="item.id"
          :class="{ 'active-bg': ativeId == item.id }"
        >
          <div class="flex flex-items-center" v-if="item.is_top">
            <el-badge
              class="position-absolute"
              :value="item.unread_num"
              style="right: -4px; top: -2px"
              :show-zero="false"
            >
            </el-badge>
            <img src="@/static/img/log.svg" alt="" class="avatar" />

            <div class="flex flex-col flex-1">
              <div class="flex-row flex flex flex-items-center flex-1">
                <div
                  class="nickname ellips-one flex-nowrap"
                  style="color: #333"
                >
                  {{
                    userStore.userInfo.role === 1
                      ? item.contact_name
                      : item.student_info?.name
                  }}
                  <span class="font-sm" style="font-size: 14px; color: #ccc">
                    {{ item.post }}</span
                  >
                </div>
                <div class="time flex-1">{{ timeFormat(item.update_at) }}</div>
              </div>
              <div class="flex flex-items-center flex-justify-between mt-10px">
                <div class="ellips-one text">
                  {{ item.msg }}
                </div>
                <div>
                  <el-dropdown @command="handleCommand">
                    <el-icon size="18" class="icon-el">
                      <ChatDotRound />
                    </el-icon>
                    <template #dropdown>
                      <el-dropdown-menu>
                        <el-dropdown-item :command="item">
                          <el-icon> <Upload /> </el-icon
                          >取消置顶</el-dropdown-item
                        >
                        <el-dropdown-item command="del">
                          <el-icon> <Delete /> </el-icon
                          >删除回话</el-dropdown-item
                        >
                      </el-dropdown-menu>
                    </template>
                  </el-dropdown>
                </div>
              </div>
            </div>
          </div>
          <div class="flex flex-items-center flex-row" v-else>
            <el-badge
              class="position-absolute"
              :value="item.unread_num"
              style="right: -4px; top: -2px"
              :show-zero="false"
            >
            </el-badge>
            <img src="@/static/img/log.svg" alt="" class="avatar" />

            <div class="flex flex-col flex-1">
              <div class="flex-row flex flex flex-items-center flex-1">
                <div
                  class="nickname ellips-one flex-nowrap"
                  style="color: #333"
                >
                  {{
                    userStore.userInfo.role === 1
                      ? item.contact_name
                      : item.student_info?.name
                  }}
                  <span class="font-sm" style="font-size: 14px; color: #ccc">
                    {{ item.post }}</span
                  >
                </div>
                <div class="time flex-1">{{ timeFormat(item.update_at) }}</div>
              </div>
              <div class="flex flex-items-center flex-justify-between mt-10px">
                <div class="ellips-one text">
                  {{ item.msg }}
                </div>
                <div>
                  <el-dropdown @command="handleCommand">
                    <el-icon size="18" class="icon-el">
                      <ChatDotRound />
                    </el-icon>
                    <template #dropdown>
                      <el-dropdown-menu>
                        <el-dropdown-item :command="item">
                          <el-icon> <Upload /> </el-icon
                          >置顶回话</el-dropdown-item
                        >
                        <el-dropdown-item command="del">
                          <el-icon> <Delete /> </el-icon
                          >删除回话</el-dropdown-item
                        >
                      </el-dropdown-menu>
                    </template>
                  </el-dropdown>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="ativeId!=-1" class="chat-body flex-1 flex-col flex" style="max-height: 800px">
        <div class="top flex-row flex flex justify-between flex-items-center">
          <div
            class="nickname flex flex-items-center flex-nowrap"
            style="width: 530px"
          >
            <span class="ellips-one" style="max-width: 120px">
              {{
                userStore.userInfo.role === 1
                  ? activeItem?.contact_name
                  : activeItem?.student_info?.name
              }}</span
            >
            <span class="address flex-1 ellips-one">{{
              userStore.userInfo.role === 1
                ? activeItem?.enterprise_info?.name
                : activeItem?.student_info?.school_name
            }}</span>
          </div>
          <el-button type="danger" class="ml-30px" @click="showComplain">
            举报<el-icon class="el-icon--right">
              <Upload />
            </el-icon>
          </el-button>
        </div>
        <div
          class="flex-1"
          style="overflow-y: scroll"
          ref="chatRef"
          @scroll="loadMore"
        >
          <div class="chat-area pt-20px">
            <template v-for="item in chatRecord" :key="item.id">
              <div
                class="chat-item flex mb-20px"
                v-if="item.sender_id !== userStore.userInfo.user_id"
              >
                <img class="avatar mr-10px" src="@/static/img/log.svg" alt="" />
                <div class="contain p-10px text">
                  {{ item.content }}
                </div>
              </div>
              <div
                v-else
                class="chat-item flex flex-justify-end flex-1 mb-20px"
              >
                <div
                  class="contain px-10px py-2px text mr-10px"
                  style="background-color: #c6e2ff"
                >
                  {{ item.content }}
                </div>
                <img class="avatar" src="@/static/img/log.svg" alt="" />
              </div>
            </template>
          </div>
        </div>
        <div class="send-area">
          <el-scrollbar class="pt-10px" style="height: 78px">
            <textarea
              v-model="message"
              rows="4"
              style="width: 100%; border: none"
            ></textarea>
          </el-scrollbar>
          <div class="flex flex-justify-end">
            <el-button class="ml-8px" type="success" size="“small" @click="send"
              >发送</el-button
            >
          </div>
        </div>
      </div>
    
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  computed,
  defineComponent,
  nextTick,
  onMounted,
  reactive,
  ref,
  watch,
  watchEffect,
} from "vue";
import { Search } from "@element-plus/icons-vue";
import { useUserSo } from "@/store/module/socket";
import formatTime from "@/utils/time";
import {
  getChatList,
  getChatRecord,
  sendMessage,
  comfirmTop,
  cancelTop,
  deletChat
} from "@/api/center";
import { useUserStore } from "@/store/module/user";
const so = useUserSo();
const userStore = useUserStore();
const props = defineProps({
  chatId: Number,
});
const message = ref("");
const activeItem = ref();
const chatRecord = ref<any[]>([]);
const ativeId = ref(so.id);
const keyword = ref("");
const info = ref("");
const recordToatl = ref(0);
const parms = reactive({
  page: 1,
  size: 20,
});
const sendLoading = ref(false);
const previousHeight = ref(0);
const isLoadMore = ref(false);
const chatRef = ref(null);
const chatList = ref<any[]>([]);
const emit = defineEmits(["readNum", "setUnread", "showComplian"]);
const timeFormat = (value) => {
  let timeStamp = +new Date(value);
  return formatTime.getTime(timeStamp);
};
const isFirst = ref(true);
onMounted(() => {
  chatList.value = JSON.parse(localStorage.getItem("chatList") as string);
  if (props.chatId != -1) {
    ativeId.value = props.chatId as number;
  } else {
    ativeId.value = JSON.parse(
      localStorage.getItem("defaultId") as string
    ) as number;
  }
});
const changeTalk = (item) => {
  ativeId.value = item.id;
  item.unread_num = 0;
  isLoadMore.value = false;
  isFirst.value = true;
  emit("readNum", item.unread_num);
};
watchEffect(() => {
  let list = JSON.parse(JSON.stringify(localStorage.getItem("chatList")));
  if (list && list.length) {
    list = JSON.parse(list);
    list!.forEach((item) => {
      if (item.id == ativeId.value) {
        activeItem.value = item;
      }
    });
  }
});
watch(ativeId, () => {
  chatRecord.value = [];
  parms.page = 1;

  getChatRecordList();
});
const getChatRecordList = () => {
  previousHeight.value = chatRef.value.scrollHeight;
  console.log(activeItem.value, "activeItem.value");
  let userId = activeItem.value?.enterprise_info
    ? activeItem.value.enterprise_info.user_id
    : activeItem.value?.student_id;
  getChatRecord(userId, ativeId.value, parms).then((res) => {
    if (res.code == 200) {
      recordToatl.value = res.data.total;
      chatRecord.value = [...res.data.list, ...chatRecord.value];
    }
    if (!isLoadMore.value) {
      isLoadMore.value = false;
      isFirst.value = false;
      nextTick(() => {
        chatRef.value!.scrollTo({
          top: chatRef.value.scrollHeight,
          behavior: "smooth",
        });
      });
    } else {
      nextTick(() => {
        chatRef.value!.scrollTo({
          top: chatRef.value.scrollHeight - previousHeight.value,
          behavior: "smooth",
        });
      });
    }
  });
};
// const searchContact =(e)=>{
// keyword.value = e
// emitChangeChatList()
// }
const showComplain = () => {
  let id =
    userStore.userInfo.role == 1
      ? activeItem.value.enterprise_id
      : activeItem.value.student_id;
  emit("showComplian", id, "user");
};
const handleCommand = (command: any) => {
  if (command === "del") {
    deletChat(ativeId.value).then((res) => {
      if (res.code === 200) {
        emitChangeChatList();
      }
    });
  } else if (!command.is_top) {
    comfirmTop(command.id).then((res) => {
      if (res.code === 200) {
        emitChangeChatList();
      }
    });
  } else if (command.is_top) {
    cancelTop(command.id).then((res) => {
      if (res.code === 200) {
        emitChangeChatList();
      }
    });
  }
};
const send = () => {
  sendLoading.value = true;
  let data = {
    conv_id: ativeId.value,
    receiver_id:
      userStore.userInfo.role == 1
        ? activeItem.value.enterprise_id
        : activeItem.value.student_id,
    content: message.value,
  };
  chatRecord.value.push({
    content: message.value,
    state: -1,
    id: -1,
    sender_id: userStore.userInfo.user_id,
  });
  nextTick(() => {
    chatRef.value!.scrollTo({
      top: chatRef.value.scrollHeight,
      behavior: "smooth",
    });
  });
  sendMessage(data).then((res) => {
    if (res.code == 200) {
      chatRecord.value[chatRecord.value.length - 1].state = 1;
      chatRecord.value[chatRecord.value.length - 1].id = res.data;
    } else {
      chatRecord.value[chatRecord.value.length - 1].state = 0;
    }
    message.value = "";
  });
};

const loadMore = () => {
  let top = chatRef.value.scrollTop;
  if (top === 0 && !isFirst.value) {
    if (recordToatl.value > chatRecord.value.length) {
      parms.page = parms.page + 1;
      isLoadMore.value = true;
      getChatRecordList();
    }
  }
};
const emitChangeChatList = () => {

  getChatList({search:keyword.value}).then((res,) => {
    if (res.code == 200) {
      if (res.data) {
        chatList.value = res.data;
        let unread = 0;
        res.data.forEach((item) => {
          unread = unread + item.unread_num;
        });
        emit("setUnread", unread);
      }
    }
  });
};
const emitPutMessage = (val) => {
  let data = JSON.parse(val);
  if (ativeId.value === data.data.conv_id) {
    chatRecord.value.push(data.data);
    nextTick(() => {
      chatRef.value!.scrollTo({
        top: chatRef.value.scrollHeight,
        behavior: "smooth",
      });
    });
  } else {
    emitChangeChatList();
  }
};
defineExpose({ emitChangeChatList, emitPutMessage });
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.chat-item {
  .avatar {
    width: 56px;
    height: 56px;
    border-radius: 50%;
  }

  .contain {
    display: flex;
    flex-direction: column;
    justify-content: center;
    position: relative;
    max-width: 420px;
    vertical-align: top;
    word-break: break-all;
    border-radius: 8px;
    background-color: #f8f8f8;
    color: #333;
  }
}

.send-area {
  border-top: 1px solid #e6e8eb;
  position: relative;

  textarea {
    overflow: hidden;
    scrollbar-width: 0;
    resize: none;
    font-size: 16px;
    color: #333;
    font-weight: 500;
    outline: none;
  }
}

.ellips-one {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.chat-body {
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 12px;

  .top {
    border-bottom: 1px solid rgb(206, 202, 202);
    height: 48px;
    line-height: 48px;

    .nickname {
      color: #333;
      font-size: 20px;

      .address {
        font-size: 16px;
        color: #b3b3b3;
        margin-left: 12px;
      }
    }
  }
}

.message {
  border-radius: 12px;
  background-color: #fff;
  padding: 10px;

  .item-message {
    cursor: pointer;
    padding: 10px;

    &:hover {
      background-color: #f8f8f8;
    }

    .avatar {
      width: 56px;
      height: 56px;
      border-radius: 50%;
      margin-right: 12px;
    }

    .nickname {
      color: #b3b3b3;
      font-size: 18px;
    }

    .time {
      font-size: 13px;
      margin-left: 8px;
      color: #b3b3b3;
      text-align: end;
    }

    .text {
      font-size: 13px;
      width: 180px;
      color: rgb(153, 153, 153);
    }

    .icon-el {
      cursor: pointer;
      color: #b3b3b3;
      opacity: 0;

      &:hover {
        color: #409eff;
      }
    }

    &:hover .icon-el {
      opacity: 1;
    }
  }
}

.active-bg {
  background-color: #d9ecff !important;
}
</style>
