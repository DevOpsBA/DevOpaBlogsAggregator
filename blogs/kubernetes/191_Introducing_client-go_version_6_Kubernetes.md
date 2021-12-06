|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/01/Introducing-Client-Go-Version-6/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-01-12 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.0621793 &#43;0300 MSK m=&#43;2.651217001  |

# Introducing client-go version 6 | Kubernetes

	
	
	
	
	The Kubernetes API server [exposes a REST interface](https://blog.openshift.com/tag/api-server/) consumable by any client. [client-go](https://github.com/kubernetes/client-go) is the official client library for the Go programming language. It is used both internally by Kubernetes itself (for example, inside kubectl) as well as by [numerous external consumers](https://github.com/search?q=k8s.io%2Fclient-go&amp;type=Code&amp;utf8=%E2%9C%93):operators like the [etcd-operator](https://github.com/coreos/etcd-operator) or [prometheus-operator;](https://github.com/coreos/prometheus-operator)higher level frameworks like [KubeLess](https://github.com/kubeless/kubeless) and [OpenShift](https://openshift.io/); and many more.
The version 6 update to client-go adds support for Kubernetes 1.9, allowing access to the latest Kubernetes features. While the [changelog](https://github.com/kubernetes/client-go/blob/master/CHANGELOG.md) contains all the gory details, this blog post highlights the most prominent changes and intends to guide on how to upgrade from version 5.
This blog post is one of a number of efforts to make client-go more accessible to third party consumers. Easier access is a joint effort by a number of people from numerous companies, all meeting in the #client-go-docs channel of the [Kubernetes Slack](http://slack.k8s.io/). We are happy to hear feedback and ideas for further improvement, and of course appreciate anybody who wants to contribute.
The following API group promotions are part of Kubernetes 1.9:
In Kubernetes 1.8 we introduced CustomResourceDefinitions (CRD) [pre-persistence schema validation](/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/#validation) as an alpha feature. With 1.9, the feature got promoted to beta and will be enabled by default. As a client-go user, you will find the API types at k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.
The [OpenAPI v3 schema](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#schemaObject) can be defined in the CRD spec as:
```
apiVersion: apiextensions.k8s.io/v1beta1  
kind: CustomResourceDefinition  
metadata: ...  
spec:  
  ...  
  validation:  
    openAPIV3Schema:  
      properties:  
        spec:  
          properties:  
            version:  
                type: string  
                enum:  
                - &#34;v1.0.0&#34;  
                - &#34;v1.0.1&#34;  
            replicas:  
                type: integer  
                minimum: 1  
                maximum: 10

```The schema in the above CRD applies following validations for the instance:
```
apiVersion: mygroup.example.com/v1  
kind: App  
metadata:  
  name: example-app  
spec:  
  version: &#34;v1.0.2&#34;  
  replicas: 15

``````$ kubectl create -f app.yaml

The App &#34;example-app&#34; is invalid: []: Invalid value: map[string]interface {}{&#34;apiVersion&#34;:&#34;mygroup.example.com/v1&#34;, &#34;kind&#34;:&#34;App&#34;, &#34;metadata&#34;:map[string]interface {}{&#34;creationTimestamp&#34;:&#34;2017-08-31T20:52:54Z&#34;, &#34;uid&#34;:&#34;5c674651-8e8e-11e7-86ad-f0761cb232d1&#34;, &#34;clusterName&#34;:&#34;&#34;, &#34;name&#34;:&#34;example-app&#34;, &#34;namespace&#34;:&#34;default&#34;, &#34;deletionTimestamp&#34;:interface {}(nil), &#34;deletionGracePeriodSeconds&#34;:(\*int64)(nil)}, &#34;spec&#34;:map[string]interface {}{&#34;replicas&#34;:15, &#34;version&#34;:&#34;v1.0.2&#34;}}:
validation failure list:  
spec.replicas in body should be less than or equal to 10  
spec.version in body should be one of [v1.0.0 v1.0.1]
```Note that with [Admission Webhooks](/docs/reference/access-authn-authz/extensible-admission-controllers/#admission-webhooks), Kubernetes 1.9 provides another beta feature to validate objects before they are created or updated. Starting with 1.9, these webhooks also allow mutation of objects (for example, to set defaults or to inject values). Of course, webhooks work with CRDs as well. Moreover, webhooks can be used to implement validations that are not easily expressible with CRD validation. Note that webhooks are harder to implement than CRD validation, so for many purposes, CRD validation is the right tool.
Often objects in one namespace or only with certain labels are to be processed in a controller. Informers [now allow](https://github.com/kubernetes/kubernetes/pull/54660) you to tweak the ListOptions used to query the API server to list and watch objects. Uninitialized objects (for consumption by [initializers](/docs/reference/access-authn-authz/extensible-admission-controllers/#what-are-initializers)) can be made visible by setting IncludeUnitialized to true. All this can be done using the new NewFilteredSharedInformerFactory constructor for shared informers:
```
import “k8s.io/client-go/informers”
...  
sharedInformers := informers.NewFilteredSharedInformerFactory(  
 client,  
 30\*time.Minute,   
 “some-namespace”,  
 func(opt \*metav1.ListOptions) {  
  opt.LabelSelector = “foo=bar”  
 },  
)  
```Note that the corresponding lister will only know about the objects matching the namespace and the given ListOptions. Note that the same restrictions apply for a List or Watch call on a client.
This [production code example](https://github.com/jetstack/cert-manager/blob/b978faa28c9f0fb0414b5d7293fab7bde65bde76/cmd/controller/app/controller.go#L123) of a cert-manager demonstrates how namespace informers can be used in real code.
Historically, only types in the extensions API group would work with autogenerated Scale clients. Furthermore, different API groups use different Scale types for their /scale subresources. To remedy these issues, k8s.io/client-go/scale provides a [polymorphic scale client](https://github.com/kubernetes/client-go/tree/master/scale) to scale different resources in different API groups in a coherent way:
```
import (


apimeta &#34;k8s.io/apimachinery/pkg/api/meta&#34;

 discocache &#34;k8s.io/client-go/discovery/cached&#34;  
 &#34;k8s.io/client-go/discovery&#34;

&#34;k8s.io/client-go/dynamic&#34;

“k8s.io/client-go/scale”  
)

...

cachedDiscovery := discocache.NewMemCacheClient(client.Discovery())  
restMapper := discovery.NewDeferredDiscoveryRESTMapper(

cachedDiscovery,

apimeta.InterfacesForUnstructured,

)  
scaleKindResolver := scale.NewDiscoveryScaleKindResolver(

client.Discovery(),

)  
scaleClient, err := scale.NewForConfig(

client, restMapper,

dynamic.LegacyAPIPathResolverFunc,

scaleKindResolver,

)
scale, err := scaleClient.Scales(&#34;default&#34;).Get(groupResource, &#34;foo&#34;)

```The returned scale object is generic and is exposed as the autoscaling/v1.Scale object. It is backed by an internal Scale type, with conversions defined to and from all the special Scale types in the API groups supporting scaling. We planto [extend this to CustomResources in 1.10](https://github.com/kubernetes/kubernetes/pull/55168).
If you’re implementing support for the scale subresource, we recommend that you expose the autoscaling/v1.Scale object.
Deeply copying an object formerly required a call to Scheme.Copy(Object) with the notable disadvantage of losing type safety. A typical piece of code from client-go version 5 required type casting:
```
newObj, err := runtime.NewScheme().Copy(node)


if err != nil {

    return fmt.Errorf(&#34;failed to copy node %v: %s”, node, err)

}


newNode, ok := newObj.(\*v1.Node)

if !ok {

    return fmt.Errorf(&#34;failed to type-assert node %v&#34;, newObj)


}

```Thanks to [k8s.io/code-generator](https://github.com/kubernetes/code-generator), Copy has now been replaced by a type-safe DeepCopy method living on each object, allowing you to simplify code significantly both in terms of volume and API error surface:
newNode := node.DeepCopy()
No error handling is necessary: this call never fails. If and only if the node is nil does DeepCopy() return nil.
To copy runtime.Objects there is an additional DeepCopyObject() method in the runtime.Object interface.
With the old method gone for good, clients need to update their copy invocations accordingly.
Using client-go’s dynamic client to access CustomResources is discouraged and superseded by type-safe code using the generators in [k8s.io/code-generator](https://github.com/kubernetes/code-generator). Check out the [Deep Dive on the Open Shift blog](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/) to learn about using code generation with client-go.
You can now place tags in the comment block just above a type or function, or in the second block above. There is no distinction anymore between these two comment blocks. This used to a be a source of [subtle errors when using the generators](https://github.com/kubernetes/kubernetes/issues/53893):
```// second block above  
// &#43;k8s:some-tag  

// first block above  
// &#43;k8s:another-tag  
type Foo struct {}
```You can now use extended tag definitions to create custom verbs . This lets you expand beyond the verbs defined by HTTP. This opens the door to higher levels of customization.
For example, this block leads to the generation of the method UpdateScale(s *autoscaling.Scale) (*autoscaling.Scale, error):
```// genclient:method=UpdateScale,verb=update,subresource=scale,input=k8s.io/kubernetes/pkg/apis/autoscaling.Scale,result=k8s.io/kubernetes/pkg/apis/autoscaling.Scale
```In more complex API groups it’s possible for Kinds, the group name, the Go package name, and the Go group alias name to conflict. This was not handled correctly prior to 1.9. The following tags resolve naming conflicts and make the generated code prettier:
```// &#43;groupName=example2.example.com  
// &#43;groupGoName=SecondExample
```These are usually [in the doc.go file of an API package](https://github.com/kubernetes/code-generator/blob/release-1.9/_examples/crd/apis/example2/v1/doc.go#L18). The first is used as the CustomResource group name when RESTfully speaking to the API server using HTTP. The second is used in the generated Golang code (for example, in the clientset) to access the group version:
clientset.SecondExampleV1()
It’s finally possible to have dots in Go package names. In this section’s example, you would put the groupName snippet into the pkg/apis/example2.example.com directory of your project.
Kubernetes 1.9 includes a number of example projects which can serve as a blueprint for your own projects:
In order to update from the previous version 5 to version 6 of client-go, the library itself as well as certain third-party dependencies must be updated. Previously, this process had been tedious due to the fact that a lot of code got refactored or relocated within the existing package layout across releases. Fortunately, far less code had to move in the latest version, which should ease the upgrade procedure for most users.
In the past [k8s.io/client-go](https://github.com/kubernetes/client-go), [k8s.io/api](https://github.com/kubernetes/api), and [k8s.io/apimachinery](https://github.com/kubernetes/apimachinery) were updated infrequently. Tags (for example, v4.0.0) were created quite some time after the Kubernetes releases. With the 1.9 release we resumed running a nightly bot that updates all the repositories for public consumption, even before manual tagging. This includes the branches:
These tags have limited test coverage, but can be used by early adopters of client-go and the other libraries. Moreover, they help to vendor the correct version of [k8s.io/api](https://github.com/kubernetes/api) and [k8s.io/apimachinery](https://github.com/kubernetes/apimachinery). Note that we only create a v6.0.3-like semantic versioning tag on [k8s.io/client-go](https://github.com/kubernetes/client-go). The corresponding tag for k8s.io/api and k8s.io/apimachinery is kubernetes-1.9.3.
Also note that only these tags correspond to tested releases of Kubernetes. If you depend on the release branch, e.g., release-1.9, your client is running on unreleased Kubernetes code.
In general, the list of which dependencies to vendor is automatically generated and written to the file Godeps/Godeps.json. Only the revisions listed there are tested. This means especially that we do not and cannot test the code-base against master branches of our dependencies. This puts us in the following situation depending on the used vendoring tool:
Even with the deficiencies of golang/dep today, dep is slowly becoming the de-facto standard in the Go ecosystem. With the necessary care and the awareness of the missing features, dep can be (and is!) used successfully. Here’s a demonstration of how to update a project with client-go 5 to the latest version 6 using dep:
(If you are still running client-go version 4 and want to play it safe by not skipping a release, now is a good time to check out [this excellent blog post](https://medium.com/@andy.goldstein/upgrading-kubernetes-client-go-from-v4-to-v5-bbd5025fe381) describing how to upgrade to version 5, put together by our friends at Heptio.)
Before starting, it is important to understand that client-go depends on two other Kubernetes projects: [k8s.io/apimachinery](https://github.com/kubernetes/apimachinery) and [k8s.io/api](https://github.com/kubernetes/api). In addition, if you are using CRDs, you probably also depend on [k8s.io/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver) for the CRD client. The first exposes lower-level API mechanics (such as schemes, serialization, and type conversion), the second holds API definitions, and the third provides APIs related to CustomResourceDefinitions. In order for client-go to operate correctly, it needs to have its companion libraries vendored in correspondingly matching versions. Each library repository provides a branch named release-*&lt;version&gt;* where *&lt;version&gt;* refers to a particular Kubernetes version; for client-go version 6, it is imperative to refer to the *release*-1.9 branch on each repository.
Assuming the latest version 5 patch release of client-go being vendored through dep, the Gopkg.toml manifest file should look something like this (possibly using branches instead of versions):
```




[[constraint]]


  name = &#34;k8s.io/api&#34;

  version = &#34;kubernetes-1.8.1&#34;


[[constraint]]

  name = &#34;k8s.io/apimachinery&#34;

  version = &#34;kubernetes-1.8.1&#34;


[[constraint]]

  name = &#34;k8s.io/apiextensions-apiserver&#34;

  version = &#34;kubernetes-1.8.1&#34;


[[constraint]]

  name = &#34;k8s.io/client-go&#34;




  version = &#34;5.0.1&#34;

```Note that some of the libraries could be missing if they are not actually needed by the client.
Upgrading to client-go version 6 means bumping the version and tag identifiers as following ( **emphasis** given):
```




[constraint]]


  name = &#34;k8s.io/api&#34;

  version = &#34;kubernetes-1.9.0&#34;


[[constraint]]

  name = &#34;k8s.io/apimachinery&#34;

  version = &#34;kubernetes-1.9.0&#34;


[[constraint]]

  name = &#34;k8s.io/apiextensions-apiserver&#34;

  version = &#34;kubernetes-1.9.0&#34;


[[constraint]]

  name = &#34;k8s.io/client-go&#34;




  version = &#34;6.0.0&#34;



```The result of the upgrade can be found [here](https://github.com/ncdc/client-go-4-to-5/tree/v5-to-v6).
A note of caution: dep cannot capture the complete set of dependencies in a reliable and reproducible fashion as described above. This means that for a 100% future-proof project you have to add constraints (or even overrides) to many other packages listed in client-go’s Godeps/Godeps.json. Be prepared to add them if something breaks. We are working with the golang/dep community to make this an easier and more smooth experience.
Finally, we need to tell dep to upgrade to the specified versions by executing dep ensure. If everything goes well, the output of the command invocation should be empty, with the only indication that it was successful being a number of updated files inside the vendor folder.
If you are using CRDs, you probably also use code-generation. The following block for Gopkg.toml will add the required code-generation packages to your project:
```
required = [  
  &#34;k8s.io/code-generator/cmd/client-gen&#34;,  
  &#34;k8s.io/code-generator/cmd/conversion-gen&#34;,  
  &#34;k8s.io/code-generator/cmd/deepcopy-gen&#34;,  
  &#34;k8s.io/code-generator/cmd/defaulter-gen&#34;,  
  &#34;k8s.io/code-generator/cmd/informer-gen&#34;,  
  &#34;k8s.io/code-generator/cmd/lister-gen&#34;,  
]


[[constraint]]

  branch = &#34;kubernetes-1.9.0&#34;


  name = &#34;k8s.io/code-generator&#34;

```Whether you would also like to prune unneeded packages (such as test files) through dep or commit the changes into the VCS at this point is up to you -- but from an upgrade perspective, you should now be ready to harness all the fancy new features that Kubernetes 1.9 brings through client-go.


	

	


