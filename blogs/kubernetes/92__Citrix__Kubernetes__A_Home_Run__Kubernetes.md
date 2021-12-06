|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/07/Citrix-Netscaler-And-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-07-14 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.4859202 &#43;0300 MSK m=&#43;4.074966001  |

#  Citrix &#43; Kubernetes = A Home Run  | Kubernetes

	
	
	
	
	*Editor’s note: today’s guest post is by Mikko Disini, a Director of Product Management at Citrix Systems, sharing their collaboration experience on a Kubernetes integration. *
Technical collaboration is like sports. If you work together as a team, you can go down the homestretch and pull through for a win. That’s our experience with the Google Cloud Platform team.
Recently, we approached Google Cloud Platform (GCP) to collaborate on behalf of Citrix customers and the broader enterprise market looking to migrate workloads. This migration required including the [NetScaler Docker load balancer](https://www.citrix.com/blogs/2016/06/20/the-best-docker-load-balancer-at-dockercon-in-seattle-this-week/), CPX, into Kubernetes nodes and resolving any issues with getting traffic into the CPX proxies.  
**Why NetScaler and Kubernetes?**
I wish all our experiences working together with a technical partner were as good as working with GCP. We had a list of issues to enable our use cases and were able to collaborate swiftly on a solution. To resolve these, GCP team offered in depth technical assistance, working with Citrix such that NetScaler CPX can spin up and take over as a client-side proxy running on each host. 
Next, NetScaler CPX needed to be inserted in the data path of GCP ingress load balancer so that NetScaler CPX can spread traffic to front end web servers. The NetScaler team made modifications so that NetScaler CPX listens to API server events and configures itself to create a VIP, IP table rules and server rules to take ingress traffic and load balance across front end applications. Google Cloud Platform team provided feedback and assistance to verify modifications made to overcome the technical hurdles. Done!
NetScaler CPX use case is supported in [Kubernetes 1.3](https://kubernetes.io/blog/2016/07/kubernetes-1-3-bridging-cloud-native-and-enterprise-workloads/). Citrix customers and the broader enterprise market will have the opportunity to leverage NetScaler with Kubernetes, thereby lowering the friction to move workloads to the cloud. 
You can learn more about NetScaler CPX [here](https://www.citrix.com/networking/microservices.html).
* -- Mikko Disini, Director of Product Management - NetScaler, Citrix Systems*


	

	


