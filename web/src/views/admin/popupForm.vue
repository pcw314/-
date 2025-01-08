<template>
    <!-- 对话框表单 -->
    <el-dialog
        class="ba-operate-dialog"
        :close-on-click-modal="false"
        :destroy-on-close="true"
        :model-value="['add', 'edit'].includes(baTable.form.operate!)"
        @close="baTable.toggleForm"
    >
        <template #header>
            <div class="title" v-drag="['.ba-operate-dialog', '.el-dialog__header']" v-zoom="'.ba-operate-dialog'">
                {{ baTable.form.operate ? t(baTable.form.operate) : '' }}
            </div>
        </template>
        <el-scrollbar v-loading="baTable.form.loading" class="ba-table-form-scrollbar">
            <div
                class="ba-operate-form"
                :class="'ba-' + baTable.form.operate + '-form'"
                :style="config.layout.shrink ? '' : 'width: calc(100% - ' + baTable.form.labelWidth! / 2 + 'px)'"
            >
                <el-form
                    ref="formRef"
                    @keyup.enter="baTable.onSubmit(formRef)"
                    :model="baTable.form.items"
                    :label-position="config.layout.shrink ? 'top' : 'right'"
                    :label-width="baTable.form.labelWidth + 'px'"
                    :rules="rules"
                    v-if="!baTable.form.loading"
                >
                    <el-form-item prop="soft_ids" :label="t('soft.name')">
                        <el-select
                            class="w100"
                            multiple
                            filterable
                            clearable
                            v-model="baTable.form.items!.soft_ids"
                            :placeholder="t('Please select field', { field: t('soft.name') })"
                        >
                            <el-option v-for="item in softOptions" :key="item.id" :label="item.name" :value="item.id" />
                        </el-select>
                    </el-form-item>
                    <el-form-item prop="username" :label="t('username')">
                        <el-input
                            v-model="baTable.form.items!.username"
                            type="string"
                            :placeholder="t('Please input field', { field: t('username') })"
                        ></el-input>
                    </el-form-item>
                    <FormItem
                        :label="t('type')"
                        v-model.number="baTable.form.items!.type"
                        type="radio"
                        :data="{ content: { 1: t('Super Admin'), 2: t('Soft Admin') }, childrenAttr: { border: true } }"
                    />
                    <el-form-item prop="password" :label="t('password')">
                        <el-input
                            v-model="baTable.form.items!.password"
                            type="string"
                            :placeholder="t('Please input field', { field: t('password') })"
                        ></el-input>
                    </el-form-item>
                    <el-form-item prop="qq" :label="t('qq')">
                        <el-input
                            v-model="baTable.form.items!.qq"
                            type="string"
                            :placeholder="t('Please input field', { field: t('qq') })"
                        ></el-input>
                    </el-form-item>
                    <FormItem
                        :label="t('state')"
                        v-model.number="baTable.form.items!.state"
                        type="radio"
                        :data="{ content: { 1: t('enable'), 0: t('disable') }, childrenAttr: { border: true } }"
                    />
                </el-form>
            </div>
        </el-scrollbar>
        <template #footer>
            <div :style="'width: calc(100% - ' + baTable.form.labelWidth! / 1.8 + 'px)'">
                <el-button
                    v-blur
                    :loading="baTable.form.submitLoading"
                    @click="
                        () => {
                            baTable.form.items!.soft_ids = baTable.form.items!.soft_ids ? baTable.form.items!.soft_ids.toString() : ''
                            baTable.onSubmit(formRef)
                        }
                    "
                    type="primary"
                >
                    {{ baTable.form.operateIds && baTable.form.operateIds.length > 1 ? t('Save and edit next item') : t('save') }}
                </el-button>
                <el-button @click="baTable.toggleForm('')">{{ t('cancel') }}</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive, inject, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import type baTableClass from '/@/utils/baTable'
import type { FormInstance, FormItemRule } from 'element-plus'
import FormItem from '/@/components/formItem/index.vue'
import { buildValidatorData } from '/@/utils/validate'
import { useConfig } from '/@/stores/config'
import { getSoftList } from '/@/api/soft'

const config = useConfig()
const formRef = ref<FormInstance>()
const baTable = inject('baTable') as baTableClass

const { t } = useI18n()

const rules: Partial<Record<string, FormItemRule[]>> = reactive({
    username: [buildValidatorData({ name: 'required', title: t('username') })],
    password: [buildValidatorData({ name: 'required', title: t('password') })],
})

const softOptions = ref<anyObj>([])
onMounted(() => {
    getSoftList({}).then((res) => {
        softOptions.value = res.data
    })
})

// 由于select是多选，要将ids转换成数组
watch(
    () => baTable.form.items,
    (newVal) => {
        baTable.form.items!.soft_ids = newVal!.soft_ids ? newVal!.soft_ids.split(',').map((item: string) => Number(item)) : []
    }
)
</script>

<style scoped lang="scss"></style>
