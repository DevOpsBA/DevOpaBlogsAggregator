|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/10/Helm-Charts-Making-It-Simple-To-Package-And-Deploy-Apps-On-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-10-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.1573027 &#43;0300 MSK m=&#43;3.746346601  |

#  Helm Charts: making it simple to package and deploy common applications on Kubernetes  | Kubernetes

	
	
	
	
	There are thousands of people and companies packaging their applications for deployment on Kubernetes. This usually involves crafting a few different Kubernetes resource definitions that configure the application runtime, as well as defining the mechanism that users and other apps leverage to communicate with the application. There are some very common applications that users regularly look for guidance on deploying, such as databases, CI tools, and content management systems. These types of applications are usually not ones that are developed and iterated on by end users, but rather their configuration is customized to fit a specific use case. Once that application is deployed users can link it to their existing systems or leverage their functionality to solve their pain points.
For best practices on how these applications should be configured, users could look at the many resources available such as: the [examples folder](https://github.com/kubernetes/kubernetes/tree/master/examples) in the Kubernetes repository, the Kubernetes [contrib repository](https://github.com/kubernetes/contrib), the [Helm Charts repository](https://github.com/helm/charts), and the [Bitnami Charts repository](https://github.com/bitnami/charts). While these different locations provided guidance, it was not always formalized or consistent such that users could leverage similar installation procedures across different applications.
So what do you do when there are too many places for things to be found?
[img](https://lh5.googleusercontent.com/l6CowJsfGRoH2wgWHlxtId4Foil2Fcs7AZ0NbOT7jGrXliESRSc6jNH8bdMmfpU-_gDRqy9UDSYCj7WaSKF1ZLK1a7t2qNo5JaIOglozee2SDIPteuOZ6aHzNMyBBJXukBv0zF9x)
[xkcd Standards](https://xkcd.com/927/)
In this case, we’re not creating Yet Another Place for Applications, rather promoting an existing one as the canonical location. As part of the Special Interest Group Apps ([SIG Apps](https://github.com/kubernetes/community/tree/master/sig-apps)) work for the [Kubernetes 1.4 release](https://kubernetes.io/blog/2016/09/kubernetes-1-4-making-it-easy-to-run-on-kuberentes-anywhere/), we began to provide a home for these Kubernetes deployable applications that provides continuous releases of well documented and user friendly packages. These packages are being created as Helm [strong](https://github.com/kubernetes/helm/blob/master/docs/charts.md) and can be installed using the Helm tool. **[Helm](https://github.com/kubernetes/helm)** allows users to easily templatize their Kubernetes manifests and provide a set of configuration parameters that allows users to customize their deployment.
**Helm is the package manager** (analogous to yum and apt) and **Charts are packages** (analogous to debs and rpms). The home for these Charts is the [Kubernetes Charts repository](https://github.com/kubernetes/charts) which provides continuous integration for pull requests, as well as automated releases of Charts in the master branch.
There are two main folders where charts reside. The [stable folder](https://github.com/kubernetes/charts/tree/master/stable) hosts those applications which meet minimum requirements such as proper documentation and inclusion of only Beta or higher Kubernetes resources. The [incubator folder](https://github.com/kubernetes/charts/tree/master/incubator) provides a place for charts to be submitted and iterated on until they’re ready for promotion to stable at which time they will automatically be pushed out to the default repository. For more information on the repository structure and requirements for being in stable, have a look at [this section in the README](https://github.com/kubernetes/charts#repository-structure).
The following applications are now available:
**Example workflow for a Chart developer**
**Example workflow for a Chart user**
```$ helm search  
NAME VERSION DESCRIPTION stable/drupal 0.3.1 One of the most versatile open source content m...stable/jenkins 0.1.0 A Jenkins Helm chart for Kubernetes. stable/mariadb 0.4.0 Chart for MariaDB stable/mysql 0.1.0 Chart for MySQL stable/redmine 0.3.1 A flexible project management web application. stable/wordpress 0.3.0 Web publishing platform for building blogs and ...
``````$ helm install stable/jenkins
``````Notes:



1. Get your &#39;admin&#39; user password by running:

  printf $(printf &#39;\%o&#39; `kubectl get secret --namespace default brawny-frog-jenkins -o jsonpath=&#34;{.data.jenkins-admin-password[*]}&#34;`);echo



2. Get the Jenkins URL to visit by running these commands in the same shell:

\*\*\*\* NOTE: It may take a few minutes for the LoadBalancer IP to be available.                      \*\*\*\*

\*\*\*\*       You can watch the status of by running &#39;kubectl get svc -w brawny-frog-jenkins&#39; \*\*\*\*

  export SERVICE\_IP=$(kubectl get svc --namespace default brawny-frog-jenkins -o jsonpath=&#39;{.status.loadBalancer.ingress[0].ip}&#39;)

  echo http://$SERVICE\_IP:8080/login
```For more information on running Jenkins on Kubernetes, visit [here](https://cloud.google.com/solutions/jenkins-on-container-engine).
**Conclusion**
Now that you’ve seen workflows for both developers and users, we hope that you’ll join us in consolidating the breadth of application deployment knowledge into a more centralized place. Together we can raise the quality bar for both developers and users of Kubernetes applications. We’re always looking for feedback on how we can better our process. Additionally, we’re looking for contributions of new charts or updates to existing ones. Join us in the following places to get engaged:
*--Vic Iglesias, Cloud Solutions Architect, Google*


	

	


