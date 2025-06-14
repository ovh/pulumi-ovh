// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an identity group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := me.GetIdentityGroup(ctx, &me.GetIdentityGroupArgs{
//				Name: "my_group_name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdentityGroup(ctx *pulumi.Context, args *LookupIdentityGroupArgs, opts ...pulumi.InvokeOption) (*LookupIdentityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdentityGroupResult
	err := ctx.Invoke("ovh:Me/getIdentityGroup:getIdentityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityGroup.
type LookupIdentityGroupArgs struct {
	// Group name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIdentityGroup.
type LookupIdentityGroupResult struct {
	// Identity URN of the group.
	GroupURN string `pulumi:"GroupURN"`
	// Creation date of this group.
	Creation string `pulumi:"creation"`
	// Is the group a default and immutable one.
	DefaultGroup bool `pulumi:"defaultGroup"`
	// Group description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Date of the last update of this group.
	LastUpdate string `pulumi:"lastUpdate"`
	Name       string `pulumi:"name"`
	// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
	Role string `pulumi:"role"`
}

func LookupIdentityGroupOutput(ctx *pulumi.Context, args LookupIdentityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIdentityGroupResultOutput, error) {
			args := v.(LookupIdentityGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Me/getIdentityGroup:getIdentityGroup", args, LookupIdentityGroupResultOutput{}, options).(LookupIdentityGroupResultOutput), nil
		}).(LookupIdentityGroupResultOutput)
}

// A collection of arguments for invoking getIdentityGroup.
type LookupIdentityGroupOutputArgs struct {
	// Group name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupIdentityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityGroup.
type LookupIdentityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityGroupResult)(nil)).Elem()
}

func (o LookupIdentityGroupResultOutput) ToLookupIdentityGroupResultOutput() LookupIdentityGroupResultOutput {
	return o
}

func (o LookupIdentityGroupResultOutput) ToLookupIdentityGroupResultOutputWithContext(ctx context.Context) LookupIdentityGroupResultOutput {
	return o
}

// Identity URN of the group.
func (o LookupIdentityGroupResultOutput) GroupURN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.GroupURN }).(pulumi.StringOutput)
}

// Creation date of this group.
func (o LookupIdentityGroupResultOutput) Creation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Creation }).(pulumi.StringOutput)
}

// Is the group a default and immutable one.
func (o LookupIdentityGroupResultOutput) DefaultGroup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) bool { return v.DefaultGroup }).(pulumi.BoolOutput)
}

// Group description.
func (o LookupIdentityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIdentityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Date of the last update of this group.
func (o LookupIdentityGroupResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

func (o LookupIdentityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
func (o LookupIdentityGroupResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityGroupResult) string { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityGroupResultOutput{})
}
