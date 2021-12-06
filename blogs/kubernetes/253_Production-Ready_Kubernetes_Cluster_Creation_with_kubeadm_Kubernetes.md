|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/12/04/production-ready-kubernetes-cluster-creation-with-kubeadm/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-12-04 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.4522963 &#43;0300 MSK m=&#43;2.041330501  |

# Production-Ready Kubernetes Cluster Creation with kubeadm | Kubernetes

	
	
	
	
	**Authors**: Lucas Käldström (CNCF Ambassador) and Luc Perkins (CNCF Developer Advocate)
[kubeadm](/docs/setup/independent/create-cluster-kubeadm/) is a tool that enables Kubernetes administrators to quickly and easily bootstrap minimum viable clusters that are fully compliant with [Certified Kubernetes](https://github.com/cncf/k8s-conformance/blob/master/terms-conditions/Certified_Kubernetes_Terms.md) guidelines. It&#39;s been under active development by [SIG Cluster Lifecycle](https://github.com/kubernetes/community/tree/master/sig-cluster-lifecycle) since 2016 and we&#39;re excited to announce that it has now graduated from beta to stable and generally available (GA)!
This GA release of kubeadm is an important event in the progression of the Kubernetes ecosystem, bringing stability to an area where stability is paramount.
The goal of kubeadm is to provide a foundational implementation for Kubernetes cluster setup and administration. kubeadm ships with best-practice defaults but can also be customized to support other ecosystem requirements or vendor-specific approaches. kubeadm is designed to be easy to integrate into larger deployment systems and tools.
kubeadm is focused on bootstrapping Kubernetes clusters on existing infrastructure and performing an essential set of maintenance tasks. The core of the kubeadm interface is quite simple: new control plane nodes are created by running [code](/docs/reference/setup-tools/kubeadm/kubeadm-init/) and worker nodes are joined to the control plane by running [code](/docs/reference/setup-tools/kubeadm/kubeadm-join/). Also included are utilities for managing already bootstrapped clusters, such as control plane upgrades and token and certificate renewal.
To keep kubeadm lean, focused, and vendor/infrastructure agnostic, the following tasks are out of its scope:
Infrastructure provisioning, for example, is left to other SIG Cluster Lifecycle projects, such as the [Cluster API](https://github.com/kubernetes-sigs/cluster-api). Instead, kubeadm covers only the common denominator in every Kubernetes cluster: the [control plane](/docs/concepts/overview/components/#control-plane-components). The user may install their preferred networking solution and other add-ons on top of Kubernetes *after* cluster creation.
General Availability means different things for different projects. For kubeadm, going GA means not only that the process of creating a conformant Kubernetes cluster is now stable, but also that kubeadm is flexible enough to support a wide variety of deployment options.
We now consider kubeadm to have achieved GA-level maturity in each of these important domains:
SIG Cluster Lifecycle has identified a handful of likely kubeadm user profiles, although we expect that kubeadm at GA can satisfy many other scenarios as well.
Here&#39;s our list:
All these users can benefit from kubeadm graduating to a stable GA state.
Although kubeadm is GA, the SIG Cluster Lifecycle will continue to be committed to improving the user experience in managing Kubernetes clusters. We&#39;re launching a survey to collect community feedback about kubeadm for the sake of future improvement.
The survey is available at [https://bit.ly/2FPfRiZ](https://bit.ly/2FPfRiZ). Your participation would be highly valued!
This release wouldn&#39;t have been possible without the help of the great people that have been contributing to the SIG. SIG Cluster Lifecycle would like to thank a few key kubeadm contributors:
We also want to thank all the companies making it possible for their developers to work on Kubernetes, and all the other people that have contributed in various ways towards making kubeadm as stable as it is today!


	

	


