<template>
  <div class="h-full flex flex-col">
    <div class="flex flex-items-center">
      <el-button type="primary" @click="addData()">
        <i class="iconfont icon-jiahao1 font-size-3 mr-1"></i>
        新增
      </el-button>
      <el-button
        type="danger"
        :disabled="!selectedIds.length"
        @click="deleteData"
      >
        <i class="iconfont icon-shanchu-copy font-size-3 mr-1"></i>
        删除
      </el-button>
    </div>
<div class="flex-1">
  <el-auto-resizer>
      <template #default="{ height, width }">
        <el-table-v2
          :border="true"
          v-loading="isLoading"
          v-model:expanded-row-keys="expandedRowKeys"
          :columns="columns"
          :data="treeData"
          :width="width"
          :expand-column-key="expandColumnKey"
          :height="height"
          fixed
        >
        </el-table-v2>
      </template>
    </el-auto-resizer>
</div>

    <div v-if="isShowDiv">
      <superMenuForm
          :isShowDiv="isShowDiv"
          @close="closeForm"
          @getList="getList"
          @editMenuList="editMenuList"
          @addMenuList="addMenuList"
          @addNextMenuList="addNextMenuList"
          :type="type"
          :formList="formList"
          :sendPather="sendPather"
          :SuperiorMenu="SuperiorMenu"
        ></superMenuForm>
    </div>
  </div>
</template>

<script lang="tsx" setup>
import superMenuForm from "./components/operate.vue";
import { getMenuList,createMnue,editMenue,deletMenu } from "@/api/menu";
import { MenuItem } from "@/api/menu/super/type";
import {
  CheckboxValueType,
  ElButton,
  ElCheckbox,
  ElMessage,
  ElPopconfirm,
  ElTag,
} from "element-plus";
import { computed, FunctionalComponent, onMounted, ref } from "vue";
const isLoading = ref<boolean>(false);
const selectedIds = ref<(number | string)[]>([]);
const isShowDiv = ref<boolean>(false);
let tableData = ref<any>([]);
const type = ref<string>("add");
let SuperiorMenu = ref<any>([]);
const minBOxHeight = ref();
// 添加子菜单
const sendPather = ref<any>();
onMounted(async () => {
  await getStoreMenuData();
});
// 获取数据列表
const getStoreMenuData = async () => {
  isLoading.value = true;
  const res = await getMenuList().then((res) => {
    if (res.code == 200) {
      tableData.value = [...res.data];
      console.log(tableData.value, "tableData.value");
      ElMessage.success(res.Message);
    } else {
      ElMessage.error(res.Message);
    }
    isLoading.value = false;
  });
  // 只将 res.list 赋值给 tableData，不影响 SuperiorMenu
};
// 更新表单
const editMenuList = async (form) => {
  try {
    await editMenue(form.id,{...form,parent_id:form.parent_id*1}).then(res=>{
      if(res.code==200){
        getStoreMenuData()
      }
    })
   
    ElMessage({
      message: "更新菜单成功",
      type: "success",
    });
  } catch (error) {
    console.log("更新失败", error);
  }
};
// 添加下一级表单数据
const addNextMenuList = async (form) => {
  try {
    await createMnue({...form,parent_id:form.parent_id*1});
    await getStoreMenuData();
    ElMessage({
      message: "添加下一级菜单成功",
      type: "success",
    });
  } catch (error) {
    console.log("添加下一级菜单失败", error);
  }
};
const addMenuList = async (form) => {
  try {
    await createMnue({...form,parent_id:form.parent_id*1});
    await getStoreMenuData();
    ElMessage({
      message: "添加菜单成功",
      type: "success",
    });
  } catch (error) {
    console.log("添加菜单失败", error);
  }
};
// 子组件回调函数
const getList = async () => {
  await getStoreMenuData();
};
// 删除
const deleteData = async () => {
  for (const ids of selectedIds.value) {
    await deletMenu(ids );
  }
  await getStoreMenuData();
  ElMessage({
    message: "删除成功",
    type: "success",
  });
  selectedIds.value = [];
  isLoading.value = false;
};
const closeForm = () => {
  isShowDiv.value = false;
  sendPather.value = "0";
};
// 新增
const addData = () => {
  type.value = "add";
  isShowDiv.value = true;
};

const handleAddEextMenu = async (rowData: MenuItem) => {
  // 传送父id
  type.value = "nextAdd";
  sendPather.value = rowData.id;
  // const resId = await getStoreMenuListApi();
  // 复制 res.list 后再对 SuperiorMenu 进行修改
  // SuperiorMenu.value = [...resId.list];
  const tempObj = {
    title: "无上级菜单",
    name: "无上级菜单",
    id: "0",
    parent_id: "0",
  };
  SuperiorMenu.value.unshift(tempObj);
  isShowDiv.value = true;
};
// 编辑
const formList = ref({});
const redact = async (rowData: MenuItem) => {
  type.value = "edit";
  isShowDiv.value = true;
  formList.value = { ...rowData };
  sendPather.value = rowData.parent_id;
  const resId = await getNotIdStoreMenuListApi({ not_id: rowData.id });
  // 复制 res.list 后再对 SuperiorMenu 进行修改
  SuperiorMenu.value = [...resId.list];
  const tempObj = {
    title: "无上级菜单",
    name: "无上级菜单",
    id: "0",
    parent_id: "0",
  };
  SuperiorMenu.value.unshift(tempObj);
};
// 删除单行
const handleDelete = async (rowData: MenuItem) => {
  await deletMenu(rowData.id);
  await getStoreMenuData();
  ElMessage({
    message: "删除成功",
    type: "success",
  });
};
type SelectionCellProps = {
  value: boolean;
  intermediate?: boolean;
  onChange: (value: CheckboxValueType) => any;
};
// 添加勾选框
const SelectionCell: FunctionalComponent<SelectionCellProps> = ({
  value,
  intermediate = false,
  onChange,
}) => {
  return (
    <ElCheckbox
      onChange={onChange}
      modelValue={value}
      indeterminate={intermediate}
    />
  );
};
const TagCellRenderer = ({ rowData }) => {
  const tagType = rowData.type == 1 ? "success" : "info";
  const tagText = ref<string>("菜单");
  if (rowData.type === 1) {
    tagText.value = "目录";
  } else if (rowData.type === 2) {
    tagText.value = "菜单";
  } else {
    tagText.value = "模块";
  }

  return <ElTag type={tagType}>{tagText.value}</ElTag>;
};
// 列 column 的配置数组
const columns = [
  // 勾选框
  {
    key: "selection",
    width: 50,
    cellRenderer: ({ rowData }) => {
      const onCheckboxChange = (value: CheckboxValueType) => {
        rowData.checked = value;
        if (value) {
          selectedIds.value.push(rowData.id);
        } else {
          selectedIds.value = selectedIds.value.filter(
            (id) => id !== rowData.id
          );
        }
        return value;
      };
      return (
        <SelectionCell value={rowData.checked} onChange={onCheckboxChange} />
      );
    },
    headerCellRenderer: () => {
      const onChange = (value) =>
        (tableData.value = tableData.value.map((row) => {
          row.checked = value;

          if (value) {
            if (!selectedIds.value.includes(row.id)) {
              selectedIds.value.push(row.id);
            }
          } else {
            selectedIds.value = selectedIds.value.filter((id) => id !== row.id);
          }
          return row;
        }));

      const allSelected = tableData.value.every((row) => row.checked);
      const containsChecked = tableData.value.some((row) => row.checked);

      return (
        <SelectionCell
          value={allSelected}
          intermediate={containsChecked && !allSelected}
          onChange={onChange}
        />
      );
    },
  },
  {
    key: "id",
    title: "ID",
    width: 100,
    cellRenderer: ({ rowData }) => {

      return rowData.id;
    },
  },
  {
    key: "title",
    title: "菜单名称",
    width: 100,
    cellRenderer: ({ rowData }) => {

return rowData.title;
},
  },
  {
    key: "icon",
    title: "图标",
    width: 150,
    cellRenderer: ({ rowData }) => {
      const icon = rowData.icon_type === 1 ? rowData.icon : undefined;

      const localIcon = rowData.icon_type === 2 ? rowData.icon : undefined;

      return (
        <div>
          <i class={`iconfont icon-${rowData.icon} font-size-[20px]`}></i>
        </div>
      );
    },
  },
  {
    key: "type",
    title: "菜单类型",
    width: 100,
    cellRenderer: TagCellRenderer,
  },
  {
    key: "method",
    title: "请求类型",
    width: 100,
    cellRenderer: ({ rowData }) => rowData.method,
  },

  {
    key: "path",
    title: "路由路径",
    width: 200,
    cellRenderer: ({ rowData }) => rowData.path,
  },
  {
    key: "is_hide",
    title: "是否隐藏",
    align: "center",
    width: 50,
    cellRenderer: ({ rowData }) => {
      const tagType = rowData.is_hide === 0 ? "success" : "danger";
      const tagText = rowData.is_hide === 0 ? "否" : "是";

      return <ElTag type={tagType}>{tagText}</ElTag>;
    },
  },
  {
    key: "parent_id",
    title: "父级id",
    align: "center",
    width: 100,
    cellRenderer: ({ rowData }) => rowData.parent_id,
  },
  {
    key: "sort",
    title: "排序",
    align: "center",
    width: 100,
    cellRenderer: ({ rowData }) => rowData.sort,
  },
  {
    key: "operations",
    title: "操作",
    align: "center",
    cellRenderer: ({ rowData }) => (
      <div>
        <ElButton size="small" onClick={() => handleAddEextMenu(rowData)}>
          添加子菜单
        </ElButton>
        <ElButton
          size="small"
          type="primary"
          plain
          onClick={() => redact(rowData)}
        >
          编辑
        </ElButton>
        {rowData.children && rowData.children.length > 0 ? null : (
          <ElPopconfirm
            title="确认删除?"
            confirm-button-text="确认"
            cancel-button-text="取消"
            onConfirm={() => handleDelete(rowData)}
          >
            {{
              reference: () => (
                <ElButton size="small" type="danger">
                  删除
                </ElButton>
              ),
            }}
          </ElPopconfirm>
        )}
      </div>
    ),
    width: 300,
  },
];
// 设置要缩进的列
const expandColumnKey = "id";

const expandedRowKeys = ref<string[]>([]);
// 列表展示的数据
const treeData = computed(() => tableData.value);
</script>

<style scoped lang="scss"></style>
