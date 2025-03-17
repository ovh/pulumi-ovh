// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enable / disable DNSSEC on a domain zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/domain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := domain.NewZoneDNSSec(ctx, "dnssec", &domain.ZoneDNSSecArgs{
//				ZoneName: pulumi.String("mysite.ovh"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ZoneDNSSec struct {
	pulumi.CustomResourceState

	// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of the domain zone
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewZoneDNSSec registers a new resource with the given unique name, arguments, and options.
func NewZoneDNSSec(ctx *pulumi.Context,
	name string, args *ZoneDNSSecArgs, opts ...pulumi.ResourceOption) (*ZoneDNSSec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ZoneDNSSec
	err := ctx.RegisterResource("ovh:Domain/zoneDNSSec:ZoneDNSSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneDNSSec gets an existing ZoneDNSSec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneDNSSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneDNSSecState, opts ...pulumi.ResourceOption) (*ZoneDNSSec, error) {
	var resource ZoneDNSSec
	err := ctx.ReadResource("ovh:Domain/zoneDNSSec:ZoneDNSSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneDNSSec resources.
type zoneDNSSecState struct {
	// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
	Status *string `pulumi:"status"`
	// The name of the domain zone
	ZoneName *string `pulumi:"zoneName"`
}

type ZoneDNSSecState struct {
	// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
	Status pulumi.StringPtrInput
	// The name of the domain zone
	ZoneName pulumi.StringPtrInput
}

func (ZoneDNSSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneDNSSecState)(nil)).Elem()
}

type zoneDNSSecArgs struct {
	// The name of the domain zone
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a ZoneDNSSec resource.
type ZoneDNSSecArgs struct {
	// The name of the domain zone
	ZoneName pulumi.StringInput
}

func (ZoneDNSSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneDNSSecArgs)(nil)).Elem()
}

type ZoneDNSSecInput interface {
	pulumi.Input

	ToZoneDNSSecOutput() ZoneDNSSecOutput
	ToZoneDNSSecOutputWithContext(ctx context.Context) ZoneDNSSecOutput
}

func (*ZoneDNSSec) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneDNSSec)(nil)).Elem()
}

func (i *ZoneDNSSec) ToZoneDNSSecOutput() ZoneDNSSecOutput {
	return i.ToZoneDNSSecOutputWithContext(context.Background())
}

func (i *ZoneDNSSec) ToZoneDNSSecOutputWithContext(ctx context.Context) ZoneDNSSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneDNSSecOutput)
}

// ZoneDNSSecArrayInput is an input type that accepts ZoneDNSSecArray and ZoneDNSSecArrayOutput values.
// You can construct a concrete instance of `ZoneDNSSecArrayInput` via:
//
//	ZoneDNSSecArray{ ZoneDNSSecArgs{...} }
type ZoneDNSSecArrayInput interface {
	pulumi.Input

	ToZoneDNSSecArrayOutput() ZoneDNSSecArrayOutput
	ToZoneDNSSecArrayOutputWithContext(context.Context) ZoneDNSSecArrayOutput
}

type ZoneDNSSecArray []ZoneDNSSecInput

func (ZoneDNSSecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneDNSSec)(nil)).Elem()
}

func (i ZoneDNSSecArray) ToZoneDNSSecArrayOutput() ZoneDNSSecArrayOutput {
	return i.ToZoneDNSSecArrayOutputWithContext(context.Background())
}

func (i ZoneDNSSecArray) ToZoneDNSSecArrayOutputWithContext(ctx context.Context) ZoneDNSSecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneDNSSecArrayOutput)
}

// ZoneDNSSecMapInput is an input type that accepts ZoneDNSSecMap and ZoneDNSSecMapOutput values.
// You can construct a concrete instance of `ZoneDNSSecMapInput` via:
//
//	ZoneDNSSecMap{ "key": ZoneDNSSecArgs{...} }
type ZoneDNSSecMapInput interface {
	pulumi.Input

	ToZoneDNSSecMapOutput() ZoneDNSSecMapOutput
	ToZoneDNSSecMapOutputWithContext(context.Context) ZoneDNSSecMapOutput
}

type ZoneDNSSecMap map[string]ZoneDNSSecInput

func (ZoneDNSSecMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneDNSSec)(nil)).Elem()
}

func (i ZoneDNSSecMap) ToZoneDNSSecMapOutput() ZoneDNSSecMapOutput {
	return i.ToZoneDNSSecMapOutputWithContext(context.Background())
}

func (i ZoneDNSSecMap) ToZoneDNSSecMapOutputWithContext(ctx context.Context) ZoneDNSSecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneDNSSecMapOutput)
}

type ZoneDNSSecOutput struct{ *pulumi.OutputState }

func (ZoneDNSSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneDNSSec)(nil)).Elem()
}

func (o ZoneDNSSecOutput) ToZoneDNSSecOutput() ZoneDNSSecOutput {
	return o
}

func (o ZoneDNSSecOutput) ToZoneDNSSecOutputWithContext(ctx context.Context) ZoneDNSSecOutput {
	return o
}

// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
func (o ZoneDNSSecOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneDNSSec) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The name of the domain zone
func (o ZoneDNSSecOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneDNSSec) pulumi.StringOutput { return v.ZoneName }).(pulumi.StringOutput)
}

type ZoneDNSSecArrayOutput struct{ *pulumi.OutputState }

func (ZoneDNSSecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneDNSSec)(nil)).Elem()
}

func (o ZoneDNSSecArrayOutput) ToZoneDNSSecArrayOutput() ZoneDNSSecArrayOutput {
	return o
}

func (o ZoneDNSSecArrayOutput) ToZoneDNSSecArrayOutputWithContext(ctx context.Context) ZoneDNSSecArrayOutput {
	return o
}

func (o ZoneDNSSecArrayOutput) Index(i pulumi.IntInput) ZoneDNSSecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneDNSSec {
		return vs[0].([]*ZoneDNSSec)[vs[1].(int)]
	}).(ZoneDNSSecOutput)
}

type ZoneDNSSecMapOutput struct{ *pulumi.OutputState }

func (ZoneDNSSecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneDNSSec)(nil)).Elem()
}

func (o ZoneDNSSecMapOutput) ToZoneDNSSecMapOutput() ZoneDNSSecMapOutput {
	return o
}

func (o ZoneDNSSecMapOutput) ToZoneDNSSecMapOutputWithContext(ctx context.Context) ZoneDNSSecMapOutput {
	return o
}

func (o ZoneDNSSecMapOutput) MapIndex(k pulumi.StringInput) ZoneDNSSecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneDNSSec {
		return vs[0].(map[string]*ZoneDNSSec)[vs[1].(string)]
	}).(ZoneDNSSecOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneDNSSecInput)(nil)).Elem(), &ZoneDNSSec{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneDNSSecArrayInput)(nil)).Elem(), ZoneDNSSecArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneDNSSecMapInput)(nil)).Elem(), ZoneDNSSecMap{})
	pulumi.RegisterOutputType(ZoneDNSSecOutput{})
	pulumi.RegisterOutputType(ZoneDNSSecArrayOutput{})
	pulumi.RegisterOutputType(ZoneDNSSecMapOutput{})
}
