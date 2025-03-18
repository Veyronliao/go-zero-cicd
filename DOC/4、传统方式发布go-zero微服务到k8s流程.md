## 	传统发布go-zero微服务到K8s的流程

在完成go-zero微服务程序编码后我们需要将代码发布到k8s测试环境进行测试，在没有gitlab和Jenkins实现cicd的情况下，手动发布流程如下：

1、上传代码到k8s服务器节点;
2、在服务器节点上使用goctl命令创建api、rpc服务的dockerfile文件；
   ```shell
   $ goctl docker -go user.go
   ```
3、复制dockerfile文件到项目跟目录，使用docker或者conternerd生成镜像文件，生成镜像命令：
   ```shell
   $ docker build -t bolog-user-rpc:v1 .
   ```
4、将镜像打标签：
   ```shell
   $ docker tag bolog-user-rpc:v1 veyron-bolog-user-rpc:v1
   ```
5、将镜像保存为tar包：
   ```shell
   $ docker save veyron-bolog-user-rpc -o  veyron-bolog-rpc.tar
   ```
6、将保存的tar包复制到k8s的各个node节点上：
   ```shell
   scp veyron-bolog-rpc.tar root@master01:/home/
   ```
7、在每个k8s节点服务器上将镜像导入到containnerd中:
   ```shell
   ctr -n k8s.io i import veyron-bolog-rpc.tar
   ```
   查看镜像包是否成功导入到containerd中：
   ```shell
   ctr -n k8s.io images list | grep veyron-bolog-user-rpc:v1
   ```

8、在k8s集群创建serviceAccount账号，要具有k8s内部的endpoint资源的list、watch、get权限；
```shell
   #创建账号
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: go-zero-bolog
  name: find-endpoints

---
#创建角色对应操作
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: discov-endpoints
rules:
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["get","list","watch"]

---
#给账号绑定角色
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: find-endpoints-discov-endpoints
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: discov-endpoints
subjects:
- kind: ServiceAccount
  name: find-endpoints
  namespace: go-zero-bolog
```
9、使用goctl生成api、rpc的k8s yaml部署文件，并指定serviceAccount，命令：
   生成api的yaml文件：
   ```shell
   goctl kube deploy -replicas 3 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name bolog-user-rpc -namespace go-zero-bolog -image veyron-bolog-user-rpc:v2 -o veyron-bolog-user-rpc.yaml -port 15010 --serviceAccount find-endpoints
   ```
   生成rpc的yaml文件：
   ```shell
   goctl kube deploy -nodePort 32000 -replicas 3 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name veyron-bolog-user-api  -namespace go-zero-bolog -image veyron-bolog-user-api:v2 -o veyron-bolog-user-api.yaml -port 8888 --serviceAccount find-endpoints
   ```

10、部署（网关配置nodeport）
11、测试访问

注意：由于我的K8s的runtim是containerd,所以需要在每个node节点上导入镜像到containerd中，如果使用的是旧版本的k8s，runtime是docker则不需要执行导入镜像这一步
总结：从上面的发布流程来看，在没有cicd的条件下发布go-zero微服务到k8s是一件重复繁琐的事情，而且很容易出错，一旦k8s的节点多了，手动发布对发布程序的人来说将是一个灾难，因此实现cicd是非常必要的，下一章我们将实现cicd。


























​	























