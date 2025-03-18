## 	Jenkins安装kubernetes插件

### jenkins动态slave
Jenkins可以使用静态节点(slave)和动态节点两种方式来构建任务，使用动态节点构建任务可以更好的利用资源，所谓动态节点构建任务就是在kubernetes集群里动态的创建一个pod,在这个pod里面跑我们写的pipeline脚本来构建任务，构建任务完成后这个pod会自动销毁，

### Jenkins安装kubernetes插件
Jenkins使用kubernetes插件主要用于完成两方面的工作：一是用于在kubernetes集群内动态生成一个pod作为Jenkins 的slave节点，提供流水线执行的工作环境；二是用于将应用代码持续部署到kubernetes集群中，我们通过编写pipeline脚本集成kubernetes插件生成动态slave。
<img src="./images/5/plugin-01.png" alt="plugin-01.png" style="zoom:50%;" />
点击可用插件选项
<img src="./images/5/plugin-02.png" alt="plugin-02.png" style="zoom:50%;" />
搜索框输入kubernetes
<img src="./images/5/plugin-04.png" alt="plugin-04.png" style="zoom:50%;" />
选中kubernetes插件然后点击右上角的安装按钮等待安装完成。

### 配置Jenkins连接Kubernetes
jenkins与kubernetes的集成，主要是通过调用Kubernetes的API去kubernetes集群中进行工作的。大多数公司在安装kubernetes集群配置apiserver服务时使用了证书，所以在配置jenkins连接kubernetes集群时，需要根据kubernetes的配置文件生成一系列证书以及key，并将证书上传到Jenkins用来对apiserver进行认证。

#### 部署在非kubernetes集群内Jenkins连接kubernetes配置

安装好插件以后，进入Jenkins首页，点击菜单”Manage Jenkins(系统管理)-–> Clouds(云配置)” 在跳转到的界面中，到右上角，点击” new cloud（新建一个云）–> kubernetes“。

如下所示：

<img src="./images/5/plugin-05.png" alt="plugin-05.png" style="zoom:50%;" />

输入云的名称后点击create按钮跳转到填写详细信息页面，如下所示：

<img src="./images/5/plugin-06.png" alt="plugin-06.png" style="zoom:50%;" />

其中：

名称：这里用于填写要添加的这个云（cloud）的名称，默认为”kubernetes”，如果不想用这个可以自定义。在编写pipeline的时候会用到。

kubernetes 地址：用于填写kubernetes集群的地址，做了多master集群高可用的环境直接写vip地址加端口；只有单个master的环境直接写master加端口地址即可。

Kubernetes 服务证书 key：用于填写与kubernetes集群认证的证书内容。

Kubernetes 命名空间：用于填写调用kubernetes时生成的pod工作的namespace。

Credentials（凭据）：用于连接kubernetes的凭证。

Jenkins 地址：Jenkins的连接地址。

### Kubernetes服务证书key

kubernetes集群安装的时候生成了一系列证书以及key，并且在配置kubernetes中kubectl客户端命令权限的时候，根据这些证书以及key生成了一个kubeconfig文件，用于kubectl与集群通信，这个文件默认为/root/.kube/config文件，对集群有最高操作权限（如果给了cluster-admin权限）。Jenkins需要根据这个文件生成的证书与集群通信，所以我们在生产环境配置Jenkins连接kubernetes集群的时候，需要注意一下kubeconfig文件绑定的用户的权限，最好从新生成一个低权限的kubeconfig文件，而不要用kubectl命令使用的文件。



























​	























