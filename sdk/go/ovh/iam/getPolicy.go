// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve am IAM policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.GetPolicy(ctx, &iam.GetPolicyArgs{
//				Id: "my_policy_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("ovh:Iam/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// Set of actions allowed by the policy.
	Allows []string `pulumi:"allows"`
	// Set of actions that will be denied no matter what policy exists.
	Denies []string `pulumi:"denies"`
	// Group description.
	Description *string `pulumi:"description"`
	// Set of actions that will be subtracted from the `allow` list.
	Excepts []string `pulumi:"excepts"`
	// UUID of the policy.
	Id string `pulumi:"id"`
	// Set of permissions groups that apply to the policy.
	PermissionsGroups []string `pulumi:"permissionsGroups"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// Set of actions allowed by the policy.
	Allows []string `pulumi:"allows"`
	// Creation date of this group.
	CreatedAt string `pulumi:"createdAt"`
	// Set of actions that will be denied no matter what policy exists.
	Denies []string `pulumi:"denies"`
	// Group description.
	Description *string `pulumi:"description"`
	// Set of actions that will be subtracted from the `allow` list.
	Excepts []string `pulumi:"excepts"`
	Id      string   `pulumi:"id"`
	// Set of identities affected by the policy.
	Identities []string `pulumi:"identities"`
	// Name of the policy.
	Name string `pulumi:"name"`
	// Owner of the policy.
	Owner string `pulumi:"owner"`
	// Set of permissions groups that apply to the policy.
	PermissionsGroups []string `pulumi:"permissionsGroups"`
	// Indicates that the policy is a default one.
	ReadOnly bool `pulumi:"readOnly"`
	// Set of resources affected by the policy.
	Resources []string `pulumi:"resources"`
	// Date of the last update of this group.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPolicyResultOutput, error) {
			args := v.(LookupPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Iam/getPolicy:getPolicy", args, LookupPolicyResultOutput{}, options).(LookupPolicyResultOutput), nil
		}).(LookupPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyOutputArgs struct {
	// Set of actions allowed by the policy.
	Allows pulumi.StringArrayInput `pulumi:"allows"`
	// Set of actions that will be denied no matter what policy exists.
	Denies pulumi.StringArrayInput `pulumi:"denies"`
	// Group description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Set of actions that will be subtracted from the `allow` list.
	Excepts pulumi.StringArrayInput `pulumi:"excepts"`
	// UUID of the policy.
	Id pulumi.StringInput `pulumi:"id"`
	// Set of permissions groups that apply to the policy.
	PermissionsGroups pulumi.StringArrayInput `pulumi:"permissionsGroups"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

// Set of actions allowed by the policy.
func (o LookupPolicyResultOutput) Allows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Allows }).(pulumi.StringArrayOutput)
}

// Creation date of this group.
func (o LookupPolicyResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Set of actions that will be denied no matter what policy exists.
func (o LookupPolicyResultOutput) Denies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Denies }).(pulumi.StringArrayOutput)
}

// Group description.
func (o LookupPolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Set of actions that will be subtracted from the `allow` list.
func (o LookupPolicyResultOutput) Excepts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Excepts }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of identities affected by the policy.
func (o LookupPolicyResultOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Identities }).(pulumi.StringArrayOutput)
}

// Name of the policy.
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Owner of the policy.
func (o LookupPolicyResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Set of permissions groups that apply to the policy.
func (o LookupPolicyResultOutput) PermissionsGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.PermissionsGroups }).(pulumi.StringArrayOutput)
}

// Indicates that the policy is a default one.
func (o LookupPolicyResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicyResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

// Set of resources affected by the policy.
func (o LookupPolicyResultOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// Date of the last update of this group.
func (o LookupPolicyResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
