// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to list the existing IAM policies of an account.
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
//			_, err := iam.GetPolicies(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPolicies(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPoliciesResult
	err := ctx.Invoke("ovh:Iam/getPolicies:getPolicies", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of the policies IDs.
	Policies []string `pulumi:"policies"`
}

func GetPoliciesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetPoliciesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetPoliciesResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ovh:Iam/getPolicies:getPolicies", nil, GetPoliciesResultOutput{}, options).(GetPoliciesResultOutput), nil
	}).(GetPoliciesResultOutput)
}

// A collection of values returned by getPolicies.
type GetPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesResult)(nil)).Elem()
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutput() GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutputWithContext(ctx context.Context) GetPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of the policies IDs.
func (o GetPoliciesResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoliciesResultOutput{})
}
