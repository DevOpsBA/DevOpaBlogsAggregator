|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/04/Borg-Predecessor-To-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-04-23 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.5847416 &#43;0300 MSK m=&#43;5.173793601  |

#  Borg: The Predecessor to Kubernetes  | Kubernetes

	
	
	
	
	Google has been running containerized workloads in production for more than a decade. Whether it&#39;s service jobs like web front-ends and stateful servers, infrastructure systems like [Bigtable](http://research.google.com/archive/bigtable.html) and [Spanner](http://research.google.com/archive/spanner.html), or batch frameworks like [MapReduce](http://research.google.com/archive/mapreduce.html) and [Millwheel](http://research.google.com/pubs/pub41378.html), virtually everything at Google runs as a container. Today, we took the wraps off of Borg, Google’s long-rumored internal container-oriented cluster-management system, publishing details at the academic computer systems conference [Eurosys](http://eurosys2015.labri.fr/). You can find the paper [here](https://research.google.com/pubs/pub43438.html).
Kubernetes traces its lineage directly from Borg. Many of the developers at Google working on Kubernetes were formerly developers on the Borg project. We&#39;ve incorporated the best ideas from Borg in Kubernetes, and have tried to address some pain points that users identified with Borg over the years.
To give you a flavor, here are four Kubernetes features that came from our experiences with Borg:
Borg has a similar abstraction, called an alloc (short for “resource allocation”). Popular uses of allocs in Borg include running a web server that generates logs alongside a lightweight log collection process that ships the log to a cluster filesystem (not unlike fluentd or logstash); running a web server that serves data from a disk directory that is populated by a process that reads data from a cluster filesystem and prepares/stages it for the web server (not unlike a Content Management System); and running user-defined processing functions alongside a storage shard. Pods not only support these use cases, but they also provide an environment similar to running multiple processes in a single VM -- Kubernetes users can deploy multiple co-located, cooperating processes in a pod without having to give up the simplicity of a one-application-per-container deployment model.
Kubernetes supports more flexible collections than Borg by organizing pods using labels, which are arbitrary key/value pairs that users attach to pods (and in fact to any object in the system). Users can create groupings equivalent to Borg Jobs by using a “job:&lt;jobname&gt;” label on their pods, but they can also use additional labels to tag the service name, service instance (production, staging, test), and in general, any subset of their pods. A label query (called a “label selector”) is used to select which set of pods an operation should be applied to. Taken together, labels and [replication controllers](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/replication-controller.md) allow for very flexible update semantics, as well as for operations that span the equivalent of Borg Jobs.
Thanks to the advent of software-defined overlay networks such as [flannel](https://coreos.com/blog/introducing-rudder/) or those built into [public clouds](https://cloud.google.com/compute/docs/networking), Kubernetes is able to give every pod and service its own IP address. This removes the infrastructure complexity of managing ports, and allows developers to choose any ports they want rather than requiring their software to adapt to the ones chosen by the infrastructure. The latter point is crucial for making it easy to run off-the-shelf open-source applications on Kubernetes--pods can be treated much like VMs or physical hosts, with access to the full port space, oblivious to the fact that they may be sharing the same physical machine with other pods.
With the growing popularity of container-based microservice architectures, the lessons Google has learned from running such systems internally have become of increasing interest to the external DevOps community. By revealing some of the inner workings of our cluster manager Borg, and building our next-generation cluster manager as both an open-source project (Kubernetes) and a publicly available hosted service ([Google Container Engine](http://cloud.google.com/container-engine)), we hope these lessons can benefit the broader community outside of Google and advance the state-of-the-art in container scheduling and cluster management.  


	

	


