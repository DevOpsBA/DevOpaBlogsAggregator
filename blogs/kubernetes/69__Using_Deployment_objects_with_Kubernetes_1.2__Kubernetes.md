|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/04/Using-Deployment-Objects-With/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-04-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.8113976 &#43;0300 MSK m=&#43;4.400445201  |

#  Using Deployment objects with Kubernetes 1.2  | Kubernetes

	
	
	
	
	*Editor&#39;s note: this is the seventh post in a [series of in-depth posts](https://kubernetes.io/blog/2016/03/five-days-of-kubernetes-12) on what&#39;s new in Kubernetes 1.2*
Kubernetes has made deploying and managing applications very straightforward, with most actions a single API or command line away, including rolling out new applications, canary testing and upgrading. So why would we need Deployments?
Deployment objects automate deploying and rolling updating applications. Compared with kubectl rolling-update, Deployment API is much faster, is declarative, is implemented server-side and has more features (for example, you can rollback to any previous revision even after the rolling update is done).
In today’s blogpost, we’ll cover how to use Deployments to:
[img](https://4.bp.blogspot.com/-M9Xc21XYtLA/Vv7ImzURFxI/AAAAAAAACg0/jlHU3nJ-qYwC74DMiD-joaDPqQfebj3-g/s1600/image03.gif)
Without further ado, let’s start playing around with Deployments!
If you want to try this example, basically you’ll need 3 things:
The configuration files contain a static website. First, we want to start serving its static content. From the root of the Kubernetes repository, run:
```$ kubectl proxy --www=docs/user-guide/update-demo/local/ &amp;  
```Starting to serve on …
This runs a proxy on the default port 8001. You may now visit [http://localhost:8001/static/](http://localhost:8001/static/) the demo website (and it should be a blank page for now). Now we want to run an app and show it on the website.
```$ kubectl run update-demo   
--image=gcr.io/google\_containers/update-demo:nautilus --port=80 -l name=update-demo  

deployment “update-demo” created  
```This deploys 1 replica of an app with the image “update-demo:nautilus” and you can see it visually on [http://localhost:8001/static/](http://localhost:8001/static/).1
[img](https://3.bp.blogspot.com/-EYXhcEK1upw/Vv7JL4rOAtI/AAAAAAAACg4/uy9oKePGjA82xPHhX6ak2_NiHPZ3FU8gw/s1600/deployment-API-5.png)
The card showing on the website represents a Kubernetes pod, with the pod’s name (ID), status, image, and labels.
Now we want more copies of this app!
$ kubectl scale deployment/update-demo --replicas=4
deployment &#34;update-demo&#34; scaled
[img](https://1.bp.blogspot.com/-6YXQqogAGcY/Vv7JnU7g_FI/AAAAAAAAChE/00pqgQvUXkcgjPzi7NfDnSSRJeBUHFaGQ/s1600/deployment-API-2.png)
How about updating the app?
``` $ kubectl edit deployment/update-demo  

 This opens up your default editor, and you can update the deployment on the fly. Find .spec.template.spec.containers[0].image and change nautilus to kitty. Save the file, and you’ll see:  

 deployment &#34;update-demo&#34; edited   
```You’re now updating the image of this app from “update-demo:nautilus” to “update-demo:kitty”.  Deployments allow you to update the app progressively, without a service outage.
[img](https://2.bp.blogspot.com/-x4FmFXdzw30/Vv7KAAQ21wI/AAAAAAAAChM/QWv8Y03lIsU4JBqjE3XFQU2EtzZgogylA/s1600/deployment-API-3.png)
After a while, you’ll find the update seems stuck. What happened?
If you look closer, you’ll find that the pods with the new “kitty” tagged image stays pending. The Deployment automatically stops the rollout if it’s failing. Let’s look at one of the new pod to see what happened:
```$ kubectl describe pod/update-demo-1326485872-a4key  
```Looking at the events of this pod, you’ll notice that Kubernetes failed to pull the image because the “kitty” tag wasn’t found:
Failed to pull image &#34;gcr.io/google_containers/update-demo:kitty&#34;: Tag kitty not found in repository gcr.io/google_containers/update-demo
Ok, now we want to undo the changes and then take our time to figure out which image tag we should use.
```$ kubectl rollout undo deployment/update-demo   
deployment &#34;update-demo&#34; rolled back  
```[img](https://1.bp.blogspot.com/-6YXQqogAGcY/Vv7JnU7g_FI/AAAAAAAAChE/00pqgQvUXkcgjPzi7NfDnSSRJeBUHFaGQ/s1600/deployment-API-2.png)
Everything’s back to normal, phew!
To learn more about rollback, visit [rolling back a Deployment](/docs/user-guide/deployments/#rolling-back-a-deployment).
After a while, we finally figure that the right image tag is “kitten”, instead of “kitty”. Now change .spec.template.spec.containers[0].image tag from “nautilus“ to “kitten“.
```$ kubectl edit deployment/update-demo  
deployment &#34;update-demo&#34; edited  
```[img](https://4.bp.blogspot.com/-u7qPUSQOMLE/Vv7JndUqKaI/AAAAAAAAChA/jHoysiDbnNQU2prPJn19ZFOtLiatzPsMg/s1600/deployment-API-1.png)
Now you see there are 4 cute kittens on the demo website, which means we’ve updated the app successfully! If you want to know the magic behind this, look closer at the Deployment:
```$ kubectl describe deployment/update-demo  
```[img](https://1.bp.blogspot.com/-3U1OTNqdz1s/Vv7Kfw4uGYI/AAAAAAAAChU/CgF6Mv5J6b8_lANXkpEIFytRGo9x0Bn_A/s1600/deployment-API-6.png)
From the events section, you’ll find that the Deployment is managing another resource called [Replica Set](/docs/user-guide/replicasets/), each controls the number of replicas of a different pod template. The Deployment enables progressive rollout by scaling up and down Replica Sets of new and old pod templates.
Now, you’ve learned the basic use of Deployment objects:
***Note:***  *In Kubernetes 1.2, Deployment (beta release) is now feature-complete and enabled by default. For those of you who have tried Deployment in Kubernetes 1.1, please **delete all Deployment 1.1 resources** (including the Replication Controllers and Pods they manage) before trying out Deployments in 1.2. This is necessary because we made some non-backward-compatible changes to the API.*
If you’re interested in Kubernetes and configuration, you’ll want to participate in:
-- *Janet Kuo, Software Engineer, Google*
**1** “kubectl run” outputs the type and name of the resource(s) it creates. In 1.2, it now creates a deployment resource. You can use that in subsequent commands, such as &#34;kubectl get deployment &#34;, or &#34;kubectl expose deployment &#34;. If you want to write a script to do that automatically, in a forward-compatible manner, use &#34;-o name&#34; flag with &#34;kubectl run&#34;, and it will generate short output &#34;deployments/&#34;, which can also be used on subsequent command lines. The &#34;--generator&#34; flag can be used with &#34;kubectl run&#34; to generate other types of resources, for example, set it to &#34;run/v1&#34; to create a Replication Controller, which was the default in 1.1 and 1.0, and to &#34;run-pod/v1&#34; to create a Pod, such as for --restart=Never pods.


	

	


