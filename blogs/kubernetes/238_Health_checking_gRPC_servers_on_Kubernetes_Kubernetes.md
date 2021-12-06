|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/10/01/health-checking-grpc-servers-on-kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-10-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.5849869 &#43;0300 MSK m=&#43;2.174021801  |

# Health checking gRPC servers on Kubernetes | Kubernetes

	
	
	
	
	**Author**: [Ahmet Alp Balkan](https://twitter.com/ahmetb) (Google)
[gRPC](https://grpc.io) is on its way to becoming the lingua franca for
communication between cloud-native microservices. If you are deploying gRPC
applications to Kubernetes today, you may be wondering about the best way to
configure health checks. In this article, we will talk about
[grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe/), a
Kubernetes-native way to health check gRPC apps.
If you&#39;re unfamiliar, Kubernetes [health
checks](/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)
(liveness and readiness probes) is what&#39;s keeping your applications available
while you&#39;re sleeping. They detect unresponsive pods, mark them unhealthy, and
cause these pods to be restarted or rescheduled.
Kubernetes [does not
support](https://github.com/kubernetes/kubernetes/issues/21493) gRPC health
checks natively. This leaves the gRPC developers with the following three
approaches when they deploy to Kubernetes:
[img](/images/blog/2019-09-30-health-checking-grpc/options.png)
Can we do better? Absolutely.
To standardize the &#34;exec probe&#34; approach mentioned above, we need:
Thankfully, gRPC has a [standard health checking
protocol](https://github.com/grpc/grpc/blob/v1.15.0/doc/health-checking.md). It
can be used easily from any language. Generated code and the utilities for
setting the health status are shipped in nearly all language implementations of
gRPC.
If you
[implement](https://github.com/grpc/grpc/blob/v1.15.0/src/proto/grpc/health/v1/health.proto)
this health check protocol in your gRPC apps, you can then use a standard/common
tool to invoke this ```Check()``` method to determine server status.
The next thing you need is the &#34;standard tool&#34;, and it&#39;s the
[strong](https://github.com/grpc-ecosystem/grpc-health-probe/).
[
    
    img
    
    img
](/images/blog/2019-09-30-health-checking-grpc/grpc_health_probe.png)With this tool, you can use the same health check configuration in all your gRPC
applications. This approach requires you to:
In this case, executing &#34;grpc_health_probe&#34; will call your gRPC server over
```localhost```, since they are in the same pod.
**grpc-health-probe** project is still in its early days and it needs your
feedback. It supports a variety of features like communicating with TLS servers
and configurable connection/RPC timeouts.
If you are running a gRPC server on Kubernetes today, try using the gRPC Health
Protocol and try the grpc-health-probe in your deployments, and [give
feedback](https://github.com/grpc-ecosystem/grpc-health-probe/).


	

	


