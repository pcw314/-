<template>
    <div class="default-main ba-table-box">
        <el-alert class="ba-table-alert" v-if="baTable.table.remark" :title="baTable.table.remark" type="info" show-icon />

        <!-- 表格顶部菜单 -->
        <TableHeader
            :buttons="['refresh', 'add', 'edit', 'delete', 'comSearch', 'columnDisplay']"
            :quick-search-placeholder="t('Quick search placeholder', { fields: t('soft.name') })"
        />

        <!-- 表格 -->
        <!-- 要使用`el-table`组件原有的属性，直接加在Table标签上即可 -->
        <Table ref="tableRef" />

        <!-- 表单 -->
        <PopupForm />
    </div>
</template>

<script setup lang="ts">
import { ref, provide } from 'vue'
import baTableClass from '/@/utils/baTable'
import PopupForm from './popupForm.vue'
import Table from '/@/components/table/index.vue'
import TableHeader from '/@/components/table/header/index.vue'
import { defaultOptButtons } from '/@/components/table'
import { baTableApi } from '/@/api/common'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'soft',
})

const { t } = useI18n()
const tableRef = ref()
const baTable = new baTableClass(
    new baTableApi('/soft'),
    {
        column: [
            { type: 'selection', align: 'center', operator: false },
            { label: t('id'), prop: 'id', align: 'center', operator: '=', operatorPlaceholder: t('id'), width: 70 },
            {
                label: t('soft.name'),
                prop: 'name',
                align: 'center',
                operator: 'LIKE',
                operatorPlaceholder: t('soft.name'),
                render: 'tag',
            },
            {
                label: t('type'),
                prop: 'type',
                align: 'center',
                render: 'tag',
                custom: { 2: 'success', 1: 'danger' },
                replaceValue: { 1: t('charge'), 2: t('free') },
                width: 100,
            },
            {
                label: t('uuid'),
                prop: 'uuid',
                align: 'center',
                width: 330,
                operator: 'LIKE',
                operatorPlaceholder: t('uuid'),
            },
            {
                label: t('secret'),
                prop: 'secret',
                align: 'center',
                width: 160,
                operator: 'LIKE',
                operatorPlaceholder: t('secret'),
            },
            {
                label: t('heart'),
                prop: 'heart',
                align: 'center',
                operator: '=',
                operatorPlaceholder: t('heart'),
            },
            {
                label: t('Create Time'),
                prop: 'created_at',
                align: 'center',
                render: 'datetime',
                width: 160,
                sortable: 'custom',
                operator: 'RANGE',
            },
            {
                label: t('state'),
                prop: 'state',
                align: 'center',
                render: 'tag',
                custom: { 1: 'success', 0: 'danger' },
                replaceValue: { 1: t('enable'), 0: t('disable') },
                width: 100,
            },
            {
                label: t('operate'),
                align: 'center',
                width: 100,
                render: 'buttons',
                buttons: defaultOptButtons(['edit', 'delete']),
                operator: false,
            },
        ],
        dblClickNotEditColumn: [undefined],
    },
    {
        defaultItems: {
            type: 2,
            state: 1,
        },
    }
)

baTable.mount()
baTable.getIndex()

provide('baTable', baTable)
</script>

<style scoped lang="scss"></style>
