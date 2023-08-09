package main

import (
	"github.com/dirien/pulumi-ovh/sdk/go/ovh/cloudproject"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mykube, err := cloudproject.NewKube(ctx, "mykube", &cloudproject.KubeArgs{
			ServiceName: pulumi.String("serviceName-dirien"),
			Name:        pulumi.String("my_desired_cluster"),
			Region:      pulumi.String("GRA5"),
		})
		if err != nil {
			return err
		}
		ctx.Export("name", mykube.Kubeconfig)
		return nil
	})
}
