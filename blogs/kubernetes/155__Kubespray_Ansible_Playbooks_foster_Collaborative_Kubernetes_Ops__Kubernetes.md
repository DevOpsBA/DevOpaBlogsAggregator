|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/05/Kubespray-Ansible-Collaborative-Kubernetes-Ops/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-05-19 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.548934 &#43;0300 MSK m=&#43;3.137974401  |

#  Kubespray Ansible Playbooks foster Collaborative Kubernetes Ops  | Kubernetes

	
	
	
	
	*Today’s guest post is by Rob Hirschfeld, co-founder of open infrastructure automation project, Digital Rebar and co-chair of the SIG Cluster Ops.  *
**Why Kubespray?**
Making Kubernetes operationally strong is a widely held priority and I track many deployment efforts around the project. The [incubated Kubespray project](https://github.com/kubernetes-incubator/kubespray) is of particular interest for me because it uses the popular Ansible toolset to build robust, upgradable clusters on both cloud and physical targets. I believe using tools familiar to operators grows our community.
We’re excited to see the breadth of platforms enabled by Kubespray and how well it handles a wide range of options like integrating Ceph for [StatefulSet](/docs/concepts/workloads/controllers/statefulset/) persistence and Helm for easier application uploads. Those additions have allowed us to fully integrate the [OpenStack Helm charts](https://github.com/att-comdev/openstack-helm) ([demo video](https://www.youtube.com/watch?v=wZ0vMrdx4a4&amp;list=PLXPBeIrpXjfjabMbwYyDULOX3kZmlxEXK&amp;index=2)).
By working with the upstream source instead of creating different install scripts, we get the benefits of a larger community. This requires some extra development effort; however, we believe helping share operational practices makes the whole community stronger. That was also the motivation behind the [SIG-Cluster Ops](https://github.com/kubernetes/community/tree/master/sig-cluster-ops).
**With Kubespray delivering robust installs, we can focus on broader operational concerns.**
For example, we can now drive parallel deployments, so it’s possible to fully exercise the options enabled by Kubespray simultaneously for development and testing.  
That’s helpful to built-test-destroy coordinated Kubernetes installs on CentOS, Red Hat and Ubuntu as part of an automation pipeline. We can also set up a full classroom environment from a single command using [Digital Rebar’s](https://github.com/digitalrebar/digitalrebar) providers, tenants and cluster definition JSON.
**Let’s explore the classroom example:**
First, we define a [student cluster in JSON](https://github.com/digitalrebar/digitalrebar/blob/master/deploy/workloads/cluster/deploy-001.json) like the snippet below
|
{
 &#34;attribs&#34;: {
   &#34;k8s-version&#34;: &#34;v1.6.0&#34;,
   &#34;k8s-kube_network_plugin&#34;: &#34;calico&#34;,
   &#34;k8s-docker_version&#34;: &#34;1.12&#34;
 },
 &#34;name&#34;: &#34;cluster01&#34;,
 &#34;tenant&#34;: &#34;cluster01&#34;,
 &#34;public_keys&#34;: {
   &#34;cluster01&#34;: &#34;ssh-rsa AAAAB..... [user@example.com](mailto:user@example.com)&#34;
 },
 &#34;provider&#34;: {
   &#34;name&#34;: &#34;google-provider&#34;
 },
 &#34;nodes&#34;: [
   {
     &#34;roles&#34;: [&#34;etcd&#34;,&#34;k8s-addons&#34;, &#34;k8s-master&#34;],
     &#34;count&#34;: 1
   },
   {
     &#34;roles&#34;: [&#34;k8s-worker&#34;],
     &#34;count&#34;: 3
   }
 ]
}
|
Then we run the [Digital Rebar workloads Multideploy.sh](https://github.com/digitalrebar/digitalrebar/blob/master/deploy/workloads/multideploy.sh) reference script which inspects the deployment files to pull out key information.  Basically, it automates the following steps:
|
rebar provider create {“name”:“google-provider”, [secret stuff]}
rebar tenants create {“name”:“cluster01”}
rebar deployments create [contents from cluster01 file]
|
The deployments create command will automatically request nodes from the provider. Since we’re using tenants and SSH key additions, each student only gets access to their own cluster. When we’re done, adding the --destroy flag will reverse the process for the nodes and deployments but leave the providers and tenants.
**We are invested in operational scripts like this example using Kubespray and Digital Rebar because if we cannot manage variation in a consistent way then we’re doomed to operational fragmentation.  **
I am excited to see and be part of the community progress towards enterprise-ready Kubernetes operations on both cloud and on-premises. That means I am seeing reasonable patterns emerge with sharable/reusable automation. I strongly recommend watching (or better, collaborating in) these efforts if you are deploying Kubernetes even at experimental scale. Being part of the community requires more upfront effort but returns dividends as you get the benefits of shared experience and improvement.
**When deploying at scale, how do you set up a system to be both repeatable and multi-platform without compromising scale or security?**
With Kubespray and Digital Rebar as a repeatable base, extensions get much faster and easier. Even better, using upstream directly allows improvements to be quickly cycled back into upstream. That means we’re closer to building a community focused on the operational side of Kubernetes with an [SRE mindset](https://rackn.com/sre).
If this is interesting, please engage with us in the [Cluster Ops SIG](https://github.com/kubernetes/community/tree/master/sig-cluster-ops), [Kubespray](https://github.com/kubernetes-incubator/kubespray) or [Digital Rebar](http://rebar.digital/) communities. 
*-- Rob Hirschfeld, co-founder of RackN and co-chair of the Cluster Ops SIG*


	

	


