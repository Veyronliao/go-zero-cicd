## 	Jenkins安装kubernetes插件

### jenkins动态slave
Jenkins可以使用静态节点(slave)和动态节点两种方式来构建任务，使用动态节点构建任务可以更好的利用资源，所谓动态节点构建任务就是在kubernetes集群里动态的创建一个pod,在这个pod里面跑我们写的pipeline脚本来构建任务，构建任务完成后这个pod会自动销毁，

### Jenkins安装kubernetes插件
Jenkins使用kubernetes插件主要用于完成两方面的工作：一是用于在kubernetes集群内动态生成一个pod作为Jenkins 的slave节点，提供流水线执行的工作环境；二是用于将应用代码持续部署到kubernetes集群中，我们通过编写pipeline脚本集成kubernetes插件生成动态slave。
<img src="./images/5/plugin-01.png" alt="plugin-01.png" style="zoom:50%;" />
<img src="./images/5/plugin-02.png" alt="plugin-02.png" style="zoom:50%;" />
点击可用插件选项
<img src="./images/5/plugin-02.png" alt="plugin-02.png" style="zoom:50%;" />
搜索框输入kubernetes
<img src="./images/5/plugin-04.png" alt="plugin-04.png" style="zoom:50%;" />




































​	























