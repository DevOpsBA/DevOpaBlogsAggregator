|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/04/11/migrating-the-kubernetes-blog/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-04-11 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.8895555 &#43;0300 MSK m=&#43;2.478592201  |

# Migrating the Kubernetes Blog | Kubernetes

	
	
	
	
	We recently migrated the Kubernetes Blog from the Blogger platform to GitHub. With the change in platform comes a change in URL: formerly at [http://blog.kubernetes.io](http://blog.kubernetes.io), the blog now resides at [https://kubernetes.io/blog](https://kubernetes.io/blog).
All existing posts redirect from their former URLs with ```&lt;rel=canonical&gt;``` tags, preserving SEO values.
Our primary reasons for migrating were to streamline blog submissions and reviews, and to make the overall blog process faster and more transparent. Blogger&#39;s web interface made it difficult to provide drafts to multiple reviewers without also granting unnecessary access permissions and compromising security. GitHub&#39;s review process offered clear improvements.
We learned from [Jim Brikman](https://www.ybrikman.com)&#39;s experience during [his own site migration](https://www.ybrikman.com/writing/2015/04/20/migrating-from-blogger-to-github-pages/) away from Blogger.
Our migration was broken into several pull requests, but you can see the work that went into the [primary migration PR](https://github.com/kubernetes/website/pull/7247).
We hope that making blog submissions more accessible will encourage greater community involvement in creating and reviewing blog content.
You can submit a blog post for consideration one of two ways:
If you have a post that you want to remain confidential until your publish date, please submit your post via the Google form. Otherwise, you can choose your submission process based on your comfort level and preferred workflow.
The Kubernetes Blog needs more reviewers! If you&#39;re interested in contributing to the Kubernetes project and can participate on a regular, weekly basis, send an introductory email to [k8sblog@linuxfoundation.org](mailto:k8sblog@linuxfoundation.org).


	

	


