|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/09/Kubernetes-Performance-Measurements-And/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-09-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.2449562 &#43;0300 MSK m=&#43;4.834006301  |

#  Kubernetes Performance Measurements and Roadmap  | Kubernetes

	
	
	
	
	No matter how flexible and reliable your container orchestration system is, ultimately, you have some work to be done, and you want it completed quickly. For big problems, a common answer is to just throw more machines at the problem. After all, more compute = faster, right?
Interestingly, adding more nodes is a little like the [tyranny of the rocket equation](http://www.nasa.gov/mission_pages/station/expeditions/expedition30/tryanny.html) - in some systems, adding more machines can actually make your processing slower. However, unlike the rocket equation, we can do better. Kubernetes in v1.0 version supports clusters with up to 100 nodes. However, we have a goal to 10x the number of nodes we will support by the end of 2015. This blog post will cover where we are and how we intend to achieve the next level of performance.
The first question we need to answer is: “what does it mean that Kubernetes can manage an N-node cluster?” Users expect that it will handle all operations “reasonably quickly,” but we need a precise definition of that. We decided to define performance and scalability goals based on the following two metrics:
Note that for “pod startup time” we explicitly assume that all images necessary to run a pod are already pre-pulled on the machine where it will be running. In our experiments, there is a high degree of variability (network throughput, size of image, etc) between images, and these variations have little to do with Kubernetes’ overall performance.
The decision to choose those metrics was made based on our experience spinning up 2 billion containers a week at Google. We explicitly want to measure the latency of user-facing flows since that’s what customers will actually care about.
To monitor performance improvements and detect regressions we set up a continuous testing infrastructure. Every 2-3 hours we create a 100-node cluster from [HEAD](https://github.com/kubernetes/kubernetes) and run our scalability tests on it. We use a GCE n1-standard-4 (4 cores, 15GB of RAM) machine as a master and n1-standard-1 (1 core, 3.75GB of RAM) machines for nodes.
In scalability tests, we explicitly focus only on the full-cluster case (full N-node cluster is a cluster with 30 * N pods running in it) which is the most demanding scenario from a  performance point of view. To reproduce what a customer might actually do, we run through the following steps:
It is worth emphasizing that the main parts of the test are done on full clusters (30 pods per node, 100 nodes) - starting a pod in an empty cluster, even if it has 100 nodes will be much faster.
To measure pod startup latency we are using very simple pods with just a single container running the “gcr.io/google_containers/pause:go” image, which starts and then sleeps forever. The container is guaranteed to be already pre-pulled on nodes (we use it as the so-called pod-infra-container).
The following table contains percentiles (50th, 90th and 99th) of pod startup time in 100-node clusters which are 10%, 25%, 50% and 100% full.
As for api-responsiveness, the following graphs present 50th, 90th and 99th percentiles of latencies of API calls grouped by kind of operation and resource type. However, note that this also includes internal system API calls, not just those issued by users (in this case issued by the test itself).



Some resources only appear on certain graphs, based on what was running during that operation (e.g. no namespace was put at that time).
As you can see in the results, we are ahead of target for our 100-node cluster with pod startup time even in a fully-packed cluster occurring 14% faster in the 99th percentile than 5 seconds. It’s interesting to point out that  LISTing pods is significantly slower than any other operation. This makes sense: in a full cluster there are 3000 pods and each of pod is roughly few kilobytes of data, meaning megabytes of data that need to processed for each LIST.
#####Work done and some future plans
The initial performance work to make 100-node clusters stable enough to run any tests on them involved a lot of small fixes and tuning, including increasing the limit for file descriptors in the apiserver and reusing tcp connections between different requests to etcd.
However, building a stable performance test was just step one to increasing the number of nodes our cluster supports by tenfold. As a result of this work, we have already taken on significant effort to remove future bottlenecks, including:
Looking further out to our 1000-node cluster goal, proposed improvements include:
This is by no means an exhaustive list. We will be adding new elements (or removing existing ones) based on the observed bottlenecks while running the existing scalability tests and newly-created ones. If there are particular use cases or scenarios that you’d like to see us address, please join in!


	

	


