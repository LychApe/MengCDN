import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "FallSoft CDN",
  description: "A fast and free CDN for developers",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '主页', link: '/' },
      { text: '介绍', link: '/introduction' }
    ],

    sidebar: [
      {
        text: '目录',
        items: [
          { text: '介绍', link: '/introduction' },
          { text: '快速开始', link: '/start' }
        ]
      }
    ],

    footer: {
      copyright: 'Copyright © 2021-Present FallSoft',
      message: '<a href="http://beian.miit.gov.cn/" target="_blank" rel="nofollow">豫ICP备2022026413号-2</a>',
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
