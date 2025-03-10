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

type Firewall struct {
	pulumi.CustomResourceState

	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringOutput `pulumi:"ipOnFirewall"`
	// Current state of your ip on firewall
	State pulumi.StringOutput `pulumi:"state"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpOnFirewall == nil {
		return nil, errors.New("invalid value for required argument 'IpOnFirewall'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Firewall
	err := ctx.RegisterResource("ovh:Ip/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("ovh:Ip/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	Enabled *bool `pulumi:"enabled"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip *string `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall *string `pulumi:"ipOnFirewall"`
	// Current state of your ip on firewall
	State *string `pulumi:"state"`
}

type FirewallState struct {
	Enabled pulumi.BoolPtrInput
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringPtrInput
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringPtrInput
	// Current state of your ip on firewall
	State pulumi.StringPtrInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	Enabled *bool `pulumi:"enabled"`
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip string `pulumi:"ip"`
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall string `pulumi:"ipOnFirewall"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	Enabled pulumi.BoolPtrInput
	// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
	Ip pulumi.StringInput
	// IPv4 address (e.g., 192.0.2.0)
	IpOnFirewall pulumi.StringInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//	FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//	FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct{ *pulumi.OutputState }

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

func (o FirewallOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Firewall) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)
func (o FirewallOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv4 address (e.g., 192.0.2.0)
func (o FirewallOutput) IpOnFirewall() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.IpOnFirewall }).(pulumi.StringOutput)
}

// Current state of your ip on firewall
func (o FirewallOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].([]*Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].(map[string]*Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInput)(nil)).Elem(), &Firewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallArrayInput)(nil)).Elem(), FirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMapInput)(nil)).Elem(), FirewallMap{})
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}
