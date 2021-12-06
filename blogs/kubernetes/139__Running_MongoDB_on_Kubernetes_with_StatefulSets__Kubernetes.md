|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/01/Running-Mongodb-On-Kubernetes-With-Statefulsets/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-01-30 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:50:36.5003456 &#43;0300 MSK m=&#43;5.889011101  |

#  Running MongoDB on Kubernetes with StatefulSets  | Kubernetes

	
	
	
	
	*Editor&#39;s note: Today’s post is by Sandeep Dinesh, Developer Advocate, Google Cloud Platform, showing how to run a database in a container.*
Conventional wisdom says you can’t run a database in a container. “Containers are stateless!” they say, and “databases are pointless without state!”
Of course, this is not true at all. At Google, everything runs in a container, including databases. You just need the right tools. [Kubernetes 1.5](https://kubernetes.io/blog/2016/12/kubernetes-1-5-supporting-production-workloads/) includes the new [StatefulSet](/docs/concepts/abstractions/controllers/statefulsets/) API object (in previous versions, StatefulSet was known as PetSet). With StatefulSets, Kubernetes makes it much easier to run stateful workloads such as databases.
If you’ve followed my previous posts, you know how to create a [MEAN Stack app with Docker](http://blog.sandeepdinesh.com/2015/07/running-mean-web-application-in-docker.html), then [migrate it to Kubernetes](https://medium.com/google-cloud/running-a-mean-stack-on-google-cloud-platform-with-kubernetes-149ca81c2b5d) to provide easier management and reliability, and [create a MongoDB replica set](https://medium.com/google-cloud/mongodb-replica-sets-with-kubernetes-d96606bd9474) to provide redundancy and high availability.
While the replica set in my previous blog post worked, there were some annoying steps that you needed to follow. You had to manually create a disk, a ReplicationController, and a service for each replica. Scaling the set up and down meant managing all of these resources manually, which is an opportunity for error, and would put your stateful application at risk In the previous example, we created a Makefile to ease the management of these resources, but it would have been great if Kubernetes could just take care of all of this for us.
With StatefulSets, these headaches finally go away. You can create and manage your MongoDB replica set natively in Kubernetes, without the need for scripts and Makefiles. Let’s take a look how.
*Note: StatefulSets are currently a beta resource. The [sidecar container](https://github.com/cvallance/mongo-k8s-sidecar) used for auto-configuration is also unsupported.*
**Prerequisites and Setup**
Before we get started, you’ll need a Kubernetes 1.5&#43; and the [Kubernetes command line tool](/docs/user-guide/prereqs/). If you want to follow along with this tutorial and use Google Cloud Platform, you also need the [Google Cloud SDK](http://cloud.google.com/sdk).
Once you have a [Google Cloud project created](https://console.cloud.google.com/projectcreate) and have your Google Cloud SDK setup (hint: gcloud init), we can create our cluster.
To create a Kubernetes 1.5 cluster, run the following command:
```gcloud container clusters create &#34;test-cluster&#34;
```This will make a three node Kubernetes cluster. Feel free to [customize the command](https://cloud.google.com/sdk/gcloud/reference/container/clusters/create) as you see fit.
Then, authenticate into the cluster:
```gcloud container clusters get-credentials test-cluster
```**Setting up the MongoDB replica set**
To set up the MongoDB replica set, you need three things: A [StorageClass](/docs/user-guide/persistent-volumes/#storageclasses), a [Headless Service](/docs/user-guide/services/#headless-services), and a [StatefulSet](/docs/concepts/abstractions/controllers/statefulsets/).
I’ve created the configuration files for these already, and you can clone the example from GitHub:
```git clone https://github.com/thesandlord/mongo-k8s-sidecar.git

cd /mongo-k8s-sidecar/example/StatefulSet/
```To create the MongoDB replica set, run these two commands:
```kubectl apply -f googlecloud\_ssd.yaml

kubectl apply -f mongo-statefulset.yaml
```That&#39;s it! With these two commands, you have launched all the components required to run an highly available and redundant MongoDB replica set.
At an high level, it looks something like this:

Let’s examine each piece in more detail.
**StorageClass**
The storage class tells Kubernetes what kind of storage to use for the database nodes. You can set up many different types of StorageClasses in a ton of different environments. For example, if you run Kubernetes in your own datacenter, you can use [GlusterFS](https://www.gluster.org/). On GCP, your [storage choices](https://cloud.google.com/compute/docs/disks/) are SSDs and hard disks. There are currently drivers for [AWS](/docs/user-guide/persistent-volumes/#aws), [Azure](/docs/user-guide/persistent-volumes/#azure-disk), [Google Cloud](/docs/user-guide/persistent-volumes/#gce), [GlusterFS](/docs/user-guide/persistent-volumes/#glusterfs), [OpenStack Cinder](/docs/user-guide/persistent-volumes/#openstack-cinder), [vSphere](/docs/user-guide/persistent-volumes/#vsphere), [Ceph RBD](/docs/user-guide/persistent-volumes/#ceph-rbd), and [Quobyte](/docs/user-guide/persistent-volumes/#quobyte).
The configuration for the StorageClass looks like this:
This configuration creates a new StorageClass called “fast” that is backed by SSD volumes. The StatefulSet can now request a volume, and the StorageClass will automatically create it!
Deploy this StorageClass:
```kubectl apply -f googlecloud\_ssd.yaml
```**Headless Service**
Now you have created the Storage Class, you need to make a Headless Service. These are just like normal Kubernetes Services, except they don’t do any load balancing for you. When combined with StatefulSets, they can give you unique DNS addresses that let you directly access the pods! This is perfect for creating MongoDB replica sets, because our app needs to connect to all of the MongoDB nodes individually.
The configuration for the Headless Service looks like this:
You can tell this is a Headless Service because the clusterIP is set to “None.” Other than that, it looks exactly the same as any normal Kubernetes Service.
**StatefulSet**
The pièce de résistance. The StatefulSet actually runs MongoDB and orchestrates everything together. StatefulSets differ from Kubernetes [ReplicaSets](/docs/user-guide/replicasets/) (not to be confused with MongoDB replica sets!) in certain ways that makes them more suited for stateful applications. Unlike Kubernetes ReplicaSets, pods created under a StatefulSet have a few unique attributes. The name of the pod is not random, instead each pod gets an ordinal name. Combined with the Headless Service, this allows pods to have stable identification. In addition, pods are created one at a time instead of all at once, which can help when bootstrapping a stateful system. You can read more about StatefulSets in the [documentation](/docs/concepts/abstractions/controllers/statefulsets/).
Just like before, [this “sidecar” container](https://github.com/cvallance/mongo-k8s-sidecar) will configure the MongoDB replica set automatically. A “sidecar” is a helper container which helps the main container do its work.
The configuration for the StatefulSet looks like this:
It’s a little long, but fairly straightforward.
The first second describes the StatefulSet object. Then, we move into the Metadata section, where you can specify labels and the number of replicas.
Next comes the pod spec. The terminationGracePeriodSeconds is used to gracefully shutdown the pod when you scale down the number of replicas, which is important for databases! Then the configurations for the two containers is shown. The first one runs MongoDB with command line flags that configure the replica set name. It also mounts the persistent storage volume to /data/db, the location where MongoDB saves its data. The second container runs the sidecar.
Finally, there is the volumeClaimTemplates. This is what talks to the StorageClass we created before to provision the volume. It will provision a 100 GB disk for each MongoDB replica.
**Using the MongoDB replica set**
At this point, you should have three pods created in your cluster. These correspond to the three nodes in your MongoDB replica set. You can see them with this command:
```kubectl get pods

NAME   READY STATUS RESTARTS AGE
mongo-0 2/2  Running 0     3m
mongo-1 2/2  Running 0     3m
mongo-2 2/2  Running 0     3m
```Each pod in a StatefulSet backed by a Headless Service will have a stable DNS name. The template follows this format: &lt;pod-name&gt;.&lt;service-name&gt;
This means the DNS names for the MongoDB replica set are:
```mongo-0.mongo
mongo-1.mongo
mongo-2.mongo
```You can use these names directly in the [connection string URI](http://docs.mongodb.com/manual/reference/connection-string) of your app.
In this case, the connection string URI would be:
```mongodb://mongo-0.mongo,mongo-1.mongo,mongo-2.mongo:27017/dbname\_?
```That’s it!
**Scaling the MongoDB replica set**
A huge advantage of StatefulSets is that you can scale them just like Kubernetes ReplicaSets. If you want 5 MongoDB Nodes instead of 3, just run the scale command:
```kubectl scale --replicas=5 statefulset mongo
```The sidecar container will automatically configure the new MongoDB nodes to join the replica set.
Include the two new nodes (mongo-3.mongo &amp; mongo-4.mongo) in your connection string URI and you are good to go. Too easy!
**Cleaning Up**
To clean up the deployed resources, delete the StatefulSet, Headless Service, and the provisioned volumes.
Delete the StatefulSet:
```kubectl delete statefulset mongo
```Delete the Service:
```kubectl delete svc mongo
```Delete the Volumes:
```kubectl delete pvc -l role=mongo
```Finally, you can delete the test cluster:
```gcloud container clusters delete &#34;test-cluster&#34;
```Happy Hacking!
For more cool Kubernetes and Container blog posts, follow me on [Twitter](https://twitter.com/sandeepdinesh) and [Medium](https://medium.com/@SandeepDinesh).
*--Sandeep Dinesh, Developer Advocate, Google Cloud Platform.*


	

	


