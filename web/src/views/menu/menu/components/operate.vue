<template>
    <div>
      <el-drawer
        size="450"
        v-model="isShowForm"
        :show-close="false"
        title="添加"
        @close="closeDrawer"
      >
        <template #header="{ close, titleId, titleClass }">
          <div style="border-bottom: 1px solid rgb(239, 239, 245); display: flex">
            <h4 :id="titleId" :class="titleClass">
              {{ type == "add" ? "新增" : "编辑" }}
            </h4>
          </div>
        </template>
        <div>
          <el-form :model="form" label-width="auto" label-position="top">
            <el-form-item label="菜单类型">
              <el-radio-group v-model="form.type">
                <el-radio :value="1">目录</el-radio>
                <el-radio :value="2">菜单</el-radio>
                <el-radio :value="3">模块</el-radio>
                <el-radio :value="4">接口</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="菜单类型">
              <el-radio-group v-model="form.method">
                <el-radio value="GET">GET</el-radio>
                <el-radio value="POST">POST</el-radio>
                <el-radio value="PUT">PUT</el-radio>
                <el-radio value="DELETE">DELETE</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="菜单名称">
              <el-input v-model="form.title" />
            </el-form-item>
            <el-form-item label="父级分类" prop="parent_id">
              <el-tree-select
                v-model="form.parent_id"
                placeholder="请选择"
                :props="{
                  label: 'title',
                  value: 'id',
                  children: 'children',
                }"
                :data="industryOptions"
                :render-after-expand="false"
                style="width: 240px"
                @change="selectChange"
                :check-strictly="true"
              />
            </el-form-item>
  
            <el-form-item label="iconify图标" prop="icon">
              <el-select
                v-model="form.icon"
                placeholder="请选择"
                size="large"
                style="width: 240px"
              >
                <el-option
                  v-for="item in iconfontList.glyphs"
                  :key="item.font_class"
                  :label="item.font_class"
                  :value="item.font_class"
                />
              </el-select>
              <div>
                选择的图标：<i
                  :class="`iconfont icon-${form.icon}`"
                  class="font-size-[30px]"
                ></i>
              </div>
            </el-form-item>
            <div class="h-5"></div>
  
            <el-form-item label="路由路径">
              <el-input v-model="form.path" />
            </el-form-item>
            <el-form-item label="排序">
              <el-input v-model="form.sort" />
            </el-form-item>
            <el-form-item label="隐藏菜单">
              <el-radio-group v-model="form.is_hide">
                <el-radio :value="0">显示</el-radio>
                <el-radio :value="1">隐藏</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-form>
        </div>
        <template #footer>
          <div class="flex justify-end items-center gap-4 p-4 bg-white">
            <el-button
              @click="closeForm"
              style="
                border-radius: 2px;
                height: 39px;
                line-height: 39px;
                padding: 0 20px;
              "
            >
              取消
            </el-button>
            <el-button
              type="primary"
              @click="onSubmit"
              style="
                border-radius: 2px;
                height: 39px;
                line-height: 39px;
                padding: 0 20px;
              "
            >
              确定
            </el-button>
          </div>
        </template>
      </el-drawer>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { Item } from "@/api/menu/super/type";
  import { onMounted, ref, watch } from "vue";
  import iconfont from "@/static/iconfont/iconfont.json";
  const iconfontList = iconfont;
  const isDialog = ref<boolean>(false);
  const isShowForm = ref<boolean>(false);
  const industryOptions = ref<any>([]);
  
  onMounted(async () => {
    if (type === "add") {
      resetFormData();
      form.value.parent_id = "0";
    } else if (type === "nextAdd") {
      resetFormData();
      form.value.parent_id = String(sendPather);
    } else {
      form.value.parent_id = String(formList.parent_id);
    }
  });
  let defaultForm = {
    parent_id: "1",
    method: "GET",
    title: "",
    path: "",
    icon_type: 1,
    icon: "",
    component: "",
    is_hide: 0,
    sort: 0,
    type: 1,
  };
  let form = ref<Item>({
    ...defaultForm,
  });
  // 重置表单
  const resetFormData = () => {
    Object.assign(form.value, { ...defaultForm });
  };
  
  const {
    type,
    isShowDiv = false,
    formList = {},
    sendPather = 0,
    SuperiorMenu = [],
  } = defineProps({
    type: {
      type: String,
    },
    isShowDiv: {
      type: Boolean,
    },
    formList: {
      type: Object,
    },
    sendPather: {
      type: Number || String,
    },
    SuperiorMenu: {
      type: Array,
      default: () => [],
    },
  });
  const emit = defineEmits([
    "close",
    "confirm",
    "getList",
    "editMenuList",
    "addMenuList",
    "addNextMenuList",
  ]);
  watch(
    () => {
      isShowDiv;
    },
    () => {
      isShowForm.value = isShowDiv;
    },
    { immediate: true }
  );
  watch(
    () => {
      formList;
    },
    () => {
      form.value = formList as Item;
      industryOptions.value.id = form.value.id;
      form.value.parent_id = form.value.parent_id;
    },
    { immediate: true }
  );
  watch(
    () => {
      sendPather;
    },
    () => {
      form.value.parent_id = Number(sendPather);
    },
    { immediate: true }
  );
  watch(
    () => SuperiorMenu,
    (newSuperiorMenu) => {
      industryOptions.value = newSuperiorMenu;
    },
    { immediate: true }
  );
  // 关闭dialog回调
  const closeDrawer = async () => {
    isDialog.value = false;
  
    resetFormData();
    // industryOptions.value.id = 0;
    emit("close");
  };
  // 取消按钮
  const closeForm = () => {
    isDialog.value = false;
    resetFormData();
    emit("close");
  };
  //提交表单按钮
  const onSubmit = async () => {
    if (type === "edit") {
      emit("editMenuList", form.value);
      emit("getList");
      emit("close");
      isDialog.value = false;
    } else if (type == "add") {
      emit("addMenuList", form.value);
      emit("getList");
      isDialog.value = false;
      emit("close");
      resetFormData();
    } else {
      emit("addMenuList", form.value);
      emit("getList");
      isDialog.value = false;
      emit("close");
      resetFormData();
    }
  };
  // select选择回调
  const selectChange = (id) => {};
  const localIconOptions = ref();
  </script>
  
  <style scoped lang="scss"></style>
  