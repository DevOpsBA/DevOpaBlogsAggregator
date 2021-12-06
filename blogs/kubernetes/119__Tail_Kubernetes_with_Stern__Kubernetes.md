|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/10/Tail-Kubernetes-With-Stern/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-10-31 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.0898763 &#43;0300 MSK m=&#43;3.678919801  |

#  Tail Kubernetes with Stern  | Kubernetes

	
	
	
	
	*Editor’s note: today’s post is by Antti Kupila, Software Engineer, at Wercker, about building a tool to tail multiple pods and containers on Kubernetes.*
We love Kubernetes here at [Wercker](http://wercker.com/) and build all our infrastructure on top of it. When deploying anything you need to have good visibility to what&#39;s going on and logs are a first view into the inner workings of your application. Good old tail -f has been around for a long time and Kubernetes has this too, built right into [kubectl](/docs/user-guide/kubectl-overview/).
I should say that tail is by no means the tool to use for debugging issues but instead you should feed the logs into a more persistent place, such as [Elasticsearch](https://www.elastic.co/products/elasticsearch). However, there&#39;s still a place for tail where you need to quickly debug something or perhaps you don&#39;t have persistent logging set up yet (such as when developing an app in [Minikube](https://github.com/kubernetes/minikube)).
**Multiple Pods**
Kubernetes has the concept of [Replication Controllers](/docs/user-guide/replication-controller/) which ensure that n pods are running at the same time. This allows rolling updates and redundancy. Considering they&#39;re quite easy to set up there&#39;s really no reason not to do so.
However now there are multiple pods running and they all have a unique id. One issue here is that you&#39;ll need to know the exact pod id (kubectl get pods) but that changes every time a pod is created so you&#39;ll need to do this every time. Another consideration is the fact that Kubernetes load balances the traffic so you won&#39;t know at which pod the request ends up at. If you&#39;re tailing pod A but the traffic ends up at pod B you&#39;ll miss what happened.
Let&#39;s say we have a pod called service with 3 replicas. Here&#39;s what that would look like:
```$ kubectl get pods                         # get pods to find pod ids

$ kubectl log -f service-1786497219-2rbt1  # pod 1

$ kubectl log -f service-1786497219-8kfbp  # pod 2

$ kubectl log -f service-1786497219-lttxd  # pod 3
```**Multiple containers**
We&#39;re heavy users [gRPC](http://www.grpc.io/) for internal services and expose the gRPC endpoints over REST using [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway). Typically we have server and gateway living as two containers in the same pod (same binary that sets the mode by a cli flag). The gateway talks to the server in the same pod and both ports are exposed to Kubernetes. For internal services we can talk directly to the gRPC endpoint while our website communicates using standard REST to the gateway.
This poses a problem though; not only do we now have multiple pods but we also have multiple containers within the pod. When this is the case the built-in logging of kubectl requires you to specify which containers you want logs from.
If we have 3 replicas of a pod and 2 containers in the pod you&#39;ll need 6 kubectl log -f &lt;pod id&gt; &lt;container id&gt;. We work with big monitors but this quickly gets out of hand…
If our service pod has a server and gateway container we&#39;d be looking at something like this:
```$ kubectl get pods                                 # get pods to find pod ids

$ kubectl describe pod service-1786497219-2rbt1    # get containers in pod

$ kubectl log -f service-1786497219-2rbt1 server   # pod 1

$ kubectl log -f service-1786497219-2rbt1 gateway  # pod 1

$ kubectl log -f service-1786497219-8kfbp server   # pod 2

$ kubectl log -f service-1786497219-8kfbp gateway  # pod 2

$ kubectl log -f service-1786497219-lttxd server   # pod 3

$ kubectl log -f service-1786497219-lttxd gateway  # pod 3
```**Stern**
To get around this we built [Stern](https://github.com/wercker/stern). It&#39;s a super simple utility that allows you to specify both the pod id and the container id as regular expressions. Any match will be followed and the output is multiplexed together, prefixed with the pod and container id, and color-coded for human consumption (colors are stripped if piping to a file).
Here&#39;s how the service example would look:
```$ stern service
```This will match any pod containing the word service and listen to all containers within it. If you only want to see traffic to the server container you could do stern --container server service and it&#39;ll stream the logs of all the server containers from the 3 pods.
The output would look something like this:
```$ stern service

&#43; service-1786497219-2rbt1 › server

&#43; service-1786497219-2rbt1 › gateway

&#43; service-1786497219-8kfbp › server

&#43; service-1786497219-8kfbp › gateway

&#43; service-1786497219-lttxd › server

&#43; service-1786497219-lttxd › gateway

&#43; service-1786497219-8kfbp server Log message from server

&#43; service-1786497219-2rbt1 gateway Log message from gateway

&#43; service-1786497219-8kfbp gateway Log message from gateway

&#43; service-1786497219-lttxd gateway Log message from gateway

&#43; service-1786497219-lttxd server Log message from server

&#43; service-1786497219-2rbt1 server Log message from server
```In addition, if a pod is killed and recreated during a deployment Stern will stop listening to the old pod and automatically hook into the new one. There&#39;s no more need to figure out what the id of that newly created pod is.
**Configuration options**
Stern was deliberately designed to be minimal so there&#39;s not much to it. However, there are still a couple configuration options we can highlight here. They&#39;re very similar to the ones built into kubectl so if you&#39;re familiar with that you should feel right at home.
**Examples**
Tail the gateway container running inside of the envvars pod on staging
``` &#43; stern --context staging --container gateway envvars
```Show auth activity from 15min ago with timestamps
```&#43; stern -t --since 15m auth
```Follow the development of some-new-feature in minikube
```&#43; stern --context minikube some-new-feature
```View pods from another namespace
```&#43; stern --namespace kube-system kubernetes-dashboard
```**Get Stern**
Stern is open source and [available on GitHub](https://github.com/wercker/stern), we&#39;d love your contributions or ideas. If you don&#39;t want to build from source you can also download a precompiled binary from [GitHub releases](https://github.com/wercker/stern/releases).
[img](https://4.bp.blogspot.com/-oNscZEvpzVw/WBeWc4cW4zI/AAAAAAAAAyw/71okg07IPHM6dtBOubO_0kxdYxzwoUGOACLcB/s1600/stern-long.gif)


	

	


