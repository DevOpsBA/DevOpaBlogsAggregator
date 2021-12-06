|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/10/10/kubernetes-v1.12-introducing-runtimeclass/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-10-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.5408726 &#43;0300 MSK m=&#43;2.129907301  |

# Kubernetes v1.12: Introducing RuntimeClass | Kubernetes

	
	
	
	
	**Author**: Tim Allclair (Google)
Kubernetes originally launched with support for Docker containers running native applications on a Linux host. Starting with [rkt](https://kubernetes.io/blog/2016/07/rktnetes-brings-rkt-container-engine-to-kubernetes/) in Kubernetes 1.3 more runtimes were coming, which lead to the development of the [Container Runtime Interface](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/) (CRI). Since then, the set of alternative runtimes has only expanded: projects like [Kata Containers](https://katacontainers.io/) and [gVisor](https://github.com/google/gvisor) were announced for stronger workload isolation, and Kubernetes&#39; Windows support has been [steadily progressing](https://kubernetes.io/blog/2018/01/kubernetes-v19-beta-windows-support/).
With runtimes targeting so many different use cases, a clear need for mixed runtimes in a cluster arose. But all these different ways of running containers have brought a new set of problems to deal with:
**RuntimeClass** aims to solve these issues.
RuntimeClass was recently introduced as an alpha feature in Kubernetes 1.12. The initial implementation focuses on providing a runtime selection API, and paves the way to address the other open problems.
The RuntimeClass resource represents a container runtime supported in a Kubernetes cluster. The cluster provisioner sets up, configures, and defines the concrete runtimes backing the RuntimeClass. In its current form, a RuntimeClassSpec holds a single field, the **RuntimeHandler**. The RuntimeHandler is interpreted by the CRI implementation running on a node, and mapped to the actual runtime configuration. Meanwhile the PodSpec has been expanded with a new field, **RuntimeClassName**, which names the RuntimeClass that should be used to run the pod.
Why is RuntimeClass a pod level concept? The Kubernetes resource model expects certain resources to be shareable between containers in the pod. If the pod is made up of different containers with potentially different resource models, supporting the necessary level of resource sharing becomes very challenging. For example, it is extremely difficult to support a loopback (localhost) interface across a VM boundary, but this is a common model for communication between two containers in a pod.
The RuntimeClass resource is an important foundation for surfacing runtime properties to the control plane. For example, to implement scheduler support for clusters with heterogeneous nodes supporting different runtimes, we might add [NodeAffinity](/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity) terms to the RuntimeClass definition. Another area to address is managing the variable resource requirements to run pods of different runtimes. The [Pod Overhead proposal](https://docs.google.com/document/d/1EJKT4gyl58-kzt2bnwkv08MIUZ6lkDpXcxkHqCvvAp4/preview) was an early take on this that aligns nicely with the RuntimeClass design, and may be pursued further.
Many other RuntimeClass extensions have also been proposed, and will be revisited as the feature continues to develop and mature. A few more extensions that are being considered include:
RuntimeClass will be under active development at least through 2019, and we’re excited to see the feature take shape, starting with the RuntimeClass alpha in Kubernetes 1.12.


	

	


