// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a rule on an IP firewall.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/ip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ip.GetFirewallRule(ctx, &ip.GetFirewallRuleArgs{
//				Ip:           "XXXXXX",
//				IpOnFirewall: "XXXXXX",
//				Sequence:     0,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("ovh:Ip/getFirewallRule:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallRule.
type LookupFirewallRuleArgs struct {
	// The IP or the CIDR
	Ip string `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall string `pulumi:"ipOnFirewall"`
	// Rule position in the rules array
	Sequence float64 `pulumi:"sequence"`
}

// A collection of values returned by getFirewallRule.
type LookupFirewallRuleResult struct {
	// Possible values for action (deny|permit)
	Action string `pulumi:"action"`
	// Creation date of the rule
	CreationDate string `pulumi:"creationDate"`
	// Destination IP for your rule
	Destination string `pulumi:"destination"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort string `pulumi:"destinationPort"`
	// Fragments option
	Fragments bool `pulumi:"fragments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IP or the CIDR
	Ip string `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall string `pulumi:"ipOnFirewall"`
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol string `pulumi:"protocol"`
	// Description of the rule
	Rule string `pulumi:"rule"`
	// Rule position in the rules array
	Sequence float64 `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source string `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort string `pulumi:"sourcePort"`
	// Current state of your rule
	State string `pulumi:"state"`
	// TCP option on your rule (syn|established)
	TcpOption string `pulumi:"tcpOption"`
}

func LookupFirewallRuleOutput(ctx *pulumi.Context, args LookupFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFirewallRuleResultOutput, error) {
			args := v.(LookupFirewallRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Ip/getFirewallRule:getFirewallRule", args, LookupFirewallRuleResultOutput{}, options).(LookupFirewallRuleResultOutput), nil
		}).(LookupFirewallRuleResultOutput)
}

// A collection of arguments for invoking getFirewallRule.
type LookupFirewallRuleOutputArgs struct {
	// The IP or the CIDR
	Ip pulumi.StringInput `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall pulumi.StringInput `pulumi:"ipOnFirewall"`
	// Rule position in the rules array
	Sequence pulumi.Float64Input `pulumi:"sequence"`
}

func (LookupFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallRule.
type LookupFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleResult)(nil)).Elem()
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutput() LookupFirewallRuleResultOutput {
	return o
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutputWithContext(ctx context.Context) LookupFirewallRuleResultOutput {
	return o
}

// Possible values for action (deny|permit)
func (o LookupFirewallRuleResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Action }).(pulumi.StringOutput)
}

// Creation date of the rule
func (o LookupFirewallRuleResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Destination IP for your rule
func (o LookupFirewallRuleResultOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Destination }).(pulumi.StringOutput)
}

// Destination port for your rule. Only with TCP/UDP protocol
func (o LookupFirewallRuleResultOutput) DestinationPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.DestinationPort }).(pulumi.StringOutput)
}

// Fragments option
func (o LookupFirewallRuleResultOutput) Fragments() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) bool { return v.Fragments }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IP or the CIDR
func (o LookupFirewallRuleResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Ip }).(pulumi.StringOutput)
}

// IPv4 address
func (o LookupFirewallRuleResultOutput) IpOnFirewall() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.IpOnFirewall }).(pulumi.StringOutput)
}

// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
func (o LookupFirewallRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Description of the rule
func (o LookupFirewallRuleResultOutput) Rule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Rule }).(pulumi.StringOutput)
}

// Rule position in the rules array
func (o LookupFirewallRuleResultOutput) Sequence() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFirewallRuleResult) float64 { return v.Sequence }).(pulumi.Float64Output)
}

// IPv4 CIDR notation (e.g., 192.0.2.0/24)
func (o LookupFirewallRuleResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Source }).(pulumi.StringOutput)
}

// Source port for your rule. Only with TCP/UDP protocol
func (o LookupFirewallRuleResultOutput) SourcePort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.SourcePort }).(pulumi.StringOutput)
}

// Current state of your rule
func (o LookupFirewallRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.State }).(pulumi.StringOutput)
}

// TCP option on your rule (syn|established)
func (o LookupFirewallRuleResultOutput) TcpOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.TcpOption }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}
