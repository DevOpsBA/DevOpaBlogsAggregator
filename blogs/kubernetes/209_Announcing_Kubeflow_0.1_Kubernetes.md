|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/05/04/announcing-kubeflow-0.1/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-05-04 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.8013845 &#43;0300 MSK m=&#43;2.390420701  |

# Announcing Kubeflow 0.1 | Kubernetes

	
	
	
	
	Since the [initial announcement](https://kubernetes.io/blog/2017/12/introducing-kubeflow-composable) of Kubeflow at [the last KubeCon&#43;CloudNativeCon](https://kccncna17.sched.com/event/CU5v/hot-dogs-or-not-at-scale-with-kubernetes-i-vish-kannan-david-aronchick-google), we have been both surprised and delighted by the excitement for building great ML stacks for Kubernetes. In just over five months, the [Kubeflow project](https://github.com/kubeflow) now has:
and already is among the top 2% of GitHub projects ***ever***.
People are excited to chat about Kubeflow as well! The Kubeflow community has also held meetups, talks and public sessions all around the world with thousands of attendees. With all this help, we’ve started to make substantial in every step of ML, from building your first model all the way to building a production-ready, high-scale deployments. At the end of the day, our mission remains the same: we want to let data scientists and software engineers focus on the things they do well by giving them an easy-to-use, portable and scalable ML stack.
Today, we’re proud to announce the availability of Kubeflow 0.1, which provides a minimal set of packages to begin developing, training and deploying ML. In just a few commands, you can get:
To get started, it’s just as easy as it always has been:
```# Create a namespace for kubeflow deployment
NAMESPACE=kubeflow
kubectl create namespace ${NAMESPACE}
VERSION=v0.1.3

# Initialize a ksonnet app. Set the namespace for it&#39;s default environment.
APP_NAME=my-kubeflow
ks init ${APP_NAME}
cd ${APP_NAME}
ks env set default --namespace ${NAMESPACE}

# Install Kubeflow components
ks registry add kubeflow github.com/kubeflow/kubeflow/tree/${VERSION}/kubeflow
ks pkg install kubeflow/core@${VERSION}
ks pkg install kubeflow/tf-serving@${VERSION}
ks pkg install kubeflow/tf-job@${VERSION}

# Create templates for core components
ks generate kubeflow-core kubeflow-core

# Deploy Kubeflow
ks apply default -c kubeflow-core
```And thats it! JupyterHub is deployed so we can now use Jupyter to begin developing models. Once we have python code to build our model we can build a docker image and train our model using our TFJob operator by running commands like the following:
```ks generate tf-job my-tf-job --name=my-tf-job --image=gcr.io/my/image:latest
ks apply default -c my-tf-job

We could then deploy the model by doing

ks generate tf-serving ${MODEL_COMPONENT} --name=${MODEL_NAME}
ks param set ${MODEL_COMPONENT} modelPath ${MODEL_PATH}
ks apply ${ENV} -c ${MODEL_COMPONENT}
```Within just a few commands, data scientists and software engineers can now create even complicated ML solutions and focus on what they do best: answering business critical questions.
It’d be impossible to have gotten where we are without enormous help from everyone in the community. Some specific contributions that we want to highlight include:
It’s difficult to overstate how much the community has helped bring all these projects (and more) to fruition. Just a few of the contributing companies include: Alibaba Cloud, Ant Financial, Caicloud, Canonical, Cisco, Datawire, Dell, GitHub, Google, Heptio, Huawei, Intel, Microsoft, Momenta, One Convergence, Pachyderm, Project Jupyter, Red Hat, Seldon, Uber and Weaveworks.
If you’d like to try out Kubeflow, we have a number of options for you:
There were also a number of sessions at KubeCon &#43; CloudNativeCon  EU 2018 covering Kubeflow. The links to the talks are here; the associated videos will be posted in the coming days.
Our next major release will be 0.2 coming this summer. In it, we expect to land the following new features:
But the most important feature is the one we haven’t heard yet. Please tell us! Some options for making your voice heard include:
Thank you for all your support so far!
*Jeremy Lewi &amp; David Aronchick* Google


	

	


