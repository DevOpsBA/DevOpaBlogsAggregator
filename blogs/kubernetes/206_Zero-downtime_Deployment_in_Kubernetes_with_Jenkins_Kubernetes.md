|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/04/30/zero-downtime-deployment-kubernetes-jenkins/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-04-30 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.8775618 &#43;0300 MSK m=&#43;2.466598401  |

# Zero-downtime Deployment in Kubernetes with Jenkins | Kubernetes

	
	
	
	
	Ever since we added the [Kubernetes Continuous Deploy](https://aka.ms/azjenkinsk8s) and [Azure Container Service](https://aka.ms/azjenkinsacs) plugins to the Jenkins update center, &#34;How do I create zero-downtime deployments&#34; is one of our most frequently-asked questions. We created a quickstart template on Azure to demonstrate what zero-downtime deployments can look like. Although our example uses Azure, the concept easily applies to all Kubernetes installations.
Kubernetes supports the RollingUpdate strategy to replace old pods with new ones gradually, while continuing to serve clients without incurring downtime. To perform a RollingUpdate deployment:
We&#39;ll use deployment of the official Tomcat image to demonstrate this:
```apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: tomcat-deployment-rolling-update
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: tomcat
        role: rolling-update
    spec:
      containers:
      - name: tomcat-container
        image: tomcat:${TOMCAT_VERSION}
        ports:
        - containerPort: 8080
        readinessProbe:
          httpGet:
            path: /
            port: 8080
  strategy:
    type: RollingUpdate
    rollingUp      maxSurge: 50%
```If the Tomcat running in the current deployments is version 7, we can replace ```${TOMCAT_VERSION}``` with 8 and apply this to the Kubernetes cluster. With the [Kubernetes Continuous Deploy](https://aka.ms/azjenkinsk8s) or the [Azure Container Service](https://aka.ms/azjenkinsacs) plugin, the value can be fetched from an environment variable which eases the deployment process.
Behind the scenes, Kubernetes manages the update like so:

The Rolling Update strategy ensures we always have some Ready backend pods serving client requests, so there&#39;s no service downtime. However, some extra care is required:
*Blue/green deployment quoted from TechTarget*
Container technology offers a stand-alone environment to run the desired service, which makes it super easy to create identical environments as required in the blue/green deployment. The loosely coupled Services - ReplicaSets, and the label/selector-based service routing in Kubernetes make it easy to switch between different backend environments. With these techniques, the blue/green deployments in Kubernetes can be done as follows:
```apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: tomcat-deployment-${TARGET_ROLE}
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: tomcat
        role: ${TARGET_ROLE}
    spec:
      containers:
      - name: tomcat-container
        image: tomcat:${TOMCAT_VERSION}
        ports:
        - containerPort: 8080
        readinessProbe:
          httpGet:
            path: /
            port: 8080
``````kind: Service
apiVersion: v1
metadata:
  name: tomcat-service
  labels:
    app: tomcat
    role: ${TARGET_ROLE}
    env: prod
spec:
  type: LoadBalancer
  selector:
    app: tomcat
    role: ${TARGET_ROLE}
  ports:
    - port: 80
      targetPort: 8080
``````kind: Service
apiVersion: v1
metadata:
  name: tomcat-test-${TARGET_ROLE}
  labels:
    app: tomcat
    role: test-${TARGET_ROLE}
spec:
  type: LoadBalancer
  selector:
    app: tomcat
    role: ${TARGET_ROLE}
  ports:
    - port: 80
      targetPort: 8080
```
As compared to Rolling Update, the blue/green up* The public service is either routed to the old applications, or new applications, but never both at the same time.
Jenkins provides easy-to-setup workflow to automate your deployments. With [Pipeline](https://jenkins.io/doc/book/pipeline/) support, it is flexible to build the zero-downtime deployment workflow, and visualize the deployment steps.
To facilitate the deployment process for Kubernetes resources, we published the [Kubernetes Continuous Deploy](https://aka.ms/azjenkinsk8s) and the [Azure Container Service](https://aka.ms/azjenkinsacs) plugins built based on the [kubernetes-client](https://github.com/fabric8io/kubernetes-client). You can deploy the resource to Azure Kubernetes Service (AKS) or the general Kubernetes clusters without the need of kubectl, and it supports variable substitution in the resource configuration so you can deploy environment-specific resources to the clusters without updating the resource config.
We created a Jenkins Pipeline to demonstrate the blue/green deployment to AKS. The flow is like the following:

```acsDeploy azureCredentialsId: &#39;stored-azure-credentials-id&#39;,
          configFilePaths: &#34;glob/path/to/*/resource-config-*.yml&#34;,
          containerService: &#34;aks-name | AKS&#34;,
          resourceGroupName: &#34;resource-group-name&#34;,
          enableConfigSubstitution: true
```For the Rolling Update strategy, simply deploy the deployment configuration to the Kubernetes cluster, which is a simple, single step.
We built a quickstart template on Azure to demonstrate how we can do the zero-downtime deployment to AKS (Kubernetes) with Jenkins. Go to [Jenkins Blue-Green Deployment on Kubernetes](https://aka.ms/azjenkinsk8sqs) and click the button Deploy to Azure to get the working demo. This template will provision:
```stage(&#39;Confirm&#39;) {
    mail (to: &#39;to@example.com&#39;,
        subject: &#34;Job &#39;${env.JOB_NAME}&#39; (${env.BUILD_NUMBER}) is waiting for input&#34;,
        body: &#34;Please go to ${env.BUILD_URL}.&#34;)
    input &#39;Ready to go?&#39;
}
```Follow the [Steps](https://github.com/Azure/azure-quickstart-templates/tree/master/301-jenkins-aks-zero-downtime-deployment#steps) to setup the resources and you can try it out by start the Jenkins build jobs.


	

	


