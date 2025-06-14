// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an IP Load Balancing product
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/iploadbalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iploadbalancing.GetIpLoadBalancing(ctx, &iploadbalancing.GetIpLoadBalancingArgs{
//				ServiceName: pulumi.StringRef("XXXXXX"),
//				State:       pulumi.StringRef("ok"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIpLoadBalancing(ctx *pulumi.Context, args *GetIpLoadBalancingArgs, opts ...pulumi.InvokeOption) (*GetIpLoadBalancingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpLoadBalancingResult
	err := ctx.Invoke("ovh:IpLoadBalancing/getIpLoadBalancing:getIpLoadBalancing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpLoadBalancing.
type GetIpLoadBalancingArgs struct {
	// the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing *string `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 *string `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing
	Ipv6 *string `pulumi:"ipv6"`
	// The offer of your IP load balancing
	Offer *string `pulumi:"offer"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Can take any of the following value: "intermediate", "modern"
	SslConfiguration *string `pulumi:"sslConfiguration"`
	// Current state of your IP. Can take any of the following value: "blacklisted", "deleted", "free", "ok", "quarantined", "suspended"
	State *string `pulumi:"state"`
	// Vrack eligibility. Takes a boolean value.
	VrackEligibility *bool `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName *string `pulumi:"vrackName"`
	// Location where your service is. This takes an array of values.
	Zones []string `pulumi:"zones"`
}

// A collection of values returned by getIpLoadBalancing.
type GetIpLoadBalancingResult struct {
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IpLoadbalancing string `pulumi:"ipLoadbalancing"`
	Ipv4            string `pulumi:"ipv4"`
	Ipv6            string `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing This attribute is sensitive.
	MetricsToken string `pulumi:"metricsToken"`
	Offer        string `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones   []GetIpLoadBalancingOrderableZone `pulumi:"orderableZones"`
	ServiceName      string                            `pulumi:"serviceName"`
	SslConfiguration string                            `pulumi:"sslConfiguration"`
	State            string                            `pulumi:"state"`
	// The URN of the load balancer, to be used in IAM policies
	Urn              string   `pulumi:"urn"`
	VrackEligibility bool     `pulumi:"vrackEligibility"`
	VrackName        string   `pulumi:"vrackName"`
	Zones            []string `pulumi:"zones"`
}

func GetIpLoadBalancingOutput(ctx *pulumi.Context, args GetIpLoadBalancingOutputArgs, opts ...pulumi.InvokeOption) GetIpLoadBalancingResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIpLoadBalancingResultOutput, error) {
			args := v.(GetIpLoadBalancingArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:IpLoadBalancing/getIpLoadBalancing:getIpLoadBalancing", args, GetIpLoadBalancingResultOutput{}, options).(GetIpLoadBalancingResultOutput), nil
		}).(GetIpLoadBalancingResultOutput)
}

// A collection of arguments for invoking getIpLoadBalancing.
type GetIpLoadBalancingOutputArgs struct {
	// the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing pulumi.StringPtrInput `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringPtrInput `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing
	Ipv6 pulumi.StringPtrInput `pulumi:"ipv6"`
	// The offer of your IP load balancing
	Offer pulumi.StringPtrInput `pulumi:"offer"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Can take any of the following value: "intermediate", "modern"
	SslConfiguration pulumi.StringPtrInput `pulumi:"sslConfiguration"`
	// Current state of your IP. Can take any of the following value: "blacklisted", "deleted", "free", "ok", "quarantined", "suspended"
	State pulumi.StringPtrInput `pulumi:"state"`
	// Vrack eligibility. Takes a boolean value.
	VrackEligibility pulumi.BoolPtrInput `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringPtrInput `pulumi:"vrackName"`
	// Location where your service is. This takes an array of values.
	Zones pulumi.StringArrayInput `pulumi:"zones"`
}

func (GetIpLoadBalancingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadBalancingArgs)(nil)).Elem()
}

// A collection of values returned by getIpLoadBalancing.
type GetIpLoadBalancingResultOutput struct{ *pulumi.OutputState }

func (GetIpLoadBalancingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadBalancingResult)(nil)).Elem()
}

func (o GetIpLoadBalancingResultOutput) ToGetIpLoadBalancingResultOutput() GetIpLoadBalancingResultOutput {
	return o
}

func (o GetIpLoadBalancingResultOutput) ToGetIpLoadBalancingResultOutputWithContext(ctx context.Context) GetIpLoadBalancingResultOutput {
	return o
}

func (o GetIpLoadBalancingResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpLoadBalancingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) IpLoadbalancing() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.IpLoadbalancing }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.Ipv4 }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.Ipv6 }).(pulumi.StringOutput)
}

// The metrics token associated with your IP load balancing This attribute is sensitive.
func (o GetIpLoadBalancingResultOutput) MetricsToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.MetricsToken }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.Offer }).(pulumi.StringOutput)
}

// Available additional zone for your Load Balancer
func (o GetIpLoadBalancingResultOutput) OrderableZones() GetIpLoadBalancingOrderableZoneArrayOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) []GetIpLoadBalancingOrderableZone { return v.OrderableZones }).(GetIpLoadBalancingOrderableZoneArrayOutput)
}

func (o GetIpLoadBalancingResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) SslConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.SslConfiguration }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.State }).(pulumi.StringOutput)
}

// The URN of the load balancer, to be used in IAM policies
func (o GetIpLoadBalancingResultOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.Urn }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) VrackEligibility() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) bool { return v.VrackEligibility }).(pulumi.BoolOutput)
}

func (o GetIpLoadBalancingResultOutput) VrackName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) string { return v.VrackName }).(pulumi.StringOutput)
}

func (o GetIpLoadBalancingResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpLoadBalancingResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpLoadBalancingResultOutput{})
}
