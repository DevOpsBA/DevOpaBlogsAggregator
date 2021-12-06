|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/11/Kubernetes-As-Foundation-For-Cloud-Native-Paas/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-11-03 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.2313426 &#43;0300 MSK m=&#43;4.820392601  |

#  Kubernetes as Foundation for Cloud Native PaaS  | Kubernetes

	
	
	
	
	With Kubernetes continuing to gain momentum as a critical tool for building and scaling container based applications, we’ve been thrilled to see a growing number of platform as a service (PaaS) offerings adopt it as a foundation. PaaS developers have been drawn to Kubernetes by its rapid rate of maturation, the soundness of its core architectural concepts, and the strength of its contributor community. The [Kubernetes ecosystem](https://kubernetes.io/blog/2015/07/the-growing-kubernetes-ecosystem) continues to grow, and these PaaS projects are great additions to it.
[img](https://1.bp.blogspot.com/-xX93tnoIlGo/Vjj2fSc_CDI/AAAAAAAAAi0/lvTkT9jyFog/s1600/k8%2Bipaas%2B1.png)
[img](https://1.bp.blogspot.com/-1XZFGRHGb34/Vjj2wUtA6pI/AAAAAAAAAi8/SD-qRhVIiIs/s1600/k8%2Bipaas%2B2.png)
[OpenShift](http://www.openshift.org/) by Red Hat helps organizations accelerate application delivery by enabling development and IT operations teams to be more agile, responsive and efficient. OpenShift Enterprise 3 is the first fully supported, enterprise-ready, web-scale container application platform that natively integrates the Docker container runtime and packaging format, Kubernetes container orchestration and management engine, on a foundation of Red Hat Enterprise Linux 7, all fully supported by Red Hat from the operating system to application runtimes.
[img](https://2.bp.blogspot.com/-t3L1CANyhUs/Vjj28Zpf9WI/AAAAAAAAAjE/Ef-PLLmHGvU/s1600/k8%2Bipaas%2B3.png)
Huawei, a leading global ICT technology solution provider, will offer container as a service (CaaS) built on Kubernetes in the public cloud for customers with Docker based applications. Huawei CaaS services will manage multiple clusters across data centers, and deploy, monitor and scale containers with high availability and high resource utilization for their customers. For example, one of Huawei’s current software products for their telecom customers utilizes tens of thousands of modules and hundreds of instances in virtual machines. By moving to a container based PaaS platform powered by Kubernetes, Huawei is migrating this product into a micro-service based, cloud native architecture. By decoupling the modules, they’re creating a high performance, scalable solution that runs hundreds, even thousands of containers in the system. Decoupling existing heavy modules could have been a painful exercise. However, using several key concepts introduced by Kubernetes, such as pods, services, labels, and proxies, Huawei has been able to re-architect their software with great ease.
Huawei has made Kubernetes the core runtime engine for container based applications/services, and they’ve been building other PaaS components or capabilities around Kubernetes, such as user access management, composite API, Portal and multiple cluster management. Additionally, as part of the migration to the new platform, they’re enhancing their PaaS solution in the areas of advanced scheduling algorithm, multi tenant support and enhanced container network communication to support customer needs.
[img](https://2.bp.blogspot.com/-Ys0Zn4IQzn0/Vjj3JIE0BVI/AAAAAAAAAjM/ktwltzVa1GE/s1600/k8%2Bipaas%2B4.png)
[Gondor](https://gondor.io/)is a PaaS with a focus on application hosting throughout the lifecycle, from development to testing to staging to production. It supports Python, Go, and Node.js applications as well as technologies such as Postgres, Redis and Elasticsearch. The Gondor team recently re-architected Gondor to incorporate Kubernetes, and discussed this in a [blog post.](https://gondor.io/blog/2015/07/21/rebuilding-gondor-kubernetes/)


	

	


