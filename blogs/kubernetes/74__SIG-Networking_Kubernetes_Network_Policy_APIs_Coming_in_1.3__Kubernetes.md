|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/04/Kubernetes-Network-Policy-APIs/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-04-18 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.7415422 &#43;0300 MSK m=&#43;4.330589401  |

#  SIG-Networking: Kubernetes Network Policy APIs Coming in 1.3  | Kubernetes

	
	
	
	
	*Editor’s note: This week we’re featuring [Kubernetes Special Interest Groups](https://github.com/kubernetes/kubernetes/wiki/Special-Interest-Groups-(SIGs)); Today’s post is by the Network-SIG team describing network policy APIs coming in 1.3 - policies for security, isolation and multi-tenancy.*
The [Kubernetes network SIG](https://kubernetes.slack.com/messages/sig-network/) has been meeting regularly since late last year to work on bringing network policy to Kubernetes and we’re starting to see the results of this effort.
One problem many users have is that the open access network policy of Kubernetes is not suitable for applications that need more precise control over the traffic that accesses a pod or service. Today, this could be a multi-tier application where traffic is only allowed from a tier’s neighbor. But as new Cloud Native applications are built by composing microservices, the ability to control traffic as it flows among these services becomes even more critical.
In most IaaS environments (both public and private) this kind of control is provided by allowing VMs to join a ‘security group’ where traffic to members of the group is defined by a network policy or Access Control List (ACL) and enforced by a network packet filter.
The Network SIG started the effort by identifying [specific use case scenarios](https://docs.google.com/document/d/1blfqiH4L_fpn33ZrnQ11v7LcYP0lmpiJ_RaapAPBbNU/edit?pref=2&amp;pli=1#) that require basic network isolation for enhanced security. Getting the API right for these simple and common use cases is important because they are also the basis for the more sophisticated network policies necessary for multi-tenancy within Kubernetes.
From these scenarios several possible approaches were considered and a minimal [policy specification](https://docs.google.com/document/d/1qAm-_oSap-f1d6a-xRTj6xaH1sYQBfK36VyjB5XOZug/edit) was defined. The basic idea is that if isolation were enabled on a per namespace basis, then specific pods would be selected where specific traffic types would be allowed.
The simplest way to quickly support this experimental API is in the form of a ThirdPartyResource extension to the API Server, which is possible today in Kubernetes 1.2.
If you’re not familiar with how this works, the Kubernetes API can be extended by defining ThirdPartyResources that create a new API endpoint at a specified URL.
```kind: ThirdPartyResource

apiVersion: extensions/v1beta1

metadata:

 &amp;nbsp;name: network-policy.net.alpha.kubernetes.io

description: &#34;Network policy specification&#34;

versions:

- name: v1alpha1
``````$kubectl create -f third-party-res-def.yaml
```This will create an API endpoint (one for each namespace):
```/net.alpha.kubernetes.io/v1alpha1/namespace/default/networkpolicys/
```Third party network controllers can now listen on these endpoints and react as necessary when resources are created, modified or deleted. *Note: With the upcoming release of Kubernetes 1.3 - when the Network Policy API is released in beta form - there will be no need to create a ThirdPartyResource API endpoint as shown above.* 
Network isolation is off by default so that all pods can communicate as they normally do. However, it’s important to know that once network isolation is enabled, all traffic to all pods, in all namespaces is blocked, which means that enabling isolation is going to change the behavior of your pods
Network isolation is enabled by defining the *network-isolation* annotation on namespaces as shown below:
```net.alpha.kubernetes.io/network-isolation: [on | off]
```Once network isolation is enabled, explicit network policies **must be applied** to enable pod communication.
A policy specification can be applied to a namespace to define the details of the policy as shown below:
```POST /apis/net.alpha.kubernetes.io/v1alpha1/namespaces/tenant-a/networkpolicys/


{

&#34;kind&#34;: &#34;NetworkPolicy&#34;,

&#34;metadata&#34;: {

&#34;name&#34;: &#34;pol1&#34;

},

&#34;spec&#34;: {

&#34;allowIncoming&#34;: {

&#34;from&#34;: [

{ &#34;pods&#34;: { &#34;segment&#34;: &#34;frontend&#34; } }

],

&#34;toPorts&#34;: [

{ &#34;port&#34;: 80, &#34;protocol&#34;: &#34;TCP&#34; }

]

},

&#34;podSelector&#34;: { &#34;segment&#34;: &#34;backend&#34; }

}

}
```In this example, the ‘ **tenant-a** ’ namespace would get policy ‘ **pol1** ’ applied as indicated. Specifically, pods with the **segment** label ‘ **backend** ’ would allow TCP traffic on port 80 from pods with the **segment** label ‘ **frontend** ’ to be received.
Today, [Romana](http://romana.io/), [OpenShift](https://www.openshift.com/), [OpenContrail](http://www.opencontrail.org/) and [Calico](http://projectcalico.org/) support network policies applied to namespaces and pods. Cisco and VMware are working on implementations as well. Both Romana and Calico demonstrated these capabilities with Kubernetes 1.2 recently at KubeCon. You can watch their presentations here: [Romana](https://www.youtube.com/watch?v=f-dLKtK6qCs) ([slides](http://www.slideshare.net/RomanaProject/kubecon-london-2016-ronana-cloud-native-sdn)), [Calico](https://www.youtube.com/watch?v=p1zfh4N4SX0) ([slides](http://www.slideshare.net/kubecon/kubecon-eu-2016-secure-cloudnative-networking-with-project-calico)). 
**How does it work?**
Each solution has their their own specific implementation details. Today, they rely on some kind of on-host enforcement mechanism, but future implementations could also be built that apply policy on a hypervisor, or even directly by the network itself. 
External policy control software (specifics vary across implementations) will watch the new API endpoint for pods being created and/or new policies being applied. When an event occurs that requires policy configuration, the listener will recognize the change and a controller will respond by configuring the interface and applying the policy.  The diagram below shows an API listener and policy controller responding to updates by applying a network policy locally via a host agent. The network interface on the pods is configured by a CNI plugin on the host (not shown).

If you’ve been holding back on developing applications with Kubernetes because of network isolation and/or security concerns, these new network policies go a long way to providing the control you need. No need to wait until Kubernetes 1.3 since network policy is available now as an experimental API enabled as a ThirdPartyResource.
If you’re interested in Kubernetes and networking, there are several ways to participate - join us at:
The Networking “Special Interest Group,” which meets bi-weekly at 3pm (15h00) Pacific Time at [SIG-Networking hangout](https://zoom.us/j/5806599998). 
*--Chris Marino, Co-Founder, Pani Networks*


	

	


