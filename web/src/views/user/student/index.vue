<template>
  <div class="main h-full pt-10px flex-col flex">
    <div class="flex justify-between top">
      <el-input
        clearable
        @clear="search"
        v-model="keyword"
        style="max-width: 600px"
        placeholder="请输入检索的账号"
        class="input-with-select"
      >
        <template #append>
          <el-button :icon="Search" @click="search" />
        </template>
      </el-input>
      <el-button size="large" type="primary" class="ml-20px" @click="addStudent"
        >新增学生</el-button
      >
    </div>
    <div class="flex-1 flex flex-col">
      <div class="flex-1">
        <el-auto-resizer>
          <template #default="{ height, width }">
            <el-table-v2
              :columns="columns"
              :data="tableData"
              :width="width"
              :height="height"
              fixed
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
    <student-info
      v-if="showStudentInfo"
      @close="close"
      :isAdd="isAdd"
      :item="item"
    ></student-info>
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
import studentInfo from "@/components/student-info.vue";
import { deleteStudentByID, banUserById, initUserPassword, } from "@/api/common";
const showStudentInfo = ref(false);
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
    title: "名称",
    dataKey: "name",
    width: 150,
    align: "center",
  },
  {
    key: "sex",
    title: "性别",
    dataKey: "sex",
    align: "center",

    width: 50,
    cellRenderer: ({ cellData: sex }) => {
      let str = "";
      if (sex == 1) {
        str = (
          <div style="width:8px;height:8px;border-radius:50%;background-color:rgb(121.3, 187.1, 255)"></div>
        );
      } else {
        str = (
          <div style="width:8px;height:8px;border-radius:50%;background-color:rgb(248, 152.1, 152.1)"></div>
        );
      }
      return (
        <div class="flex flex-items-center">
          {str}
          <span class="ml-4px">{sex == 1 ? "男" : "女"}</span>
        </div>
      );
    },
  },
  {
    key: "school_name",
    title: "学校",
    dataKey: "school_name",
    width: 200,
    align: "center",
  },
  {
    key: "major",
    title: "专业",
    dataKey: "major",
    width: 150,
    align: "center",
  },
  {
    key: "age",
    title: "年龄",
    dataKey: "age",
    width: 50,
    align: "center",
  },
  {
    key: "phone",
    title: "联系方式",
    dataKey: "phone",
    width: 150,
    align: "center",
  },
  {
    key: "username",
    title: "账号",
    dataKey: "username",
    width: 150,
    align: "center",
  },
  {
    key: "options",
    title: "操作",
    dataKey: "options",
    width: 500,
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
              deleteStudent(row.id);
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
  showStudentInfo.value = true;
  isAdd.value = false;
};
const loadTable = () => {
  getUserList({ search: keyword.value, ...param }).then((res) => {
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
  showStudentInfo.value = false;
  if (isAdd.value) {
    param.page = 1;
  }
  loadTable();
};
const addStudent = () => {
  isAdd.value = true;
  showStudentInfo.value = true;
};
const deleteStudent = (id) => {
  console.log(id, "2222");
  deleteStudentByID(id).then((res) => {
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
const initPassword = (id)=>{
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
