|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/08/Create-Couchbase-Cluster-Using-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-08-15 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.3454365 &#43;0300 MSK m=&#43;3.934481501  |

#  Create a Couchbase cluster using Kubernetes  | Kubernetes

	
	
	
	
	*Editor’s note: today’s guest post is by Arun Gupta, Vice President Developer Relations at Couchbase, showing how to setup a Couchbase cluster with Kubernetes.*
[Couchbase Server](http://www.couchbase.com/nosql-databases/couchbase-server) is an open source, distributed NoSQL document-oriented database. It exposes a fast key-value store with managed cache for submillisecond data operations, purpose-built indexers for fast queries and a query engine for executing SQL queries. For mobile and Internet of Things (IoT) environments, [Couchbase Lite](http://developer.couchbase.com/mobile) runs native on-device and manages sync to Couchbase Server.
Couchbase Server 4.5 was [recently announced](http://blog.couchbase.com/2016/june/announcing-couchbase-server-4.5), bringing [many new features](http://developer.couchbase.com/documentation/server/4.5/introduction/whats-new.html), including [production certified support for Docker](http://www.couchbase.com/press-releases/couchbase-announces-support-for-docker-containers). Couchbase is supported on a wide variety of orchestration frameworks for Docker containers, such as Kubernetes, Docker Swarm and Mesos, for full details visit [this page](http://couchbase.com/containers).
This blog post will explain how to create a Couchbase cluster using Kubernetes. This setup is tested using Kubernetes 1.3.3, Amazon Web Services, and Couchbase 4.5 Enterprise Edition.
Like all good things, this post is standing on the shoulder of giants. The design pattern used in this blog was defined in a [Friday afternoon hack](https://twitter.com/arungupta/status/703378246432231424) with [@saturnism](https://twitter.com/saturnism). A working version of the configuration files was [contributed](https://twitter.com/arungupta/status/759059647680552962) by [@r_schmiddy](http://twitter.com/r_schmiddy).
**Couchbase Cluster**
A cluster of Couchbase Servers is typically deployed on commodity servers. Couchbase Server has a peer-to-peer topology where all the nodes are equal and communicate to each other on demand. There is no concept of master nodes, slave nodes, config nodes, name nodes, head nodes, etc, and all the software loaded on each node is identical. It allows the nodes to be added or removed without considering their “type”. This model works particularly well with cloud infrastructure in general. For Kubernetes, this means that we can use the exact same container image for all Couchbase nodes.
A typical Couchbase cluster creation process looks like:
In order to automate using Kubernetes, the cluster creation is split into a “master” and “worker” Replication Controller (RC).
The master RC has only one replica and is also published as a Service. This provides a single reference point to start the cluster creation. By default services are visible only from inside the cluster. This service is also exposed as a load balancer. This allows the [Couchbase Web Console](http://developer.couchbase.com/documentation/server/current/admin/ui-intro.html) to be accessible from outside the cluster.
The worker RC use the exact same image as master RC. This keeps the cluster homogenous which allows to scale the cluster easily.

Configuration files used in this blog are available [here](http://github.com/arun-gupta/couchbase-kubernetes/tree/master/cluster). Let’s create the Kubernetes resources to create the Couchbase cluster.
**Create Couchbase “master” Replication Controller**
Couchbase master RC can be created using the following configuration file:
```apiVersion: v1  
kind: ReplicationController  
metadata:  
  name: couchbase-master-rc  
spec:  
  replicas: 1  
  selector:  
    app: couchbase-master-pod  
  template:  
    metadata:  
      labels:  
        app: couchbase-master-pod  
    spec:  
      containers:  
      - name: couchbase-master  
        image: arungupta/couchbase:k8s  
        env:  
          - name: TYPE  
            value: MASTER  
        ports:  
        - containerPort: 8091  
----  
apiVersion: v1  
kind: Service  
metadata:   
  name: couchbase-master-service  
  labels:   
    app: couchbase-master-service  
spec:   
  ports:  
    - port: 8091  
  selector:   
    app: couchbase-master-pod  
  type: LoadBalancer
```This configuration file creates a couchbase-master-rc Replication Controller. This RC has one replica of the pod created using the arungupta/couchbase:k8s image. This image is created using the Dockerfile [here](http://github.com/arun-gupta/couchbase-kubernetes/blob/master/cluster/Dockerfile). This Dockerfile uses a [configuration script](https://github.com/arun-gupta/couchbase-kubernetes/blob/master/cluster/configure-node.sh) to configure the base Couchbase Docker image. First, it uses [Couchbase REST API](http://developer.couchbase.com/documentation/server/current/rest-api/rest-endpoints-all.html) to setup memory quota, setup index, data and query services, security credentials, and loads a sample data bucket. Then, it invokes the appropriate [Couchbase CLI](http://developer.couchbase.com/documentation/server/current/cli/cbcli-intro.html) commands to add the Couchbase node to the cluster or add the node and rebalance the cluster. This is based upon three environment variables:
For this first configuration file, the TYPE environment variable is set to MASTER and so no additional configuration is done on the Couchbase image.
Let’s create and verify the artifacts.
Create Couchbase master RC:
```kubectl create -f cluster-master.yml   
replicationcontroller &#34;couchbase-master-rc&#34; created  
service &#34;couchbase-master-service&#34; created
```List all the services:
```kubectl get svc  
NAME                       CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE  
couchbase-master-service   10.0.57.201                 8091/TCP   30s  
kubernetes                 10.0.0.1      \&lt;none\&gt;        443/TCP    5h
```Output shows that couchbase-master-service is created.
Get all the pods:
```kubectl get po  
NAME                        READY     STATUS    RESTARTS   AGE  
couchbase-master-rc-97mu5   1/1       Running   0          1m
```A pod is created using the Docker image specified in the configuration file.
Check the RC:
```kubectl get rc  
NAME                  DESIRED   CURRENT   AGE  
couchbase-master-rc   1         1         1m
```It shows that the desired and current number of pods in the RC are matching.
Describe the service:
```kubectl describe svc couchbase-master-service  
Name: couchbase-master-service  
Namespace: default  
Labels: app=couchbase-master-service  
Selector: app=couchbase-master-pod  
Type: LoadBalancer  
IP: 10.0.57.201  
LoadBalancer Ingress: a94f1f286590c11e68e100283628cd6c-1110696566.us-west-2.elb.amazonaws.com  
Port: \&lt;unset\&gt; 8091/TCP  
NodePort: \&lt;unset\&gt; 30019/TCP  
Endpoints: 10.244.2.3:8091  
Session Affinity: None  
Events:

  FirstSeen LastSeen Count From SubobjectPath Type Reason Message

  --------- -------- ----- ---- ------------- -------- ------ -------

  2m 2m 1 {service-controller } Normal CreatingLoadBalancer Creating load balancer

  2m 2m 1 {service-controller } Normal CreatedLoadBalancer Created load balancer
```Among other details, the address shown next to LoadBalancer Ingress is relevant for us. This address is used to access the Couchbase Web Console.
Wait for ~3 mins for the load balancer to be ready to receive requests. Couchbase Web Console is accessible at &lt;ip&gt;:8091 and looks like:

The image used in the configuration file is configured with the Administrator username and password password. Enter the credentials to see the console:

Click on Server Nodes to see how many Couchbase nodes are part of the cluster. As expected, it shows only one node:

Click on Data Buckets to see a sample bucket that was created as part of the image:

This shows the travel-sample bucket is created and has 31,591 JSON documents.
**Create Couchbase “worker” Replication Controller**
Now, let’s create a worker replication controller. It can be created using the configuration file:
```apiVersion: v1  
kind: ReplicationController  
metadata:  
  name: couchbase-worker-rc  
spec:  
  replicas: 1  
  selector:  
    app: couchbase-worker-pod  
  template:  
    metadata:  
      labels:  
        app: couchbase-worker-pod  
    spec:  
      containers:  
      - name: couchbase-worker  
        image: arungupta/couchbase:k8s  
        env:  
          - name: TYPE  
            value: &#34;WORKER&#34;  
          - name: COUCHBASE\_MASTER  
            value: &#34;couchbase-master-service&#34;  
          - name: AUTO\_REBALANCE  
            value: &#34;false&#34;  
        ports:  
        - containerPort: 8091
```This RC also creates a single replica of Couchbase using the same arungupta/couchbase:k8s image. The key differences here are:
```kubectl create -f cluster-worker.yml   
replicationcontroller &#34;couchbase-worker-rc&#34; created
```Check the RC:
```kubectl get rc  
NAME                  DESIRED   CURRENT   AGE  
couchbase-master-rc   1         1         6m  
couchbase-worker-rc   1         1         22s
```A new couchbase-worker-rc is created where the desired and the current number of instances are matching.
Get all pods:
```kubectl get po  
NAME                        READY     STATUS    RESTARTS   AGE  
couchbase-master-rc-97mu5   1/1       Running   0          6m  
couchbase-worker-rc-4ik02   1/1       Running   0          46s
```An additional pod is now created. Each pod’s name is prefixed with the corresponding RC’s name. For example, a worker pod is prefixed with couchbase-worker-rc.
Couchbase Web Console gets updated to show that a new Couchbase node is added. This is evident by red circle with the number 1 on the Pending Rebalance tab.

Clicking on the tab shows the IP address of the node that needs to be rebalanced:

**Scale Couchbase cluster**
Now, let’s scale the Couchbase cluster by scaling the replicas for worker RC:
```kubectl scale rc couchbase-worker-rc --replicas=3  
replicationcontroller &#34;couchbase-worker-rc&#34; scaled
```Updated state of RC shows that 3 worker pods have been created:
```kubectl get rc  
NAME                  DESIRED   CURRENT   AGE  
couchbase-master-rc   1         1         8m  
couchbase-worker-rc   3         3         2m
```This can be verified again by getting the list of pods:
```kubectl get po  
NAME                        READY     STATUS    RESTARTS   AGE  
couchbase-master-rc-97mu5   1/1       Running   0          8m  
couchbase-worker-rc-4ik02   1/1       Running   0          2m  
couchbase-worker-rc-jfykx   1/1       Running   0          53s  
couchbase-worker-rc-v8vdw   1/1       Running   0          53s
```Pending Rebalance tab of Couchbase Web Console shows that 3 servers have now been added to the cluster and needs to be rebalanced.

Rebalance Couchbase Cluster
Finally, click on Rebalance button to rebalance the cluster. A message window showing the current state of rebalance is displayed:

Once all the nodes are rebalanced, Couchbase cluster is ready to serve your requests:

In addition to creating a cluster, Couchbase Server supports a range of [high availability and disaster recovery](http://developer.couchbase.com/documentation/server/current/ha-dr/ha-dr-intro.html) (HA/DR) strategies. Most HA/DR strategies rely on a multi-pronged approach of maximizing availability, increasing redundancy within and across data centers, and performing regular backups.
Now that your Couchbase cluster is ready, you can run your first [sample application](http://developer.couchbase.com/documentation/server/current/travel-app/index.html).
For further information check out the Couchbase [Developer Portal](http://developer.couchbase.com/server) and [Forums](https://forums.couchbase.com/), or see questions on [Stack Overflow](http://stackoverflow.com/questions/tagged/couchbase).
*--Arun Gupta, Vice President Developer Relations at Couchbase*


	

	


