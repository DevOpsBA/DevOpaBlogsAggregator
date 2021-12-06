|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/09/Kubernetes-18-Security-Workloads-And/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-09-29 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.3906526 &#43;0300 MSK m=&#43;2.979692101  |

#  Kubernetes 1.8: Security, Workloads and Feature Depth  | Kubernetes

	
	
	
	
	*Editor&#39;s note: today&#39;s post is by Aparna Sinha, Group Product Manager, Kubernetes, Google; Ihor Dvoretskyi, Developer Advocate, CNCF; Jaice Singer DuMars, Kubernetes Ambassador, Microsoft; and Caleb Miles, Technical Program Manager, CoreOS on the latest release of Kubernetes 1.8.*
We’re pleased to announce the delivery of Kubernetes 1.8, our third release this year. Kubernetes 1.8 represents a snapshot of many exciting enhancements and refinements underway. In addition to functional improvements, we’re increasing project-wide focus on maturing [process](https://github.com/kubernetes/sig-release), formalizing [architecture](https://github.com/kubernetes/community/tree/master/sig-architecture), and strengthening Kubernetes’ [governance model](https://github.com/kubernetes/community/tree/master/community/elections/2017). The evolution of mature processes clearly signals that sustainability is a driving concern, and helps to ensure that Kubernetes is a viable and thriving project far into the future.
Kubernetes 1.8 graduates support for [role based access control](https://en.wikipedia.org/wiki/Role-based_access_control) (RBAC) to stable. RBAC allows cluster administrators to [dynamically define roles](/docs/reference/access-authn-authz/rbac/) to enforce access policies through the Kubernetes API. Beta support for filtering outbound traffic through [network policies](/docs/concepts/services-networking/network-policies/) augments existing support for filtering inbound traffic to a pod. RBAC and Network Policies are two powerful tools for enforcing organizational and regulatory security requirements within Kubernetes.
Transport Layer Security (TLS) [certificate rotation](/docs/admin/kubelet-tls-bootstrapping/) for the Kubelet graduates to beta. Automatic certificate rotation eases secure cluster operation.
Kubernetes 1.8 promotes the core Workload APIs to beta with the apps/v1beta2 group and version. The beta contains the current version of Deployment, DaemonSet, ReplicaSet, and StatefulSet. The Workloads APIs provide a stable foundation for migrating existing workloads to Kubernetes as well as developing cloud native applications that target Kubernetes natively.
For those considering running Big Data workloads on Kubernetes, the Workloads API now enables [native Kubernetes support](https://apache-spark-on-k8s.github.io/userdocs/) in Apache Spark.
Batch workloads, such as nightly ETL jobs, will benefit from the graduation of [CronJobs](/docs/concepts/workloads/controllers/cron-jobs/) to beta.
[Custom Resource Definitions](/docs/concepts/api-extension/custom-resources/) (CRDs) remain in beta for Kubernetes 1.8. A CRD provides a powerful mechanism to extend Kubernetes with user-defined API objects. One use case for CRDs is the automation of complex stateful applications such as [key-value stores](https://github.com/coreos/etcd-operator), databases and [storage engines](https://rook.io/) through the Operator Pattern. Expect continued enhancements to CRDs such as [validation](/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/#validation) as stabilization continues.
Volume snapshots, PV resizing, automatic taints, priority pods, kubectl plugins, oh my!
In addition to stabilizing existing functionality, Kubernetes 1.8 offers a number of alpha features that preview new functionality.
Each Special Interest Group (SIG) in the community continues to deliver the most requested user features for their area. For a complete list, please visit the [release notes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG.md#v180).
Kubernetes 1.8 is available for [download on GitHub](https://github.com/kubernetes/kubernetes/releases/tag/v1.8.0). To get started with Kubernetes, check out these [interactive tutorials](/docs/tutorials/kubernetes-basics/).
The [Release team](https://github.com/kubernetes/features/blob/master/release-1.8/release_team.md) for 1.8 was led by Jaice Singer DuMars, Kubernetes Ambassador at Microsoft, and was comprised of 14 individuals responsible for managing all aspects of the release, from documentation to testing, validation, and feature completeness.
As the Kubernetes community has grown, our release process has become an amazing demonstration of collaboration in open source software development. Kubernetes continues to gain new users at a rapid clip. This growth creates a positive feedback cycle where more contributors commit code creating a more vibrant ecosystem.
According to [Redmonk](http://redmonk.com/fryan/2017/09/10/cloud-native-technologies-in-the-fortune-100/), 54 percent of Fortune 100 companies are running Kubernetes in some form with adoption coming from every sector across the world. Recent user stories from the community include:
Is Kubernetes helping your team? [Share your story](https://docs.google.com/a/google.com/forms/d/e/1FAIpQLScuI7Ye3VQHQTwBASrgkjQDSS5TP0g3AXfFhwSM9YpHgxRKFA/viewform) with the community.
Announced on September 11, [Kubernetes Certified Service Providers](https://www.cncf.io/certification/kcsp/) (KCSPs) are pre-qualified [organizations](https://kubernetes.io/partners/#kcsp) with deep experience helping enterprises successfully adopt Kubernetes. Individual professionals can now [register](https://www.cncf.io/certification/expert/) for the new Certified Kubernetes Administrator (CKA) program and exam, which requires passing an online, proctored, performance-based exam that tests one’s ability to solve multiple issues in a hands-on, command-line environment.
CNCF also offers [online training](https://www.cncf.io/certification/training/) that teaches the skills needed to create and configure a real-world Kubernetes cluster.
Join the community at [KubeCon &#43; CloudNativeCon](http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-north-america) in Austin, December 6-8 for the largest Kubernetes gathering ever. The premiere Kubernetes event will feature technical sessions, case studies, developer deep dives, salons and more! A full schedule of events and speakers will be available [here](http://events.linuxfoundation.org/events/kubecon-and-cloudnativecon-north-america/program/schedule) on September 28. Discounted [registration](https://www.regonline.com/registration/Checkin.aspx?EventID=1903774&amp;_ga=2.224109086.464556664.1498490094-1623727562.1496428006) ends October 6.
Ihor Dvoretskyi, Kubernetes 1.8 features release lead, will [present](https://osseu17.sched.com/event/C4AA) new features and enhancements at Open Source Summit EU in Prague, October 23. Registration is [still open](http://events.linuxfoundation.org/events/open-source-summit-europe/attend/register).
The simplest way to get involved with Kubernetes is by joining one of the many [Special Interest Groups](https://github.com/kubernetes/community/blob/master/sig-list.md) (SIGs) that align with your interests. Have something you’d like to broadcast to the Kubernetes community? Share your voice at our weekly [community meeting](https://github.com/kubernetes/community/blob/master/communication.md#weekly-meeting), and through the channels below.


	

	


