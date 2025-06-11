<template>
    <div class="main h-full pt-10px flex-col flex">
        <div class="flex justify-between top">
            <el-input clearable @clear="search" v-model="keyword" style="max-width: 600px" placeholder="请输入您检索的校区名称"
                class="input-with-select">
                <template #append>
                    <el-button :icon="Search" @click="search" />
                </template>
            </el-input>
            <el-button size="large" type="primary" class="ml-20px" @click="addSchool">新增校区</el-button>
        </div>
        <div class="flex-1 flex flex-col">
            <div class="flex-1">
                <el-auto-resizer>
                    <template #default="{ height, width }">
                        <el-table-v2 :columns="columns" :data="tableData" :width="width" :height="height"  />
                    </template>
                </el-auto-resizer>
            </div>
            <div class="flex flex-justify-end pb-20px pr-30px">
                <el-pagination @change="changePageSize" :page-size="param.size" background layout="prev, pager, next"
                    :total="total" />
            </div>
        </div>

        <el-dialog v-model="showDialog" :title="editId ? '修改校区' : '新增校区'" width="500" destroy-on-close center>
            <div>
                <el-form ref="ruleFormRef" :model="form" label-width="auto">
                    <el-form-item label="校区名称" required prop="name">
                        <el-input v-model="form.name"></el-input>

                    </el-form-item>
                    <el-form-item label="所在省份" required prop="province_id">
                        <el-select @change="provinceChange" v-model="form.province_id" placeholder="省份">
                            <el-option v-for="item in provinceList" :key="item.id" :label="item.name"
                                :value="item.id" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="所在城市" required prop="city_id">
                        <el-select :disabled="!form.province_id" v-model="form.city_id" placeholder="城市">
                            <el-option v-for="item in cityList" :key="item.id" :label="item.name" :value="item.id" />
                        </el-select>
                    </el-form-item> <el-form-item label="状态" required prop="state">
                        <el-radio-group v-model="form.state">
                            <el-radio :value="1">启用</el-radio>
                            <el-radio :value="-1">禁用</el-radio>
                        </el-radio-group>

                    </el-form-item>

                </el-form>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="showDialog = false">取消</el-button>
                    <el-button type="primary" @click="comfirm">
                        {{ editId ? '确认修改' : "保存" }}
                    </el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="tsx" setup>
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column, FormInstance } from "element-plus";
import { Search } from "@element-plus/icons-vue";


import {
    ElButton,
    ElMessage,
    ElTag,
} from "element-plus";

import { getProvinceList, getSchoolList } from "@/api/common";
import { adminAddSchool, deleteSchoolApi, editAddSchoolApi } from "@/api/area";

const total = ref(0);
const tableData = ref([]);
const showDialog = ref(false);
const editId = ref<number | undefined>(undefined); // 使用 undefined 更清晰表示初始状态
const param = reactive({
    page: 1,
    size: 10,
});
const ruleFormRef = ref<FormInstance>();

const form = reactive({
    name: "",
    province_id: "",
    city_id: "",
    state: 1, // 设置一个默认值
});
const provinceList = reactive<any[]>([]);
const cityList = reactive<any[]>([
]);
const keyword = ref("");

const columns: Column<any>[] = [
    {
        key: "name",
        title: "校区",
        dataKey: "name",
        width: 250,
        align: "center",
    },
    {
        key: "province",
        title: "所在省",
        dataKey: "province",
        width: 250,
        align: "center",
    },
    {
        key: "province_id",
        title: "省编",
        dataKey: "province_id",
        width: 250,
        align: "center",
    },
    {
        key: "city",
        title: "所在市",
        dataKey: "city",
        width: 250,
        flexGrow: 1,
        align: "center",
    },
    {
        key: "city_id",
        title: "市编",
        dataKey: "city_id",
        width: 250,
        flexGrow: 1,
        align: "center",
    },
    {
        key: "state",
        title: "状态",
        dataKey: "state",
        width: 250,
        flexGrow: 1,
        align: "center",
        cellRenderer: ({ rowData }) => {
            const tagType = rowData.state === -1 ? "danger" : "success";
            const tagText = rowData.state === -1 ? "禁用" : "启用";
            return <ElTag type={tagType}>{tagText}</ElTag>;
        },
    },
    {
        key: "options",
        title: "操作",
        dataKey: "options",
        width: 600,
        align: "center",
        flexGrow: 1,
        cellRenderer: ({ rowData: row }) => {
            return (
                <div class="flex flex-items-center">
                    <el-button
                        type="primary"
                        onClick={() => {
                            editId.value = row.id;
                            form.name = row.name;
                            form.province_id = row.province_id;
                            form.city_id = row.city_id;
                            form.state = row.state;
                            showDialog.value = true;
                            provinceChange(row.province_id);
                        }}
                    >
                        编辑
                    </el-button>
                    <el-button
                        type="danger"
                        onClick={() => {
                            deleteSchoolById(row.id);
                        }}
                    >
                        删除
                    </el-button>
                </div>
            );
        },
    },
];

onMounted(() => {
    loadTable();
    getProvinceList(1).then(res => {
        if (res.code === 200) {
            provinceList.push(...res.data);
        } else {
            ElMessage.error(res.Message);
        }
    });
});

const addSchool = () => {
    editId.value = undefined;
    // 清空表单数据
    form.name = "";
    form.province_id = "";
    form.city_id = "";
    form.state = 1; // 恢复默认值
    cityList.length = 0; // 清空城市列表
    showDialog.value = true;
};

const loadTable = () => {
    getSchoolList(0, { ...param, search: keyword.value }).then((res) => {
        if (res.code === 200) {
            tableData.value = res.data.list;
            total.value = res.data.total;
        }
    });
};

const changePageSize = (e: number) => {
    param.page = e;
    loadTable();
};

const search = () => {
    param.page = 1;
    loadTable();
};

const provinceChange = (provinceId: number | string) => {
    form.city_id = ""; // 在切换省份时清空城市选择
    if (provinceId) {
        getProvinceList(provinceId).then(res => {
            if (res.code === 200) {
                cityList.length = 0; // 清空原有数据
                cityList.push(...res.data);
            } else {
                ElMessage.error(res.Message);
            }
        });
    } else {
        cityList.length = 0; // 如果没有选择省份，清空城市列表
    }
};

const comfirm = async () => {
    await ruleFormRef.value!.validate(async (valid) => {
        if (valid) {
            let res;
            if (!editId.value) {
                res = await adminAddSchool(form);
            } else {
                res = await editAddSchoolApi(editId.value, form);
            }

            if (res?.code === 200) {
                ElMessage.success(res.Message);
                showDialog.value = false;
                // 清空表单数据
                form.name = "";
                form.province_id = "";
                form.city_id = "";
                form.state = 1; // 恢复默认值
                cityList.length = 0; // 清空城市列表
                loadTable(); // 重新加载表格数据
            } else {
                ElMessage.error(res?.Message || '操作失败');
            }
        }
    });
};

const deleteSchoolById = (id: number) => {
    deleteSchoolApi(id).then(res => {
        if (res.code === 200) {
            ElMessage.success(res.Message);
            search(); // 重新加载表格数据
        } else {
            ElMessage.error(res.Message);
        }
    });
};
</script>

<style scoped>
.main {
    padding-left: 8px;

    .top {
        padding: 10px;
        border-radius: 4px;
        background-color: aliceblue;
    }

    :deep(.el-table-v2__main) {
        top: 0;
        right: 0;
        left: 0;
        bottom: 0;
    }

    :deep(.el-table-v2__empty) {
        top: 50% !important;
        left: 50% !important;
        transform: translate(-50%, -50%);
    }
}
</style>