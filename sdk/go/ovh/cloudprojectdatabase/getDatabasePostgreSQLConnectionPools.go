// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.
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
//			testPools, err := cloudprojectdatabase.GetDatabasePostgreSQLConnectionPools(ctx, &cloudprojectdatabase.GetDatabasePostgreSQLConnectionPoolsArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("connectionPoolIds", testPools.ConnectionPoolIds)
//			return nil
//		})
//	}
//
// ```
func GetDatabasePostgreSQLConnectionPools(ctx *pulumi.Context, args *GetDatabasePostgreSQLConnectionPoolsArgs, opts ...pulumi.InvokeOption) (*GetDatabasePostgreSQLConnectionPoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasePostgreSQLConnectionPoolsResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabasePostgreSQLConnectionPools.
type GetDatabasePostgreSQLConnectionPoolsArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabasePostgreSQLConnectionPools.
type GetDatabasePostgreSQLConnectionPoolsResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// The list of patterns ids of the opensearch cluster associated with the project.
	ConnectionPoolIds []string `pulumi:"connectionPoolIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetDatabasePostgreSQLConnectionPoolsOutput(ctx *pulumi.Context, args GetDatabasePostgreSQLConnectionPoolsOutputArgs, opts ...pulumi.InvokeOption) GetDatabasePostgreSQLConnectionPoolsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabasePostgreSQLConnectionPoolsResultOutput, error) {
			args := v.(GetDatabasePostgreSQLConnectionPoolsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", args, GetDatabasePostgreSQLConnectionPoolsResultOutput{}, options).(GetDatabasePostgreSQLConnectionPoolsResultOutput), nil
		}).(GetDatabasePostgreSQLConnectionPoolsResultOutput)
}

// A collection of arguments for invoking getDatabasePostgreSQLConnectionPools.
type GetDatabasePostgreSQLConnectionPoolsOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabasePostgreSQLConnectionPoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasePostgreSQLConnectionPoolsArgs)(nil)).Elem()
}

// A collection of values returned by getDatabasePostgreSQLConnectionPools.
type GetDatabasePostgreSQLConnectionPoolsResultOutput struct{ *pulumi.OutputState }

func (GetDatabasePostgreSQLConnectionPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasePostgreSQLConnectionPoolsResult)(nil)).Elem()
}

func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) ToGetDatabasePostgreSQLConnectionPoolsResultOutput() GetDatabasePostgreSQLConnectionPoolsResultOutput {
	return o
}

func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) ToGetDatabasePostgreSQLConnectionPoolsResultOutputWithContext(ctx context.Context) GetDatabasePostgreSQLConnectionPoolsResultOutput {
	return o
}

// See Argument Reference above.
func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasePostgreSQLConnectionPoolsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The list of patterns ids of the opensearch cluster associated with the project.
func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) ConnectionPoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabasePostgreSQLConnectionPoolsResult) []string { return v.ConnectionPoolIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasePostgreSQLConnectionPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabasePostgreSQLConnectionPoolsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasePostgreSQLConnectionPoolsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasePostgreSQLConnectionPoolsResultOutput{})
}
