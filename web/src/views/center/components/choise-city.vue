<template>
  <div class="mask"></div>
  <div class="dialog">
    <div class="header ">
      <el-icon class="close" size="24" @click="close"><Close /></el-icon>
      <div class="flex flex-items-center">
        <div class="flex flex-items-center">
          <span>按照省份选择</span>
          <el-select @change="provinceChange" v-model="province" class="m-2" placeholder="省份" style="width: 140px">
            <el-option v-for="item in provinceList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </div>
        <div class="flex flex-items-center ml-10px">
          <span>选择城市</span>
          <el-select  @change="cityChange" :disabled="!province" v-model="city" class="m-2" placeholder="城市" style="width: 140px">
            <el-option v-for="item in cityList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </div>
        <div class="flex flex-items-center ml-10px">
          <span>选择校区</span>
          <el-select @change="selectSchool" :disabled="!city" v-model="school" class="m-2" placeholder="校区" style="width: 140px">
            <el-option v-for="item in schoolList" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </div>
      </div>

      <div class="flex flex-items-center mt-10px " style="cursor: pointer;">
        <span class="mr-10px">直接搜索校区</span>
        <el-select style="cursor: pointer;width: 240px" @change="selectSchool"  v-model="school" filterable  placeholder="收入您检索的校区"
         >
          <el-option v-for="item in allSchoolList" 
          :key="item.id" 
          :label="item.name" 
          :value="item.id" 
          />
        </el-select>
      </div>
    </div>
    <div class="body">
      <div class="hot-school">
        <h4 class="mb-20px">热门校区</h4>
        <div class="flex flex-itens-center flex-wrap">
          <div @click="selectHotSchool(item.id,item.name)" class="hot-item flex flex-itens-center  flex-justify-center flex-wrap " v-for="item in hotSchool"
            :key="item.code">{{ item.name }}</div>

        </div>
      </div>
      <!-- <div class="mt-20px">
        <div class="flex flex-items-center">
          <span style="font-weight: 700;">按字母选择 ：</span>
          <span class="alpha" v-for="i in alpha">{{ i }}</span>
        </div>
        <div class="school-group-section  ">
          <div class="school-item mt-14px flex " v-for="(item, index) in schoolSiteList" :key="index">

            <span class="school-group-char ml-10px">{{ item.firstChar }}</span>
            <div class="school-group-item-list flex-1" >
              <span class="school-name" v-for="(iitem, iindex) in item.cityList" :key="index">{{ iitem.name }}</span>
            </div>

          </div>
        </div>

      </div> -->
    </div>
  </div>
</template>

<script lang='ts' setup>
import { defineComponent, reactive, ref,defineEmits, onMounted, } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { getHotSchool, getProvinceList, getSchoolList } from '@/api/common'
import { ElMessage } from 'element-plus'

const emit =  defineEmits(['close','selectSchool'])
const province = ref('')
const city = ref('')
const school = ref('')
const allSchoolList = reactive<any[]>([])
const schoolList = reactive<any[]>([])
const provinceList = reactive<any[]>([])
const cityList = reactive<any[]>([
 ])
const hotSchool = reactive<any[]>([
  ])
const remoteMethod = (query: string) => {
console.log(11111111)
}
const close = ()=>{
  emit('close')
}
onMounted(()=>{
  getHotSchool({page:1,size:20}).then(res=>{
    if(res.code !==200){
      ElMessage.success(res.Message ? res.Message : "数据异常")
    }else{
      hotSchool.push(...res.data.list)
    }
  })
  getProvinceList(1).then(res=>{
    if(res.code !==200){

    }else{
      provinceList.push(...res.data)
    }
  })
  getSchoolList(0).then(res=>{
    if(res.code !==200){
      ElMessage.success(res.Message ? res.Message : "数据异常")
    }else{
      allSchoolList.push(...res.data.list)
    }
  })
})
const provinceChange =(e)=>{
  city.value = ""
  school.value= ""
 schoolList.length=0
 cityList.length = 0
  getProvinceList(e).then(res=>{
    if(res.code !==200){
      ElMessage.success(res.Message ? res.Message : "数据异常")
    }else{
      cityList.push(...res.data)
    }
  })
}
const cityChange =  (e)=>{
  school.value= ""
 schoolList.length=0

  getSchoolList(e).then(res=>{
    if(res.code !==200){
      ElMessage.success(res.Message ? res.Message : "数据异常")
    }else{
      schoolList.push(...res.data.list)
    }
  })
}

const selectSchool = (e)=>{
  let name = ""
  allSchoolList.forEach(item=>{
    if(item.id ===e){
      name=item.name
    }
  })
 
emit('selectSchool',{id:e,name:name})
}
const selectHotSchool = (id,name)=>{
emit('selectSchool',{id,name})
}
</script>

<style scoped>
.mask{
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #302e2e;
  opacity: 0.4;
}
.close{
  cursor: pointer;
  position: absolute;
  right: 10px;
  top: 10px;
}
.dialog {
  position: fixed;
  top: 30%;
  left: 50%;
  width: 900px;
  transform: translate(-50%,-50%);
  border-radius: 8px;

  background-color: #409eff;

}

.header {
  border-top-right-radius: 8px;
  border-top-left-radius: 8px;
  color: #fff;
  padding: 20px;
  background-color: #409eff;
  /* .el-select{
        --el-select-input-focus-border-color:#fff;
    } */
}

.body {
  border-top-right-radius: 8px;
  border-top-left-radius: 8px;
  padding: 20px;
  background-color: #fff;

  .hot-school {
    .hot-item {
      cursor: pointer;
      margin-right: 4px;
      margin-bottom: 6px;
      flex:1;
      min-width: 130px;
      padding: 4px;
      background-color: #f8f8f8;

      &:hover {
        background-color: #c6e2ff;
        color: #409eff;
      }
    }
  }

  .alpha {
    flex: 1;
    text-align: justify;
    color: #999;
    transition: all 0.2s;
    cursor: pointer;

    &:hover {
      color: #409eff;
    }
  }

  /* .school-group-section {
    position: relative;
    flex: 1;
    margin-right: -30px;
    padding-right: 30px;
    padding-bottom: 20px;
    max-height: 420px;
    overflow-y: scroll;

    .school-item {
      position: relative;
      overflow: hidden;
      padding: 6px 16px 6px 22px;
      border-radius: 3px;
      box-sizing: border-box;
      max-height: 45px;
      transition: background .2slinear;

      &::before {
        content: ' ';
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        z-index: 1;
        width: 6px;
        background: #f2f2f2;
      }

      .school-group-char {
        float: left;
        font-size: 24px;
        font-weight: 600;
        color: #999;
        line-height: 32px;
      }

      .school-group-item-list {
        text-align: justify;
        margin-right: -10px;
        margin-left: 34px;
        overflow: hidden;

        .school-name {
          float: left;
          margin: 6px 10px;
          display: block;
          font-size: 14px;
          color: #222;
          line-height: 20px;
          transition: all .2slinear;
        }
      }

    }
  } */

}

.item {
  color: #fff;
  background-color: antiquewhite;
}
</style>