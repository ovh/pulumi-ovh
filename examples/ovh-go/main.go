package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/cloudproject"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mykube, err := cloudproject.NewKube(ctx, "mykube", &cloudproject.KubeArgs{
			ServiceName: pulumi.String("serviceName-scraly"),
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
