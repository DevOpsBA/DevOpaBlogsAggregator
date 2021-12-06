|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/07/Strong-Simple-Ssl-For-Kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-07-14 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.3141189 &#43;0300 MSK m=&#43;4.903169401  |

#  Strong, Simple SSL for Kubernetes Services  | Kubernetes

	
	
	
	
	Hi, I’m Evan Brown [(@evandbrown](http://twitter.com/evandbrown)) and I work on the solutions architecture team for Google Cloud Platform. I recently wrote an [article](https://cloud.google.com/solutions/automated-build-images-with-jenkins-kubernetes) and [tutorial](https://github.com/GoogleCloudPlatform/kube-jenkins-imager) about using Jenkins on Kubernetes to automate the Docker and GCE image build process. Today I’m going to discuss how I used Kubernetes services and secrets to add SSL to the Jenkins web UI. After reading this, you’ll be able to add SSL termination (and HTTP-&gt;HTTPS redirects &#43; basic auth) to your public HTTP Kubernetes services.
In the spirit of minimum viability, the first version of Jenkins-on-Kubernetes I built was very basic but functional:
Here’s a visual of that first version:
[img](https://1.bp.blogspot.com/-ccmpTmulrng/VaVxOs7gysI/AAAAAAAAAU8/bCEzgGGm-pE/s1600/0.png)
This works, but I have a few problems with it. First, authentication isn’t configured in a default Jenkins installation. The leader is sitting on the public Internet, accessible to anyone, until you connect and configure authentication. And since there’s no encryption, configuring authentication is kind of a symbolic gesture. We need SSL, and we need it now!
For a few milliseconds I considered trying to get SSL working directly on Jenkins. I’d never done it before, and I caught myself wondering if it would be as straightforward as working with SSL on [Nginx](http://nginx.org/), something I do have experience with. I’m all for learning new things, but this seemed like a great place to not invent a new wheel: SSL on Nginx is straightforward and well documented (as are its reverse-proxy capabilities), and Kubernetes is all about building functionality by orchestrating and composing containers. Let’s use Nginx, and add a few bonus features that Nginx makes simple: HTTP-&gt;HTTPS redirection, and basic access authentication.
I started by putting together a [Dockerfile](https://github.com/GoogleCloudPlatform/nginx-ssl-proxy/blob/master/Dockerfile) that inherited from the standard nginx image, copied a few Nginx config files, and added a custom entrypoint (start.sh). The entrypoint script checks an environment variable (ENABLE_SSL) and activates the correct Nginx config accordingly (meaning that unencrypted HTTP reverse proxy is possible, but that defeats the purpose). The script also configures basic access authentication if it’s enabled (the ENABLE_BASIC_AUTH env var).
Finally, start.sh evaluates the SERVICE_HOST_ENV_NAME and SERVICE_PORT_ENV_NAME env vars. These variables should be set to the names of the environment variables for the Kubernetes service you want to proxy to. In this example, the service for our Jenkins leader is cleverly named jenkins, which means pods in the cluster will see an environment variable named JENKINS_SERVICE_HOST and JENKINS_SERVICE_PORT_UI (the port that 8080 is mapped to on the Jenkins leader). SERVICE_HOST_ENV_NAME and SERVICE_PORT_ENV_NAME simply reference the correct service to use for a particular scenario, allowing the image to be used generically across deployments.
LIke every other pod in this example, we’ll deploy Nginx with a replication controller, allowing us to scale out or in, and recover automatically from container failures. This excerpt from a[complete descriptor in the sample app](https://github.com/GoogleCloudPlatform/kube-jenkins-imager/blob/master/ssl_proxy.yaml#L20-L48) shows some relevant bits of the pod spec:
```  spec:

    containers:

      -

        name: &#34;nginx-ssl-proxy&#34;

        image: &#34;gcr.io/cloud-solutions-images/nginx-ssl-proxy:latest&#34;

        env:

          -

            name: &#34;SERVICE\_HOST\_ENV\_NAME&#34;

            value: &#34;JENKINS\_SERVICE\_HOST&#34;

          -

            name: &#34;SERVICE\_PORT\_ENV\_NAME&#34;

            value: &#34;JENKINS\_SERVICE\_PORT\_UI&#34;

          -

            name: &#34;ENABLE\_SSL&#34;

            value: &#34;true&#34;

          -

            name: &#34;ENABLE\_BASIC\_AUTH&#34;

            value: &#34;true&#34;

        ports:

          -

            name: &#34;nginx-ssl-proxy-http&#34;

            containerPort: 80

          -

            name: &#34;nginx-ssl-proxy-https&#34;

            containerPort: 443
```The pod will have a service exposing TCP 80 and 443 to a public load balancer. Here’s the service descriptor [(also available in the sample app](https://github.com/GoogleCloudPlatform/kube-jenkins-imager/blob/master/service_ssl_proxy.yaml)):
```  kind: &#34;Service&#34;

  apiVersion: &#34;v1&#34;

  metadata:

    name: &#34;nginx-ssl-proxy&#34;

    labels:

      name: &#34;nginx&#34;

      role: &#34;ssl-proxy&#34;

  spec:

    ports:

      -

        name: &#34;https&#34;

        port: 443

        targetPort: &#34;nginx-ssl-proxy-https&#34;

        protocol: &#34;TCP&#34;

      -

        name: &#34;http&#34;

        port: 80

        targetPort: &#34;nginx-ssl-proxy-http&#34;

        protocol: &#34;TCP&#34;

    selector:

      name: &#34;nginx&#34;

      role: &#34;ssl-proxy&#34;

    type: &#34;LoadBalancer&#34;
```And here’s an overview with the SSL termination proxy in place. Notice that Jenkins is no longer directly exposed to the public Internet:
[img](https://3.bp.blogspot.com/-0B1BEQo_fWc/VaVxVUBkf3I/AAAAAAAAAVE/5yCCnA29C88/s1600/0%2B%25281%2529.png)
Now, how did the Nginx pods get ahold of the super-secret SSL key/cert and htpasswd file (for basic access auth)?
Kubernetes has an [API and resource for Secrets](https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/secrets.md). Secrets “are intended to hold sensitive information, such as passwords, OAuth tokens, and ssh keys. Putting this information in a secret is safer and more flexible than putting it verbatim in a pod definition or in a docker image.”
You can create secrets in your cluster in 3 simple steps:
Base64-encode your secret data (i.e., SSL key pair or htpasswd file)
```$ cat ssl.key | base64  
   LS0tLS1CRUdJTiBDRVJUS...
```Create a json document describing your secret, and add the base64-encoded values:
```  apiVersion: &#34;v1&#34;

  kind: &#34;Secret&#34;

  metadata:

    name: &#34;ssl-proxy-secret&#34;

    namespace: &#34;default&#34;

  data:

    proxycert: &#34;LS0tLS1CRUd...&#34;

    proxykey: &#34;LS0tLS1CR...&#34;

    htpasswd: &#34;ZXZhb...&#34;
```Create the secrets resource:
```$ kubectl create -f secrets.json
```To access the secrets from a container, specify them as a volume mount in your pod spec. Here’s the relevant excerpt from the [Nginx proxy template](https://github.com/GoogleCloudPlatform/kube-jenkins-imager/blob/master/ssl_proxy.yaml###L41-L48) we saw earlier:
```  spec:

    containers:

      -

        name: &#34;nginx-ssl-proxy&#34;

        image: &#34;gcr.io/cloud-solutions-images/nginx-ssl-proxy:latest&#34;

        env: [...]

        ports: ...[]

        volumeMounts:

          -

            name: &#34;secrets&#34;

            mountPath: &#34;/etc/secrets&#34;

            readOnly: true

    volumes:

      -

        name: &#34;secrets&#34;

        secret:

          secretName: &#34;ssl-proxy-secret&#34;
```A volume of type secret that points to the ssl-proxy-secret secret resource is defined, and then mounted into /etc/secrets in the container. The secrets spec in the earlier example defined data.proxycert, data.proxykey, and data.htpasswd, so we would see those files appear (base64-decoded) in /etc/secrets/proxycert, /etc/secrets/proxykey, and /etc/secrets/htpasswd for the Nginx process to access.
All together now
I have “containers and Kubernetes are fun and cool!” moments all the time, like probably every day. I’m beginning to have “containers and Kubernetes are extremely useful and powerful and are adding value to what I do by helping me do important things with ease” more frequently. This SSL termination proxy with Nginx example is definitely one of the latter. I didn’t have to waste time learning a new way to use SSL. I was able to solve my problem using well-known tools, in a reusable way, and quickly (from idea to working took about 2 hours).
Check out the complete [Automated Image Builds with Jenkins, Packer, and Kubernetes](https://github.com/GoogleCloudPlatform/kube-jenkins-imager) repo to see how the SSL termination proxy is used in a real cluster, or dig into the details of the proxy image in the [nginx-ssl-proxy repo](https://github.com/GoogleCloudPlatform/nginx-ssl-proxy) (complete with a Dockerfile and Packer template so you can build the image yourself).


	

	


