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
- 添加Chart仓库
```shell
helm repo add cita-cloud-operator-proxy https://cita-cloud.github.io/operator-proxy
```
- 创建安装`operator-proxy`的命名空间
```shell
kubectl create ns cita
```
- 执行安装
```shell
helm install cita-cloud-operator-proxy cita-cloud-operator-proxy/cita-cloud-operator-proxy -n=cita
```
- 验证安装
```
kubectl get pod -ncita | grep cita-cloud-operator-proxy
```
如果实际输出的`STATUS`状态是`Running`, 表示`operator-proxy`已经成功安装。

### `cco-cli`的安装
- 最新版本安装
```shell
curl -sLS https://raw.githubusercontent.com/cita-cloud/operator-proxy/master/install-cli.sh | bash
```
- 手动安装指定版本
可根据运行平台下载对应版本的二进制至可执行目录下：<https://github.com/cita-cloud/operator-proxy/releases>

- 设置环境变量(以`Linux`为例)
获得`operator-proxy`暴露的`NodePort`端口(示例中为30194)
```shell
kubectl describe service cita-cloud-operator-proxy -ncita | grep NodePort
```
设置环境变量，以`Kubernetes`集群的任意一个节点`IP`和`NodePort`端口作为`endpoint`
```shell
export OPERATOR_PROXY_ENDPOINT=192.168.10.120:30194
```

- 执行帮助命令
```shell
cco-cli -h
The cita-cloud operator command line interface lets you create and manage CITA-CLOUD chain.

Usage:
  cco-cli [command]
```

示例
----

### 创建一条名为`test-chain`的链(默认3节点，并创建在名为`cita`的环境变量中)
该命令是多个子命令的集合
```shell
cco-cli all-in-one create test-chain
```
可以看到有3个`Pod`被创建出来
```shell
kubectl get pod -ncita
NAME                                          READY   STATUS    RESTARTS   AGE
test-chain-9ba8b85938c7-0                     6/6     Running   0          18h
test-chain-af41db9a3064-0                     6/6     Running   0          18h
test-chain-efe071517f54-0                     6/6     Running   0          18h
```

### chain command
```shell
$ cco-cli chain -h
Chain related commands

Usage:
  cco-cli chain [command]

Available Commands:
  delete      Delete a chain in the k8s cluster
  describe    Show chain detail in the k8s cluster
  init        Initialize a chain into the k8s cluster
  list        List chain in the k8s cluster
  online      Online a chain into the k8s cluster
```

- 初始化一条名为`test-chain`的链
```shell
$ cco-cli chain init test-chain
init chain [cita/test-chain] success
```
- 此时，这条链的状态为`Publicizing`，代表该链的信息需要向各参与方公示，可以通过`describe`命令查看详情
```shell
$ cco-cli chain describe test-chain
Chain Base Info:
+-----------------+--------------------------------------------------------------------+
|      FIELD      |                               VALUE                                |
+-----------------+--------------------------------------------------------------------+
| Name            | test-chain                                                         |
| Namespace       | cita                                                               |
| Id              | 63586a3c0255f337c77a777ff54f0040b8c388da04f23ecee6bfd4953a6512b4   |
| Timestamp       | 1644466132459760                                                   |
| PrevHash        | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| BlockInterval   | 3                                                                  |
| BlockLimit      | 100                                                                |
| EnableTls       | false                                                              |
| ConsensusType   | Raft                                                               |
| NetworkImage    | citacloud/network_p2p:v6.3.0                                       |
| ConsensusImage  | citacloud/consensus_raft:v6.3.0                                    |
| ExecutorImage   | citacloud/executor_evm:v6.3.0                                      |
| StorageImage    | citacloud/storage_rocksdb:v6.3.0                                   |
| ControllerImage | citacloud/controller:v6.3.0                                        |
| KmsImage        | citacloud/kms_sm:v6.3.0                                            |
| Status          | Publicizing                                                             |
+-----------------+--------------------------------------------------------------------+
Admin Account:
+-------+-----------+------------+-------+--------+
| NAME  | NAMESPACE |   CHAIN    | ROLE  | DOMAIN |
+-------+-----------+------------+-------+--------+
| admin |   cita    | test-chain | Admin |        |
+-------+-----------+------------+-------+--------+
Node Info:
+--------+-----------+------------+---------+------+---------+
|  NAME  | NAMESPACE |   CHAIN    | ACCOUNT | SIZE | STATUS  |
+--------+-----------+------------+---------+------+---------+
| node-1 |   cita    | test-chain |  alice  | 10Gi | Running |
| node-2 |   cita    | test-chain |   bob   | 10Gi | Running |
| node-3 |   cita    | test-chain | carlos  | 10Gi | Running |
+--------+-----------+------------+---------+------+---------+
```
- 经各方确认通过后，上线这条链，上线前需要创建好`Admin`账户和共识节点账户，参考[account command](#account command)
```shell
$ cco-cli chain online test-chain
online chain [cita/test-chain] success
```
- 查看一条链的详情
```shell
$ cco-cli chain describe test-chain
Chain Base Info:
+-----------------+--------------------------------------------------------------------+
|      FIELD      |                               VALUE                                |
+-----------------+--------------------------------------------------------------------+
| Name            | test-chain                                                         |
| Namespace       | cita                                                               |
| Id              | 63586a3c0255f337c77a777ff54f0040b8c388da04f23ecee6bfd4953a6512b4   |
| Timestamp       | 1644466132459760                                                   |
| PrevHash        | 0x0000000000000000000000000000000000000000000000000000000000000000 |
| BlockInterval   | 3                                                                  |
| BlockLimit      | 100                                                                |
| EnableTls       | false                                                              |
| ConsensusType   | Raft                                                               |
| NetworkImage    | citacloud/network_p2p:v6.3.0                                       |
| ConsensusImage  | citacloud/consensus_raft:v6.3.0                                    |
| ExecutorImage   | citacloud/executor_evm:v6.3.0                                      |
| StorageImage    | citacloud/storage_rocksdb:v6.3.0                                   |
| ControllerImage | citacloud/controller:v6.3.0                                        |
| KmsImage        | citacloud/kms_sm:v6.3.0                                            |
| Status          | Online                                                             |
+-----------------+--------------------------------------------------------------------+
Admin Account:
+-------+-----------+------------+-------+--------+
| NAME  | NAMESPACE |   CHAIN    | ROLE  | DOMAIN |
+-------+-----------+------------+-------+--------+
| admin |   cita    | test-chain | Admin |        |
+-------+-----------+------------+-------+--------+
Node Info:
+--------+-----------+------------+---------+------+---------+
|  NAME  | NAMESPACE |   CHAIN    | ACCOUNT | SIZE | STATUS  |
+--------+-----------+------------+---------+------+---------+
| node-1 |   cita    | test-chain |  alice  | 10Gi | Running |
| node-2 |   cita    | test-chain |   bob   | 10Gi | Running |
| node-3 |   cita    | test-chain | carlos  | 10Gi | Running |
+--------+-----------+------------+---------+------+---------+
```
- 列出命名空间下的所有链
```shell
$ cco-cli chain list -n cita
+------------+-----------+--------+
|    NAME    | NAMESPACE | STATUS |
+------------+-----------+--------+
| test-chain |   cita    | Online |
+------------+-----------+--------+
```
- 删除一条链
```shell
$ cco-cli chain delete test-chain -n cita
delete chain [cita/test-chain] success
```

### account command
```shell
$ cco-cli account -h
Account related commands

Usage:
  cco-cli account [command]

Available Commands:
  create      Create a node account for chain
  list        List node account in the k8s cluster
```
- 创建`Admin`账户: `admin`
```shell
$ cco-cli account create admin --chain test-chain --kmsPassword 123456 --role Admin
create account [cita/admin] success
```
若创建时指定`address`字段，则不会生成新的`admin`账户地址，链的配置会将此地址作为`admin`账户的地址
- 创建共识账户: `alice`，若链的网络选择`tls`，则必须加上`domain`参数
```shell
$ cco-cli account create alice --chain test-chain --kmsPassword 123456 --role Consensus --domain alice.cita.com
create account [cita/alice] success
```
- 创建普通账户: `davis`，若链的网络选择`tls`，则必须加上`domain`参数
```shell
$ cco-cli account create davis --chain test-chain --kmsPassword 123456 --role Ordinary --domain davis.cita.com
create account [cita/davis] success
```
- 查看命名空间下所有用户
```shell
$ cco-cli account list -n cita
+--------+-----------+------------+-----------+-----------------+
|  NAME  | NAMESPACE |   CHAIN    |   ROLE    |   DOMAIN        |
+--------+-----------+------------+-----------+-----------------+
| admin  |   cita    | test-chain |   Admin   |                 |
| alice  |   cita    | test-chain | Consensus | alice.cita.com  |
|  bob   |   cita    | test-chain | Consensus |  bob.cita.com   |
| carlos |   cita    | test-chain | Consensus | carlos.cita.com |
| davis  |   cita    | test-chain | Ordinary  | davis.cita.com  |
+--------+-----------+------------+-----------+-----------------+
```

### node command
```shell
$ cco-cli node -h
Node related commands

Usage:
  cco-cli node [command]

Available Commands:
  delete      Delete a node
  init        Init a node for chain
  list        List node in the k8s cluster
  reload      Reload the node config, usually used to add or delete nodes in a chain
  start       Start a node
  stop        Stop a node
```
- 初始化链下的一个节点：`node1`，需匹配对应的链名和账户
可初始化与共识账户对应的节点数量
```shell
$ cco-cli node init node-1 --account alice --chain test-chain --storageClassName nas-client-provisioner --storageSize 10737418240
init node [cita/node-1] success
```
- 启动对应的各个节点
```shell
$ cco-cli node start node-1
start node [cita/node-1] success
```
其他共识节点同上
- 列出对应链下的所有节点
```shell
$ cco-cli node list --chain test-chain
+--------+-----------+------------+---------+------+---------+
|  NAME  | NAMESPACE |   CHAIN    | ACCOUNT | SIZE | STATUS  |
+--------+-----------+------------+---------+------+---------+
| node-2 |   cita    | test-chain |   bob   | 10Gi | Running |
| node-3 |   cita    | test-chain | carlos  | 10Gi | Running |
| node-1 |   cita    | test-chain |  alice  | 10Gi | Running |
+--------+-----------+------------+---------+------+---------+
```
- 新增普通节点`node-4`，对应账户为`davis`
```shell
# 初始化
cco-cli node init node-4 --account davis --chain test-chain
# 启动
cco-cli node start node-4
```
新增节点后，原有节点均需要执行`reload`操作，以便能与新节点进行网络交互
```shell
$ cco-cli node reload node-1
reload node [cita/node-1] success
$ cco-cli node reload node-2
reload node [cita/node-2] success
$ cco-cli node reload node-3
reload node [cita/node-3] success
$ cco-cli node reload node-4
reload node [cita/node-4] success
```