|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/03/1000-Nodes-And-Beyond-Updates-To-Kubernetes-Performance-And-Scalability-In-12/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-03-28 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.8803044 &#43;0300 MSK m=&#43;4.469352401  |

#  1000 nodes and beyond: updates to Kubernetes performance and scalability in 1.2  | Kubernetes

	
	
	
	
	*Editor&#39;s note: this is the first in a [series of in-depth posts](https://kubernetes.io/blog/2016/03/five-days-of-kubernetes-12) on what&#39;s new in Kubernetes 1.2*
We&#39;re proud to announce that with the [release of 1.2](https://kubernetes.io/blog/2016/03/kubernetes-1-2-even-more-performance-upgrades-plus-easier-application-deployment-and-management), Kubernetes now supports 1000-node clusters, with a reduction of 80% in 99th percentile tail latency for most API operations. This means in just six months, we&#39;ve increased our overall scale by 10 times while maintaining a great user experience — the 99th percentile pod startup times are less than 3 seconds, and 99th percentile latency of most API operations is tens of milliseconds (the exception being LIST operations, which take hundreds of milliseconds in very large clusters).
Words are fine, but nothing speaks louder than a demo. Check this out!
In the above video, you saw the cluster scale up to 10 M queries per second (QPS) over 1,000 nodes, including a rolling update, with zero downtime and no impact to tail latency. That’s big enough to be one of the top 100 sites on the Internet!
In this blog post, we’ll cover the work we did to achieve this result, and discuss some of our future plans for scaling even higher.
We benchmark Kubernetes scalability against the following Service Level Objectives (SLOs):
Kubernetes offers high-level abstractions for users to represent their applications. For example, the ReplicationController is an abstraction representing a collection of [pods](/docs/user-guide/pods/). Listing all ReplicationControllers or listing all pods from a given ReplicationController is a very common use case. On the other hand, there is little reason someone would want to list all pods in the system — for example, 30,000 pods (1000 nodes with 30 pods per node) represent ~150MB of data (~5kB/pod * 30k pods). So this test uses ReplicationControllers.
For this test (assuming N to be number of nodes in the cluster), we:
For the v1.3 release, we plan to extend this test by also creating Services, Deployments, DaemonSets, and other API objects.
Users are also very interested in how long it takes Kubernetes to schedule and start a pod. This is true not only upon initial creation, but also when a ReplicationController needs to create a replacement pod to take over from one whose node failed.
We (assuming N to be the number of nodes in the cluster):
While we could have decreased the “pod startup time” substantially by excluding for example waiting for report via watch, or creating pods directly rather than through ReplicationControllers, we believe that a broad definition that maps to the most realistic use cases is the best for real users to understand the performance they can expect from the system.
So what was the result?We run our tests on Google Compute Engine, setting the size of the master VM based on the size of the Kubernetes cluster. In particular for 1000-node clusters we use a n1-standard-32 VM for the master (32 cores, 120GB RAM).
The following two charts present a comparison of 99th percentile API call latencies for the Kubernetes 1.2 release and the 1.0 release on 100-node clusters. (Smaller bars are better)

We present results for LIST operations separately, since these latencies are significantly higher. Note that we slightly modified our tests in the meantime, so running current tests against v1.0 would result in higher latencies than they used to.

We also ran these tests against 1000-node clusters. Note: We did not support clusters larger than 100 on GKE, so we do not have metrics to compare these results to. However, customers have reported running on 1,000&#43; node clusters since Kubernetes 1.0.

Since LIST operations are significantly larger, we again present them separately: All latencies, in both cluster sizes, are well within our 1 second SLO.

The results for “pod startup latency” (as defined in the “Pod-Startup end-to-end latency” section) are presented in the following graph. For reference we are presenting also results from v1.0 for 100-node clusters in the first part of the graph.

As you can see, we substantially reduced tail latency in 100-node clusters, and now deliver low pod startup latency up to the largest cluster sizes we have measured. It is noteworthy that the metrics for 1000-node clusters, for both API latency and pod startup latency, are generally better than those reported for 100-node clusters just six months ago!
To make these significant gains in scale and performance over the past six months, we made a number of improvements across the whole system. Some of the most important ones are listed below.
Since most Kubernetes control logic operates on an ordered, consistent snapshot kept up-to-date by etcd watches (via the API server), a slight delay in that arrival of that data has no impact on the correct operation of the cluster. These independent controller loops, distributed by design for extensibility of the system, are happy to trade a bit of latency for an increase in overall throughput.
In Kubernetes 1.2 we exploited this fact to improve performance and scalability by adding an API server read cache. With this change, the API server’s clients can read data from an in-memory cache in the API server instead of reading it from etcd. The cache is updated directly from etcd via watch in the background. Those clients that can tolerate latency in retrieving data (usually the lag of cache is on the order of tens of milliseconds) can be served entirely from cache, reducing the load on etcd and increasing the throughput of the server. This is a continuation of an optimization begun in v1.1, where we added support for serving watch directly from the API server instead of etcd:[https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/apiserver-watch.md](https://github.com/kubernetes/kubernetes/blob/master/docs/proposals/apiserver-watch.md). 
Thanks to contributions from Wojciech Tyczynski at Google and Clayton Coleman and Timothy St. Clair at Red Hat, we were able to join careful system design with the unique advantages of etcd to improve the scalability and performance of Kubernetes. 
Kubernetes 1.2 also improved density from a pods-per-node perspective — for v1.2 we test and advertise up to 100 pods on a single node (vs 30 pods in the 1.1 release). This improvement was possible because of diligent work by the Kubernetes community through an implementation of the Pod Lifecycle Event Generator (PLEG).
The Kubelet (the Kubernetes node agent) has a worker thread per pod which is responsible for managing the pod’s lifecycle. In earlier releases each worker would periodically poll the underlying container runtime (Docker) to detect state changes, and perform any necessary actions to ensure the node’s state matched the desired state (e.g. by starting and stopping containers). As pod density increased, concurrent polling from each worker would overwhelm the Docker runtime, leading to serious reliability and performance issues (including additional CPU utilization which was one of the limiting factors for scaling up).
To address this problem we introduced a new Kubelet subcomponent — the PLEG — to centralize state change detection and generate lifecycle events for the workers. With concurrent polling eliminated, we were able to lower the steady-state CPU usage of Kubelet and the container runtime by 4x. This also allowed us to adopt a shorter polling period, so as to detect and react to changes more quickly. 
 
After surveying the Go JSON landscape and conducting some initial tests, we found the [ugorji codec](https://github.com/ugorji/go) library offered the most significant speedups - a 200% improvement in encoding and decoding JSON when using generated serializers, with a significant reduction in object allocations. After contributing fixes to the upstream library to deal with some of our complex structures, we switched Kubernetes and the go-etcd client library over. Along with some other important optimizations in the layers above and below JSON, we were able to slash the cost in CPU time of almost all API operations, especially reads. 

In both cases, the problem was debugged and/or fixed by Kubernetes community members, including Andy Goldstein and Jordan Liggitt from Red Hat, and Liang Mingqiang from NetEase. 
Of course, our job is not finished. We will continue to invest in improving Kubernetes performance, as we would like it to scale to many thousands of nodes, just like Google’s [Borg](http://static.googleusercontent.com/media/research.google.com/en//pubs/archive/43438.pdf). Thanks to our investment in testing infrastructure and our focus on how teams use containers in production, we have already identified the next steps on our path to improving scale. 
On deck for Kubernetes 1.3: 
In the last six months we’ve significantly improved Kubernetes scalability, allowing v1.2 to run 1000-node clusters with the same excellent responsiveness (as measured by our SLOs) as we were previously achieving only on much smaller clusters. But that isn’t enough — we want to push Kubernetes even further and faster. Kubernetes v1.3 will improve the system’s scalability and responsiveness further, while continuing to add features that make it easier to build and run the most demanding container-based applications. 
Please join our community and help us build the future of Kubernetes! There are many ways to participate. If you’re particularly interested in scalability, you’ll be interested in: 

---
[strong](https://www.blogger.com/null)We exclude operations on “events” since these are more like system logs and are not required for the system to operate properly.[strong](https://www.blogger.com/null)This is test/e2e/load.go from the Kubernetes github repository.[strong](https://www.blogger.com/null)This is test/e2e/density.go test from the Kubernetes github repository [strong](https://www.blogger.com/null)We are looking into optimizing this in the next release, but for now using a smaller master can result in significant (order of magnitude) performance degradation. We encourage anyone running benchmarking against Kubernetes or attempting to replicate these findings to use a similarly sized master, or performance will suffer.


	

	


