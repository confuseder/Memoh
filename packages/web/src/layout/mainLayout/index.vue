<template>
  <section class="flex">
    <sidebar-provider>
      <slot name="sidebar" />
      <slot name="main" />
    </sidebar-provider>
  </section>
</template>
<script setup lang="ts">
import { SidebarProvider } from '@memoh/ui'

</script>
