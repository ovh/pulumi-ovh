// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a domain zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/domain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := domain.GetZone(ctx, &domain.GetZoneArgs{
//				Name: "mysite.ovh",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupZoneResult
	err := ctx.Invoke("ovh:Domain/getZone:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZone.
type LookupZoneArgs struct {
	// The name of the domain zone.
	Name string `pulumi:"name"`
}

// A collection of values returned by getZone.
type LookupZoneResult struct {
	// URN of the DNS zone
	ZoneURN string `pulumi:"ZoneURN"`
	// Is DNSSEC supported by this zone
	DnssecSupported bool `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast bool `pulumi:"hasDnsAnycast"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Last update date of the DNS zone
	LastUpdate string `pulumi:"lastUpdate"`
	Name       string `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers []string `pulumi:"nameServers"`
}

func LookupZoneOutput(ctx *pulumi.Context, args LookupZoneOutputArgs, opts ...pulumi.InvokeOption) LookupZoneResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupZoneResultOutput, error) {
			args := v.(LookupZoneArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Domain/getZone:getZone", args, LookupZoneResultOutput{}, options).(LookupZoneResultOutput), nil
		}).(LookupZoneResultOutput)
}

// A collection of arguments for invoking getZone.
type LookupZoneOutputArgs struct {
	// The name of the domain zone.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneArgs)(nil)).Elem()
}

// A collection of values returned by getZone.
type LookupZoneResultOutput struct{ *pulumi.OutputState }

func (LookupZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneResult)(nil)).Elem()
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutput() LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutputWithContext(ctx context.Context) LookupZoneResultOutput {
	return o
}

// URN of the DNS zone
func (o LookupZoneResultOutput) ZoneURN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.ZoneURN }).(pulumi.StringOutput)
}

// Is DNSSEC supported by this zone
func (o LookupZoneResultOutput) DnssecSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZoneResult) bool { return v.DnssecSupported }).(pulumi.BoolOutput)
}

// hasDnsAnycast flag of the DNS zone
func (o LookupZoneResultOutput) HasDnsAnycast() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupZoneResult) bool { return v.HasDnsAnycast }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last update date of the DNS zone
func (o LookupZoneResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

func (o LookupZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name servers that host the DNS zone
func (o LookupZoneResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupZoneResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneResultOutput{})
}
