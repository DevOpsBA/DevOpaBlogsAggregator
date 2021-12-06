|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/12/From-Network-Policies-To-Security-Policies/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-12-08 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.985586 &#43;0300 MSK m=&#43;3.574628901  |

#  From Network Policies to Security Policies  | Kubernetes

	
	
	
	
	*Editor&#39;s note: Today’s post is by Bernard Van De Walle, Kubernetes Lead Engineer, at Aporeto, showing how they took a new approach to the Kubernetes network policy enforcement.*
**Kubernetes Network Policies **
Kubernetes supports a [new API for network policies](/docs/user-guide/networkpolicies/) that provides a sophisticated model for isolating applications and reducing their attack surface. This feature, which came out of the [SIG-Network group](https://github.com/kubernetes/community/wiki/SIG-Network), makes it very easy and elegant to define network policies by using the built-in labels and selectors Kubernetes constructs.
Kubernetes has left it up to third parties to implement these network policies and does not provide a default implementation.
We want to introduce a new way to think about “Security” and “Network Policies”. We want to show that security and reachability are two different problems, and that security policies defined using endpoints (pods labels for example) do not specifically need to be implemented using network primitives.
Most of us at [Aporeto](https://www.aporeto.com/) come from a Network/SDN background, and we knew how to implement those policies by using traditional networking and firewalling: Translating the pods identity and policy definitions to network constraints, such as IP addresses, subnets, and so forth.
However, we also knew from past experiences that using an external control plane also introduces a whole new set of challenges: This distribution of ACLs requires very tight synchronization between Kubernetes workers; and every time a new pod is instantiated, ACLs need to be updated on all other pods that have some policy related to the new pod. Very tight synchronization is fundamentally a quadratic state problem and, while shared state mechanisms can work at a smaller scale, they often have convergence, security, and eventual consistency issues in large scale clusters. 
**From Network Policies to Security Policies**
At Aporeto, we took a different approach to the network policy enforcement, by actually decoupling the network from the policy. We open sourced our solution as [Trireme](https://github.com/aporeto-inc/trireme), which translates the network policy to an authorization policy, and it implements a transparent authentication and authorization function for any communication between pods. Instead of using IP addresses to identify pods, it defines a cryptographically signed identity for each pod as the set of its associated labels. Instead of using ACLs or packet filters to enforce policy, it uses an authorization function where a container can only receive traffic from containers with an identity that matches the policy requirements. 
The authentication and authorization function in Trireme is overlaid on the TCP negotiation sequence. Identity (i.e. set of labels) is captured as a JSON Web Token (JWT), signed by local keys, and exchanged during the Syn/SynAck negotiation. The receiving worker validates that the JWTs are signed by a trusted authority (authentication step) and validates against a cached copy of the policy that the connection can be accepted. Once the connection is accepted, the rest of traffic flows through the Linux kernel and all of the protections that it can potentially offer (including conntrack capabilities if needed). The current implementation uses a simple user space process that captures the initial negotiation packets and attaches the authorization information as payload. The JWTs include nonces that are validated during the Ack packet and can defend against man-in-the-middle or replay attacks.

The Trireme implementation talks directly to the Kubernetes master without an external controller and receives notifications on policy updates and pod instantiations so that it can maintain a local cache of the policy and update the authorization rules as needed. There is no requirement for any shared state between Trireme components that needs to be synchronized. Trireme can be deployed either as a standalone process in every worker or by using [Daemon Sets](/docs/admin/daemons/). In the latter case, Kubernetes takes ownership of the lifecycle of the Trireme pods. 
Trireme&#39;s simplicity is derived from the separation of security policy from network transport. Policy enforcement is linked directly to the labels present on the connection, irrespective of the networking scheme used to make the pods communicate. This identity linkage enables tremendous flexibility to operators to use any networking scheme they like without tying security policy enforcement to network implementation details. Also, the implementation of security policy across the federated clusters becomes simple and viable.
**Kubernetes and Trireme deployment**
Kubernetes is unique in its ability to scale and provide an extensible security support for the deployment of containers and microservices. Trireme provides a simple, secure, and scalable mechanism for enforcing these policies. 
You can deploy and try Trireme on top of Kubernetes by using a provided Daemon Set. You&#39;ll need to modify some of the YAML parameters based on your cluster architecture. All the steps are described in detail in the [deployment GitHub folder](https://github.com/aporeto-inc/trireme-kubernetes/tree/master/deployment). The same folder contains an example 3-tier policy that you can use to test the traffic pattern.
To learn more, download the code, and engage with the project, visit:
*--Bernard Van De Walle, Kubernetes lead engineer, [Aporeto](https://www.aporeto.com/)*


	

	


