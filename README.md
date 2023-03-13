# Introduction

这是 [DIYgod/RSSHub](https://github.com/DIYgod/RSSHub) 的 Go 版本

项目基于 [GoFrame web 框架](https://github.com/gogf/gf) 实现，框架文档 [GoFrame (ZH)](https://itician.org/pages/viewpage.action?pageId=1114119)

目前正在迁移 [DIYgod/RSSHub](https://github.com/DIYgod/RSSHub) 已适配的网站


# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.