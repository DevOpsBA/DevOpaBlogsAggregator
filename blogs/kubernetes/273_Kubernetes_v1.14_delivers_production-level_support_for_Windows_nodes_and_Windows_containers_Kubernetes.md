|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2019/04/01/kubernetes-v1.14-delivers-production-level-support-for-windows-nodes-and-windows-containers/        |
| Tags              | [kubernetes]       |
| Date Create       | 2019-04-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.2662656 &#43;0300 MSK m=&#43;1.855298701  |

# Kubernetes v1.14 delivers production-level support for Windows nodes and Windows containers | Kubernetes

	
	
	
	
	**Authors:** Michael Michael (VMware), Patrick Lang (Microsoft)
The first release of Kubernetes in 2019 brings a highly anticipated feature -  production-level support for Windows workloads. Up until now Windows node support in Kubernetes has been in beta, allowing many users to experiment and see the value of Kubernetes for Windows containers. While in beta, developers in the Kubernetes community and Windows Server team worked together to improve the container runtime, build a continuous testing process, and complete features needed for a good user experience. Kubernetes now officially supports adding Windows nodes as worker nodes and scheduling Windows containers, enabling a vast ecosystem of Windows applications to leverage the power of our platform.
As Windows developers and devops engineers have been adopting containers over the last few years, they&#39;ve been looking for a way to manage all their workloads with a common interface. Kubernetes has taken the lead for container orchestration, and this gives users a consistent way to manage their container workloads whether they need to run on Linux or Windows.
The journey to a stable release of Windows in Kubernetes was not a walk in the park. The community has been working on Windows support for 3 years, delivering an alpha release with v1.5, a beta with v1.9, and now a stable release with v1.14. We would not be here today without rallying broad support and getting significant contributions from companies including Microsoft, Docker, VMware, Pivotal, Cloudbase Solutions, Google and Apprenda. During this journey, there were 3 critical points in time that significantly advanced our progress.
With v1.14, we&#39;re declaring that Windows node support is stable, well-tested, and ready for adoption in production scenarios. This is a huge milestone for many reasons. For Kubernetes, it strengthens its position in the industry, enabling a vast ecosystem of Windows-based applications to be deployed on the platform. For Windows operators and developers, this means they can use the same tools and processes to manage their Windows and Linux workloads, taking full advantage of the efficiencies of the cloud-native ecosystem powered by Kubernetes. Let’s dig in a little bit into these.
“Using Kubernetes on Windows allows us to run our internal web applications as microservices. This provides quick scaling in response to load, smoother upgrades, and allows for different development groups to build without worry of other group&#39;s version dependencies. We save money because development times are shorter and operation&#39;s time is not spent maintaining multiple virtual machine environments,” said Jeremy, a lead devops engineer working for a top multinational legal firm, one of the early adopters of Windows on Kubernetes.
There are many features that are surfaced with this release. We want to turn your attention to a few key features and enablers of Windows support in Kubernetes. For a detailed list of supported functionality, you can read our [documentation](https://kubernetes.io/docs/setup/windows/intro-windows-in-kubernetes/#supported-functionality).
When looking at Windows support in Kubernetes, many start drawing comparisons to Linux containers. Although some of the comparisons that highlight limitations are fair, it is important to distinguish between **operational limitations and differences between the Windows and Linux operating systems**. From a container management standpoint, we must  strike a balance between preserving OS-specific behaviors required for application compatibility, and reaching operational consistency in Kubernetes across multiple operating systems. For example, some Linux-specific file system features, user IDs and permissions exposed through Kubernetes will not work on Windows today, and users are familiar with these fundamental differences. We will also be adding support for Windows-specific configurations to meet the needs of Windows customers that may not exist on Linux. The alpha support for Windows Group Managed Service Accounts is one example. Other areas such as memory reservations for Windows pods and the Windows kubelet are a work in progress and highlight an operational limitation. We will continue working on operational limitations based on what’s important to our community in future releases.
Today, Kubernetes master components will continue to run on Linux. That way users can add Windows nodes without having to create a separate Kubernetes cluster. As always, our future direction is set by the community, so more components, features and deployment methods will come over time. Users should understand the differences between Windows and Linux and utilize the advantages of each platform. Our goal with this release is not to make Windows interchangeable with Linux or to answer the question of Windows vs Linux. We offer consistency in management. Managing workloads without automation is tedious and expensive. Rewriting or re-architecting workloads is even more expensive. Containers provide a clear path forward whether your app runs on Linux or Windows, and Kubernetes brings an IT organization operational consistency.
As a community, our work is not complete. As already mentioned , we still have a fair bit of [limitations](https://kubernetes.io/docs/setup/windows/intro-windows-in-kubernetes/#limitations) and a healthy [roadmap](https://kubernetes.io/docs/setup/windows/intro-windows-in-kubernetes/#what-s-next). We will continue making progress and enhancing Windows container support in Kubernetes, with some notable upcoming features including:
We welcome you to get involved and join our community to share feedback and deployment stories, and contribute to code, docs, and improvements of any kind.
Thank you and feel free to reach us individually if you have any questions.
Michael Michael

SIG-Windows Chair

Director of Product Management, VMware

@michmike77 on Twitter

@m2 on Slack
Patrick Lang

SIG-Windows Chair

Senior Software Engineer, Microsoft

@PatrickLang on Slack


	

	


