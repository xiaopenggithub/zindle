zindle(go-zero)从极速开发到完美上线 企业级微服务架构落地实战 [讲解演示](https://www.bilibili.com/video/BV1Z54y1V7Ze/)



## 1.前言

zindle：基于go-zero开发的bookstore "kindle"

极速开发：丰富的工具支持:goctl 除了生成标准化代码框架外，还包括生成k8s部署脚本，从开发到部署，直接用goctl。

企业级：框架支持日活千万，适合大部分中小型企业的业务需求了。

![image-20210524161004117](https://user-images.githubusercontent.com/20268389/119350328-c5f0a000-bcd1-11eb-8ee8-56137177e1c9.png)

本系统的客户端有web,App(ios,android)

1. 开发工具
   1. goland
   2. vscode
   3. hbuilder
   4. App模拟器：网易Mumu(如果发布App Store，需要苹果开发者证书)

2. 环境准备
   1. 运行环境
      - [ ] mysql
      - [ ] redis
      - [ ] etcd
      - [ ] kubernetes1.20.1
   2. 前端开发
      - [ ] nodejs
   3. 后端开发
      - [ ] golang

开发环境要求，大家可以参考官方文档:https://go-zero.dev

## 2.系统架构图

|                                                              |
| ------------------------------------------------------------ |
| ![image-20210524175216643](https://user-images.githubusercontent.com/20268389/119350501-005a3d00-bcd2-11eb-80a7-1bc83d1dd8c8.png) |



## 3.读者端App原型

https://org.modao.cc/app/aa11e8c03777a7653d18145be11b7f3e70eb1a13#screen=skow3pw1nx0atg6



## 4.代码目录说明

```
├── code-dir
│   ├── app  // app代码
│   ├── backend // 后台接口，rpc
│   ├── backendweb // 后台vue页面代码
│   ├── script // 数据库脚本，简化的kubernetes部署脚本
│   ├── .gitignore // git控制忽略文件
│   ├── LICENSE // LICENSE文件，使用的是MIT LICENSE

```

## 5.本地代码运行

1. goland(编写go代码)
2. vscode(编写vue代码)
3. hbuilder(编写app代码)
4. App模拟器：网易Mumu

### 5.1运行相关命令详细说明

请提前准备数据库，并导入script目录中的sql

1. app直接用hbuilder打开
   1. 开启模拟器调度
2. backend用goland打开
   1. k8s
   2. backend-rpc
   3. backend-api
   4. bookstore
3. backendweb用vscode打开
   1. 使用cnpm install命令安装相关依赖包

## 6.后台演示



## 7.用户App端演示



## 8.部署上线说明

我的线上k8s环境：1master,2node，k8s版本:v1.20.4

部署脚本已经简化，去除了健康检查相关声明配置。



## 9.后续开发

1. 接入ES提高图书检索效率
2. 精简代码，对系统不断优化改进



## 10.总结

本系统zindle是使用go-zero开发的一个完整的全端系统，基于官方的bookstore案例情景开发。后台功能包括权限角色管理、菜单管理等企业系统必备模块，在此基础上，很方便的进行业务开发。

### 10.1. 知识储备

1. 前端：vue

2. 后端：golang

3. 部署：kubernetes

4. 一些架构思想，统筹全局

### 10.2.本系统参考文献

参考到的开源库都非常优秀，欢迎大家下载本地体验并给作者star以支持

1. go-zero微服务框架 https://github.com/zeromicro/go-zero
   
   系统内核,前后端接口使用go-zero制作


2. gin-vue-admin 全栈开发框架 https://github.com/flipped-aurora/gin-vue-admin

   后台界面基于此源码库制作

3. ElementUI vue相关组件 https://element.eleme.cn
   
   后台界面相关组件参考此源码库官方文档

### 10.3. 个人wx（备注zindle）
![image-20210524151130130](https://user-images.githubusercontent.com/20268389/119350257-b3766680-bcd1-11eb-8c20-56ad40a03803.png)


### 10.4. zindle交流群（不要闲聊）
![zindle交流群](https://user-images.githubusercontent.com/20268389/120781070-49748180-c55b-11eb-8607-3a6bbbd066a4.jpeg)


### 10.5. 欢迎star

如果本系统能给你带来帮助，麻烦给我一个star。
