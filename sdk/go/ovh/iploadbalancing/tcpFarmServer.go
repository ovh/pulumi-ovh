// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server entry linked to loadbalancing group (farm)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/iploadbalancing"
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
//			farmName, err := iploadbalancing.NewTcpFarm(ctx, "farmName", &iploadbalancing.TcpFarmArgs{
//				Port:        pulumi.Int(8080),
//				ServiceName: pulumi.String(lb.ServiceName),
//				Zone:        pulumi.String("all"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iploadbalancing.NewTcpFarmServer(ctx, "backend", &iploadbalancing.TcpFarmServerArgs{
//				Address:              pulumi.String("4.5.6.7"),
//				Backup:               pulumi.Bool(true),
//				DisplayName:          pulumi.String("mybackend"),
//				FarmId:               farmName.ID(),
//				Port:                 pulumi.Int(80),
//				Probe:                pulumi.Bool(true),
//				ProxyProtocolVersion: pulumi.String("v2"),
//				ServiceName:          pulumi.String(lb.ServiceName),
//				Ssl:                  pulumi.Bool(false),
//				Status:               pulumi.String("active"),
//				Weight:               pulumi.Int(2),
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
// TCP farm server can be imported using the following format `serviceName`, the `id` of the farm and the `id` of the server separated by "/" e.g.
type TcpFarmServer struct {
	pulumi.CustomResourceState

	// Address of the backend server (IP from either internal or OVHcloud network)
	Address pulumi.StringOutput `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrOutput   `pulumi:"backup"`
	Chain  pulumi.StringPtrOutput `pulumi:"chain"`
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

// NewTcpFarmServer registers a new resource with the given unique name, arguments, and options.
func NewTcpFarmServer(ctx *pulumi.Context,
	name string, args *TcpFarmServerArgs, opts ...pulumi.ResourceOption) (*TcpFarmServer, error) {
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
	var resource TcpFarmServer
	err := ctx.RegisterResource("ovh:IpLoadBalancing/tcpFarmServer:TcpFarmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTcpFarmServer gets an existing TcpFarmServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTcpFarmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TcpFarmServerState, opts ...pulumi.ResourceOption) (*TcpFarmServer, error) {
	var resource TcpFarmServer
	err := ctx.ReadResource("ovh:IpLoadBalancing/tcpFarmServer:TcpFarmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TcpFarmServer resources.
type tcpFarmServerState struct {
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address *string `pulumi:"address"`
	// is it a backup server used in case of failure of all the non-backup backends
	Backup *bool   `pulumi:"backup"`
	Chain  *string `pulumi:"chain"`
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

type TcpFarmServerState struct {
	// Address of the backend server (IP from either internal or OVHcloud network)
	Address pulumi.StringPtrInput
	// is it a backup server used in case of failure of all the non-backup backends
	Backup pulumi.BoolPtrInput
	Chain  pulumi.StringPtrInput
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

func (TcpFarmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpFarmServerState)(nil)).Elem()
}

type tcpFarmServerArgs struct {
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

// The set of arguments for constructing a TcpFarmServer resource.
type TcpFarmServerArgs struct {
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

func (TcpFarmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpFarmServerArgs)(nil)).Elem()
}

type TcpFarmServerInput interface {
	pulumi.Input

	ToTcpFarmServerOutput() TcpFarmServerOutput
	ToTcpFarmServerOutputWithContext(ctx context.Context) TcpFarmServerOutput
}

func (*TcpFarmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpFarmServer)(nil)).Elem()
}

func (i *TcpFarmServer) ToTcpFarmServerOutput() TcpFarmServerOutput {
	return i.ToTcpFarmServerOutputWithContext(context.Background())
}

func (i *TcpFarmServer) ToTcpFarmServerOutputWithContext(ctx context.Context) TcpFarmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpFarmServerOutput)
}

// TcpFarmServerArrayInput is an input type that accepts TcpFarmServerArray and TcpFarmServerArrayOutput values.
// You can construct a concrete instance of `TcpFarmServerArrayInput` via:
//
//	TcpFarmServerArray{ TcpFarmServerArgs{...} }
type TcpFarmServerArrayInput interface {
	pulumi.Input

	ToTcpFarmServerArrayOutput() TcpFarmServerArrayOutput
	ToTcpFarmServerArrayOutputWithContext(context.Context) TcpFarmServerArrayOutput
}

type TcpFarmServerArray []TcpFarmServerInput

func (TcpFarmServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpFarmServer)(nil)).Elem()
}

func (i TcpFarmServerArray) ToTcpFarmServerArrayOutput() TcpFarmServerArrayOutput {
	return i.ToTcpFarmServerArrayOutputWithContext(context.Background())
}

func (i TcpFarmServerArray) ToTcpFarmServerArrayOutputWithContext(ctx context.Context) TcpFarmServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpFarmServerArrayOutput)
}

// TcpFarmServerMapInput is an input type that accepts TcpFarmServerMap and TcpFarmServerMapOutput values.
// You can construct a concrete instance of `TcpFarmServerMapInput` via:
//
//	TcpFarmServerMap{ "key": TcpFarmServerArgs{...} }
type TcpFarmServerMapInput interface {
	pulumi.Input

	ToTcpFarmServerMapOutput() TcpFarmServerMapOutput
	ToTcpFarmServerMapOutputWithContext(context.Context) TcpFarmServerMapOutput
}

type TcpFarmServerMap map[string]TcpFarmServerInput

func (TcpFarmServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpFarmServer)(nil)).Elem()
}

func (i TcpFarmServerMap) ToTcpFarmServerMapOutput() TcpFarmServerMapOutput {
	return i.ToTcpFarmServerMapOutputWithContext(context.Background())
}

func (i TcpFarmServerMap) ToTcpFarmServerMapOutputWithContext(ctx context.Context) TcpFarmServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpFarmServerMapOutput)
}

type TcpFarmServerOutput struct{ *pulumi.OutputState }

func (TcpFarmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpFarmServer)(nil)).Elem()
}

func (o TcpFarmServerOutput) ToTcpFarmServerOutput() TcpFarmServerOutput {
	return o
}

func (o TcpFarmServerOutput) ToTcpFarmServerOutputWithContext(ctx context.Context) TcpFarmServerOutput {
	return o
}

// Address of the backend server (IP from either internal or OVHcloud network)
func (o TcpFarmServerOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// is it a backup server used in case of failure of all the non-backup backends
func (o TcpFarmServerOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.BoolPtrOutput { return v.Backup }).(pulumi.BoolPtrOutput)
}

func (o TcpFarmServerOutput) Chain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringPtrOutput { return v.Chain }).(pulumi.StringPtrOutput)
}

// Label for the server
func (o TcpFarmServerOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// ID of the farm this server is attached to
func (o TcpFarmServerOutput) FarmId() pulumi.IntOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.IntOutput { return v.FarmId }).(pulumi.IntOutput)
}

// enable action when backend marked down. (`shutdown-sessions`)
func (o TcpFarmServerOutput) OnMarkedDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringPtrOutput { return v.OnMarkedDown }).(pulumi.StringPtrOutput)
}

// Port that backend will respond on
func (o TcpFarmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// defines if backend will be probed to determine health and keep as active in farm if healthy
func (o TcpFarmServerOutput) Probe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.BoolPtrOutput { return v.Probe }).(pulumi.BoolPtrOutput)
}

// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
func (o TcpFarmServerOutput) ProxyProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringPtrOutput { return v.ProxyProtocolVersion }).(pulumi.StringPtrOutput)
}

// The internal name of your IP load balancing
func (o TcpFarmServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// is the connection ciphered with SSL (TLS)
func (o TcpFarmServerOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// backend status - `active` or `inactive`
func (o TcpFarmServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// used in loadbalancing algorithm
func (o TcpFarmServerOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpFarmServer) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

type TcpFarmServerArrayOutput struct{ *pulumi.OutputState }

func (TcpFarmServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpFarmServer)(nil)).Elem()
}

func (o TcpFarmServerArrayOutput) ToTcpFarmServerArrayOutput() TcpFarmServerArrayOutput {
	return o
}

func (o TcpFarmServerArrayOutput) ToTcpFarmServerArrayOutputWithContext(ctx context.Context) TcpFarmServerArrayOutput {
	return o
}

func (o TcpFarmServerArrayOutput) Index(i pulumi.IntInput) TcpFarmServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TcpFarmServer {
		return vs[0].([]*TcpFarmServer)[vs[1].(int)]
	}).(TcpFarmServerOutput)
}

type TcpFarmServerMapOutput struct{ *pulumi.OutputState }

func (TcpFarmServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpFarmServer)(nil)).Elem()
}

func (o TcpFarmServerMapOutput) ToTcpFarmServerMapOutput() TcpFarmServerMapOutput {
	return o
}

func (o TcpFarmServerMapOutput) ToTcpFarmServerMapOutputWithContext(ctx context.Context) TcpFarmServerMapOutput {
	return o
}

func (o TcpFarmServerMapOutput) MapIndex(k pulumi.StringInput) TcpFarmServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TcpFarmServer {
		return vs[0].(map[string]*TcpFarmServer)[vs[1].(string)]
	}).(TcpFarmServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TcpFarmServerInput)(nil)).Elem(), &TcpFarmServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpFarmServerArrayInput)(nil)).Elem(), TcpFarmServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpFarmServerMapInput)(nil)).Elem(), TcpFarmServerMap{})
	pulumi.RegisterOutputType(TcpFarmServerOutput{})
	pulumi.RegisterOutputType(TcpFarmServerArrayOutput{})
	pulumi.RegisterOutputType(TcpFarmServerMapOutput{})
}
