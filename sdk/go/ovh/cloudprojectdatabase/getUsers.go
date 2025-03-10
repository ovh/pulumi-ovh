// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	ClusterId   string `pulumi:"clusterId"`
	Engine      string `pulumi:"engine"`
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	ClusterId string `pulumi:"clusterId"`
	Engine    string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	ServiceName string   `pulumi:"serviceName"`
	UserIds     []string `pulumi:"userIds"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetUsersResultOutput, error) {
			args := v.(GetUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getUsers:getUsers", args, GetUsersResultOutput{}, options).(GetUsersResultOutput), nil
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	ClusterId   pulumi.StringInput `pulumi:"clusterId"`
	Engine      pulumi.StringInput `pulumi:"engine"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.UserIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
