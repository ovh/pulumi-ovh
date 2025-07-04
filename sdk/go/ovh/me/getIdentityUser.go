// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an identity user.
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
//			_, err := me.GetIdentityUser(ctx, &me.GetIdentityUserArgs{
//				User: "my_user_login",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdentityUser(ctx *pulumi.Context, args *LookupIdentityUserArgs, opts ...pulumi.InvokeOption) (*LookupIdentityUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdentityUserResult
	err := ctx.Invoke("ovh:Me/getIdentityUser:getIdentityUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdentityUser.
type LookupIdentityUserArgs struct {
	// User's login.
	User string `pulumi:"user"`
}

// A collection of values returned by getIdentityUser.
type LookupIdentityUserResult struct {
	// User's identity URN.
	UserURN string `pulumi:"UserURN"`
	// Creation date of this user.
	Creation string `pulumi:"creation"`
	// User description.
	Description string `pulumi:"description"`
	// User's email.
	Email string `pulumi:"email"`
	// User's group.
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Last update of this user.
	LastUpdate string `pulumi:"lastUpdate"`
	// User's login suffix.
	Login string `pulumi:"login"`
	// When the user changed his password for the last time.
	PasswordLastUpdate string `pulumi:"passwordLastUpdate"`
	// Current user's status.
	Status string `pulumi:"status"`
	User   string `pulumi:"user"`
}

func LookupIdentityUserOutput(ctx *pulumi.Context, args LookupIdentityUserOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIdentityUserResultOutput, error) {
			args := v.(LookupIdentityUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Me/getIdentityUser:getIdentityUser", args, LookupIdentityUserResultOutput{}, options).(LookupIdentityUserResultOutput), nil
		}).(LookupIdentityUserResultOutput)
}

// A collection of arguments for invoking getIdentityUser.
type LookupIdentityUserOutputArgs struct {
	// User's login.
	User pulumi.StringInput `pulumi:"user"`
}

func (LookupIdentityUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityUserArgs)(nil)).Elem()
}

// A collection of values returned by getIdentityUser.
type LookupIdentityUserResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityUserResult)(nil)).Elem()
}

func (o LookupIdentityUserResultOutput) ToLookupIdentityUserResultOutput() LookupIdentityUserResultOutput {
	return o
}

func (o LookupIdentityUserResultOutput) ToLookupIdentityUserResultOutputWithContext(ctx context.Context) LookupIdentityUserResultOutput {
	return o
}

// User's identity URN.
func (o LookupIdentityUserResultOutput) UserURN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.UserURN }).(pulumi.StringOutput)
}

// Creation date of this user.
func (o LookupIdentityUserResultOutput) Creation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Creation }).(pulumi.StringOutput)
}

// User description.
func (o LookupIdentityUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// User's email.
func (o LookupIdentityUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// User's group.
func (o LookupIdentityUserResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIdentityUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last update of this user.
func (o LookupIdentityUserResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

// User's login suffix.
func (o LookupIdentityUserResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Login }).(pulumi.StringOutput)
}

// When the user changed his password for the last time.
func (o LookupIdentityUserResultOutput) PasswordLastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.PasswordLastUpdate }).(pulumi.StringOutput)
}

// Current user's status.
func (o LookupIdentityUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIdentityUserResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdentityUserResult) string { return v.User }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityUserResultOutput{})
}
