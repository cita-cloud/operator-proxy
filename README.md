简介
----
本项目包括`cita-cloud-operator`的代理服务端`operator-proxy`和客户端`cco-cli`( cita-cloud-operator command-line interface ) 两部分
- `cita-cloud-operator`: 一个管理`cita cloud`链的生命周期的自定义控制器
- `operator-proxy`: `cita-cloud-operator` 的服务端代理程序，通过RPC暴露接口 
- `cco-cli`: 连接`operator-proxy`的客户端工具，可通过命令的方式快速创建链和节点

安装
----
### `cita-cloud-operator`的安装
`cita-cloud-operator`的`Helm`安装方式建该链接: <https://github.com/cita-cloud/cita-cloud-operator>

### `operator-proxy`的安装
- 安装Chart仓库
```
helm repo add cita-cloud-operator-proxy https://cita-cloud.github.io/operator-proxy
```
- 创建安装`operator-proxy`的命名空间
```
kubectl create ns cita
```
- 执行安装
```
helm install cita-cloud-operator-proxy cita-cloud-operator-proxy/cita-cloud-operator-proxy -n=cita
```
- 验证安装
```
kubectl get pod -ncita | grep cita-cloud-operator-proxy
```
如果实际输出的`STATUS`状态是`Running`, 表示`operator-proxy`已经成功安装。

### `cco-cli`的安装
- 下载地址
可根据运行平台下载对应的二进制程序：<https://github.com/cita-cloud/operator-proxy/releases>

- 设置环境变量(以`Linux`为例)
获得`operator-proxy`暴露的`NodePort`端口(示例中为30194)
```
kubectl describe service cita-cloud-operator-proxy -ncita | grep NodePort
```
设置环境变量，以`Kubernetes`集群的任意一个节点`IP`和`NodePort`端口作为`endpoint`
```
export OPERATOR_PROXY_ENDPOINT=192.168.10.120:30194
```

- 执行帮助命令
```
cco-cli -h
The cita-cloud operator command line interface lets you create and manage CITA-CLOUD chain.

Usage:
  cco-cli [command]
```

示例
----

### 创建一条名为`test-chain`的链(默认3节点)
```
cco-cli all-in-one create test-chain
```