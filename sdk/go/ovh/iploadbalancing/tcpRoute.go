// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage TCP route for a loadbalancer service
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
//			_, err := iploadbalancing.NewTcpRoute(ctx, "tcp_reject", &iploadbalancing.TcpRouteArgs{
//				ServiceName: pulumi.String("loadbalancer-xxxxxxxxxxxxxxxxxx"),
//				Weight:      pulumi.Int(1),
//				Action: &iploadbalancing.TcpRouteActionArgs{
//					Type: pulumi.String("reject"),
//				},
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
// TCP route can be imported using the following format `service_name` and the `id` of the route separated by "/" e.g.
//
// bash
//
// ```sh
// $ pulumi import ovh:IpLoadBalancing/tcpRoute:TcpRoute tcpreject service_name/route_id
// ```
type TcpRoute struct {
	pulumi.CustomResourceState

	// Action triggered when all rules match
	Action TcpRouteActionOutput `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId pulumi.IntOutput `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules TcpRouteRuleTypeArrayOutput `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringOutput `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewTcpRoute registers a new resource with the given unique name, arguments, and options.
func NewTcpRoute(ctx *pulumi.Context,
	name string, args *TcpRouteArgs, opts ...pulumi.ResourceOption) (*TcpRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TcpRoute
	err := ctx.RegisterResource("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTcpRoute gets an existing TcpRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTcpRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TcpRouteState, opts ...pulumi.ResourceOption) (*TcpRoute, error) {
	var resource TcpRoute
	err := ctx.ReadResource("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TcpRoute resources.
type tcpRouteState struct {
	// Action triggered when all rules match
	Action *TcpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules []TcpRouteRuleType `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status *string `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight *int `pulumi:"weight"`
}

type TcpRouteState struct {
	// Action triggered when all rules match
	Action TcpRouteActionPtrInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// List of rules to match to trigger action
	Rules TcpRouteRuleTypeArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringPtrInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntPtrInput
}

func (TcpRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteState)(nil)).Elem()
}

type tcpRouteArgs struct {
	// Action triggered when all rules match
	Action TcpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a TcpRoute resource.
type TcpRouteArgs struct {
	// Action triggered when all rules match
	Action TcpRouteActionInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
	Weight pulumi.IntPtrInput
}

func (TcpRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteArgs)(nil)).Elem()
}

type TcpRouteInput interface {
	pulumi.Input

	ToTcpRouteOutput() TcpRouteOutput
	ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput
}

func (*TcpRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRoute)(nil)).Elem()
}

func (i *TcpRoute) ToTcpRouteOutput() TcpRouteOutput {
	return i.ToTcpRouteOutputWithContext(context.Background())
}

func (i *TcpRoute) ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteOutput)
}

// TcpRouteArrayInput is an input type that accepts TcpRouteArray and TcpRouteArrayOutput values.
// You can construct a concrete instance of `TcpRouteArrayInput` via:
//
//	TcpRouteArray{ TcpRouteArgs{...} }
type TcpRouteArrayInput interface {
	pulumi.Input

	ToTcpRouteArrayOutput() TcpRouteArrayOutput
	ToTcpRouteArrayOutputWithContext(context.Context) TcpRouteArrayOutput
}

type TcpRouteArray []TcpRouteInput

func (TcpRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpRoute)(nil)).Elem()
}

func (i TcpRouteArray) ToTcpRouteArrayOutput() TcpRouteArrayOutput {
	return i.ToTcpRouteArrayOutputWithContext(context.Background())
}

func (i TcpRouteArray) ToTcpRouteArrayOutputWithContext(ctx context.Context) TcpRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteArrayOutput)
}

// TcpRouteMapInput is an input type that accepts TcpRouteMap and TcpRouteMapOutput values.
// You can construct a concrete instance of `TcpRouteMapInput` via:
//
//	TcpRouteMap{ "key": TcpRouteArgs{...} }
type TcpRouteMapInput interface {
	pulumi.Input

	ToTcpRouteMapOutput() TcpRouteMapOutput
	ToTcpRouteMapOutputWithContext(context.Context) TcpRouteMapOutput
}

type TcpRouteMap map[string]TcpRouteInput

func (TcpRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpRoute)(nil)).Elem()
}

func (i TcpRouteMap) ToTcpRouteMapOutput() TcpRouteMapOutput {
	return i.ToTcpRouteMapOutputWithContext(context.Background())
}

func (i TcpRouteMap) ToTcpRouteMapOutputWithContext(ctx context.Context) TcpRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteMapOutput)
}

type TcpRouteOutput struct{ *pulumi.OutputState }

func (TcpRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRoute)(nil)).Elem()
}

func (o TcpRouteOutput) ToTcpRouteOutput() TcpRouteOutput {
	return o
}

func (o TcpRouteOutput) ToTcpRouteOutputWithContext(ctx context.Context) TcpRouteOutput {
	return o
}

// Action triggered when all rules match
func (o TcpRouteOutput) Action() TcpRouteActionOutput {
	return o.ApplyT(func(v *TcpRoute) TcpRouteActionOutput { return v.Action }).(TcpRouteActionOutput)
}

// Human readable name for your route, this field is for you
func (o TcpRouteOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Route traffic for this frontend
func (o TcpRouteOutput) FrontendId() pulumi.IntOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.IntOutput { return v.FrontendId }).(pulumi.IntOutput)
}

// List of rules to match to trigger action
func (o TcpRouteOutput) Rules() TcpRouteRuleTypeArrayOutput {
	return o.ApplyT(func(v *TcpRoute) TcpRouteRuleTypeArrayOutput { return v.Rules }).(TcpRouteRuleTypeArrayOutput)
}

// The internal name of your IP load balancing
func (o TcpRouteOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Route status. Routes in "ok" state are ready to operate
func (o TcpRouteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
func (o TcpRouteOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *TcpRoute) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type TcpRouteArrayOutput struct{ *pulumi.OutputState }

func (TcpRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpRoute)(nil)).Elem()
}

func (o TcpRouteArrayOutput) ToTcpRouteArrayOutput() TcpRouteArrayOutput {
	return o
}

func (o TcpRouteArrayOutput) ToTcpRouteArrayOutputWithContext(ctx context.Context) TcpRouteArrayOutput {
	return o
}

func (o TcpRouteArrayOutput) Index(i pulumi.IntInput) TcpRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TcpRoute {
		return vs[0].([]*TcpRoute)[vs[1].(int)]
	}).(TcpRouteOutput)
}

type TcpRouteMapOutput struct{ *pulumi.OutputState }

func (TcpRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpRoute)(nil)).Elem()
}

func (o TcpRouteMapOutput) ToTcpRouteMapOutput() TcpRouteMapOutput {
	return o
}

func (o TcpRouteMapOutput) ToTcpRouteMapOutputWithContext(ctx context.Context) TcpRouteMapOutput {
	return o
}

func (o TcpRouteMapOutput) MapIndex(k pulumi.StringInput) TcpRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TcpRoute {
		return vs[0].(map[string]*TcpRoute)[vs[1].(string)]
	}).(TcpRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteInput)(nil)).Elem(), &TcpRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteArrayInput)(nil)).Elem(), TcpRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteMapInput)(nil)).Elem(), TcpRouteMap{})
	pulumi.RegisterOutputType(TcpRouteOutput{})
	pulumi.RegisterOutputType(TcpRouteArrayOutput{})
	pulumi.RegisterOutputType(TcpRouteMapOutput{})
}
