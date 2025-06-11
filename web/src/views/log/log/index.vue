<template>
    <div class="main h-full pt-10px flex-col flex">
        <div class="flex justify-between top">
            <el-input clearable @clear="search" v-model="keyword" style="max-width: 600px" placeholder="请输入您检索的日志关键字"
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
import { getLogList, deleteLog,} from "@/api/log";

const total = ref(0);
const tableData = ref([]);
const param = reactive({
    page: 1,
    size: 10,
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
        key: "method",
        title: "请求方法",
        dataKey: "method",
        width: 100,
        align: "center",
    },
    {
        key: "path",
        title: "路径",
        dataKey: "path",
        flexGrow: 1,
        align: "center",
        width: 100
    },
    {
        key: "query",
        title: "查询参数",
        dataKey: "query",
        flexGrow: 1,
        align: "center",
        width:100
    },
    {
        key: "status",
        title: "状态码",
        dataKey: "status",
        width: 100,
        align: "center",
    },
    {
        key: "ip",
        title: "IP地址",
        dataKey: "ip",
        width: 150,
        align: "center",
    },
    {
        key: "latency",
        title: "延迟(ms)",
        dataKey: "latency",
        width: 120,
        align: "center",
    },
    {
        key: "timestamp",
        title: "时间",
        dataKey: "timestamp",
        width: 200,
        align: "center",
        cellRenderer: ({ rowData }) => {
            const date = new Date(rowData.timestamp);
            return <span>{date.toLocaleString()}</span>;
        },
    },
    {
        key: "options",
        title: "操作",
        dataKey: "options",
        width: 50,
        align: "center",
        flexGrow: 1,
        cellRenderer: ({ rowData: row }) => {
            return (
                <div class="flex flex-items-center">
                    <el-button
                        type="danger"
                        onClick={() => {
                            deleteLogByID(row.id);
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
    // 假设这是获取日志信息的API调用
    getLogList({ ...param, search: keyword.value }).then((res) => {
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

const deleteLogByID = (id) => {
    deleteLog(id).then(res => {
        if (res.code == 200) {
            ElMessage.success(res.Message)
        } else {
            ElMessage.error(res.Message)

        }
        search()
    })
}
</script>