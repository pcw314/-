<template>
    <div class="main h-full pt-10px flex-col flex">
      <div class="flex-1 flex flex-col">
        <!-- <div class="top flex flex-items-center"> -->
          <!-- <el-input
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
          </el-select> -->
        <!-- </div> -->
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
  
  import { allowOrNOComplain, getUserComplain } from "@/api/audit";
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
      key: "informer_name",
      title: "举报者",
      dataKey: "informer_name",
      width: 250,

      flexGrow: 1,
      align: "center",
      cellRenderer: ({ rowData: row }) => {
        return row.informer_name;
      },
    },
    {
      key: "reason",
      title: "举报原因",
      dataKey: "reason",
      width: 250,

      flexGrow: 1,
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
          />
        ))}
      </div>
    );
  }
},
  {
      key: "name",
      title: "被举报者",
      dataKey: "name",
      width: 250,
      flexGrow: 1,
      align: "center",
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
  search()
  }
  const loadTable = () => {
    getUserComplain({  ...param,}).then((res) => {
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
    tableData.value = []
    loadTable();
  };
  const allow = (id) => {
    allowOrNOComplain(id,{state:1}).then((res) => {
      if (res.code == 200) {
        ElMessage.success(res.Message);
        search()
      } else {
        ElMessage.error(res.Message);
      }
    });
  };
  const refuse = ()=>{
    allowOrNOComplain(complainId.value,{state:-2,reply:reply.value}).then((res) => {
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
  