<template>
  <section>
    <h1 @click="open=!open">
      <!-- 主体 -->
    </h1>
  </section>
</template>

<script setup lang="ts">
  import { inject } from 'vue'

  const open=inject('sideBarIsOpen')
</script>