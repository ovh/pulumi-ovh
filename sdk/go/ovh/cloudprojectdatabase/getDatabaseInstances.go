// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of databases of a database cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloudprojectdatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			databases, err := cloudprojectdatabase.GetDatabaseInstances(ctx, &cloudprojectdatabase.GetDatabaseInstancesArgs{
//				ServiceName: "XXXX",
//				Engine:      "YYYY",
//				ClusterId:   "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("databaseIds", databases.DatabaseIds)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseInstances(ctx *pulumi.Context, args *GetDatabaseInstancesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabaseInstancesResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabaseInstances:getDatabaseInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstances.
type GetDatabaseInstancesArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to list databases. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabaseInstances.
type GetDatabaseInstancesResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// The list of databases ids of the database cluster associated with the project.
	DatabaseIds []string `pulumi:"databaseIds"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetDatabaseInstancesOutput(ctx *pulumi.Context, args GetDatabaseInstancesOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseInstancesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabaseInstancesResultOutput, error) {
			args := v.(GetDatabaseInstancesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getDatabaseInstances:getDatabaseInstances", args, GetDatabaseInstancesResultOutput{}, options).(GetDatabaseInstancesResultOutput), nil
		}).(GetDatabaseInstancesResultOutput)
}

// A collection of arguments for invoking getDatabaseInstances.
type GetDatabaseInstancesOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want to list databases. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabaseInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseInstances.
type GetDatabaseInstancesResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseInstancesResult)(nil)).Elem()
}

func (o GetDatabaseInstancesResultOutput) ToGetDatabaseInstancesResultOutput() GetDatabaseInstancesResultOutput {
	return o
}

func (o GetDatabaseInstancesResultOutput) ToGetDatabaseInstancesResultOutputWithContext(ctx context.Context) GetDatabaseInstancesResultOutput {
	return o
}

// See Argument Reference above.
func (o GetDatabaseInstancesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstancesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The list of databases ids of the database cluster associated with the project.
func (o GetDatabaseInstancesResultOutput) DatabaseIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabaseInstancesResult) []string { return v.DatabaseIds }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetDatabaseInstancesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstancesResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabaseInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseInstancesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstancesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseInstancesResultOutput{})
}
