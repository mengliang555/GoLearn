# [基础概念](https://www.runoob.com/docker/docker-architecture.html)

- 镜像(Image): 一个 root 文件系统
- 容器:(Container):容器是镜像运行时的实体
- 仓库(Repository):用于存储镜像的相关配置

## For Docker

- Client:Docker 客户端通过命令行或者其他工具使用 Docker SDK (https://docs.docker.com/develop/sdk/) 与 Docker 的守护进程通信。
- Host:一个物理或者虚拟的机器用于执行 Docker 守护进程和容器。
- Docker Registry:Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库,[Docker Hub](https://hub.docker.com) 提供了庞大的镜像集合供使用
- 