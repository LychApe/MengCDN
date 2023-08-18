<script lang="ts" setup>
import {Icon} from '@arco-design/web-vue';
import router from "@/router";
import { emitter } from "@/plugin/BusJs/bus";
import {onBeforeMount, ref} from "vue";

const IconFont = Icon.addFromIconFontCn({src: '//at.alicdn.com/t/c/font_4197143_zzqrjqtplq.js'});
const selectedKeys = ref();

onBeforeMount(() => {
  emitter.on("getSelectedKeys", (v: any) => {
    selectedKeys.value = v.value;
  });
});
</script>

<template>
  <div style="height: 100vh;">
    <a-layout class="layout-demo">
      <a-layout-sider>
        <div class="logo">
          <h3 style="padding: 7px 46px;">MengCDN</h3>
        </div>
        <a-menu
            :default-selected-keys="['1']"
            :selected-keys="selectedKeys"
            :style="{ width: '100%' }"
        >
          <a-menu-item key="1" @click="router.push('/')">
            <icon-dashboard/>
            仪表盘
          </a-menu-item>
          <a-menu-item key="2" @click="router.push('/Github')">
            <icon-font type="icon-github"/>
            Github 设置
          </a-menu-item>
          <a-menu-item key="3" @click="router.push('/NPM')">
            <icon-font type="icon-npm"/>
            NPM 设置
          </a-menu-item>
          <a-menu-item key="4" @click="router.push('/WordPressPlugin')">
            <icon-font type="icon-wordpress"/>
            WPplugin 设置
          </a-menu-item>
          <a-menu-item key="5" @click="router.push('/WordPressTheme')">
            <icon-font type="icon-wordpress"/>
            WPtheme 设置
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <RouterView></RouterView>
    </a-layout>
  </div>
</template>


<style scoped>
.layout-demo {
  height: -webkit-fill-available;
  background: var(--color-fill-2);
  border: 1px solid var(--color-border);
}

.layout-demo :deep(.arco-layout-sider) .logo {
  height: 32px;
  margin: 12px 8px;
  background: rgba(255, 255, 255, 0.2);
}

.layout-demo :deep(.arco-layout-sider-light) .logo {
  background: var(--color-fill-2);
}

.layout-demo :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout-demo :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
  border-radius: 13px;
}
</style>