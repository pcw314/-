<template>
    <div class="flex-1 p-10px flex-col h-full flex">
        <!-- 抽屉组件 -->
        <div v-if="isShow">
            <Drawer :type="type" :isShow="isShow" :nextParentId="nextParentId" @confirm="DealDrawerConfirm"
                :editData="ruleForm" @close="isShow = false"></Drawer>
        </div>
        <div class="flex flex-items-center justify-between">
            <el-button type="primary" :icon="Plus" @click="addCity()">新增</el-button>
            <el-button type="danger" :icon="Delete" @click="delCistData" :disabled="!selectedIds.length">删除</el-button>
        </div>
        <div class="flex-1">
            <el-auto-resizer>
                <template #default="{ height, width }">
                    <el-table-v2 :border="true" v-loading="isLoading" v-model:expanded-row-keys="expandedRowKeys"
                        :columns="columns" :data="treeData" :expand-column-key="expandColumnKey" :width="width"
                        :height="height">

                    </el-table-v2>
                </template>
            </el-auto-resizer>

        </div>
    </div>
</template>


<script lang="tsx" setup>
import {
    reactive,
    ref,
    computed,
    onMounted,
    unref,
    FunctionalComponent,
} from "vue";
import { areaRegionList, deleteAreaRegionList } from '@/api/area';
import { Plus, Delete } from '@element-plus/icons-vue'
import type { Column, CheckboxValueType } from 'element-plus';
import { ElMessage } from 'element-plus'
import Drawer from "./components/drawer.vue";
import {
    ElButton,
    ElCheckbox,
    ElPopconfirm,
    ElTag,
} from "element-plus";
const isLoading = ref(false);
const selectedIds = ref<Number[]>([]);
let tableData = ref<any>([]);

const isShow = ref<boolean>(false);

const type = ref<string>("add");
const addCity = () => {
    isShow.value = true;
    type.value = "add";
};
// 抽屉

interface RuleForm {
    id: number;
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
    checked?: boolean;
}

let ruleForm = reactive<RuleForm>({
    id: 0,
    parent_id: 0,
    name: "",
    type: 0,
    pinyin: "",
    zip: "",
    sort: 0,
    longitude: 0,
    latitude: 0,
    state: 1,
    is_search: 0,
});

onMounted(async () => {
    // 获取城市数据
    await getCityList();
});
// 获取城市数据
const getCityList = async () => {
    try {
        isLoading.value = true;
        const res = await areaRegionList();
        if (res.data) {
            tableData.value = res.data

        } else {
            tableData.value = [];
        }
        isLoading.value = false;

    } catch (error) {
        isLoading.value = false;

    }
};

// 刷新城市数据
const updata = async () => {
    await getCityList();
};

const treeProps = reactive({
    checkStrictly: false,
});

// 编辑城市
const redact = (data: RuleForm) => {
    ruleForm.id = data.id
    ruleForm.parent_id = data.parent_id
    ruleForm.type = data.type
    ruleForm.pinyin = data.pinyin
    ruleForm.zip = data.zip
    ruleForm.sort = data.sort
    ruleForm.longitude = data.longitude
    ruleForm.latitude = data.latitude
    ruleForm.state = data.state
    ruleForm.is_search = data.is_search
    type.value = "edit";
    type.value = "edit";
    isShow.value = true;
    if (ruleForm && typeof ruleForm === "object") {
        delete ruleForm.checked;
    }
};
const DealDrawerConfirm = () => {
    getCityList();
};

const nextParentId = ref<number>(1);
// 添加下一级
const handleAddChild = async (data) => {
    isShow.value = true;
    type.value = "add";

    ruleForm.parent_id = data.parent_id;
    nextParentId.value = Number(data.id);
};

// 删除 城市
const handleDelete = async (data: { id: number | string }) => {
    try {
        await deleteAreaRegionList(data.id);
        await getCityList();
        ElMessage({
            type: "success",
            message: "删除成功",
        });
    } catch (error) {
        console.log("城市删除失败", error);
    }
};

const delCistData = async () => {
    try {
        // 确保 selectedIds.value 是 (string | number) 类型数组
        const idsToDelete = selectedIds.value.map((id: any) => {
            return typeof id === "object" && "valueOf" in id ? id.valueOf() : id;
        });
        for (const ids of idsToDelete) {
            await deleteAreaRegionList(ids);
        }
        selectedIds.value = [];
        await getCityList();
        ElMessage({
            type: "success",
            message: "删除成功",
        });
    } catch (error) {
        console.log("城市删除失败", error);
    }
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
        key: "name",
        title: "地区名称",
        width: 200,
        flexGrow:1,

        cellRenderer: ({ rowData }) => rowData.name,
    },
    {
        key: "id",
        title: "编号",
        width: 300,
        flexGrow:1,

        cellRenderer: ({ rowData }) => rowData.id,
    },
    {
        key: "state",
        title: "是否禁用",
        flexGrow:1,

        width: 300,
        cellRenderer: ({ rowData }) => {
            const tagType = rowData.state === 0 ? "danger" : "success";
            const tagText = rowData.state === 0 ? "禁用" : "启用";

            return <ElTag type={tagType}>{tagText}</ElTag>;
        },
    },

    {
        key: "operations",
        title: "操作",
        align: "center",
        flexGroup: 1,
        with: 200,
        cellRenderer: ({ rowData }) => (
            <div>
                <ElButton size="small" onClick={() => handleAddChild(rowData)}>
                    添加下级
                </ElButton>
                <ElButton
                    size="small"
                    type="primary"
                    plain
                    onClick={() => redact(rowData)}
                >
                    编辑
                </ElButton>
                {rowData.children && rowData.children?.length > 0 ? null : (
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
const expandColumnKey = "name";

const expandedRowKeys = ref<string[]>([]);
// 列表展示的数据
const treeData = computed(() => tableData.value);
</script>

<style scoped lang="scss">
// :deep(.el-table-v2) {
//   border: var(--el-table-border) !important;
// }</style>
