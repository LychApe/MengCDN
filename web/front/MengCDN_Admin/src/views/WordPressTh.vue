<script setup lang="ts">
import {emitter} from "@/plugin/BusJs/bus";
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import axios from "axios";

let SelectedKeys: any = ref(["5"]);
emitter.emit("getSelectedKeys", SelectedKeys);
window.document.title = "MengCDN - WordPress|theme 设置";

let form = reactive({
  whiteList: "",
  modelSW: false,
  whiteListSW: false,
  msw1: "",
  msw2: ""
});

const data = reactive([{
  label: '模块名称',
  value: 'WordPress 加速模块',
}, {
  label: '模块版本',
  value: 'V1.230801',
}]);

const apiInfoData = [{
  label: '文件加速接口',
  value: '/wp/theme/仓库/版本号/文件或目录',
}];

axios.post(
    "/api/v1/cdnWL/WPth"
).then(res => {
  if (res.data.status !== 200) {
    return window.alert("异常")
  } else {
    form.whiteList = res.data.data
  }
})

axios.post(
    "/api/v1/cdnSW/0/WPth"
).then(res => {
  if (res.status !== 200) {
    return window.alert("异常")
  } else {
    if (res.data.status == 200) {
      form.modelSW = true
      form.msw1 = "1"
    } else {
      form.modelSW = false
      form.msw1 = "0"
    }
  }
})

axios.post(
    "/api/v1/cdnSW/1/WPth"
).then(res => {
  if (res.status !== 200) {
    return window.alert("异常")
  } else {
    if (res.data.status == 200) {
      form.whiteListSW = true
      form.msw2 = "1"
    } else {
      form.whiteListSW = false
      form.msw2 = "0"
    }
  }
})

let handleSubmit = () => {
  // 请求拦截器
  axios.interceptors.request.use((config: any) => {
    // 在发送请求之前做些什么
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem(
        "token"
    )}`;
    return config;
  });
  axios.put('/api/v1/cdnWL/WPth', {
    "value": form.whiteList
  }).then((res) => {
    if (res.data.status == 200) {
      Message.success(res.data.message)
    } else {
      Message.success(res.data.message)
    }
  })
};

let AXSW01 = () => {
  // 请求拦截器
  axios.interceptors.request.use((config: any) => {
    // 在发送请求之前做些什么
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem(
        "token"
    )}`;
    return config;
  });
  let msw11
  if (form.msw1 == "0") {
    msw11 = "1"
  } else {
    msw11 = "0"
  }
  axios.put('/api/v1/cdnSW/0/WPth/' + msw11, {
    "value": form.whiteList
  }).then((res) => {
    if (res.data.status == 200) {
      Message.success(res.data.message)
    } else {
      Message.error(res.data.message)
    }
  })
};

let AXSW02 = () => {
  // 请求拦截器
  axios.interceptors.request.use((config: any) => {
    // 在发送请求之前做些什么
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem(
        "token"
    )}`;
    return config;
  });
  let msw22
  if (form.msw2 == "0") {
    msw22 = "1"
  } else {
    msw22 = "0"
  }
  axios.put('/api/v1/cdnSW/1/WPth/' + msw22, {
    "value": form.whiteList
  }).then((res) => {
    if (res.data.status == 200) {
      Message.success(res.data.message)
    } else {
      Message.error(res.data.message)
    }
  })
};
</script>

<template>
  <a-layout>
    <a-layout-header style="padding-left: 20px;">
      WordPress|theme 设置中心
    </a-layout-header>

    <a-layout style="padding: 24px;">
      <a-layout-content>
        <div style="margin: 15px;">

          <a-space direction="vertical" size="large" fill>
            <a-descriptions :data="data" title="基础信息" :column="{xs:1, md:3, lg:5}">
              <a-descriptions-item v-for="item of data" :label="item.label">
                <a-tag>{{ item.value }}</a-tag>
              </a-descriptions-item>
            </a-descriptions>
            <a-descriptions :data="apiInfoData" title=接口信息 bordered/>
          </a-space>

          <br/><br/><br/><br/>

          <a-form :model="form" @submit="handleSubmit">
            <a-row>
              <a-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12" :xxl="12">
                <a-form-item field="whiteList" tooltip="仓库" label="白名单">
                  <a-mention v-model="form.whiteList" type="textarea" style="height: 450px"/>
                </a-form-item>
                <a-form-item>
                  <a-button html-type="submit">Submit</a-button>
                </a-form-item>
              </a-col>
              <a-col :xs="24" :sm="24" :md="24" :lg="1" :xl="1" :xxl="1"></a-col>
              <a-col :xs="24" :sm="24" :md="24" :lg="10" :xl="10" :xxl="10">
                <a-form-item field="modelSW" label="模块开关: ">
                  <a-switch v-model="form.modelSW" type="round" @change=AXSW01>
                    <template #checked>
                      ON
                    </template>
                    <template #unchecked>
                      OFF
                    </template>
                  </a-switch>
                </a-form-item>
                <br/>
                <a-form-item field="whiteListSW" label="白名单开关:">
                  <a-switch v-model="form.whiteListSW" type="round" @change=AXSW02>
                    <template #checked>
                      ON
                    </template>
                    <template #unchecked>
                      OFF
                    </template>
                  </a-switch>

                </a-form-item>
              </a-col>
            </a-row>

          </a-form>

        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style scoped>

</style>