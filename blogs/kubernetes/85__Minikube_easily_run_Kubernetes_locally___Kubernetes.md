|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/07/Minikube-Easily-Run-Kubernetes-Locally/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-07-11 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:50:37.6786638 &#43;0300 MSK m=&#43;7.067336001  |

#  Minikube: easily run Kubernetes locally   | Kubernetes

	
	
	
	
	*Editor&#39;s note: This is the first post in a [series of in-depth articles](https://kubernetes.io/blog/2016/07/five-days-of-kubernetes-1-3) on what&#39;s new in Kubernetes 1.3 *
While Kubernetes is one of the best tools for managing containerized applications available today, and has been production-ready for over a year, Kubernetes has been missing a great local development platform.
For the past several months, several of us from the Kubernetes community have been working to fix this in the [Minikube](http://github.com/kubernetes/minikube) repository on GitHub. Our goal is to build an easy-to-use, high-fidelity Kubernetes distribution that can be run locally on Mac, Linux and Windows workstations and laptops with a single command.
Thanks to lots of help from members of the community, we&#39;re proud to announce the official release of Minikube. This release comes with support for [Kubernetes 1.3](https://kubernetes.io/blog/2016/07/kubernetes-1-3-bridging-cloud-native-and-enterprise-workloads/), new commands to make interacting with your local cluster easier and experimental drivers for xhyve (on macOS) and KVM (on Linux).
**Using Minikube**
Minikube ships as a standalone Go binary, so installing it is as simple as downloading Minikube and putting it on your path:
Minikube currently requires that you have VirtualBox installed, which you can download [here](https://www.virtualbox.org/).
_(This is for Mac, for Linux substitute “minikube-darwin-amd64” with “minikube-linux-amd64”)*curl -Lo minikube [https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64](https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64) &amp;&amp; chmod &#43;x minikube &amp;&amp; sudo mv minikube /usr/local/bin/*
To start a Kubernetes cluster in Minikube, use the ```minikube start``` command:
```$ minikube start

Starting local Kubernetes cluster...

Kubernetes is available at https://192.168.99.100:443

Kubectl is now configured to use the cluster
```
At this point, you have a running single-node Kubernetes cluster on your laptop! Minikube also configures ```kubectl``` for you, so you&#39;re also ready to run containers with no changes.
Minikube creates a Host-Only network interface that routes to your node. To interact with running pods or services, you should send traffic over this address. To find out this address, you can use the ```minikube ip``` command:

Minikube also comes with the Kubernetes Dashboard. To open this up in your browser, you can use the built-in ```minikube dashboard``` command:


In general, Minikube supports everything you would expect from a Kubernetes cluster. You can use ```kubectl exec``` to get a bash shell inside a pod in your cluster. You can use the ```kubectl port-forward``` and ```kubectl proxy``` commands to forward traffic from localhost to a pod or the API server.
Since Minikube is running locally instead of on a cloud provider, certain provider-specific features like LoadBalancers and PersistentVolumes will not work out-of-the-box. However, you can use NodePort LoadBalancers and HostPath PersistentVolumes.
**Architecture**
Minikube is built on top of Docker&#39;s [libmachine](https://github.com/docker/machine/tree/master/libmachine), and leverages the driver model to create, manage and interact with locally-run virtual machines.
[RedSpread](https://redspread.com/) was kind enough to donate their [localkube](https://github.com/redspread/localkube) codebase to the Minikube repo, which we use to spin up a single-process Kubernetes cluster inside a VM. Localkube bundles etcd, DNS, the Kubelet and all the Kubernetes master components into a single Go binary, and runs them all via separate goroutines.
**Upcoming Features**
Minikube has been a lot of fun to work on so far, and we&#39;re always looking to improve Minikube to make the Kubernetes development experience better. If you have any ideas for features, don&#39;t hesitate to let us know in the [issue tracker](https://github.com/kubernetes/minikube/issues). 
Here&#39;s a list of some of the things we&#39;re hoping to add to Minikube soon:
**Community**
We&#39;d love to hear feedback on Minikube. To join the community:
Please give Minikube a try, and let us know how it goes!
*--Dan Lorenc, Software Engineer, Google*


	

	


