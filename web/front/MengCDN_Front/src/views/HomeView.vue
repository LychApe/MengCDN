<script setup lang="ts">
import axios from "axios";
import {reactive, ref} from "vue";
import router from "@/router";
import {useRoute} from "vue-router";

const route = useRoute()
let getGhPath = <any>reactive([])
let ifGetGhPath
router.push({
  query: {ghPath: <any>getGhPath}
})

if (route.query.ghPath == null) {
  ifGetGhPath = ''
} else {
  ifGetGhPath = route.query.ghPath
}

const owner = router.currentRoute.value.params.owner;
const repo = router.currentRoute.value.params.repo;
const branch = router.currentRoute.value.params.branch;

document.title = "Github文件列表 - " + owner + "/" + repo + "/" + branch;

//列表数据
let ghRawList: any = reactive([])

let autoUrl: any = ref('https://api.github.com/repos/' + owner + '/' + repo + '/contents/' + ifGetGhPath + '?ref=' + branch)

function getGhRawList(autoUrl: string) {
  axios({
    method: 'get',
    url: autoUrl,
  }).then(function (resp) {
    let num = resp.data.length - 1
    //console.log(num)
    ghRawList.value = resp.data

    for (let i = 0; i <= num; i++) {
      if (String(ghRawList.value[i].download_url) != "null") {
        ghRawList.value[i].download_url = String(ghRawList.value[i].download_url).replace('https://raw.githubusercontent.com/', 'https://cdn.fallsoft.cn/gh/')
      }
    }

  });
  getGhPath = route.query.ghPath
}

function setAutoUrl(path: any) {
  autoUrl = 'https://api.github.com/repos/' + owner + '/' + repo + '/contents/' + path + '?ref=' + branch
  getGhRawList(autoUrl)
}

function backAutoUrl() {
  autoUrl = 'https://api.github.com/repos/' + owner + '/' + repo + '/contents/' + '/' + '?ref=' + branch
  getGhRawList(autoUrl)
}

function lookUrl(path: any) {
  window.location.href = "http://cdn.fallsoft.cn/gh/" + owner + "/" + repo + "@" + branch + "/" + path
}

getGhRawList(autoUrl.value)

</script>

<template>
  <a-row>
    <a-col :xs="0" :sm="0" :md="0" :lg="24" :xl="24" :xxl="24">
      <div style="background-color: #f2f3f5">
        <div style="padding: 85px">
          <h1 style="color: #10b981;font-family: CustomFont,sans-serif">FallSoft CDN</h1>
          <h2 style="color: #3c3c43;font-family: CustomFont,sans-serif">一个为各大开发者提供的免费、开放、快速的CDN !</h2>
          <h3 style="color: #6d6d72;font-family: CustomFont,sans-serif">简单易用、高速稳定、免费开放</h3>
        </div>
      </div>
    </a-col>
    <a-col :xs="24" :sm="24" :md="24" :lg="0" :xl="0" :xxl="0">
      <div style="background-color: #f2f3f5">
        <div style="padding: 30px">
          <h1 style="color: #10b981;font-family: CustomFont,sans-serif">FallSoft CDN</h1>
          <h2 style="color: #3c3c43;font-family: CustomFont,sans-serif">一个为各大开发者提供的免费、开放、快速的CDN !</h2>
          <h3 style="color: #6d6d72;font-family: CustomFont,sans-serif">简单易用、高速稳定、免费开放</h3>
        </div>
      </div>
    </a-col>
  </a-row>
  <a-row>
    <a-col :xs="0" :sm="0" :md="0" :lg="24" :xl="24" :xxl="24">
      <div style="padding: 120px">

        <a-card :bordered="false" style="background-color: rgb(242 243 245);border-radius: 9px">
          <a-button @click="backAutoUrl">
            <icon-left-circle/>
          </a-button>
          <a-list>
            <a-list-item v-for="(item ,index) in ghRawList.value">
              <a-list-item-meta
                  :title=item.name
              >
                <template #avatar>
                  <div v-if="item.type === 'dir' " @click="setAutoUrl(item.path)">
                    <a-avatar shape="square">
                      <icon-folder/>
                    </a-avatar>
                  </div>
                  <div v-if="item.type === 'file' " @click="lookUrl(item.path)">
                    <a-avatar shape="square">
                      <icon-file/>
                    </a-avatar>
                  </div>
                </template>
              </a-list-item-meta>
              <template #actions>
                <div v-if="item.type === 'file' ">
                  <a :href=item.html_url style="margin-left: 5px;text-decoration: none;color: #1d2129;">
                    <icon-code-square/>
                  </a>
                </div>
              </template>
            </a-list-item>
          </a-list>

        </a-card>
      </div>
    </a-col>
    <a-col :xs="24" :sm="24" :md="24" :lg="0" :xl="0" :xxl="0">
      <div style="padding: 10px">

        <a-card :bordered="false" style="background-color: rgb(242 243 245);border-radius: 9px">
          <a-button @click="backAutoUrl">
            <icon-left-circle/>
          </a-button>
          <a-list>
            <a-list-item v-for="(item ,index) in ghRawList.value">
              <a-list-item-meta
                  :title=item.name
              >
                <template #avatar>
                  <div v-if="item.type === 'dir' " @click="setAutoUrl(item.path)">
                    <a-avatar shape="square">
                      <icon-folder/>
                    </a-avatar>
                  </div>
                  <div v-if="item.type === 'file' " @click="lookUrl(item.path)">
                    <a-avatar shape="square">
                      <icon-file/>
                    </a-avatar>
                  </div>
                </template>
              </a-list-item-meta>
              <template #actions>
                <div v-if="item.type === 'file' ">
                  <a :href=item.html_url style="margin-left: 5px;text-decoration: none;color: #1d2129;">
                    <icon-code-square/>
                  </a>
                </div>
              </template>
            </a-list-item>
          </a-list>

        </a-card>
      </div>
    </a-col>
    <a-col :xs="0" :sm="0" :md="0" :lg="0" :xl="0" :xxl="0">
    </a-col>
    <a-col :xs="0" :sm="0" :md="0" :lg="0" :xl="0" :xxl="0">

    </a-col>
  </a-row>


</template>

<style scoped>
@font-face {
  font-family: CustomFont;
  src: url(https://js.fallsoft.cn/LXGWWenKai-Regular.ttf);
}

h1 h2 h3 h4 {
  font-family: CustomFont, sans-serif;
}
</style>