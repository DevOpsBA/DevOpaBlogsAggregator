|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/01/Kubernetes-V19-Beta-Windows-Support/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-01-09 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.1284793 &#43;0300 MSK m=&#43;2.717517301  |

# Kubernetes 1.9 对 Windows Server 容器提供 Beta 版本支持 | Kubernetes

	
	
	
	
	随着 Kubernetes v1.9 的发布，我们确保所有人在任何地方都能正常运行 Kubernetes 的使命前进了一大步。我们的 Beta 版本对 Windows Server 的支持进行了升级，并且在 Kubernetes 和 Windows 平台上都提供了持续的功能改进。为了在 Kubernetes 上运行许多特定于 Windows 的应用程序和工作负载，SIG-Windows 自2016年3月以来一直在努力，大大扩展了 Kubernetes 的实现场景和企业适用范围。
各种规模的企业都在 .NET 和基于 Windows 的应用程序上进行了大量投资。如今许多企业产品组合都包含 .NET 和 Windows，Gartner 声称 [80%](http://www.gartner.com/document/3446217) 的企业应用都在 Windows 上运行。根据 StackOverflow Insights，40% 的专业开发人员使用 .NET 编程语言（包括 .NET Core）。
但为什么这些信息都很重要？这意味着企业既有传统的，也有新生的云（microservice）应用程序，利用了大量的编程框架。业界正在大力推动将现有/遗留应用程序现代化到容器中，使用类似于“提升和转移”的方法。同时，也能灵活地向其他 Windows 或 Linux 容器引入新功能。容器正在成为打包、部署和管理现有程序和微服务应用程序的业界标准。IT 组织正在寻找一种更简单且一致的方法来跨 Linux 和 Windows 环境进行协调和管理容器。Kubernetes v1.9 现在对 Windows Server 容器提供了 Beta 版本支持，使之成为策划任何类型容器的明确选择。
Kubernetes 中对 Windows Server 容器的 Alpha 支持是非常有用的，尤其是对于概念项目和可视化 Kubernetes 中 Windows 支持的路线图。然而，Alpha 版本有明显的缺点，并且缺少许多特性，特别是在网络方面。SIG Windows、Microsoft、Cloudbase Solutions、Apprenda 和其他社区成员联合创建了一个全面的 Beta 版本，使 Kubernetes 用户能够开始评估和使用 Windows。
Kubernetes 对 Windows 服务器容器的一些关键功能改进包括：
如果您需要继续在 Linux 中运行 Kubernetes Control Plane 和 Master Components，现在也可以将 Windows Server 作为 Kubernetes 中的一个节点引入。对一个社区来说，这是一个巨大的里程碑和成就。现在，我们将会在 Kubernetes 中看到 .NET，.NET Core，ASP.NET，IIS，Windows 服务，Windows 可执行文件以及更多基于 Windows 的应用程序。
这个 Beta 版本进行了大量工作，但是社区意识到在将 Windows 支持作为生产工作负载发布为 GA（General Availability）之前，我们需要更多领域的投资。2018年前两个季度的重点关注领域包括：
随着我们在 Kubernetes 的普遍可用性方向不断取得进展，我们欢迎您参与进来，贡献代码、提供反馈，将 Windows 服务器容器部署到 Kubernetes 集群，或者干脆加入我们的社区。
谢谢大家，
Michael Michael (@michmike77)
SIG-Windows 领导人
Apprenda 产品管理高级总监


	

	


