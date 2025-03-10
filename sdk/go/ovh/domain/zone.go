// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Zone struct {
	pulumi.CustomResourceState

	ZoneURN pulumi.StringOutput `pulumi:"ZoneURN"`
	// Is DNSSEC supported by this zone
	DnssecSupported pulumi.BoolOutput `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast pulumi.BoolOutput `pulumi:"hasDnsAnycast"`
	// Last update date of the DNS zone
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// Zone name
	Name pulumi.StringOutput `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// Details about an Order
	Orders ZoneOrderArrayOutput `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan ZonePlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions ZonePlanOptionArrayOutput `pulumi:"planOptions"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		args = &ZoneArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Zone
	err := ctx.RegisterResource("ovh:Domain/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("ovh:Domain/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	ZoneURN *string `pulumi:"ZoneURN"`
	// Is DNSSEC supported by this zone
	DnssecSupported *bool `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast *bool `pulumi:"hasDnsAnycast"`
	// Last update date of the DNS zone
	LastUpdate *string `pulumi:"lastUpdate"`
	// Zone name
	Name *string `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers []string `pulumi:"nameServers"`
	// Details about an Order
	Orders []ZoneOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *ZonePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []ZonePlanOption `pulumi:"planOptions"`
}

type ZoneState struct {
	ZoneURN pulumi.StringPtrInput
	// Is DNSSEC supported by this zone
	DnssecSupported pulumi.BoolPtrInput
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast pulumi.BoolPtrInput
	// Last update date of the DNS zone
	LastUpdate pulumi.StringPtrInput
	// Zone name
	Name pulumi.StringPtrInput
	// Name servers that host the DNS zone
	NameServers pulumi.StringArrayInput
	// Details about an Order
	Orders ZoneOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan ZonePlanPtrInput
	// Product Plan to order
	PlanOptions ZonePlanOptionArrayInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// Details about an Order
	Orders []ZoneOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *ZonePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []ZonePlanOption `pulumi:"planOptions"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// Details about an Order
	Orders ZoneOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan ZonePlanPtrInput
	// Product Plan to order
	PlanOptions ZonePlanOptionArrayInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//	ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//	ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func (o ZoneOutput) ZoneURN() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.ZoneURN }).(pulumi.StringOutput)
}

// Is DNSSEC supported by this zone
func (o ZoneOutput) DnssecSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v *Zone) pulumi.BoolOutput { return v.DnssecSupported }).(pulumi.BoolOutput)
}

// hasDnsAnycast flag of the DNS zone
func (o ZoneOutput) HasDnsAnycast() pulumi.BoolOutput {
	return o.ApplyT(func(v *Zone) pulumi.BoolOutput { return v.HasDnsAnycast }).(pulumi.BoolOutput)
}

// Last update date of the DNS zone
func (o ZoneOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// Zone name
func (o ZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name servers that host the DNS zone
func (o ZoneOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

// Details about an Order
func (o ZoneOutput) Orders() ZoneOrderArrayOutput {
	return o.ApplyT(func(v *Zone) ZoneOrderArrayOutput { return v.Orders }).(ZoneOrderArrayOutput)
}

// Ovh Subsidiary
func (o ZoneOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode
//
// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
func (o ZoneOutput) PaymentMean() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.PaymentMean }).(pulumi.StringPtrOutput)
}

// Product Plan to order
func (o ZoneOutput) Plan() ZonePlanOutput {
	return o.ApplyT(func(v *Zone) ZonePlanOutput { return v.Plan }).(ZonePlanOutput)
}

// Product Plan to order
func (o ZoneOutput) PlanOptions() ZonePlanOptionArrayOutput {
	return o.ApplyT(func(v *Zone) ZonePlanOptionArrayOutput { return v.PlanOptions }).(ZonePlanOptionArrayOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
