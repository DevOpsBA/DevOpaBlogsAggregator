|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2019/01/15/container-storage-interface-ga/        |
| Tags              | [kubernetes]       |
| Date Create       | 2019-01-15 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.4349552 &#43;0300 MSK m=&#43;2.023989301  |

# Container Storage Interface (CSI) for Kubernetes GA | Kubernetes

	
	
	
	
	
**Author:** Saad Ali, Senior Software Engineer, Google
The Kubernetes implementation of the [Container Storage Interface](https://github.com/container-storage-interface/spec/blob/master/spec.md) (CSI) has been promoted to GA in the Kubernetes v1.13 release. Support for CSI was [introduced as alpha](http://blog.kubernetes.io/2018/01/introducing-container-storage-interface.html) in Kubernetes v1.9 release, and [promoted to beta](https://kubernetes.io/blog/2018/04/10/container-storage-interface-beta/) in the Kubernetes v1.10 release.
The GA milestone indicates that Kubernetes users may depend on the feature and its API without fear of backwards incompatible changes in future causing regressions. GA features are protected by the [Kubernetes deprecation policy](/docs/reference/using-api/deprecation-policy/).
Although prior to CSI Kubernetes provided a powerful volume plugin system, it was challenging to add support for new volume plugins to Kubernetes: volume plugins were “in-tree” meaning their code was part of the core Kubernetes code and shipped with the core Kubernetes binaries—vendors wanting to add support for their storage system to Kubernetes (or even fix a bug in an existing volume plugin) were forced to align with the Kubernetes release process. In addition, third-party storage code caused reliability and security issues in core Kubernetes binaries and the code was often difficult (and in some cases impossible) for Kubernetes maintainers to test and maintain.
CSI was developed as a standard for exposing arbitrary block and file storage storage systems to containerized workloads on Container Orchestration Systems (COs) like Kubernetes. With the adoption of the Container Storage Interface, the Kubernetes volume layer becomes truly extensible. Using CSI, third-party storage providers can write and deploy plugins exposing new storage systems in Kubernetes without ever having to touch the core Kubernetes code. This gives Kubernetes users more options for storage and makes the system more secure and reliable.
With the promotion to GA, the Kubernetes implementation of CSI introduces the following changes:
Kubernetes users interested in how to deploy or manage an existing CSI driver on Kubernetes should look at the documentation provided by the author of the CSI driver.
Assuming a CSI storage plugin is already deployed on a Kubernetes cluster, users can use CSI volumes through the familiar Kubernetes storage API objects: ```PersistentVolumeClaims```, ```PersistentVolumes```, and ```StorageClasses```. Documented [here](/docs/concepts/storage/volumes/#csi).
Although the Kubernetes implementation of CSI is a GA feature in Kubernetes v1.13, it may require the following flag:
You can enable automatic creation/deletion of volumes for CSI Storage plugins that support dynamic provisioning by creating a ```StorageClass``` pointing to the CSI plugin.
The following StorageClass, for example, enables dynamic creation of “```fast-storage```” volumes by a CSI volume plugin called “```csi-driver.example.com```”.
```kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: fast-storage
provisioner: csi-driver.example.com
parameters:
  type: pd-ssd
  csi.storage.k8s.io/provisioner-secret-name: mysecret
  csi.storage.k8s.io/provisioner-secret-namespace: mynamespace
```New for GA, the [CSI external-provisioner](https://github.com/kubernetes-csi/external-provisioner) (v1.0.1&#43;) reserves the parameter keys prefixed with ```csi.storage.k8s.io/```. If the keys do not correspond to a set of known keys the values are simply ignored (and not passed to the CSI driver). The older secret parameter keys (```csiProvisionerSecretName```, ```csiProvisionerSecretNamespace```, etc.) are also supported by CSI external-provisioner v1.0.1 but are deprecated and may be removed in future releases of the CSI external-provisioner.
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
```When volume provisioning is invoked, the parameter type: ```pd-ssd``` and the secret any referenced secret(s) are passed to the CSI plugin ```csi-driver.example.com``` via a ```CreateVolume``` call. In response, the external volume plugin provisions a new volume and then automatically create a ```PersistentVolume``` object to represent the new volume. Kubernetes then binds the new ```PersistentVolume``` object to the ```PersistentVolumeClaim```, making it ready to use.
If the ```fast-storage  StorageClass``` is marked as “default”, there is no need to include the ```storageClassName``` in the ```PersistentVolumeClaim```, it will be used by default.
You can always expose a pre-existing volume in Kubernetes by manually creating a PersistentVolume object to represent the existing volume. The following ```PersistentVolume```, for example, exposes a volume with the name “```existingVolumeName```” belonging to a CSI storage plugin called “```csi-driver.example.com```”.
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
    driver: csi-driver.example.com
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
For more details please see the CSI implementation [design doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md) and [documentation](/docs/concepts/storage/volumes/#csi).
The [kubernetes-csi](https://kubernetes-csi.github.io/) site details how to develop, deploy, and test a CSI driver on Kubernetes. In general, CSI Drivers should be deployed on Kubernetes along with the following sidecar (helper) containers:
Storage vendors can build Kubernetes deployments for their plugins using these components, while leaving their CSI driver completely unaware of Kubernetes.
CSI drivers are developed and maintained by third parties. You can find a non-definitive list of CSI drivers [here](https://kubernetes-csi.github.io/docs/drivers.html).
There is a plan to migrate most of the persistent, remote in-tree volume plugins to CSI. For more details see [design doc](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/csi-migration.md).
The GA implementation of CSI has the following limitations:
The Kubernetes Slack channel [wg-csi](https://kubernetes.slack.com/messages/C8EJ01Z46/details/) and the Google group [kubernetes-sig-storage-wg-csi](https://groups.google.com/forum/#!forum/kubernetes-sig-storage-wg-csi) along with any of the standard [SIG storage communication channels](https://github.com/kubernetes/community/blob/master/sig-storage/README.md#contact) are all great mediums to reach out to the SIG Storage team.
This project, like all of Kubernetes, is the result of hard work by many contributors from diverse backgrounds working together. We offer a huge thank you to the new contributors who stepped up this quarter to help the project reach GA:
If you’re interested in getting involved with the design and development of CSI or any part of the Kubernetes Storage system, join the [Kubernetes Storage Special Interest Group](https://github.com/kubernetes/community/tree/master/sig-storage) (SIG). We’re rapidly growing and always welcome new contributors.


	

	


