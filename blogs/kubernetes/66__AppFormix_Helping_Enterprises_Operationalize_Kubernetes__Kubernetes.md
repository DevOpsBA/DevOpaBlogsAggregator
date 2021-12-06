|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/03/Appformix-Helping-Enterprises/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-03-29 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.8550109 &#43;0300 MSK m=&#43;4.444058801  |

#  AppFormix: Helping Enterprises Operationalize Kubernetes  | Kubernetes

	
	
	
	
	*Today’s guest post is written Sumeet Singh, founder and CEO of [AppFormix](http://www.appformix.com/), a cloud infrastructure performance optimization service helping enterprise operators streamline their cloud operations on any OpenStack or Kubernetes cloud.*
If you run clouds for a living, you’re well aware that the tools we&#39;ve used since the client/server era for monitoring, analytics and optimization just don’t cut it when applied to the agile, dynamic and rapidly changing world of modern cloud infrastructure.
And, if you’re an operator of enterprise clouds, you know that implementing containers and container cluster management is all about giving your application developers a more agile, responsive and efficient cloud infrastructure. Applications are being rewritten and new ones developed – not for legacy environments where relatively static workloads are the norm, but for dynamic, scalable cloud environments. The dynamic nature of cloud native applications coupled with the shift to continuous deployment means that the demands placed by the applications on the infrastructure are constantly changing.
This shift necessitates infrastructure transparency and real-time monitoring and analytics. Without these key pieces, neither applications nor their underlying plumbing can deliver the low-latency user experience end users have come to expect.
  **AppFormix Architectural Review**
From an operational standpoint, it is necessary to understand how applications are consuming infrastructure resources in order to maximize ROI and guarantee SLAs. AppFormix software empowers operators and developers to monitor, visualize, and control how physical resources are utilized by cloud workloads. 
At the center of the software, the AppFormix Data Platform provides a distributed analysis engine that performs configurable, real-time evaluation of in-depth, high-resolution metrics. On each host, the resource-efficient AppFormix Agent collects and evaluates multi-layer metrics from the hardware, virtualization layer, and up to the application. Intelligent agents offer sub-second response times that make it possible to detect and solve problems before they start to impact applications and users. The raw data is associated with the elements that comprise a cloud-native environment: applications, virtual machines, containers, hosts. The AppFormix Agent then publishes metrics and events to a Data Manager that stores and forwards the data to Analytics modules. Events are based on predefined or dynamic conditions set by users or infrastructure operators to make sure that SLAs and policies are being met.
|  |
| Figure 1: Roll-up summary view of the Kubernetes cluster. Operators and Users can define their SLA policies and AppFormix provides with a real-time view of the health of all elements in the Kubernetes cluster.  |
|  |
| Figure 2: Real-Time visualization of telemetry from a Kubernetes node provides a quick overview of resource utilization on the host as well as resources consumed by the pods and containers. The user defined Labels make is easy to capture namespaces, and other metadata. |
Additional subsystems are the Policy Controller and Analytics. The Policy Controller manages policies for resource monitoring, analysis, and control. It also provides role-based access control. The Analytics modules analyze metrics and events produced by Data Platform, enabling correlation across multiple elements to provide higher-level information to operators and developers. The Analytics modules may also configure policies in Policy Controller in response to conditions in the infrastructure.
AppFormix organizes elements of cloud infrastructure around hosts and instances (either containers or virtual machines), and logical groups of such elements. AppFormix integrates with cloud platforms using Adapter modules that discover the physical and virtual elements in the environment and configure those elements into the Policy Controller.
**Integrating AppFormix with Kubernetes**
Enterprises often run many environments located on- or off-prem, as well as running different compute technologies (VMs, containers, bare metal). The analytics platform we’ve developed at AppFormix gives Kubernetes users a single pane of glass from which to monitor and manage container clusters in private and hybrid environments.
The AppFormix Kubernetes Adapter leverages the REST-based APIs of Kubernetes to discover nodes, pods, containers, services, and replication controllers. With the relational information about each element, Kubernetes Adapter is able to represent all of these elements in our system. A pod is a group of containers. A service and a replication controller are both different types of pod groups. In addition, using the watch endpoint, Kubernetes Adapter stays aware of changes to the environment.
**DevOps in the Enterprise with AppFormix**
With AppFormix, developers and operators can work collaboratively to optimize applications and infrastructure. Users can access a self-service IT experience that delivers visibility into CPU, memory, storage, and network consumption by each layer of the stack: physical hardware, platform, and application software. 
As you can see, we’re working hard to give Kubernetes users a useful, performant toolset for both OpenStack and Kubernetes environments that allows operators to deliver self-service IT to their application developers. We’re excited to be partner contributing to the Kubernetes ecosystem and community.
*-- Sumeet Singh, Founder and CEO, AppFormix*


	

	


