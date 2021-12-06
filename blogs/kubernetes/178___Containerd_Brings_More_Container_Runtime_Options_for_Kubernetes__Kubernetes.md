|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/11/Containerd-Container-Runtime-Options-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-11-02 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.2636111 &#43;0300 MSK m=&#43;2.852649901  |

#   Containerd Brings More Container Runtime Options for Kubernetes  | Kubernetes

	
	
	
	
	***Editor&#39;s note: Today&#39;s post is by Lantao Liu, Software Engineer at Google, and Mike Brown, Open Source Developer Advocate at IBM.***
A *container runtime* is software that executes containers and manages container images on a node. Today, the most widely known container runtime is [Docker](https://www.docker.com/), but there are other container runtimes in the ecosystem, such as [rkt](https://coreos.com/rkt/), [containerd](https://containerd.io/), and [lxd](https://linuxcontainers.org/lxd/). Docker is by far the most common container runtime used in production Kubernetes environments, but Docker’s smaller offspring, containerd, may prove to be a better option. This post describes using containerd with Kubernetes.
Kubernetes 1.5 introduced an internal plugin API named [Container Runtime Interface (CRI)](https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes) to provide easy access to different container runtimes. CRI enables Kubernetes to use a variety of container runtimes without the need to recompile. In theory, Kubernetes could use any container runtime that implements CRI to manage pods, containers and container images.
Over the past 6 months, engineers from Google, Docker, IBM, ZTE, and ZJU have worked to implement CRI for containerd. The project is called [cri-containerd](https://github.com/kubernetes-incubator/cri-containerd), which had its [feature complete v1.0.0-alpha.0 release](https://github.com/kubernetes-incubator/cri-containerd/releases/tag/v1.0.0-alpha.0) on September 25, 2017. With cri-containerd, users can run Kubernetes clusters using containerd as the underlying runtime without Docker installed.
[Containerd](https://containerd.io/) is an [OCI](https://www.opencontainers.org/) compliant core container runtime designed to be embedded into larger systems. It provides the minimum set of functionality to execute containers and manages images on a node. It was initiated by Docker Inc. and [donated to CNCF](https://www.cncf.io/announcement/2017/03/29/containerd-joins-cloud-native-computing-foundation/) in March of 2017. The Docker engine itself is built on top of earlier versions of containerd, and will soon be updated to the newest version. Containerd is close to a feature complete stable release, with [1.0.0-beta.1](https://github.com/containerd/containerd/releases/tag/v1.0.0-beta.1) available right now.
Containerd has a much smaller scope than Docker, provides a golang client API, and is more focused on being embeddable.The smaller scope results in a smaller codebase that’s easier to maintain and support over time, matching Kubernetes requirements as shown in the following table:
[Cri-containerd](https://github.com/kubernetes-incubator/cri-containerd) is exactly that: an implementation of CRI for containerd. It operates on the same node as the Kubelet and containerd. Layered between Kubernetes and containerd, cri-containerd handles all CRI service requests from the Kubelet and uses containerd to manage containers and container images. Cri-containerd manages these service requests in part by forming containerd service requests while adding sufficient additional function to support the CRI requirements.

Compared with the current Docker CRI implementation ([dockershim](https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/dockershim)), cri-containerd eliminates an extra hop in the stack, making the stack more stable and efficient.
Cri-containerd uses containerd to manage the full container lifecycle and all container images. As also shown below, cri-containerd manages pod networking via [CNI](https://github.com/containernetworking/cni) (another CNCF project).

Let’s use an example to demonstrate how cri-containerd works for the case when Kubelet creates a single-container pod:
Cri-containerd v1.0.0-alpha.0 was released on Sep. 25, 2017.
It is feature complete. All Kubernetes features are supported.
All [CRI validation tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/cri-validation.md) have passed. (A CRI validation is a test framework for validating whether a CRI implementation meets all the requirements expected by Kubernetes.)
All regular [node e2e tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md) have passed. (The Kubernetes test framework for testing Kubernetes node level functionalities such as managing pods, mounting volumes etc.)
To learn more about the v1.0.0-alpha.0 release, see the [project repository](https://github.com/kubernetes-incubator/cri-containerd/releases/tag/v1.0.0-alpha.0).
For a multi-node cluster installer and bring up steps using ansible and kubeadm, see [this repo link](https://github.com/kubernetes-incubator/cri-containerd/blob/master/contrib/ansible/README.md).
For creating a cluster from scratch on Google Cloud, see [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way).
For a custom installation from release tarball, see [this repo link](https://github.com/kubernetes-incubator/cri-containerd/blob/master/docs/installation.md).
For a installation with LinuxKit on a local VM, see [this repo link](https://github.com/linuxkit/linuxkit/tree/master/projects/kubernetes).
We are focused on stability and usability improvements as our next steps.
We plan to release our v1.0.0-beta.0 by the end of 2017.
Cri-containerd is a Kubernetes incubator project located at [https://github.com/kubernetes-incubator/cri-containerd](https://github.com/kubernetes-incubator/cri-containerd). Any contributions in terms of ideas, issues, and/or fixes are welcome. The [getting started guide for developers](https://github.com/kubernetes-incubator/cri-containerd#getting-started-for-developers) is a good place to start for contributors.
Cri-containerd is developed and maintained by the Kubernetes SIG-Node community. We’d love to hear feedback from you. To join the community:


	

	


