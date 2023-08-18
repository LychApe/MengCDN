<script lang="ts" setup>
import {reactive} from 'vue';
import axios from "axios";
import {Message, Notification} from "@arco-design/web-vue";
import router from "@/router";

window.document.title = "FallSoft CDN Admin"

const form = reactive({
  user: '',
  password: '',
});

const handleSubmit = () => {

  axios({
    method: 'post',
    url: '/api/v1/login',
    data: form
  }).then(res => {
    if (res.data.status == 200) {
      window.sessionStorage.setItem("token", res.data.token)
      Message.success("登陆成功!")
      router.push("/")
    } else {
      return Message.error(res.data.message)
    }
  })
};

</script>

<template>
  <a-row>

    <a-col flex="700px">
      <div style="height: 100vh;background: #222529;">
        <div class="Login-R-1">
          <h1 style="color: #ecf0f1;">MengCDN | Admin</h1>
        </div>
      </div>
    </a-col>

    <a-col flex="auto">
      <div class="Login-L-0">
        <div class="Login-L-1">
          <a-card :bordered="false" style="border-radius: 10px;background-color: rgb(50 50 50 / 54%);">
            <a-form :model="form" style="width: 400px;" @submit="handleSubmit">
              <a-form-item field="user" label="Username">
                <a-input v-model="form.user"/>
              </a-form-item>
              <a-form-item field="password" label="Password">
                <a-input type="password" v-model="form.password"/>
              </a-form-item>
              <a-form-item>
                <a-button html-type="submit">Login</a-button>
              </a-form-item>
            </a-form>
          </a-card>
        </div>
      </div>
    </a-col>

  </a-row>
</template>

<style>
.Login-R-1 {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
}

.Login-L-0 {
  background: url("https://2b.chat/img/ambient.jpg") no-repeat 50% transparent;
  background-size: cover;
  height: 100vh;
}

.Login-L-1 {
  display: flex;
  height: 100%;
  justify-content: center;
  align-items: center;
}
</style>
