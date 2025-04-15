// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach an IP Load Balancing to a VRack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/vrack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vrack.NewIpLoadbalancing(ctx, "viplb", &vrack.IpLoadbalancingArgs{
//				ServiceName:     pulumi.String("xxx"),
//				LoadbalancingId: pulumi.String("yyy"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpLoadbalancing struct {
	pulumi.CustomResourceState

	// The id of the IP Load Balancing.
	LoadbalancingId pulumi.StringOutput `pulumi:"LoadbalancingId"`
	// The id of the vrack.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewIpLoadbalancing registers a new resource with the given unique name, arguments, and options.
func NewIpLoadbalancing(ctx *pulumi.Context,
	name string, args *IpLoadbalancingArgs, opts ...pulumi.ResourceOption) (*IpLoadbalancing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadbalancingId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancingId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpLoadbalancing
	err := ctx.RegisterResource("ovh:Vrack/ipLoadbalancing:IpLoadbalancing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadbalancing gets an existing IpLoadbalancing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadbalancing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadbalancingState, opts ...pulumi.ResourceOption) (*IpLoadbalancing, error) {
	var resource IpLoadbalancing
	err := ctx.ReadResource("ovh:Vrack/ipLoadbalancing:IpLoadbalancing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadbalancing resources.
type ipLoadbalancingState struct {
	// The id of the IP Load Balancing.
	LoadbalancingId *string `pulumi:"LoadbalancingId"`
	// The id of the vrack.
	ServiceName *string `pulumi:"serviceName"`
}

type IpLoadbalancingState struct {
	// The id of the IP Load Balancing.
	LoadbalancingId pulumi.StringPtrInput
	// The id of the vrack.
	ServiceName pulumi.StringPtrInput
}

func (IpLoadbalancingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadbalancingState)(nil)).Elem()
}

type ipLoadbalancingArgs struct {
	// The id of the IP Load Balancing.
	LoadbalancingId string `pulumi:"LoadbalancingId"`
	// The id of the vrack.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a IpLoadbalancing resource.
type IpLoadbalancingArgs struct {
	// The id of the IP Load Balancing.
	LoadbalancingId pulumi.StringInput
	// The id of the vrack.
	ServiceName pulumi.StringInput
}

func (IpLoadbalancingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadbalancingArgs)(nil)).Elem()
}

type IpLoadbalancingInput interface {
	pulumi.Input

	ToIpLoadbalancingOutput() IpLoadbalancingOutput
	ToIpLoadbalancingOutputWithContext(ctx context.Context) IpLoadbalancingOutput
}

func (*IpLoadbalancing) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadbalancing)(nil)).Elem()
}

func (i *IpLoadbalancing) ToIpLoadbalancingOutput() IpLoadbalancingOutput {
	return i.ToIpLoadbalancingOutputWithContext(context.Background())
}

func (i *IpLoadbalancing) ToIpLoadbalancingOutputWithContext(ctx context.Context) IpLoadbalancingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadbalancingOutput)
}

// IpLoadbalancingArrayInput is an input type that accepts IpLoadbalancingArray and IpLoadbalancingArrayOutput values.
// You can construct a concrete instance of `IpLoadbalancingArrayInput` via:
//
//	IpLoadbalancingArray{ IpLoadbalancingArgs{...} }
type IpLoadbalancingArrayInput interface {
	pulumi.Input

	ToIpLoadbalancingArrayOutput() IpLoadbalancingArrayOutput
	ToIpLoadbalancingArrayOutputWithContext(context.Context) IpLoadbalancingArrayOutput
}

type IpLoadbalancingArray []IpLoadbalancingInput

func (IpLoadbalancingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadbalancing)(nil)).Elem()
}

func (i IpLoadbalancingArray) ToIpLoadbalancingArrayOutput() IpLoadbalancingArrayOutput {
	return i.ToIpLoadbalancingArrayOutputWithContext(context.Background())
}

func (i IpLoadbalancingArray) ToIpLoadbalancingArrayOutputWithContext(ctx context.Context) IpLoadbalancingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadbalancingArrayOutput)
}

// IpLoadbalancingMapInput is an input type that accepts IpLoadbalancingMap and IpLoadbalancingMapOutput values.
// You can construct a concrete instance of `IpLoadbalancingMapInput` via:
//
//	IpLoadbalancingMap{ "key": IpLoadbalancingArgs{...} }
type IpLoadbalancingMapInput interface {
	pulumi.Input

	ToIpLoadbalancingMapOutput() IpLoadbalancingMapOutput
	ToIpLoadbalancingMapOutputWithContext(context.Context) IpLoadbalancingMapOutput
}

type IpLoadbalancingMap map[string]IpLoadbalancingInput

func (IpLoadbalancingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadbalancing)(nil)).Elem()
}

func (i IpLoadbalancingMap) ToIpLoadbalancingMapOutput() IpLoadbalancingMapOutput {
	return i.ToIpLoadbalancingMapOutputWithContext(context.Background())
}

func (i IpLoadbalancingMap) ToIpLoadbalancingMapOutputWithContext(ctx context.Context) IpLoadbalancingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadbalancingMapOutput)
}

type IpLoadbalancingOutput struct{ *pulumi.OutputState }

func (IpLoadbalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadbalancing)(nil)).Elem()
}

func (o IpLoadbalancingOutput) ToIpLoadbalancingOutput() IpLoadbalancingOutput {
	return o
}

func (o IpLoadbalancingOutput) ToIpLoadbalancingOutputWithContext(ctx context.Context) IpLoadbalancingOutput {
	return o
}

// The id of the IP Load Balancing.
func (o IpLoadbalancingOutput) LoadbalancingId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadbalancing) pulumi.StringOutput { return v.LoadbalancingId }).(pulumi.StringOutput)
}

// The id of the vrack.
func (o IpLoadbalancingOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadbalancing) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type IpLoadbalancingArrayOutput struct{ *pulumi.OutputState }

func (IpLoadbalancingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadbalancing)(nil)).Elem()
}

func (o IpLoadbalancingArrayOutput) ToIpLoadbalancingArrayOutput() IpLoadbalancingArrayOutput {
	return o
}

func (o IpLoadbalancingArrayOutput) ToIpLoadbalancingArrayOutputWithContext(ctx context.Context) IpLoadbalancingArrayOutput {
	return o
}

func (o IpLoadbalancingArrayOutput) Index(i pulumi.IntInput) IpLoadbalancingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadbalancing {
		return vs[0].([]*IpLoadbalancing)[vs[1].(int)]
	}).(IpLoadbalancingOutput)
}

type IpLoadbalancingMapOutput struct{ *pulumi.OutputState }

func (IpLoadbalancingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadbalancing)(nil)).Elem()
}

func (o IpLoadbalancingMapOutput) ToIpLoadbalancingMapOutput() IpLoadbalancingMapOutput {
	return o
}

func (o IpLoadbalancingMapOutput) ToIpLoadbalancingMapOutputWithContext(ctx context.Context) IpLoadbalancingMapOutput {
	return o
}

func (o IpLoadbalancingMapOutput) MapIndex(k pulumi.StringInput) IpLoadbalancingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadbalancing {
		return vs[0].(map[string]*IpLoadbalancing)[vs[1].(string)]
	}).(IpLoadbalancingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadbalancingInput)(nil)).Elem(), &IpLoadbalancing{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadbalancingArrayInput)(nil)).Elem(), IpLoadbalancingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadbalancingMapInput)(nil)).Elem(), IpLoadbalancingMap{})
	pulumi.RegisterOutputType(IpLoadbalancingOutput{})
	pulumi.RegisterOutputType(IpLoadbalancingArrayOutput{})
	pulumi.RegisterOutputType(IpLoadbalancingMapOutput{})
}
