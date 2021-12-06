|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/08/Kompose-Helps-Developers-Move-Docker/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-08-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.4697261 &#43;0300 MSK m=&#43;3.058766101  |

#  Kompose Helps Developers Move Docker Compose Files to Kubernetes  | Kubernetes

	
	
	
	
	*Editor&#39;s note: today&#39;s post is by Charlie Drage, Software Engineer at Red Hat giving an update about the Kubernetes project Kompose.*
I&#39;m pleased to announce that [Kompose](https://github.com/kubernetes/kompose), a conversion tool for developers to transition Docker Compose applications to Kubernetes, has graduated from the [Kubernetes Incubator](https://github.com/kubernetes/community/blob/master/incubator.md) to become an official part of the project.
Since our first commit on June 27, 2016, Kompose has achieved 13 releases over 851 commits, gaining 21 contributors since the inception of the project. Our work started at Skippbox (now part of [Bitnami](https://bitnami.com/)) and grew through contributions from Google and Red Hat.
The Kubernetes Incubator allowed contributors to get to know each other across companies, as well as collaborate effectively under guidance from Kubernetes contributors and maintainers. Our incubation led to the development and release of a new and useful tool for the Kubernetes ecosystem.
We’ve created a reliable, scalable Kubernetes environment from an initial Docker Compose file. We worked hard to convert as many keys as possible to their Kubernetes equivalent. Running a single command gets you up and running on Kubernetes:  kompose up.
We couldn’t have done it without feedback and contributions from the community!
If you haven’t yet tried [Kompose on GitHub](https://github.com/kubernetes/kompose) check it out!
Kubernetes guestbook
The go-to example for Kubernetes is the famous [guestbook](https://github.com/kubernetes/examples/blob/master/guestbook), which we use as a base for conversion.
Here is an example from the official [kompose.io](https://kompose.io/) site, starting with a simple Docker Compose [file](https://raw.githubusercontent.com/kubernetes/kompose/master/examples/docker-compose.yaml)).
First, we’ll retrieve the file:
```$ wget https://raw.githubusercontent.com/kubernetes/kompose/master/examples/docker-compose.yaml
```You can test it out by first deploying to Docker Compose:
```$ docker-compose up -d

Creating network &#34;examples\_default&#34; with the default driver

Creating examples\_redis-slave\_1

Creating examples\_frontend\_1

Creating examples\_redis-master\_1
```And when you’re ready to deploy to Kubernetes:
```$ kompose up


We are going to create Kubernetes Deployments, Services and PersistentVolumeClaims for your Dockerized application.


If you need different kind of resources, use the kompose convert and kubectl create -f commands instead.


INFO Successfully created Service: redis          

INFO Successfully created Service: web            

INFO Successfully created Deployment: redis       

INFO Successfully created Deployment: web         


Your application has been deployed to Kubernetes. You can run kubectl get deployment,svc,pods,pvc for details
```Check out [other examples](https://github.com/kubernetes/kompose/tree/master/examples) of what Kompose can do.
Converting to alternative Kubernetes controllers
Kompose can also convert to specific Kubernetes controllers with the use of flags:
```$ kompose convert --help  

Usage:

  kompose convert [file] [flags]


Kubernetes Flags:

      --daemon-set               Generate a Kubernetes daemonset object

  -d, --deployment               Generate a Kubernetes deployment object

  -c, --chart                    Create a Helm chart for converted objects

      --replication-controller   Generate a Kubernetes replication controller object

…
```For example, let’s convert our [guestbook](https://github.com/kubernetes/examples/blob/master/guestbook) example to a DaemonSet:
```$ kompose convert --daemon-set

INFO Kubernetes file &#34;frontend-service.yaml&#34; created

INFO Kubernetes file &#34;redis-master-service.yaml&#34; created

INFO Kubernetes file &#34;redis-slave-service.yaml&#34; created

INFO Kubernetes file &#34;frontend-daemonset.yaml&#34; created

INFO Kubernetes file &#34;redis-master-daemonset.yaml&#34; created

INFO Kubernetes file &#34;redis-slave-daemonset.yaml&#34; created
```Key Kompose 1.0 features
With our graduation, comes the release of Kompose 1.0.0, here’s what’s new:
What’s ahead?
As we continue development, we will strive to convert as many Docker Compose keys as possible for all future and current Docker Compose releases, converting each one to their Kubernetes equivalent. All future releases will be backwards-compatible.
--Charlie Drage, Software Engineer, Red Hat


	

	


