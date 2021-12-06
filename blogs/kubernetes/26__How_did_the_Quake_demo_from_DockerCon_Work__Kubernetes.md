|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2015/07/How-Did-Quake-Demo-From-Dockercon-Work/        |
| Tags              | [kubernetes]       |
| Date Create       | 2015-07-02 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:50:38.5895571 &#43;0300 MSK m=&#43;7.978234501  |

#  How did the Quake demo from DockerCon Work?  | Kubernetes

	
	
	
	
	Shortly after its release in 2013, Docker became a very popular open source container management tool for Linux.  Docker has a rich set of commands to control the execution of a container. Commands such as start, stop, restart, kill, pause, and unpause. However, what is still missing is the ability to Checkpoint and Restore (C/R) a container natively via Docker itself.
We’ve been actively working with upstream and community developers to add support in Docker for native C/R and hope that checkpoint and restore commands will be introduced in Docker 1.8.  As of this writing, it’s possible to C/R a container externally because this functionality was recently merged in libcontainer.
External container C/R was demo’d at DockerCon 2015:

Container C/R offers many benefits including the following:
CRIU
Implementing C/R functionality from scratch is a major undertaking and a daunting task.  Fortunately, there is a powerful open source tool written in C that has been used in production for checkpointing and restoring entire process trees in Linux.  The tool is called CRIU which stands for Checkpoint Restore In Userspace ([http://criu.org](http://criu.org)).  CRIU works by:
In April 2014, we decided to find out if CRIU could checkpoint and restore Docker containers to facilitate container migration.
The first phase of this effort invoking CRIU directly to dump a process tree running inside a container and determining why the checkpoint or restore operation failed.  There were quite a few issues that caused CRIU failure.  The following three issues were among the more challenging ones.
Docker sets up /etc/{hostname,hosts,resolv.conf} as targets with source files outside the container&#39;s mount namespace.
The --ext-mount-map command line option was added to CRIU to specify the path of the external bind mounts.  For example, assuming default Docker configuration, /etc/hostname in the container&#39;s mount namespace is bind mounted from the source at /var/lib/docker/containers/&lt;container-id&gt;/hostname.  When checkpointing, we tell CRIU to record /etc/hostname&#39;s &#34;map&#34; as, say, etc_hostname. When restoring, we tell CRIU that the file previously recorded as etc_hostname should be mapped from the external bind mount at /var/lib/docker/containers/&lt;container-id&gt;/hostname.

Docker initially used AUFS as its preferred filesystem which is still in wide usage (the preferred filesystem is now OverlayFS)..  Due to a bug, the AUFS symbolic link paths of /proc/&lt;pid&gt;/map_files point inside AUFS branches instead of their pathnames relative to the container&#39;s root. This problem has been fixed in AUFS source code but hasn&#39;t made it to all the distros yet.  CRIU would get confused seeing the same file in its physical location (in the branch) and its logical location (from the root of mount namespace).
The --root command line option that was used only during restore was generalized to understand the root of the mount namespace during checkpoint and automatically &#34;fix&#34; the exposed AUFS pathnames.
After checkpointing, the Docker daemon removes the container’s cgroups subdirectories (because the container has “exited”).  This causes restore to fail.
The --manage-cgroups command line option was added to CRIU to dump and restore the process&#39;s cgroups along with their properties.
The CRIU command lines are a simple container are shown below:
```$ docker run -d busybox:latest /bin/sh -c &#39;i=0; while true; do echo $i \&gt;\&gt; /foo; i=$(expr $i &#43; 1); sleep 3; done&#39;  

$ docker ps  
CONTAINER ID  IMAGE           COMMAND           CREATED        STATUS  
168aefb8881b  busybox:latest  &#34;/bin/sh -c &#39;i=0; 6 seconds ago  Up 4 seconds  

$ sudo criu dump -o dump.log -v4 -t 17810 \  
        -D /tmp/img/\&lt;container\_id\&gt; \  
        --root /var/lib/docker/aufs/mnt/\&lt;container\_id\&gt; \  
        --ext-mount-map /etc/resolv.conf:/etc/resolv.conf \  
        --ext-mount-map /etc/hosts:/etc/hosts \  
        --ext-mount-map /etc/hostname:/etc/hostname \  
        --ext-mount-map /.dockerinit:/.dockerinit \  
        --manage-cgroups \  
        --evasive-devices  

$ docker ps -a  
CONTAINER ID  IMAGE           COMMAND           CREATED        STATUS  
168aefb8881b  busybox:latest  &#34;/bin/sh -c &#39;i=0; 6 minutes ago  Exited (-1) 4 minutes ago  

$ sudo mount -t aufs -o br=\  
/var/lib/docker/aufs/diff/\&lt;container\_id\&gt;:\  
/var/lib/docker/aufs/diff/\&lt;container\_id\&gt;-init:\  
/var/lib/docker/aufs/diff/a9eb172552348a9a49180694790b33a1097f546456d041b6e82e4d7716ddb721:\  
/var/lib/docker/aufs/diff/120e218dd395ec314e7b6249f39d2853911b3d6def6ea164ae05722649f34b16:\  
/var/lib/docker/aufs/diff/42eed7f1bf2ac3f1610c5e616d2ab1ee9c7290234240388d6297bc0f32c34229:\  
/var/lib/docker/aufs/diff/511136ea3c5a64f264b78b5433614aec563103b4d4702f3ba7d4d2698e22c158:\  
none /var/lib/docker/aufs/mnt/\&lt;container\_id\&gt;  

$ sudo criu restore -o restore.log -v4 -d  
        -D /tmp/img/\&lt;container\_id\&gt; \  
        --root /var/lib/docker/aufs/mnt/\&lt;container\_id\&gt; \  
        --ext-mount-map /etc/resolv.conf:/var/lib/docker/containers/\&lt;container\_id\&gt;/resolv.conf \  
        --ext-mount-map /etc/hosts:/var/lib/docker/containers/\&lt;container\_id\&gt;/hosts \  
        --ext-mount-map /etc/hostname:/var/lib/docker/containers/\&lt;container\_id\&gt;/hostname \  
        --ext-mount-map /.dockerinit:/var/lib/docker/init/dockerinit-1.0.0 \  
        --manage-cgroups \  
        --evasive-devices  

$ ps -ef | grep /bin/sh  
root     18580     1  0 12:38 ?        00:00:00 /bin/sh -c i=0; while true; do echo $i \&gt;\&gt; /foo; i=$(expr $i &#43; 1); sleep 3; done  

$ docker ps -a  
CONTAINER ID  IMAGE           COMMAND           CREATED        STATUS  
168aefb8881b  busybox:latest  &#34;/bin/sh -c &#39;i=0; 7 minutes ago  Exited (-1) 5 minutes ago  

docker\_cr.sh
```Since the command line arguments to CRIU were long, a helper script called docker_cr.sh was provided in the CRIU source tree to simplify the process.  So, for the above container, one would simply C/R the container as follows (for details see [http://criu.org/Docker](http://criu.org/Docker)):
```$ sudo docker\_cr.sh -c 4397   
dump successful  

$ sudo docker\_cr.sh -r 4397  
restore successful  
```At the end of Phase 1, it was possible to externally checkpoint and restore a Docker 1.0 container using either VFS, AUFS, or UnionFS storage drivers with CRIU v1.3.
While external C/R served as a successful proof of concept for container C/R, it had the following limitations:
Therefore, the second phase of the effort concentrated on adding native checkpoint and restore commands to Docker.
Libcontainer is Docker’s native execution driver.  It provides a set of APIs to create and manage containers.  The first step of adding native support was the introduction of two methods, checkpoint() and restore(), to libcontainer and the corresponding checkpoint and restore subcommands to nsinit.  Nsinit is a simple utility that is used to test and debug libcontainer.
With C/R support in libcontainer, the next step was adding checkpoint and restore subcommands to Docker itself. A big challenge in this step was to rebuild the “plumbing” between the container and the daemon.  When the daemon initially starts a container, it sets up individual pipes between itself (parent) and the standard input, output, and error file descriptors of the container (child).  This is how docker logs can show the output of a container.
When a container exits after being checkpointed, the pipes between it and the daemon are deleted.  During container restore, it’s actually CRIU that is the parent.  Therefore, setting up a pipe between the child (container) and an unrelated process (Docker daemon) required is not a bit of challenge.
To address this issue, the --inherit-fd command line option was added to CRIU.  Using this option, the Docker daemon tells CRIU to let the restored container “inherit” certain file descriptors passed from the daemon to CRIU.
The first version of native C/R was demo&#39;ed at the Linux Plumbers Conference (LPC) in October 2014 ([http://linuxplumbersconf.org/2014/ocw/proposals/1899](http://linuxplumbersconf.org/2014/ocw/proposals/1899)).

The LPC demo was done with a simple container that did not require network connectivity.  Support for restoring network connections was done in early 2015 and demonstrated in this 2-minute [video clip](https://www.youtube.com/watch?v=HFt9v6yqsXo).
In May 2015, the criu branch of libcontainer was merged into master.  Using the newly-introduced lightweight [runC](https://blog.docker.com/2015/06/runc/) container runtime, container migration was demo’ed at DockerCon15.  In this
[img](https://www.youtube.com/watch?v=?mL9AFkJJAq0) (minute 23:00), a container running Quake was checkpointed and restored on a different machine, effectively implementing container migration.
At the time of this writing, there are two repos on GitHub that have native C/R support in Docker:
Work is underway to merge C/R functionality into Docker.  You can use either of the above repositories to experiment with Docker C/R.  If you are using OverlayFS or your container workload uses AIO, please note the following:
When OverlayFS support was officially merged into the Linux kernel version 3.18, it became the preferred storage driver (instead of AUFS) .  However, OverlayFS in 3.18 has the following issues:
Both issues are fixed in this [patch](https://lkml.org/lkml/2015/3/20/372) but the patch has not been merged upstream yet.
If you are using a kernel older than 3.19 and your container uses AIO, you need the following kernel patches from 3.19:


	

	


