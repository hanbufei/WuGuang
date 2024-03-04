# 后端
![img.png](../res/Icon.png)

## 简介
雾光笔记是一款基于"简洁、高效、方便"为目的而开发的个人知识管理系统。 
本项目为其后端项目。后端以本地文件系统为依托，使用wg后缀作为笔记文件格式。
采用golang + goframe搭建。

## 开发指南
### 编译
`go build`
### or 交叉编译
`GOOS=windows go build -ldflags -H=windowsgui`
### 部署
将生成的可执行文件放入launcher程序同目录