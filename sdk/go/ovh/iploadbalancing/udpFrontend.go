// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server group (frontend) to be used by loadbalancing frontend(s)
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
//			lb, err := iploadbalancing.GetIpLoadBalancing(ctx, &iploadbalancing.GetIpLoadBalancingArgs{
//				ServiceName: pulumi.StringRef("ip-1.2.3.4"),
//				State:       pulumi.StringRef("ok"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iploadbalancing.NewUdpFrontend(ctx, "test_frontend", &iploadbalancing.UdpFrontendArgs{
//				ServiceName: pulumi.String(lb.ServiceName),
//				DisplayName: pulumi.String("ingress-8080-gra"),
//				Zone:        pulumi.String("all"),
//				Port:        pulumi.String("10,11"),
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
// UDP frontend can be imported using the following format `service_name` and the `id` of the frontend separated by "/" e.g.
//
// bash
//
// ```sh
// $ pulumi import ovh:IpLoadBalancing/udpFrontend:UdpFrontend testfrontend service_name/frontend_id
// ```
type UdpFrontend struct {
	pulumi.CustomResourceState

	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayOutput `pulumi:"dedicatedIpfos"`
	// Default UDP Farm of your frontend
	DefaultFarmId pulumi.Float64PtrOutput `pulumi:"defaultFarmId"`
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Human readable name for your frontend
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Id of your frontend
	FrontendId pulumi.Float64Output `pulumi:"frontendId"`
	// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringOutput `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewUdpFrontend registers a new resource with the given unique name, arguments, and options.
func NewUdpFrontend(ctx *pulumi.Context,
	name string, args *UdpFrontendArgs, opts ...pulumi.ResourceOption) (*UdpFrontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UdpFrontend
	err := ctx.RegisterResource("ovh:IpLoadBalancing/udpFrontend:UdpFrontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUdpFrontend gets an existing UdpFrontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUdpFrontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UdpFrontendState, opts ...pulumi.ResourceOption) (*UdpFrontend, error) {
	var resource UdpFrontend
	err := ctx.ReadResource("ovh:IpLoadBalancing/udpFrontend:UdpFrontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UdpFrontend resources.
type udpFrontendState struct {
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default UDP Farm of your frontend
	DefaultFarmId *float64 `pulumi:"defaultFarmId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend
	DisplayName *string `pulumi:"displayName"`
	// Id of your frontend
	FrontendId *float64 `pulumi:"frontendId"`
	// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
	Port *string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone *string `pulumi:"zone"`
}

type UdpFrontendState struct {
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default UDP Farm of your frontend
	DefaultFarmId pulumi.Float64PtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend
	DisplayName pulumi.StringPtrInput
	// Id of your frontend
	FrontendId pulumi.Float64PtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringPtrInput
}

func (UdpFrontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*udpFrontendState)(nil)).Elem()
}

type udpFrontendArgs struct {
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	// Default UDP Farm of your frontend
	DefaultFarmId *float64 `pulumi:"defaultFarmId"`
	// Disable your frontend. Default: 'false'
	Disabled *bool `pulumi:"disabled"`
	// Human readable name for your frontend
	DisplayName *string `pulumi:"displayName"`
	// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
	Port string `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a UdpFrontend resource.
type UdpFrontendArgs struct {
	// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
	DedicatedIpfos pulumi.StringArrayInput
	// Default UDP Farm of your frontend
	DefaultFarmId pulumi.Float64PtrInput
	// Disable your frontend. Default: 'false'
	Disabled pulumi.BoolPtrInput
	// Human readable name for your frontend
	DisplayName pulumi.StringPtrInput
	// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
	Port pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringInput
}

func (UdpFrontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*udpFrontendArgs)(nil)).Elem()
}

type UdpFrontendInput interface {
	pulumi.Input

	ToUdpFrontendOutput() UdpFrontendOutput
	ToUdpFrontendOutputWithContext(ctx context.Context) UdpFrontendOutput
}

func (*UdpFrontend) ElementType() reflect.Type {
	return reflect.TypeOf((**UdpFrontend)(nil)).Elem()
}

func (i *UdpFrontend) ToUdpFrontendOutput() UdpFrontendOutput {
	return i.ToUdpFrontendOutputWithContext(context.Background())
}

func (i *UdpFrontend) ToUdpFrontendOutputWithContext(ctx context.Context) UdpFrontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFrontendOutput)
}

// UdpFrontendArrayInput is an input type that accepts UdpFrontendArray and UdpFrontendArrayOutput values.
// You can construct a concrete instance of `UdpFrontendArrayInput` via:
//
//	UdpFrontendArray{ UdpFrontendArgs{...} }
type UdpFrontendArrayInput interface {
	pulumi.Input

	ToUdpFrontendArrayOutput() UdpFrontendArrayOutput
	ToUdpFrontendArrayOutputWithContext(context.Context) UdpFrontendArrayOutput
}

type UdpFrontendArray []UdpFrontendInput

func (UdpFrontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UdpFrontend)(nil)).Elem()
}

func (i UdpFrontendArray) ToUdpFrontendArrayOutput() UdpFrontendArrayOutput {
	return i.ToUdpFrontendArrayOutputWithContext(context.Background())
}

func (i UdpFrontendArray) ToUdpFrontendArrayOutputWithContext(ctx context.Context) UdpFrontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFrontendArrayOutput)
}

// UdpFrontendMapInput is an input type that accepts UdpFrontendMap and UdpFrontendMapOutput values.
// You can construct a concrete instance of `UdpFrontendMapInput` via:
//
//	UdpFrontendMap{ "key": UdpFrontendArgs{...} }
type UdpFrontendMapInput interface {
	pulumi.Input

	ToUdpFrontendMapOutput() UdpFrontendMapOutput
	ToUdpFrontendMapOutputWithContext(context.Context) UdpFrontendMapOutput
}

type UdpFrontendMap map[string]UdpFrontendInput

func (UdpFrontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UdpFrontend)(nil)).Elem()
}

func (i UdpFrontendMap) ToUdpFrontendMapOutput() UdpFrontendMapOutput {
	return i.ToUdpFrontendMapOutputWithContext(context.Background())
}

func (i UdpFrontendMap) ToUdpFrontendMapOutputWithContext(ctx context.Context) UdpFrontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFrontendMapOutput)
}

type UdpFrontendOutput struct{ *pulumi.OutputState }

func (UdpFrontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UdpFrontend)(nil)).Elem()
}

func (o UdpFrontendOutput) ToUdpFrontendOutput() UdpFrontendOutput {
	return o
}

func (o UdpFrontendOutput) ToUdpFrontendOutputWithContext(ctx context.Context) UdpFrontendOutput {
	return o
}

// Only attach frontend on these ip. No restriction if null. List of Ip blocks.
func (o UdpFrontendOutput) DedicatedIpfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.StringArrayOutput { return v.DedicatedIpfos }).(pulumi.StringArrayOutput)
}

// Default UDP Farm of your frontend
func (o UdpFrontendOutput) DefaultFarmId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.Float64PtrOutput { return v.DefaultFarmId }).(pulumi.Float64PtrOutput)
}

// Disable your frontend. Default: 'false'
func (o UdpFrontendOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Human readable name for your frontend
func (o UdpFrontendOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Id of your frontend
func (o UdpFrontendOutput) FrontendId() pulumi.Float64Output {
	return o.ApplyT(func(v *UdpFrontend) pulumi.Float64Output { return v.FrontendId }).(pulumi.Float64Output)
}

// Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
func (o UdpFrontendOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o UdpFrontendOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
func (o UdpFrontendOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *UdpFrontend) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type UdpFrontendArrayOutput struct{ *pulumi.OutputState }

func (UdpFrontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UdpFrontend)(nil)).Elem()
}

func (o UdpFrontendArrayOutput) ToUdpFrontendArrayOutput() UdpFrontendArrayOutput {
	return o
}

func (o UdpFrontendArrayOutput) ToUdpFrontendArrayOutputWithContext(ctx context.Context) UdpFrontendArrayOutput {
	return o
}

func (o UdpFrontendArrayOutput) Index(i pulumi.IntInput) UdpFrontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UdpFrontend {
		return vs[0].([]*UdpFrontend)[vs[1].(int)]
	}).(UdpFrontendOutput)
}

type UdpFrontendMapOutput struct{ *pulumi.OutputState }

func (UdpFrontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UdpFrontend)(nil)).Elem()
}

func (o UdpFrontendMapOutput) ToUdpFrontendMapOutput() UdpFrontendMapOutput {
	return o
}

func (o UdpFrontendMapOutput) ToUdpFrontendMapOutputWithContext(ctx context.Context) UdpFrontendMapOutput {
	return o
}

func (o UdpFrontendMapOutput) MapIndex(k pulumi.StringInput) UdpFrontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UdpFrontend {
		return vs[0].(map[string]*UdpFrontend)[vs[1].(string)]
	}).(UdpFrontendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFrontendInput)(nil)).Elem(), &UdpFrontend{})
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFrontendArrayInput)(nil)).Elem(), UdpFrontendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFrontendMapInput)(nil)).Elem(), UdpFrontendMap{})
	pulumi.RegisterOutputType(UdpFrontendOutput{})
	pulumi.RegisterOutputType(UdpFrontendArrayOutput{})
	pulumi.RegisterOutputType(UdpFrontendMapOutput{})
}
