// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new IP whitelist on your private cloud database instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/hosting"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := hosting.NewPrivateDatabaseAllowlist(ctx, "ip", &hosting.PrivateDatabaseAllowlistArgs{
//				ServiceName: pulumi.String("XXXXXX"),
//				Ip:          pulumi.String("1.2.3.4"),
//				Name:        pulumi.String("A name for your IP address"),
//				Service:     pulumi.Bool(true),
//				Sftp:        pulumi.Bool(true),
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
// OVHcloud database whitelist can be imported using the `service_name` and the `ip`, separated by "/" E.g.,
//
// ```sh
// $ pulumi import ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist ip service_name/ip
// ```
type PrivateDatabaseAllowlist struct {
	pulumi.CustomResourceState

	// The whitelisted IP in your instance.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Custom name for your Whitelisted IP.
	Name pulumi.StringOutput `pulumi:"name"`
	// Authorize this IP to access service port. Values can be `true` or `false`
	Service pulumi.BoolOutput `pulumi:"service"`
	// The internal name of your private database.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Authorize this IP to access SFTP port. Values can be `true` or `false`
	Sftp pulumi.BoolOutput `pulumi:"sftp"`
}

// NewPrivateDatabaseAllowlist registers a new resource with the given unique name, arguments, and options.
func NewPrivateDatabaseAllowlist(ctx *pulumi.Context,
	name string, args *PrivateDatabaseAllowlistArgs, opts ...pulumi.ResourceOption) (*PrivateDatabaseAllowlist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Sftp == nil {
		return nil, errors.New("invalid value for required argument 'Sftp'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateDatabaseAllowlist
	err := ctx.RegisterResource("ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDatabaseAllowlist gets an existing PrivateDatabaseAllowlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDatabaseAllowlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDatabaseAllowlistState, opts ...pulumi.ResourceOption) (*PrivateDatabaseAllowlist, error) {
	var resource PrivateDatabaseAllowlist
	err := ctx.ReadResource("ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDatabaseAllowlist resources.
type privateDatabaseAllowlistState struct {
	// The whitelisted IP in your instance.
	Ip *string `pulumi:"ip"`
	// Custom name for your Whitelisted IP.
	Name *string `pulumi:"name"`
	// Authorize this IP to access service port. Values can be `true` or `false`
	Service *bool `pulumi:"service"`
	// The internal name of your private database.
	ServiceName *string `pulumi:"serviceName"`
	// Authorize this IP to access SFTP port. Values can be `true` or `false`
	Sftp *bool `pulumi:"sftp"`
}

type PrivateDatabaseAllowlistState struct {
	// The whitelisted IP in your instance.
	Ip pulumi.StringPtrInput
	// Custom name for your Whitelisted IP.
	Name pulumi.StringPtrInput
	// Authorize this IP to access service port. Values can be `true` or `false`
	Service pulumi.BoolPtrInput
	// The internal name of your private database.
	ServiceName pulumi.StringPtrInput
	// Authorize this IP to access SFTP port. Values can be `true` or `false`
	Sftp pulumi.BoolPtrInput
}

func (PrivateDatabaseAllowlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseAllowlistState)(nil)).Elem()
}

type privateDatabaseAllowlistArgs struct {
	// The whitelisted IP in your instance.
	Ip string `pulumi:"ip"`
	// Custom name for your Whitelisted IP.
	Name *string `pulumi:"name"`
	// Authorize this IP to access service port. Values can be `true` or `false`
	Service bool `pulumi:"service"`
	// The internal name of your private database.
	ServiceName string `pulumi:"serviceName"`
	// Authorize this IP to access SFTP port. Values can be `true` or `false`
	Sftp bool `pulumi:"sftp"`
}

// The set of arguments for constructing a PrivateDatabaseAllowlist resource.
type PrivateDatabaseAllowlistArgs struct {
	// The whitelisted IP in your instance.
	Ip pulumi.StringInput
	// Custom name for your Whitelisted IP.
	Name pulumi.StringPtrInput
	// Authorize this IP to access service port. Values can be `true` or `false`
	Service pulumi.BoolInput
	// The internal name of your private database.
	ServiceName pulumi.StringInput
	// Authorize this IP to access SFTP port. Values can be `true` or `false`
	Sftp pulumi.BoolInput
}

func (PrivateDatabaseAllowlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseAllowlistArgs)(nil)).Elem()
}

type PrivateDatabaseAllowlistInput interface {
	pulumi.Input

	ToPrivateDatabaseAllowlistOutput() PrivateDatabaseAllowlistOutput
	ToPrivateDatabaseAllowlistOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistOutput
}

func (*PrivateDatabaseAllowlist) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseAllowlist)(nil)).Elem()
}

func (i *PrivateDatabaseAllowlist) ToPrivateDatabaseAllowlistOutput() PrivateDatabaseAllowlistOutput {
	return i.ToPrivateDatabaseAllowlistOutputWithContext(context.Background())
}

func (i *PrivateDatabaseAllowlist) ToPrivateDatabaseAllowlistOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseAllowlistOutput)
}

// PrivateDatabaseAllowlistArrayInput is an input type that accepts PrivateDatabaseAllowlistArray and PrivateDatabaseAllowlistArrayOutput values.
// You can construct a concrete instance of `PrivateDatabaseAllowlistArrayInput` via:
//
//	PrivateDatabaseAllowlistArray{ PrivateDatabaseAllowlistArgs{...} }
type PrivateDatabaseAllowlistArrayInput interface {
	pulumi.Input

	ToPrivateDatabaseAllowlistArrayOutput() PrivateDatabaseAllowlistArrayOutput
	ToPrivateDatabaseAllowlistArrayOutputWithContext(context.Context) PrivateDatabaseAllowlistArrayOutput
}

type PrivateDatabaseAllowlistArray []PrivateDatabaseAllowlistInput

func (PrivateDatabaseAllowlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseAllowlist)(nil)).Elem()
}

func (i PrivateDatabaseAllowlistArray) ToPrivateDatabaseAllowlistArrayOutput() PrivateDatabaseAllowlistArrayOutput {
	return i.ToPrivateDatabaseAllowlistArrayOutputWithContext(context.Background())
}

func (i PrivateDatabaseAllowlistArray) ToPrivateDatabaseAllowlistArrayOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseAllowlistArrayOutput)
}

// PrivateDatabaseAllowlistMapInput is an input type that accepts PrivateDatabaseAllowlistMap and PrivateDatabaseAllowlistMapOutput values.
// You can construct a concrete instance of `PrivateDatabaseAllowlistMapInput` via:
//
//	PrivateDatabaseAllowlistMap{ "key": PrivateDatabaseAllowlistArgs{...} }
type PrivateDatabaseAllowlistMapInput interface {
	pulumi.Input

	ToPrivateDatabaseAllowlistMapOutput() PrivateDatabaseAllowlistMapOutput
	ToPrivateDatabaseAllowlistMapOutputWithContext(context.Context) PrivateDatabaseAllowlistMapOutput
}

type PrivateDatabaseAllowlistMap map[string]PrivateDatabaseAllowlistInput

func (PrivateDatabaseAllowlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseAllowlist)(nil)).Elem()
}

func (i PrivateDatabaseAllowlistMap) ToPrivateDatabaseAllowlistMapOutput() PrivateDatabaseAllowlistMapOutput {
	return i.ToPrivateDatabaseAllowlistMapOutputWithContext(context.Background())
}

func (i PrivateDatabaseAllowlistMap) ToPrivateDatabaseAllowlistMapOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseAllowlistMapOutput)
}

type PrivateDatabaseAllowlistOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseAllowlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseAllowlist)(nil)).Elem()
}

func (o PrivateDatabaseAllowlistOutput) ToPrivateDatabaseAllowlistOutput() PrivateDatabaseAllowlistOutput {
	return o
}

func (o PrivateDatabaseAllowlistOutput) ToPrivateDatabaseAllowlistOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistOutput {
	return o
}

// The whitelisted IP in your instance.
func (o PrivateDatabaseAllowlistOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseAllowlist) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Custom name for your Whitelisted IP.
func (o PrivateDatabaseAllowlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseAllowlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Authorize this IP to access service port. Values can be `true` or `false`
func (o PrivateDatabaseAllowlistOutput) Service() pulumi.BoolOutput {
	return o.ApplyT(func(v *PrivateDatabaseAllowlist) pulumi.BoolOutput { return v.Service }).(pulumi.BoolOutput)
}

// The internal name of your private database.
func (o PrivateDatabaseAllowlistOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseAllowlist) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Authorize this IP to access SFTP port. Values can be `true` or `false`
func (o PrivateDatabaseAllowlistOutput) Sftp() pulumi.BoolOutput {
	return o.ApplyT(func(v *PrivateDatabaseAllowlist) pulumi.BoolOutput { return v.Sftp }).(pulumi.BoolOutput)
}

type PrivateDatabaseAllowlistArrayOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseAllowlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseAllowlist)(nil)).Elem()
}

func (o PrivateDatabaseAllowlistArrayOutput) ToPrivateDatabaseAllowlistArrayOutput() PrivateDatabaseAllowlistArrayOutput {
	return o
}

func (o PrivateDatabaseAllowlistArrayOutput) ToPrivateDatabaseAllowlistArrayOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistArrayOutput {
	return o
}

func (o PrivateDatabaseAllowlistArrayOutput) Index(i pulumi.IntInput) PrivateDatabaseAllowlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateDatabaseAllowlist {
		return vs[0].([]*PrivateDatabaseAllowlist)[vs[1].(int)]
	}).(PrivateDatabaseAllowlistOutput)
}

type PrivateDatabaseAllowlistMapOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseAllowlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseAllowlist)(nil)).Elem()
}

func (o PrivateDatabaseAllowlistMapOutput) ToPrivateDatabaseAllowlistMapOutput() PrivateDatabaseAllowlistMapOutput {
	return o
}

func (o PrivateDatabaseAllowlistMapOutput) ToPrivateDatabaseAllowlistMapOutputWithContext(ctx context.Context) PrivateDatabaseAllowlistMapOutput {
	return o
}

func (o PrivateDatabaseAllowlistMapOutput) MapIndex(k pulumi.StringInput) PrivateDatabaseAllowlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateDatabaseAllowlist {
		return vs[0].(map[string]*PrivateDatabaseAllowlist)[vs[1].(string)]
	}).(PrivateDatabaseAllowlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseAllowlistInput)(nil)).Elem(), &PrivateDatabaseAllowlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseAllowlistArrayInput)(nil)).Elem(), PrivateDatabaseAllowlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseAllowlistMapInput)(nil)).Elem(), PrivateDatabaseAllowlistMap{})
	pulumi.RegisterOutputType(PrivateDatabaseAllowlistOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseAllowlistArrayOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseAllowlistMapOutput{})
}
