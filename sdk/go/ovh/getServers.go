// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of dedicated servers associated with your OVHcloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.GetServers(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServersResult
	err := ctx.Invoke("ovh:index/getServers:getServers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getServers.
type GetServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of dedicated servers IDs associated with your OVHcloud Account.
	Results []string `pulumi:"results"`
}

func GetServersOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetServersResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetServersResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ovh:index/getServers:getServers", nil, GetServersResultOutput{}, options).(GetServersResultOutput), nil
	}).(GetServersResultOutput)
}

// A collection of values returned by getServers.
type GetServersResultOutput struct{ *pulumi.OutputState }

func (GetServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServersResult)(nil)).Elem()
}

func (o GetServersResultOutput) ToGetServersResultOutput() GetServersResultOutput {
	return o
}

func (o GetServersResultOutput) ToGetServersResultOutputWithContext(ctx context.Context) GetServersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of dedicated servers IDs associated with your OVHcloud Account.
func (o GetServersResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServersResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServersResultOutput{})
}
