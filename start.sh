#!/bin/bash
echo ----------------start----------------
#定义镜像名
image_name="my_test_app"
#定义容器名
container_name="my_test_app"
#先删除容器
docker rm -f ${container_name}
#再删除镜像
docker rmi -f ${image_name}
#执行Dockerfile文件，生成镜像
docker build . -t ${image_name}
#生成容器并启动
docker run -p 8888:8888 -d --name=${container_name} ${image_name}
#查看容器是否启动成功
docker ps


