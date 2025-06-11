<template>
  <div class="main h-full pt-10px flex-col flex">
    <div class="flex-1 flex flex-col">
      <div class="top flex flex-items-center">
        <el-input
          clearable
          @clear="search"
          v-model="keyword"
          style="max-width: 400px"
          placeholder="请输入您检索的名称"
          class="input-with-select"
        >
          <template #append>
            <el-button :icon="Search" @click="search" />
          </template>
        </el-input>
        <el-select class="ml-20px" v-model="state" style="width:200px" @change="change">
          <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
        </el-select>
      </div>
      <div class="flex-1">
        <el-auto-resizer>
          <template #default="{ height, width }">
            <el-table-v2
              :estimated-row-height="40"
              :columns="columns"
              :data="tableData"
              :width="width"
              :height="height"
            />
          </template>
        </el-auto-resizer>
      </div>
      <div class="flex flex-justify-end pb-20px pr-30px">
        <el-pagination
          @change="changePageSize"
          :page-size="param.size"
          background
          layout="prev, pager, next"
          :total="total"
        />
      </div>
    </div>
    <el-dialog v-model="showComplain" width="500" destroy-on-close center>
      <div class="flex flex-items-center">
        <span>审核回复</span>
        <el-input
          class="ml-10px"
          style="width: 300px"
          v-model="reply"
        ></el-input>
      </div>
      <template #footer>
        <div class="flex flex-justify-end">
          <el-button type="primary" @click="refuse"> 驳回 </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="tsx" setup>
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column } from "element-plus";
import { Search } from '@element-plus/icons-vue'
import { ElButton, ElMessage } from "element-plus";
import { checkJob, getJobList } from "@/api/audit";
const reply = ref("");
const total = ref(0);
const tableData = ref([]);
const param = reactive({
  page: 1,
  size: 10,
});
const showComplain =ref(false)
const complainId = ref<Number>();
const state = ref(0)
const options =[{
  value: 0,
  label: '待审核',
},{
  value: 1,
  label: '审核通过',
},{
  value: -2,
  label: '驳回',
}]
const columns: Column<any>[] = [
  {
    key: "post",
    title: "联系电话",
    dataKey: "post",
    width: 150,
    flexGrow: 1,
    align: "center",
  },
  {
    key: "companyname",
    title: "公司名称",
    dataKey: "contact_number",
    width: 200,
    flexGrow: 1,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return row.enterprise_info.name;
    },
  },
  {
    key: "contact_name",
    title: "发布者",
    dataKey: "contact_name",
    width: 150,
    flexGrow: 1,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return row.contact_name;
    },
  },
  {
    key: "contact_number",
    title: "联系电话",
    dataKey: "contact_number",
    width: 150,
    flexGrow: 1,
    align: "center",
  },
  {
    key: "salary",
    title: "薪资",
    dataKey: "salary",
    width: 250,
    align: "center",
    flexGrow: 1,

    cellRenderer: ({ rowData: row }) => {
      return <el-tag>{row.salary}</el-tag>;
    },
  },
  {
    key: "working_time",
    title: "工作时间",
    dataKey: "working_time",
    width: 150,
    flexGrow: 1,
    align: "center",
  },
  {
    key: "place",
    title: "工作地点",
    dataKey: "place",
    width: 250,
    align: "center",
    flexGrow: 1,
  },

  {
    key: "desc",
    title: "工作描述",
    dataKey: "desc",
    width: 250,
    flexGrow: 1,

    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return <div style="padding: 10px 0;">{row.description}</div>;
    },
  },
  {
    key: "require",
    title: "任职要求",
    dataKey: "require",
    width: 250,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return <div style="padding: 10px 0;">{row.require}</div>;
    },
  },
  {
    key: "state",
    title: "状态",
    dataKey: "require",
    width: 150,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      let type =
        row.state === 0
          ? "primary"
          : row.state === 1
            ? "success"
            : row.state === -1
              ? "danger"
              : "info";
      let text =
        row.state === 0
          ? "待审核"
          : row.state === 1
            ? "审核通过"
            : row.state === -1
              ? "审核不通过"
              : "已下架";
      return <el-tag type={type}>{text}</el-tag>;
    },
  },

  {
    key: "options",
    title: "操作",
    dataKey: "options",
    width: 600,
    flexGrow: 1,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return (
        <div class="flex flex-items-center">
          <el-button
            type="primary"
            disabled={row.state != 0}
            onClick={() => {
              allow(row.id);
            }}
          >
            通过审核
          </el-button>
          <el-button
            disabled={row.state !== 0}
            type="danger"
            onClick={() => {
              showComplain.value = true;
              complainId.value = row.id
            }}
          >
            驳回
          </el-button>
        </div>
      );
    },
  },
];
const keyword = ref("");
onMounted(() => {
  loadTable();
});
const change = (state)=>{
console.log(state)
search()
}
const loadTable = () => {
  getJobList({ search: keyword.value, ...param,state:state.value }).then((res) => {
    if ((res.code = 200)) {
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
  param.page = 1;
  loadTable();
};
const allow = (id) => {
  checkJob(id,{state:1}).then((res) => {
    if (res.code == 200) {
      ElMessage.success(res.Message);
      search()
    } else {
      ElMessage.error(res.Message);
    }
  });
};
const refuse = ()=>{
  checkJob(complainId.value,{state:-2,reply:reply.value}).then((res) => {
    if (res.code == 200) {
      ElMessage.success(res.Message);
      showComplain.value =false
      search()
    } else {
      ElMessage.error(res.Message);
    }
  });
}
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
