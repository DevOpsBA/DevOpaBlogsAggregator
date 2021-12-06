|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/09/Creating-Postgresql-Cluster-Using-Helm/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-09-09 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.2452576 &#43;0300 MSK m=&#43;3.834302001  |

#  Creating a PostgreSQL Cluster using Helm  | Kubernetes

	
	
	
	
	*Editor’s note: Today’s guest post is by Jeff McCormick, a developer at Crunchy Data, showing how to deploy a PostgreSQL cluster using Helm, a Kubernetes package manager.*
[Crunchy Data](http://www.crunchydata.com/) supplies a set of open source PostgreSQL and PostgreSQL related containers. The Crunchy PostgreSQL Container Suite includes containers that deploy, monitor, and administer the open source PostgreSQL database, for more details view this GitHub [repository](https://github.com/crunchydata/crunchy-containers).
In this post we’ll show you how to deploy a PostgreSQL cluster using [Helm](https://github.com/kubernetes/helm), a Kubernetes package manager. For reference, the Crunchy Helm Chart examples used within this post are located [here](https://github.com/CrunchyData/crunchy-containers/tree/master/examples/kubehelm/crunchy-postgres), and the pre-built containers can be found on DockerHub at [this location](https://hub.docker.com/u/crunchydata/dashboard/).
This example will create the following in your Kubernetes cluster:

This example creates a simple Postgres streaming replication deployment with a master (read-write), and a single asynchronous replica (read-only). You can scale up the number of replicas dynamically.
**Contents**
The example is made up of various Chart files as follows:
**Installation**
[Install Helm](https://github.com/kubernetes/helm#install) according to their GitHub documentation and then install the examples as follows:
```helm init

cd crunchy-containers/examples/kubehelm

helm install ./crunchy-postgres
```**Testing**
After installing the Helm chart, you will see the following services:
```kubectl get services  
NAME              CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE  
crunchy-master    10.0.0.171   \&lt;none\&gt;        5432/TCP   1h  
crunchy-replica   10.0.0.31    \&lt;none\&gt;        5432/TCP   1h  
kubernetes        10.0.0.1     \&lt;none\&gt;        443/TCP    1h
```It takes about a minute for the replica to begin replicating with the master. To test out replication, see if replication is underway with this command, enter password for the password when prompted:
```psql -h crunchy-master -U postgres postgres -c &#39;table pg\_stat\_replication&#39;
```If you see a line returned from that query it means the master is replicating to the slave. Try creating some data on the master:
```psql -h crunchy-master -U postgres postgres -c &#39;create table foo (id int)&#39;

psql -h crunchy-master -U postgres postgres -c &#39;insert into foo values (1)&#39;
```Then verify that the data is replicated to the slave:
```psql -h crunchy-replica -U postgres postgres -c &#39;table foo&#39;
```You can scale up the number of read-only replicas by running the following kubernetes command:
```kubectl scale rc crunchy-replica --replicas=2
```It takes 60 seconds for the replica to start and begin replicating from the master.
The Kubernetes Helm and Charts projects provide a streamlined way to package up complex applications and deploy them on a Kubernetes cluster.  Deploying PostgreSQL clusters can sometimes prove challenging, but the task is greatly simplified using Helm and Charts.
*--Jeff McCormick, Developer, Crunchy Data*


	

	


