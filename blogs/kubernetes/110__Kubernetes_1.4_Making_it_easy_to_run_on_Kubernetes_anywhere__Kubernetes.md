|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/09/Kubernetes-1-4-Making-It-Easy-To-Run-On-Kuberentes-Anywhere/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-09-26 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.2415198 &#43;0300 MSK m=&#43;3.830564201  |

#  Kubernetes 1.4: Making it easy to run on Kubernetes anywhere  | Kubernetes

	
	
	
	
	Today we’re happy to announce the release of Kubernetes 1.4.
Since the release to general availability just over 15 months ago, Kubernetes has continued to grow and achieve broad adoption across the industry. From brand new startups to large-scale businesses, users have described how big a difference Kubernetes has made in building, deploying and managing distributed applications. However, one of our top user requests has been making Kubernetes itself easier to install and use. We’ve taken that feedback to heart, and 1.4 has several major improvements.
These setup and usability enhancements are the result of concerted, coordinated work across the community - more than 20 contributors from SIG-Cluster-Lifecycle came together to greatly simplify the Kubernetes user experience, covering improvements to installation, startup, certificate generation, discovery, networking, and application deployment.
Additional product highlights in this release include simplified cluster deployment on any cloud, easy installation of stateful apps, and greatly expanded Cluster Federation capabilities, enabling a straightforward deployment across multiple clusters, and multiple clouds.
**What’s new:**
**Cluster creation with two commands -** To get started with Kubernetes a user must provision nodes, install Kubernetes and bootstrap the cluster. A common request from users is to have an easy, portable way to do this on any cloud (public, private, or bare metal).
**Expanded stateful application support -** While cloud-native applications are built to run in containers, many existing applications need additional features to make it easy to adopt containers. Most commonly, these include stateful applications such as batch processing, databases and key-value stores. In Kubernetes 1.4, we have introduced a number of features simplifying the deployment of such applications, including: 
**Cluster federation API additions -** One of the most requested capabilities from our global customers has been the ability to build applications with clusters that span regions and clouds. 
**Container security support -** Administrators of multi-tenant clusters require the ability to provide varying sets of permissions among tenants, infrastructure components, and end users of the system.
**Infrastructure enhancements - ** We continue adding to the scheduler, storage and client capabilities in Kubernetes based on user and ecosystem needs.
**Kubernetes Dashboard UI -** lastly, a great looking Kubernetes [Dashboard UI](https://github.com/kubernetes/dashboard#kubernetes-dashboard) with 90% CLI parity for at-a-glance management.
For a complete list of updates see the [release notes](https://github.com/kubernetes/kubernetes/pull/33410) on GitHub. Apart from features the most impressive aspect of Kubernetes development is the community of contributors. This is particularly true of the 1.4 release, the full breadth of which will unfold in upcoming weeks.
**Availability**
Kubernetes 1.4 is available for download at [get.k8s.io](http://get.k8s.io/) and via the open source repository hosted on [GitHub](http://github.com/kubernetes/kubernetes). To get started with Kubernetes try the [Hello World app](/docs/hellonode/).
To get involved with the project, join the [weekly community meeting](https://groups.google.com/forum/#!forum/kubernetes-community-video-chat) or start contributing to the project here (marked help). 
**Users and Case Studies**
Over the past fifteen months since the Kubernetes 1.0 GA release, the [adoption and enthusiasm](http://kubernetes.io/case-studies/) for this project has surpassed everyone&#39;s imagination. Kubernetes runs in production at hundreds of organization and thousands more are in development. Here are a few unique highlights of companies running Kubernetes: 
We’re very grateful to our community of over 900 contributors who contributed more than 5,000 commits to make this release possible. To get a closer look on how the community is using Kubernetes, join us at the user conference [KubeCon](http://events.linuxfoundation.org/events/kubecon) to hear directly from users and contributors.
**Connect**
Thank you for your support! 
*-- Aparna Sinha, Product Manager, Google*


	

	


