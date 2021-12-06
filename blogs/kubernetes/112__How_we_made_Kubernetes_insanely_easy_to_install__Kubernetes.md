|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/09/How-We-Made-Kubernetes-Easy-To-Install/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-09-28 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.216995 &#43;0300 MSK m=&#43;3.806039201  |

#  How we made Kubernetes insanely easy to install  | Kubernetes

	
	
	
	
	*Editor&#39;s note: Today’s post is by [Luke Marsden](https://twitter.com/lmarsden), Head of Developer Experience, at Weaveworks, showing the Special Interest Group Cluster-Lifecycle’s recent work on kubeadm, a tool to make installing Kubernetes much simpler.*
Over at [SIG-cluster-lifecycle](https://github.com/kubernetes/community/blob/master/sig-cluster-lifecycle/README.md), we&#39;ve been hard at work the last few months on kubeadm, a tool that makes Kubernetes dramatically easier to install. We&#39;ve heard from users that installing Kubernetes is harder than it should be, and we want folks to be focused on writing great distributed apps not wrangling with infrastructure!
There are three stages in setting up a Kubernetes cluster, and we decided to focus on the second two (to begin with):
They use lots of different cloud providers, private clouds, bare metal, or even Raspberry Pi&#39;s, and almost always have their own preferred tools for automating provisioning machines: Terraform or CloudFormation, Chef, Puppet or Ansible, or even PXE booting bare metal. So we made an important decision: **kubeadm would not provision machines**. Instead, the only assumption it makes is that the user has some [computers running Linux](/docs/getting-started-guides/kubeadm/#prerequisites).
Another important constraint was we didn&#39;t want to just build another tool that &#34;configures Kubernetes from the outside, by poking all the bits into place&#34;. There are many external projects out there for doing this, but we wanted to aim higher. We chose to actually improve the Kubernetes core itself to make it easier to install. Luckily, a lot of the groundwork for making this happen had already been started.
We realized that if we made Kubernetes insanely easy to install manually, it should be obvious to users how to automate that process using any tooling.
So, enter [kubeadm](/docs/getting-started-guides/kubeadm/). It has no infrastructure dependencies, and satisfies the requirements above. It&#39;s easy to use and should be easy to automate. It&#39;s still in **alpha** , but it works like this:
For a video walkthrough, check this out:
Follow the [kubeadm getting started guide](/docs/getting-started-guides/kubeadm/) to try it yourself, and please give us [feedback on GitHub](https://github.com/kubernetes/kubernetes/issues/new), mentioning **@kubernetes/sig-cluster-lifecycle**!
Finally, I want to give a huge shout-out to so many people in the SIG-cluster-lifecycle, without whom this wouldn&#39;t have been possible. I&#39;ll mention just a few here:
This truly has been an excellent cross-company and cross-timezone achievement, with a lovely bunch of people. There&#39;s lots more work to do in SIG-cluster-lifecycle, so if you’re interested in these challenges join our SIG. Looking forward to collaborating with you all!
*--[Luke Marsden](https://twitter.com/lmarsden), Head of Developer Experience at [Weaveworks](https://twitter.com/weaveworks)*


	

	


