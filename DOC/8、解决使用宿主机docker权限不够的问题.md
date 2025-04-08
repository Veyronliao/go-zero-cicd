## 	解决使用宿主机docker权限不够的问题

### 概述
在前面章节中我们在使用宿主机的docker的时候会报权限不足的错误，如下所示：
```shell
permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post "http://%2Fvar%2Frun%2Fdocker.sock/v1.45/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=bolog-rpc-user%3A57c2bf8&target=&ulimits=%5B%5D&version=1": dial unix /var/run/docker.sock: connect: permission denied
```
### 解决办法
经过一阵摸索发现是jnlp是使用jenkins用户来运行pod的，所以在访问宿主机docker的时候会出现权限不足的情况，我们只需要指定containner的运行用户即可解决问题，我们只需要通过runAsUser属性来指定使用root用户来运行container即可，具体脚本如下：

```shell
def POD_LABEL = "jenkinspod-${UUID.randomUUID().toString()}"
podTemplate(label: POD_LABEL, cloud: 'veyron-k8s', serviceAccount: 'default',containers: [
    containerTemplate(
    name: 'jnlp', 
    image:'jenkins/inbound-agent:latest', 
    args: '${computer.jnlpmac} ${computer.name}',
    ttyEnabled: true,
    privileged: true,
    alwaysPullImage: false),
    containerTemplate(name: 'goctl', image: 'kevinwan/goctl', ttyEnabled: true, command: 'cat',alwaysPullImage: false),
    containerTemplate(name: 'veyron', image: '192.168.211.150:8077/veyron/jnlp-agent:latest', ttyEnabled: true,runAsUser : "0", command: 'cat',alwaysPullImage: false)
  ],volumes: [
        hostPathVolume(hostPath: '/usr/bin/docker', mountPath:'/usr/bin/docker'),
        hostPathVolume(hostPath: '/var/run/docker.sock', mountPath:'/var/run/docker.sock'),
        hostPathVolume(hostPath: '/etc/docker/daemon.json', mountPath:'/etc/docker/daemon.json')
    ]) {
    node(POD_LABEL) {
        stage('拉取GitLab仓库代码') {
            container('jnlp'){
                checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[credentialsId: 'd837f2f8-745e-44b7-b078-c650ac4bd7e8', url: 'http://192.168.211.150/liaoweilong/go-zero-bolog.git']])
                env.commit_id = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                sh 'echo ${commit_id}'
                sh 'echo 构建服务类型：${JOB_NAME}-$service'
            }
        }
        stage('编译') {
            container('goctl') {
                sh 'echo goctl版本检测'
                sh 'goctl --version'
                sh 'rm User/Rpc/Dockerfile -rf'
                sh 'echo 使用goctl生成dockerfile'
                sh 'cd User/Rpc/ && goctl docker -go ${service}.go && ls -l'
            }
        }
        stage('打包镜像') {
            container('veyron') {
                sh 'whoami'
                sh 'echo docker --version'
                sh 'docker --version'
                env.image = sh(returnStdout: true, script: 'echo ${JOB_NAME}-${service}:${commit_id}').trim()
                sh 'echo 镜像名称：${image} && cp User/Rpc/Dockerfile ./  && ls -l && docker build  -t ${image} .'
            }    
        }
        stage('上传到镜像仓库') {
            container('veyron') {
                echo "上传到镜像仓库"
                sh 'docker login --username=${dockerusername} --password=${dockerpwd} http://${dockerrepo}'
                sh 'docker tag  ${image} ${dockerrepo}/go-zero-bolog/${image}'
                sh 'docker push ${dockerrepo}/go-zero-bolog/${image}'
            }    
        }
        stage('部署到k8s集群') {
            container('goctl'){
                echo "goctl生成yaml部署文件"
                env.deployYaml = sh(returnStdout: true, script: 'echo ${JOB_NAME}-${service}-deploy.yaml').trim()
                env.port='15010'
                sh 'echo ${port}'
                sh 'rm -f ${deployYaml}'
                sh 'goctl kube deploy -replicas 1 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name ${JOB_NAME}-${service} -namespace go-zero-bolog -image ${dockerrepo}/go-zero-bolog/${image} -o ${deployYaml} -port ${port} -serviceAccount find-endpoints '
                sh "sed -i 's/v2beta2/v2/g' ${deployYaml}"
            }
            container('veyron') {
                withKubeConfig(caCertificate: '', clusterName: 'kubernetes', contextName: '', credentialsId: 'k8scert', namespace: 'go-zero-bolog', restrictKubeConfigAccess: false, serverUrl: 'https://192.168.211.131:6443') {
                    echo "部署到k8s"
                    sh 'kubectl delete -f ${deployYaml}'
                    sh 'cat ${deployYaml}'
                }
            }
       }
    }
}
```
我们通过设置“veyron”这个container的runAsUser为“0”（0表示使用root用户运行），在这个container运行期间即可访问宿主机的docker

