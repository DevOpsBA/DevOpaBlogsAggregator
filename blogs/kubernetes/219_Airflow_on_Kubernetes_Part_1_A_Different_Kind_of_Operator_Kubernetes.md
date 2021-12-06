|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/06/28/airflow-on-kubernetes-part-1-a-different-kind-of-operator/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-06-28 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.7460259 &#43;0300 MSK m=&#43;2.335061801  |

# Airflow on Kubernetes (Part 1): A Different Kind of Operator | Kubernetes

	
	
	
	
	**Author**: Daniel Imberman (Bloomberg LP)
As part of Bloomberg&#39;s [continued commitment to developing the Kubernetes ecosystem](https://www.techatbloomberg.com/blog/bloomberg-awarded-first-cncf-end-user-award-contributions-kubernetes/), we are excited to announce the Kubernetes Airflow Operator; a mechanism for [Apache Airflow](https://airflow.apache.org/), a popular workflow orchestration framework to natively launch arbitrary Kubernetes Pods using the Kubernetes API.
Apache Airflow is one realization of the DevOps philosophy of &#34;Configuration As Code.&#34; Airflow allows users to launch multi-step pipelines using a simple Python object DAG (Directed Acyclic Graph). You can define dependencies, programmatically construct complex workflows, and monitor scheduled jobs in an easy to read UI.

Since its inception, Airflow&#39;s greatest strength has been its flexibility. Airflow offers a wide range of integrations for services ranging from Spark and HBase, to services on various cloud providers. Airflow also offers easy extensibility through its plug-in framework. However, one limitation of the project is that Airflow users are confined to the frameworks and clients that exist on the Airflow worker at the moment of execution. A single organization can have varied Airflow workflows ranging from data science pipelines to application deployments. This difference in use-case creates issues in dependency management as both teams might use vastly different libraries for their workflows.
To address this issue, we&#39;ve utilized Kubernetes to allow users to launch arbitrary Kubernetes pods and configurations. Airflow users can now have full power over their run-time environments, resources, and secrets, basically turning Airflow into an &#34;any job you want&#34; workflow orchestrator.
Before we move any further, we should clarify that an [Operator](https://airflow.apache.org/concepts.html#operators) in Airflow is a task definition. When a user creates a DAG, they would use an operator like the &#34;SparkSubmitOperator&#34; or the &#34;PythonOperator&#34; to submit/monitor a Spark job or a Python function respectively. Airflow comes with built-in operators for frameworks like Apache Spark, BigQuery, Hive, and EMR. It also offers a Plugins entrypoint that allows DevOps engineers to develop their own connectors.
Airflow users are always looking for ways to make deployments and ETL pipelines simpler to manage. Any opportunity to decouple pipeline steps, while increasing monitoring, can reduce future outages and fire-fights. The following is a list of benefits provided by the Airflow Kubernetes Operator:

The Kubernetes Operator uses the [Kubernetes Python Client](https://github.com/kubernetes-client/Python) to generate a request that is processed by the APIServer (1). Kubernetes will then launch your pod with whatever specs you&#39;ve defined (2). Images will be loaded with all the necessary environment variables, secrets and dependencies, enacting a single command. Once the job is launched, the operator only needs to monitor the health of track logs (3). Users will have the choice of gathering logs locally to the scheduler or to any distributed logging service currently in their Kubernetes cluster.
The following DAG is probably the simplest example we could write to show how the Kubernetes Operator works. This DAG creates two pods on Kubernetes: a Linux distro with Python and a base Ubuntu distro without it. The Python pod will run the Python request correctly, while the one without Python will report a failure to the user. If the Operator is working correctly, the ```passing-task``` pod should complete, while the ```failing-task``` pod returns a failure to the Airflow webserver.

While this example only uses basic images, the magic of Docker is that this same DAG will work for any image/command pairing you want. The following is a recommended CI/CD pipeline to run production-ready code on an Airflow DAG.
Use Travis or Jenkins to run unit and integration tests, bribe your favorite team-mate into PR&#39;ing your code, and merge to the master branch to trigger an automated CI build.
[Generate your Docker images and bump release version within your Jenkins build](https://getintodevops.com/blog/building-your-first-docker-image-with-jenkins-2-guide-for-developers).
Finally, update your DAGs to reflect the new release version and you should be ready to go!
Since the Kubernetes Operator is not yet released, we haven&#39;t released an official [helm](https://helm.sh/) chart or operator (however both are currently in progress). However, we are including instructions for a basic deployment below and are actively looking for foolhardy beta testers to try this new feature. To try this system out please follow these steps:
Run ```git clone https://github.com/apache/incubator-airflow.git``` to clone the official Airflow repo.
To run this basic deployment, we are co-opting the integration testing script that we currently use for the Kubernetes Executor (which will be explained in the next article of this series). To launch this deployment, run these three commands:
```sed -ie &#34;s/KubernetesExecutor/LocalExecutor/g&#34; scripts/ci/kubernetes/kube/configmaps.yaml
./scripts/ci/kubernetes/Docker/build.sh
./scripts/ci/kubernetes/kube/deploy.sh
```Before we move on, let&#39;s discuss what these commands are doing:
The Kubernetes Executor is another Airflow feature that allows for dynamic allocation of tasks as idempotent pods. The reason we are switching this to the LocalExecutor is simply to introduce one feature at a time. You are more then welcome to skip this step if you would like to try the Kubernetes Executor, however we will go into more detail in a future article.
This script will tar the Airflow master source code build a Docker container based on the Airflow distribution
Finally, we create a full Airflow deployment on your cluster. This includes Airflow configs, a postgres backend, the webserver &#43; scheduler, and all necessary services between. One thing to note is that the role binding supplied is a cluster-admin, so if you do not have that level of permission on the cluster, you can modify this at scripts/ci/kubernetes/kube/airflow.yaml
Now that your Airflow instance is running let&#39;s take a look at the UI! The UI lives in port 8080 of the Airflow pod, so simply run
```WEB=$(kubectl get pods -o go-template --template &#39;{{range .items}}{{.metadata.name}}{{&#34;\n&#34;}}{{end}}&#39; | grep &#34;airflow&#34; | head -1)
kubectl port-forward $WEB 8080:8080
```Now the Airflow UI will exist on http://localhost:8080. To log in simply enter ```airflow```/```airflow``` and you should have full access to the Airflow web UI.
To modify/add your own DAGs, you can use ```kubectl cp``` to upload local files into the DAG folder of the Airflow scheduler. Airflow will then read the new DAG and automatically upload it to its system. The following command will upload any local file into the correct directory:
```kubectl cp &lt;local file&gt; &lt;namespace&gt;/&lt;pod&gt;:/root/airflow/dags -c scheduler```
While this feature is still in the early stages, we hope to see it released for wide release in the next few months.
This feature is just the beginning of multiple major efforts to improves Apache Airflow integration into Kubernetes. The Kubernetes Operator has been merged into the [1.10 release branch of Airflow](https://github.com/apache/incubator-airflow/tree/v1-10-test) (the executor in experimental mode), along with a fully k8s native scheduler called the Kubernetes Executor (article to come). These features are still in a stage where early adopters/contributers can have a huge influence on the future of these features.
For those interested in joining these efforts, I&#39;d recommend checkint out these steps:
Special thanks to the Apache Airflow and Kubernetes communities, particularly Grant Nicholas, Ben Goldberg, Anirudh Ramanathan, Fokko Dreisprong, and Bolke de Bruin, for your awesome help on these features as well as our future efforts.


	

	


