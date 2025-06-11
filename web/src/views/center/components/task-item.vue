<template>
  <div class="card position-relative" @click="selectItem">
    <el-button v-if="isSelf && userStore.userInfo.role==2" plain :type="bgC" class="judge">{{ judgeText}}</el-button>
    <div style="padding: 12px">
      <div class="flex flex-justify-between">
        <div class="title ellips-one">{{ item.post }} <span class="ml-4px ">{{ item.place }}</span></div>
        <div class="price">{{ item.salary }}</div>
      </div>
      <div class="desc ellips-one mt-10px" v-html="item.description"></div>
    </div>
    <div class="flex flex-items-center info">
      <img
        :src="item.enterprise_info.avatar"
        alt=""
        style="width: 40px; height: 40px; border-radius: 50%"
      />
      <span class="ml-2 flex-1 ellips-one">{{ item.enterprise_info.name }}</span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineComponent, reactive,computed  } from "vue";
import { useUserStore } from "@/store/module/user";
const userStore = useUserStore();

const emit = defineEmits(["selectItem"]);

const props = defineProps({
  item: Object,
});
const isSelf = computed(() => {
  return userStore.getUserInfo?.id == props.item?.enterprise_info?.id;
});
const judgeText = computed(() => {
  return props.item.state==0?'待审核':props.item.state==1?"已发布":props.item.state==-1?'已下架':'被驳回';
});
const bgC = computed(() => {
  return props.item.state==0?'primary':props.item.state==1?"success":props.item.state==-1?'warning':'danger';
});
const selectItem = ()=>{
emit('selectItem',props.item)
}
</script>

<style scoped>
.card {
  overflow: hidden;
  margin-bottom: 20px;
  margin-right: 10px;
  cursor: pointer;
  background-color: #fff;

  border-radius: 10px;
  width: 360px;

  &:hover .title,
  &:hover .nickname {
    color: #409eff;
  }

  &:hover {
    box-shadow: 0 16px 40px 0 rgba(138, 149, 158, 0.2);
  }

  .title {
    color: #222;
    line-height: 22px;
    font-weight: 500;
    span{
        font-size: 13px;
        color: #666;
    }
  }

  .price {
    color: #fe574a;
    line-height: 22px;
    font-weight: 500;
  }

  .desc {
    color: #666;
    font-size: 14px;
  }

  .ellips-one {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .info {
    padding: 12px;
    margin-top: 10px;
    background: linear-gradient(90deg, #f5fcfc 0, #fcfbfa 100%);
  }
}
.judge{

  width: 110px;
  position: absolute;
  transform: rotate(40deg);
  top: 20px;
  right: 0;
}
</style>
