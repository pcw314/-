<template>
    <div class="main h-full pt-10px flex-col flex">
        <div class="flex justify-between top">
            <el-input clearable @clear="search" v-model="keyword" style="max-width: 600px" placeholder="请输入您检索的图片关键字"
                class="input-with-select">
                <template #append>
                    <el-button :icon="Search" @click="search" />
                </template>
            </el-input>
        </div>
        <div class="flex-1 flex flex-col">
            <div class="flex-1">
                <el-auto-resizer>
                    <template #default="{ height, width }">
                        <el-table-v2 :columns="columns" :data="tableData" :width="width" :height="height" />
                    </template>
                </el-auto-resizer>
            </div>
            <div class="flex flex-justify-end pb-20px pr-30px">
                <el-pagination @change="changePageSize" :page-size="param.size" background layout="prev, pager, next"
                    :total="total" />
            </div>
        </div>
    </div>
</template>

<script lang="tsx" setup>
import {
    ElButton,
    ElMessage,
} from "element-plus";
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import { getOssList, deleteOss } from "@/api/oss"; // 假设这是获取图片信息的API调用

const total = ref(0);
const tableData = ref([]);
const param = reactive({
    page: 1,
    size: 10, // 根据提供的数据调整每页显示数量
});
const keyword = ref("");

const columns: Column<any>[] = [
    {
        key: "id",
        title: "ID",
        dataKey: "id",
        width: 80,
        align: "center",
    },
    {
        key: "name",
        title: "名称",
        dataKey: "name",
        flexGrow: 1,
        align: "center",
        width: 200,
    },
    {
        key: "type",
        title: "类型",
        dataKey: "type",
        width: 100,
        align: "center",
    },
    {
        key: "size",
        title: "大小",
        dataKey: "size",
        width: 120,
        align: "center",
        cellRenderer: ({ rowData }) => {
            return <span>{parseInt(rowData.size) / 1024} KB</span>;
        },
    },
    {
        key: "path",
        title: "预览",
        dataKey: "path",
        width: 150,
        align: "center",
        cellRenderer: ({ rowData }) => {
            return <img src={rowData.path} alt={rowData.name} style={{ width: '80px', height: 'auto' }} />;
        },
    },
    {
    key: "owner_type",
    title: "所有者角色",
    dataKey: "owner_type",
    width: 100, // 增加宽度以便更好地显示文本
    align: "center",
    cellRenderer: ({ rowData }) => {
        const roleMap = {
            '1': '学生',
            '2': '商户',
            '3': '管理员'
        };
        return <span>{roleMap[rowData.owner_type] || '未知'}</span>;
    },
},
    {
        key: "owner_id",
        title: "所有者ID",
        dataKey: "owner_id",
        width: 80,
        align: "center",
    },
    {
        key: "state",
        title: "状态",
        dataKey: "state",
        width: 100,
        align: "center",
        cellRenderer: ({ rowData }) => {
            const stateText = rowData.state === 0 ? "正常" : "异常";
            return <span>{stateText}</span>;
        },
    },
    {
        key: "created_at",
        title: "创建时间",
        dataKey: "created_at",
        width: 200,
        align: "center",
        cellRenderer: ({ rowData }) => {
            const date = new Date(rowData.created_at);
            return <span>{date.toLocaleString()}</span>;
        },
    },
    {
        key: "options",
        title: "操作",
        dataKey: "options",
        width: 100,
        align: "center",
        cellRenderer: ({ rowData: row }) => {
            return (
                <div class="flex flex-items-center">
                    <el-button
                        type="danger"
                        onClick={() => {
                            deleteImageByID(row.id);
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
});

const loadTable = () => {
    // 假设这是获取图片信息的API调用
    getOssList({ ...param, search: keyword.value }).then((res) => {
        if (res.code === 200) {
            tableData.value = res.data.list;
            total.value = res.data.total;
        }
    });
};

const changePageSize = (e) => {
    param.page = e;
    loadTable();
};

const search = () => {
    loadTable();
};

const deleteImageByID = (id) => {
    deleteOss(id).then(res => {
        if (res.code == 200) {
            ElMessage.success(res.Message)
        } else {
            ElMessage.error(res.Message)
        }
        search()
    })
}
</script>