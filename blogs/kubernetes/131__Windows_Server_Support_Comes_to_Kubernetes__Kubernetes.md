|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/12/Windows-Server-Support-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-12-21 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.9158593 &#43;0300 MSK m=&#43;3.504901801  |

#  Windows Server Support Comes to Kubernetes  | Kubernetes

	
	
	
	
	*Editor’s note: this post is part of a [series of in-depth articles](https://kubernetes.io/blog/2016/12/five-days-of-kubernetes-1-5/) on what&#39;s new in Kubernetes 1.5*
Extending on the theme of giving users choice, [Kubernetes 1.5 release](https://kubernetes.io/blog/2016/12/kubernetes-1-5-supporting-production-workloads/) includes the support for Windows Servers. WIth more than [80%](http://www.gartner.com/document/3446217) of enterprise apps running Java on Linux or .Net on Windows, Kubernetes is previewing capabilities that extends its reach to the mass majority of enterprise workloads. 
The new Kubernetes Windows Server 2016 and Windows Container support includes public preview with the following features:
The process to bring Windows Server to Kubernetes has been a truly multi-vendor effort and championed by the [Windows Special Interest Group (SIG)](https://github.com/kubernetes/community/blob/master/sig-windows/README.md) - Apprenda, Google, Red Hat and Microsoft were all involved in bringing Kubernetes to Windows Server. On the community effort to bring Kubernetes to Windows Server, Taylor Brown, Principal Program Manager at Microsoft stated that “This new Kubernetes community work furthers Windows Server container support options for popular orchestrators, reinforcing Microsoft’s commitment to choice and flexibility for both Windows and Linux ecosystems.”
**Guidance for Current Usage**
|
Where to use Windows Server support?
|
Right now organizations should start testing Kubernetes on Windows Server and provide feedback. Most organizations take months to set up hardened production environments and general availability should be available in next few releases of Kubernetes.
|
|
What works?
|
Most of the Kubernetes constructs, such as Pods, Services, Labels, etc. work with Windows Containers.
|
|
What doesn’t work yet?
|
|
|
When will it be ready for all production workloads (general availability)?
|
The goal is to refine the networking and other areas that need work to get Kubernetes users a production version of Windows Server 2016 - including with Windows Nano Server and Windows Server Core installation options - support in the next couple releases.
|
**Technical Demo**
**Roadmap**
Support for Windows Server-based containers is in alpha release mode for Kubernetes 1.5, but the community is not stopping there. Customers want enterprise hardened container scheduling and management for their entire tech portfolio. That has to include full parity of features among Linux and Windows Server in production. The [Windows Server SIG](https://github.com/kubernetes/community/blob/master/sig-windows/README.md) will deliver that parity within the next one or two releases of Kubernetes through a few key areas of investment:
To get started with Kubernetes on Windows Server 2016, please visit the [GitHub guide](/docs/getting-started-guides/windows/) for more details.
If you want to help with Windows Server support, then please connect with the [Windows Server SIG](https://github.com/kubernetes/community/blob/master/sig-windows/README.md) or connect directly with Michael Michael, the SIG lead, on [GitHub](https://github.com/michmike). 
*--[Michael Michael](https://twitter.com/michmike77), Senior Director of Product Management, Apprenda *
|  |
| Kubernetes on Windows Server 2016 Architecture |


	

	


