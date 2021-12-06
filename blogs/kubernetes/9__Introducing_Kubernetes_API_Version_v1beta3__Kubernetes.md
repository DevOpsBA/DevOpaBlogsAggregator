|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/04/Introducing-Kubernetes-V1Beta3/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-04-16 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:50:38.8295226 &#43;0300 MSK m=&#43;8.218201301  |

#  Introducing Kubernetes API Version v1beta3  | Kubernetes

	
	
	
	
	We&#39;ve been hard at work on cleaning up the API over the past several months (see [https://github.com/GoogleCloudPlatform/kubernetes/issues/1519](https://github.com/GoogleCloudPlatform/kubernetes/issues/1519) for details). The result is v1beta3, which is considered to be the release candidate for the v1 API.
We would like you to move to this new API version as soon as possible. v1beta1 and v1beta2 are deprecated, and will be removed by the end of June, shortly after we introduce the v1 API.
As of the latest release, v0.15.0, v1beta3 is the primary, default API. We have changed the default kubectl and client API versions as well as the default storage version (which means objects persisted in etcd will be converted from v1beta1 to v1beta3 as they are rewritten). 
You can take a look at v1beta3 examples such as:
[https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/guestbook/v1beta3](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/guestbook/v1beta3)
[https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/walkthrough/v1beta3](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/walkthrough/v1beta3)
[https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/update-demo/v1beta3](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/examples/update-demo/v1beta3)
To aid the transition, we&#39;ve also created a conversion [tool](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/cluster_management.md#switching-your-config-files-to-a-new-api-version) and put together a list of important [different API changes](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api.md#v1beta3-conversion-tips).
And the most recently generated Swagger specification of the API is here:
[http://kubernetes.io/third_party/swagger-ui/#!/v1beta3](http://kubernetes.io/third_party/swagger-ui/#!/v1beta3)
More details about our approach to API versioning and the transition can be found here:
[https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api.md](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api.md)
Another change we discovered is that with the change to the default API version in kubectl, commands that use &#34;-o template&#34; will break unless you specify &#34;--api-version=v1beta1&#34; or update to v1beta3 syntax. An example of such a change can be seen here:
[https://github.com/GoogleCloudPlatform/kubernetes/pull/6377/files](https://github.com/GoogleCloudPlatform/kubernetes/pull/6377/files)
If you use &#34;-o template&#34;, I recommend always explicitly specifying the API version rather than relying upon the default. We may add this setting to kubeconfig in the future.
Let us know if you have any questions. As always, we&#39;re available on IRC (#google-containers) and github issues.


	

	


