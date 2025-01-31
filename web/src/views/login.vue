<template>
    <div>
        <div class="switch-language">
            <el-dropdown size="large" :hide-timeout="50" placement="bottom-end" :hide-on-click="true">
                <Icon name="fa fa-globe" color="var(--el-text-color-secondary)" size="28" />
                <template #dropdown>
                    <el-dropdown-menu class="chang-lang">
                        <el-dropdown-item v-for="item in config.lang.langArray" :key="item.name" @click="editDefaultLang(item.name)">
                            {{ item.value }}
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
        <div @contextmenu.stop="" id="bubble" class="bubble">
            <canvas id="bubble-canvas" class="bubble-canvas"></canvas>
        </div>
        <div class="login">
            <div class="login-box">
                <div class="head">
                    <img src="~assets/login-header.png" alt="" />
                </div>
                <div class="form">
                    <img class="profile-avatar" src="~assets/avatar.png" alt="" />
                    <div class="content">
                        <el-form @keyup.enter="onSubmitPre()" ref="formRef" :rules="rules" size="large" :model="form">
                            <el-form-item prop="username">
                                <el-input
                                    ref="usernameRef"
                                    type="text"
                                    clearable
                                    v-model="form.username"
                                    :placeholder="t('login.Please enter an account')"
                                >
                                    <template #prefix>
                                        <Icon name="fa fa-user" class="form-item-icon" size="16" color="var(--el-input-icon-color)" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input
                                    ref="passwordRef"
                                    v-model="form.password"
                                    type="password"
                                    :placeholder="t('login.Please input a password')"
                                    show-password
                                >
                                    <template #prefix>
                                        <Icon name="fa fa-unlock-alt" class="form-item-icon" size="16" color="var(--el-input-icon-color)" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="role">
                                <el-select
        ref="roleRef"
        v-model="form.role"
        :placeholder="t('请选择身份')"
        class="w-full"
    >
        <el-option
            v-for="item in roleOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
    </el-select>
                            </el-form-item>
                            <el-checkbox v-model="form.keep" :label="t('login.Hold session')" size="default"></el-checkbox>
                            <el-form-item>
                                <el-button
                                    :loading="state.submitLoading"
                                    class="submit-button"
                                    round
                                    type="primary"
                                    size="large"
                                    @click="onSubmitPre()"
                                >
                                    {{ t('login.Sign in') }}
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, reactive, ref } from 'vue'
import * as pageBubble from '/@/utils/pageBubble'
import type { FormInstance, InputInstance } from 'element-plus'
import { useI18n } from 'vue-i18n'
import { editDefaultLang } from '/@/lang/index'
import { useConfig } from '/@/stores/config'
import { useAdminInfo } from '../stores/adminInfo'
import { login } from '../api/admin'
import { uuid } from '/@/utils/random'
import { buildValidatorData } from '/@/utils/validate'
import router from '/@/router'
import clickCaptcha from '/@/components/clickCaptcha'
import toggleDark from '/@/utils/useDark'
import { baseRoutePath } from '/@/router/static/base'

// 定义定时器变量
let timer: number

// 获取配置和管理员信息的store
const config = useConfig()
const adminInfo = useAdminInfo()
// 设置深色/浅色模式
toggleDark(config.layout.isDark)

// 表单和输入框的引用
const formRef = ref<FormInstance>()
const usernameRef = ref<InputInstance>()
const passwordRef = ref<InputInstance>()
const roleRef = ref<InputInstance>()

// 组件状态
const state = reactive({
    showCaptcha: false,         // 是否显示验证码
    submitLoading: false,       // 提交按钮加载状态
})
// 角色选项
const roleOptions = [
    {
        value: 1,
        label: '学生'
    },
    {
        value: 2,
        label: '企业'
    },
    {
        value: 3,
        label: '员工'
    },
    // 可以根据需要添加更多角色
]

// 表单数据
const form = reactive({
    username: '',              // 用户名
    password: '',             // 密码
    keep: false,              // 是否记住登录状态
    captchaId: uuid(),        // 验证码ID
    captchaInfo: '',          // 验证码信息
    role: '',
})

// 国际化
const { t } = useI18n()

// 表单验证规则
const rules = reactive({
    username: [buildValidatorData({ name: 'required', message: t('login.Please enter an account') })],
    password: [buildValidatorData({ name: 'required', message: t('login.Please input a password') })],
})

// 组件挂载时初始化页面气泡效果
onMounted(() => {
    timer = window.setTimeout(() => {
        pageBubble.init()
    }, 1000)
})

// 组件卸载前清理定时器和事件监听
onBeforeUnmount(() => {
    clearTimeout(timer)
    pageBubble.removeListeners()
})

// 提交前的验证处理
const onSubmitPre = () => {
    formRef.value?.validate((valid) => {
        if (valid) {
            // 如果需要验证码，则显示验证码组件
            if (state.showCaptcha) {
                clickCaptcha(form.captchaId, (captchaInfo: string) => onSubmit(captchaInfo))
            } else {
                onSubmit()
            }
        }
    })
}

// 提交登录请求
const onSubmit = (captchaInfo = '') => {
    state.submitLoading = true
    form.captchaInfo = captchaInfo
    login(form)
        .then((res) => {
            const  access_token  = res.data
            console.log(res.data)
            // 填充管理员信息
            adminInfo.dataFill({
                type: form.username == 'admin' ? 1 : 2,    // 根据用户名判断用户类型
                username: form.username,
                token: access_token,
            })
            // 登录成功后跳转到首页
            router.push({ path: baseRoutePath })
        })
        .finally(() => {
            state.submitLoading = false
        })
}
</script>

<style scoped lang="scss">
.switch-language {
    position: fixed;
    top: 20px;
    right: 20px;
    z-index: 1;
}
.bubble {
    overflow: hidden;
    background: url(/@/assets/bg.jpg) repeat;
}
.form-item-icon {
    height: auto;
}
.login {
    position: absolute;
    top: 0;
    display: flex;
    width: 100vw;
    height: 100vh;
    align-items: center;
    justify-content: center;
    .login-box {
        overflow: hidden;
        width: 430px;
        padding: 0;
        background: var(--ba-bg-color-overlay);
        margin-bottom: 80px;
    }
    .head {
        background: #ccccff;
        img {
            display: block;
            width: 430px;
            margin: 0 auto;
            user-select: none;
        }
    }
    .form {
        position: relative;
        .profile-avatar {
            display: block;
            position: absolute;
            height: 100px;
            width: 100px;
            border-radius: 50%;
            border: 4px solid var(--ba-bg-color-overlay);
            top: -50px;
            right: calc(50% - 50px);
            z-index: 2;
            user-select: none;
        }
        .content {
            padding: 100px 40px 40px 40px;
        }
        .submit-button {
            width: 100%;
            letter-spacing: 2px;
            font-weight: 300;
            margin-top: 15px;
            --el-button-bg-color: var(--el-color-primary);
        }
    }
}

@media screen and (max-width: 720px) {
    .login {
        display: flex;
        align-items: center;
        justify-content: center;
        .login-box {
            width: 340px;
            margin-top: 0;
        }
    }
}
.chang-lang :deep(.el-dropdown-menu__item) {
    justify-content: center;
}
.content :deep(.el-input__prefix) {
    display: flex;
    align-items: center;
}

// 暗黑样式
@at-root .dark {
    .bubble {
        background: url(/@/assets/bg-dark.jpg) repeat;
    }
    .login {
        .login-box {
            background: #161b22;
        }
        .head {
            img {
                filter: brightness(61%);
            }
        }
        .form {
            .submit-button {
                --el-button-bg-color: var(--el-color-primary-light-5);
                --el-button-border-color: rgba(240, 252, 241, 0.1);
            }
        }
    }
}
@media screen and (max-height: 800px) {
    .login .login-box {
        margin-bottom: 0;
    }
}
</style>
