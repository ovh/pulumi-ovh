// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source get details about a resource group.
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
//			_, err := iam.GetResourceGroup(ctx, &iam.GetResourceGroupArgs{
//				Id: "my_resource_group_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupResourceGroup(ctx *pulumi.Context, args *LookupResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupResourceGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceGroupResult
	err := ctx.Invoke("ovh:Iam/getResourceGroup:getResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceGroup.
type LookupResourceGroupArgs struct {
	// Id of the resource group
	Id string `pulumi:"id"`
}

// A collection of values returned by getResourceGroup.
type LookupResourceGroupResult struct {
	// URN of the resource group, used when writing policies
	GroupURN string `pulumi:"GroupURN"`
	// Date of the creation of the resource group
	CreatedAt string `pulumi:"createdAt"`
	Id        string `pulumi:"id"`
	// Name of the resource group
	Name string `pulumi:"name"`
	// Name of the account owning the resource group
	Owner string `pulumi:"owner"`
	// Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
	ReadOnly bool `pulumi:"readOnly"`
	// Set of the URNs of the resources contained in the resource group
	Resources []string `pulumi:"resources"`
	// Date of the last modification of the resource group
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupResourceGroupOutput(ctx *pulumi.Context, args LookupResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResourceGroupResultOutput, error) {
			args := v.(LookupResourceGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Iam/getResourceGroup:getResourceGroup", args, LookupResourceGroupResultOutput{}, options).(LookupResourceGroupResultOutput), nil
		}).(LookupResourceGroupResultOutput)
}

// A collection of arguments for invoking getResourceGroup.
type LookupResourceGroupOutputArgs struct {
	// Id of the resource group
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupArgs)(nil)).Elem()
}

// A collection of values returned by getResourceGroup.
type LookupResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupResult)(nil)).Elem()
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutput() LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutputWithContext(ctx context.Context) LookupResourceGroupResultOutput {
	return o
}

// URN of the resource group, used when writing policies
func (o LookupResourceGroupResultOutput) GroupURN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.GroupURN }).(pulumi.StringOutput)
}

// Date of the creation of the resource group
func (o LookupResourceGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource group
func (o LookupResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name of the account owning the resource group
func (o LookupResourceGroupResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
func (o LookupResourceGroupResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

// Set of the URNs of the resources contained in the resource group
func (o LookupResourceGroupResultOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// Date of the last modification of the resource group
func (o LookupResourceGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGroupResultOutput{})
}
