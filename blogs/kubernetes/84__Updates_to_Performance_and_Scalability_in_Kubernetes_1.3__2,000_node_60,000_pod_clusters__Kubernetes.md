|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/07/Update-On-Kubernetes-For-Windows-Server-Containers/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-07-07 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.6133579 &#43;0300 MSK m=&#43;4.202404401  |

#  Updates to Performance and Scalability in Kubernetes 1.3 -- 2,000 node 60,000 pod clusters  | Kubernetes

	
	
	
	
	We are proud to announce that with the [release of version 1.3](https://kubernetes.io/blog/2016/07/kubernetes-1-3-bridging-cloud-native-and-enterprise-workloads/), Kubernetes now supports 2000-node clusters with even better end-to-end pod startup time. The latency of our API calls are within our one-second [Service Level Objective (SLO)](https://en.wikipedia.org/wiki/Service_level_objective) and most of them are even an order of magnitude better than that. It is possible to run larger deployments than a 2,000 node cluster, but performance may be degraded and it may not meet our strict SLO.
In this blog post we discuss the detailed performance results from Kubernetes 1.3 and what changes we made from version 1.2 to achieve these results. We also describe Kubemark, a performance testing tool that we’ve integrated into our continuous testing framework to detect performance and scalability regressions.
**Evaluation Methodology**
We have described our test scenarios in a [previous blog post](https://kubernetes.io/blog/2016/03/1000-nodes-and-beyond-updates-to-Kubernetes-performance-and-scalability-in-12). The biggest change since the 1.2 release is that in our API responsiveness tests we now create and use multiple namespaces. In particular for the 2000-node/60000 pod cluster tests we create 8 namespaces. The change was done because we believe that users of such very large clusters are likely to use many namespaces, certainly at least 8 in the cluster in total.
**Metrics from Kubernetes 1.3**
So, what is the performance of Kubernetes version 1.3? The following graph shows the end-to-end pod startup latency with a 2000 and 1000 node cluster. For comparison we show the same metric from Kubernetes 1.2 with a 1000-node cluster.

The next graphs show API response latency for a v1.3 2000-node cluster.


**How did we achieve these improvements?**
The biggest change that we made for scalability in Kubernetes 1.3 was adding an efficient [Protocol Buffer](https://developers.google.com/protocol-buffers/)-based serialization format to the API as an alternative to JSON. It is primarily intended for communication between Kubernetes control plane components, but all API server clients can use this format. All Kubernetes control plane components now use it for their communication, but the system continues to support JSON for backward compatibility.
We didn’t change the format in which we store cluster state in etcd to Protocol Buffers yet, as we’re still working on the upgrade mechanism. But we’re very close to having this ready, and we expect to switch the storage format to Protocol Buffers in Kubernetes 1.4. Our experiments show that this should reduce pod startup end-to-end latency by another 30%.
**How do we test Kubernetes at scale?**
Spawning clusters with 2000 nodes is expensive and time-consuming. While we need to do this at least once for each release to collect real-world performance and scalability data, we also need a lighter-weight mechanism that can allow us to quickly evaluate our ideas for different performance improvements, and that we can run continuously to detect performance regressions. To address this need we created a tool call “Kubemark.”
**What is “Kubemark”?**
Kubemark is a performance testing tool which allows users to run experiments on emulated clusters. We use it for measuring performance in large clusters.
A Kubemark cluster consists of two parts: a real master node running the normal master components, and a set of “hollow” nodes. The prefix “hollow” means an implementation/instantiation of a component with some “moving parts” mocked out. The best example is hollow-kubelet, which pretends to be an ordinary Kubelet, but doesn’t start any containers or mount any volumes. It just claims it does, so from master components’ perspective it behaves like a real Kubelet.
Since we want a Kubemark cluster to be as similar to a real cluster as possible, we use the real Kubelet code with an injected fake Docker client. Similarly hollow-proxy (KubeProxy equivalent) reuses the real KubeProxy code with injected no-op Proxier interface (to avoid mutating iptables).
Thanks to those changes
**How do we set up Kubemark clusters?**
To create a Kubemark cluster we use the power the Kubernetes itself gives us - we run Kubemark clusters on Kubernetes. Let’s describe this in detail.
In order to create a N-node Kubemark cluster, we:
One thing worth mentioning here is that while running Kubemark, underneath we’re also testing Kubernetes correctness. Obviously your Kubemark cluster will not work correctly if the base Kubernetes cluster under it doesn’t work. 
**Performance measured in real clusters vs Kubemark**
Crucially, the performance of Kubemark clusters is mostly similar to the performance of real clusters. For the pod startup end-to-end latency, as shown in the graph below, the difference is negligible:

For the API-responsiveness, the differences are higher, though generally less than 2x. However, trends are exactly the same: an improvement/regression in a real cluster is visible as a similar percentage drop/increase in metrics in Kubemark.
**Conclusion**
We continue to improve the performance and scalability of Kubernetes. In this blog post we 
showed that the 1.3 release scales to 2000 nodes while meeting our responsiveness SLOs
explained the major change we made to improve scalability from the 1.2 release, and 
described Kubemark, our emulation framework that allows us to quickly evaluate the performance impact of code changes, both when experimenting with performance improvement ideas and to detect regressions as part of our continuous testing infrastructure.
Please join our community and help us build the future of Kubernetes! If you’re particularly interested in scalability, participate by:
For more information about the Kubernetes project, visit [kubernetes.io](http://kubernetes.io/) and follow us on Twitter [@Kubernetesio](https://twitter.com/kubernetesio). 
*-- Wojciech Tyczynski, Software Engineer, Google*


	

	


