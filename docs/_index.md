---
title: OVH
meta_desc: Provides an overview of the OVH Provider for Pulumi.
layout: overview
---

The `OVH` provider for Pulumi can be used to provision any of the resources available in [OVHcloud](https://www.ovhcloud.com/fr/).
The `OVH` provider must be configured with credentials to deploy and update resources in `OVH`.

## Example

TODO: goooooooooo!!!



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

{{% choosable language typescript %}}

```typescript
TODO
```

{{% /choosable %}}
{{% choosable language python %}}

```python
TODO
```

{{% /choosable %}}
{{< /chooser >}}

