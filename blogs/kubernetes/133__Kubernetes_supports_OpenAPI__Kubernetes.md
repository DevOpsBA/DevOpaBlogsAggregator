|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/12/Kubernetes-Supports-Openapi/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-12-23 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.851784 &#43;0300 MSK m=&#43;3.440826201  |

#  Kubernetes supports OpenAPI  | Kubernetes

	
	
	
	
	*Editor’s note: this post is part of a [series of in-depth articles](https://kubernetes.io/blog/2016/12/five-days-of-kubernetes-1-5/) on what&#39;s new in Kubernetes 1.5*
[OpenAPI](https://www.openapis.org/) allows API providers to define their operations and models, and enables developers to automate their tools and generate their favorite language’s client to talk to that API server. Kubernetes has supported swagger 1.2 (older version of OpenAPI spec) for a while, but the spec was incomplete and invalid, making it hard to generate tools/clients based on it.
In Kubernetes 1.4, we introduced alpha support for the OpenAPI spec (formerly known as swagger 2.0 before it was donated to the [Open API Initiative](https://www.openapis.org/about)) by upgrading the current models and operations. Beginning in [Kubernetes 1.5](https://kubernetes.io/blog/2016/12/kubernetes-1-5-supporting-production-workloads/), the support for the OpenAPI spec has been completed by auto-generating the spec directly from Kubernetes source, which will keep the spec--and documentation--completely in sync with future changes in operations/models.
The new spec enables us to have better API documentation and we have even introduced a supported [python client](https://github.com/kubernetes-incubator/client-python).
The spec is modular, divided by GroupVersion: this is future-proof, since we intend to allow separate GroupVersions to be served out of separate API servers.
The structure of spec is explained in detail in [OpenAPI spec definition](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md). We used [operation’s tags](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#tag-object) to separate each GroupVersion and filled as much information as we can about paths/operations and models. For a specific operation, all parameters, method of call, and responses are documented.
For example, OpenAPI spec for reading a pod information is:
```{

...  
  &#34;paths&#34;: {

&#34;/api/v1/namespaces/{namespace}/pods/{name}&#34;: {  
    &#34;get&#34;: {  
     &#34;description&#34;: &#34;read the specified Pod&#34;,  
     &#34;consumes&#34;: [  
      &#34;\*/\*&#34;  
     ],  
     &#34;produces&#34;: [  
      &#34;application/json&#34;,  
      &#34;application/yaml&#34;,  
      &#34;application/vnd.kubernetes.protobuf&#34;  
     ],  
     &#34;schemes&#34;: [  
      &#34;https&#34;  
     ],  
     &#34;tags&#34;: [  
      &#34;core\_v1&#34;  
     ],  
     &#34;operationId&#34;: &#34;readCoreV1NamespacedPod&#34;,  
     &#34;parameters&#34;: [  
      {  
       &#34;uniqueItems&#34;: true,  
       &#34;type&#34;: &#34;boolean&#34;,  
       &#34;description&#34;: &#34;Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;.&#34;,  
       &#34;name&#34;: &#34;exact&#34;,  
       &#34;in&#34;: &#34;query&#34;  
      },  
      {  
       &#34;uniqueItems&#34;: true,  
       &#34;type&#34;: &#34;boolean&#34;,  
       &#34;description&#34;: &#34;Should this value be exported.  Export strips fields that a user can not specify.&#34;,  
       &#34;name&#34;: &#34;export&#34;,  
       &#34;in&#34;: &#34;query&#34;  
      }  
     ],  
     &#34;responses&#34;: {  
      &#34;200&#34;: {  
       &#34;description&#34;: &#34;OK&#34;,  
       &#34;schema&#34;: {  
        &#34;$ref&#34;: &#34;#/definitions/v1.Pod&#34;  
       }  
      },  
      &#34;401&#34;: {  
       &#34;description&#34;: &#34;Unauthorized&#34;  
      }  
     }  
    },

…

}

…
```Using this information and the URL of ```kube-apiserver```, one should be able to make the call to the given url (/api/v1/namespaces/{namespace}/pods/{name}) with parameters such as ```name```, ```exact```, ```export```, etc. to get pod’s information. Client libraries generators would also use this information to create an API function call for reading pod’s information. For example, [python client](https://github.com/kubernetes-incubator/client-python) makes it easy to call this operation like this:
```from kubernetes import client

ret = client.CoreV1Api().read\_namespaced\_pod(name=&#34;pods\_name&#34;, namespace=&#34;default&#34;)
```A simplified version of generated read_namespaced_pod, can be found [here](https://gist.github.com/mbohlool/d5ec1dace27ef90cf742555c05480146).
Swagger-codegen document generator would also be able to create documentation using the same information:
```GET /api/v1/namespaces/{namespace}/pods/{name}

(readCoreV1NamespacedPod)

read the specified Pod

Path parameters

name (required)

Path Parameter — name of the Pod

namespace (required)

Path Parameter — object name and auth scope, such as for teams and projects

Consumes

This API call consumes the following media types via the Content-Type request header:

-
\*/\*


Query parameters

pretty (optional)

Query Parameter — If &#39;true&#39;, then the output is pretty printed.

exact (optional)

Query Parameter — Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;.

export (optional)

Query Parameter — Should this value be exported. Export strips fields that a user can not specify.

Return type

v1.Pod


Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

-
application/json
-
application/yaml
-
application/vnd.kubernetes.protobuf

Responses

200

OK v1.Pod

401

Unauthorized
```There are two ways to access OpenAPI spec:
There are numerous [tools](http://swagger.io/tools/) that works with this spec. For example, you can use the [swagger editor](http://swagger.io/swagger-editor/) to open the spec file and render documentation, as well as generate clients; or you can directly use [swagger codegen](http://swagger.io/swagger-codegen/) to generate documentation and clients. The clients this generates will mostly work out of the box--but you will need some support for authorization and some Kubernetes specific utilities. Use [python client](https://github.com/kubernetes-incubator/client-python) as a template to create your own client.
If you want to get involved in development of OpenAPI support, client libraries, or report a bug, you can get in touch with developers at [SIG-API-Machinery](https://github.com/kubernetes/community/tree/master/sig-api-machinery).
*--Mehdy Bohlool, Software Engineer, Google*


	

	


