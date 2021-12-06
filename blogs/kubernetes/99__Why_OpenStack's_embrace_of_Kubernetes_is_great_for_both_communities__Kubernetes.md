|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/07/openstack-kubernetes-communities/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-07-26 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.3138614 &#43;0300 MSK m=&#43;3.902906201  |

#  Why OpenStack&#39;s embrace of Kubernetes is great for both communities  | Kubernetes

	
	
	
	
	Today, [Mirantis](https://www.mirantis.com/), the leading contributor to [OpenStack](http://stackalytics.com/?release=mitaka), [announced](https://techcrunch.com/2016/07/25/openstack-will-soon-be-able-to-run-on-top-of-kubernetes/) that it will re-write its private cloud platform to use Kubernetes as its underlying orchestration engine. We think this is a great step forward for both the OpenStack and Kubernetes communities. With Kubernetes under the hood, OpenStack users will benefit from the tremendous efficiency, manageability and resiliency that Kubernetes brings to the table, while positioning their applications to use more cloud-native patterns. The Kubernetes community, meanwhile, can feel confident in their choice of orchestration framework, while gaining the ability to manage both container- and VM-based applications from a single platform.
**The Path to Cloud Native**
Google spent over ten years developing, applying and refining the principles of cloud native computing. A cloud-native application is:
These principles have enabled us to build the largest, most efficient, most powerful cloud infrastructure in the world, which anyone can access via [Google Cloud Platform](http://cloud.google.com/). They are the same principles responsible for the recent surge in popularity of Linux containers. Two years ago, we open-sourced Kubernetes to spur adoption of containers and scalable, microservices-based applications, and the recently released [Kubernetes version 1.3](https://kubernetes.io/blog/2016/07/kubernetes-1-3-bridging-cloud-native-and-enterprise-workloads/) introduces a number of features to bridge enterprise and cloud native workloads. We expect that adoption of cloud-native principles will drive the same benefits within the OpenStack community, as well as smoothing the path between OpenStack and the public cloud providers that embrace them.
**Making OpenStack better**
We hear from enterprise customers that they want to move towards cloud-native infrastructure and application patterns. Thus, it is hardly surprising that OpenStack would also move in this direction [1], with large OpenStack users such as [eBay](http://fortune.com/2016/04/23/ebay-parlays-new-age-tools/) and [GoDaddy](http://thenewstack.io/tns-analysts-show-95-consider-containerizing-openstack/) adopting Kubernetes as key components of their stack. Kubernetes and cloud-native patterns will improve OpenStack lifecycle management by enabling rolling updates, versioning, and canary deployments of new components and features. In addition, OpenStack users will benefit from self-healing infrastructure, making OpenStack easier to manage and more resilient to the failure of core services and individual compute nodes. Finally, OpenStack users will realize the developer and resource efficiencies that come with a container-based infrastructure.
**OpenStack is a great tool for Kubernetes users**
Conversely, incorporating Kubernetes into OpenStack will give Kubernetes users access to a robust framework for deploying and managing applications built on virtual machines. As users move to the cloud-native model, they will be faced with the challenge of managing hybrid application architectures that contain some mix of virtual machines and Linux containers. The combination of Kubernetes and OpenStack means that they can do so on the same platform using a common set of tools.
We are excited by the ever increasing momentum of the cloud-native movement as embodied by Kubernetes and related projects, and look forward to working with Mirantis, its partner Intel, and others within the OpenStack community to brings the benefits of cloud-native to their applications and infrastructure.
*--Martin Buhr, Product Manager, Strategic Initiatives, Google*
[1] Check out the announcement of Kubernetes-OpenStack Special Interest Group [here](https://kubernetes.io/blog/2016/04/introducing-kubernetes-openstack-sig), and a great talk about OpenStack on Kubernetes by CoreOS CEO Alex Polvi at the most recent OpenStack summit [here](https://www.youtube.com/watch?v=e-j9FOO-i84).


	

	


