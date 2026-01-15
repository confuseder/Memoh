<template>
  <aside>
    <SidebarProvider :open="open as boolean">
      <Sidebar>
        <SidebarHeader>
          <img
            src="../../../public/logo.png"
            width="80"
            class="m-auto"
            alt="logo.png"
          >
          <h4 class="scroll-m-20 text-xl font-semibold tracking-tight text-center text-muted-foreground">
            Memoh
          </h4>
          <!-- <SidebarMenu>
            <SidebarMenuItem>
              <SidebarMenuButton size="lg">
                <div
                  class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
                >
                  <GalleryVerticalEnd class="size-4" />
                </div>
                <div class="grid flex-1 text-left text-sm leading-tight">
                  <span class="truncate font-semibold">Acme Inc</span>
                  <span class="truncate text-xs">Enterprise</span>
                </div>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu> -->
        </SidebarHeader>
        <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>
              对话操作
            </SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                <SidebarMenuItem
                  v-for="sidebarItem in sidebarInfo"
                  :key="sidebarItem.title"
                >
                  <CollapsibleTrigger as-child>
                    <SidebarMenuButton>
                      <svg-icon
                        type="mdi"
                        :path="sidebarItem.icon"
                      />
                      <span>{{ sidebarItem.title }}</span>
                    </SidebarMenuButton>
                  </CollapsibleTrigger>
                </SidebarMenuItem>
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>
        <SidebarFooter />
        <SidebarRail />
      </Sidebar>
    </SidebarProvider>
  </aside>
</template>
<script setup lang="ts">
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,  
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail, 
  
} from '@memoh/ui'
import { reactive, inject } from 'vue'
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiRobot, mdiChatOutline, mdiCogBox, mdiListBox } from '@mdi/js'


const open=inject('sideBarIsOpen')

const sidebarInfo = reactive([{
  title: '创建对话',
  path: '/',
  icon: mdiChatOutline
}, {
  title: '模型配置',
  path: '/',
  icon: mdiRobot
  }, {
  title: '环境设置',
  path: '/',
  icon: mdiCogBox
  }, {
  title: '调度规则',
  path: '/',
  icon: mdiListBox
}])
</script>