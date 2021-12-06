|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/05/Kubernetes-Release-0170/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-05-15 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.5014125 &#43;0300 MSK m=&#43;5.090464101  |

#  Kubernetes Release: 0.17.0  | Kubernetes

	
	
	
	
	Release Notes:
To download, please visit [https://github.com/GoogleCloudPlatform/kubernetes/releases/tag/v0.17.0](https://github.com/GoogleCloudPlatform/kubernetes/releases/tag/v0.17.0)
Simple theme. Powered by [Blogger][385].
[ ![][327] ][386]
[146]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7873](https://github.com/GoogleCloudPlatform/kubernetes/pull/7873) &#34;Fix bug in Service documentation: incorrect location of &#34;selector&#34; in JSON&#34;
[147]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7876](https://github.com/GoogleCloudPlatform/kubernetes/pull/7876) &#34;Fix controller-manager manifest for providers that don&#39;t specify CLUSTER_IP_RANGE&#34;
[148]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7867](https://github.com/GoogleCloudPlatform/kubernetes/pull/7867) &#34;Fix controller unittests&#34;
[149]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7751](https://github.com/GoogleCloudPlatform/kubernetes/pull/7751) &#34;Enable GCM and GCL instead of InfluxDB on GCE&#34;
[150]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7862](https://github.com/GoogleCloudPlatform/kubernetes/pull/7862) &#34;Remove restriction that cluster-cidr be a class-b&#34;
[151]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7591](https://github.com/GoogleCloudPlatform/kubernetes/pull/7591) &#34;Fix OpenShift example&#34;
[152]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7718](https://github.com/GoogleCloudPlatform/kubernetes/pull/7718) &#34;API Server - pass path name in context of create request for subresource&#34;
[153]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7575](https://github.com/GoogleCloudPlatform/kubernetes/pull/7575) &#34;Rolling Updates: Add support for --rollback.&#34;
[154]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7820](https://github.com/GoogleCloudPlatform/kubernetes/pull/7820) &#34;Update to container-vm-v20150505 (Also updates GCE to Docker 1.6)&#34;
[155]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7830](https://github.com/GoogleCloudPlatform/kubernetes/pull/7830) &#34;Fix metric label&#34;
[156]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7838](https://github.com/GoogleCloudPlatform/kubernetes/pull/7838) &#34;Fix v1beta1 typos in v1beta2 conversions&#34;
[157]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7764](https://github.com/GoogleCloudPlatform/kubernetes/pull/7764) &#34;skydns: use the etcd-2.x native syntax, enable IANA attributed ports.&#34;
[158]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7794](https://github.com/GoogleCloudPlatform/kubernetes/pull/7794) &#34;Added port 6443 to kube-proxy default IP address for api-server&#34;
[159]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7834](https://github.com/GoogleCloudPlatform/kubernetes/pull/7834) &#34;Added client header info for authentication doc.&#34;
[160]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7827](https://github.com/GoogleCloudPlatform/kubernetes/pull/7827) &#34;Clean up safe_format_and_mount spam in the startup logs&#34;
[161]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7829](https://github.com/GoogleCloudPlatform/kubernetes/pull/7829) &#34;Set allocate_node_cidrs to be blank by default.&#34;
[162]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/5246](https://github.com/GoogleCloudPlatform/kubernetes/pull/5246) &#34;Make nodecontroller configure nodes&#39; pod IP ranges&#34;
[163]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7799](https://github.com/GoogleCloudPlatform/kubernetes/pull/7799) &#34;Fix sync problems in #5246&#34;
[164]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7823](https://github.com/GoogleCloudPlatform/kubernetes/pull/7823) &#34;Fix event doc link&#34;
[165]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7776](https://github.com/GoogleCloudPlatform/kubernetes/pull/7776) &#34;Cobra update and bash completions fix&#34;
[166]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7675](https://github.com/GoogleCloudPlatform/kubernetes/pull/7675) &#34;Fix kube2sky flakes. Fix tools.GetEtcdVersion to work with etcd &gt; 2.0.7&#34;
[167]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7154](https://github.com/GoogleCloudPlatform/kubernetes/pull/7154) &#34;Change kube2sky to use token-system-dns secret, point at https endpoint ...&#34;
[168]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7468](https://github.com/GoogleCloudPlatform/kubernetes/pull/7468) &#34;replica: serialize created-by reference&#34;
[169]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7702](https://github.com/GoogleCloudPlatform/kubernetes/pull/7702) &#34;Inject mounter into volume plugins&#34;
[170]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/6973](https://github.com/GoogleCloudPlatform/kubernetes/pull/6973) &#34;bringing CoreOS cloud-configs up-to-date (against 0.15.x and latest OS&#39; alpha) &#34;
[171]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7787](https://github.com/GoogleCloudPlatform/kubernetes/pull/7787) &#34;Update kubeconfig-file doc.&#34;
[172]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7780](https://github.com/GoogleCloudPlatform/kubernetes/pull/7780) &#34;Throw an API error when deleting namespace in termination&#34;
[173]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7773](https://github.com/GoogleCloudPlatform/kubernetes/pull/7773) &#34;Fix command field PodExecOptions&#34;
[174]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7785](https://github.com/GoogleCloudPlatform/kubernetes/pull/7785) &#34;Start ImageManager housekeeping in Run().&#34;
[175]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7769](https://github.com/GoogleCloudPlatform/kubernetes/pull/7769) &#34;fix DeepCopy to properly support runtime.EmbeddedObject&#34;
[176]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7273](https://github.com/GoogleCloudPlatform/kubernetes/pull/7273) &#34;fix master service endpoint system for multiple masters&#34;
[177]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7757](https://github.com/GoogleCloudPlatform/kubernetes/pull/7757) &#34;Add genbashcomp to KUBE_TEST_TARGETS&#34;
[178]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7669](https://github.com/GoogleCloudPlatform/kubernetes/pull/7669) &#34;Change the cloud provider TCPLoadBalancerExists function to GetTCPLoadBalancer...&#34;
[179]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7772](https://github.com/GoogleCloudPlatform/kubernetes/pull/7772) &#34;Add containerized option to kubelet binary&#34;
[180]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7779](https://github.com/GoogleCloudPlatform/kubernetes/pull/7779) &#34;Fix swagger spec&#34;
[181]: [https://github.com/GoogleCloudPlatform/kubernetes/issues/7750](https://github.com/GoogleCloudPlatform/kubernetes/issues/7750) &#34;Hyperkube image requires root certificates to work with cloud-providers (at least AWS)&#34;
[182]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7755](https://github.com/GoogleCloudPlatform/kubernetes/pull/7755) &#34;FIX: Issue #7750 - Hyperkube docker image needs certificates to connect to cloud-providers&#34;
[183]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7752](https://github.com/GoogleCloudPlatform/kubernetes/pull/7752) &#34;Add build labels to rkt&#34;
[184]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7672](https://github.com/GoogleCloudPlatform/kubernetes/pull/7672) &#34;Check license boilerplate for python files&#34;
[185]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7705](https://github.com/GoogleCloudPlatform/kubernetes/pull/7705) &#34;Reliable updates in rollingupdate&#34;
[186]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7650](https://github.com/GoogleCloudPlatform/kubernetes/pull/7650) &#34;Don&#39;t exit abruptly if there aren&#39;t yet any minions right after the cluster is created.&#34;
[187]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7742](https://github.com/GoogleCloudPlatform/kubernetes/pull/7742) &#34;Make changes suggested in #7675&#34;
[188]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7357](https://github.com/GoogleCloudPlatform/kubernetes/pull/7357) &#34;A guide to set up kubernetes multiple nodes cluster with flannel on fedora&#34;
[189]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7760](https://github.com/GoogleCloudPlatform/kubernetes/pull/7760) &#34;Setup generators in factory&#34;
[190]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7737](https://github.com/GoogleCloudPlatform/kubernetes/pull/7737) &#34;Reduce usage of time.After&#34;
[191]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7735](https://github.com/GoogleCloudPlatform/kubernetes/pull/7735) &#34;Remove node status from &#34;componentstatuses&#34; call.&#34;
[192]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7614](https://github.com/GoogleCloudPlatform/kubernetes/pull/7614) &#34;React to failure by growing the remaining clusters&#34;
[193]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7725](https://github.com/GoogleCloudPlatform/kubernetes/pull/7725) &#34;Fix typo in runtime_cache.go&#34;
[194]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7740](https://github.com/GoogleCloudPlatform/kubernetes/pull/7740) &#34;Update non-GCE Salt distros to 1.6.0, fallback to ContainerVM Docker version on GCE&#34;
[195]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7744](https://github.com/GoogleCloudPlatform/kubernetes/pull/7744) &#34;Skip SaltStack install if it&#39;s already installed&#34;
[196]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7712](https://github.com/GoogleCloudPlatform/kubernetes/pull/7712) &#34;Expose pod name as a label on containers.&#34;
[197]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7732](https://github.com/GoogleCloudPlatform/kubernetes/pull/7732) &#34;Log which SSH key is used in e2e SSH test&#34;
[198]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7649](https://github.com/GoogleCloudPlatform/kubernetes/pull/7649) &#34;Add a central simple getting started guide with kubernetes guide.&#34;
[199]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7443](https://github.com/GoogleCloudPlatform/kubernetes/pull/7443) &#34;Explicitly state the lack of support for &#39;Requests&#39; for the purposes of scheduling&#34;
[200]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7721](https://github.com/GoogleCloudPlatform/kubernetes/pull/7721) &#34;Select IPv4-only from host interfaces&#34;
[201]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7723](https://github.com/GoogleCloudPlatform/kubernetes/pull/7723) &#34;Metrics tests can&#39;t run on Mac&#34;
[202]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7727](https://github.com/GoogleCloudPlatform/kubernetes/pull/7727) &#34;Add step to API changes doc for swagger regen&#34;
[203]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7703](https://github.com/GoogleCloudPlatform/kubernetes/pull/7703) &#34;Add NsenterMounter mount implementation&#34;
[204]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7509](https://github.com/GoogleCloudPlatform/kubernetes/pull/7509) &#34;add StringSet.HasAny&#34;
[205]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/6941](https://github.com/GoogleCloudPlatform/kubernetes/pull/6941) &#34;Add an integration test that checks for the metrics we expect to be exported from the master&#34;
[206]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7722](https://github.com/GoogleCloudPlatform/kubernetes/pull/7722) &#34;Minor bash update found by shellcheck.net&#34;
[207]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7536](https://github.com/GoogleCloudPlatform/kubernetes/pull/7536) &#34;Add --hostport to run-container.&#34;
[208]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7659](https://github.com/GoogleCloudPlatform/kubernetes/pull/7659) &#34;Have rkt implement the container Runtime interface&#34;
[209]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7629](https://github.com/GoogleCloudPlatform/kubernetes/pull/7629) &#34;Change the order the different versions of API are registered &#34;
[210]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7699](https://github.com/GoogleCloudPlatform/kubernetes/pull/7699) &#34;expose: Create objects in a generic way&#34;
[211]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7643](https://github.com/GoogleCloudPlatform/kubernetes/pull/7643) &#34;Requeue rc if a single get/put retry on status.Replicas fails&#34;
[212]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7316](https://github.com/GoogleCloudPlatform/kubernetes/pull/7316) &#34;logs for master components&#34;
[213]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7522](https://github.com/GoogleCloudPlatform/kubernetes/pull/7522) &#34;cloudproviders: add ovirt getting started guide&#34;
[214]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7671](https://github.com/GoogleCloudPlatform/kubernetes/pull/7671) &#34;Make rkt-install a oneshot.&#34;
[215]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7665](https://github.com/GoogleCloudPlatform/kubernetes/pull/7665) &#34;Provide container_runtime flag to Kubelet in CoreOS.&#34;
[216]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7654](https://github.com/GoogleCloudPlatform/kubernetes/pull/7654) &#34;Boilerplate speedup&#34;
[217]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7700](https://github.com/GoogleCloudPlatform/kubernetes/pull/7700) &#34;Log host for failed pod in Density test&#34;
[218]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7655](https://github.com/GoogleCloudPlatform/kubernetes/pull/7655) &#34;Removes spurious quotation mark&#34;
[219]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7694](https://github.com/GoogleCloudPlatform/kubernetes/pull/7694) &#34;Add kubectl_label to custom functions in bash completion&#34;
[220]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7696](https://github.com/GoogleCloudPlatform/kubernetes/pull/7696) &#34;Enable profiling in kube-controller&#34;
[221]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7690](https://github.com/GoogleCloudPlatform/kubernetes/pull/7690) &#34;Set vagrant test cluster default NUM_MINIONS=2&#34;
[222]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7695](https://github.com/GoogleCloudPlatform/kubernetes/pull/7695) &#34;Add metrics to measure cache hit ratio&#34;
[223]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7662](https://github.com/GoogleCloudPlatform/kubernetes/pull/7662) &#34;Change IP to IP(S) in service columns for kubectl get&#34;
[224]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7076](https://github.com/GoogleCloudPlatform/kubernetes/pull/7076) &#34;annotate required flags for bash_completions&#34;
[225]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7685](https://github.com/GoogleCloudPlatform/kubernetes/pull/7685) &#34;(minor) Add pgrep debugging to etcd error&#34;
[226]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7676](https://github.com/GoogleCloudPlatform/kubernetes/pull/7676) &#34;Fixed nil pointer issue in describe when volume is unbound&#34;
[227]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7691](https://github.com/GoogleCloudPlatform/kubernetes/pull/7691) &#34;Removed unnecessary closing bracket&#34;
[228]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7432](https://github.com/GoogleCloudPlatform/kubernetes/pull/7432) &#34;Added TerminationGracePeriod field to PodSpec and grace-period flag to kubectl stop&#34;
[229]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7689](https://github.com/GoogleCloudPlatform/kubernetes/pull/7689) &#34;Fix boilerplate in test/e2e/scale.go&#34;
[230]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7628](https://github.com/GoogleCloudPlatform/kubernetes/pull/7628) &#34;Update expiration timeout based on observed latencies&#34;
[231]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7644](https://github.com/GoogleCloudPlatform/kubernetes/pull/7644) &#34;Output generated conversion functions/names&#34;
[232]: [https://github.com/GoogleCloudPlatform/kubernetes/issues/7645](https://github.com/GoogleCloudPlatform/kubernetes/issues/7645) &#34;Move the scale tests into a separate file&#34;
[233]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7646](https://github.com/GoogleCloudPlatform/kubernetes/pull/7646) &#34;Moved the Scale tests into a scale file. #7645&#34;
[234]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7609](https://github.com/GoogleCloudPlatform/kubernetes/pull/7609) &#34;Truncate GCE load balancer names to 63 chars&#34;
[235]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7603](https://github.com/GoogleCloudPlatform/kubernetes/pull/7603) &#34;Add SyncPod() and remove Kill/Run InContainer().&#34;
[236]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7663](https://github.com/GoogleCloudPlatform/kubernetes/pull/7663) &#34;Merge release 0.16 to master&#34;
[237]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7637](https://github.com/GoogleCloudPlatform/kubernetes/pull/7637) &#34;Update license boilerplate for examples/rethinkdb&#34;
[238]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7268](https://github.com/GoogleCloudPlatform/kubernetes/pull/7268) &#34;First part of improved rolling update, allow dynamic next replication controller generation.&#34;
[239]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7638](https://github.com/GoogleCloudPlatform/kubernetes/pull/7638) &#34;Add license boilerplate to examples/phabricator&#34;
[240]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7597](https://github.com/GoogleCloudPlatform/kubernetes/pull/7597) &#34;Use generic copyright holder name in license boilerplate&#34;
[241]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7633](https://github.com/GoogleCloudPlatform/kubernetes/pull/7633) &#34;Retry incrementing quota if there is a conflict&#34;
[242]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7568](https://github.com/GoogleCloudPlatform/kubernetes/pull/7568) &#34;Remove GetContainers from Runtime interface&#34;
[243]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7578](https://github.com/GoogleCloudPlatform/kubernetes/pull/7578) &#34;Add image-related methods to DockerManager&#34;
[244]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7586](https://github.com/GoogleCloudPlatform/kubernetes/pull/7586) &#34;Remove more docker references in kubelet&#34;
[245]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7601](https://github.com/GoogleCloudPlatform/kubernetes/pull/7601) &#34;Add KillContainerInPod in DockerManager&#34;
[246]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7652](https://github.com/GoogleCloudPlatform/kubernetes/pull/7652) &#34;Kubelet: Add container runtime option.&#34;
[247]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7626](https://github.com/GoogleCloudPlatform/kubernetes/pull/7626) &#34;bump heapster to v0.11.0 and grafana to v0.7.0&#34;
[248]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7593](https://github.com/GoogleCloudPlatform/kubernetes/pull/7593) &#34;Build github.com/onsi/ginkgo/ginkgo as a part of the release&#34;
[249]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7490](https://github.com/GoogleCloudPlatform/kubernetes/pull/7490) &#34;Do not automatically decode runtime.RawExtension&#34;
[250]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7500](https://github.com/GoogleCloudPlatform/kubernetes/pull/7500) &#34;Update changelog.&#34;
[251]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7610](https://github.com/GoogleCloudPlatform/kubernetes/pull/7610) &#34;Add SyncPod() to DockerManager and use it in Kubelet&#34;
[252]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7602](https://github.com/GoogleCloudPlatform/kubernetes/pull/7602) &#34;Build: Push .md5 and .sha1 files for every file we push to GCS&#34;
[253]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7540](https://github.com/GoogleCloudPlatform/kubernetes/pull/7540) &#34;Fix rolling update --image &#34;
[254]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7636](https://github.com/GoogleCloudPlatform/kubernetes/pull/7636) &#34;Update license boilerplate for docs/man/md2man-all.sh&#34;
[255]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7632](https://github.com/GoogleCloudPlatform/kubernetes/pull/7632) &#34;Include shell license boilerplate in examples/k8petstore&#34;
[256]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7277](https://github.com/GoogleCloudPlatform/kubernetes/pull/7277) &#34;Add --cgroup_parent flag to Kubelet to set the parent cgroup for pods&#34;
[257]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7209](https://github.com/GoogleCloudPlatform/kubernetes/pull/7209) &#34;change the current dir to the config dir&#34;
[258]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7158](https://github.com/GoogleCloudPlatform/kubernetes/pull/7158) &#34;Set Weave To 0.9.0 And Update Etcd Configuration For Azure&#34;
[259]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7467](https://github.com/GoogleCloudPlatform/kubernetes/pull/7467) &#34;Augment describe to search for matching things if it doesn&#39;t match the original resource.&#34;
[260]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7559](https://github.com/GoogleCloudPlatform/kubernetes/pull/7559) &#34;Add a simple cache for objects stored in etcd.&#34;
[261]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7549](https://github.com/GoogleCloudPlatform/kubernetes/pull/7549) &#34;Rkt gc&#34;
[262]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7550](https://github.com/GoogleCloudPlatform/kubernetes/pull/7550) &#34;Rkt pull&#34;
[263]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/6400](https://github.com/GoogleCloudPlatform/kubernetes/pull/6400) &#34;Implement Mount interface using mount(8) and umount(8)&#34;
[264]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7588](https://github.com/GoogleCloudPlatform/kubernetes/pull/7588) &#34;Trim Fleuntd tag for Cloud Logging&#34;
[265]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7569](https://github.com/GoogleCloudPlatform/kubernetes/pull/7569) &#34;GCE CoreOS cluster - set master name based on variable&#34;
[266]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7535](https://github.com/GoogleCloudPlatform/kubernetes/pull/7535) &#34;Capitalization of KubeProxyVersion wrong in JSON&#34;
[267]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7530](https://github.com/GoogleCloudPlatform/kubernetes/pull/7530) &#34;Make nodes report their external IP rather than the master&#39;s.&#34;
[268]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7539](https://github.com/GoogleCloudPlatform/kubernetes/pull/7539) &#34;Trim cluster log tags to pod name and container name&#34;
[269]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7541](https://github.com/GoogleCloudPlatform/kubernetes/pull/7541) &#34;Handle conversion of boolean query parameters with a value of &#34;false&#34;&#34;
[270]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7532](https://github.com/GoogleCloudPlatform/kubernetes/pull/7532) &#34;Add image-related methods to Runtime interface.&#34;
[271]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7560](https://github.com/GoogleCloudPlatform/kubernetes/pull/7560) &#34;Test whether auto-generated conversions weren&#39;t manually edited&#34;
[272]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7484](https://github.com/GoogleCloudPlatform/kubernetes/pull/7484) &#34;Mention :latest behavior for image version tag&#34;
[273]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7487](https://github.com/GoogleCloudPlatform/kubernetes/pull/7487) &#34;readinessProbe calls livenessProbe.Exec.Command which cause &#34;invalid memory address or nil pointer dereference&#34;.&#34;
[274]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7520](https://github.com/GoogleCloudPlatform/kubernetes/pull/7520) &#34;Add RuntimeHooks to abstract Kubelet logic&#34;
[275]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7546](https://github.com/GoogleCloudPlatform/kubernetes/pull/7546) &#34;Expose URL() on Request to allow building URLs&#34;
[276]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7288](https://github.com/GoogleCloudPlatform/kubernetes/pull/7288) &#34;Add a simple cache for objects stored in etcd&#34;
[277]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7431](https://github.com/GoogleCloudPlatform/kubernetes/pull/7431) &#34;Prepare for chaining autogenerated conversion methods &#34;
[278]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7353](https://github.com/GoogleCloudPlatform/kubernetes/pull/7353) &#34;Increase maxIdleConnection limit when creating etcd client in apiserver.&#34;
[279]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7354](https://github.com/GoogleCloudPlatform/kubernetes/pull/7354) &#34;Improvements to generator of conversion methods.&#34;
[280]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7107](https://github.com/GoogleCloudPlatform/kubernetes/pull/7107) &#34;Code to automatically generate conversion methods&#34;
[281]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7407](https://github.com/GoogleCloudPlatform/kubernetes/pull/7407) &#34;Support recovery for anonymous roll outs&#34;
[282]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7527](https://github.com/GoogleCloudPlatform/kubernetes/pull/7527) &#34;Bump kube2sky to 1.2. Point it at https endpoint (3rd try).&#34;
[283]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7526](https://github.com/GoogleCloudPlatform/kubernetes/pull/7526) &#34;cluster/gce/coreos: Add metadata-service in node.yaml&#34;
[284]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7480](https://github.com/GoogleCloudPlatform/kubernetes/pull/7480) &#34;Move ComputePodChanges to the Docker runtime&#34;
[285]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7510](https://github.com/GoogleCloudPlatform/kubernetes/pull/7510) &#34;Cobra rebase&#34;
[286]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/6718](https://github.com/GoogleCloudPlatform/kubernetes/pull/6718) &#34;Adding system oom events from kubelet&#34;
[287]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7479](https://github.com/GoogleCloudPlatform/kubernetes/pull/7479) &#34;Move Prober to its own subpackage&#34;
[288]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7513](https://github.com/GoogleCloudPlatform/kubernetes/pull/7513) &#34;Fix parallel-e2e.sh to work on my macbook (bash v3.2)&#34;
[289]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7449](https://github.com/GoogleCloudPlatform/kubernetes/pull/7449) &#34;Move network plugin TearDown to DockerManager&#34;
[290]: [https://github.com/GoogleCloudPlatform/kubernetes/issues/7498](https://github.com/GoogleCloudPlatform/kubernetes/issues/7498) &#34;CoreOS Getting Started Guide not working&#34;
[291]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7499](https://github.com/GoogleCloudPlatform/kubernetes/pull/7499) &#34;Fixes #7498 - CoreOS Getting Started Guide had invalid cloud config&#34;
[292]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7504](https://github.com/GoogleCloudPlatform/kubernetes/pull/7504) &#34;Fix invalid character &#39;&#34;&#39; after object key:value pair&#34;
[293]: [https://github.com/GoogleCloudPlatform/kubernetes/issues/7317](https://github.com/GoogleCloudPlatform/kubernetes/issues/7317) &#34;GlusterFS Volume Plugin deletes the contents of the mounted volume upon Pod deletion&#34;
[294]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7503](https://github.com/GoogleCloudPlatform/kubernetes/pull/7503) &#34;Fixed kubelet deleting data from volumes on stop (#7317).&#34;
[295]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7482](https://github.com/GoogleCloudPlatform/kubernetes/pull/7482) &#34;Fixing hooks/description to catch API fields without description tags&#34;
[296]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7457](https://github.com/GoogleCloudPlatform/kubernetes/pull/7457) &#34;cadvisor is obsoleted so kubelet service does not require it.&#34;
[297]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7408](https://github.com/GoogleCloudPlatform/kubernetes/pull/7408) &#34;Set the default namespace for events to be &#34;default&#34;&#34;
[298]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7446](https://github.com/GoogleCloudPlatform/kubernetes/pull/7446) &#34;Fix typo in namespace conversion&#34;
[299]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7419](https://github.com/GoogleCloudPlatform/kubernetes/pull/7419) &#34;Convert Secret registry to use update/create strategy, allow filtering by Type&#34;
[300]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7102](https://github.com/GoogleCloudPlatform/kubernetes/pull/7102) &#34;Use pod namespace when looking for its GlusterFS endpoints.&#34;
[301]: [https://github.com/GoogleCloudPlatform/kubernetes/pull/7427](https://github.com/GoogleCloudPlatform/kubernetes/pull/7427) &#34;Fixed name of kube-proxy path in deployment scripts.&#34;
[


	

	


