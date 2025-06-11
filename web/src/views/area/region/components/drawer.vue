<template>
  <el-drawer v-model="showDrawer" :show-close="false" @close="closeDrawer">
    <template #header="{ close, titleId, titleClass }">
      <div style="border-bottom: 1px solid rgb(239, 239, 245); display: flex">
        <h4 :id="titleId" :class="titleClass">
          {{ type == "add" ? "新增" : "编辑" }}
        </h4>
      </div>
    </template>
    <el-form ref="form" :model="ruleForm" label-width="80px" :rules="rules">
      <el-form-item label="类型" label-position="top">
        <el-select v-model="ruleForm.type" placeholder="地区(省下面的地级市)">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="标准行政区域代码" prop="id" label-position="top">
        <el-input
          placeholder="请输入"
          :disabled="type == 'add' ? false : true"
          v-model.number="ruleForm.id"
        />
      </el-form-item>

      <el-form-item label="名称" prop="name" label-position="top">
        <el-input placeholder="请输入" v-model="ruleForm.name" />
      </el-form-item>

      <el-form-item label="地区名称拼音" prop="pinyin" label-position="top">
        <el-input placeholder="请输入" v-model="ruleForm.pinyin" />
      </el-form-item>

      <!-- <el-form-item label="地区经纬度" label-position="top">
        <el-col :span="8">
          <el-form-item prop="longitude">
            <el-input
              v-model.number="ruleForm.longitude"
              placeholder="请输入"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col class="text-center" :span="2">
          <span class="text-gray-500">-</span>
        </el-col>
        <el-col :span="8">
          <el-form-item prop="latitude">
            <el-input
              v-model.number="ruleForm.latitude"
              placeholder="请输入"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="5" :push="1">
          <el-button type="primary" @click="dialogMapVisible"
            >选择地区</el-button
          >
        </el-col>
      </el-form-item> -->

      <el-form-item label="邮编" prop="zip" label-position="top">
        <el-input placeholder="请输入" v-model="ruleForm.zip" />
      </el-form-item>

      <el-form-item label="排序" label-position="top">
        <el-input placeholder="请输入" v-model.number="ruleForm.sort" />
      </el-form-item>

      <el-form-item label="是否启用" label-position="top">
        <el-radio-group v-model="ruleForm.state">
          <el-radio :value="1">是</el-radio>
          <el-radio :value="0">否</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button
        @click="showDrawer = false"
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
        确认
      </el-button>
    </template>
  </el-drawer>

  <!-- <el-dialog v-model="MapVisible" title="选择地区" width="900">
    <div class="z-99">
      <MapControl
        @updateLocation="handleLocationUpdate"
        @closeMapBox="closeMapBox"
      ></MapControl>
    </div>
  </el-dialog> -->
</template>
<script setup lang="ts">
// import MapControl from "@/components/map-control.vue";
import { ref, reactive, watch } from "vue";
import { redactCityApi, addCityApi } from "@/api/city/index";
import { ElMessage } from "element-plus";
import { createAreaRegionList, editAreaRegionList } from "@/api/area";
const MapVisible = ref(false);
const showDrawer = ref<Boolean>(false);
const form = ref<any>(null);
interface RuleForm {
  id: number | null;
  parent_id: number;
  name: string;
  type: number;
  pinyin: string;
  zip: string;
  sort: number;
  longitude: number;
  latitude: number;
  state: number;
  is_search: number;
}
let defaultRegion = {
  id: null,
  parent_id: 1,
  name: "",
  type: 0,
  pinyin: "",
  zip: "",
  sort: 0,
  longitude: 0,
  latitude: 0,
  state: 1,
  is_search: 0,
};
const emit = defineEmits(["close", "confirm"]);
let ruleForm = reactive<RuleForm>({
  ...defaultRegion,
});
const dialogMapVisible = () => {
  MapVisible.value = true;
};
// 处理子组件传递的经纬度信息
const handleLocationUpdate = (locationData: any) => {
  ruleForm.latitude = locationData.lat;
  ruleForm.longitude = locationData.lng;
};
const closeMapBox = (isMap: boolean) => {
  MapVisible.value = isMap;
};
// 父表格传递的数据
const {
  isShow = false,
  editData = {},
  type = "add",
  nextParentId,
} = defineProps({
  isShow: {
    type: Boolean,
  },
  editData: {
    type: Object,
  },
  type: {
    type: String,
    default: "add",
  },
  nextParentId: {
    type: Number,
    default: 1,
  },
});
// 监听父表格传递过来的数据
watch(
  () => {
    return nextParentId;
  },
  () => {
    ruleForm.parent_id = nextParentId;
  },
  { immediate: true }
);
watch(
  () => {
    return isShow;
  },
  () => {
    showDrawer.value = isShow;
  },
  { immediate: true }
);
watch(
  () => {
    return editData;
  },
  () => {
    if (type === "edit") {
      Object.assign(ruleForm, editData);
    }
  },
  { immediate: true }
);
// select选择
const options = [
  {
    value: 0,
    label: "区域",
  },
  {
    value: 1,
    label: "国家",
  },
  {
    value: 2,
    label: "省/自治区/直辖市",
  },
  {
    value: 3,
    label: "地区(省下面的地级市)",
  },
  {
    value: 4,
    label: "县/市(县级市)/区",
  },
  {
    value: 5,
    label: "海外",
  },
];
interface RuleForm {
  id: number | null;
  parent_id: number;
  name: string;
  type: number;
  pinyin: string;
  zip: string;
  sort: number;
  longitude: number;
  latitude: number;
  state: number;
  is_search: number;
}
// from表单验证
const rules = reactive({
  name: [{ required: true, message: "请输入内容！", trigger: "change" }],
  id: [
    { required: true, message: "请输入内容！", trigger: "blur" },
    { pattern: /^[0-9]+$/, message: "请输入数字", trigger: "blur" },
  ],
  pinyin: [
    { required: true, message: "请输入内容！", trigger: "blur" },
    { pattern: /^[A-Za-z]+$/, message: "请输入英文字母", trigger: "blur" },
  ],
  zip: [
    { required: true, message: "请输入内容！", trigger: "blur" },
    { pattern: /^[0-9]+$/, message: "请输入数字", trigger: "blur" },
  ],
});
const closeDrawer = () => {
  resetFormData();

  emit("close");
};
const confirm = () => {
  resetFormData();
  emit("confirm");
};
// 重置表单
const resetFormData = () => {
  Object.assign(ruleForm, { ...defaultRegion });
};
const onSubmit = () => {
  form.value.validate(async (valid) => {
    if (valid) {
      if (type == "edit") {
        ruleForm.parent_id = ruleForm.parent_id * 1;
        try {
          console.log(ruleForm.id,'ruleForm.id')
          await editAreaRegionList(ruleForm.id,ruleForm);
          ElMessage({
            type: "success",
            message: "修改成功",
          });
          confirm();
          showDrawer.value = false;
        } catch (error) {
          console.log("修改失败", error);
        }
      } else {
        ruleForm.parent_id = nextParentId * 1;
        try {
          await createAreaRegionList(ruleForm);
          ElMessage({
            type: "success",
            message: "添加成功",
          });
          confirm();
          showDrawer.value = false;
        } catch (error) {
          console.log("添加失败", error);
        }
      }
    }
  });
};
</script>

<style lang="scss" scoped></style>
