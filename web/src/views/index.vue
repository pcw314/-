<template>
    <div class="default-main">
        <div class="banner">
            <el-row :gutter="10">
                <el-col :md="24" :lg="18">
                    <div class="welcome suspension">
                        <img class="welcome-img" :src="headerSvg" alt="" />
                        <div class="welcome-text">
                            <div class="welcome-title">{{ adminInfo.username + t('utils.comma') + getGreet() }}</div>
                            <div class="welcome-note">{{ state.remark }}</div>
                        </div>
                    </div>
                </el-col>
                <el-col :lg="6" class="hidden-md-and-down">
                    <div class="working">
                        <img class="working-coffee" :src="coffeeSvg" alt="" />
                        <div class="working-text">
                            {{ t('index.You have worked today') }}<span class="time">{{ state.workingTimeFormat }}</span>
                        </div>
                        <div @click="onChangeWorkState()" class="working-opt working-rest">
                            {{ state.pauseWork ? t('index.Continue to work') : t('index.have a bit of rest') }}
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
        <div class="small-panel-box">
            <el-row :gutter="20">
                <el-col :sm="12" :lg="6">
                    <div class="small-panel user-reg suspension">
                        <div class="small-panel-title">星陨官群</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>承接定制</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel file suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel users suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel addons suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
        <div class="small-panel-box">
            <el-row :gutter="20">
                <el-col :sm="12" :lg="6">
                    <div class="small-panel file suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel users suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel addons suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel addons suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
        <div class="small-panel-box">
            <el-row :gutter="20">
                <el-col :sm="12" :lg="6">
                    <div class="small-panel file suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel users suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel addons suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :sm="12" :lg="6">
                    <div class="small-panel addons suspension">
                        <div class="small-panel-title">{{ t('index.Advertising space') }}</div>
                        <div class="small-panel-content">
                            <div class="content-left">
                                <Icon color="#8595F4" size="20" name="fa fa-cloud" />
                                <span>{{ t('index.Advertising space') }}</span>
                            </div>
                            <div class="content-right">
                                <el-link type="primary" href="https://qm.qq.com/q/tqglsof3a2" target="_blank">加群</el-link>
                            </div>
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, reactive, onActivated, onBeforeMount } from 'vue'
import headerSvg from '/@/assets/index/header-1.svg'
import coffeeSvg from '/@/assets/index/coffee.svg'
// import { index } from '../api/dashboard'
import { useI18n } from 'vue-i18n'
import { Local } from '/@/utils/storage'
import { useAdminInfo } from '../stores/adminInfo'
import { WORKING_TIME } from '/@/stores/constant/cacheKey'
import { getGreet } from '/@/utils/common'
let workTimer: number

defineOptions({
    name: 'index',
})

const d = new Date()
const { t } = useI18n()
const adminInfo = useAdminInfo()

const state: {
    charts: any[]
    remark: string
    workingTimeFormat: string
    pauseWork: boolean
} = reactive({
    charts: [],
    remark: t('index.welcome'),
    workingTimeFormat: '',
    pauseWork: false,
})

const onChangeWorkState = () => {
    const time = parseInt((new Date().getTime() / 1000).toString())
    const workingTime = Local.get(WORKING_TIME)
    if (state.pauseWork) {
        // 继续工作
        workingTime.pauseTime += time - workingTime.startPauseTime
        workingTime.startPauseTime = 0
        Local.set(WORKING_TIME, workingTime)
        state.pauseWork = false
        startWork()
    } else {
        // 暂停工作
        workingTime.startPauseTime = time
        Local.set(WORKING_TIME, workingTime)
        clearInterval(workTimer)
        state.pauseWork = true
    }
}

const startWork = () => {
    const workingTime = Local.get(WORKING_TIME) || { date: '', startTime: 0, pauseTime: 0, startPauseTime: 0 }
    const currentDate = d.getFullYear() + '-' + (d.getMonth() + 1) + '-' + d.getDate()
    const time = parseInt((new Date().getTime() / 1000).toString())

    if (workingTime.date != currentDate) {
        workingTime.date = currentDate
        workingTime.startTime = time
        workingTime.pauseTime = workingTime.startPauseTime = 0
        Local.set(WORKING_TIME, workingTime)
    }

    let startPauseTime = 0
    if (workingTime.startPauseTime <= 0) {
        state.pauseWork = false
        startPauseTime = 0
    } else {
        state.pauseWork = true
        startPauseTime = time - workingTime.startPauseTime // 已暂停时间
    }

    let workingSeconds = time - workingTime.startTime - workingTime.pauseTime - startPauseTime

    state.workingTimeFormat = formatSeconds(workingSeconds)
    if (!state.pauseWork) {
        workTimer = window.setInterval(() => {
            workingSeconds++
            state.workingTimeFormat = formatSeconds(workingSeconds)
        }, 1000)
    }
}

const formatSeconds = (seconds: number) => {
    var secondTime = 0 // 秒
    var minuteTime = 0 // 分
    var hourTime = 0 // 小时
    var dayTime = 0 // 天
    var result = ''

    if (seconds < 60) {
        secondTime = seconds
    } else {
        // 获取分钟，除以60取整数，得到整数分钟
        minuteTime = Math.floor(seconds / 60)
        // 获取秒数，秒数取佘，得到整数秒数
        secondTime = Math.floor(seconds % 60)
        // 如果分钟大于60，将分钟转换成小时
        if (minuteTime >= 60) {
            // 获取小时，获取分钟除以60，得到整数小时
            hourTime = Math.floor(minuteTime / 60)
            // 获取小时后取佘的分，获取分钟除以60取佘的分
            minuteTime = Math.floor(minuteTime % 60)
            if (hourTime >= 24) {
                // 获取天数， 获取小时除以24，得到整数天
                dayTime = Math.floor(hourTime / 24)
                // 获取小时后取余小时，获取分钟除以24取余的分；
                hourTime = Math.floor(hourTime % 24)
            }
        }
    }
    result =
        hourTime +
        t('index.hour') +
        ((minuteTime >= 10 ? minuteTime : '0' + minuteTime) + t('index.minute')) +
        ((secondTime >= 10 ? secondTime : '0' + secondTime) + t('index.second'))
    if (dayTime > 0) {
        result = dayTime + t('index.day') + result
    }
    return result
}

onActivated(() => {})

onMounted(() => {
    startWork()
})

onBeforeMount(() => {
    for (const key in state.charts) {
        state.charts[key].dispose()
    }
})

onUnmounted(() => {
    clearInterval(workTimer)
})
</script>

<style scoped lang="scss">
.banner {
    margin-bottom: 20px;
}
.welcome {
    background: #e1eaf9;
    border-radius: 6px;
    display: flex;
    align-items: center;
    padding: 15px 20px !important;
    box-shadow: 0 0 30px 0 rgba(82, 63, 105, 0.05);
    .welcome-img {
        height: 100px;
        margin-right: 10px;
        user-select: none;
    }
    .welcome-title {
        font-size: 1.5rem;
        line-height: 30px;
        color: var(--ba-color-primary-light);
    }
    .welcome-note {
        padding-top: 6px;
        font-size: 15px;
        color: var(--el-text-color-primary);
    }
}
.working {
    height: 130px;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    height: 100%;
    position: relative;
    &:hover {
        .working-coffee {
            -webkit-transform: translateY(-4px) scale(1.02);
            -moz-transform: translateY(-4px) scale(1.02);
            -ms-transform: translateY(-4px) scale(1.02);
            -o-transform: translateY(-4px) scale(1.02);
            transform: translateY(-4px) scale(1.02);
            z-index: 999;
        }
    }
    .working-coffee {
        transition: all 0.3s ease;
        width: 80px;
    }
    .working-text {
        display: block;
        width: 100%;
        font-size: 15px;
        text-align: center;
        color: var(--el-text-color-primary);
    }
    .working-opt {
        position: absolute;
        top: -40px;
        right: 10px;
        background-color: rgba($color: #000000, $alpha: 0.3);
        padding: 10px 20px;
        border-radius: 20px;
        color: var(--ba-bg-color-overlay);
        transition: all 0.3s ease;
        cursor: pointer;
        opacity: 0;
        z-index: 999;
        &:active {
            background-color: rgba($color: #000000, $alpha: 0.6);
        }
    }
    &:hover {
        .working-opt {
            opacity: 1;
            top: 0;
        }
        .working-done {
            opacity: 1;
            top: 50px;
        }
    }
}
// .small-panel-box {
//     margin-top: 20px;
// }
.small-panel {
    background-color: #e9edf2;
    border-radius: var(--el-border-radius-base);
    padding: 25px;
    margin-bottom: 20px;
    .small-panel-title {
        color: #92969a;
        font-size: 15px;
    }
    .small-panel-content {
        display: flex;
        align-items: flex-end;
        margin-top: 20px;
        color: #2c3f5d;
        .content-left {
            font-size: 24px;
            .icon {
                margin-right: 10px;
            }
            span {
                display: inline-block;
                font-size: 28px;
            }
        }
        .content-right {
            font-size: 18px;
            margin-left: auto;
        }
        .color-success {
            color: var(--el-color-success);
        }
        .color-warning {
            color: var(--el-color-warning);
        }
        .color-danger {
            color: var(--el-color-danger);
        }
        .color-info {
            color: var(--el-text-color-secondary);
        }
    }
}
.growth-chart {
    margin-bottom: 20px;
}
.user-growth-chart,
.file-growth-chart {
    height: 260px;
}
.new-user-growth {
    height: 300px;
}

.user-source-chart,
.user-surname-chart {
    height: 400px;
}
.new-user-item {
    display: flex;
    align-items: center;
    padding: 20px;
    margin: 10px 15px;
    box-shadow: 0 0 30px 0 rgba(82, 63, 105, 0.05);
    background-color: var(--ba-bg-color-overlay);
    .new-user-avatar {
        height: 48px;
        width: 48px;
        border-radius: 50%;
    }
    .new-user-base {
        margin-left: 10px;
        color: #2c3f5d;
        .new-user-name {
            font-size: 15px;
        }
        .new-user-time {
            font-size: 13px;
        }
    }
    .new-user-arrow {
        margin-left: auto;
    }
}
.new-user-card :deep(.el-card__body) {
    padding: 0;
}

@media screen and (max-width: 425px) {
    .welcome-img {
        display: none;
    }
}
@media screen and (max-width: 1200px) {
    .lg-mb-20 {
        margin-bottom: 20px;
    }
}
html.dark {
    .welcome {
        background-color: var(--ba-bg-color-overlay);
    }
    .small-panel {
        background-color: var(--ba-bg-color-overlay);
        .small-panel-content {
            color: var(--el-text-color-regular);
        }
    }
    .new-user-item {
        .new-user-base {
            color: var(--el-text-color-regular);
        }
    }
}
</style>
