<template>
  <SidebarInset>
    <SidebarTrigger class="-ml-1" />
    <h1 @click="open = !open">
      <!-- 主体 -->
    </h1>
  </SidebarInset>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import { SidebarTrigger, SidebarInset } from '@memoh/ui'

const open = inject('sideBarIsOpen')
</script>