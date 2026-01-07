# go-zero-cicd

> 🚀 基于 **go-zero** 微服务框架的 CI/CD 实战项目  
> 使用 **GitLab + Jenkins + Harbor + Kubernetes** 构建从代码提交到自动部署的完整流水线。

本项目用于演示 **Go 微服务在真实生产环境中的 CI/CD 落地方案**，覆盖构建、测试、镜像管理及 Kubernetes 自动部署流程，适合作为 **DevOps / 后端 / 云原生方向的实践与展示项目**。

---

## ✨ 项目特性

- 🔧 基于 **go-zero** 的 Go 微服务示例
- 🔁 GitLab 提交自动触发 CI/CD
- 🧪 自动化测试与代码检查
- 📦 Docker 镜像构建并推送至 **Harbor**
- ☸️ 自动部署到 **Kubernetes**
- 📈 Jenkins Pipeline 流水线管理

---

## 🧱 技术栈

| 组件 | 说明 |
|----|----|
| Go | 后端开发语言 |
| go-zero | 微服务框架 |
| GitLab | 代码仓库 / CI 触发 |
| Jenkins | CI/CD Pipeline |
| Docker | 镜像构建 |
| Harbor | 私有镜像仓库 |
| Kubernetes | 容器编排与部署 |

---

## 📁 项目结构

```text
.
├── main.go                # go-zero 示例服务入口
├── go.mod
├── go.sum
├── Jenkinsfile            # Jenkins CI/CD Pipeline 定义
├── .gitlab-ci.yml         # GitLab CI 配置（可选）
├── scripts/               # 构建 / 部署脚本
├── helm/                  # Helm Chart 或 K8s 部署文件
└── README.md
```

目录：
[1、CI/CD环境搭建之安装gitlab](./DOC/1、部署CICD环境-搭建gitlab.md)<br>
[2、部署CICD环境-搭建harbor](./DOC/2、部署CICD环境-搭建harbor.md)<br>
[3、部署CICD环境-搭建Jenkins](./DOC/3、部署CICD环境-搭建Jenkins.md)<br>
[4、传统方式发布go-zero微服务到k8s](./DOC/4、传统方式发布go-zero微服务到k8s.md)<br>
[5、Jenkins安装kubernetes插件](./DOC/5、Jenkins安装kubernetes插件.md)<br>
[6、创建SCM实现提交代码触发构建任务](./DOC/6、创建SCM实现提交代码触发构建任务.md)<br>
[7、使用脚本式语法编写pipeline构建go-zero微服务发布任务](./DOC/7、使用脚本式语法编写pipeline构建go-zero微服务发布任务.md)<br>
[8、解决使用宿主机docker权限不够的问题](./DOC/8、解决使用宿主机docker权限不够的问题.md)<br>
