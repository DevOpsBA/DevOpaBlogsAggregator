|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/10/Using-Rbac-Generally-Available-18/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-10-28 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.2791075 &#43;0300 MSK m=&#43;2.868146401  |

#  Using RBAC, Generally Available in Kubernetes v1.8  | Kubernetes

	
	
	
	
	***Editor&#39;s note: this post is part of a [series of in-depth articles](https://kubernetes.io/blog/2017/10/five-days-of-kubernetes-18) on what&#39;s new in Kubernetes 1.8. Today’s post comes from Eric Chiang, software engineer, CoreOS, and SIG-Auth co-lead.***
Kubernetes 1.8 represents a significant milestone for the [role-based access control (RBAC) authorizer](/docs/reference/access-authn-authz/rbac/), which was promoted to GA in this release. RBAC is a mechanism for controlling access to the Kubernetes API, and since its [beta in 1.6](https://kubernetes.io/blog/2017/04/rbac-support-in-kubernetes), many Kubernetes clusters and provisioning strategies have enabled it by default.
Going forward, we expect to see RBAC become a fundamental building block for securing Kubernetes clusters. This post explores using RBAC to manage user and application access to the Kubernetes API.
RBAC is configured using standard Kubernetes resources. Users can be bound to a set of roles (ClusterRoles and Roles) through bindings (ClusterRoleBindings and RoleBindings). Users start with no permissions and must explicitly be granted access by an administrator.
All Kubernetes clusters install a default set of ClusterRoles, representing common buckets users can be placed in. The “edit” role lets users perform basic actions like deploying pods; “view” lets a user observe non-sensitive resources; “admin” allows a user to administer a namespace; and “cluster-admin” grants access to administer a cluster.
```
$ kubectl get clusterroles

NAME            AGE

admin           40m

cluster-admin   40m

edit            40m

# ...


view            40m

```ClusterRoleBindings grant a user, group, or service account a ClusterRole’s power across the entire cluster. Using kubectl, we can let a sample user “jane” perform basic actions in all namespaces by binding her to the “edit” ClusterRole:
```
$ kubectl create clusterrolebinding jane --clusterrole=edit --user=jane

$ kubectl get namespaces --as=jane

NAME          STATUS    AGE

default       Active    43m

kube-public   Active    43m

kube-system   Active    43m

$ kubectl auth can-i create deployments --namespace=dev --as=jane

yes

```RoleBindings grant a ClusterRole’s power within a namespace, allowing administrators to manage a central list of ClusterRoles that are reused throughout the cluster. For example, as new resources are added to Kubernetes, the default ClusterRoles are updated to automatically grant the correct permissions to RoleBinding subjects within their namespace.
Next we’ll let the group “infra” modify resources in the “dev” namespace:
```
$ kubectl create rolebinding infra --clusterrole=edit --group=infra --namespace=dev

rolebinding &#34;infra&#34; created

```Because we used a RoleBinding, these powers only apply within the RoleBinding’s namespace. In our case, a user in the “infra” group can view resources in the “dev” namespace but not in “prod”:
```
$ kubectl get deployments --as=dave --as-group=infra --namespace dev

No resources found.

$ kubectl get deployments --as=dave --as-group=infra --namespace prod

Error from server (Forbidden): deployments.extensions is forbidden: User &#34;dave&#34; cannot list deployments.extensions in the namespace &#34;prod&#34;.

```When the default ClusterRoles aren’t enough, it’s possible to create new roles that define a custom set of permissions. Since ClusterRoles are just regular API resources, they can be expressed as YAML or JSON manifests and applied using kubectl.
Each ClusterRole holds a list of permissions specifying “rules.” Rules are purely additive and allow specific HTTP verb to be performed on a set of resource. For example, the following ClusterRole holds the permissions to perform any action on &#34;deployments”, “configmaps,” or “secrets”, and to view any “pod”:
```
kind: ClusterRole

apiVersion: rbac.authorization.k8s.io/v1

metadata:

 name: deployer

rules:

- apiGroups: [&#34;apps&#34;]

 resources: [&#34;deployments&#34;]

 verbs: [&#34;get&#34;, &#34;list&#34;, &#34;watch&#34;, &#34;create&#34;, &#34;delete&#34;, &#34;update&#34;, &#34;patch&#34;]



- apiGroups: [&#34;&#34;] # &#34;&#34; indicates the core API group

 resources: [&#34;configmaps&#34;, &#34;secrets&#34;]

 verbs: [&#34;get&#34;, &#34;list&#34;, &#34;watch&#34;, &#34;create&#34;, &#34;delete&#34;, &#34;update&#34;, &#34;patch&#34;]



- apiGroups: [&#34;&#34;] # &#34;&#34; indicates the core API group

 resources: [&#34;pods&#34;]

 verbs: [&#34;get&#34;, &#34;list&#34;, &#34;watch&#34;]

```Verbs correspond to the HTTP verb of the request, while the resource and API groups refer to the resource being referenced. Consider the following Ingress resource:
```
apiVersion: extensions/v1beta1

kind: Ingress

metadata:

 name: test-ingress

spec:

 backend:

   serviceName: testsvc

   servicePort: 80

```To POST the resource, the user would need the following permissions:
```
rules:

- apiGroups: [&#34;extensions&#34;] # &#34;apiVersion&#34; without version

 resources: [&#34;ingresses&#34;]  # Plural of &#34;kind&#34;

 verbs: [&#34;create&#34;]         # &#34;POST&#34; maps to &#34;create&#34;

```When deploying containers that require access to the Kubernetes API, it’s good practice to ship an RBAC Role with your application manifests. Besides ensuring your app works on RBAC enabled clusters, this helps users audit what actions your app will perform on the cluster and consider their security implications.
A namespaced Role is usually more appropriate for an application, since apps are traditionally run inside a single namespace and the namespace&#39;s resources should be tied to the lifecycle of the app. However, Roles cannot grant access to non-namespaced resources (such as nodes) or across namespaces, so some apps may still require ClusterRoles.
The following Role allows a Prometheus instance to monitor and discover services, endpoints, and pods in the “dev” namespace:
```
kind: Role

metadata:

 name: prometheus-role

 namespace: dev

rules:

- apiGroups: [&#34;&#34;] # &#34;&#34; refers to the core API group

 Resources: [&#34;services&#34;, &#34;endpoints&#34;, &#34;pods&#34;]

 verbs: [&#34;get&#34;, &#34;list&#34;, &#34;watch&#34;]

```Containers running in a Kubernetes cluster receive service account credentials to talk to the Kubernetes API, and service accounts can be targeted by a RoleBinding. Pods normally run with the “default” service account, but it’s good practice to run each app with a unique service account so RoleBindings don’t unintentionally grant permissions to other apps.
To run a pod with a custom service account, create a ServiceAccount resource in the same namespace and specify the ```serviceAccountName``` field of the manifest.
```
apiVersion: apps/v1beta2 # Abbreviated, not a full manifest

kind: Deployment

metadata:

 name: prometheus-deployment

 namespace: dev

spec:

 replicas: 1

 template:

   spec:

     containers:

     - name: prometheus

       image: prom/prometheus:v1.8.0

       command: [&#34;prometheus&#34;, &#34;-config.file=/etc/prom/config.yml&#34;]

   # Run this pod using the &#34;prometheus-sa&#34; service account.

   serviceAccountName: prometheus-sa

---

apiVersion: v1

kind: ServiceAccount

metadata:

 name: prometheus-sa

 namespace: dev

```Development of RBAC is a community effort organized through the [Auth Special Interest Group](https://github.com/kubernetes/community/blob/master/sig-auth/README.md), one of the [many SIGs](https://github.com/kubernetes/community/blob/master/sig-list.md) responsible for maintaining Kubernetes. A great way to get involved in the Kubernetes community is to join a SIG that aligns with your interests, provide feedback, and help with the roadmap.
Eric Chiang is a software engineer and technical lead of Kubernetes development at [CoreOS](https://coreos.com/?utm_source=k8sblog&amp;utm_medium=social&amp;utm_campaign=organic), the creator of Tectonic, the enterprise-ready Kubernetes platform. Eric co-leads Kubernetes SIG Auth and maintains several open source projects and libraries on behalf of CoreOS.


	

	


