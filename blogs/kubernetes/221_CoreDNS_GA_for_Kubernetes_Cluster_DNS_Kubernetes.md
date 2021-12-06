|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/07/10/coredns-ga-for-kubernetes-cluster-dns/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-07-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.7227543 &#43;0300 MSK m=&#43;2.311790001  |

# CoreDNS GA for Kubernetes Cluster DNS | Kubernetes

	
	
	
	
	**Author**: John Belamaric (Infoblox)
**Editor’s note: this post is part of a [series of in-depth articles](https://kubernetes.io/blog/2018/06/27/kubernetes-1.11-release-announcement/) on what’s new in Kubernetes 1.11**
In Kubernetes 1.11, [CoreDNS](https://coredns.io) has reached General Availability (GA) for DNS-based service discovery, as an alternative to the kube-dns addon. This means that CoreDNS will be offered as an option in upcoming versions of the various installation tools. In fact, the kubeadm team chose to make it the default option starting with Kubernetes 1.11.
DNS-based service discovery has been part of Kubernetes for a long time with the kube-dns cluster addon. This has generally worked pretty well, but there have been some concerns around the reliability, flexibility and security of the implementation.
CoreDNS is a general-purpose, authoritative DNS server that provides a backwards-compatible, but extensible, integration with Kubernetes. It resolves the issues seen with kube-dns, and offers a number of unique features that solve a wider variety of use cases.
In this article, you will learn about the differences in the implementations of kube-dns and CoreDNS, and some of the helpful extensions offered by CoreDNS.
In kube-dns, several containers are used within a single pod: ```kubedns```, ```dnsmasq```, and ```sidecar```. The ```kubedns```
container watches the Kubernetes API and serves DNS records based on the [Kubernetes DNS specification](https://github.com/kubernetes/dns/blob/master/docs/specification.md), ```dnsmasq``` provides caching and stub domain support, and ```sidecar``` provides metrics and health checks.
This setup leads to a few issues that have been seen over time. For one, security vulnerabilities in ```dnsmasq``` have led to the need
for a security-patch release of Kubernetes in the past. Additionally, because ```dnsmasq``` handles the stub domains,
but ```kubedns``` handles the External Services, you cannot use a stub domain in an external service, which is very
limiting to that functionality (see [dns#131](https://github.com/kubernetes/dns/issues/131)).
All of these functions are done in a single container in CoreDNS, which is running a process written in Go. The
different plugins that are enabled replicate (and enhance) the functionality found in kube-dns.
In kube-dns, you can [modify a ConfigMap](https://kubernetes.io/blog/2017/04/configuring-private-dns-zones-upstream-nameservers-kubernetes/) to change the behavior of your service discovery. This allows the addition of
features such as serving stub domains, modifying upstream nameservers, and enabling federation.
In CoreDNS, you similarly can modify the ConfigMap for the CoreDNS [Corefile](https://coredns.io/2017/07/23/corefile-explained/) to change how service discovery
works. This Corefile configuration offers many more options than you will find in kube-dns, since it is the
primary configuration file that CoreDNS uses for configuration of all of its features, even those that are not
Kubernetes related.
When upgrading from kube-dns to CoreDNS using ```kubeadm```, your existing ConfigMap will be used to generate the
customized Corefile for you, including all of the configuration for stub domains, federation, and upstream nameservers. See [Using CoreDNS for Service Discovery](/docs/tasks/administer-cluster/coredns/) for more details.
There are several open issues with kube-dns that are resolved in CoreDNS, either in default configuration or with some customized configurations.
The functional behavior of the default CoreDNS configuration is the same as kube-dns. However,
one difference you need to be aware of is that the published metrics are not the same. In kube-dns,
you get separate metrics for ```dnsmasq``` and ```kubedns``` (skydns). In CoreDNS there is a completely
different set of metrics, since it is all a single process. You can find more details on these
metrics on the CoreDNS [Prometheus plugin](https://coredns.io/plugins/metrics/) page.
The standard CoreDNS Kubernetes configuration is designed to be backwards compatible with the prior
kube-dns behavior. But with some configuration changes, CoreDNS can allow you to modify how the
DNS service discovery works in your cluster. A number of these features are intended to still be
compliant with the [Kubernetes DNS specification](https://github.com/kubernetes/dns/blob/master/docs/specification.md);
they enhance functionality but remain backward compatible. Since CoreDNS is not
*only* made for Kubernetes, but is instead a general-purpose DNS server, there are many things you
can do beyond that specification.
In kube-dns, pod name records are &#34;fake&#34;. That is, any &#34;a-b-c-d.namespace.pod.cluster.local&#34; query will
return the IP address &#34;a.b.c.d&#34;. In some cases, this can weaken the identity guarantees offered by TLS. So,
CoreDNS offers a &#34;pods verified&#34; mode, which will only return the IP address if there is a pod in the
specified namespace with that IP address.
In kube-dns, when using a headless service, you can use an SRV request to get a list of
all endpoints for the service:
```dnstools# host -t srv headless
headless.default.svc.cluster.local has SRV record 10 33 0 6234396237313665.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 10 33 0 6662363165353239.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 10 33 0 6338633437303230.headless.default.svc.cluster.local.
dnstools#
```However, the endpoint DNS names are (for practical purposes) random. In CoreDNS, by default, you get endpoint
DNS names based upon the endpoint IP address:
```dnstools# host -t srv headless
headless.default.svc.cluster.local has SRV record 0 25 443 172-17-0-14.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 172-17-0-18.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 172-17-0-4.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 172-17-0-9.headless.default.svc.cluster.local.
```For some applications, it is desirable to have the pod name for this, rather than the pod IP
address (see for example [kubernetes#47992](https://github.com/kubernetes/kubernetes/issues/47992) and [coredns#1190](https://github.com/coredns/coredns/pull/1190)). To enable this in CoreDNS, you specify the &#34;endpoint_pod_names&#34; option in your Corefile, which results in this:
```dnstools# host -t srv headless
headless.default.svc.cluster.local has SRV record 0 25 443 headless-65bb4c479f-qv84p.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 headless-65bb4c479f-zc8lx.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 headless-65bb4c479f-q7lf2.headless.default.svc.cluster.local.
headless.default.svc.cluster.local has SRV record 0 25 443 headless-65bb4c479f-566rt.headless.default.svc.cluster.local.
```CoreDNS also has a special feature to improve latency in DNS requests for external names. In Kubernetes, the
DNS search path for pods specifies a long list of suffixes. This enables the use of short names when requesting
services in the cluster - for example, &#34;headless&#34; above, rather than &#34;headless.default.svc.cluster.local&#34;. However,
when requesting an external name  - &#34;infoblox.com&#34;, for example - several invalid DNS queries are made by the client,
requiring a roundtrip from the client to kube-dns each time (actually to ```dnsmasq``` and then to ```kubedns```, since [negative caching is disabled](https://github.com/kubernetes/dns/issues/121)):
In CoreDNS, an optional feature called [autopath](https://coredns.io/plugins/autopath) can be enabled that will cause this search path to be followed
*in the server*. That is, CoreDNS will figure out from the source IP address which namespace the client pod is in,
and it will walk this search list until it gets a valid answer. Since the first 3 of these are resolved internally
within CoreDNS itself, it cuts out all of the back and forth between the client and server, reducing latency.
In CoreDNS, you can use standard DNS zone transfer to export the entire DNS record set. This is useful for
debugging your services as well as importing the cluster zone into other DNS servers.
You can also filter by namespaces or a label selector. This can allow you to run specific CoreDNS instances that will only server records that match the filters, exposing only a limited set of your services via DNS.
In addition to the features described above, CoreDNS is easily extended. It is possible to build custom versions
of CoreDNS that include your own features. For example, this ability has been used to extend CoreDNS to do recursive resolution
with the [unbound plugin](https://coredns.io/explugins/unbound), to server records directly from a database with the [pdsql plugin](https://coredns.io/explugins/pdsql), and to allow multiple CoreDNS instances to share a common level 2 cache with the [redisc plugin](https://coredns.io/explugins/redisc).
Many other interesting extensions have been added, which you will find on the [External Plugins](https://coredns.io/explugins/) page of the CoreDNS site. One that is really interesting for Kubernetes and Istio users is the [kubernetai plugin](https://coredns.io/explugins/kubernetai), which allows a single CoreDNS instance to connect to multiple Kubernetes clusters and provide service discovery across all of them.
CoreDNS is an independent project, and as such is developing many features that are not directly
related to Kubernetes. However, a number of these will have applications within Kubernetes. For example,
the upcoming integration with policy engines will allow CoreDNS to make intelligent choices about which endpoint
to return when a headless service is requested. This could be used to route traffic to a local pod, or
to a more responsive pod. Many other features are in development, and of course as an open source project, we welcome you to suggest and contribute your own features!
The features and differences described above are a few examples. There is much more you can do with CoreDNS.
You can find out more on the [CoreDNS Blog](https://coredns.io/blog).
CoreDNS is an incubated [CNCF](https://cncf.io) project.
We&#39;re most active on Slack (and GitHub):
More resources can be found:


	

	


