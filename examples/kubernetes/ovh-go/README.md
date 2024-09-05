# ovh-go

This example deploys an OVHcloud Managed Kubernetes cluster in Go.

## Update the provider version

* edit the `go.mod` file and change the version of the github.com/ovh/pulumi-ovh/sdk dependency
* `go mod tidy`
* `pulumi config set serviceName <your_ovh_service_name>`
* `pulumi login --local`
* `pulumi up`