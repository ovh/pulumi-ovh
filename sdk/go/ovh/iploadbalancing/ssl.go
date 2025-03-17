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

// Creates a new custom SSL certificate on your IP Load Balancing
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
//			_, err = iploadbalancing.NewSsl(ctx, "sslname", &iploadbalancing.SslArgs{
//				Certificate: pulumi.String("..."),
//				Chain:       pulumi.String("..."),
//				DisplayName: pulumi.String("test"),
//				Key:         pulumi.String("..."),
//				ServiceName: pulumi.String(lb.ServiceName),
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
// SSL can be imported using the following format `service_name` and the `id` of the ssl, separated by "/" e.g.
//
// bash
//
// ```sh
// $ pulumi import ovh:IpLoadBalancing/ssl:Ssl sslname service_name/ssl_id
// ```
type Ssl struct {
	pulumi.CustomResourceState

	// Certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Certificate chain
	Chain pulumi.StringPtrOutput `pulumi:"chain"`
	// Readable label for loadbalancer ssl
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Expire date of your SSL certificate.
	ExpireDate pulumi.StringOutput `pulumi:"expireDate"`
	// Fingerprint of your SSL certificate.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Certificate key
	Key pulumi.StringOutput `pulumi:"key"`
	// Subject Alternative Name of your SSL certificate.
	Sans pulumi.StringArrayOutput `pulumi:"sans"`
	// Serial of your SSL certificate (Deprecated, use fingerprint instead !)
	Serial pulumi.StringOutput `pulumi:"serial"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Subject of your SSL certificate.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Type of your SSL certificate. 'built' for SSL certificates managed by the IP Load Balancing. 'custom' for user manager certificates.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSsl registers a new resource with the given unique name, arguments, and options.
func NewSsl(ctx *pulumi.Context,
	name string, args *SslArgs, opts ...pulumi.ResourceOption) (*Ssl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Key != nil {
		args.Key = pulumi.ToSecret(args.Key).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"key",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ssl
	err := ctx.RegisterResource("ovh:IpLoadBalancing/ssl:Ssl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSsl gets an existing Ssl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslState, opts ...pulumi.ResourceOption) (*Ssl, error) {
	var resource Ssl
	err := ctx.ReadResource("ovh:IpLoadBalancing/ssl:Ssl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ssl resources.
type sslState struct {
	// Certificate
	Certificate *string `pulumi:"certificate"`
	// Certificate chain
	Chain *string `pulumi:"chain"`
	// Readable label for loadbalancer ssl
	DisplayName *string `pulumi:"displayName"`
	// Expire date of your SSL certificate.
	ExpireDate *string `pulumi:"expireDate"`
	// Fingerprint of your SSL certificate.
	Fingerprint *string `pulumi:"fingerprint"`
	// Certificate key
	Key *string `pulumi:"key"`
	// Subject Alternative Name of your SSL certificate.
	Sans []string `pulumi:"sans"`
	// Serial of your SSL certificate (Deprecated, use fingerprint instead !)
	Serial *string `pulumi:"serial"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Subject of your SSL certificate.
	Subject *string `pulumi:"subject"`
	// Type of your SSL certificate. 'built' for SSL certificates managed by the IP Load Balancing. 'custom' for user manager certificates.
	Type *string `pulumi:"type"`
}

type SslState struct {
	// Certificate
	Certificate pulumi.StringPtrInput
	// Certificate chain
	Chain pulumi.StringPtrInput
	// Readable label for loadbalancer ssl
	DisplayName pulumi.StringPtrInput
	// Expire date of your SSL certificate.
	ExpireDate pulumi.StringPtrInput
	// Fingerprint of your SSL certificate.
	Fingerprint pulumi.StringPtrInput
	// Certificate key
	Key pulumi.StringPtrInput
	// Subject Alternative Name of your SSL certificate.
	Sans pulumi.StringArrayInput
	// Serial of your SSL certificate (Deprecated, use fingerprint instead !)
	Serial pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Subject of your SSL certificate.
	Subject pulumi.StringPtrInput
	// Type of your SSL certificate. 'built' for SSL certificates managed by the IP Load Balancing. 'custom' for user manager certificates.
	Type pulumi.StringPtrInput
}

func (SslState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslState)(nil)).Elem()
}

type sslArgs struct {
	// Certificate
	Certificate string `pulumi:"certificate"`
	// Certificate chain
	Chain *string `pulumi:"chain"`
	// Readable label for loadbalancer ssl
	DisplayName *string `pulumi:"displayName"`
	// Certificate key
	Key string `pulumi:"key"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ssl resource.
type SslArgs struct {
	// Certificate
	Certificate pulumi.StringInput
	// Certificate chain
	Chain pulumi.StringPtrInput
	// Readable label for loadbalancer ssl
	DisplayName pulumi.StringPtrInput
	// Certificate key
	Key pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
}

func (SslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslArgs)(nil)).Elem()
}

type SslInput interface {
	pulumi.Input

	ToSslOutput() SslOutput
	ToSslOutputWithContext(ctx context.Context) SslOutput
}

func (*Ssl) ElementType() reflect.Type {
	return reflect.TypeOf((**Ssl)(nil)).Elem()
}

func (i *Ssl) ToSslOutput() SslOutput {
	return i.ToSslOutputWithContext(context.Background())
}

func (i *Ssl) ToSslOutputWithContext(ctx context.Context) SslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslOutput)
}

// SslArrayInput is an input type that accepts SslArray and SslArrayOutput values.
// You can construct a concrete instance of `SslArrayInput` via:
//
//	SslArray{ SslArgs{...} }
type SslArrayInput interface {
	pulumi.Input

	ToSslArrayOutput() SslArrayOutput
	ToSslArrayOutputWithContext(context.Context) SslArrayOutput
}

type SslArray []SslInput

func (SslArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ssl)(nil)).Elem()
}

func (i SslArray) ToSslArrayOutput() SslArrayOutput {
	return i.ToSslArrayOutputWithContext(context.Background())
}

func (i SslArray) ToSslArrayOutputWithContext(ctx context.Context) SslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslArrayOutput)
}

// SslMapInput is an input type that accepts SslMap and SslMapOutput values.
// You can construct a concrete instance of `SslMapInput` via:
//
//	SslMap{ "key": SslArgs{...} }
type SslMapInput interface {
	pulumi.Input

	ToSslMapOutput() SslMapOutput
	ToSslMapOutputWithContext(context.Context) SslMapOutput
}

type SslMap map[string]SslInput

func (SslMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ssl)(nil)).Elem()
}

func (i SslMap) ToSslMapOutput() SslMapOutput {
	return i.ToSslMapOutputWithContext(context.Background())
}

func (i SslMap) ToSslMapOutputWithContext(ctx context.Context) SslMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslMapOutput)
}

type SslOutput struct{ *pulumi.OutputState }

func (SslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ssl)(nil)).Elem()
}

func (o SslOutput) ToSslOutput() SslOutput {
	return o
}

func (o SslOutput) ToSslOutputWithContext(ctx context.Context) SslOutput {
	return o
}

// Certificate
func (o SslOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Certificate chain
func (o SslOutput) Chain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringPtrOutput { return v.Chain }).(pulumi.StringPtrOutput)
}

// Readable label for loadbalancer ssl
func (o SslOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Expire date of your SSL certificate.
func (o SslOutput) ExpireDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.ExpireDate }).(pulumi.StringOutput)
}

// Fingerprint of your SSL certificate.
func (o SslOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Certificate key
func (o SslOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Subject Alternative Name of your SSL certificate.
func (o SslOutput) Sans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringArrayOutput { return v.Sans }).(pulumi.StringArrayOutput)
}

// Serial of your SSL certificate (Deprecated, use fingerprint instead !)
func (o SslOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Serial }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o SslOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Subject of your SSL certificate.
func (o SslOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

// Type of your SSL certificate. 'built' for SSL certificates managed by the IP Load Balancing. 'custom' for user manager certificates.
func (o SslOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Ssl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type SslArrayOutput struct{ *pulumi.OutputState }

func (SslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ssl)(nil)).Elem()
}

func (o SslArrayOutput) ToSslArrayOutput() SslArrayOutput {
	return o
}

func (o SslArrayOutput) ToSslArrayOutputWithContext(ctx context.Context) SslArrayOutput {
	return o
}

func (o SslArrayOutput) Index(i pulumi.IntInput) SslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ssl {
		return vs[0].([]*Ssl)[vs[1].(int)]
	}).(SslOutput)
}

type SslMapOutput struct{ *pulumi.OutputState }

func (SslMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ssl)(nil)).Elem()
}

func (o SslMapOutput) ToSslMapOutput() SslMapOutput {
	return o
}

func (o SslMapOutput) ToSslMapOutputWithContext(ctx context.Context) SslMapOutput {
	return o
}

func (o SslMapOutput) MapIndex(k pulumi.StringInput) SslOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ssl {
		return vs[0].(map[string]*Ssl)[vs[1].(string)]
	}).(SslOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslInput)(nil)).Elem(), &Ssl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslArrayInput)(nil)).Elem(), SslArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslMapInput)(nil)).Elem(), SslMap{})
	pulumi.RegisterOutputType(SslOutput{})
	pulumi.RegisterOutputType(SslArrayOutput{})
	pulumi.RegisterOutputType(SslMapOutput{})
}
