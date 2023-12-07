package main

import (
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		serviceName := config.Require(ctx, "serviceName")

		// Deploy a new Kubernetes cluster
		myKube, err := cloudproject.NewKube(ctx, "my_desired_cluster", &cloudproject.KubeArgs{
			ServiceName: pulumi.String(serviceName),
			Name:        pulumi.String("my_desired_cluster"),
			Region:      pulumi.String("GRA5"),
		})
		if err != nil {
			return err
		}

		// Export kubeconfig file to a secret
		ctx.Export("kubeconfig", pulumi.ToSecret(myKube.Kubeconfig))

		//Create a Node Pool
		nodePool, err := cloudproject.NewKubeNodePool(ctx, "my-desired-pool", &cloudproject.KubeNodePoolArgs{
			ServiceName:  pulumi.String(serviceName),
			KubeId:       myKube.ID(),
			Name:         pulumi.String("my-desired-pool"),
			DesiredNodes: pulumi.Int(1),
			MaxNodes:     pulumi.Int(3),
			MinNodes:     pulumi.Int(1),
			FlavorName:   pulumi.String("b2-7"),
		})
		if err != nil {
			return err
		}

		ctx.Export("nodePoolID", nodePool.ID())

		return nil
	})
}
