|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/04/10/container-storage-interface-beta/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-04-10 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.9015664 &#43;0300 MSK m=&#43;2.490603201  |

# Container Storage Interface (CSI) for Kubernetes Goes Beta | Kubernetes

	
	
	
	
	
The Kubernetes implementation of the Container Storage Interface (CSI) is now beta in Kubernetes v1.10. CSI was [introduced as alpha](https://kubernetes.io/blog/2018/01/introducing-container-storage-interface) in Kubernetes v1.9.
Kubernetes features are generally introduced as alpha and moved to beta (and eventually to stable/GA) over subsequent Kubernetes releases. This process allows Kubernetes developers to get feedback, discover and fix issues, iterate on the designs, and deliver high quality, production grade features.
Although Kubernetes already provides a powerful volume plugin system that makes it easy to consume different types of block and file storage, adding support for new volume plugins has been challenging. Because volume plugins are currently “in-tree”—volume plugins are part of the core Kubernetes code and shipped with the core Kubernetes binaries—vendors wanting to add support for their storage system to Kubernetes (or even fix a bug in an existing volume plugin) must align themselves with the Kubernetes release process.
With the adoption of the Container Storage Interface, the Kubernetes volume layer becomes truly extensible. Third party storage developers can now write and deploy volume plugins exposing new storage systems in Kubernetes without ever having to touch the core Kubernetes code. This will result in even more options for the storage that backs Kubernetes users’ stateful containerized workloads.
With the promotion to beta CSI is now enabled by default on standard Kubernetes deployments instead of being opt-in.
The move of the Kubernetes implementation of CSI to beta also means:
CSI plugin authors must provide their own instructions for deploying their plugin on Kubernetes.
The Kubernetes-CSI implementation team created a [sample hostpath CSI driver](https://kubernetes-csi.github.io/docs/example.html). The sample provides a rough idea of what the deployment process for a CSI driver looks like. Production drivers, however, would deploy node components via a DaemonSet and controller components via a StatefulSet rather than a single pod (for example, see the deployment files for the [GCE PD driver](https://github.com/GoogleCloudPlatform/compute-persistent-disk-csi-driver/blob/master/deploy/kubernetes/README.md)).
Assuming a CSI storage plugin is already deployed on your cluster, you can use it through the familiar Kubernetes storage primitives: ```PersistentVolumeClaims```, ```PersistentVolumes```, and ```StorageClasses```.
CSI is a beta feature in Kubernetes v1.10. Although it is enabled by default, it may require the following flag:
You can enable automatic creation/deletion of volumes for CSI Storage plugins that support dynamic provisioning by creating a ```StorageClass``` pointing to the CSI plugin.
The following ```StorageClass```, for example, enables dynamic creation of “```fast-storage```” volumes by a CSI volume plugin called “```com.example.csi-driver```”.
```kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: fast-storage
provisioner: com.example.csi-driver
parameters:
  type: pd-ssd
  csiProvisionerSecretName: mysecret
  csiProvisionerSecretNamespace: mynamespace
```New for beta, the [default CSI external-provisioner](https://github.com/kubernetes-csi/external-provisioner) reserves the parameter keys ```csiProvisionerSecretName``` and ```csiProvisionerSecretNamespace```. If specified, it fetches the secret and passes it to the CSI driver during provisioning.
Dynamic provisioning is triggered by the creation of a ```PersistentVolumeClaim``` object. The following ```PersistentVolumeClaim```, for example, triggers dynamic provisioning using the ```StorageClass``` above.
```apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-request-for-storage
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  storageClassName: fast-storage
```When volume provisioning is invoked, the parameter type: ```pd-ssd``` and the secret any referenced secret(s) are passed to the CSI plugin ```com.example.csi-driver``` via a ```CreateVolume call```. In response, the external volume plugin provisions a new volume and then automatically create a ```PersistentVolume``` object to represent the new volume. Kubernetes then binds the new ```PersistentVolume``` object to the ```PersistentVolumeClaim```, making it ready to use.
If the fast-storage  StorageClass is marked as “default”, there is no need to include the storageClassName in the PersistentVolumeClaim, it will be used by default.
You can always expose a pre-existing volume in Kubernetes by manually creating a ```PersistentVolume``` object to represent the existing volume. The following ```PersistentVolume```, for example, exposes a volume with the name “```existingVolumeName```” belonging to a CSI storage plugin called “```com.example.csi-driver```”.
```apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-manually-created-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  csi:
    driver: com.example.csi-driver
    volumeHandle: existingVolumeName
    readOnly: false
    fsType: ext4
    volumeAttributes:
      foo: bar
    controllerPublishSecretRef:
      name: mysecret1
      namespace: mynamespace
    nodeStageSecretRef:
      name: mysecret2
      namespace: mynamespace
    nodePublishSecretRef
      name: mysecret3
      namespace: mynamespace
```You can reference a ```PersistentVolumeClaim``` that is bound to a CSI volume in any pod or pod template.
```kind: Pod
apiVersion: v1
metadata:
  name: my-pod
spec:
  containers:
    - name: my-frontend
      image: nginx
      volumeMounts:
      - mountPath: &#34;/var/www/html&#34;
        name: my-csi-volume
  volumes:
    - name: my-csi-volume
      persistentVolumeClaim:
        claimName: my-request-for-storage
```When the pod referencing a CSI volume is scheduled, Kubernetes will trigger the appropriate operations against the external CSI plugin (```ControllerPublishVolume```, ```NodeStageVolume```, ```NodePublishVolume```, etc.) to ensure the specified volume is attached, mounted, and ready to use by the containers in the pod.
For more details please see the CSI implementation [design doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md) and [documentation](https://kubernetes-csi.github.io/).
CSI Volume Driver deployments on Kubernetes must meet some [minimum requirements](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md#third-party-csi-volume-drivers).
The minimum requirements document also outlines the [suggested mechanism](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md#recommended-mechanism-for-deploying-csi-drivers-on-kubernetes) for deploying an arbitrary containerized CSI driver on Kubernetes. This mechanism can be used by a Storage Provider to simplify deployment of containerized CSI compatible volume drivers on Kubernetes.
As part of the suggested deployment process, the Kubernetes team provides the following sidecar (helper) containers:
Storage vendors can build Kubernetes deployments for their plugins using these components, while leaving their CSI driver completely unaware of Kubernetes.
CSI drivers are developed and maintained by third parties. You can find a non-definitive list of some [sample and production CSI drivers](https://kubernetes-csi.github.io/docs/drivers.html).
As mentioned in the [alpha release blog post](https://kubernetes.io/blog/2018/01/introducing-container-storage-interface), [FlexVolume plugin](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-storage/flexvolume.md) was an earlier attempt to make the Kubernetes volume plugin system extensible. Although it enables third party storage vendors to write drivers “out-of-tree”, because it is an exec based API, FlexVolumes requires files for third party driver binaries (or scripts) to be copied to a special plugin directory on the root filesystem of every node (and, in some cases, master) machine. This requires a cluster admin to have write access to the host filesystem for each node and some external mechanism to ensure that the driver file is recreated if deleted, just to deploy a volume plugin.
In addition to being difficult to deploy, Flex did not address the pain of plugin dependencies: Volume plugins tend to have many external requirements (on mount and filesystem tools, for example). These dependencies are assumed to be available on the underlying host OS, which is often not the case.
CSI addresses these issues by not only enabling storage plugins to be developed out-of-tree, but also containerized and deployed via standard Kubernetes primitives.
If you still have questions about in-tree volumes vs CSI vs Flex, please see the [Volume Plugin FAQ](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md).
Once CSI reaches stability, we plan to migrate most of the in-tree volume plugins to CSI. Stay tuned for more details as the Kubernetes CSI implementation approaches stable.
The beta implementation of CSI has the following limitations:
Depending on feedback and adoption, the Kubernetes team plans to push the CSI implementation to GA in 1.12.
The team would like to encourage storage vendors to start developing CSI drivers, deploying them on Kubernetes, and sharing feedback with the team via the Kubernetes Slack channel [wg-csi](https://kubernetes.slack.com/messages/C8EJ01Z46/details/), the Google group [kubernetes-sig-storage-wg-csi](https://groups.google.com/forum/#!forum/kubernetes-sig-storage-wg-csi), or any of the standard [SIG storage communication channels](https://github.com/kubernetes/community/blob/master/sig-storage/README.md#contact).
This project, like all of Kubernetes, is the result of hard work by many contributors from diverse backgrounds working together.
In addition to the contributors who have been working on the Kubernetes implementation of CSI since alpha:
We offer a huge thank you to the new contributors who stepped up this quarter to help the project reach beta:
If you’re interested in getting involved with the design and development of CSI or any part of the Kubernetes Storage system, join the [Kubernetes Storage Special Interest Group](https://github.com/kubernetes/community/tree/master/sig-storage) (SIG). We’re rapidly growing and always welcome new contributors.


	

	


