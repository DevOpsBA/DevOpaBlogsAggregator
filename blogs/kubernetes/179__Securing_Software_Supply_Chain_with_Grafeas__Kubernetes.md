|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2017/11/Securing-Software-Supply-Chain-Grafeas/        |
| Tags              | [kubernetes]       |
| Date Create       | 2017-11-03 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:21.251022 &#43;0300 MSK m=&#43;2.840060701  |

#  Securing Software Supply Chain with Grafeas  | Kubernetes

	
	
	
	
	***Editor&#39;s note: This post is written by Kelsey Hightower, Staff Developer Advocate at Google, and Sandra Guo, Product Manager at Google.***
Kubernetes has evolved to support increasingly complex classes of applications, enabling the development of two major industry trends: hybrid cloud and microservices. With increasing complexity in production environments, customers—especially enterprises—are demanding better ways to manage their software supply chain with more centralized visibility and control over production deployments.
On October 12th, Google and partners [announced](https://cloudplatform.googleblog.com/2017/10/introducing-grafeas-open-source-api-.html) Grafeas, an open source initiative to define a best practice for auditing and governing the modern software supply chain. With Grafeas (“scribe” in Greek), developers can plug in components of the CI/CD pipeline into a central source of truth for tracking and enforcing policies. Google is also working on [Kritis](https://github.com/Grafeas/Grafeas/blob/master/case-studies/binary-authorization.md) (“judge” in Greek), allowing devOps teams to enforce deploy-time image policy using metadata and attestations stored in Grafeas.
Grafeas allows build, auditing and compliance tools to exchange comprehensive metadata on container images using a central API. This allows enforcing policies that provide central control over the software supply process.
[img](https://2.bp.blogspot.com/-TDD4slMA7gg/WfzDeKVLr2I/AAAAAAAAAGw/dhfWOrCMdmogSNhGr5RrA2ovr02K5nn8ACK4BGAYYCw/s1600/Screen%2BShot%2B2017-11-03%2Bat%2B12.28.13%2BPM.png)
Let’s consider a simple application, *PaymentProcessor*, that retrieves, processes and updates payment info stored in a database. This application is made up of two containers: a standard ruby container and custom logic.
Due to the sensitive nature of the payment data, the developers and DevOps team really want to make sure that the code meets certain security and compliance requirements, with detailed records on the provenance of this code. There are CI/CD stages that validate the quality of the PaymentProcessor release, but there is no easy way to centrally view/manage this information:
[img](https://1.bp.blogspot.com/-WeI6zpGd42A/WfzDkkIonFI/AAAAAAAAAG4/wKUaNaXYvaQ-an9p4_9T9J3EQB_zHkRXwCK4BGAYYCw/s1600/Screen%2BShot%2B2017-11-03%2Bat%2B12.28.23%2BPM.png)
Grafeas provides an API for customers to centrally manage metadata created by various CI/CD components and enables deploy time policy enforcement through a Kritis implementation.
[img](https://4.bp.blogspot.com/-SRMfm5z606M/WfzDpHqlz-I/AAAAAAAAAHA/y2suaInhr9E0hU0u78PacBT_kZj2D7DKgCK4BGAYYCw/s1600/Screen%2BShot%2B2017-11-03%2Bat%2B12.28.34%2BPM.png)
Let’s consider a basic example of how Grafeas can provide deploy time control for the PaymentProcessor app using a demo verification pipeline.
Assume that a PaymentProcessor container image has been created and pushed to Google Container Registry. This example uses the gcr.io/exampleApp/PaymentProcessor container for testing. You as the QA engineer want to create an attestation certifying this image for production usage. Instead of trusting an image tag like 0.0.1, which can be reused and point to a different container image later, we can trust the image digest to ensure the attestation links to the full image contents.
**1. Set up the environment**
Generate a signing key:
```gpg --quick-generate-key --yes qa\_bob@example.com
```Export the image signer&#39;s public key:
```gpg --armor --export image.signer@example.com \&gt; ${GPG\_KEY\_ID}.pub
```Create the ‘qa’ AttestationAuthority note via the Grafeas API:
```curl -X POST \  
  &#34;http://127.0.0.1:8080/v1alpha1/projects/image-signing/notes?noteId=qa&#34; \  
  -d @note.json
```Create the Kubernetes ConfigMap for admissions control and store the QA signer&#39;s public key:
```kubectl create configmap image-signature-webhook \  
  --from-file ${GPG\_KEY\_ID}.pub

kubectl get configmap image-signature-webhook -o yaml
```Set up an admissions control webhook to require QA signature during deployment.
```kubectl apply -f kubernetes/image-signature-webhook.yaml
```**2. Attempt to deploy an image without QA attestation**
Attempt to run the image in paymentProcessor.ymal before it is QA attested:
```kubectl apply -f pods/nginx.yaml

apiVersion: v1

kind: Pod

metadata:

  name: payment

spec:

  containers:

    - name: payment

      image: &#34;gcr.io/hightowerlabs/payment@sha256:aba48d60ba4410ec921f9d2e8169236c57660d121f9430dc9758d754eec8f887&#34;
```Create the paymentProcessor pod:
```kubectl apply -f pods/paymentProcessor.yaml
```Notice the paymentProcessor pod was not created and the following error was returned:
```The  &#34;&#34; is invalid: : No matched signatures for container image: gcr.io/hightowerlabs/payment@sha256:aba48d60ba4410ec921f9d2e8169236c57660d121f9430dc9758d754eec8f887
```**3. Create an image signature**
Assume the image digest is stored in Image-digest.txt, sign the image digest:
```gpg -u qa\_bob@example.com \  
  --armor \  
  --clearsign \  
  --output=signature.gpg \  
  Image-digest.txt
```**4. Upload the signature to the Grafeas API**
Generate a pgpSignedAttestation occurrence from the signature :
```cat \&gt; occurrence.json \&lt;\&lt;EOF  
{  
  &#34;resourceUrl&#34;: &#34;$(cat image-digest.txt)&#34;,  
  &#34;noteName&#34;: &#34;projects/image-signing/notes/qa&#34;,  
  &#34;attestation&#34;: {  
    &#34;pgpSignedAttestation&#34;: {  
       &#34;signature&#34;: &#34;$(cat signature.gpg)&#34;,  
       &#34;contentType&#34;: &#34;application/vnd.gcr.image.url.v1&#34;,  
       &#34;pgpKeyId&#34;: &#34;${GPG\_KEY\_ID}&#34;  
    }  
  }  
}  
EOF
```Upload the attestation through the Grafeas API:
```curl -X POST \  
  &#39;http://127.0.0.1:8080/v1alpha1/projects/image-signing/occurrences&#39; \  
  -d @occurrence.json
```**5. Verify QA attestation during a production deployment**
Attempt to run the image in paymentProcessor.ymal now that it has the correct attestation in the Grafeas API:
```kubectl apply -f pods/paymentProcessor.yaml

pod &#34;PaymentProcessor&#34; created
```With the attestation added the pod will be created as the execution criteria are met.
For more detailed information, see this [Grafeas tutorial](https://github.com/kelseyhightower/grafeas-tutorial).
The demo above showed how you can integrate your software supply chain with Grafeas and gain visibility and control over your production deployments. However, the demo verification pipeline by itself is not a full Kritis implementation. In addition to basic admission control, Kritis provides additional support for workflow enforcement, multi-authority signing, breakglass deployment and more. You can read the [Kritis whitepaper](https://github.com/Grafeas/Grafeas/blob/master/case-studies/binary-authorization.md) for more details. The team is actively working on a full open-source implementation. We’d love your feedback!
In addition, a hosted alpha implementation of Kritis, called Binary Authorization, is available on Google Container Engine and will be available for broader consumption soon.
Google, JFrog, and other partners joined forces to create Grafeas based on our common experiences building secure, large, and complex microservice deployments for internal and enterprise customers. Grafeas is an industry-wide community effort.
To learn more about Grafeas and contribute to the project:


	

	


