|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/10/02/building-a-network-bootable-server-farm-for-kubernetes-with-ltsp/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-10-02 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.5729817 &#43;0300 MSK m=&#43;2.162016601  |

# Building a Network Bootable Server Farm for Kubernetes with LTSP | Kubernetes

	
	
	
	
	**Author**: Andrei Kvapil (WEDOS)

In this post, I&#39;m going to introduce you to a cool technology for Kubernetes, LTSP. It is useful for large baremetal Kubernetes deployments.
You don&#39;t need to think about installing an OS and binaries on each node anymore. Why? You can do that automatically through Dockerfile!
You can buy and put 100 new servers into a production environment and get them working immediately - it&#39;s really amazing!
Intrigued? Let me walk you through how it works.
***Please note:** this is a cool hack, but is not officially supported in Kubernetes.*
First, we need to understand how exactly it works.
In short, for all nodes we have prepared the image with the OS, Docker, Kubelet and everything else that you need there. This image with the kernel is building automatically by CI using Dockerfile. End nodes are booting the kernel and OS from this image via the network.
Nodes are using overlays as the root filesystem and after reboot any changes will be lost (like in Docker containers). You have a config-file where you can describe mounts and some initial commands which should be executed during node boot (Example: set root user ssh-key and kubeadm join commands)
We will use LTSP project because it&#39;s gives us everything we need to organize the network booting environment. Basically, LTSP is a pack of shell-scripts which makes our life much easier.
LTSP provides a initramfs module, a few helper-scripts, and the configuration system which prepare the system during the early state of boot, before the main init process call.
**This is what the image preparation procedure looks like:**
After that, you will get the squashed image from the chroot with all the software inside. Each node will download this image during the boot and use it as the rootfs. For the update node, you can just reboot it. The new squashed image will be downloaded and mounted into the rootfs.
**The server part of LTSP includes two components in our case:**
You should also have:
**This is how the node is booting up**
As I said before, I&#39;m preparing the LTSP-server with the squashed image automatically using Dockerfile. This method is quite good because you have all steps described in your git repository.
You have versioning, branches, CI and everything that you used to use for preparing your usual Docker projects.
Otherwise, you can deploy the LTSP server manually by executing all steps by hand. This is a good practice for learning and understanding the basic principles.
Just repeat all the steps listed here by hand, just to try to install LTSP without Dockerfile.
LTSP still has some issues which authors don’t want to apply, yet. However LTSP is easy customizable so I prepared a few patches for myself and will share them here.
I’ll create a fork if the community will warmly accept my solution.
We will use [stage building](https://docs.docker.com/develop/develop-images/multistage-build/) in our Dockerfile to leave only the needed parts in our Docker image. The unused parts will be removed from the final image.
```ltsp-base
(install basic LTSP server software)
   |
   |---basesystem
   |   (prepare chroot with main software and kernel)
   |     |
   |     |---builder
   |     |   (build additional software from sources, if needed)
   |     |
   |     &#39;---ltsp-image
   |         (install additional software, docker, kubelet and build squashed image)
   |
   &#39;---final-stage
       (copy squashed image, kernel and initramfs into first stage)
```Let&#39;s start writing our Dockerfile. This is the first part:
At this stage our Docker image has already been installed:
In this stage we will prepare a chroot environment with basesystem, and install basic software with the kernel.
We will use the classic **debootstrap** instead of **ltsp-build-client** to prepare the base image, because **ltsp-build-client** will install GUI and few other things which we don&#39;t need for the server deployment.
Note that you may encounter problems with some packages, such as ```lvm2```.
They have not fully optimized for installing in an unprivileged chroot.
Their postinstall scripts try to call some privileged commands which can fail with errors and block the package installation.
Solution:
Now we can build all the necessary software and kernel modules. It&#39;s really cool that you can do that automatically in this stage.
You can skip this stage if you have nothing to do here.
Here is example for install latest MLNX_EN driver:
In this stage we will install what we built in the previous step:
Then do some additional changes to finalize our ltsp-image:
Then we will make the squashed image from our chroot:
In the final stage we will save only our squashed image and kernels with initramfs.
Ok, now we have docker image which includes:
OK, now when our docker-image with LTSP-server, kernel, initramfs and squashed rootfs fully prepared we can run the deployment with it.
We can do that as usual, but one more thing is networking.
Unfortunately, we can&#39;t use the standard Kubernetes service abstraction for our deployment, because TFTP can&#39;t work behind the NAT. During the boot, our nodes are not part of Kubernetes cluster and they requires ExternalIP, but Kubernetes always enables NAT for ExternalIPs, and there is no way to override this behavior.
For now I have two ways for avoid this: use ```hostNetwork: true``` or use [pipework](https://github.com/dreamcat4/docker-images/blob/master/pipework/3.%20Examples.md#kubernetes). The second option will also provide you redundancy because, in case of failure, the IP will be moved with the Pod to another node. Unfortunately, pipework is not native and a less secure method.
If you have some better option for that please let me know.
Here is example for deployment with hostNetwork:
As you can see it also requires configmap with **lts.conf** file.
Here is example part from mine:
```apiVersion: v1
kind: ConfigMap
metadata:
  name: ltsp-config
data:
  lts.conf: |
    [default]
    KEEP_SYSTEM_SERVICES           = &#34;ssh ureadahead dbus-org.freedesktop.login1 systemd-logind polkitd cgmanager ufw rpcbind nfs-kernel-server&#34;

    PREINIT_00_TIME                = &#34;ln -sf /usr/share/zoneinfo/Europe/Prague /etc/localtime&#34;
    PREINIT_01_FIX_HOSTNAME        = &#34;sed -i &#39;/^127.0.0.2/d&#39; /etc/hosts&#34;
    PREINIT_02_DOCKER_OPTIONS      = &#34;sed -i &#39;s|^ExecStart=.*|ExecStart=/usr/bin/dockerd -H fd:// --storage-driver overlay2 --iptables=false --ip-masq=false --log-driver=json-file --log-opt=max-size=10m --log-opt=max-file=5|&#39; /etc/systemd/system/docker.service&#34;

    FSTAB_01_SSH                   = &#34;/dev/data/ssh     /etc/ssh          ext4 nofail,noatime,nodiratime 0 0&#34;
    FSTAB_02_JOURNALD              = &#34;/dev/data/journal /var/log/journal  ext4 nofail,noatime,nodiratime 0 0&#34;
    FSTAB_03_DOCKER                = &#34;/dev/data/docker  /var/lib/docker   ext4 nofail,noatime,nodiratime 0 0&#34;

    # Each command will stop script execution when fail
    RCFILE_01_SSH_SERVER           = &#34;cp /rofs/etc/ssh/*_config /etc/ssh; ssh-keygen -A&#34;
    RCFILE_02_SSH_CLIENT           = &#34;mkdir -p /root/.ssh/; echo &#39;ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDBSLYRaORL2znr1V4a3rjDn3HDHn2CsvUNK1nv8&#43;CctoICtJOPXl6zQycI9KXNhANfJpc6iQG1ZPZUR74IiNhNIKvOpnNRPyLZ5opm01MVIDIZgi9g0DUks1g5gLV5LKzED8xYKMBmAfXMxh/nsP9KEvxGvTJB3OD&#43;/bBxpliTl5xY3Eu41&#43;VmZqVOz3Yl98&#43;X8cZTgqx2dmsHUk7VKN9OZuCjIZL9MtJCZyOSRbjuo4HFEssotR1mvANyz&#43;BUXkjqv2pEa0I2vGQPk1VDul5TpzGaN3nOfu83URZLJgCrX&#43;8whS1fzMepUYrbEuIWq95esjn0gR6G4J7qlxyguAb9 admin@kubernetes&#39; &gt;&gt; /root/.ssh/authorized_keys&#34;
    RCFILE_03_KERNEL_DEBUG         = &#34;sysctl -w kernel.unknown_nmi_panic=1 kernel.softlockup_panic=1; modprobe netconsole netconsole=@/vmbr0,@10.9.0.15/&#34;
    RCFILE_04_SYSCTL               = &#34;sysctl -w fs.file-max=20000000 fs.nr_open=20000000 net.ipv4.neigh.default.gc_thresh1=80000 net.ipv4.neigh.default.gc_thresh2=90000 net.ipv4.neigh.default.gc_thresh3=100000&#34;
    RCFILE_05_FORWARD              = &#34;echo 1 &gt; /proc/sys/net/ipv4/ip_forward&#34;
    RCFILE_06_MODULES              = &#34;modprobe br_netfilter&#34;
    RCFILE_07_JOIN_K8S             = &#34;kubeadm join --token 2a4576.504356e45fa3d365 10.9.0.20:6443 --discovery-token-ca-cert-hash sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855&#34;
```You can get more details on all the variables used from [lts.conf manpage](http://manpages.ubuntu.com/manpages/xenial/man5/lts.conf.5.html).
Now you can configure your DHCP. Basically you should set the ```next-server``` and ```filename``` options.
I use ISC-DHCP server, and here is an example ```dhcpd.conf```:
```shared-network ltsp-netowrk {
    subnet 10.9.0.0 netmask 255.255.0.0 {
        authoritative;
        default-lease-time -1;
        max-lease-time -1;

        option domain-name              &#34;example.org&#34;;
        option domain-name-servers      10.9.0.1;
        option routers                  10.9.0.1;
        next-server                     ltsp-1;  # write LTSP-server hostname here

        if option architecture = 00:07 {
            filename &#34;/ltsp/amd64/grub/x86_64-efi/core.efi&#34;;
        } else {
            filename &#34;/ltsp/amd64/grub/i386-pc/core.0&#34;;
        }

        range 10.9.200.0 10.9.250.254; 
    }
```You can start from this, but what about me, I have multiple LTSP-servers and I configure leases statically for each node via the Ansible playbook.
Try to run your first node. If everything was right, you will have a running system there.
The node also will be added to your Kubernetes cluster.
Now you can try to make your own changes.
If you need something more, note that LTSP can be easily changed to meet your needs.
Feel free to look into the source code and you can find many answers there.
***UPD:** Many people asking me: Why not simple use CoreOS and Ignition?*
*I can answer. The main feature here is image preparation process, not configuration. In case with LTSP you have classic Ubuntu system, and everything that can be installed on Ubuntu it can also be written here in the Dockerfile. In case CoreOS you have no so many freedom and you can’t easily add custom kernel modules and packages at the build stage of the boot image.*


	

	


