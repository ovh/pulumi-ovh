---
title: OVH
meta_desc: Provides an overview of the OVH Provider for Pulumi.
layout: overview
---

The `OVH` provider for Pulumi can be used to provision any of the resources available in [OVHcloud](https://www.ovhcloud.com/fr/).
The `OVH` provider must be configured with credentials to deploy and update resources in `OVH`.

## Example

{{< chooser language "go,typescript,python" >}}
{{% choosable language go %}}

```golang
// Deploy a new Kubernetes cluster
myKube, err := cloudproject.NewKube(ctx, "my_desired_cluster", &cloudproject.KubeArgs{
    ServiceName: pulumi.String("xxxxxxxxxxxxx-xxxx-xxxx-xxxxxxxxx"),
    Name:        pulumi.String("my_desired_cluster"),
    Region:      pulumi.String("GRA5"),
})
if err != nil {
    return err
}

// Export kubeconfig file to a secret
ctx.Export("kubeconfig", pulumi.ToSecret(myKube.Kubeconfig))
```

{{% choosable language python %}}

```python
import pulumi
import scraly_pulumi_ovh as ovh

config = pulumi.Config();
service_name = config.require('serviceName')

print(service_name);

my_kube_cluster = cloudproject.get_kube(service_name,
    kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx");
pulumi.export("version", my_kube_cluster.version)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
TODO
```

{{% /choosable %}}

{{< /chooser >}}

