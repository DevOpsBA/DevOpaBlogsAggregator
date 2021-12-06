|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2019/03/15/kubernetes-setup-using-ansible-and-vagrant/        |
| Tags              | [kubernetes]       |
| Date Create       | 2019-03-15 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.354224 &#43;0300 MSK m=&#43;1.943257601  |

# Kubernetes Setup Using Ansible and Vagrant | Kubernetes

	
	
	
	
	**Author:** Naresh L J (Infosys)
This blog post describes the steps required to setup a multi node Kubernetes cluster for development purposes. This setup provides a production-like cluster that can be setup on your local machine.
Multi node Kubernetes clusters offer a production-like environment which has various advantages. Even though Minikube provides an excellent platform for getting started, it doesn&#39;t provide the opportunity to work with multi node clusters which can help solve problems or bugs that are related to application design and architecture. For instance, Ops can reproduce an issue in a multi node cluster environment, Testers can deploy multiple versions of an application for executing test cases and verifying changes. These benefits enable teams to resolve issues faster which make the more agile.
Vagrant is a tool that will allow us to create a virtual environment easily and it eliminates pitfalls that cause the works-on-my-machine phenomenon. It can be used with multiple providers such as Oracle VirtualBox, VMware, Docker, and so on. It allows us to create a disposable environment by making use of configuration files.
Ansible is an infrastructure automation engine that automates software configuration management. It is agentless and allows us to use SSH keys for connecting to remote machines. Ansible playbooks are written in yaml and offer inventory management in simple text files.
We will be setting up a Kubernetes cluster that will consist of one master and two worker nodes. All the nodes will run Ubuntu Xenial 64-bit OS and Ansible playbooks will be used for provisioning.
Use the text editor of your choice and create a file with named ```Vagrantfile```, inserting the code below. The value of N denotes the number of nodes present in the cluster, it can be modified accordingly. In the below example, we are setting the value of N as 2.
Create a directory named ```kubernetes-setup``` in the same directory as the ```Vagrantfile```. Create two files named ```master-playbook.yml``` and ```node-playbook.yml``` in the directory ```kubernetes-setup```.
In the file ```master-playbook.yml```, add the code below.
We will be installing the following packages, and then adding a user named “vagrant” to the “docker” group.
Create a file named ```node-playbook.yml``` in the directory ```kubernetes-setup```.
Add the code below into ```node-playbook.yml```
Upon completion of all the above steps, the Kubernetes cluster should be up and running.
We can login to the master or worker nodes using Vagrant as follows:

	

	


