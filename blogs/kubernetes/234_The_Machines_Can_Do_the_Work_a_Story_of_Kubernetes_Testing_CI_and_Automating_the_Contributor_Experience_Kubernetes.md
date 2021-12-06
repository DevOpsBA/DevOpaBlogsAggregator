|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/08/29/the-machines-can-do-the-work-a-story-of-kubernetes-testing-ci-and-automating-the-contributor-experience/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-08-29 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.6015235 &#43;0300 MSK m=&#43;2.190558501  |

# The Machines Can Do the Work, a Story of Kubernetes Testing, CI, and Automating the Contributor Experience | Kubernetes

	
	
	
	
	**Author**: Aaron Crickenberger (Google) and Benjamin Elder (Google)
*“Large projects have a lot of less exciting, yet, hard work. We value time spent automating repetitive work more highly than toil. Where that work cannot be automated, it is our culture to recognize and reward all types of contributions. However, heroism is not sustainable.”* - [Kubernetes Community Values](https://git.k8s.io/community/values.md#automation-over-process)
Like many open source projects, Kubernetes is hosted on GitHub. We felt the barrier to participation would be lowest if the project lived where developers already worked, using tools and processes developers already knew. Thus the project embraced the service fully: it was the basis of our workflow, our issue tracker, our documentation, our blog platform, our team structure, and more.
This strategy worked. It worked so well that the project quickly scaled past its contributors’ capacity as humans. What followed was an incredible journey of automation and innovation. We didn’t just need to rebuild our airplane mid-flight without crashing, we needed to convert it into a rocketship and launch into orbit. We needed machines to do the work.
Initially, we focused on the fact that we needed to support the sheer volume of tests mandated by a complex distributed system such as Kubernetes. Real world failure scenarios had to be exercised via end-to-end (e2e) tests to ensure proper functionality. Unfortunately, e2e tests were susceptible to flakes (random failures) and took anywhere from an hour to a day to complete.
Further experience revealed other areas where machines could do the work for us:
As we developed automation to improve our situation, we followed a few guiding principles:
This led us to create [Prow](https://git.k8s.io/test-infra/prow) as the central component for our automation. Prow is sort of like an [If This, Then That](https://ifttt.com/) for GitHub events, with a built-in library of [commands](https://prow.k8s.io/command-help), [plugins](https://prow.k8s.io/plugins), and utilities. We built Prow on top of Kubernetes to free ourselves from worrying about resource management and scheduling, and ensure a more pleasant operational experience.
Prow lets us do things like:
Prow was initially developed by the engineering productivity team building Google Kubernetes Engine, and is actively contributed to by multiple members of Kubernetes SIG Testing. Prow has been adopted by several other open source projects, including Istio, JetStack, Knative and OpenShift. [Getting started with Prow](https://github.com/kubernetes/test-infra/tree/master/prow#getting-started) takes a Kubernetes cluster and ```kubectl apply starter.yaml``` (running pods on a Kubernetes cluster).
Once we had Prow in place, we began to hit other scaling bottlenecks, and so produced additional tooling to support testing at the scale required by Kubernetes, including:
With workflow automation addressed, we turned our attention to project health. We chose to use Google Cloud Storage (GCS) as our source of truth for all test data, allowing us to lean on established infrastructure, and allowed the community to contribute results. We then built a variety of tools to help individuals and the project as a whole make sense of this data, including:
We approached the Cloud Native Computing Foundation (CNCF) to develop DevStats to glean insights from our GitHub events such as:
Today, the Kubernetes project spans over 125 repos across five orgs. There are 31 Special Interests Groups and 10 Working Groups coordinating development within the project. In the last year the project has had [participation from over 13,800 unique developers](https://k8s.devstats.cncf.io/d/13/developer-activity-counts-by-repository-group?orgId=1&amp;var-period_name=Last%20year&amp;var-metric=contributions&amp;var-repogroup_name=All) on GitHub.
On any given weekday our Prow instance [runs over 10,000 CI jobs](http://velodrome.k8s.io/dashboard/db/bigquery-metrics?panelId=10&amp;fullscreen&amp;orgId=1&amp;from=now-6M&amp;to=now); from March 2017 to March 2018 it ran 4.3 million jobs. Most of these jobs involve standing up an entire Kubernetes cluster, and exercising it using real world scenarios. They allow us to ensure all supported releases of Kubernetes work across cloud providers, container engines, and networking plugins. They make sure the latest releases of Kubernetes work with various optional features enabled, upgrade safely, meet performance requirements, and work across architectures.
With today’s [announcement from CNCF](https://www.cncf.io/announcement/2018/08/29/cncf-receives-9-million-cloud-credit-grant-from-google) – noting that Google Cloud has begun transferring ownership and management of the Kubernetes project’s cloud resources to CNCF community contributors, we are excited to embark on another journey. One that allows the project infrastructure to be owned and operated by the community of contributors, following the same open governance model that has worked for the rest of the project. Sound exciting to you? Come talk to us at #sig-testing on kubernetes.slack.com.
Want to find out more? Come check out these resources:


	

	


