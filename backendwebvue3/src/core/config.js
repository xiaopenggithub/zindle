/**
 * 网站配置文件
 */

const config = {
  appName: 'Gin-Vue-Admin',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Zindle，开源地址：https://github.com/xiaopenggithub/zindle`
      )
    )
    console.log('\n')
  }
}

export default config
