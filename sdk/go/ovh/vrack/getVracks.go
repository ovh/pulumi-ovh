// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of Vrack IDs available for your OVHcloud account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/vrack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vrack.GetVracks(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVracks(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetVracksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVracksResult
	err := ctx.Invoke("ovh:Vrack/getVracks:getVracks", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getVracks.
type GetVracksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of vrack service name available for your OVHcloud account.
	Results []string `pulumi:"results"`
}

func GetVracksOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetVracksResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetVracksResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ovh:Vrack/getVracks:getVracks", nil, GetVracksResultOutput{}, options).(GetVracksResultOutput), nil
	}).(GetVracksResultOutput)
}

// A collection of values returned by getVracks.
type GetVracksResultOutput struct{ *pulumi.OutputState }

func (GetVracksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVracksResult)(nil)).Elem()
}

func (o GetVracksResultOutput) ToGetVracksResultOutput() GetVracksResultOutput {
	return o
}

func (o GetVracksResultOutput) ToGetVracksResultOutputWithContext(ctx context.Context) GetVracksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVracksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVracksResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of vrack service name available for your OVHcloud account.
func (o GetVracksResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVracksResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVracksResultOutput{})
}
