|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/11/Visualize-Kubelet-Performance-With-Node-Dashboard/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-11-17 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.0111397 &#43;0300 MSK m=&#43;3.600182801  |

#  Visualize Kubelet Performance with Node Dashboard  | Kubernetes

	
	
	
	
	*Since this article was published, the Node Performance Dashboard was retired and is no longer available.*
*This retirement happened in early 2019, as part of the* ```kubernetes/contrib```*[repository deprecation](https://github.com/kubernetes-retired/contrib/issues/3007)*.
In Kubernetes 1.4, we introduced a new node performance analysis tool, called the *node performance dashboard*, to visualize and explore the behavior of the Kubelet in much richer details. This new feature will make it easy to understand and improve code performance for Kubelet developers, and lets cluster maintainer set configuration according to provided Service Level Objectives (SLOs).
**Background**
A Kubernetes cluster is made up of both master and worker nodes. The master node manages the cluster’s state, and the worker nodes do the actual work of running and managing pods. To do so, on each worker node, a binary, called [Kubelet](/docs/admin/kubelet/), watches for any changes in pod configuration, and takes corresponding actions to make sure that containers run successfully. High performance of the Kubelet, such as low latency to converge with new pod configuration and efficient housekeeping with low resource usage, is essential for the entire Kubernetes cluster. To measure this performance, Kubernetes uses [end-to-end (e2e) tests](https://github.com/kubernetes/kubernetes/blob/master/docs/devel/e2e-tests.md#overview) to continuously monitor benchmark changes of latest builds with new features.
**Kubernetes SLOs are defined by the following benchmarks** :
*** API responsiveness** : 99% of all API calls return in less than 1s.*** Pod startup time** : 99% of pods and their containers (with pre-pulled images) start within 5s.
Prior to 1.4 release, we’ve only measured and defined these at the cluster level, opening up the risk that other factors could influence the results. Beyond these, we also want to have more performance related SLOs such as the maximum number of pods for a specific machine type allowing maximum utilization of your cluster. In order to do the measurement correctly, we want to introduce a set of tests isolated to just a node’s performance. In addition, we aim to collect more fine-grained resource usage and operation tracing data of Kubelet from the new tests.
**Data Collection**
The node specific density and resource usage tests are now added into e2e-node test set since 1.4. The resource usage is measured by a standalone cAdvisor pod for flexible monitoring interval (comparing with Kubelet integrated cAdvisor). The performance data, such as latency and resource usage percentile, are recorded in persistent test result logs. The tests also record time series data such as creation time, running time of pods, as well as real-time resource usage. Tracing data of Kubelet operations are recorded in its log stored together with test results.
**Node Performance Dashboard**
Since Kubernetes 1.4, we are continuously building the newest Kubelet code and running node performance tests. The data is collected by our new performance dashboard available at [node-perf-dash.k8s.io](http://node-perf-dash.k8s.io/). Figure 1 gives a preview of the dashboard. You can start to explore it by selecting a test, either using the drop-down list of short test names (region (a)) or by choosing test options one by one (region (b)). The test details show up in region (c) containing the full test name from Ginkgo (the Go test framework used by Kubernetes). Then select a node type (image and machine) in region (d).
| |
| Figure 1. Select a test to display in node performance dashboard. |
The &#34;BUILDS&#34; page exhibits the performance data across different builds (Figure 2). The plots include pod startup latency, pod creation throughput, and CPU/memory usage of Kubelet and runtime (currently Docker). In this way it’s easy to monitor the performance change over time as new features are checked in.
|  |
| Figure 2. Performance data across different builds. |
**Compare Different Node Configurations**
It’s always interesting to compare the performance between different configurations, such as comparing startup latency of different machine types, different numbers of pods, or comparing resource usage of hosting different number of pods. The dashboard provides a convenient way to do this. Just click the &#34;Compare it&#34; button the right up corner of test selection menu (region (e) in Figure 1). The selected tests will be added to a comparison list in the &#34;COMPARISON&#34; page, as shown in Figure 3. Data across a series of builds are aggregated to a single value to facilitate comparison and are displayed in bar charts.
|  |
| Figure 3. Compare different test configurations. |
**Time Series and Tracing: Diving Into Performance Data**
Pod startup latency is an important metric for Kubelet, especially when creating a large number of pods per node. Using the dashboard you can see the change of latency, for example, when creating 105 pods, as shown in Figure 4. When you see the highly variable lines, you might expect that the variance is due to different builds. However, as these test here were run against the same Kubernetes code, we can conclude the variance is due to performance fluctuation. The variance is close to 40s when we compare the 99% latency of build #162 and #173, which is very large. To drill into the source of the fluctuation, let’s check out the &#34;TIME SERIES&#34; page.
|  |
| Figure 4. Pod startup latency when creating 105 pods. |
Looking specifically at build #162, we are able to see that the tracing data plotted in the pod creation latency chart (Figure 5). Each curve is an accumulated histogram of the number of pod operations which have already arrive at a certain tracing probe. The timestamp of tracing pod is either collected from the performance tests or by parsing the Kubelet log. Currently we collect the following tracing data:
The time series chart illustrates that it is taking a long time for the status manager to update pod status (the data of &#34;running&#34; is not shown since it overlaps with &#34;pod_status_running&#34;). We figure out this latency is introduced due to the query per second (QPS) limits of Kubelet to the API server (default is 5). After being aware of this, we find in additional tests that by increasing QPS limits, curve &#34;running&#34; gradually converges with &#34;pod_running&#39;, and results in much lower latency. Therefore the previous e2e test pod startup results reflect the combined latency of both Kubelet and time of uploading status, the performance of Kubelet is thus under-estimated.
|  |
| Figure 5. Time series page using data from build #162. |
Further, by comparing the time series data of build #162 (Figure 5) and build #173 (Figure 6), we find that the performance pod startup latency fluctuation actually happens during updating pod statuses. Build #162 has several straggler &#34;pod_status_running&#34; events with a long latency tails. It thus provides useful ideas for future optimization. 
|  |
| Figure 6. Pod startup latency of build #173. |
In future we plan to use events in Kubernetes which has a fixed log format to collect tracing data more conveniently. Instead of extracting existing log entries, then you can insert your own tracing probes inside Kubelet and obtain the break-down latency of each segment. 
You can check the latency between any two probes across different builds in the “TRACING” page, as shown in Figure 7. For example, by selecting &#34;pod_config_change&#34; as the start probe, and &#34;pod_status_running&#39; as the end probe, it gives the latency variance of Kubelet over continuous builds without status updating overhead. With this feature, developers are able to monitor the performance change of a specific part of code inside Kubelet.
|  |
| Figure 7. Plotting latency between any two probes. |
**Future Work**
The [node performance dashboard](http://node-perf-dash.k8s.io/) is a brand new feature. It is still alpha version under active development. We will keep optimizing the data collecting and visualization, providing more tests, metrics and tools to the developers and the cluster maintainers. 
Please join our community and help us build the future of Kubernetes! If you’re particularly interested in nodes or performance testing, participate by chatting with us in our [Slack channel](https://kubernetes.slack.com/messages/sig-scale/) or join our meeting which meets every Tuesday at 10 AM PT on this [SIG-Node Hangout](https://github.com/kubernetes/community/tree/master/sig-node).
*--Zhou Fang, Software Engineering Intern, Google*


	

	


