// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a credit card payment mean associated with an OVHcloud account.
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
//			_, err := me.GetPaymentmeanCreditCard(ctx, &me.GetPaymentmeanCreditCardArgs{
//				UseDefault: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPaymentmeanCreditCard(ctx *pulumi.Context, args *GetPaymentmeanCreditCardArgs, opts ...pulumi.InvokeOption) (*GetPaymentmeanCreditCardResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPaymentmeanCreditCardResult
	err := ctx.Invoke("ovh:Me/getPaymentmeanCreditCard:getPaymentmeanCreditCard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPaymentmeanCreditCard.
type GetPaymentmeanCreditCardArgs struct {
	// a regexp used to filter credit cards on their `description` attributes.
	DescriptionRegexp *string `pulumi:"descriptionRegexp"`
	// Filter credit cards on their `state` attribute. Can be "expired", "valid", "tooManyFailures"
	States []string `pulumi:"states"`
	// Retrieve credit card marked as default payment mean.
	UseDefault *bool `pulumi:"useDefault"`
	// Retrieve the credit card that will be the last to expire according to its expiration date.
	UseLastToExpire *bool `pulumi:"useLastToExpire"`
}

// A collection of values returned by getPaymentmeanCreditCard.
type GetPaymentmeanCreditCardResult struct {
	// a boolean which tells if the retrieved credit card is marked as the default payment mean
	Default bool `pulumi:"default"`
	// the description attribute of the credit card
	Description       string  `pulumi:"description"`
	DescriptionRegexp *string `pulumi:"descriptionRegexp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the state attribute of the credit card
	State           string   `pulumi:"state"`
	States          []string `pulumi:"states"`
	UseDefault      *bool    `pulumi:"useDefault"`
	UseLastToExpire *bool    `pulumi:"useLastToExpire"`
}

func GetPaymentmeanCreditCardOutput(ctx *pulumi.Context, args GetPaymentmeanCreditCardOutputArgs, opts ...pulumi.InvokeOption) GetPaymentmeanCreditCardResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetPaymentmeanCreditCardResultOutput, error) {
			args := v.(GetPaymentmeanCreditCardArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Me/getPaymentmeanCreditCard:getPaymentmeanCreditCard", args, GetPaymentmeanCreditCardResultOutput{}, options).(GetPaymentmeanCreditCardResultOutput), nil
		}).(GetPaymentmeanCreditCardResultOutput)
}

// A collection of arguments for invoking getPaymentmeanCreditCard.
type GetPaymentmeanCreditCardOutputArgs struct {
	// a regexp used to filter credit cards on their `description` attributes.
	DescriptionRegexp pulumi.StringPtrInput `pulumi:"descriptionRegexp"`
	// Filter credit cards on their `state` attribute. Can be "expired", "valid", "tooManyFailures"
	States pulumi.StringArrayInput `pulumi:"states"`
	// Retrieve credit card marked as default payment mean.
	UseDefault pulumi.BoolPtrInput `pulumi:"useDefault"`
	// Retrieve the credit card that will be the last to expire according to its expiration date.
	UseLastToExpire pulumi.BoolPtrInput `pulumi:"useLastToExpire"`
}

func (GetPaymentmeanCreditCardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPaymentmeanCreditCardArgs)(nil)).Elem()
}

// A collection of values returned by getPaymentmeanCreditCard.
type GetPaymentmeanCreditCardResultOutput struct{ *pulumi.OutputState }

func (GetPaymentmeanCreditCardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPaymentmeanCreditCardResult)(nil)).Elem()
}

func (o GetPaymentmeanCreditCardResultOutput) ToGetPaymentmeanCreditCardResultOutput() GetPaymentmeanCreditCardResultOutput {
	return o
}

func (o GetPaymentmeanCreditCardResultOutput) ToGetPaymentmeanCreditCardResultOutputWithContext(ctx context.Context) GetPaymentmeanCreditCardResultOutput {
	return o
}

// a boolean which tells if the retrieved credit card is marked as the default payment mean
func (o GetPaymentmeanCreditCardResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// the description attribute of the credit card
func (o GetPaymentmeanCreditCardResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetPaymentmeanCreditCardResultOutput) DescriptionRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) *string { return v.DescriptionRegexp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPaymentmeanCreditCardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) string { return v.Id }).(pulumi.StringOutput)
}

// the state attribute of the credit card
func (o GetPaymentmeanCreditCardResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetPaymentmeanCreditCardResultOutput) States() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) []string { return v.States }).(pulumi.StringArrayOutput)
}

func (o GetPaymentmeanCreditCardResultOutput) UseDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) *bool { return v.UseDefault }).(pulumi.BoolPtrOutput)
}

func (o GetPaymentmeanCreditCardResultOutput) UseLastToExpire() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPaymentmeanCreditCardResult) *bool { return v.UseLastToExpire }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPaymentmeanCreditCardResultOutput{})
}
