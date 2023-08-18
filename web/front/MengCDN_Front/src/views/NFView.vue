<script setup lang="ts">
document.title = "查询错误"
</script>

<template>
  <link rel="stylesheet" href="https://unpkg.com/mdui@1.0.2/dist/css/mdui.min.css"/>

  <div class="mdui-container">
    <br/>
    <br/>
    <div style="padding: 50px; background: #fafafa; border-radius: 5px">
      <h1 style="font-weight: 300; font-family: CustomFont, sans-serif;">Error : M100003</h1>
      <h3 style="font-weight: 350; font-family: CustomFont, sans-serif;">查询错误</h3>
      <br/>
      <h2 style="font-family: CustomFont, sans-serif;">什么问题?</h2>
      <p style="font-family: CustomFont, sans-serif;">查询路径信息错误</p>
      <br/>
      <h2 style="font-family: CustomFont, sans-serif;">如何解决?</h2>
      <p style="font-family: CustomFont, sans-serif;">请检查你的路径信息。<br/>正确方法: /browseGH/gh/用户名/仓库@分支?ghPath=/要查询的仓库内目录</p>
      <br/>
      <hr style="border-radius: 30px;border-style: dashed;"/>
      <h5 style="font-weight: 450; font-family: CustomFont, sans-serif; ">MengCDN V1.0.230728<br/>Fallsoft | Lychape</h5>
    </div>
  </div>
</template>

<style>
@font-face {
  font-family: CustomFont;
  src: url(https://js.fallsoft.cn/LXGWWenKai-Regular.ttf);
}
</style>

