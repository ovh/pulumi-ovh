// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information the list of existing OAuth2 service account IDs.
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
//			_, err := me.GetAPIOAuth2Client(ctx, &me.GetAPIOAuth2ClientArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAPIOAuth2Clients(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAPIOAuth2ClientsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAPIOAuth2ClientsResult
	err := ctx.Invoke("ovh:Me/getAPIOAuth2Clients:getAPIOAuth2Clients", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAPIOAuth2Clients.
type GetAPIOAuth2ClientsResult struct {
	// The list of all the existing client IDs.
	ClientIds []string `pulumi:"clientIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetAPIOAuth2ClientsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetAPIOAuth2ClientsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetAPIOAuth2ClientsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ovh:Me/getAPIOAuth2Clients:getAPIOAuth2Clients", nil, GetAPIOAuth2ClientsResultOutput{}, options).(GetAPIOAuth2ClientsResultOutput), nil
	}).(GetAPIOAuth2ClientsResultOutput)
}

// A collection of values returned by getAPIOAuth2Clients.
type GetAPIOAuth2ClientsResultOutput struct{ *pulumi.OutputState }

func (GetAPIOAuth2ClientsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAPIOAuth2ClientsResult)(nil)).Elem()
}

func (o GetAPIOAuth2ClientsResultOutput) ToGetAPIOAuth2ClientsResultOutput() GetAPIOAuth2ClientsResultOutput {
	return o
}

func (o GetAPIOAuth2ClientsResultOutput) ToGetAPIOAuth2ClientsResultOutputWithContext(ctx context.Context) GetAPIOAuth2ClientsResultOutput {
	return o
}

// The list of all the existing client IDs.
func (o GetAPIOAuth2ClientsResultOutput) ClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAPIOAuth2ClientsResult) []string { return v.ClientIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAPIOAuth2ClientsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAPIOAuth2ClientsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAPIOAuth2ClientsResultOutput{})
}
