<script setup lang="ts">
// import type { Payment } from '@/components/columns'
import { watch, h, computed } from 'vue'
import CreateModel from '@/components/CreateModel/index.vue'
import { useQuery } from '@pinia/colada'

import DataTable from '@/components/DataTable/index.vue'
import request from '@/utils/request'
import {type ColumnDef } from '@tanstack/vue-table'


interface ModelType {
  apiKey:string,
  baseUrl: string,
  clientType: 'OpenAI'|'Anthropic'|'Google',
  modelId:string,
  name:string,
  type:'chat'|'embedding'
}

const columns:ColumnDef<ModelType>[] = [
  {
    accessorKey: 'modelId',
    header: () => h('div', { class: 'text-left py-4' }, 'Name'),
    cell({row}) {
      return h('div',{ class: 'text-left py-4' },row.getValue('modelId'))
    }
  },
  {
    accessorKey: 'baseUrl',
    header: () => h('div', { class: 'text-left' }, 'Base Url'),

  },

  {
    accessorKey: 'apiKey',
    header: () => h('div', { class: 'text-left' }, 'Api Key'),
  },
  {
    accessorKey: 'clientType',
    header: () => h('div', { class: 'text-left' }, 'Client Type'),
  },
  {
    accessorKey: 'Name',
    header: () => h('div', { class: 'text-left' }, 'Name'),
  },
  {
    accessorKey: 'type',
    header: () => h('div', { class: 'text-left' }, 'Type'),
  }
]

const {data:modelData}=useQuery({
  key: ['models'],
  query() {
    return request({
      url: '/model'
    })
  }
})
const displayFormat = computed(() => {
  return modelData.value?.data?.items?.map((currentModel:{model: ModelType,id:'string' })=>currentModel.model)??[]
})

</script>

<template>
  <div class="w-full py-10 mx-auto">
    <div class="flex mb-4">
      <CreateModel />
    </div>
    <DataTable
      :columns="columns"
      :data="displayFormat"
    />
  </div>
</template>