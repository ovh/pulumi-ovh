// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a domain zone DNSSEC status.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/domain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := domain.GetZoneDNSSec(ctx, &domain.GetZoneDNSSecArgs{
//				ZoneName: "mysite.ovh",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupZoneDNSSec(ctx *pulumi.Context, args *LookupZoneDNSSecArgs, opts ...pulumi.InvokeOption) (*LookupZoneDNSSecResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupZoneDNSSecResult
	err := ctx.Invoke("ovh:Domain/getZoneDNSSec:getZoneDNSSec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZoneDNSSec.
type LookupZoneDNSSecArgs struct {
	// The name of the domain zone
	ZoneName string `pulumi:"zoneName"`
}

// A collection of values returned by getZoneDNSSec.
type LookupZoneDNSSecResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
	Status   string `pulumi:"status"`
	ZoneName string `pulumi:"zoneName"`
}

func LookupZoneDNSSecOutput(ctx *pulumi.Context, args LookupZoneDNSSecOutputArgs, opts ...pulumi.InvokeOption) LookupZoneDNSSecResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupZoneDNSSecResultOutput, error) {
			args := v.(LookupZoneDNSSecArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Domain/getZoneDNSSec:getZoneDNSSec", args, LookupZoneDNSSecResultOutput{}, options).(LookupZoneDNSSecResultOutput), nil
		}).(LookupZoneDNSSecResultOutput)
}

// A collection of arguments for invoking getZoneDNSSec.
type LookupZoneDNSSecOutputArgs struct {
	// The name of the domain zone
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}

func (LookupZoneDNSSecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneDNSSecArgs)(nil)).Elem()
}

// A collection of values returned by getZoneDNSSec.
type LookupZoneDNSSecResultOutput struct{ *pulumi.OutputState }

func (LookupZoneDNSSecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneDNSSecResult)(nil)).Elem()
}

func (o LookupZoneDNSSecResultOutput) ToLookupZoneDNSSecResultOutput() LookupZoneDNSSecResultOutput {
	return o
}

func (o LookupZoneDNSSecResultOutput) ToLookupZoneDNSSecResultOutputWithContext(ctx context.Context) LookupZoneDNSSecResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupZoneDNSSecResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneDNSSecResult) string { return v.Id }).(pulumi.StringOutput)
}

// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
func (o LookupZoneDNSSecResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneDNSSecResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupZoneDNSSecResultOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneDNSSecResult) string { return v.ZoneName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneDNSSecResultOutput{})
}
