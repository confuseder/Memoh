<template>
  <aside class="[&_[data-state=collapsed]_.title-container]:hidden">
    <Sidebar collapsible="icon">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <img
              src="../../../public/logo.png"
              width="80"
              class="m-auto"
              alt="logo.png"
            >
            <h4
              class="scroll-m-20 text-xl font-semibold tracking-tight text-center text-muted-foreground title-container"
            >
              Memoh
            </h4>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>
            对话操作
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <Collapsible
                v-for="sidebarItem in sidebarInfo"
                :key="sidebarItem.title"
                class="group/collapsible"
              >
                <SidebarMenuItem>
                  <CollapsibleTrigger as-child>
                    <SidebarMenuButton
                      :tooltip="sidebarItem.title"
                      @click="router.push({ name: sidebarItem.name })"
                    >
                      <svg-icon
                        type="mdi"
                        :path="sidebarItem.icon"
                      />
                      <span>{{ sidebarItem.title }}</span>
                    </SidebarMenuButton>
                  </CollapsibleTrigger>
                </SidebarMenuItem>
              </Collapsible>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarRail />
    </Sidebar>
  </aside>
</template>
<script setup lang="ts">
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
  CollapsibleTrigger,
  Collapsible

} from '@memoh/ui'
import { reactive } from 'vue'
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiRobot, mdiChatOutline, mdiCogBox, mdiListBox, mdiHome, mdiBookArrowDown } from '@mdi/js'
import { useRouter } from 'vue-router'

const router = useRouter()

const sidebarInfo = reactive([{
  title: '创建对话',
  name: 'chat',
  icon: mdiChatOutline
}, {
  title: '主页',
  name: 'home',
  icon: mdiHome
},
{
  title: '模型配置',
  name: 'models',
  icon: mdiRobot
}, {
  title: '设置',
  name: 'settings',
  icon: mdiCogBox
}, {
  title: 'MCP',
  name: 'mcp',
  icon: mdiListBox
}, {
  title: '平台',
  name: 'platform',
  icon: mdiBookArrowDown
}])
</script>