<template>
  <div class="main h-full pt-10px flex-col flex">
    <div class="flex justify-between top">
      <el-input clearable @clear="search" v-model="keyword" style="max-width: 600px" placeholder="请输入检索的账号"
        class="input-with-select">
        <template #append>
          <el-button :icon="Search" @click="search" />
        </template>
      </el-input>
      <el-button size="large" type="primary" class="ml-20px" @click="addStudent">新增商户</el-button>
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
    <company-info v-if="showCompanyInfo" @close="close" :isAdd="isAdd" :item="item"></company-info>
  </div>
</template>

<script lang="tsx" setup>
import { defineComponent, onMounted, reactive, ref } from "vue";
import type { Column } from "element-plus";
import { Search } from "@element-plus/icons-vue";
import {
  ElButton,
  ElIcon,
  ElMessage,
  ElTag,
  ElTooltip,
  TableV2FixedDir,
} from "element-plus";
import { getUserList } from "@/api/studen";
import companyInfo from "@/components/company-info.vue";
import { deleteCompany, banUserById, initUserPassword, getCompanyList } from "@/api/common";
const showCompanyInfo = ref(false);
const total = ref(0);
const tableData = ref([]);
const param = reactive({
  page: 1,
  size: 10,
});
const item = ref({});
const columns: Column<any>[] = [
  {
    key: "name",
    title: "公司名称",
    dataKey: "name",
    width: 250,
    align: "center",
    flexGrow: 1,

  },
  {
    key: "lega_person",
    title: "法人",
    dataKey: "lega_person",
    width: 250,
    align: "center",
    flexGrow: 1,

  },

  {
    key: "school_name",
    title: "发布校区",
    dataKey: "school_name",
    width: 250,
    align: "center",
    flexGrow: 1,

  },
  {
    key: "phone",
    title: "联系方式",
    dataKey: "phone",
    width: 250,
    align: "center",
    flexGrow: 1,

  },
  {
    key: "username",
    title: "账号",
    dataKey: "username",
    width: 250,
    align: "center",
    flexGrow: 1,

  },
  {
    key: "options",
    title: "操作",
    dataKey: "options",
    width: 600,
    flexGrow: 2,
    align: "center",
    cellRenderer: ({ rowData: row }) => {
      return (
        <div class="flex flex-items-center">
          <el-button
            type="primary"
            onClick={() => {
              editStudentInfo(row);
            }}
          >
            编辑资料
          </el-button>
          <el-button
            type="danger"
            onClick={() => {
              deleteCom(row.id);
            }}
          >
            删除
          </el-button>
          <el-tooltip
            class="box-item"
            effect="light"
            content="初始化后密码为123456"
            placement="top-start"
          >
            <el-button
              style="background-color:#6E65E6;color:#fff;"

              onClick={() => {
                initPassword(row.user_id);
              }}
            >
              修改密码
            </el-button>
          </el-tooltip>
          <el-button
            type="warning"
            onClick={() => {
              ban(row.user_id);
            }}
          >
            {row.state === 1 ? "禁用" : "解除"}
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
const editStudentInfo = (row) => {
  item.value = row;
  showCompanyInfo.value = true;
  isAdd.value = false;
};
const loadTable = () => {
  getCompanyList({ search: keyword.value, ...param }).then((res) => {
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
const isAdd = ref(true);
const close = (isFresh = false) => {
  showCompanyInfo.value = false;
  if (isAdd.value) {
    param.page = 1;
  }
  loadTable();
};
const addStudent = () => {
  isAdd.value = true;
  showCompanyInfo.value = true;
};
const deleteCom = (id) => {
  deleteCompany(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
  });
};
const ban = (id) => {
  banUserById(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
  });
};
const initPassword = (id) => {
  initUserPassword(id).then((res) => {
    if (res.code == 200) {
      search();
      return ElMessage.success(res.Message);
    }
    ElMessage.error(res.Message);
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