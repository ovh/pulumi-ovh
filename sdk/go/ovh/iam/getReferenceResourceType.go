// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to list all the IAM resource types.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.GetReferenceResourceType(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetReferenceResourceType(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetReferenceResourceTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReferenceResourceTypeResult
	err := ctx.Invoke("ovh:Iam/getReferenceResourceType:getReferenceResourceType", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getReferenceResourceType.
type GetReferenceResourceTypeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of resource types
	Types []string `pulumi:"types"`
}

func GetReferenceResourceTypeOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetReferenceResourceTypeResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetReferenceResourceTypeResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ovh:Iam/getReferenceResourceType:getReferenceResourceType", nil, GetReferenceResourceTypeResultOutput{}, options).(GetReferenceResourceTypeResultOutput), nil
	}).(GetReferenceResourceTypeResultOutput)
}

// A collection of values returned by getReferenceResourceType.
type GetReferenceResourceTypeResultOutput struct{ *pulumi.OutputState }

func (GetReferenceResourceTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReferenceResourceTypeResult)(nil)).Elem()
}

func (o GetReferenceResourceTypeResultOutput) ToGetReferenceResourceTypeResultOutput() GetReferenceResourceTypeResultOutput {
	return o
}

func (o GetReferenceResourceTypeResultOutput) ToGetReferenceResourceTypeResultOutputWithContext(ctx context.Context) GetReferenceResourceTypeResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetReferenceResourceTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReferenceResourceTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of resource types
func (o GetReferenceResourceTypeResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetReferenceResourceTypeResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReferenceResourceTypeResultOutput{})
}
