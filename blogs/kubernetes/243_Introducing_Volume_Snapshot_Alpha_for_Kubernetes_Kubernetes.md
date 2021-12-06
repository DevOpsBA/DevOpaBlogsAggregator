|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/10/09/introducing-volume-snapshot-alpha-for-kubernetes/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-10-09 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.5430384 &#43;0300 MSK m=&#43;2.132073101  |

# Introducing Volume Snapshot Alpha for Kubernetes | Kubernetes

	
	
	
	
	**Author**: Jing Xu (Google) Xing Yang (Huawei), Saad Ali (Google)
Kubernetes v1.12 introduces alpha support for volume snapshotting. This feature allows creating/deleting volume snapshots, and the ability to create new volumes from a snapshot natively using the Kubernetes API.
Many storage systems (like Google Cloud Persistent Disks, Amazon Elastic Block Storage, and many on-premise storage systems) provide the ability to create a &#34;snapshot&#34; of a persistent volume. A snapshot represents a point-in-time copy of a volume. A snapshot can be used either to provision a new volume (pre-populated with the snapshot data) or to restore the existing volume to a previous state (represented by the snapshot).
The Kubernetes volume plugin system already provides a powerful abstraction that automates the provisioning, attaching, and mounting of block and file storage.
Underpinning all these features is the Kubernetes goal of workload portability: Kubernetes aims to create an abstraction layer between distributed systems applications and underlying clusters so that applications can be agnostic to the specifics of the cluster they run on and application deployment requires no “cluster specific” knowledge.
The [Kubernetes Storage SIG](https://github.com/kubernetes/community/tree/master/sig-storage) identified snapshot operations as critical functionality for many stateful workloads. For example, a database administrator may want to snapshot a database volume before starting a database operation.
By providing a standard way to trigger snapshot operations in the Kubernetes API, Kubernetes users can now handle use cases like this without having to go around the Kubernetes API (and manually executing storage system specific operations).
Instead, Kubernetes users are now empowered to incorporate snapshot operations in a cluster agnostic way into their tooling and policy with the comfort of knowing that it will work against arbitrary Kubernetes clusters regardless of the underlying storage.
Additionally these Kubernetes snapshot primitives act as basic building blocks that unlock the ability to develop advanced, enterprise grade, storage administration features for Kubernetes: such as data protection, data replication, and data migration.
Kubernetes supports three types of volume plugins: in-tree, Flex, and CSI. See [Kubernetes Volume Plugin FAQ](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md) for details.
Snapshots are only supported for CSI drivers (not for in-tree or Flex). To use the Kubernetes snapshots feature, ensure that a CSI Driver that implements snapshots is deployed on your cluster.
As of the publishing of this blog, the following CSI drivers support snapshots:
Snapshot support for other [drivers](https://kubernetes-csi.github.io/docs/drivers.html) is pending, and should be available soon. Read the “[Container Storage Interface (CSI) for Kubernetes Goes Beta](https://kubernetes.io/blog/2018/04/10/container-storage-interface-beta/)” blog post to learn more about CSI and how to deploy CSI drivers.
Similar to the API for managing Kubernetes Persistent Volumes, Kubernetes Volume Snapshots introduce three new API objects for managing snapshots:
It is important to note that unlike the core Kubernetes Persistent Volume objects, these Snapshot objects are defined as [CustomResourceDefinitions (CRDs)](/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions). The Kubernetes project is moving away from having resource types pre-defined in the API server, and is moving towards a model where the API server is independent of the API objects. This allows the API server to be reused for projects other than Kubernetes, and consumers (like Kubernetes) can simply install the resource types they require as CRDs.
[CSI Drivers](https://kubernetes-csi.github.io/docs/drivers.html) that support snapshots will automatically install the required CRDs. Kubernetes end users only need to verify that a CSI driver that supports snapshots is deployed on their Kubernetes cluster.
In addition to these new objects, a new, DataSource field has been added to the ```PersistentVolumeClaim``` object:
```type PersistentVolumeClaimSpec struct {
	AccessModes []PersistentVolumeAccessMode
	Selector *metav1.LabelSelector
	Resources ResourceRequirements
	VolumeName string
	StorageClassName *string
	VolumeMode *PersistentVolumeMode
	DataSource *TypedLocalObjectReference
}
```This new alpha field enables a new volume to be created and automatically pre-populated with data from an existing snapshot.
Before using Kubernetes Volume Snapshotting, you must:
Before creating a snapshot, you also need to specify CSI driver information for snapshots by creating a ```VolumeSnapshotClass``` object and setting the ```snapshotter``` field to point to your CSI driver. In the example of ```VolumeSnapshotClass``` below, the CSI driver is ```com.example.csi-driver```. You need at least one ```VolumeSnapshotClass``` object per snapshot provisioner. You can also set a default ```VolumeSnapshotClass``` for each individual CSI driver by putting an annotation ```snapshot.storage.kubernetes.io/is-default-class: &#34;true&#34;``` in the class definition.
```apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshotClass
metadata:
  name: default-snapclass
  annotations:
    snapshot.storage.kubernetes.io/is-default-class: &#34;true&#34;
snapshotter: com.example.csi-driver


apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshotClass
metadata:
  name: csi-snapclass
snapshotter: com.example.csi-driver
parameters:
  fakeSnapshotOption: foo
  csiSnapshotterSecretName: csi-secret
  csiSnapshotterSecretNamespace: csi-namespace
```You must set any required opaque parameters based on the documentation for your CSI driver. As the example above shows,  the parameter ```fakeSnapshotOption: foo``` and any referenced secret(s) will be passed to CSI driver during snapshot creation and deletion. The [default CSI external-snapshotter](https://github.com/kubernetes-csi/external-snapshotter) reserves the parameter keys ```csiSnapshotterSecretName``` and ```csiSnapshotterSecretNamespace```. If specified, it fetches the secret and passes it to the CSI driver when creating and deleting a snapshot.
And finally, before creating a snapshot, you must provision a volume using your CSI driver and populate it with some data that you want to snapshot (see the [CSI blog post](https://kubernetes.io/blog/2018/04/10/container-storage-interface-beta/) on how to create and use CSI volumes).
Once a ```VolumeSnapshotClass``` object is defined and you have a volume you want to snapshot, you may create a new snapshot by creating a ```VolumeSnapshot``` object.
The source of the snapshot specifies the volume to create a snapshot from. It has two parameters:
The namespace of the volume to snapshot is assumed to be the same as the namespace of the ```VolumeSnapshot``` object.
```apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: new-snapshot-demo
  namespace: demo-namespace
spec:
  snapshotClassName: csi-snapclass
  source:
    name: mypvc
    kind: PersistentVolumeClaim
```In the ```VolumeSnapshot``` spec, user can specify the ```VolumeSnapshotClass``` which has the information about which CSI driver should be used for creating the snapshot . When the ```VolumeSnapshot``` object is created, the parameter ```fakeSnapshotOption: foo``` and any referenced secret(s) from the ```VolumeSnapshotClass``` are passed to the CSI plugin ```com.example.csi-driver``` via a ```CreateSnapshot``` call.
In response, the CSI driver triggers a snapshot of the volume and then automatically creates a ```VolumeSnapshotContent``` object to represent the new snapshot, and binds the new ```VolumeSnapshotContent``` object to the ```VolumeSnapshot```, making it ready to use. If the  CSI driver fails to create the snapshot and returns error, the snapshot controller reports the error in the status of ```VolumeSnapshot``` object and does not retry (this is different from other controllers in Kubernetes, and is to prevent snapshots from being taken at an unexpected time).
If a snapshot class is not specified, the external snapshotter will try to find and set a default snapshot class for the snapshot. The ```CSI driver``` specified by ```snapshotter``` in the default snapshot class must match the ```CSI driver``` specified by the ```provisioner``` in the storage class of the PVC.
Please note that the alpha release of Kubernetes Snapshot does not provide any consistency guarantees. You have to prepare your application (pause application, freeze filesystem etc.) before taking the snapshot for data consistency.
You can verify that the ```VolumeSnapshot``` object is created and bound with ```VolumeSnapshotContent``` by running ```kubectl describe volumesnapshot```:
You can always import an existing snapshot to Kubernetes by manually creating a ```VolumeSnapshotContent``` object to represent the existing snapshot. Because ```VolumeSnapshotContent``` is a non-namespace API object, only a system admin may have the permission to create it. Once a ```VolumeSnapshotContent``` object is created, the user can create a ```VolumeSnapshot``` object pointing to the ```VolumeSnapshotContent``` object. The external-snapshotter controller will mark snapshot as ready after verifying the snapshot exists and the binding between ```VolumeSnapshot``` and ```VolumeSnapshotContent``` objects is correct. Once bound, the snapshot is ready to use in Kubernetes.
A ```VolumeSnapshotContent``` object should be created with the following fields to represent a pre-provisioned snapshot:
```apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshotContent
metadata:
  name: static-snapshot-content
spec:
  csiVolumeSnapshotSource:
    driver: com.example.csi-driver
    snapshotHandle: snapshotcontent-example-id
  volumeSnapshotRef:
    kind: VolumeSnapshot
    name: static-snapshot-demo
    namespace: demo-namespace
```A ```VolumeSnapshot``` object should be created to allow a user to use the snapshot:
```apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshot
metadata:
  name: static-snapshot-demo
  namespace: demo-namespace
spec:
  snapshotClassName: csi-snapclass
  snapshotContentName: static-snapshot-content
```Once these objects are created, the snapshot controller will bind them together, and set the field Ready (under ```Status```) to True to indicate the snapshot is ready to use.
To provision a new volume pre-populated with data from a snapshot object, use the new dataSource field in the ```PersistentVolumeClaim```. It has three parameters:
The namespace of the source ```VolumeSnapshot``` object is assumed to be the same as the namespace of the ```PersistentVolumeClaim``` object.
```apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-restore
  Namespace: demo-namespace
spec:
  storageClassName: csi-storageclass
  dataSource:
    name: new-snapshot-demo
    kind: VolumeSnapshot
    apiGroup: snapshot.storage.k8s.io
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```When the ```PersistentVolumeClaim``` object is created, it will trigger provisioning of a new volume that is pre-populated with data from the specified snapshot.
To implement the snapshot feature, a CSI driver MUST add support for additional controller capabilities ```CREATE_DELETE_SNAPSHOT``` and ```LIST_SNAPSHOTS```, and implement additional controller RPCs: ```CreateSnapshot```, ```DeleteSnapshot```, and ```ListSnapshots```. For details, see [the CSI spec](https://github.com/container-storage-interface/spec/blob/master/spec.md).
Although Kubernetes is as [minimally prescriptive](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md#third-party-csi-volume-drivers) on the packaging and deployment of a CSI Volume Driver as possible, it provides a [suggested mechanism](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/container-storage-interface.md#recommended-mechanism-for-deploying-csi-drivers-on-kubernetes) for deploying an arbitrary containerized CSI driver on Kubernetes to simplify deployment of containerized CSI compatible volume drivers.
As part of this recommended deployment process, the Kubernetes team provides a number of sidecar (helper) containers, including a new [external-snapshotter](https://github.com/kubernetes-csi/external-snapshotter) sidecar container.
The external-snapshotter watches the Kubernetes API server for ```VolumeSnapshot``` and ```VolumeSnapshotContent``` objects and triggers CreateSnapshot and DeleteSnapshot operations against a CSI endpoint. The CSI [external-provisioner](https://github.com/kubernetes-csi/external-provisioner) sidecar container has also been updated to support restoring volume from snapshot using the new ```dataSource``` PVC field.
In order to support snapshot feature, it is recommended that storage vendors deploy the external-snapshotter sidecar containers in addition to the external provisioner the external attacher, along with their CSI driver in a statefulset as shown in the following diagram.

In this [example deployment yaml](https://github.com/kubernetes-csi/external-snapshotter/blob/e011fe31df548813d2eb6dacb278c0ca58533b34/deploy/kubernetes/setup-csi-snapshotter.yaml) file, two sidecar containers, the external provisioner and the external snapshotter, and CSI drivers are deployed together with the hostpath CSI plugin in the statefulset pod. Hostpath CSI plugin is a sample plugin, not for production.
The alpha implementation of snapshots for Kubernetes has the following limitations:
Depending on feedback and adoption, the Kubernetes team plans to push the CSI Snapshot implementation to beta in either 1.13 or 1.14.
Check out additional documentation on the snapshot feature here: [http://k8s.io/docs/concepts/storage/volume-snapshots](http://k8s.io/docs/concepts/storage/volume-snapshots) and [https://kubernetes-csi.github.io/docs/](https://kubernetes-csi.github.io/docs/)
This project, like all of Kubernetes, is the result of hard work by many contributors from diverse backgrounds working together.
In addition to the contributors who have been working on the Snapshot feature:
We offer a huge thank you to all the contributors in Kubernetes Storage SIG and CSI community who helped review the design and implementation of the project, including but not limited to the following:
If you’re interested in getting involved with the design and development of CSI or any part of the Kubernetes Storage system, join the [Kubernetes Storage Special Interest Group](https://github.com/kubernetes/community/tree/master/sig-storage) (SIG). We’re rapidly growing and always welcome new contributors.


	

	


