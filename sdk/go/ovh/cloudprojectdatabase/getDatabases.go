// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	Engine      string `pulumi:"engine"`
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	ClusterIds []string `pulumi:"clusterIds"`
	Engine     string   `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ServiceName string `pulumi:"serviceName"`
}

func GetDatabasesOutput(ctx *pulumi.Context, args GetDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabasesResultOutput, error) {
			args := v.(GetDatabasesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getDatabases:getDatabases", args, GetDatabasesResultOutput{}, options).(GetDatabasesResultOutput), nil
		}).(GetDatabasesResultOutput)
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesOutputArgs struct {
	Engine      pulumi.StringInput `pulumi:"engine"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o GetDatabasesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}
