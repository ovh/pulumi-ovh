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

type HttpFarmServer struct {
	pulumi.CustomResourceState

	Address              pulumi.StringOutput    `pulumi:"address"`
	Backup               pulumi.BoolPtrOutput   `pulumi:"backup"`
	Chain                pulumi.StringPtrOutput `pulumi:"chain"`
	Cookie               pulumi.StringOutput    `pulumi:"cookie"`
	DisplayName          pulumi.StringPtrOutput `pulumi:"displayName"`
	FarmId               pulumi.IntOutput       `pulumi:"farmId"`
	OnMarkedDown         pulumi.StringPtrOutput `pulumi:"onMarkedDown"`
	Port                 pulumi.IntPtrOutput    `pulumi:"port"`
	Probe                pulumi.BoolPtrOutput   `pulumi:"probe"`
	ProxyProtocolVersion pulumi.StringPtrOutput `pulumi:"proxyProtocolVersion"`
	ServiceName          pulumi.StringOutput    `pulumi:"serviceName"`
	Ssl                  pulumi.BoolPtrOutput   `pulumi:"ssl"`
	Status               pulumi.StringOutput    `pulumi:"status"`
	Weight               pulumi.IntPtrOutput    `pulumi:"weight"`
}

// NewHttpFarmServer registers a new resource with the given unique name, arguments, and options.
func NewHttpFarmServer(ctx *pulumi.Context,
	name string, args *HttpFarmServerArgs, opts ...pulumi.ResourceOption) (*HttpFarmServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.FarmId == nil {
		return nil, errors.New("invalid value for required argument 'FarmId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HttpFarmServer
	err := ctx.RegisterResource("ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpFarmServer gets an existing HttpFarmServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpFarmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpFarmServerState, opts ...pulumi.ResourceOption) (*HttpFarmServer, error) {
	var resource HttpFarmServer
	err := ctx.ReadResource("ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpFarmServer resources.
type httpFarmServerState struct {
	Address              *string `pulumi:"address"`
	Backup               *bool   `pulumi:"backup"`
	Chain                *string `pulumi:"chain"`
	Cookie               *string `pulumi:"cookie"`
	DisplayName          *string `pulumi:"displayName"`
	FarmId               *int    `pulumi:"farmId"`
	OnMarkedDown         *string `pulumi:"onMarkedDown"`
	Port                 *int    `pulumi:"port"`
	Probe                *bool   `pulumi:"probe"`
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	ServiceName          *string `pulumi:"serviceName"`
	Ssl                  *bool   `pulumi:"ssl"`
	Status               *string `pulumi:"status"`
	Weight               *int    `pulumi:"weight"`
}

type HttpFarmServerState struct {
	Address              pulumi.StringPtrInput
	Backup               pulumi.BoolPtrInput
	Chain                pulumi.StringPtrInput
	Cookie               pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	FarmId               pulumi.IntPtrInput
	OnMarkedDown         pulumi.StringPtrInput
	Port                 pulumi.IntPtrInput
	Probe                pulumi.BoolPtrInput
	ProxyProtocolVersion pulumi.StringPtrInput
	ServiceName          pulumi.StringPtrInput
	Ssl                  pulumi.BoolPtrInput
	Status               pulumi.StringPtrInput
	Weight               pulumi.IntPtrInput
}

func (HttpFarmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFarmServerState)(nil)).Elem()
}

type httpFarmServerArgs struct {
	Address              string  `pulumi:"address"`
	Backup               *bool   `pulumi:"backup"`
	Chain                *string `pulumi:"chain"`
	DisplayName          *string `pulumi:"displayName"`
	FarmId               int     `pulumi:"farmId"`
	OnMarkedDown         *string `pulumi:"onMarkedDown"`
	Port                 *int    `pulumi:"port"`
	Probe                *bool   `pulumi:"probe"`
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	ServiceName          string  `pulumi:"serviceName"`
	Ssl                  *bool   `pulumi:"ssl"`
	Status               string  `pulumi:"status"`
	Weight               *int    `pulumi:"weight"`
}

// The set of arguments for constructing a HttpFarmServer resource.
type HttpFarmServerArgs struct {
	Address              pulumi.StringInput
	Backup               pulumi.BoolPtrInput
	Chain                pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	FarmId               pulumi.IntInput
	OnMarkedDown         pulumi.StringPtrInput
	Port                 pulumi.IntPtrInput
	Probe                pulumi.BoolPtrInput
	ProxyProtocolVersion pulumi.StringPtrInput
	ServiceName          pulumi.StringInput
	Ssl                  pulumi.BoolPtrInput
	Status               pulumi.StringInput
	Weight               pulumi.IntPtrInput
}

func (HttpFarmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFarmServerArgs)(nil)).Elem()
}

type HttpFarmServerInput interface {
	pulumi.Input

	ToHttpFarmServerOutput() HttpFarmServerOutput
	ToHttpFarmServerOutputWithContext(ctx context.Context) HttpFarmServerOutput
}

func (*HttpFarmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFarmServer)(nil)).Elem()
}

func (i *HttpFarmServer) ToHttpFarmServerOutput() HttpFarmServerOutput {
	return i.ToHttpFarmServerOutputWithContext(context.Background())
}

func (i *HttpFarmServer) ToHttpFarmServerOutputWithContext(ctx context.Context) HttpFarmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmServerOutput)
}

// HttpFarmServerArrayInput is an input type that accepts HttpFarmServerArray and HttpFarmServerArrayOutput values.
// You can construct a concrete instance of `HttpFarmServerArrayInput` via:
//
//	HttpFarmServerArray{ HttpFarmServerArgs{...} }
type HttpFarmServerArrayInput interface {
	pulumi.Input

	ToHttpFarmServerArrayOutput() HttpFarmServerArrayOutput
	ToHttpFarmServerArrayOutputWithContext(context.Context) HttpFarmServerArrayOutput
}

type HttpFarmServerArray []HttpFarmServerInput

func (HttpFarmServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFarmServer)(nil)).Elem()
}

func (i HttpFarmServerArray) ToHttpFarmServerArrayOutput() HttpFarmServerArrayOutput {
	return i.ToHttpFarmServerArrayOutputWithContext(context.Background())
}

func (i HttpFarmServerArray) ToHttpFarmServerArrayOutputWithContext(ctx context.Context) HttpFarmServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmServerArrayOutput)
}

// HttpFarmServerMapInput is an input type that accepts HttpFarmServerMap and HttpFarmServerMapOutput values.
// You can construct a concrete instance of `HttpFarmServerMapInput` via:
//
//	HttpFarmServerMap{ "key": HttpFarmServerArgs{...} }
type HttpFarmServerMapInput interface {
	pulumi.Input

	ToHttpFarmServerMapOutput() HttpFarmServerMapOutput
	ToHttpFarmServerMapOutputWithContext(context.Context) HttpFarmServerMapOutput
}

type HttpFarmServerMap map[string]HttpFarmServerInput

func (HttpFarmServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFarmServer)(nil)).Elem()
}

func (i HttpFarmServerMap) ToHttpFarmServerMapOutput() HttpFarmServerMapOutput {
	return i.ToHttpFarmServerMapOutputWithContext(context.Background())
}

func (i HttpFarmServerMap) ToHttpFarmServerMapOutputWithContext(ctx context.Context) HttpFarmServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmServerMapOutput)
}

type HttpFarmServerOutput struct{ *pulumi.OutputState }

func (HttpFarmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFarmServer)(nil)).Elem()
}

func (o HttpFarmServerOutput) ToHttpFarmServerOutput() HttpFarmServerOutput {
	return o
}

func (o HttpFarmServerOutput) ToHttpFarmServerOutputWithContext(ctx context.Context) HttpFarmServerOutput {
	return o
}

func (o HttpFarmServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

func (o HttpFarmServerOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o HttpFarmServerOutput) Chain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.Chain }).(pulumi.StringPtrOutput)
}

func (o HttpFarmServerOutput) Cookie() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Cookie }).(pulumi.StringOutput)
}

func (o HttpFarmServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o HttpFarmServerOutput) FarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.IntOutput { return v.FarmId }).(pulumi.IntOutput)
}

func (o HttpFarmServerOutput) OnMarkedDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.OnMarkedDown }).(pulumi.StringPtrOutput)
}

func (o HttpFarmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o HttpFarmServerOutput) Probe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Probe }).(pulumi.BoolPtrOutput)
}

func (o HttpFarmServerOutput) ProxyProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.ProxyProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o HttpFarmServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o HttpFarmServerOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

func (o HttpFarmServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o HttpFarmServerOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

type HttpFarmServerArrayOutput struct{ *pulumi.OutputState }

func (HttpFarmServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFarmServer)(nil)).Elem()
}

func (o HttpFarmServerArrayOutput) ToHttpFarmServerArrayOutput() HttpFarmServerArrayOutput {
	return o
}

func (o HttpFarmServerArrayOutput) ToHttpFarmServerArrayOutputWithContext(ctx context.Context) HttpFarmServerArrayOutput {
	return o
}

func (o HttpFarmServerArrayOutput) Index(i pulumi.IntInput) HttpFarmServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpFarmServer {
		return vs[0].([]*HttpFarmServer)[vs[1].(int)]
	}).(HttpFarmServerOutput)
}

type HttpFarmServerMapOutput struct{ *pulumi.OutputState }

func (HttpFarmServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFarmServer)(nil)).Elem()
}

func (o HttpFarmServerMapOutput) ToHttpFarmServerMapOutput() HttpFarmServerMapOutput {
	return o
}

func (o HttpFarmServerMapOutput) ToHttpFarmServerMapOutputWithContext(ctx context.Context) HttpFarmServerMapOutput {
	return o
}

func (o HttpFarmServerMapOutput) MapIndex(k pulumi.StringInput) HttpFarmServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpFarmServer {
		return vs[0].(map[string]*HttpFarmServer)[vs[1].(string)]
	}).(HttpFarmServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmServerInput)(nil)).Elem(), &HttpFarmServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmServerArrayInput)(nil)).Elem(), HttpFarmServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmServerMapInput)(nil)).Elem(), HttpFarmServerMap{})
	pulumi.RegisterOutputType(HttpFarmServerOutput{})
	pulumi.RegisterOutputType(HttpFarmServerArrayOutput{})
	pulumi.RegisterOutputType(HttpFarmServerMapOutput{})
}
