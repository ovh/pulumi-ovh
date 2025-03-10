// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vps

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vps struct {
	pulumi.CustomResourceState

	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// Set the name displayed in Manager for your VPS (max 50 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// IAM resource metadata
	Iam VpsIamOutput `pulumi:"iam"`
	// KVM keyboard layout on VPS Cloud
	Keymap      pulumi.StringOutput  `pulumi:"keymap"`
	MemoryLimit pulumi.Float64Output `pulumi:"memoryLimit"`
	// A structure describing characteristics of a VPS model
	Model VpsModelOutput `pulumi:"model"`
	// Ip blocks for OVH monitoring servers
	MonitoringIpBlocks pulumi.StringArrayOutput `pulumi:"monitoringIpBlocks"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	// All values a VPS netboot mode can be in
	NetbootMode pulumi.StringOutput `pulumi:"netbootMode"`
	// All offers a VPS can have
	OfferType pulumi.StringOutput `pulumi:"offerType"`
	// Details about an Order
	Order VpsOrderOutput `pulumi:"order"`
	// OVH subsidiaries
	OvhSubsidiary pulumi.StringPtrOutput   `pulumi:"ovhSubsidiary"`
	PlanOptions   VpsPlanOptionArrayOutput `pulumi:"planOptions"`
	Plans         VpsPlanArrayOutput       `pulumi:"plans"`
	// The internal name of your VPS offer
	ServiceName   pulumi.StringOutput `pulumi:"serviceName"`
	SlaMonitoring pulumi.BoolOutput   `pulumi:"slaMonitoring"`
	// All states a VPS can be in
	State pulumi.StringOutput  `pulumi:"state"`
	Vcore pulumi.Float64Output `pulumi:"vcore"`
	Zone  pulumi.StringOutput  `pulumi:"zone"`
}

// NewVps registers a new resource with the given unique name, arguments, and options.
func NewVps(ctx *pulumi.Context,
	name string, args *VpsArgs, opts ...pulumi.ResourceOption) (*Vps, error) {
	if args == nil {
		args = &VpsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vps
	err := ctx.RegisterResource("ovh:Vps/vps:Vps", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVps gets an existing Vps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpsState, opts ...pulumi.ResourceOption) (*Vps, error) {
	var resource Vps
	err := ctx.ReadResource("ovh:Vps/vps:Vps", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vps resources.
type vpsState struct {
	Cluster *string `pulumi:"cluster"`
	// Set the name displayed in Manager for your VPS (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// IAM resource metadata
	Iam *VpsIam `pulumi:"iam"`
	// KVM keyboard layout on VPS Cloud
	Keymap      *string  `pulumi:"keymap"`
	MemoryLimit *float64 `pulumi:"memoryLimit"`
	// A structure describing characteristics of a VPS model
	Model *VpsModel `pulumi:"model"`
	// Ip blocks for OVH monitoring servers
	MonitoringIpBlocks []string `pulumi:"monitoringIpBlocks"`
	Name               *string  `pulumi:"name"`
	// All values a VPS netboot mode can be in
	NetbootMode *string `pulumi:"netbootMode"`
	// All offers a VPS can have
	OfferType *string `pulumi:"offerType"`
	// Details about an Order
	Order *VpsOrder `pulumi:"order"`
	// OVH subsidiaries
	OvhSubsidiary *string         `pulumi:"ovhSubsidiary"`
	PlanOptions   []VpsPlanOption `pulumi:"planOptions"`
	Plans         []VpsPlan       `pulumi:"plans"`
	// The internal name of your VPS offer
	ServiceName   *string `pulumi:"serviceName"`
	SlaMonitoring *bool   `pulumi:"slaMonitoring"`
	// All states a VPS can be in
	State *string  `pulumi:"state"`
	Vcore *float64 `pulumi:"vcore"`
	Zone  *string  `pulumi:"zone"`
}

type VpsState struct {
	Cluster pulumi.StringPtrInput
	// Set the name displayed in Manager for your VPS (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// IAM resource metadata
	Iam VpsIamPtrInput
	// KVM keyboard layout on VPS Cloud
	Keymap      pulumi.StringPtrInput
	MemoryLimit pulumi.Float64PtrInput
	// A structure describing characteristics of a VPS model
	Model VpsModelPtrInput
	// Ip blocks for OVH monitoring servers
	MonitoringIpBlocks pulumi.StringArrayInput
	Name               pulumi.StringPtrInput
	// All values a VPS netboot mode can be in
	NetbootMode pulumi.StringPtrInput
	// All offers a VPS can have
	OfferType pulumi.StringPtrInput
	// Details about an Order
	Order VpsOrderPtrInput
	// OVH subsidiaries
	OvhSubsidiary pulumi.StringPtrInput
	PlanOptions   VpsPlanOptionArrayInput
	Plans         VpsPlanArrayInput
	// The internal name of your VPS offer
	ServiceName   pulumi.StringPtrInput
	SlaMonitoring pulumi.BoolPtrInput
	// All states a VPS can be in
	State pulumi.StringPtrInput
	Vcore pulumi.Float64PtrInput
	Zone  pulumi.StringPtrInput
}

func (VpsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpsState)(nil)).Elem()
}

type vpsArgs struct {
	// Set the name displayed in Manager for your VPS (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// KVM keyboard layout on VPS Cloud
	Keymap      *string  `pulumi:"keymap"`
	MemoryLimit *float64 `pulumi:"memoryLimit"`
	// A structure describing characteristics of a VPS model
	Model *VpsModel `pulumi:"model"`
	// Ip blocks for OVH monitoring servers
	MonitoringIpBlocks []string `pulumi:"monitoringIpBlocks"`
	Name               *string  `pulumi:"name"`
	// All values a VPS netboot mode can be in
	NetbootMode *string `pulumi:"netbootMode"`
	// All offers a VPS can have
	OfferType *string `pulumi:"offerType"`
	// OVH subsidiaries
	OvhSubsidiary *string         `pulumi:"ovhSubsidiary"`
	PlanOptions   []VpsPlanOption `pulumi:"planOptions"`
	Plans         []VpsPlan       `pulumi:"plans"`
	SlaMonitoring *bool           `pulumi:"slaMonitoring"`
	// All states a VPS can be in
	State *string  `pulumi:"state"`
	Vcore *float64 `pulumi:"vcore"`
	Zone  *string  `pulumi:"zone"`
}

// The set of arguments for constructing a Vps resource.
type VpsArgs struct {
	// Set the name displayed in Manager for your VPS (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// KVM keyboard layout on VPS Cloud
	Keymap      pulumi.StringPtrInput
	MemoryLimit pulumi.Float64PtrInput
	// A structure describing characteristics of a VPS model
	Model VpsModelPtrInput
	// Ip blocks for OVH monitoring servers
	MonitoringIpBlocks pulumi.StringArrayInput
	Name               pulumi.StringPtrInput
	// All values a VPS netboot mode can be in
	NetbootMode pulumi.StringPtrInput
	// All offers a VPS can have
	OfferType pulumi.StringPtrInput
	// OVH subsidiaries
	OvhSubsidiary pulumi.StringPtrInput
	PlanOptions   VpsPlanOptionArrayInput
	Plans         VpsPlanArrayInput
	SlaMonitoring pulumi.BoolPtrInput
	// All states a VPS can be in
	State pulumi.StringPtrInput
	Vcore pulumi.Float64PtrInput
	Zone  pulumi.StringPtrInput
}

func (VpsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpsArgs)(nil)).Elem()
}

type VpsInput interface {
	pulumi.Input

	ToVpsOutput() VpsOutput
	ToVpsOutputWithContext(ctx context.Context) VpsOutput
}

func (*Vps) ElementType() reflect.Type {
	return reflect.TypeOf((**Vps)(nil)).Elem()
}

func (i *Vps) ToVpsOutput() VpsOutput {
	return i.ToVpsOutputWithContext(context.Background())
}

func (i *Vps) ToVpsOutputWithContext(ctx context.Context) VpsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpsOutput)
}

// VpsArrayInput is an input type that accepts VpsArray and VpsArrayOutput values.
// You can construct a concrete instance of `VpsArrayInput` via:
//
//	VpsArray{ VpsArgs{...} }
type VpsArrayInput interface {
	pulumi.Input

	ToVpsArrayOutput() VpsArrayOutput
	ToVpsArrayOutputWithContext(context.Context) VpsArrayOutput
}

type VpsArray []VpsInput

func (VpsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vps)(nil)).Elem()
}

func (i VpsArray) ToVpsArrayOutput() VpsArrayOutput {
	return i.ToVpsArrayOutputWithContext(context.Background())
}

func (i VpsArray) ToVpsArrayOutputWithContext(ctx context.Context) VpsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpsArrayOutput)
}

// VpsMapInput is an input type that accepts VpsMap and VpsMapOutput values.
// You can construct a concrete instance of `VpsMapInput` via:
//
//	VpsMap{ "key": VpsArgs{...} }
type VpsMapInput interface {
	pulumi.Input

	ToVpsMapOutput() VpsMapOutput
	ToVpsMapOutputWithContext(context.Context) VpsMapOutput
}

type VpsMap map[string]VpsInput

func (VpsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vps)(nil)).Elem()
}

func (i VpsMap) ToVpsMapOutput() VpsMapOutput {
	return i.ToVpsMapOutputWithContext(context.Background())
}

func (i VpsMap) ToVpsMapOutputWithContext(ctx context.Context) VpsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpsMapOutput)
}

type VpsOutput struct{ *pulumi.OutputState }

func (VpsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vps)(nil)).Elem()
}

func (o VpsOutput) ToVpsOutput() VpsOutput {
	return o
}

func (o VpsOutput) ToVpsOutputWithContext(ctx context.Context) VpsOutput {
	return o
}

func (o VpsOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// Set the name displayed in Manager for your VPS (max 50 chars)
func (o VpsOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// IAM resource metadata
func (o VpsOutput) Iam() VpsIamOutput {
	return o.ApplyT(func(v *Vps) VpsIamOutput { return v.Iam }).(VpsIamOutput)
}

// KVM keyboard layout on VPS Cloud
func (o VpsOutput) Keymap() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.Keymap }).(pulumi.StringOutput)
}

func (o VpsOutput) MemoryLimit() pulumi.Float64Output {
	return o.ApplyT(func(v *Vps) pulumi.Float64Output { return v.MemoryLimit }).(pulumi.Float64Output)
}

// A structure describing characteristics of a VPS model
func (o VpsOutput) Model() VpsModelOutput {
	return o.ApplyT(func(v *Vps) VpsModelOutput { return v.Model }).(VpsModelOutput)
}

// Ip blocks for OVH monitoring servers
func (o VpsOutput) MonitoringIpBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringArrayOutput { return v.MonitoringIpBlocks }).(pulumi.StringArrayOutput)
}

func (o VpsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// All values a VPS netboot mode can be in
func (o VpsOutput) NetbootMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.NetbootMode }).(pulumi.StringOutput)
}

// All offers a VPS can have
func (o VpsOutput) OfferType() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.OfferType }).(pulumi.StringOutput)
}

// Details about an Order
func (o VpsOutput) Order() VpsOrderOutput {
	return o.ApplyT(func(v *Vps) VpsOrderOutput { return v.Order }).(VpsOrderOutput)
}

// OVH subsidiaries
func (o VpsOutput) OvhSubsidiary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringPtrOutput { return v.OvhSubsidiary }).(pulumi.StringPtrOutput)
}

func (o VpsOutput) PlanOptions() VpsPlanOptionArrayOutput {
	return o.ApplyT(func(v *Vps) VpsPlanOptionArrayOutput { return v.PlanOptions }).(VpsPlanOptionArrayOutput)
}

func (o VpsOutput) Plans() VpsPlanArrayOutput {
	return o.ApplyT(func(v *Vps) VpsPlanArrayOutput { return v.Plans }).(VpsPlanArrayOutput)
}

// The internal name of your VPS offer
func (o VpsOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o VpsOutput) SlaMonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vps) pulumi.BoolOutput { return v.SlaMonitoring }).(pulumi.BoolOutput)
}

// All states a VPS can be in
func (o VpsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o VpsOutput) Vcore() pulumi.Float64Output {
	return o.ApplyT(func(v *Vps) pulumi.Float64Output { return v.Vcore }).(pulumi.Float64Output)
}

func (o VpsOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Vps) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpsArrayOutput struct{ *pulumi.OutputState }

func (VpsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vps)(nil)).Elem()
}

func (o VpsArrayOutput) ToVpsArrayOutput() VpsArrayOutput {
	return o
}

func (o VpsArrayOutput) ToVpsArrayOutputWithContext(ctx context.Context) VpsArrayOutput {
	return o
}

func (o VpsArrayOutput) Index(i pulumi.IntInput) VpsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vps {
		return vs[0].([]*Vps)[vs[1].(int)]
	}).(VpsOutput)
}

type VpsMapOutput struct{ *pulumi.OutputState }

func (VpsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vps)(nil)).Elem()
}

func (o VpsMapOutput) ToVpsMapOutput() VpsMapOutput {
	return o
}

func (o VpsMapOutput) ToVpsMapOutputWithContext(ctx context.Context) VpsMapOutput {
	return o
}

func (o VpsMapOutput) MapIndex(k pulumi.StringInput) VpsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vps {
		return vs[0].(map[string]*Vps)[vs[1].(string)]
	}).(VpsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpsInput)(nil)).Elem(), &Vps{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpsArrayInput)(nil)).Elem(), VpsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpsMapInput)(nil)).Elem(), VpsMap{})
	pulumi.RegisterOutputType(VpsOutput{})
	pulumi.RegisterOutputType(VpsArrayOutput{})
	pulumi.RegisterOutputType(VpsMapOutput{})
}
