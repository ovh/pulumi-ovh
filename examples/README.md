# Examples

The following examples demonstrate various ways to create and work with [OVH](https://www.ovhcloud.com/fr/) Cloud provider.

## Managed Kubernetes

1. [Create an OVHcloud Managed Kubernetes cluster and his Node Pool in Golang/Go](./kubernetes/ovh-go/)
1. [Create an OVHcloud MKS and his Node Pool in Java](./kubernetes/ovh-java/)
1. [Get a an OVHcloud MKS cluster version in Python](./kubernetes/ovh-python)
1. [Get a an OVHcloud MKS cluster version in Typescript](./kubernetes/ovh-typescript)
1. [Get a an OVHcloud MKS cluster version in C#/Dotnet](./kubernetes/ovh-csharp)

Import an existing MKS cluster:
```bash
$ pulumi import ovh:CloudProject/kube:Kube my_kube_cluster_name <service_name>/xxxxx-xxx-xxx-xxx-xxxxxxxx
```

## Managed Private Registry

1. [Create an OVHcloud MPR in Golang/Go](./registry/ovh-registry-go/)

## Managed Databases

1. [Create an OVHcloud Managed MySQL Database in Golang/Go](./databases/ovh-db-go/)
