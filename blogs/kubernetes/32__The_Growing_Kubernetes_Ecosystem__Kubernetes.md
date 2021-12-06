|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/07/The-Growing-Kubernetes-Ecosystem/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-07-24 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.2751402 &#43;0300 MSK m=&#43;4.864190501  |

#  The Growing Kubernetes Ecosystem  | Kubernetes

	
	
	
	
	Over the past year, we’ve seen fantastic momentum in the Kubernetes project, culminating with the release of [Kubernetes v1](https://tectonic.com/) earlier this week. We’ve also witnessed the ecosystem around Kubernetes blossom, and wanted to draw attention to some of the cooler offerings we’ve seen.
| ----- |
|

|
[CloudBees](https://www.hds.com/corporate/press-analyst-center/press-releases/2015/gl150721.html) and the Jenkins community have created a Kubernetes plugin, allowing Jenkins slaves to be built as Docker images and run in Docker hosts managed by Kubernetes, either on the Google Cloud Platform or on a more local Kubernetes instance. These elastic slaves are then brought online as Jenkins schedules jobs for them and destroyed after their builds are complete, ensuring masters have steady access to clean workspaces and minimizing builds’ resource footprint.
|
|

|
[CoreOS](https://www.kismatic.com/) has created launched Tectonic, an opinionated enterprise distribution of Kubernetes, CoreOS and Docker. Tectonic includes a management console for workflows and dashboards, an integrated registry to build and share containers, and additional tools to automate deployment and customize rolling updates. At KuberCon, CoreOS launched Tectonic Preview, giving users easy access to Kubernetes 1.0, 24x7 enterprise ready support, Kubernetes guides and Kubernetes training to help enterprises begin experiencing the power of Kubernetes, CoreOS and Docker.
|
|

|
[Hitachi Data Systems](http://info.meteor.com/blog/meteor-and-a-galaxy-of-containers-with-kubernetes) has announced that Kubernetes now joins the list of solutions validated to run on their enterprise Unified Computing Platform. With this announcement Hitachi has validated Kubernetes and VMware running side-by-side on the UCP platform, providing an enterprise solution for container-based applications and traditional virtualized workloads.
|
|

|
[Kismatic](https://mesosphere.com/training/kubernetes/) is providing enterprise support for pure play open source Kubernetes. They have announced open source and commercially supported Kubernetes plug-ins specifically built for production-grade enterprise environments. Any Kubernetes deployment can now benefit from modular role-based access controls (RBAC), Kerberos for bedrock authentication, LDAP/AD integration, rich auditing and platform-agnostic Linux distro packages.
|
|

|
[Meteor Development Group](https://www.mirantis.com/blog/kubernetes-docker-mirantis-openstack-6-1/), creators of Meteor, a JavaScript App Platform, are using Kubernetes to build [Galaxy](https://www.mirantis.com/blog/kubernetes-docker-mirantis-openstack-6-1/) to run Meteor apps in production. Galaxy will scale from free test apps to production-suitable high-availability hosting.
|
|

|
Mesosphere has incorporated Kubernetes into its Data Center Operating System (DCOS) platform as a first class citizen. Using DCOS, enterprises can deploy Kubernetes across thousands of nodes, both bare-metal and virtualized machines that can run on-premise and in the cloud.  Mesosphere also launched a beta of their [Kubernetes Training Bootcamp](http://www.opencontrail.org/opencontrail-kubernetes-integration/) and will be offering more in the future.
|
|

|
[Mirantis](http://pachyderm.io/) is enabling hybrid cloud applications across OpenStack and other clouds supporting Kubernetes. An OpenStack Murano app package supports full application lifecycle actions such as deploy, create cluster, create pod, add containers to pods, scale up and scale down.
|
|

|
[OpenContrail](http://www.platalytics.com/) is creating a kubernetes-contrail plugin designed to stitch the cluster management capabilities of Kubernetes with the network service automation capabilities of OpenContrail. Given the event-driven abstractions of pods and services inherent in Kubernetes, it is a simple extension to address network service enforcement by leveraging OpenContrail’s Virtual Network policy approach and programmatic API’s.
|
|

|
[Pachyderm](https://github.com/metral/corekube) is a containerized data analytics engine which provides the broad functionality of Hadoop with the ease of use of Docker. Users simply provide containers with their data analysis logic and Pachyderm will distribute that computation over the data. They have just released full deployment on Kubernetes for on premise deployments, and on Google Container Engine, eliminating all the operational overhead of running a cluster yourself.
|
|

|
[Platalytics, Inc](http://www.redhat.com/en/about/blog/welcoming-kubernetes-officially-enterprise-open-source-world). and announced the release of one-touch deploy-anywhere feature for its Spark Application Platform. Based on Kubernetes, Docker, and CoreOS, it allows simple and automated deployment of Apache Hadoop, Spark, and Platalytics platform, with a single click, to all major public clouds, including Google, Amazon, Azure, DigitalOcean, and private on-premise clouds. It also enables hybrid cloud scenarios, where resources on public and private clouds can be mixed.
|
|

|
[Rackspace](https://github.com/metral/corekube) has created Corekube as a simple, quick way to deploy Kubernetes on OpenStack. By using a decoupled infrastructure that is coordinated by etcd, fleet and flannel, it enables users to try Kubernetes and CoreOS without all the fuss of setting things up by hand.
|
|

|
[Red Hat](http://www.redhat.com/en/about/blog/welcoming-kubernetes-officially-enterprise-open-source-world) is a long time proponent of Kubernetes, and a significant contributor to the project. In their own words, “From Red Hat Enterprise Linux 7 and Red Hat Enterprise Linux Atomic Host to OpenShift Enterprise 3 and the forthcoming Red Hat Atomic Enterprise Platform, we are well-suited to bring container innovations into the enterprise, leveraging Kubernetes as the common backbone for orchestration.”
|
|

|
[Redapt](http://www.redapt.com/kubernetes/%20%E2%80%8E) has launching a variety of turnkey, on-premises Kubernetes solutions co-engineered with other partners in the Kubernetes partner ecosystem. These include appliances built to leverage the CoreOS/Tectonic, Mirantis OpenStack, and Mesosphere platforms for management and provisioning. Redapt also offers private, public, and multi-cloud solutions that help customers accelerate their Kubernetes deployments successfully into production.
|
| ----- |
|
|
|
We’ve also seen a community of services partners spring up to assist in adopting Kubernetes and containers:
| ----- |
|

|
[Biarca](http://biarca.io/building-distributed-multi-cloud-applications-using-kubernetes-and-containers/) is using Kubernetes to ease application deployment and scale on demand across available hybrid and multi-cloud clusters through strategically managed policy. A video on their website illustrates how to use Kubernetes to deploy applications in a private cloud infrastructure based on OpenStack and use a public cloud like GCE to address bursting demand for applications.
|
|

|
[Cloud Technology Partners](http://www.cloudtp.com/container-adoption-services/) has developed a Container Services Offering featuring Kubernetes to assist enterprises with container best practices, adoption and implementation. This offering helps organizations understand how containers deliver competitive edge.
|
|

|
[DoIT International](http://doit-intl.com/kubernetes) is offering a Kubernetes Bootcamp which consists of a series of hands-on exercises interleaved with mini-lectures covering hands on topics such as Container Basics, Using Docker, Kubernetes and Google Container Engine.
|
|

|
[OpenCredo](https://www.opencredo.com/2015/04/20/kubernetes/) provides a practical, lab style container and scheduler course in addition to consulting and solution delivery.  The three-day course allows development teams to quickly ramp up and make effective use of containers in real world scenarios, covering containers in general along with Docker and Kubernetes.
|
|

|
[Pythian](http://www.pythian.com/google-kubernetes/) focuses on helping clients design, implement, and manage systems that directly contribute to revenue and business success. They provide small, [dedicated teams of highly trained and experienced data experts](http://www.pythian.com/blog/lessons-learned-kubernetes/) have the deep Kubernetes and container experience necessary to help companies solve Big Data problems with containers.
|
- Martin Buhr, Product Manager at Google


	

	


