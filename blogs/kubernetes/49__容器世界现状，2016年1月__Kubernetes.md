|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2016/02/State-Of-Container-World-January-2016/        |
| Tags              | [kubernetes]       |
| Date Create       | 2016-02-01 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:23.0623706 &#43;0300 MSK m=&#43;4.651419701  |

#  容器世界现状，2016年1月  | Kubernetes

	
	
	
	
	新年伊始，我们进行了一项调查，以评估容器世界的现状。
我们已经准备好发送[ 2 月版](https://docs.google.com/forms/d/13yxxBqb5igUhwrrnDExLzZPjREiCnSs-AH-y4SSZ-5c/viewform)但在此之前，让我们从 119 条回复中看一下 1 月的数据（感谢您的参与！）。
关于这些数字的注释：
首先，您可能会注意到，这些数字加起来并不是 100％，在大多数情况下，选择不是唯一的，因此给出的百分比是选择特定选择的所有受访者的百分比。
其次，虽然我们尝试覆盖广泛的云社区，但调查最初是通过 Twitter 发送给[@brendandburns](https://twitter.com/brendandburns)，[@kelseyhightower](https://twitter.com/kelseyhightower)，[@sarahnovotny](https://twitter.com/sarahnovotny)，[@juliaferraioli](https://twitter.com/juliaferraioli)，[@thagomizer_rb](https://twitter.com/thagomizer_rb)，因此受众覆盖可能并不完美。
我们正在努力扩大样本数量（我是否提到过2月份的调查？[点击立即参加](https://docs.google.com/forms/d/13yxxBqb5igUhwrrnDExLzZPjREiCnSs-AH-y4SSZ-5c/viewform))。
首先，很多人正在使用容器！目前有 71％ 的人正在使用容器，而其中有 24％ 的人正在考虑尽快使用它们。
显然，这表明样本集有些偏颇。
在更广泛的社区中，容器使用的数量有所不同，但绝对低于 71％。
因此，对这些数字的其余部分要持怀疑态度。
那么人们在使用容器做什么呢？
超过 80％ 的受访者使用容器进行开发，但只有 50％ 的人在生产环境下使用容器。
但是他们有计划很快投入到生产环境之中，78% 的容器用户表示了意愿。
你们在哪里部署容器？
你的笔记本电脑显然是赢家，53% 的人使用笔记本电脑。
接下来是 44％ 的人在自己的 VM 上运行（Vagrant？OpenStack？我们将在2月的调查中尝试深入研究），然后是 33％ 的人在物理基础架构上运行，而 31％ 的人在公共云 VM 上运行。
如何部署容器？
你们当中有 54% 的人使用 Kubernetes，虽然看起来有点受样本集的偏见（请参阅上面的注释），但真是令人惊讶，但有 45％ 的人在使用 shell 脚本。
是因为 Kubernetes 存储库中正在运行大量（且很好）的 Bash 脚本吗？
继续下去，我们可以看到真相……
数据显示，25% 使用 CAPS (Chef/Ansible/Puppet/Salt)系统，约 13% 使用 Docker Swarm、Mesos 或其他系统。
最后，我们让人们自由回答使用容器的挑战。
这儿有一些进行了分组和复制的最有趣的答案：
*下载完整的调查结果 [链接](https://docs.google.com/spreadsheets/d/18wZe7wEDvRuT78CEifs13maXoSGem_hJvbOSmsuJtkA/pub?gid=530616014&amp;single=true&amp;output=csv) (CSV 文件）。*
_Up-- Brendan Burns，Google 软件工程师


	

	


