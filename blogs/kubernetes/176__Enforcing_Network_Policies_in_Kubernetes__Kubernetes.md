|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/10/Enforcing-Network-Policies-In-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-10-30 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.2761113 &#43;0300 MSK m=&#43;2.865150201  |

#  Enforcing Network Policies in Kubernetes  | Kubernetes

	
	
	
	
	***Editor&#39;s note: this post is part of a [series of in-depth articles](https://kubernetes.io/blog/2017/10/five-days-of-kubernetes-18) on what&#39;s new in Kubernetes 1.8. Today’s post comes from Ahmet Alp Balkan, Software Engineer, Google.***
Kubernetes now offers functionality to enforce rules about which pods can communicate with each other using [network policies](/docs/concepts/services-networking/network-policies/). This feature is has become stable Kubernetes 1.7 and is ready to use with supported networking plugins. The Kubernetes 1.8 release has added better capabilities to this feature.
In a Kubernetes cluster configured with default settings, all pods can discover and communicate with each other without any restrictions. The new Kubernetes object type NetworkPolicy lets you allow and block traffic to pods.
If you’re running multiple applications in a Kubernetes cluster or sharing a cluster among multiple teams, it’s a security best practice to create firewalls that permit pods to talk to each other while blocking other network traffic. Networking policy corresponds to the Security Groups concepts in the Virtual Machines world.
Networking Policies are implemented by networking plugins. These plugins typically install an overlay network in your cluster to enforce the Network Policies configured. A number of networking plugins, including [Calico](/docs/tasks/configure-pod-container/calico-network-policy/), [Romana](/docs/tasks/configure-pod-container/romana-network-policy/) and [Weave Net](/docs/tasks/configure-pod-container/weave-network-policy/), support using Network Policies.
Google Container Engine (GKE) also provides beta support for [Network Policies](https://cloud.google.com/container-engine/docs/network-policy) using the Calico networking plugin when you create clusters with the following command:
gcloud beta container clusters create --enable-network-policy
Once you install a networking plugin that implements Network Policies, you need to create a Kubernetes resource of type NetworkPolicy. This object describes two set of label-based pod selector fields, matching:
The following example of a network policy blocks all in-cluster traffic to a set of web server pods, except the pods allowed by the policy configuration.

To achieve this setup, create a NetworkPolicy with the following manifest:
```kind: NetworkPolicy

apiVersion: networking.k8s.io/v1

metadata:

  name: access-nginx

spec:

  podSelector:

    matchLabels:

      app: nginx

  ingress:

  - from:

    - podSelector:

        matchLabels:

          app: foo
```Once you apply this configuration, only pods with label **app: foo** can talk to the pods with the label **app: nginx**. For a more detailed tutorial, see the [Kubernetes documentation](/docs/tasks/administer-cluster/declare-network-policy/).
If you specify the spec.podSelector field as empty, the set of pods the network policy matches to all pods in the namespace, blocking all traffic between pods by default. In this case, you must explicitly create network policies whitelisting all communication between the pods.

You can enable a policy like this by applying the following manifest in your Kubernetes cluster:
```apiVersion: networking.k8s.io/v1

kind: NetworkPolicy

metadata:

  name: default-deny

spec:

  podSelector:
```In addition to the previous examples, you can make the Network Policy API enforce more complicated rules:


	

	


