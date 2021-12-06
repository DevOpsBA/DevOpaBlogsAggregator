|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/06/Container-Design-Patterns/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-06-21 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.6238371 &#43;0300 MSK m=&#43;4.212883701  |

#  Container Design Patterns  | Kubernetes

	
	
	
	
	Kubernetes automates deployment, operations, and scaling of applications, but our goals in the Kubernetes project extend beyond system management -- we want Kubernetes to help developers, too. Kubernetes should make it easy for them to write the distributed applications and services that run in cloud and datacenter environments. To enable this, Kubernetes defines not only an API for administrators to perform management actions, but also an API for containerized applications to interact with the management platform.
Our work on the latter is just beginning, but you can already see it manifested in a few features of Kubernetes. For example:
More generally, we see Kubernetes enabling a new generation of design patterns, similar to [object oriented design patterns](https://en.wikipedia.org/wiki/Object-oriented_programming#Design_patterns), but this time for containerized applications. That design patterns would emerge from containerized architectures is not surprising -- containers provide many of the same benefits as software objects, in terms of modularity/packaging, abstraction, and reuse. Even better, because containers generally interact with each other via HTTP and widely available data formats like JSON, the benefits can be provided in a language-independent way.
This week Kubernetes co-founder Brendan Burns is presenting a [strong](https://www.usenix.org/conference/hotcloud16/workshop-program/presentation/burns) outlining our thoughts on this topic at the [8th Usenix Workshop on Hot Topics in Cloud Computing](https://www.usenix.org/conference/hotcloud16) (HotCloud ‘16), a venue where academic researchers and industry practitioners come together to discuss ideas at the forefront of research in private and public cloud technology. The paper describes three classes of patterns: management patterns (such as those described above), patterns involving multiple cooperating containers running on the same node, and patterns involving containers running across multiple nodes. We don’t want to spoil the fun of reading the paper, but we will say that you’ll see that the [Pod](/docs/user-guide/pods/) abstraction is a key enabler for the last two types of patterns.
As the Kubernetes project continues to bring our decade of experience with [Borg](https://queue.acm.org/detail.cfm?id=2898444) to the open source community, we aim not only to make application deployment and operations at scale simple and reliable, but also to make it easy to create “cloud-native” applications in the first place. Our work on documenting our ideas around design patterns for container-based services, and Kubernetes’s enabling of such patterns, is a first step in this direction. We look forward to working with the academic and practitioner communities to identify and codify additional patterns, with the aim of helping containers fulfill the promise of bringing increased simplicity and reliability to the entire software lifecycle, from development, to deployment, to operations.
To learn more about the Kubernetes project visit [kubernetes.io](http://kubernetes.io/) or chat with us on Slack at [slack.kubernetes.io](http://slack.kubernetes.io/).
-*-Brendan Burns and David Oppenheimer, Software Engineers, Google*


	

	


