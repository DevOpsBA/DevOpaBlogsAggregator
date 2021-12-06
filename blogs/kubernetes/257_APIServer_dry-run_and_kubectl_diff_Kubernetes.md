|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2019/01/14/apiserver-dry-run-and-kubectl-diff/        |
| Tags              | [kubernetes]       |
| Date Create       | 2019-01-14 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.4255819 &#43;0300 MSK m=&#43;2.014615901  |

# APIServer dry-run and kubectl diff | Kubernetes

	
	
	
	
	**Author**: Antoine Pelisse (Google Cloud, @apelisse)
Declarative configuration management, also known as configuration-as-code, is
one of the key strengths of Kubernetes. It allows users to commit the desired state of
the cluster, and to keep track of the different versions, improve auditing and
automation through CI/CD pipelines. The [Apply working-group](https://groups.google.com/forum/#!forum/kubernetes-wg-apply)
is working on fixing some of the gaps, and is happy to announce that Kubernetes
1.13 promoted server-side dry-run and ```kubectl diff``` to beta. These
two features are big improvements for the Kubernetes declarative model.
A few pieces are still missing in order to have a seamless declarative
experience with Kubernetes, and we tried to address some of these:
The working group has tried to address these problems.
[APIServer dry-run](/docs/reference/using-api/api-concepts/#dry-run) was implemented to address these two problems:
While dynamic admission controllers are not supposed to have side-effects on
each request, dry-run requests are only processed if all admission controllers
explicitly announce that they don&#39;t have any dry-run side-effects.
Server-side dry-run is enabled through a feature-gate. Now that the feature is
Beta in 1.13, it should be enabled by default, but still can be enabled/disabled
using ```kube-apiserver --feature-gates DryRun=true```.
If you have dynamic admission controllers, you might have to fix them to:
You can trigger the feature from kubectl by using ```kubectl apply --server-dry-run```, which will decorate the request with the dryRun flag
and return the object as it would have been applied, or an error if it would
have failed.
APIServer dry-run is convenient because it lets you see how the object would be
processed, but it can be hard to identify exactly what changed if the object is
big. ```kubectl diff``` does exactly what you want by showing the differences between
the current &#34;live&#34; object and the new &#34;dry-run&#34; object. It makes it very
convenient to focus on only the changes that are made to the object, how the
server has merged these and how the mutating webhooks affects the output.
```kubectl diff``` is meant to be as similar as possible to ```kubectl apply```:
```kubectl diff -f some-resources.yaml``` will show a diff for the resources in the yaml file. One can even use the diff program of their choice by using the KUBECTL_EXTERNAL_DIFF environment variable, for example:
```KUBECTL_EXTERNAL_DIFF=meld kubectl diff -f some-resources.yaml
```The working group is still busy trying to improve some of these things:


	

	


