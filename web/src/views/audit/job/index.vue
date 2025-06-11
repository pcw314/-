<template>
  <div class="main h-full pt-10px flex-col flex">

    <div class="flex-1 flex flex-col">
      <div class="flex-1">
        <el-auto-resizer>
          <template #default="{ height, width }">
            <el-table-v2 :estimated-row-height="40" :columns="columns" :data="tableData" :width="width"
              :height="height" />
          </template>
        </el-auto-resizer>
      </div>
      <div class="flex flex-justify-end pb-20px pr-30px">
        <el-pagination @change="changePageSize" :page-size="param.size" background layout="prev, pager, next"
          :total="total" />
      </div>
    </div>
    <el-dialog v-model="showComplain" width="500" destroy-on-close center>
      <div class="flex flex-items-center">
        <span>审核回复(通过可不填)</span>
        <el-input class="ml-10px" style="width: 300px;" v-model="reply"></el-input>
      </div>
      <template #footer>
        <div class=" flex flex-justify-end">
          <el-button @click="allow">审核通过</el-button>
          <el-button type="primary" @click="refuse">
            审核不过
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="tsx" setup>
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column } from "element-plus";
import { Search } from "@element-plus/icons-vue";


import {
  ElButton,
  ElMessage,
} from "element-plus";

import { jobComplainList } from "@/api/common";
import { allowOrNOComplain, deleteComplain } from "@/api/audit";
const reply = ref('')
const showComplain = ref(false)
const showStudentInfo = ref(false);
const total = ref(0);
const tableData = ref([]);
const param = reactive({
  page: 1,
  size: 10,
});
const complainId = ref<Number>()
const item = ref({});
const columns: Column<any>[] = [

  {
    key: "informer_name",
    title: "举报人",
    dataKey: "informer_name",
    width: 150,
    align: "center",
  },

  {
    key: "job_info.contact_name",
    title: "联系人",
    dataKey: "contact_name",
    width: 150,

    align: "center",
    cellRenderer: ({ rowData: row }) => {
 
      return row.job_info?.contact_name
    }
  },
  {
    key: "place",
    title: "工作地点",
    dataKey: "place",
    width: 250,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return row.job_info?.place
    }
  },
  {
    key: "salary",
    title: "薪资",
    dataKey: "salary",
    width: 250,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return <el-tag>{row.job_info?.salary?row.job_info?.salary:'-'}</el-tag>
    }
  },
  {
    key: "desc",
    title: "工作描述",
    dataKey: "desc",
    width: 250,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return <div style="padding: 10px 0;">{row.job_info?.description}</div>
    }
  },
  {
    key: "require",
    title: "任职要求",
    dataKey: "require",
    width: 250,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return <div style="padding: 10px 0;">{row.job_info?.require}</div>
    }
  },
  {
    key: "reason",
    title: "举报原因",
    dataKey: "reason",
    width: 250,
    align: "center",

  },
  {
  key: "image",
  title: "举报图片",
  dataKey: "image",
  width: 250,
  align: "center",
  cellRenderer: ({ rowData: row }) => {
    const imgs = row.image.split(',');

    if (!imgs.length || !imgs[0]) {
      return <span></span>;
    }

    return (
      <div class="flex flex-items-center">
        {imgs.map((item, index) => (
          <el-image
            key={index}
            src={item}
            style="width: 40px; height: 40px; margin: 0 2px; cursor: pointer;"
            preview-src-list={imgs}
            initial-index={index}
            hide-on-click-modal={false}
          >
            {/* 保留原有的 img 渲染方式 */}
            <img
              slot="error" // 如果图片加载失败，仍然显示默认 img
              src={item}
              style="width: 40px; height: 40px;"
            />
          </el-image>
        ))}
      </div>
    );
  }
},
{
      key: "state",
      title: "状态",
      dataKey: "require",
      width: 150,
      flexGrow:1,
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
              : '被驳回';
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
              complainId.value = row.id
              showComplain.value = true
            }}
          >
            审核
          </el-button>
          <el-button
            disabled={row.state !== 0}
            type="danger"
            onClick={() => {
              deleteCom(row.id);
            }}
          >
            删除举报
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

const loadTable = () => {
  jobComplainList({ search: keyword.value, ...param }).then((res) => {
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


const allow = () => {
  allowOrNOComplain(complainId.value as Number, { state: 1, reply: reply.value }).then(res => {
    if (res.code == 200) {
      ElMessage.success(res.Message)
    } else {
      ElMessage.error(res.Message)

    }
    showComplain.value = false
    search()

  })
}
const refuse = () => {
  if (!reply.value) {
    ElMessage.error("审核不过,必须给予原因")

    return
  }
  allowOrNOComplain(complainId.value as Number, { state: -1, reply: reply.value }).then(res => {
    if (res.code == 200) {
      ElMessage.success(res.Message)
    } else {
      ElMessage.error(res.Message)

    }
    showComplain.value = false
    search()

  })
}
const deleteCom = (id) => {
  deleteComplain(id).then(res => {
    if (res.code == 200) {
      ElMessage.success(res.Message)
    } else {
      ElMessage.error(res.Message)

    }
    search()
  })
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