// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vps

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of VPS associated with your OVH Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Vps"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vps.GetVpss(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVpss(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetVpssResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpssResult
	err := ctx.Invoke("ovh:Vps/getVpss:getVpss", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getVpss.
type GetVpssResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of VPS IDs associated with your OVH Account.
	Results []string `pulumi:"results"`
}

func GetVpssOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetVpssResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetVpssResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetVpssResult
		secret, err := ctx.InvokePackageRaw("ovh:Vps/getVpss:getVpss", nil, &rv, "", opts...)
		if err != nil {
			return GetVpssResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetVpssResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetVpssResultOutput), nil
		}
		return output, nil
	}).(GetVpssResultOutput)
}

// A collection of values returned by getVpss.
type GetVpssResultOutput struct{ *pulumi.OutputState }

func (GetVpssResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpssResult)(nil)).Elem()
}

func (o GetVpssResultOutput) ToGetVpssResultOutput() GetVpssResultOutput {
	return o
}

func (o GetVpssResultOutput) ToGetVpssResultOutputWithContext(ctx context.Context) GetVpssResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpssResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpssResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of VPS IDs associated with your OVH Account.
func (o GetVpssResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpssResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpssResultOutput{})
}
