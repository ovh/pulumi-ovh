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

// Creates a backend server entry linked to HTTP loadbalancing group (farm)
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
//			farmname, err := iploadbalancing.NewHttpFarm(ctx, "farmname", &iploadbalancing.HttpFarmArgs{
//				ServiceName: pulumi.String(lb.ServiceName),
//				Port:        pulumi.Int(8080),
//				Zone:        pulumi.String("all"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iploadbalancing.NewHttpFarmServer(ctx, "backend", &iploadbalancing.HttpFarmServerArgs{
//				ServiceName:          pulumi.String(lb.ServiceName),
//				FarmId:               farmname.ID(),
//				DisplayName:          pulumi.String("mybackend"),
//				Address:              pulumi.String("4.5.6.7"),
//				Status:               pulumi.String("active"),
//				Port:                 pulumi.Int(80),
//				ProxyProtocolVersion: pulumi.String("v2"),
//				Weight:               pulumi.Int(2),
//				Probe:                pulumi.Bool(true),
//				Ssl:                  pulumi.Bool(false),
//				Backup:               pulumi.Bool(true),
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
// HTTP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.
//
// bash
//
// ```sh
// $ pulumi import ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer backend service_name/farm_id/server_id
// ```
type HttpFarmServer struct {
	pulumi.CustomResourceState

	// Address of the backend server (IP from either internal or OVHcloud network)
	Address pulumi.StringOutput `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrOutput   `pulumi:"backup"`
	Chain  pulumi.StringPtrOutput `pulumi:"chain"`
	// Value of the stickiness cookie used for this backend.
	Cookie pulumi.StringOutput `pulumi:"cookie"`
	// Label for the server
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId pulumi.IntOutput `pulumi:"farmId"`
	// enable action when backend marked down. (`shutdown-sessions`)
	OnMarkedDown pulumi.StringPtrOutput `pulumi:"onMarkedDown"`
	// Port that backend will respond on
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrOutput `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrOutput `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status pulumi.StringOutput `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
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
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address *string `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup *bool   `pulumi:"backup"`
	Chain  *string `pulumi:"chain"`
	// Value of the stickiness cookie used for this backend.
	Cookie *string `pulumi:"cookie"`
	// Label for the server
	DisplayName *string `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId *int `pulumi:"farmId"`
	// enable action when backend marked down. (`shutdown-sessions`)
	OnMarkedDown *string `pulumi:"onMarkedDown"`
	// Port that backend will respond on
	Port *int `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe *bool `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl *bool `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status *string `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight *int `pulumi:"weight"`
}

type HttpFarmServerState struct {
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address pulumi.StringPtrInput
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrInput
	Chain  pulumi.StringPtrInput
	// Value of the stickiness cookie used for this backend.
	Cookie pulumi.StringPtrInput
	// Label for the server
	DisplayName pulumi.StringPtrInput
	// ID of the farm this server is attached to
	FarmId pulumi.IntPtrInput
	// enable action when backend marked down. (`shutdown-sessions`)
	OnMarkedDown pulumi.StringPtrInput
	// Port that backend will respond on
	Port pulumi.IntPtrInput
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrInput
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrInput
	// backend status - `active` or `inactive`
	Status pulumi.StringPtrInput
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrInput
}

func (HttpFarmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFarmServerState)(nil)).Elem()
}

type httpFarmServerArgs struct {
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address string `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup *bool   `pulumi:"backup"`
	Chain  *string `pulumi:"chain"`
	// Label for the server
	DisplayName *string `pulumi:"displayName"`
	// ID of the farm this server is attached to
	FarmId int `pulumi:"farmId"`
	// enable action when backend marked down. (`shutdown-sessions`)
	OnMarkedDown *string `pulumi:"onMarkedDown"`
	// Port that backend will respond on
	Port *int `pulumi:"port"`
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe *bool `pulumi:"probe"`
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// is the connection ciphered with SSL (TLS)
	Ssl *bool `pulumi:"ssl"`
	// backend status - `active` or `inactive`
	Status string `pulumi:"status"`
	// used in loadbalancing algorithm
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a HttpFarmServer resource.
type HttpFarmServerArgs struct {
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address pulumi.StringInput
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrInput
	Chain  pulumi.StringPtrInput
	// Label for the server
	DisplayName pulumi.StringPtrInput
	// ID of the farm this server is attached to
	FarmId pulumi.IntInput
	// enable action when backend marked down. (`shutdown-sessions`)
	OnMarkedDown pulumi.StringPtrInput
	// Port that backend will respond on
	Port pulumi.IntPtrInput
	// defines if backend will be probed to determine health and keep as active in farm if healthy
	Probe pulumi.BoolPtrInput
	// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
	ProxyProtocolVersion pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// is the connection ciphered with SSL (TLS)
	Ssl pulumi.BoolPtrInput
	// backend status - `active` or `inactive`
	Status pulumi.StringInput
	// used in loadbalancing algorithm
	Weight pulumi.IntPtrInput
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

// Address of the backend server (IP from either internal or OVHcloud network)
func (o HttpFarmServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// is it a backup server used in case of failure of all the non-backup backends
func (o HttpFarmServerOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o HttpFarmServerOutput) Chain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.Chain }).(pulumi.StringPtrOutput)
}

// Value of the stickiness cookie used for this backend.
func (o HttpFarmServerOutput) Cookie() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Cookie }).(pulumi.StringOutput)
}

// Label for the server
func (o HttpFarmServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// ID of the farm this server is attached to
func (o HttpFarmServerOutput) FarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.IntOutput { return v.FarmId }).(pulumi.IntOutput)
}

// enable action when backend marked down. (`shutdown-sessions`)
func (o HttpFarmServerOutput) OnMarkedDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.OnMarkedDown }).(pulumi.StringPtrOutput)
}

// Port that backend will respond on
func (o HttpFarmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// defines if backend will be probed to determine health and keep as active in farm if healthy
func (o HttpFarmServerOutput) Probe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Probe }).(pulumi.BoolPtrOutput)
}

// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
func (o HttpFarmServerOutput) ProxyProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringPtrOutput { return v.ProxyProtocolVersion }).(pulumi.StringPtrOutput)
}

// The internal name of your IP load balancing
func (o HttpFarmServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// is the connection ciphered with SSL (TLS)
func (o HttpFarmServerOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// backend status - `active` or `inactive`
func (o HttpFarmServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarmServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// used in loadbalancing algorithm
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
