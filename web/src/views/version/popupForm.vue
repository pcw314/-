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
                    <el-form-item prop="soft_id" :label="t('soft.name')">
                        <el-select
                            class="w100"
                            filterable
                            v-model="baTable.form.items!.soft_id"
                            :placeholder="t('Please input field', { field: t('soft.name') })"
                        >
                            <el-option v-for="item in softOptions" :key="item.id" :label="item.name" :value="item.id" />
                        </el-select>
                    </el-form-item>
                    <el-form-item prop="name" :label="t('version.name')">
                        <el-input
                            v-model="baTable.form.items!.name"
                            type="string"
                            :placeholder="t('Please input field', { field: t('version.name') })"
                        ></el-input>
                    </el-form-item>
                    <FormItem :label="t('log')" v-model="baTable.form.items!.log" type="textarea" />
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
                <el-button v-blur :loading="baTable.form.submitLoading" @click="baTable.onSubmit(formRef)" type="primary">
                    {{ baTable.form.operateIds && baTable.form.operateIds.length > 1 ? t('Save and edit next item') : t('save') }}
                </el-button>
                <el-button @click="baTable.toggleForm('')">{{ t('cancel') }}</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive, inject } from 'vue'
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
    name: [buildValidatorData({ name: 'required', title: t('version.name') })],
    soft_id: [buildValidatorData({ name: 'required', message: t('Please select field', { field: t('soft.name') }) })],
})

const softOptions = ref<anyObj>([])
onMounted(() => {
    getSoftList({}).then((res) => {
        softOptions.value = res.data
    })
})
</script>

<style scoped lang="scss"></style>
