|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/10/Globally-Distributed-Services-Kubernetes-Cluster-Federation/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-10-14 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.1385266 &#43;0300 MSK m=&#43;3.727570401  |

#  Building Globally Distributed Services using Kubernetes Cluster Federation  | Kubernetes

	
	
	
	
	*Editor&#39;s note: Today’s post is by Allan Naim, Product Manager, and Quinton Hoole, Staff Engineer at Google, showing how to deploy a multi-homed service behind a global load balancer and have requests sent to the closest cluster.*
In Kubernetes 1.3, we announced Kubernetes Cluster Federation and introduced the concept of Cross Cluster Service Discovery, enabling developers to deploy a service that was sharded across a federation of clusters spanning different zones, regions or cloud providers. This enables developers to achieve higher availability for their applications, without sacrificing quality of service, as detailed in our [previous](https://kubernetes.io/blog/2016/07/cross-cluster-services) blog post.
In the latest release, [Kubernetes 1.4](https://kubernetes.io/blog/2016/09/kubernetes-1-4-making-it-easy-to-run-on-kuberentes-anywhere/), we&#39;ve extended Cluster Federation to support Replica Sets, Secrets, Namespaces and Ingress objects. This means that you no longer need to deploy and manage these objects individually in each of your federated clusters. Just create them once in the federation, and have its built-in controllers automatically handle that for you.
[strong](/docs/user-guide/federation/replicasets/) leverage the same configuration as non-federated Kubernetes Replica Sets and automatically distribute Pods across one or more federated clusters. By default, replicas are evenly distributed across all clusters, but for cases where that is not the desired behavior, we&#39;ve introduced Replica Set preferences, which allow replicas to be distributed across only some clusters, or in non-equal proportions ([define annotations](https://github.com/kubernetes/kubernetes/blob/master/federation/apis/federation/types.go#L114)).
Starting with Google Cloud Platform (GCP), we’ve introduced [strong](/docs/user-guide/federation/federated-ingress/) as a Kubernetes 1.4 alpha feature which enables external clients point to a single IP address and have requests sent to the closest cluster with usable capacity in any region, zone of the Federation.
[strong](/docs/user-guide/federation/secrets/) automatically create and manage secrets across all clusters in a Federation, automatically ensuring that these are kept globally consistent and up-to-date, even if some clusters are offline when the original updates are applied.
[strong](/docs/user-guide/federation/namespaces/) are similar to the traditional [Kubernetes Namespaces](/docs/user-guide/namespaces/) providing the same functionality. Creating them in the Federation control plane ensures that they are synchronized across all the clusters in Federation.
[strong](/docs/user-guide/federation/events/) are similar to the traditional Kubernetes Events providing the same functionality. Federation Events are stored only in Federation control plane and are not passed on to the underlying kubernetes clusters.
Let’s walk through how all this stuff works. We’re going to provision 3 clusters per region, spanning 3 continents (Europe, North America and Asia).
[img](https://2.bp.blogspot.com/-Gj83DdcKqTI/WAE8pwAEZYI/AAAAAAAAAwI/9dbyBFipvDIGkPQWRB1dRxNwkrvzlcYMwCLcB/s1600/k8s%2Bfed%2Bmap.png)
The next step is to federate these clusters. Kelsey Hightower developed a [tutorial](https://github.com/kelseyhightower/kubernetes-cluster-federation) for setting up a Kubernetes Cluster Federation. Follow the tutorial to configure a Cluster Federation with clusters in 3 zones in each of the 3 GCP regions, us-central1, europe-west1 and asia-east1. For the purpose of this blog post, we’ll provision the Federation Control Plane in the us-central1-b zone. Note that more highly available, multi-cluster deployments are also available, but not used here in the interests of simplicity.
The rest of the blog post assumes that you have a running Kubernetes Cluster Federation provisioned.
Let’s verify that we have 9 clusters in 3 regions running.
```$ kubectl --context=federation-cluster get clusters


NAME              STATUS    AGE  
gce-asia-east1-a     Ready     17m  
gce-asia-east1-b     Ready     15m  
gce-asia-east1-c     Ready     10m  
gce-europe-west1-b   Ready     7m  
gce-europe-west1-c   Ready     7m  
gce-europe-west1-d   Ready     4m  
gce-us-central1-a    Ready     1m  
gce-us-central1-b    Ready     53s  
gce-us-central1-c    Ready     39s
```In our example, we’ll be deploying the service and ingress object using the federated control plane. The [ConfigMap](/docs/user-guide/configmap/) object isn’t currently supported by Federation, so we’ll be deploying it manually in each of the underlying Federation clusters. Our cluster deployment will look as follows:
We’re going to deploy a Service that is sharded across our 9 clusters. The backend deployment will consist of a Pod with 2 containers:
Let’s start by creating a federated service object in the federation-cluster context.
```$ kubectl --context=federation-cluster create -f services/nginx.yaml
```It will take a few minutes for the service to propagate across the 9 clusters.
```$ kubectl --context=federation-cluster describe services nginx


Name:                   nginx  
Namespace:              default  
Labels:                 app=nginx  
Selector:               app=nginx  
Type:                   LoadBalancer  
IP:  
LoadBalancer Ingress:   108.59.xx.xxx, 104.199.xxx.xxx, ...  
Port:                   http    80/TCP

NodePort:               http    30061/TCP  
Endpoints:              &lt;none&gt;  
Session Affinity:       None
```Let’s now create a Federated Ingress. Federated Ingresses are created in much that same way as traditional [Kubernetes Ingresses](/docs/user-guide/ingress/): by making an API call which specifies the desired properties of your logical ingress point. In the case of Federated Ingress, this API call is directed to the Federation API endpoint, rather than a Kubernetes cluster API endpoint. The API for Federated Ingress is 100% compatible with the API for traditional Kubernetes Services.
```$ cat ingress/ingress.yaml   

apiVersion: extensions/v1beta1  
kind: Ingress  
metadata:  
  name: nginx  
spec:  
  backend:  
    serviceName: nginx  
    servicePort: 80
``````$ kubectl --context=federation-cluster create -f ingress/ingress.yaml   
ingress &#34;nginx&#34; created
```Once created, the Federated Ingress controller automatically:
```$ for c in $(kubectl config view -o jsonpath=&#39;{.contexts[*].name}&#39;); do kubectl --context=$c get ingress; done  

NAME      HOSTS     ADDRESS   PORTS     AGE  
nginx     \*                   80        1h  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        40m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        1h  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        26m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        1h  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        25m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        38m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        3m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        57m  
NAME      HOSTS     ADDRESS          PORTS     AGE  
nginx     \*         130.211.40.xxx   80        56m
```Note that in the case of Google Cloud Platform, the logical L7 load balancer is not a single physical device (which would present both a single point of failure, and a single global network routing choke point), but rather a [truly global, highly available load balancing managed service](https://cloud.google.com/load-balancing/), globally reachable via a single, static IP address.
Clients inside your federated Kubernetes clusters (i.e. Pods) will be automatically routed to the cluster-local shard of the Federated Service backing the Ingress in their cluster if it exists and is healthy, or the closest healthy shard in a different cluster if it does not. Note that this involves a network trip to the HTTP(S) load balancer, which resides outside your local Kubernetes cluster but inside the same GCP region.
The next step is to schedule the service backends. Let’s first create the ConfigMap in each cluster in the Federation.
We do this by submitting the ConfigMap to each cluster in the Federation.
```$ for c in $(kubectl config view -o jsonpath=&#39;{.contexts[\*].name}&#39;); do kubectl --context=$c create -f configmaps/zonefetch.yaml; done
```Let’s have a quick peek at our Replica Set:
```$ cat replicasets/nginx-rs.yaml


apiVersion: extensions/v1beta1  
kind: ReplicaSet  
metadata:  
  name: nginx  
  labels:  
    app: nginx  
    type: demo  
spec:  
  replicas: 9  
  template:  
    metadata:  
      labels:  
        app: nginx  
    spec:  
      containers:  
      - image: nginx  
        name: frontend  
        ports:  
          - containerPort: 80  
        volumeMounts:  
        - name: html-dir  
          mountPath: /usr/share/nginx/html  
      - image: busybox  
        name: zone-fetcher  
        command:  
          - &#34;/bin/sh&#34;  
          - &#34;-c&#34;  
          - &#34;/zonefetch/zonefetch.sh&#34;  
        volumeMounts:  
        - name: zone-fetch  
          mountPath: /zonefetch  
        - name: html-dir  
          mountPath: /usr/share/nginx/html  
      volumes:  
        - name: zone-fetch  
          configMap:  
            defaultMode: 0777  
            name: zone-fetch  
        - name: html-dir  
          emptyDir:  
            medium: &#34;&#34;
```The Replica Set consists of 9 replicas, spread evenly across 9 clusters within the Cluster Federation. Annotations can also be used to control which clusters Pods are scheduled to. This is accomplished by adding annotations to the Replica Set spec, as follows:
```apiVersion: extensions/v1beta1  
kind: ReplicaSet  
metadata:  
  name: nginx-us  
  annotations:  
    federation.kubernetes.io/replica-set-preferences: ```  
        {  
            &#34;rebalance&#34;: true,  
            &#34;clusters&#34;: {  
                &#34;gce-us-central1-a&#34;: {  
                    &#34;minReplicas&#34;: 2,  
                    &#34;maxReplicas&#34;: 4,  
                    &#34;weight&#34;: 1  
                },  
                &#34;gce-us-central10b&#34;: {  
                    &#34;minReplicas&#34;: 2,  
                    &#34;maxReplicas&#34;: 4,  
                    &#34;weight&#34;: 1  
                }  
            }  
        }
```For the purpose of our demo, we’ll keep things simple and spread our Pods evenly across the Cluster Federation.
Let’s create the federated Replica Set:
```$ kubectl --context=federation-cluster create -f replicasets/nginx-rs.yaml
```Verify the Replica Sets and Pods were created in each cluster:
```$ for c in $(kubectl config view -o jsonpath=&#39;{.contexts[\*].name}&#39;); do kubectl --context=$c get rs; done  

NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         42s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         14m  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         45s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         46s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         47s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         48s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         49s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         49s  
NAME      DESIRED   CURRENT   READY     AGE  
nginx     1         1         1         49s


$ for c in $(kubectl config view -o jsonpath=&#39;{.contexts[\*].name}&#39;); do kubectl --context=$c get po; done  

NAME          READY     STATUS    RESTARTS   AGE  
nginx-ph8zx   2/2       Running   0          25s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-sbi5b   2/2       Running   0          27s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-pf2dr   2/2       Running   0          28s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-imymt   2/2       Running   0          30s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-9cd5m   2/2       Running   0          31s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-vxlx4   2/2       Running   0          33s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-itagl   2/2       Running   0          33s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-u7uyn   2/2       Running   0          33s  
NAME          READY     STATUS    RESTARTS   AGE  
nginx-i0jh6   2/2       Running   0          34s
```Below is an illustration of how the nginx service and associated ingress deployed. To summarize, we have a global VIP (130.211.23.176) exposed using a Global L7 load balancer that forwards requests to the closest cluster with available capacity.
[img](https://1.bp.blogspot.com/-vDz5dEG_-yI/WAE81YPVlYI/AAAAAAAAAwM/jvt46qwIViQbsbftCqFenUocGfssuLbjwCLcB/s1600/Copy%2Bof%2BFederation%2BBlog%2BDrawing%2B%25281%2529.png)
To test this out, we’re going to spin up 2 Google Cloud Engine (GCE) instances, one in us-west1-b and the other in asia-east1-a. All client requests are automatically routed, via the shortest network path, to a healthy Pod in the closest cluster to the origin of the request. So for example, HTTP(S) requests from Asia will be routed directly to the closest cluster in Asia that has available capacity. If there are no such clusters in Asia, the request will be routed to the next closest cluster (in this case the U.S.). This works irrespective of whether the requests originate from a GCE instance or anywhere else on the internet. We only use a GCE instance for simplicity in the demo.

We can SSH directly into the VMs using the Cloud Console or by issuing a gcloud SSH command.
```$ gcloud compute ssh test-instance-asia --zone asia-east1-a

-----

user@test-instance-asia:~$ curl 130.211.40.186  
&lt;!DOCTYPE html&gt;  
&lt;html&gt;  
&lt;head&gt;  
&lt;title&gt;Welcome to the global site!&lt;/title&gt;  
&lt;/head&gt;  
&lt;body&gt;  
&lt;h1&gt;Welcome to the global site! You are being served from asia-east1-b&lt;/h1&gt;  
&lt;p&gt;Congratulations!&lt;/p&gt;


user@test-instance-asia:~$ exit

----


$ gcloud compute ssh test-instance-us --zone us-west1-b

----

user@test-instance-us:~$ curl 130.211.40.186  
&lt;!DOCTYPE html&gt;  
&lt;html&gt;  
&lt;head&gt;  
&lt;title&gt;Welcome to the global site!&lt;/title&gt;  
&lt;/head&gt;  
&lt;body&gt;  
&lt;h1&gt;Welcome to the global site! You are being served from us-central1-b&lt;/h1&gt;  
&lt;p&gt;Congratulations!&lt;/p&gt;


----
```Federations of Kubernetes Clusters can include clusters running in different cloud providers (e.g. GCP, AWS), and on-premises (e.g. on OpenStack). However, in Kubernetes 1.4, Federated Ingress is only supported across Google Cloud Platform clusters. In future versions we intend to support hybrid cloud Ingress-based deployments.
To summarize, we walked through leveraging the Kubernetes 1.4 Federated Ingress alpha feature to deploy a multi-homed service behind a global load balancer. External clients point to a single IP address and are sent to the closest cluster with usable capacity in any region, zone of the Federation, providing higher levels of availability without sacrificing latency or ease of operation.
We&#39;d love to hear feedback on Kubernetes Cross Cluster Services. To join the community:


	

	


