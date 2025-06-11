<template>
    <div class="contain flex flex-row justify-center">
        <el-card class="box-card" :class="{ 'opacity-bg': showCity }">
            <div class="flex flex-row flex-1">
                <div class="left flex">
                    <h2 style="color: #4481eb;font-weight: 700;">xx平台</h2>
                    <h2 class="mt-10px mb-20px">Hi，欢迎{{ role == 1 ? '您成为接单者' : '您成为发布者' }}</h2>
                    <img src="@/static/img/function.svg" class="bg" />
                </div>
                <div class="flex-1 p-20px right">
                    <h2 class="mb-30px"> {{ role == 1 ? '创建个人名片' : '创建商户信息' }}</h2>
                    <template v-if="role == 2">
                        <div class="flex flex-row mb-20px flex-content-center flex-wrap">
                            <div class="gray-text">头像：</div>
                            <div class="circle flex justify-center flex-items-center" @click="showDialog = true">
                                <el-icon v-if="!imageUrl" color="#fff" size="22">
                                    <Camera />
                                </el-icon>

                                <img v-else :src="imageUrl" alt="" class="circle">
                            </div>
                        </div>
                        <div class="flex flex-row mb-20px align-content-center flex-wrap">
                            <div class="gray-text">公司名称：</div>
                            <el-input class="input" v-model="name" placeholder="Please input" size="large" />
                        </div>
                        <div class="flex flex-row mb-20px align-content-center flex-wrap">
                            <div class="gray-text">法人：</div>
                            <el-input class="input" v-model="lega_person" placeholder="Please input" size="large" />
                        </div>
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">电话：</div>
                            <el-input class="input" v-model="phone" placeholder="Please input" size="large" />
                        </div>
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">所在校区：</div>
                            <span style="line-height: 32px;color: #409eff;cursor: pointer;" @click="showCity = true">{{
                                school.name }}</span>
                            <el-button class="ml-10px" type="primary"
                                @click="showCity = true">{{ school.name ? "重新选择" : "点击选择" }}</el-button>
                        </div>
                    </template>

                    <template v-if="step == 1 && role == 1">
                        <div class="flex flex-row mb-20px flex-content-center flex-wrap">
                            <div class="gray-text">头像：</div>
                            <div class="circle flex justify-center flex-items-center" @click="showDialog = true">
                                <el-icon v-if="!imageUrl" color="#fff" size="22">
                                    <Camera />
                                </el-icon>

                                <img v-else class="circle" :src="imageUrl" alt="">
                            </div>
                        </div>
                        <div class="flex flex-row mb-20px align-content-center flex-wrap">
                            <div class="gray-text">昵称：</div>
                            <el-input class="input" v-model="name" placeholder="Please input" size="large" />
                        </div>
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">电话：</div>
                            <el-input class="input"  v-model="phone" placeholder="Please input" size="large" />
                        </div>
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">年龄：</div>
                            <el-input class="input" type="number" v-model="age" placeholder="Please input" size="large" />
                        </div>

                        <div class="flex flex-row align-content-center flex-wrap mb-20px">
                            <div class="gray-text">性别：</div>
                            <el-radio-group v-model="sex">
                                <el-radio :value=1 size="large">男</el-radio>
                                <el-radio :value=2 size="large">女</el-radio>
                            </el-radio-group>
                        </div>
                    </template>
                    <template v-if="step == 2 && role == 1">
                      
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">所在校区：</div>
                            <span style="line-height: 32px;color: #409eff;cursor: pointer;" @click="showCity = true">{{
                                school.name }}</span>
                            <el-button class="ml-10px" type="primary"
                                @click="showCity = true">{{ school.name ? "重新选择" : "点击选择" }}</el-button>
                        </div>
                        <div class="flex flex-row align-content-center flex-wrap mb-20px ">
                            <div class="gray-text">专业：</div>
                            <el-input class="input" v-model="major" placeholder="Please input" size="large" />
                        </div>
                    </template>
                </div>
            </div>
            <template #footer>
                <div v-if="role == 1">
                    <div class="flex flex-justify-end flex-1" v-if="step === 1">
                        <el-button type="primary" size="large" @click="step = 2">下一步</el-button>
                    </div>
                    <div class="flex flex-justify-end flex-1" v-if="step === 2">
                        <el-button type="primary" size="large" @click="step = 1">返回</el-button>
                        <el-button type="primary" size="large" @click="editStudent">确定</el-button>
                    </div>
                </div>
                <div v-if="role === 2" class="flex flex-justify-end flex-1" @click="loginCompany">
                    <el-button type="primary" size="large">确定</el-button>
                </div>
            </template>
        </el-card>
    </div>
    <el-dialog v-model="showDialog" width="600px" align-center destroy-on-close center @close="clearImage">
        <el-upload ref="uploadRef" v-if="showDialog" class="avatar-uploader" :auto-upload="false"
            v-model:file-list="fileList" :show-file-list="false" :on-change="onchange"
            :before-upload="beforeAvatarUpload" accept=".jpeg,.png">
            <template #trigger>
                <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon">
                    <Plus />
                </el-icon>
                <input id="upload" class="position-absolute" style="z-index: -11;opacity: 0;"> </input>
            </template>

        </el-upload>
        <template #footer>
            <div class="flex flex-justify-end">
                <label for="upload" v-if="imageUrl">
                    <el-button :loading="loading" class="btn"> 更换图片</el-button>
                </label>
                <label for="">
                    <el-button :loading="loading" class="btn  " style="margin-left: 20px;background-color: #409eff;"
                        @click="upload">确定</el-button>
                </label>
            </div>
        </template>
    </el-dialog>
    <choise-city v-if="showCity" @close="close" @selectSchool="selectSchool"></choise-city>
</template>

<script lang="ts" setup>
import { defineComponent, reactive, ref } from "vue";
import { ElMessage, UploadInstance, UploadUserFile, type UploadProps } from 'element-plus'
import { useUserStore } from "@/store/module/user";
import { editCompanyInfo, editStudentInfo, uploadFile } from "@/api/common";
const userStore = useUserStore();
import choiseCity from "../center/components/choise-city.vue";
import { useRouter } from "vue-router";


const route = useRouter()
const showDialog = ref<Boolean>(false)
const isSuccess = ref(false)
const imageUrl = ref('')
const uploadRef = ref<UploadInstance>()
const showCity = ref(false)
const major = ref("")
const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
}
const school = ref({ name: "", id: "" })
const fileList = ref<UploadUserFile[]>([])
const raw = ref(null)
const step = ref(1)
const name = ref<string>()
const lega_person = ref<string>()
const phone = ref<string>()
const age = ref<Number | null>(null)
const sex = ref(1)
const role = userStore.userInfo.role
const loading = ref(false)

const onchange = (rawFile) => {
    if (rawFile.raw.type !== 'image/jpeg' && rawFile.raw.type !== "image/png") {
        ElMessage.error('Avatar picture must be JPG or PNG format!')
        return false
    }
    imageUrl.value = URL.createObjectURL(rawFile.raw)
    console.log(rawFile)
    raw.value = rawFile

}
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {

    if (rawFile.type !== 'image/jpeg' && rawFile.type !== "image/png") {
        ElMessage.error('Avatar picture must be JPG or PNG format!')
        return false
    }

    return true
}
const clearImage = () => {
    if (!isSuccess.value) {
        imageUrl.value = ""

    } else {
        isSuccess.value = false

    }
}
const upload = () => {
    if (fileList.value.length == 0) {
        return ElMessage.success('请先上传图片')

    }
    loading.value = true
    console.log(fileList.value, 'fileList')
    let formData = new FormData()
    let img = fileList.value[0].raw
    formData.append('files', img as Blob)
    uploadFile(formData).then(res => {
        if (res.code === 200) {
            imageUrl.value = res.data[0].path
            showDialog.value = false
            isSuccess.value = true
        ElMessage.success( "上传成功")

        }
        loading.value = false

    }).catch(res => {
        showDialog.value = false
        ElMessage.success('上传失败')

    })
}
const close = () => {
    showCity.value = false
}
const selectSchool = (e) => {
    school.value = e
    showCity.value = false
}
const loginCompany = () => {
    if(!name.value || !lega_person.value ||!phone .value|| !school.value || !imageUrl.value){
         ElMessage.error('请完成初始信息填写');
return 
    }
    editCompanyInfo(userStore.userInfo.id, { name: name.value, lega_person: lega_person.value, phone: phone.value, school_id: school.value.id, avatar: imageUrl.value }).then(res => {
        if (res.code == 200) {
            route.push({ path: "/center/index" })
        }
        ElMessage.success(res.Message ? res.Message : "数据异常")

    })
}
const editStudent = ()=>{
   if(!name.value || !age.value ||!phone .value|| !school.value || !imageUrl.value || !sex.value || !major.value  ){
         ElMessage.error('请完成初始信息填写');
return 
    }
    editStudentInfo(userStore.userInfo.id, { name: name.value,  phone: phone.value, school_id: school.value.id, avatar: imageUrl.value,age:Number(age.value),sex:Number(sex.value),major:major.value }).then(res => {
        if (res.code == 200) {
            route.push({ path: "/center/index" })
        }
        ElMessage.success(res.Message ? res.Message : "数据异常")

    })
}
</script>

<style scoped lang="scss"></style>
<style>
.contain {
    height: 100vh;
    width: 100vw;
    background-color: #409eff;
    opacity: 0.8;
    align-items: center;

    .box-card {

        width: 1180px;
        border-radius: 10px;
        height: 600px;

        .left {
            flex-direction: column;
            align-items: center;
            width: 472px;

            .bg {
                width: 100%;
            }

        }

        .right {
            .gray-text {
                text-align: justify;
                display: flex;
                align-items: center;
                font-size: 16px;
                font-weight: 400;
                color: #8d97a6;
                padding-right: 5px;
            }

            .input {
                width: 400px;
            }

            .circle {
                cursor: pointer;
                width: 60px;
                height: 60px;
                border-radius: 50%;
                background-color: #4481eb;

                &:hover {
                    background-color: #75a6fa;
                }
            }
        }
    }

}

.avatar-uploader {
    display: flex;
    justify-content: center;
}

.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}

.avatar-uploader .avatar {
    width: 100%;
    height: 178px;
    display: block;
}

#upload{
    opacity: 0;
}
</style>