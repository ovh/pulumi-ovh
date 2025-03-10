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

type Services struct {
	pulumi.CustomResourceState

	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// vrackServices service name
	VrackServices pulumi.StringOutput `pulumi:"vrackServices"`
}

// NewServices registers a new resource with the given unique name, arguments, and options.
func NewServices(ctx *pulumi.Context,
	name string, args *ServicesArgs, opts ...pulumi.ResourceOption) (*Services, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VrackServices == nil {
		return nil, errors.New("invalid value for required argument 'VrackServices'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Services
	err := ctx.RegisterResource("ovh:Vrack/services:Services", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServices gets an existing Services resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServices(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicesState, opts ...pulumi.ResourceOption) (*Services, error) {
	var resource Services
	err := ctx.ReadResource("ovh:Vrack/services:Services", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Services resources.
type servicesState struct {
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
	// vrackServices service name
	VrackServices *string `pulumi:"vrackServices"`
}

type ServicesState struct {
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
	// vrackServices service name
	VrackServices pulumi.StringPtrInput
}

func (ServicesState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicesState)(nil)).Elem()
}

type servicesArgs struct {
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
	// vrackServices service name
	VrackServices string `pulumi:"vrackServices"`
}

// The set of arguments for constructing a Services resource.
type ServicesArgs struct {
	// The internal name of your vrack
	ServiceName pulumi.StringInput
	// vrackServices service name
	VrackServices pulumi.StringInput
}

func (ServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicesArgs)(nil)).Elem()
}

type ServicesInput interface {
	pulumi.Input

	ToServicesOutput() ServicesOutput
	ToServicesOutputWithContext(ctx context.Context) ServicesOutput
}

func (*Services) ElementType() reflect.Type {
	return reflect.TypeOf((**Services)(nil)).Elem()
}

func (i *Services) ToServicesOutput() ServicesOutput {
	return i.ToServicesOutputWithContext(context.Background())
}

func (i *Services) ToServicesOutputWithContext(ctx context.Context) ServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesOutput)
}

// ServicesArrayInput is an input type that accepts ServicesArray and ServicesArrayOutput values.
// You can construct a concrete instance of `ServicesArrayInput` via:
//
//	ServicesArray{ ServicesArgs{...} }
type ServicesArrayInput interface {
	pulumi.Input

	ToServicesArrayOutput() ServicesArrayOutput
	ToServicesArrayOutputWithContext(context.Context) ServicesArrayOutput
}

type ServicesArray []ServicesInput

func (ServicesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Services)(nil)).Elem()
}

func (i ServicesArray) ToServicesArrayOutput() ServicesArrayOutput {
	return i.ToServicesArrayOutputWithContext(context.Background())
}

func (i ServicesArray) ToServicesArrayOutputWithContext(ctx context.Context) ServicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesArrayOutput)
}

// ServicesMapInput is an input type that accepts ServicesMap and ServicesMapOutput values.
// You can construct a concrete instance of `ServicesMapInput` via:
//
//	ServicesMap{ "key": ServicesArgs{...} }
type ServicesMapInput interface {
	pulumi.Input

	ToServicesMapOutput() ServicesMapOutput
	ToServicesMapOutputWithContext(context.Context) ServicesMapOutput
}

type ServicesMap map[string]ServicesInput

func (ServicesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Services)(nil)).Elem()
}

func (i ServicesMap) ToServicesMapOutput() ServicesMapOutput {
	return i.ToServicesMapOutputWithContext(context.Background())
}

func (i ServicesMap) ToServicesMapOutputWithContext(ctx context.Context) ServicesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesMapOutput)
}

type ServicesOutput struct{ *pulumi.OutputState }

func (ServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Services)(nil)).Elem()
}

func (o ServicesOutput) ToServicesOutput() ServicesOutput {
	return o
}

func (o ServicesOutput) ToServicesOutputWithContext(ctx context.Context) ServicesOutput {
	return o
}

// The internal name of your vrack
func (o ServicesOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Services) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// vrackServices service name
func (o ServicesOutput) VrackServices() pulumi.StringOutput {
	return o.ApplyT(func(v *Services) pulumi.StringOutput { return v.VrackServices }).(pulumi.StringOutput)
}

type ServicesArrayOutput struct{ *pulumi.OutputState }

func (ServicesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Services)(nil)).Elem()
}

func (o ServicesArrayOutput) ToServicesArrayOutput() ServicesArrayOutput {
	return o
}

func (o ServicesArrayOutput) ToServicesArrayOutputWithContext(ctx context.Context) ServicesArrayOutput {
	return o
}

func (o ServicesArrayOutput) Index(i pulumi.IntInput) ServicesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Services {
		return vs[0].([]*Services)[vs[1].(int)]
	}).(ServicesOutput)
}

type ServicesMapOutput struct{ *pulumi.OutputState }

func (ServicesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Services)(nil)).Elem()
}

func (o ServicesMapOutput) ToServicesMapOutput() ServicesMapOutput {
	return o
}

func (o ServicesMapOutput) ToServicesMapOutputWithContext(ctx context.Context) ServicesMapOutput {
	return o
}

func (o ServicesMapOutput) MapIndex(k pulumi.StringInput) ServicesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Services {
		return vs[0].(map[string]*Services)[vs[1].(string)]
	}).(ServicesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesInput)(nil)).Elem(), &Services{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesArrayInput)(nil)).Elem(), ServicesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicesMapInput)(nil)).Elem(), ServicesMap{})
	pulumi.RegisterOutputType(ServicesOutput{})
	pulumi.RegisterOutputType(ServicesArrayOutput{})
	pulumi.RegisterOutputType(ServicesMapOutput{})
}
