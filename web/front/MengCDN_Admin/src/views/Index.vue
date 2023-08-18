<script setup lang="ts">
import {emitter} from "@/plugin/BusJs/bus";
import {ref} from "vue";

let SelectedKeys: any = ref(["1"]);
emitter.emit("getSelectedKeys", SelectedKeys);
window.document.title = "MengCDN - 控制台";

const apiInfoData0 = [{
  label: '文件加速接口',
  value: '/gh/用户名/仓库@分支/文件或目录',
}, {
  label: '目录列表接口',
  value: '/browseGH/gh/用户名/仓库@分支',
}];
const apiInfoData1 = [{
  label: '文件加速接口',
  value: '/npm/仓库/版本号/文件',
}, {
  label: '目录列表接口',
  value: '/npm/browse/仓库/版本号/文件或目录',
}];
const apiInfoData2 = [{
  label: 'plugin加速接口',
  value: '/wp/plugin/仓库/文件或目录',
}, {
  label: 'theme加速接口',
  value: '/wp/theme/仓库/版本号/文件或目录',
}];
</script>

<template>
  <a-layout>
    <a-layout-header style="padding-left: 20px;">
    </a-layout-header>

    <a-layout style="padding: 24px;">
      <a-layout-content>
        <div style="margin: 15px;">

          <a-space direction="vertical" size="large" fill>
            <a-descriptions :data="apiInfoData0" title=Github接口信息 bordered/>
            <a-descriptions :data="apiInfoData1" title=NPM接口信息 bordered/>
            <a-descriptions :data="apiInfoData2" title=WordPress接口信息 bordered/>
          </a-space>


        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>


<style scoped>

</style>