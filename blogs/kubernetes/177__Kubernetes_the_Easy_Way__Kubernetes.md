|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/11/Kubernetes-Easy-Way/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-11-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.2706086 &#43;0300 MSK m=&#43;2.859647501  |

#  Kubernetes the Easy Way  | Kubernetes

	
	
	
	
	***Editor&#39;s note: Today&#39;s post is by Dan Garfield, VP of Marketing at Codefresh, on how to set up and easily deploy a Kubernetes cluster.***
Kelsey Hightower wrote an invaluable guide for Kubernetes called [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way). It’s an awesome resource for those looking to understand the ins and outs of Kubernetes—but what if you want to put Kubernetes on easy mode? That’s something we’ve been working on together with Google Cloud. In this guide, we’ll show you how to get a cluster up and running, as well as how to actually deploy your code to that cluster and run it.
This is Kubernetes the easy way. 
We made Codefresh free for open-source projects and offer 200 builds/mo free for private projects, to make adopting Kubernetes as easy as possible. Deploy as much as you like on as many clusters as you like. 
**Note:** If you’re using a Cluster outside of Google Cloud, you can skip this step.
Google Container Engine is Google Cloud’s managed Kubernetes service. In our testing, it’s both powerful and easy to use.
If you’re new to the platform, you can get a $500 credit at the end of this process.

We’re done with step 1. In my experience it usually takes less than 5 minutes for a cluster to be created. 
First go to [Codefresh and create an account using GitHub, Bitbucket, or Gitlab](https://codefresh.io/kubernetes-deploy/). As mentioned previously, Codefresh is free for both open source and smaller private projects. We’ll use it to create the configuration Yaml necessary to deploy our application to Kubernetes. Then we&#39;ll deploy our application and automate the process to happen every time we commit code changes. Here are the steps:
To connect your Clusters in Google Container Engine, go to *Account Settings &gt; Integrations &gt; Kubernetes* and click **Authenticate**. This prompts you to login with your Google credentials.
Once you log in, all of your clusters are available within Codefresh.

To add your cluster, click the down arrow, and then click **add cluster**, select the project and cluster name. You can now deploy images!
To connect a non-GKE cluster we’ll need to add a token and certificate to Codefresh. Go to *Account Settings (bottom left) &gt; Integrations &gt; Kubernetes &gt; Configure &gt; Add Provider &gt; Custom Providers*. Expand the dropdown and click **Add Cluster**.

Follow the instructions on how to generate the needed information and click Save. Your cluster now appears under the Kubernetes tab. 
Now for the fun part! Codefresh provides an easily modifiable boilerplate that takes care of the heavy lifting of configuring Kubernetes for your application.
Think of namespaces as acting a bit like VLANs on a Kubernetes cluster. Each namespace can contain all the services that need to talk to each other on a Kubernetes cluster. For now, we’ll just work off the default namespace (the easy way!).
You can use the [demo application I mentioned earlier](https://github.com/containers101/demochat) that has a Node.js frontend with a MongoDB.

Here’s the info we need to pass:
**Cluster** - This is the cluster we added earlier, our application will be deployed there.**Namespace** - We’ll use default for our namespace but you can create and use a new one if you’d prefer. Namespaces are discrete units for grouping all the services associated with an application.**Service name** - You can name the service whatever you like. Since we’re deploying Mongo, I’ll just name it mongo!**Expose port** - We don’t need to expose the port outside of our cluster so we won’t check the box for now but we will specify a port where other containers can talk to this service. Mongo’s default port is ‘27017’.**Image** - Mongo is a public image on Dockerhub, so I can reference it by name and tag, ‘mongo:latest’.**Internal Ports** - This is the port the mongo application listens on, in this case it’s ‘27017’ again.
We can ignore the other options for now.

Boom! You’ve just deployed this image to Kubernetes. You can see by clicking on the status that the service, deployment, replicas, and pods are all configured and running. If you click Edit &gt; Advanced, you can see and edit all the raw YAML files associated with this application, or copy them and put them into your repository for use on any cluster. 
To get the rest of our demo application up and running we need to build and deploy the Node.js portion of the application. To do that we’ll need to add our repository to Codefresh.

We have the option to use a dockerfile, or to use a template if we need help creating a dockerfile. In this case, the demochat repo already has a dockerfile so we’ll select that. Click through the next few screens until the image builds.
Once the build is finished the image is automatically saved inside of the Codefresh docker registry. You can also add any [other registry to your account](https://docs.codefresh.io/v1.0/docs/docker-registry) and use that instead.
To deploy the image we’ll need
The pull secret is a token that the Kubernetes cluster can use to access a private Docker registry. To create one, we’ll need to generate the token and save it to Codefresh.

We’ll now be able to create our secret later on when we deploy our image.

We’re now ready to deploy the image we built.

Now let’s expose the port so we can access this application. This provisions an IP address and automatically configures ingress.

From this view you can scale the replicas, see application status, and similar tasks.

At this point you should have your entire application up and running! Not so bad huh? Now to automate deployment!
Every time we make a change to our application, we want to build a new image and deploy it to our cluster. We’ve already set up automated builds, but to automate deployment:


You can see the option to use a deployment file from your repo, or to use the deployment file that you just generated.
You’re done with deployment automation! Now whenever a change is made, the image will build, test, and deploy. 
We want to make it easy for every team, not just big enterprise teams, to adopt Kubernetes while preserving all of Kubernetes’ power and flexibility. At any point on the Kubernetes service screen you can switch to YAML to view all of the YAMLfiles generated by the configuration you performed in this walkthrough. You can tweak the file content, copy and paste them into local files, etc.
This walkthrough gives everyone a solid base to start with. When you’re ready, you can tweak the entities directly to specify the exact configuration you’d like.
We’d love your feedback! Please share with us on [Twitter](https://twitter.com/codefresh), or [reach out directly](https://codefresh.io/contact-us/).
**Do you have a video to walk me through this?** [You bet](https://www.youtube.com/watch?v=oFwFuUxxFdI&amp;list=PL8mgsmlx4BWV_j_L5oq-q8JdPnlJc3bUv).
**Does this work with Helm Charts?** Yes! We’re currently piloting Helm Charts with a limited set of users. Ping us if you’d like to try it early.
**Does this work with any Kubernetes cluster?** It should work with any Kubernetes cluster and is tested for Kubernetes 1.5 forward.
**Can I deploy Codefresh in my own data center?** Sure, Codefresh is built on top of Kubernetes using Helm Charts. Codefresh cloud is free for open source, and 200 builds/mo. Codefresh on prem is currently for enterprise users only.
**Won’t the database be wiped every time we update?** Yes, in this case we skipped creating a persistent volume. It’s a bit more work to get the persistent volume configured, if you’d like, [feel free to reach out](https://codefresh.io/contact-us/) and we’re happy to help!


	

	


