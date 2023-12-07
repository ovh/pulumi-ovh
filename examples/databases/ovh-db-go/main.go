package main

import (
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		serviceName := config.Require(ctx, "serviceName")

		myMysqlDB, err := cloudproject.NewDatabase(ctx, "mysqldb", &cloudproject.DatabaseArgs{
			ServiceName: pulumi.String(serviceName),
			Description: pulumi.String("my-first-mysql"),
			Engine:      pulumi.String("mysql"),
			Version:     pulumi.String("8"),
			Plan:        pulumi.String("essential"),
			Nodes: cloudproject.DatabaseNodeArray{
				&cloudproject.DatabaseNodeArgs{
					Region: pulumi.String("SBG"),
				},
			},
			Flavor: pulumi.String("db1-4"),
			AdvancedConfiguration: pulumi.StringMap{
				"mysql.sql_mode":                pulumi.String("ANSI,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE,STRICT_ALL_TABLES"),
				"mysql.sql_require_primary_key": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("myMysqlDBID", myMysqlDB.ID())
		ctx.Export("myMysqlDBEndpoints", myMysqlDB.Endpoints)

		return nil
	})
}
