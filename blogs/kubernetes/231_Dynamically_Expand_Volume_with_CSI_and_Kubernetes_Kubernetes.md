|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/08/02/dynamically-expand-volume-with-csi-and-kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-08-02 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.6376834 &#43;0300 MSK m=&#43;2.226718601  |

# Dynamically Expand Volume with CSI and Kubernetes | Kubernetes

	
	
	
	
	**Author**: Orain Xiong (Co-Founder, WoquTech)
*There is a very powerful storage subsystem within Kubernetes itself, covering a fairly broad spectrum of use cases. Whereas, when planning to build a product-grade relational database platform with Kubernetes, we face a big challenge: coming up with storage. This article describes how to extend latest Container Storage Interface 0.2.0 and integrate with Kubernetes, and demonstrates the essential facet of dynamically expanding volume capacity.*
As we focalize our customers, especially in financial space, there is a huge upswell in the adoption of container orchestration technology.
They are looking forward to open source solutions to redesign already existing monolithic applications, which have been running for several years on virtualization infrastructure or bare metal.
Considering extensibility and the extent of technical maturity, Kubernetes and Docker are at the very top of the list. But migrating monolithic applications to a distributed orchestration like Kubernetes is challenging, the relational database is critical for the migration.
With respect to the relational database, we should pay attention to storage. There is a very powerful storage subsystem within Kubernetes itself. It is very useful and covers a fairly broad spectrum of use cases. When planning to run a relational database with Kubernetes in production, we face a big challenge: coming up with storage. There are still some fundamental functionalities which are left unimplemented. Specifically, dynamically expanding volume. It sounds boring but is highly required, except for actions like create and delete and mount and unmount.
Currently, expanding volume is only available with those storage provisioners:
In order to enable this feature, we should set feature gate ```ExpandPersistentVolumes``` true and turn on the ```PersistentVolumeClaimResize``` admission plugin. Once ```PersistentVolumeClaimResize``` has been enabled, resizing will be allowed by a Storage Class whose ```allowVolumeExpansion``` field is set to true.
Unfortunately, dynamically expanding volume through the Container Storage Interface (CSI) and Kubernetes is unavailable, even though the underlying storage providers have this feature.
This article will give a simplified view of CSI, followed by a walkthrough of how to introduce a new expanding volume feature on the existing CSI and Kubernetes. Finally, the article will demonstrate how to dynamically expand volume capacity.
To have a better understanding of what we&#39;re going to do, the first thing we need to know is what the Container Storage Interface is. Currently, there are still some problems for already existing storage subsystem within Kubernetes. Storage driver code is maintained in the Kubernetes core repository which is difficult to test. But beyond that, Kubernetes needs to give permissions to storage vendors to check code into the Kubernetes core repository. Ideally, that should be implemented externally.
CSI is designed to define an industry standard that will enable storage providers who enable CSI to be available across container orchestration systems that support CSI.
This diagram depicts a kind of high-level Kubernetes archetypes integrated with CSI:

For more details, please visit: [https://github.com/container-storage-interface/spec/blob/master/spec.md](https://github.com/container-storage-interface/spec/blob/master/spec.md)
In order to enable the feature of expanding volume atop Kubernetes, we should extend several components including CSI specification, “in-tree” volume plugin, external-provisioner and external-attacher.
The feature of expanding volume is still undefined in latest CSI 0.2.0. The new 3 RPCs, including ```RequiresFSResize``` and ```ControllerResizeVolume``` and ```NodeResizeVolume```, should be introduced.
```service Controller {
 rpc CreateVolume (CreateVolumeRequest)
   returns (CreateVolumeResponse) {}
……
 rpc RequiresFSResize (RequiresFSResizeRequest)
   returns (RequiresFSResizeResponse) {}
 rpc ControllerResizeVolume (ControllerResizeVolumeRequest)
   returns (ControllerResizeVolumeResponse) {}
}

service Node {
 rpc NodeStageVolume (NodeStageVolumeRequest)
   returns (NodeStageVolumeResponse) {}
……
 rpc NodeResizeVolume (NodeResizeVolumeRequest)
   returns (NodeResizeVolumeResponse) {}
}
```In addition to the extend CSI specification, the ```csiPlugin﻿``` interface within Kubernetes should also implement ```expandablePlugin```. The ```csiPlugin``` interface will expand ```PersistentVolumeClaim``` representing for ```ExpanderController```.
Finally, to abstract complexity of the implementation, we should hard code the separate storage provider management logic into the following functions which is well-defined in the CSI specification:
Let’s demonstrate this feature with a concrete user case.
The Prometheus and Grafana integration allows us to visualize corresponding critical metrics.

We notice that the middle reading shows MySQL datafile size increasing slowly during bulk inserting. At the same time, the bottom reading shows file system expanding twice in about 20 minutes, from 300 GiB to 400 GiB and then 500 GiB. Meanwhile, the upper reading shows the whole process of expanding volume immediately completes and hardly impacts MySQL QPS.
Regardless of whatever infrastructure applications have been running on, the database is always a critical resource. It is essential to have a more advanced storage subsystem out there to fully support database requirements. This will help drive the more broad adoption of cloud native technology.


	

	


