|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2019/02/06/poseidon-firmament-scheduler-flow-network-graph-based-scheduler/        |
| Tags              | [kubernetes]       |
| Date Create       | 2019-02-06 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.392109 &#43;0300 MSK m=&#43;1.981142801  |

# Poseidon-Firmament Scheduler – Flow Network Graph Based Scheduler | Kubernetes

	
	
	
	
	**Authors:** Deepak Vij (Huawei), Shivram Shrivastava (Huawei)
Cluster Management systems such as Mesos, Google Borg, Kubernetes etc. in a cloud scale datacenter environment (also termed as ***Datacenter-as-a-Computer*** or ***Warehouse-Scale Computing - WSC***) typically manage application workloads by performing tasks such as tracking machine live-ness, starting, monitoring, terminating workloads and more importantly using a **Cluster Scheduler** to decide on workload placements.
A **Cluster Scheduler** essentially performs the scheduling of workloads to compute resources – combining the global placement of work across the WSC environment makes the “warehouse-scale computer” more efficient, increases utilization, and saves energy. **Cluster Scheduler** examples are Google Borg, Kubernetes, Firmament, Mesos, Tarcil, Quasar, Quincy, Swarm, YARN, Nomad, Sparrow, Apollo etc.
In this blog post, we briefly describe the novel Firmament flow network graph based scheduling approach ([OSDI paper](https://www.usenix.org/conference/osdi16/technical-sessions/presentation/gog)) in Kubernetes. We specifically describe the Firmament Scheduler and how it integrates with the Kubernetes cluster manager using Poseidon as the integration glue. We have seen extremely impressive scheduling throughput performance benchmarking numbers with this novel scheduling approach. Originally, Firmament Scheduler was conceptualized, designed and implemented by University of Cambridge researchers, [Malte Schwarzkopf](http://www.malteschwarzkopf.de/) &amp; [Ionel Gog](http://ionelgog.org/).
At a very high level, [Poseidon-Firmament scheduler](/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/) augments the current Kubernetes scheduling capabilities by incorporating novel flow network graph based scheduling capabilities alongside the default Kubernetes Scheduler. It models the scheduling problem as a constraint-based optimization over a flow network graph – by reducing scheduling to a min-cost max-flow optimization problem. Due to the inherent rescheduling capabilities, the new scheduler enables a globally optimal scheduling environment that constantly keeps refining the workloads placements dynamically.
Flow graph scheduling based [Poseidon-Firmament scheduler](/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/) provides the following key advantages:
Firmament scheduler runs a min-cost flow algorithm over the flow network to find an optimal flow, from which it extracts the implied workload (pod placements). A flow network is a directed graph whose arcs carry flow from source nodes (i.e. pod nodes) to a sink node. A cost and capacity associated with each arc constrain the flow, and specify preferential routes for it.
Figure 1 below shows an example of a flow network for a cluster with two tasks (workloads or pods) and four machines (nodes) – each workload on the left hand side, is a source of one unit of flow. All such flow must be drained into the sink node (S) for a feasible solution to the optimization problem.




Poseidon is a service that acts as the integration glue for the Firmament scheduler with Kubernetes. It augments the current Kubernetes scheduling capabilities by incorporating new flow network graph based Firmament scheduling capabilities alongside the default Kubernetes Scheduler; multiple schedulers running simultaneously. Figure 2 below describes the high level overall design as far as how Poseidon integration glue works in conjunction with the underlying Firmament flow network graph based scheduler.




As part of the Kubernetes multiple schedulers support, each new pod is typically scheduled by the default scheduler, but Kubernetes can be instructed to use another scheduler by specifying the name of another custom scheduler (in our case, [Poseidon-Firmament](/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/)) at the time of pod deployment. In this case, the default scheduler will ignore that Pod and allow Poseidon scheduler to schedule the Pod to a relevant node.
[Poseidon-Firmament scheduler](/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/) enables extremely high throughput scheduling environment at scale due to its bulk scheduling approach superiority versus K8s pod-at-a-time approach. In our extensive tests, we have observed substantial throughput benefits as long as resource requirements (CPU/Memory) for incoming Pods is uniform across jobs (Replicasets/Deployments/Jobs), mainly due to efficient amortization of work across jobs.
Although, [Poseidon-Firmament scheduler](/docs/concepts/extend-kubernetes/poseidon-firmament-alternate-scheduler/) is capable of scheduling various types of workloads (service, batch, etc.), following are the few use cases where it excels the most:
Currently Poseidon-Firmament project is an incubation project. Alpha Release is available at [https://github.com/kubernetes-sigs/poseidon](https://github.com/kubernetes-sigs/poseidon).


	

	


