// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to manage a rule on an IP firewall.
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
//			_, err := ip.NewFirewallRule(ctx, "my_firewall_rule", &ip.FirewallRuleArgs{
//				Ip:           pulumi.String("XXXXXX"),
//				IpOnFirewall: pulumi.String("XXXXXX"),
//				Sequence:     pulumi.Float64(0),
//				Action:       pulumi.String("deny"),
//				Protocol:     pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// The resource can be imported using the properties `ip`, `ip_on_firewall` and `sequence`, separated by "|" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Ip/firewallRule:FirewallRule my_firewall_rule '127.0.0.1|127.0.0.2|0'
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// Possible values for action (deny|permit)
	Action pulumi.StringOutput `pulumi:"action"`
	// Creation date of the rule
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Destination IP for your rule
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64Output `pulumi:"destinationPort"`
	// String description of field `destinationPort`
	DestinationPortDesc pulumi.StringOutput `pulumi:"destinationPortDesc"`
	// Fragments option
	Fragments pulumi.BoolOutput `pulumi:"fragments"`
	// The IP or the CIDR
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall pulumi.StringOutput `pulumi:"ipOnFirewall"`
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Description of the rule
	Rule pulumi.StringOutput `pulumi:"rule"`
	// Rule position in the rules array
	Sequence pulumi.Float64Output `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringOutput `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64Output `pulumi:"sourcePort"`
	// String description of field `sourcePort`
	SourcePortDesc pulumi.StringOutput `pulumi:"sourcePortDesc"`
	// Current state of your rule
	State pulumi.StringOutput `pulumi:"state"`
	// TCP option on your rule (syn|established)
	TcpOption pulumi.StringOutput `pulumi:"tcpOption"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpOnFirewall == nil {
		return nil, errors.New("invalid value for required argument 'IpOnFirewall'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Sequence == nil {
		return nil, errors.New("invalid value for required argument 'Sequence'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallRule
	err := ctx.RegisterResource("ovh:Ip/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("ovh:Ip/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// Possible values for action (deny|permit)
	Action *string `pulumi:"action"`
	// Creation date of the rule
	CreationDate *string `pulumi:"creationDate"`
	// Destination IP for your rule
	Destination *string `pulumi:"destination"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `pulumi:"destinationPort"`
	// String description of field `destinationPort`
	DestinationPortDesc *string `pulumi:"destinationPortDesc"`
	// Fragments option
	Fragments *bool `pulumi:"fragments"`
	// The IP or the CIDR
	Ip *string `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall *string `pulumi:"ipOnFirewall"`
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol *string `pulumi:"protocol"`
	// Description of the rule
	Rule *string `pulumi:"rule"`
	// Rule position in the rules array
	Sequence *float64 `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `pulumi:"sourcePort"`
	// String description of field `sourcePort`
	SourcePortDesc *string `pulumi:"sourcePortDesc"`
	// Current state of your rule
	State *string `pulumi:"state"`
	// TCP option on your rule (syn|established)
	TcpOption *string `pulumi:"tcpOption"`
}

type FirewallRuleState struct {
	// Possible values for action (deny|permit)
	Action pulumi.StringPtrInput
	// Creation date of the rule
	CreationDate pulumi.StringPtrInput
	// Destination IP for your rule
	Destination pulumi.StringPtrInput
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64PtrInput
	// String description of field `destinationPort`
	DestinationPortDesc pulumi.StringPtrInput
	// Fragments option
	Fragments pulumi.BoolPtrInput
	// The IP or the CIDR
	Ip pulumi.StringPtrInput
	// IPv4 address
	IpOnFirewall pulumi.StringPtrInput
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol pulumi.StringPtrInput
	// Description of the rule
	Rule pulumi.StringPtrInput
	// Rule position in the rules array
	Sequence pulumi.Float64PtrInput
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringPtrInput
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64PtrInput
	// String description of field `sourcePort`
	SourcePortDesc pulumi.StringPtrInput
	// Current state of your rule
	State pulumi.StringPtrInput
	// TCP option on your rule (syn|established)
	TcpOption pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// Possible values for action (deny|permit)
	Action string `pulumi:"action"`
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort *float64 `pulumi:"destinationPort"`
	// Fragments option
	Fragments *bool `pulumi:"fragments"`
	// The IP or the CIDR
	Ip string `pulumi:"ip"`
	// IPv4 address
	IpOnFirewall string `pulumi:"ipOnFirewall"`
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol string `pulumi:"protocol"`
	// Rule position in the rules array
	Sequence float64 `pulumi:"sequence"`
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source *string `pulumi:"source"`
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort *float64 `pulumi:"sourcePort"`
	// TCP option on your rule (syn|established)
	TcpOption *string `pulumi:"tcpOption"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Possible values for action (deny|permit)
	Action pulumi.StringInput
	// Destination port for your rule. Only with TCP/UDP protocol
	DestinationPort pulumi.Float64PtrInput
	// Fragments option
	Fragments pulumi.BoolPtrInput
	// The IP or the CIDR
	Ip pulumi.StringInput
	// IPv4 address
	IpOnFirewall pulumi.StringInput
	// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
	Protocol pulumi.StringInput
	// Rule position in the rules array
	Sequence pulumi.Float64Input
	// IPv4 CIDR notation (e.g., 192.0.2.0/24)
	Source pulumi.StringPtrInput
	// Source port for your rule. Only with TCP/UDP protocol
	SourcePort pulumi.Float64PtrInput
	// TCP option on your rule (syn|established)
	TcpOption pulumi.StringPtrInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//	FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//	FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

// Possible values for action (deny|permit)
func (o FirewallRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Creation date of the rule
func (o FirewallRuleOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Destination IP for your rule
func (o FirewallRuleOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Destination port for your rule. Only with TCP/UDP protocol
func (o FirewallRuleOutput) DestinationPort() pulumi.Float64Output {
	return o.ApplyT(func(v *FirewallRule) pulumi.Float64Output { return v.DestinationPort }).(pulumi.Float64Output)
}

// String description of field `destinationPort`
func (o FirewallRuleOutput) DestinationPortDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.DestinationPortDesc }).(pulumi.StringOutput)
}

// Fragments option
func (o FirewallRuleOutput) Fragments() pulumi.BoolOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolOutput { return v.Fragments }).(pulumi.BoolOutput)
}

// The IP or the CIDR
func (o FirewallRuleOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv4 address
func (o FirewallRuleOutput) IpOnFirewall() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.IpOnFirewall }).(pulumi.StringOutput)
}

// Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
func (o FirewallRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Description of the rule
func (o FirewallRuleOutput) Rule() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Rule }).(pulumi.StringOutput)
}

// Rule position in the rules array
func (o FirewallRuleOutput) Sequence() pulumi.Float64Output {
	return o.ApplyT(func(v *FirewallRule) pulumi.Float64Output { return v.Sequence }).(pulumi.Float64Output)
}

// IPv4 CIDR notation (e.g., 192.0.2.0/24)
func (o FirewallRuleOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Source port for your rule. Only with TCP/UDP protocol
func (o FirewallRuleOutput) SourcePort() pulumi.Float64Output {
	return o.ApplyT(func(v *FirewallRule) pulumi.Float64Output { return v.SourcePort }).(pulumi.Float64Output)
}

// String description of field `sourcePort`
func (o FirewallRuleOutput) SourcePortDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.SourcePortDesc }).(pulumi.StringOutput)
}

// Current state of your rule
func (o FirewallRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// TCP option on your rule (syn|established)
func (o FirewallRuleOutput) TcpOption() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.TcpOption }).(pulumi.StringOutput)
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
