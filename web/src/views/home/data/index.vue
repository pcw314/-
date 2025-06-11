<template>
  <div>
    <el-card shadow="always " class="position-relative">
      <img style="width: 100%;height: 140px;" src="@/static/img/admin-bg.jfif" alt="">
      <span class="text">好久不见，</span>
      <span class="text-bt">欢迎回来！</span>
    </el-card>
    <el-scrollbar>
      <div class="flex flex-items-center justify-between mt-40px">
        <div class="card card-one flex flex-items-center justify-center p-10px ">
          <img class="ml-40px" src="@/static/img/log.svg" alt="" style="width: 100px;height: 100px;">
          <div class="bg-text flex-1 justify-center flex-col flex-items-center">
            <div class="text-center">总学生数</div>
            <div class="text-center mt-10px">{{ studentSum }}</div>
          </div>
        </div>
        <div class="card card-two">
          <div class="card card-two flex flex-items-center justify-center p-10px ">
            <img class="ml-40px" src="@/static/img/log.svg" alt="" style="width: 100px;height: 100px;">
            <div class="bg-text flex-1 justify-center flex-col flex-items-center">
              <div class="text-center">总商户数</div>
              <div class="text-center mt-10px">{{ enterpriseSum }}</div>
            </div>
          </div>
        </div>
        <div class="card card-three">
          <div class="card card-three flex flex-items-center justify-center p-10px ">
            <img class="ml-40px" src="@/static/img/log.svg" alt="" style="width: 100px;height: 100px;">
            <div class="bg-text flex-1 justify-center flex-col flex-items-center">
              <div class="text-center">近七日新增学生数</div>
              <div class="text-center mt-10px">{{ newStudentSum }}</div>
            </div>
          </div>
        </div>
        <div class="card card-four">
          <div class="card card-four flex flex-items-center justify-center p-10px ">
            <img class="ml-40px" src="@/static/img/log.svg" alt="" style="width: 100px;height: 100px;">
            <div class="bg-text flex-1 justify-center flex-col flex-items-center">
              <div class="text-center">近七日新增商户数</div>
              <div class="text-center mt-10px">{{ newEnterpriseSum }}</div>
            </div>
          </div>
        </div>
      </div>

    </el-scrollbar>
    <el-scrollbar class="mt-100px">
      <div class="flex flex-items-center ">
        <div ref="peoppleChart" class="flex-1 mr-20px" style="min-height: 300px;"></div>
        <div class="flex-1 mr-20px" ref="taskRef" style="min-height: 400px;"></div>

      </div>
    </el-scrollbar>

  </div>
</template>

<script lang='ts' setup>
import { getStaffInfo, getDashBoard } from '@/api/common';
import studentInfo from '@/components/student-info.vue';
import { useUserStore } from '@/store/module/user';
import * as echarts from 'echarts';
import { ElMessage } from 'element-plus';
import { defineComponent, reactive, ref, onMounted, markRaw } from 'vue'
const userStore = useUserStore()
const studentSum = ref(0)
const enterpriseSum = ref(0)
const newEnterpriseSum = ref(0)
const newStudentSum = ref(0)
const peoppleChart = ref('null')
const taskRef = ref(null)
const taskRef2 = ref(null)
let chartInstance = ref(null)
let Task = ref(null)
const taskOp = {
  title: {
    text: '近七日兼职发布数'
  },
  tooltip: {
    trigger: 'axis'
  },

  grid: {
    left: '3%',
    right: '5%',
    bottom: '15%',
    containLabel: true
  },

  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      type: 'line',
      stack: 'Total',
      data: [320, 332, 301, 334, 390, 330, 320]
    },

  ]
}
const SevenNew = {
  title: {
    text: '近七日登录人数'
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {},
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: [
    {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    }
  ],
  yAxis: [
    {
      type: 'value'
    }
  ],
  series: [
    {
      name: '近七日商户登录数',
      type: 'bar',
      emphasis: {
        focus: 'series'
      },
      itemStyle: {
        color: "#6E65E6"
      },
      data: [320, 332, 301, 334, 390, 330, 320]
    },

    {
      name: '近七日学生登录数',
      type: 'bar',

      itemStyle: {
        color: "#409eff"
      },
      data: [862, 1018, 964, 1026, 1679, 1600, 1570],
      emphasis: {
        focus: 'series'
      },
      markLine: {
        lineStyle: {
          type: 'dashed'
        },
      }
    },

  ]
}
onMounted(() => {
  getDashBoard().then(res => {
    if (res.code === 200) {
      console.log(res.data)
      studentSum.value = res.data.student_sum
      enterpriseSum.value = res.data.enterprise_sum
      newStudentSum.value = res.data.new_student_sum
      newEnterpriseSum.value = res.data.new_enterprise_sum
      SevenNew.xAxis[0].data = res.data.date
      SevenNew.series[0].data  = res.data.new_enterprise
      SevenNew.series[1].data  = res.data.new_student
      taskOp.series[0].data = res.data.new_job
      taskOp.xAxis.data = res.data.date
      chartInstance.value = markRaw(echarts.init(peoppleChart.value));
      console.log(SevenNew,'SevenNew')
      chartInstance.value.setOption(SevenNew, true);
      Task.value = markRaw(echarts.init(taskRef.value));
      Task.value.setOption(taskOp, true);
    } else {
      ElMessage.error(res.message)
    }
  })

  getStaffInfo().then((res) => {
    if (res.code == 200) {
      userStore.setUserInfo(res.data)
    } else {
      ElMessage.error(res.message)
    }
  })
})
window.addEventListener('resize', () => {
  setTimeout(() => {
    chartInstance.value.resize()
    Task.value.resize()
  }, 100);
});
</script>

<style scoped>
.text,
.text-bt {
  position: absolute;
  color: #fff;
  font-size: 22px;
  top: 10%;
  left: 10%;
  transform: translate(-50%);
  text-shadow: 1px 0 #9a9ea1;
}

.text-bt {
  font-size: 22px;
  top: 40%;
  left: 12%;
  transform: translate(-50%);
}

:deep(.el-card__body) {
  padding: 0px !important;
}

.user-data {
  background-color: #ecf5ff;
}

.bg-white {
  border-radius: 4px;
  width: 100px;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.card {
  font-size: 22px;
  font-weight: 700;
  height: 200px;
  flex: 1;
  min-width: 300px;
  background-color: #9a9ea1;
  margin-right: 30px;
  color: #fff;
  border-radius: 6px;
  transition: all 0.5s;
  transform-origin: center;

  &:hover {
    box-shadow: 1px 2px #d7d8da;
    transform: scale(1.02);
  }

  &:last-child {
    margin-right: 0;

  }

}

.card-one {
  background-color: #409eff;
}

.card-two {
  background-color: #6E65E6;
}

.card-three {
  background-color: #62D0D4;
}

.card-four {
  background-color: #6BB860;
}
</style>