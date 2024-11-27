# OVHcloud Pulumi templates

## Use a template to deploy an OVHcloud MKS in Go

Create a project folder and go into it:

```
mkdir test-template
cd test-template
```

Define the OVHcloud credentials

```
export OVH_ENDPOINT="ovh-eu" #or ovh-ca or ovh-us
export OVH_APPLICATION_KEY="xxx"
export OVH_APPLICATION_SECRET="xxx"
export OVH_CONSUMER_KEY="xxx"
```

Use the template to deploy your MKS:
```
pulumi new ../kubernetes-ovh-go
pulumi preview
pulumi up
```

## Python

Use the template to deploy your MKS:
```
pulumi new ../kubernetes-ovh-python
pulumi preview
pulumi up
```

## TypeScript

Use the template to deploy your MKS:
```
pulumi new ../kubernetes-ovh-typescript
pulumi preview
pulumi up
```