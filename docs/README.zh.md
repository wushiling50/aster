<div align="center">
  <h1 style="display: inline-block; vertical-align: middle;">Aster</h1>
</div>

# Overview
Aster 是一个 GitHub 用户分析工具，利用 GitHub 开源项目数据评估开发者的技术水平。通过分析项目影响力和开发者贡献度，生成生成开发者评分，并推测开发者的国籍和专长。

# Project Structure
```bash
├── go.mod
├── go.sum
├── LICENSE
├── Makefile        # 一些 make 命令
├── README.md
├── api             # gateway/api 服务
├── config          # 配置文件
├── docker          # Docker 相关部署文件
├── gen             # goctl rpc protoc 生成的代码
├── idl             # 接口定义文件
├── pkg             # 各类工具以及实用函数
├── rpc             # 微服务的实现
└── script          # 脚本文件
```
# Quick Start
## 预备操作
添加 config/config.yaml 文件，配置示例请参考 config.example.yaml。

## 本地部署
### 项目启动
```sh
# 构筑服务镜像
make aster-build-all
# 启动环境基础容器
make env-up
# 启动服务容器
make aster-run-all
```
### 项目关闭
```sh
# 关闭环境基础容器
make env-down
# 关闭服务容器
make aster-remove-all
```

## 项目设计
具体内容请参考: [Aster 设计文档](https://west2-online.feishu.cn/wiki/CY0cwiHZAiPiFSkwuoac8VrEnJg)

## LICENSE
本项目使用 Apache-2.0 License 开源，具体请见 [LICENSE](./LICENSE) 文件。
