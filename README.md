# Covid-19 疫情信息导航网站

## 前言

2020年的春节，疫情的阴影笼罩全国，随着确诊人数的增长，恐慌在人群中蔓延，然而在这种时刻，依然有人挺身逆行，毅然投身最严重的疫区！

疫情没有边界，技术亦无边界，作为一名普通的开发者，我们致力于用技术改变世界，受能力职业所限，我们不能前往疫情一线战斗，但我相信我们仍然可以为这个世界做些什么。

恐慌源于猜疑，造谣源自无知！为了让获取疫情信息更加便利，信息流转的效率更高，我做了一个简单的疫情信息导航网站，网站上线一天后，我在网络上看到了很多和我一样的开发者或者其他岗位的个人做出的努力，2020武汉加油、手工整理百余家医院物资求助信息、医疗物资供应平台等等都是大家努力的产物！

疫情终将消散，明天一定更好！向夜以继日奋战在疫情防控一线的天使和勇士们致敬！天佑中华！武汉加油！

<!--more-->

## 项目简介

疫情信息导航旨在为大家提供一个便利而全面的疫情信息入口，整理了涵盖疫情信息的各个方面的网络通道，目前计划提供的信息入口有

- 实时疫情
- 防疫指南
- 物资救援
- 举报通道
- 在线义症
- 谣言鉴定
- 发热门诊
- 病患寻人（未上线）

近期将持续更新添加其他模块


## 技术架构

### 前端部署
前端部分仅需静态部署public目录下的静态文件即可，为保证网站的高可用，线上服务采用了无服务器(Serveless)架构中最基本的一种模式，网站部署于云服务商的对象存储服务上，后端接口使用go服务实现高性能代理， 前期使用了AWS的API Gateway及Lambda提供服务支持

![](http://image.werty.cn/source_blog/20200127200841.png)

相关的架构介绍可以参考这篇文章：[AWS - Serverless 和 lambda](https://blog.51cto.com/beanxyz/2348953)

### 后端服务
最新版本后端服务采用go语言编写，部署接口时需要安装go语言环境(1.10以上版本)及相关依赖，然后将项目目录放入GOPATH路径下，完成后进入项目目录后执行
```shell script
go build
# 假设项目目录为nCov-API
nohup ./nCov-API &
```
也可直接使用部署好的开放接口： [开放接口文档]()
> 注意：
> 原有http接口已升级为https协议 使用go语言直接配置的https服务，实现http与https共存较为复杂，且会接口性能有一定影响，因此不做共存支持。
> 同时，目前https证书是直接配置到代码中的，实测https对比http本地QPS下降90%以上，因此性能要求较高的最好还是走http
>
> http接口端口为2020 https端口为2021 如需使用https请及时替换
```
# 企业维护更稳定
## 最新总数统计
http://ncov-api.werty.cn:2020/latest/tx/
## 含省份数据 相关资讯
http://ncov-api.werty.cn:2020/original/tx/

#源自github开源项目 个人维护 偶尔不稳定
http://ncov-api.werty.cn:2020/latest/dxy/
http://ncov-api.werty.cn:2020/original/dxy/

#患者同行查询接口 结果数据分页  分页参数 page 页码， limit 每页条数
http://ncov-api.werty.cn:2020/together/

#疫情信息导航数据接口
http://ncov-api.werty.cn:2020/nav/app

```


###

## 线上入口：

[http://nav.werty.cn](http://nav.werty.cn)

## 友情链接

[2020 援助武汉](https://wuhan2020.github.io/#)

[2020 援助武汉 Github项目地址](https://github.com/wuhan2020/wuhan2020)

[2019新型冠状病毒疫情实时爬虫 Github项目地址](https://github.com/BlankerL/DXY-Covid-19-Crawler) 首页实时疫情数据基于该项目提供接口，使用APIGateway代理实现跨域及缓存功能

## 许可

基于MIT许可，同时如果直接使用本项目建站，请在页脚添加Github项目入口或源站[nav.werty.cn](http://nav.werty.cn)地址

本站基于[webstack.cc](https://webstack.cc)开源项目开发

## 其他

在疫情结束前，我会尽力添加更多的优质信息入口到疫情导航，如果你有优质的信息来源，或者其他建议，欢迎在[issue](https://github.com/wertycn/nCoV/issues)中反馈，也可联系邮箱[debugicu@163.com](mailto:debugicu@163.com)
