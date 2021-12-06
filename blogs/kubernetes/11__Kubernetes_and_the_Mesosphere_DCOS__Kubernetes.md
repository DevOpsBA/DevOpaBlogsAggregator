|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/04/Kubernetes-And-Mesosphere-Dcos/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-04-22 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.5935801 &#43;0300 MSK m=&#43;5.182632201  |

#  Kubernetes and the Mesosphere DCOS  | Kubernetes

	
	
	
	
	Today Mesosphere announced the addition of Kubernetes as a standard part of their [DCOS](https://mesosphere.com/product/) offering.  This is a great step forwards in bringing cloud native application management to the world, and should lay to rest many questions we hear about &#39;Kubernetes or Mesos, which one should I use?&#39;.  Now you can have your cake and eat it too:  use both.  Today&#39;s announcement extends the reach of Kubernetes to a new class of users, and add some exciting new capabilities for everyone.
By way of background, Kubernetes is a cluster management framework that was started by Google nine months ago, inspired by the internal system known as Borg.  You can learn a little more about Borg by checking out this [paper](http://research.google.com/pubs/pub43438.html).  At the heart of it Kubernetes offers what has been dubbed &#39;cloud native&#39; application management.  To us, there are three things that together make something &#39;cloud native&#39;:
Kubernetes was designed from the start to make these capabilities available to everyone, and built by the same engineers that built the system internally known as Borg.  For many users the promise of &#39;Google style app management&#39; is interesting, but they want to run these new classes of applications on the same set of physical resources as their existing workloads like Hadoop, Spark, Kafka, etc.  Now they will have access to commercially supported offering that brings the two worlds together.
Mesosphere, one of the earliest supporters of the Kubernetes project, has been working closely with the core Kubernetes team to create a natural experience for users looking to get the best of both worlds, adding Kubernetes to every Mesos deployment they instantiate, whether it be in the public cloud, private cloud, or in a hybrid deployment model.  This is well aligned with the overall Kubernetes vision of creating ubiquitous management framework that runs anywhere a container can.  It will be interesting to see how you blend together the old world and the new on a commercially supported, versatile platform.
Craig McLuckie
Product Manager, Google and Kubernetes co-founder


	

	


