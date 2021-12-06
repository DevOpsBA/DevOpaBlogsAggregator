|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/09/27/kubernetes-1.12-kubelet-tls-bootstrap-and-azure-virtual-machine-scale-sets-vmss-move-to-general-availability/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-09-27 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.5849869 &#43;0300 MSK m=&#43;2.174021801  |

# Kubernetes 1.12: Kubelet TLS Bootstrap and Azure Virtual Machine Scale Sets (VMSS) Move to General Availability | Kubernetes

	
	
	
	
	**Author**: The 1.12 [Release Team](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.12/release_team.md)
We’re pleased to announce the delivery of Kubernetes 1.12, our third release of 2018!
Today’s release continues to focus on internal improvements and graduating features to stable in Kubernetes. This newest version graduates key features such as security and Azure. Notable additions in this release include two highly-anticipated features graduating to general availability: Kubelet TLS Bootstrap and Support for Azure Virtual Machine Scale Sets (VMSS).
These new features mean increased security, availability, resiliency, and ease of use to get production applications to market faster. The release also signifies the increasing maturation and sophistication of Kubernetes on the developer side.
Let’s dive into the key features of this release:
We’re excited to announce General Availability (GA) of [Kubelet TLS Bootstrap](https://github.com/kubernetes/features/issues/43). In Kubernetes 1.4, we introduced an API for requesting certificates from a cluster-level Certificate Authority (CA). The original intent of this API is to enable provisioning of TLS client certificates for kubelets. This feature allows for a kubelet to bootstrap itself into a TLS-secured cluster. Most importantly, it automates the provision and distribution of signed certificates.
Before, when a kubelet ran for the first time, it had to be given client credentials in an out-of-band process during cluster startup. The burden was on the operator to provision these credentials. Because this task was so onerous to manually execute and complex to automate, many operators deployed clusters with a single credential and single identity for all kubelets. These setups prevented deployment of node lockdown features like the Node authorizer and the NodeRestriction admission controller.
To alleviate this, [SIG Auth](https://github.com/kubernetes/community/tree/master/sig-auth) introduced a way for kubelet to generate a private key and a CSR for submission to a cluster-level certificate signing process. The v1 (GA) designation indicates production hardening and readiness, and comes with the guarantee of long-term backwards compatibility.
Alongside this, [Kubelet server certificate bootstrap and rotation](https://github.com/kubernetes/features/issues/267) is moving to beta. Currently, when a kubelet first starts, it generates a self-signed certificate/key pair that is used for accepting incoming TLS connections. This feature introduces a process for generating a key locally and then issuing a Certificate Signing Request to the cluster API server to get an associated certificate signed by the cluster’s root certificate authority. Also, as certificates approach expiration, the same mechanism will be used to request an updated certificate.
Azure Virtual Machine Scale Sets (VMSS) allow you to create and manage a homogenous VM pool that can automatically increase or decrease based on demand or a set schedule. This enables you to easily manage, scale, and load balance multiple VMs to provide high availability and application resiliency, ideal for large-scale applications that can run as Kubernetes workloads.
With this new stable feature, Kubernetes supports the [scaling of containerized applications with Azure VMSS](https://github.com/kubernetes/features/issues/514), including the ability to [integrate it with cluster-autoscaler](https://github.com/kubernetes/features/issues/513) to automatically adjust the size of the Kubernetes clusters based on the same conditions.
[code](https://github.com/kubernetes/features/issues/585) is a new cluster-scoped resource that surfaces container runtime properties to the control plane being released as an alpha feature.
[Snapshot / restore functionality for Kubernetes and CSI](https://github.com/kubernetes/features/issues/177) is being introduced as an alpha feature. This provides standardized APIs design (CRDs) and adds PV snapshot/restore support for CSI volume drivers.
[Topology aware dynamic provisioning](https://github.com/kubernetes/features/issues/561) is now in beta, meaning storage resources can now understand where they live. This also includes beta support to [AWS EBS](https://github.com/kubernetes/features/issues/567) and [GCE PD](https://github.com/kubernetes/features/issues/558).
[Configurable pod process namespace sharing](https://github.com/kubernetes/features/issues/495) is moving to beta, meaning users can configure containers within a pod to share a common PID namespace by setting an option in the PodSpec.
[Taint node by condition](https://github.com/kubernetes/features/issues/382) is now in beta, meaning users have the ability to represent node conditions that block scheduling by using taints.
[Arbitrary / Custom Metrics](https://github.com/kubernetes/features/issues/117) in the Horizontal Pod Autoscaler is moving to a second beta to test some additional feature enhancements. This reworked Horizontal Pod Autoscaler functionality includes support for custom metrics and status conditions.
Improvements that will allow the [Horizontal Pod Autoscaler to reach proper size faster](https://github.com/kubernetes/features/issues/591) are moving to beta.
[Vertical Scaling of Pods](https://github.com/kubernetes/features/issues/21) is now in beta, which makes it possible to vary the resource limits on a pod over its lifetime. In particular, this is valuable for pets (i.e., pods that are very costly to destroy and re-create).
[Encryption at rest via KMS](https://github.com/kubernetes/features/issues/460) is now in beta. This adds multiple encryption providers, including Google Cloud KMS, Azure Key Vault, AWS KMS, and Hashicorp Vault, that will encrypt data as it is stored to etcd.
Kubernetes 1.12 is available for [download on GitHub](https://github.com/kubernetes/kubernetes/releases/tag/v1.12.0). To get started with Kubernetes, check out these [interactive tutorials](/docs/tutorials/). You can also install 1.12 using [Kubeadm](/docs/setup/independent/create-cluster-kubeadm/).
If you’re interested in exploring these features more in depth, check back next week for our 5 Days of Kubernetes series where we’ll highlight detailed walkthroughs of the following features:
This release is made possible through the effort of hundreds of individuals who contributed both technical and non-technical content. Special thanks to the [release team](https://github.com/kubernetes/sig-release/blob/master/releases/release-1.12/release_team.md) led by Tim Pepper, Orchestration &amp; Containers Lead, at VMware Open Source Technology Center. The 36 individuals on the release team coordinate many aspects of the release, from documentation to testing, validation, and feature completeness.
As the Kubernetes community has grown, our release process represents an amazing demonstration of collaboration in open source software development. Kubernetes continues to gain new users at a rapid clip. This growth creates a positive feedback cycle where more contributors commit code creating a more vibrant ecosystem. Kubernetes has over 22,000 individual contributors to date and an active community of more than 45,000 people.
The CNCF has continued refining DevStats, an ambitious project to visualize the myriad contributions that go into the project. [K8s DevStats](https://devstats.k8s.io) illustrates the breakdown of contributions from major company contributors, as well as an impressive set of preconfigured reports on everything from individual contributors to pull request lifecycle times. On average, 259 different companies and over 1,400 individuals contribute to Kubernetes each month. [Check out DevStats](https://k8s.devstats.cncf.io/d/11/companies-contributing-in-repository-groups?orgId=1&amp;var-period=m&amp;var-repogroup_name=All) to learn more about the overall velocity of the Kubernetes project and community.
Established, global organizations are using [Kubernetes in production](https://kubernetes.io/case-studies/) at massive scale. Recently published user stories from the community include:
Is Kubernetes helping your team? [Share your story](https://docs.google.com/a/google.com/forms/d/e/1FAIpQLScuI7Ye3VQHQTwBASrgkjQDSS5TP0g3AXfFhwSM9YpHgxRKFA/viewform) with the community.
The world’s largest Kubernetes gathering, KubeCon &#43; CloudNativeCon is coming to [Shanghai](https://events.linuxfoundation.cn/events/kubecon-cloudnativecon-china-2018/) from November 13-15, 2018 and [Seattle](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-north-america-2018/) from December 10-13, 2018. This conference will feature technical sessions, case studies, developer deep dives, salons and more! [Register today](https://www.cncf.io/community/kubecon-cloudnativecon-events/)!
Join members of the Kubernetes 1.12 release team on November 6th at 10am PDT to learn about the major features in this release. Register [here](https://zoom.us/webinar/register/WN_DYMejau3TMaTbk91oC3YkA).
The simplest way to get involved with Kubernetes is by joining one of the many [Special Interest Groups](https://github.com/kubernetes/community/blob/master/sig-list.md) (SIGs) that align with your interests. Have something you’d like to broadcast to the Kubernetes community? Share your voice at our weekly [community meeting](https://github.com/kubernetes/community/blob/master/communication.md#weekly-meeting), and through the channels below.
Thank you for your continued feedback and support.


	

	


