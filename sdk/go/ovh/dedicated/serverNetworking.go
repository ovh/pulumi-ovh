// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/internal"
)

// Manage dedicated server networking interface on SCALE and HIGH-GRADE range.
//
// !> The API route targeted by this resource are restricted to OVHCloud users (`Internal API`) with additional restrictions.
//
// ## Example Usage
//
// # The following example aims to bind all interfaces in a vRack
//
// The following example aims to attach the server to two different vRack.
//
// ## Import
//
// A dedicated server networking configuration can be imported using the `service_name`. bash <break><break>```sh<break> $ pulumi import ovh:Dedicated/serverNetworking:ServerNetworking server service_name <break>```<break><break>
type ServerNetworking struct {
	pulumi.CustomResourceState

	// Operation description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Interface or interfaces aggregation.
	Interfaces ServerNetworkingInterfaceArrayOutput `pulumi:"interfaces"`
	// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// status of the networking configuration (should be `active`).
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewServerNetworking registers a new resource with the given unique name, arguments, and options.
func NewServerNetworking(ctx *pulumi.Context,
	name string, args *ServerNetworkingArgs, opts ...pulumi.ResourceOption) (*ServerNetworking, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interfaces == nil {
		return nil, errors.New("invalid value for required argument 'Interfaces'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerNetworking
	err := ctx.RegisterResource("ovh:Dedicated/serverNetworking:ServerNetworking", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerNetworking gets an existing ServerNetworking resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerNetworking(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerNetworkingState, opts ...pulumi.ResourceOption) (*ServerNetworking, error) {
	var resource ServerNetworking
	err := ctx.ReadResource("ovh:Dedicated/serverNetworking:ServerNetworking", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerNetworking resources.
type serverNetworkingState struct {
	// Operation description.
	Description *string `pulumi:"description"`
	// Interface or interfaces aggregation.
	Interfaces []ServerNetworkingInterface `pulumi:"interfaces"`
	// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
	ServiceName *string `pulumi:"serviceName"`
	// status of the networking configuration (should be `active`).
	Status *string `pulumi:"status"`
}

type ServerNetworkingState struct {
	// Operation description.
	Description pulumi.StringPtrInput
	// Interface or interfaces aggregation.
	Interfaces ServerNetworkingInterfaceArrayInput
	// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
	ServiceName pulumi.StringPtrInput
	// status of the networking configuration (should be `active`).
	Status pulumi.StringPtrInput
}

func (ServerNetworkingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverNetworkingState)(nil)).Elem()
}

type serverNetworkingArgs struct {
	// Interface or interfaces aggregation.
	Interfaces []ServerNetworkingInterface `pulumi:"interfaces"`
	// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServerNetworking resource.
type ServerNetworkingArgs struct {
	// Interface or interfaces aggregation.
	Interfaces ServerNetworkingInterfaceArrayInput
	// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
	ServiceName pulumi.StringInput
}

func (ServerNetworkingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverNetworkingArgs)(nil)).Elem()
}

type ServerNetworkingInput interface {
	pulumi.Input

	ToServerNetworkingOutput() ServerNetworkingOutput
	ToServerNetworkingOutputWithContext(ctx context.Context) ServerNetworkingOutput
}

func (*ServerNetworking) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerNetworking)(nil)).Elem()
}

func (i *ServerNetworking) ToServerNetworkingOutput() ServerNetworkingOutput {
	return i.ToServerNetworkingOutputWithContext(context.Background())
}

func (i *ServerNetworking) ToServerNetworkingOutputWithContext(ctx context.Context) ServerNetworkingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerNetworkingOutput)
}

// ServerNetworkingArrayInput is an input type that accepts ServerNetworkingArray and ServerNetworkingArrayOutput values.
// You can construct a concrete instance of `ServerNetworkingArrayInput` via:
//
//	ServerNetworkingArray{ ServerNetworkingArgs{...} }
type ServerNetworkingArrayInput interface {
	pulumi.Input

	ToServerNetworkingArrayOutput() ServerNetworkingArrayOutput
	ToServerNetworkingArrayOutputWithContext(context.Context) ServerNetworkingArrayOutput
}

type ServerNetworkingArray []ServerNetworkingInput

func (ServerNetworkingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerNetworking)(nil)).Elem()
}

func (i ServerNetworkingArray) ToServerNetworkingArrayOutput() ServerNetworkingArrayOutput {
	return i.ToServerNetworkingArrayOutputWithContext(context.Background())
}

func (i ServerNetworkingArray) ToServerNetworkingArrayOutputWithContext(ctx context.Context) ServerNetworkingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerNetworkingArrayOutput)
}

// ServerNetworkingMapInput is an input type that accepts ServerNetworkingMap and ServerNetworkingMapOutput values.
// You can construct a concrete instance of `ServerNetworkingMapInput` via:
//
//	ServerNetworkingMap{ "key": ServerNetworkingArgs{...} }
type ServerNetworkingMapInput interface {
	pulumi.Input

	ToServerNetworkingMapOutput() ServerNetworkingMapOutput
	ToServerNetworkingMapOutputWithContext(context.Context) ServerNetworkingMapOutput
}

type ServerNetworkingMap map[string]ServerNetworkingInput

func (ServerNetworkingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerNetworking)(nil)).Elem()
}

func (i ServerNetworkingMap) ToServerNetworkingMapOutput() ServerNetworkingMapOutput {
	return i.ToServerNetworkingMapOutputWithContext(context.Background())
}

func (i ServerNetworkingMap) ToServerNetworkingMapOutputWithContext(ctx context.Context) ServerNetworkingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerNetworkingMapOutput)
}

type ServerNetworkingOutput struct{ *pulumi.OutputState }

func (ServerNetworkingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerNetworking)(nil)).Elem()
}

func (o ServerNetworkingOutput) ToServerNetworkingOutput() ServerNetworkingOutput {
	return o
}

func (o ServerNetworkingOutput) ToServerNetworkingOutputWithContext(ctx context.Context) ServerNetworkingOutput {
	return o
}

// Operation description.
func (o ServerNetworkingOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerNetworking) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Interface or interfaces aggregation.
func (o ServerNetworkingOutput) Interfaces() ServerNetworkingInterfaceArrayOutput {
	return o.ApplyT(func(v *ServerNetworking) ServerNetworkingInterfaceArrayOutput { return v.Interfaces }).(ServerNetworkingInterfaceArrayOutput)
}

// The serviceName of your dedicated server. The full list of available dedicated servers can be found using the `getServers` datasource.
func (o ServerNetworkingOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerNetworking) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// status of the networking configuration (should be `active`).
func (o ServerNetworkingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerNetworking) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ServerNetworkingArrayOutput struct{ *pulumi.OutputState }

func (ServerNetworkingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerNetworking)(nil)).Elem()
}

func (o ServerNetworkingArrayOutput) ToServerNetworkingArrayOutput() ServerNetworkingArrayOutput {
	return o
}

func (o ServerNetworkingArrayOutput) ToServerNetworkingArrayOutputWithContext(ctx context.Context) ServerNetworkingArrayOutput {
	return o
}

func (o ServerNetworkingArrayOutput) Index(i pulumi.IntInput) ServerNetworkingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerNetworking {
		return vs[0].([]*ServerNetworking)[vs[1].(int)]
	}).(ServerNetworkingOutput)
}

type ServerNetworkingMapOutput struct{ *pulumi.OutputState }

func (ServerNetworkingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerNetworking)(nil)).Elem()
}

func (o ServerNetworkingMapOutput) ToServerNetworkingMapOutput() ServerNetworkingMapOutput {
	return o
}

func (o ServerNetworkingMapOutput) ToServerNetworkingMapOutputWithContext(ctx context.Context) ServerNetworkingMapOutput {
	return o
}

func (o ServerNetworkingMapOutput) MapIndex(k pulumi.StringInput) ServerNetworkingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerNetworking {
		return vs[0].(map[string]*ServerNetworking)[vs[1].(string)]
	}).(ServerNetworkingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerNetworkingInput)(nil)).Elem(), &ServerNetworking{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerNetworkingArrayInput)(nil)).Elem(), ServerNetworkingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerNetworkingMapInput)(nil)).Elem(), ServerNetworkingMap{})
	pulumi.RegisterOutputType(ServerNetworkingOutput{})
	pulumi.RegisterOutputType(ServerNetworkingArrayOutput{})
	pulumi.RegisterOutputType(ServerNetworkingMapOutput{})
}
