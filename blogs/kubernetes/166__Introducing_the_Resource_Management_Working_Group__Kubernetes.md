|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/09/Introducing-Resource-Management-Working/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-09-21 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.40967 &#43;0300 MSK m=&#43;2.998709601  |

#  Introducing the Resource Management Working Group  | Kubernetes

	
	
	
	
	***Editor&#39;s note: today&#39;s post is by Jeremy Eder, Senior Principal Software Engineer at Red Hat, on the formation of the Resource Management Working Group***
Kubernetes has evolved to support diverse and increasingly complex classes of applications. We can onboard and scale out modern, cloud-native web applications based on microservices, batch jobs, and stateful applications with persistent storage requirements.
However, there are still opportunities to improve Kubernetes; for example, the ability to run workloads that require specialized hardware or those that perform measurably better when hardware topology is taken into account. These conflicts can make it difficult for application classes (particularly in established verticals) to adopt Kubernetes.
We see an unprecedented opportunity here, with a high cost if it’s missed. The Kubernetes ecosystem must create a consumable path forward to the next generation of system architectures by catering to needs of as-yet unserviced workloads in meaningful ways. The Resource Management Working Group, along with other SIGs, must demonstrate the vision customers want to see, while enabling solutions to run well in a fully integrated, thoughtfully planned end-to-end stack.
 
Kubernetes Working Groups are created when a particular challenge requires cross-SIG collaboration. The Resource Management Working Group, for example, works primarily with sig-node and sig-scheduling to drive support for additional resource management capabilities in Kubernetes. We make sure that key contributors from across SIGs are frequently consulted because working groups are not meant to make system-level decisions on behalf of any SIG.
 
An example and key benefit of this is the working group’s relationship with sig-node.  We were able to ensure completion of several releases of node reliability work (complete in 1.6) before contemplating feature design on top. Those designs are use-case driven: research into technical requirements for a variety of workloads, then sorting based on measurable impact to the largest cross-section.
One of the working group’s key design tenets is that user experience must remain clean and portable, while still surfacing infrastructure capabilities that are required by businesses and applications.
 
While not representing any commitment, we hope in the fullness of time that Kubernetes can optimally run financial services workloads, machine learning/training, grid schedulers, map-reduce, animation workloads, and more. As a use-case driven group, we account for potential application integration that can also facilitate an ecosystem of complementary independent software vendors to flourish on top of Kubernetes.

Kubernetes covers generic web hosting capabilities very well, so why go through the effort of expanding workload coverage for Kubernetes at all? The fact is that workloads elegantly covered by Kubernetes today, only represent a fraction of the world’s compute usage. We have a tremendous opportunity to safely and methodically expand upon the set of workloads that can run optimally on Kubernetes.
To date, there’s demonstrable progress in the areas of expanded workload coverage:
As a consequence, we began advocating for increasing the scope of workloads covered by Kubernetes, from overall concepts to specific features. Our aim is to put control and choice in users hands, helping them move with confidence towards whatever infrastructure strategy they choose. In this advocacy, we quickly found a large group of like-minded companies interested in broadening the types of workloads that Kubernetes can orchestrate. And thus the working group was born.
After extensive development/feature [discussions](https://docs.google.com/document/d/1p7scsTPzPyouktBFTxu4RhRwW8yUn5Lj7VGY9SaOo-8/edit?ts=5824ee1f) during the Kubernetes Developer Summit 2016 after [CloudNativeCon | KubeCon Seattle](http://events.linuxfoundation.org/events/kubecon/program/schedule), we decided to [formalize](https://groups.google.com/d/msg/kubernetes-dev/Sb0VlXOM8eQ/La3YCe2-CgAJ) our loosely organized group. In January 2017, the Kubernetes *[Resource Management Working Group](https://github.com/kubernetes/community/tree/master/wg-resource-management)* was formed. This group (led by Derek Carr from Red Hat and Vishnu Kannan from Google) was originally cast as a temporary initiative to provide guidance back to sig-node and sig-scheduling (primarily). However, due to the cross-cutting nature of the goals within the working group, and the depth of [roadmap](https://docs.google.com/spreadsheets/d/1NWarIgtSLsq3izc5wOzV7ItdhDNRd-6oBVawmvs-LGw/edit) quickly uncovered, the Resource Management Working Group became its own entity within the first few months.
Recently, Brian Grant from Google (@bgrant0607) posted the following image on his [Twitter feed](https://twitter.com/bgrant0607/status/862091393723842561). This image helps to explain the role of each SIG, and shows where the Resource Management Working Group fits into the overall project organization.
{.big-img}
To help bootstrap this effort, the Resource Management Working Group had its first face-to-face kickoff meeting in May 2017. Thanks to Google for hosting!

Folks from Intel, NVIDIA, Google, IBM, Red Hat. and Microsoft (among others) participated. 
You can read the outcomes of that 3-day meeting [here](https://docs.google.com/document/d/13_nk75eItkpbgZOt62In3jj0YuPbGPC_NnvSCHpgvUM/edit).
The group’s prioritized list of features for increasing workload coverage on Kubernetes enumerated in the [charter](https://github.com/kubernetes/community/tree/master/wg-resource-management) of the Resource Management Working group includes:
The set of initially targeted use-cases share one or more of the following characteristics:
In the months leading up to our recent face-to-face, we had discussed how to safely abstract resources in a way that retains portability and clean user experience, while still meeting application requirements. The working group came away with a multi-release [roadmap](https://docs.google.com/spreadsheets/d/1NWarIgtSLsq3izc5wOzV7ItdhDNRd-6oBVawmvs-LGw/edit) that included 4 short- to mid-term targets with great overlap between target workloads:
Our charter document includes a [Contact Us](https://github.com/kubernetes/community/tree/master/wg-resource-management#contact-us) section with links to our mailing list, Slack channel, and Zoom meetings. Recordings of previous meetings are uploaded to [Youtube](https://www.youtube.com/channel/UCyfvrmhAGcsFlJeGgZQvZ6g). We plan to discuss these topics and more at the 2017 Kubernetes Developer Summit at [CloudNativeCon | KubeCon](http://events.linuxfoundation.org/events/cloudnativecon-and-kubecon-north-america) in Austin. Please come and join one of our meetings (users, customers, software and hardware vendors are all welcome) and contribute to the working group!


	

	


