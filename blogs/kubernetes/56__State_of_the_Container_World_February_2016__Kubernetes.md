|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/03/State-Of-Container-World-February-2016/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-03-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:22.9681022 &#43;0300 MSK m=&#43;4.557150701  |

#  State of the Container World, February 2016  | Kubernetes

	
	
	
	
	Hello, and welcome to the second installment of the Kubernetes state of the container world survey. At the beginning of February we sent out a survey about people’s usage of containers, and wrote about the [results from the January survey](https://kubernetes.io/blog/2016/02/state-of-container-world-january-2016). Here we are again, as before, while we try to reach a large and representative set of respondents, this survey was publicized across the social media account of myself and others on the Kubernetes team, so I expect some pro-container and Kubernetes bias in the data.We continue to try to get as large an audience as possible, and in that vein, please go and take the [March survey](https://docs.google.com/a/google.com/forms/d/1hlOEyjuN4roIbcAAUbDhs7xjNMoM8r-hqtixf6zUsp4/viewform) and share it with your friends and followers everywhere! Without further ado, the numbers...
In January, 71% of respondents were currently using containers, in February, 89% of respondents were currently using containers. The percentage of users not even considering containers also shrank from 4% in January to a surprising 0% in February. Will see if that holds consistent in March.Likewise, the usage of containers continued to march across the dev/canary/prod lifecycle. In all parts of the lifecycle, container usage increased:
What is striking in this is that pre-production growth continued, even as workloads were clearly transitioned into true production. Likewise the share of people considering containers for production rose from 78% in January to 82% in February. Again we’ll see if the trend continues into March.
We asked some new questions in the survey too, around container and cluster sizes, and there were some interesting numbers:
How many containers are you running?

How many machines are you running containers on?

So while container usage continues to grow, the size and scope continues to be quite modest, with more than 50% of users running fewer than 50 containers on fewer than 10 machines.
Across the orchestration space, things seemed pretty consistent between January and February (Kubernetes is quite popular with folks (54% -&gt; 57%), though again, please see the note at the top about the likely bias in our survey population. Shell scripts likewise are also quite popular and holding steady. You all certainly love your Bash (don’t worry, we do too ;)
Likewise people continue to use cloud services both in raw IaaS form (10% on GCE, 30% on EC2, 2% on Azure) as well as cloud container services (16% for GKE, 11% on ECS, 1% on ACS). Though the most popular deployment target by far remains your laptop/desktop at ~53%.
As always, the complete raw data is available in a spreadsheet [here](https://docs.google.com/spreadsheets/d/126nnv9Q9avxDvC82irJGUDK3UODokILZOQe5X_WB9VQ/edit?usp=sharing).
Containers continue to gain in popularity and usage. The world of orchestration is somewhat stabilizing, and cloud services continue to be a common place to run containers, though your laptop is even more popular.
And if you are just getting started with containers (or looking to move beyond your laptop) please visit us at [kubernetes.io](http://kubernetes.io/) and [Google Container Engine](https://cloud.google.com/container-engine/). ‘Till next month, please get your friends, relatives and co-workers to take our [March survey](https://docs.google.com/a/google.com/forms/d/1hlOEyjuN4roIbcAAUbDhs7xjNMoM8r-hqtixf6zUsp4/viewform)!
Thanks!
*-- Brendan Burns, Software Engineer, Google*


	

	


