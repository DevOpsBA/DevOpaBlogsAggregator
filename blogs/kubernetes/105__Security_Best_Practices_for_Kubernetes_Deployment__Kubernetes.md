|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/08/Security-Best-Practices-Kubernetes-Deployment/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-08-31 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.2790209 &#43;0300 MSK m=&#43;3.868065501  |

#  Security Best Practices for Kubernetes Deployment  | Kubernetes

	
	
	
	
	*Note: some of the recommendations in this post are no longer current. Current cluster hardening options are described in this [documentation](/docs/tasks/administer-cluster/securing-a-cluster/).*
*Editor’s note: today’s post is by Amir Jerbi and Michael Cherny of Aqua Security, describing security best practices for Kubernetes deployments, based on data they’ve collected from various use-cases seen in both on-premises and cloud deployments.*
Kubernetes provides many controls that can greatly improve your application security. Configuring them requires intimate knowledge with Kubernetes and the deployment’s security requirements. The best practices we highlight here are aligned to the container lifecycle: build, ship and run, and are specifically tailored to Kubernetes deployments. We adopted these best practices in [our own SaaS deployment](http://blog.aquasec.com/running-a-security-service-in-google-cloud-real-world-example) that runs Kubernetes on Google Cloud Platform.
The following are our recommendations for deploying a secured Kubernetes application:
**Ensure That Images Are Free of Vulnerabilities **
Having running containers with vulnerabilities opens your environment to the risk of being easily compromised. Many of the attacks can be mitigated simply by making sure that there are no software components that have known vulnerabilities.
**Ensure That Only Authorized Images are Used in Your Environment** 
Without a process that ensures that only images adhering to the organization’s policy are allowed to run, the organization is open to risk of running vulnerable or even malicious containers. Downloading and running images from unknown sources is dangerous. It is equivalent to running software from an unknown vendor on a production server. Don’t do that.
Use private registries to store your approved images - make sure you only push approved images to these registries. This alone already narrows the playing field, reducing the number of potential images that enter your pipeline to a fraction of the hundreds of thousands of publicly available images. Build a CI pipeline that integrates security assessment (like vulnerability scanning), making it part of the build process.
The CI pipeline should ensure that only vetted code (approved for production) is used for building the images. Once an image is built, it should be scanned for security vulnerabilities, and only if no issues are found then the image would be pushed to a private registry, from which deployment to production is done. A failure in the security assessment should create a failure in the pipeline, preventing images with bad security quality from being pushed to the image registry.
There is work in progress being done in Kubernetes for image authorization plugins (expected in Kubernetes 1.4), which will allow preventing the shipping of unauthorized images. For more info see this [pull request](https://github.com/kubernetes/kubernetes/pull/27129).
**Limit Direct Access to Kubernetes Nodes**
You should limit SSH access to Kubernetes nodes, reducing the risk for unauthorized access to host resource. Instead you should ask users to use &#34;kubectl exec&#34;, which will provide direct access to the container environment without the ability to access the host.
You can use Kubernetes [Authorization Plugins](/docs/reference/access-authn-authz/authorization/) to further control user access to resources. This allows defining fine-grained-access control rules for specific namespace, containers and operations.
**Create Administrative Boundaries between Resources**
Limiting the scope of user permissions can reduce the impact of mistakes or malicious activities. A Kubernetes namespace allows you to partition created resources into logically named groups. Resources created in one namespace can be hidden from other namespaces. By default, each resource created by a user in Kubernetes cluster runs in a default namespace, called default. You can create additional namespaces and attach resources and users to them. You can use Kubernetes Authorization plugins to create policies that segregate access to namespace resources between different users.
For example: the following policy will allow ‘alice’ to read pods from namespace ‘fronto’.
```{

  &#34;apiVersion&#34;: &#34;abac.authorization.kubernetes.io/v1beta1&#34;,

  &#34;kind&#34;: &#34;Policy&#34;,

  &#34;spec&#34;: {

    &#34;user&#34;: &#34;alice&#34;,

    &#34;namespace&#34;: &#34;fronto&#34;,

    &#34;resource&#34;: &#34;pods&#34;,

    &#34;readonly&#34;: true

  }

}
```**Define Resource Quota**
An option of running resource-unbound containers puts your system in risk of DoS or “noisy neighbor” scenarios. To prevent and minimize those risks you should define resource quotas. By default, all resources in Kubernetes cluster are created with unbounded CPU and memory requests/limits. You can create resource quota policies, attached to Kubernetes namespace, in order to limit the CPU and memory a pod is allowed to consume.
The following is an example for namespace resource quota definition that will limit number of pods in the namespace to 4, limiting their CPU requests between 1 and 2 and memory requests between 1GB to 2GB.
compute-resources.yaml:
```apiVersion: v1  
kind: ResourceQuota  
metadata:  
  name: compute-resources  
spec:  
  hard:  
    pods: &#34;4&#34;  
    requests.cpu: &#34;1&#34;  
    requests.memory: 1Gi  
    limits.cpu: &#34;2&#34;  
    limits.memory: 2Gi
```Assign a resource quota to namespace:
```kubectl create -f ./compute-resources.yaml --namespace=myspace
```**Implement Network Segmentation**
Running different applications on the same Kubernetes cluster creates a risk of one compromised application attacking a neighboring application. Network segmentation is important to ensure that containers can communicate only with those they are supposed to.
One of the challenges in Kubernetes deployments is creating network segmentation between pods, services and containers. This is a challenge due to the “dynamic” nature of container network identities (IPs), along with the fact that containers can communicate both inside the same node or between nodes.
Users of Google Cloud Platform can benefit from automatic firewall rules, preventing cross-cluster communication. A similar implementation can be deployed on-premises using network firewalls or SDN solutions. There is work being done in this area by the Kubernetes [Network SIG](https://github.com/kubernetes/community/wiki/SIG-Network), which will greatly improve the pod-to-pod communication policies. A new network policy API should address the need to create firewall rules around pods, limiting the network access that a containerized can have.
The following is an example of a network policy that controls the network for “backend” pods, only allowing inbound network access from “frontend” pods:
```POST /apis/net.alpha.kubernetes.io/v1alpha1/namespaces/tenant-a/networkpolicys  
{  
  &#34;kind&#34;: &#34;NetworkPolicy&#34;,

  &#34;metadata&#34;: {

    &#34;name&#34;: &#34;pol1&#34;

  },

  &#34;spec&#34;: {

    &#34;allowIncoming&#34;: {

      &#34;from&#34;: [{

        &#34;pods&#34;: { &#34;segment&#34;: &#34;frontend&#34; }

      }],

      &#34;toPorts&#34;: [{

        &#34;port&#34;: 80,

        &#34;protocol&#34;: &#34;TCP&#34;

      }]

    },

    &#34;podSelector&#34;: {

      &#34;segment&#34;: &#34;backend&#34;

    }

  }

}
```Read more about Network policies [here](https://kubernetes.io/blog/2016/04/Kubernetes-Network-Policy-APIs).
**Apply Security Context to Your Pods and Containers**
When designing your containers and pods, make sure that you configure the security context for your pods, containers and volumes. A security context is a property defined in the deployment yaml. It controls the security parameters that will be assigned to the pod/container/volume. Some of the important parameters are:
The following is an example for pod definition with security context parameters:
```apiVersion: v1  
kind: Pod  
metadata:  
  name: hello-world  
spec:  
  containers:  
  # specification of the pod’s containers  
  # ...  
  securityContext:  
    readOnlyRootFilesystem: true  
    runAsNonRoot: true
```Reference [here](/docs/api-reference/v1/definitions/#_v1_podsecuritycontext).
In case you are running containers with elevated privileges (--privileged) you should consider using the “DenyEscalatingExec” admission control. This control denies exec and attach commands to pods that run with escalated privileges that allow host access. This includes pods that run as privileged, have access to the host IPC namespace, and have access to the host PID namespace. For more details on admission controls, see the Kubernetes [documentation](/docs/reference/access-authn-authz/admission-controllers/).
**Log Everything**
Kubernetes supplies cluster-based logging, allowing to log container activity into a central log hub. When a cluster is created, the standard output and standard error output of each container can be ingested using a Fluentd agent running on each node into either Google Stackdriver Logging or into Elasticsearch and viewed with Kibana.
**Summary**
Kubernetes supplies many options to create a secured deployment. There is no one-size-fit-all solution that can be used everywhere, so a certain degree of familiarity with these options is required, as well as an understanding of how they can enhance your application’s security.
We recommend implementing the best practices that were highlighted in this blog, and use Kubernetes flexible configuration capabilities to incorporate security processes into the continuous integration pipeline, automating the entire process with security seamlessly “baked in”.
*--Michael Cherny, Head of Security Research, and Amir Jerbi, CTO and co-founder [Aqua Security](https://www.aquasec.com/)*


	

	


