|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/06/Kubernetes-1-7-Security-Hardening-Stateful-Application-Extensibility-Updates/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-06-30 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.5131366 &#43;0300 MSK m=&#43;3.102176801  |

#  Kubernetes 1.7: Security Hardening, Stateful Application Updates and Extensibility  | Kubernetes

	
	
	
	
	Today we’re announcing Kubernetes 1.7, a milestone release that adds security, storage and extensibility features motivated by widespread production use of Kubernetes in the most demanding enterprise environments. 
At-a-glance, security enhancements in this release include encrypted secrets, network policy for pod-to-pod communication, node authorizer to limit kubelet access and client / server TLS certificate rotation. 
For those of you running scale-out databases on Kubernetes, this release has a major feature that adds automated updates to StatefulSets and enhances updates for DaemonSets. We are also announcing alpha support for local storage and a burst mode for scaling StatefulSets faster. 
Also, for power users, API aggregation in this release allows user-provided apiservers to be served along with the rest of the Kubernetes API at runtime. Additional highlights include support for extensible admission controllers, pluggable cloud providers, and container runtime interface (CRI) enhancements.
**What’s New**
Security:
Stateful workloads:
Extensibility:
Additional Features:
Deprecation: 
The above are a subset of the feature highlights in Kubernetes 1.7. For a complete list please visit the [release notes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG.md#v170).
**Adoption**
This release is possible thanks to our vast and open community. Together, we’ve already pushed more than 50,000 commits in just three years, and that’s only in the main Kubernetes repo. Additional extensions to Kubernetes are contributed in associated repos bringing overall stability to the project. This velocity makes Kubernetes one of the fastest growing open source projects -- ever. 
Kubernetes adoption has been coming from every sector across the world. Recent user stories from the community include: 
Huge kudos and thanks go out to the Kubernetes 1.7 [release team](https://github.com/kubernetes/features/blob/master/release-1.7/release_team.md), led by Dawn Chen of Google. 
**Availability**
Kubernetes 1.7 is available for [download on GitHub](https://github.com/kubernetes/kubernetes/releases/tag/v1.7.0). To get started with Kubernetes, try one of the these [interactive tutorials](/docs/tutorials/kubernetes-basics/). 
**Get Involved**
Join the community at [CloudNativeCon &#43; KubeCon](http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-north-america) in Austin Dec. 6-8 for the largest Kubernetes gathering ever. [Speaking submissions](http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-north-america/program/cfp) are open till August 21 and [discounted registration](https://www.regonline.com/registration/Checkin.aspx?EventID=1903774&amp;_ga=2.224109086.464556664.1498490094-1623727562.1496428006) ends October 6.
The simplest way to get involved is joining one of the many [Special Interest Groups](https://github.com/kubernetes/community/blob/master/sig-list.md) (SIGs) that align with your interests. Have something you’d like to broadcast to the Kubernetes community? Share your voice at our weekly [community meeting](https://github.com/kubernetes/community/blob/master/communication.md#weekly-meeting), and these channels:
Many thanks to our vast community of contributors and supporters in making this and all releases possible.
*-- Aparna Sinha, Group Product Manager, Kubernetes Google and Ihor Dvoretskyi, Program Manager, Kubernetes Mirantis*


	

	


