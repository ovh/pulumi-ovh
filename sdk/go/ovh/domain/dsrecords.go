// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to manage a domain's DS records.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/domain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := domain.NewDSRecords(ctx, "dsRecords", &domain.DSRecordsArgs{
//				Domain: pulumi.String("mydomain.ovh"),
//				DsRecords: domain.DSRecordsDsRecordArray{
//					&domain.DSRecordsDsRecordArgs{
//						Algorithm: pulumi.String("RSASHA1_NSEC3_SHA1"),
//						Flags:     pulumi.String("KEY_SIGNING_KEY"),
//						PublicKey: pulumi.String("my_base64_encoded_public_key"),
//						Tag:       pulumi.Int(12345),
//					},
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
// DS records can be imported using their `domain`.
//
// Using the following configuration:
//
// hcl
//
// import {
//
//	to = ovh_domain_ds_records.ds_records
//
//	id = "<domain name>"
//
// }
//
// You can then run:
//
// bash
//
// $ pulumi preview -generate-config-out=ds_records.tf
//
// $ pulumi up
//
// The file `ds_records.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.
//
// See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type DSRecords struct {
	pulumi.CustomResourceState

	// Domain name for which to manage DS records
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Details about a DS record
	DsRecords DSRecordsDsRecordArrayOutput `pulumi:"dsRecords"`
}

// NewDSRecords registers a new resource with the given unique name, arguments, and options.
func NewDSRecords(ctx *pulumi.Context,
	name string, args *DSRecordsArgs, opts ...pulumi.ResourceOption) (*DSRecords, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.DsRecords == nil {
		return nil, errors.New("invalid value for required argument 'DsRecords'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DSRecords
	err := ctx.RegisterResource("ovh:Domain/dSRecords:DSRecords", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDSRecords gets an existing DSRecords resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDSRecords(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DSRecordsState, opts ...pulumi.ResourceOption) (*DSRecords, error) {
	var resource DSRecords
	err := ctx.ReadResource("ovh:Domain/dSRecords:DSRecords", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DSRecords resources.
type dsrecordsState struct {
	// Domain name for which to manage DS records
	Domain *string `pulumi:"domain"`
	// Details about a DS record
	DsRecords []DSRecordsDsRecord `pulumi:"dsRecords"`
}

type DSRecordsState struct {
	// Domain name for which to manage DS records
	Domain pulumi.StringPtrInput
	// Details about a DS record
	DsRecords DSRecordsDsRecordArrayInput
}

func (DSRecordsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dsrecordsState)(nil)).Elem()
}

type dsrecordsArgs struct {
	// Domain name for which to manage DS records
	Domain string `pulumi:"domain"`
	// Details about a DS record
	DsRecords []DSRecordsDsRecord `pulumi:"dsRecords"`
}

// The set of arguments for constructing a DSRecords resource.
type DSRecordsArgs struct {
	// Domain name for which to manage DS records
	Domain pulumi.StringInput
	// Details about a DS record
	DsRecords DSRecordsDsRecordArrayInput
}

func (DSRecordsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dsrecordsArgs)(nil)).Elem()
}

type DSRecordsInput interface {
	pulumi.Input

	ToDSRecordsOutput() DSRecordsOutput
	ToDSRecordsOutputWithContext(ctx context.Context) DSRecordsOutput
}

func (*DSRecords) ElementType() reflect.Type {
	return reflect.TypeOf((**DSRecords)(nil)).Elem()
}

func (i *DSRecords) ToDSRecordsOutput() DSRecordsOutput {
	return i.ToDSRecordsOutputWithContext(context.Background())
}

func (i *DSRecords) ToDSRecordsOutputWithContext(ctx context.Context) DSRecordsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DSRecordsOutput)
}

// DSRecordsArrayInput is an input type that accepts DSRecordsArray and DSRecordsArrayOutput values.
// You can construct a concrete instance of `DSRecordsArrayInput` via:
//
//	DSRecordsArray{ DSRecordsArgs{...} }
type DSRecordsArrayInput interface {
	pulumi.Input

	ToDSRecordsArrayOutput() DSRecordsArrayOutput
	ToDSRecordsArrayOutputWithContext(context.Context) DSRecordsArrayOutput
}

type DSRecordsArray []DSRecordsInput

func (DSRecordsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DSRecords)(nil)).Elem()
}

func (i DSRecordsArray) ToDSRecordsArrayOutput() DSRecordsArrayOutput {
	return i.ToDSRecordsArrayOutputWithContext(context.Background())
}

func (i DSRecordsArray) ToDSRecordsArrayOutputWithContext(ctx context.Context) DSRecordsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DSRecordsArrayOutput)
}

// DSRecordsMapInput is an input type that accepts DSRecordsMap and DSRecordsMapOutput values.
// You can construct a concrete instance of `DSRecordsMapInput` via:
//
//	DSRecordsMap{ "key": DSRecordsArgs{...} }
type DSRecordsMapInput interface {
	pulumi.Input

	ToDSRecordsMapOutput() DSRecordsMapOutput
	ToDSRecordsMapOutputWithContext(context.Context) DSRecordsMapOutput
}

type DSRecordsMap map[string]DSRecordsInput

func (DSRecordsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DSRecords)(nil)).Elem()
}

func (i DSRecordsMap) ToDSRecordsMapOutput() DSRecordsMapOutput {
	return i.ToDSRecordsMapOutputWithContext(context.Background())
}

func (i DSRecordsMap) ToDSRecordsMapOutputWithContext(ctx context.Context) DSRecordsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DSRecordsMapOutput)
}

type DSRecordsOutput struct{ *pulumi.OutputState }

func (DSRecordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DSRecords)(nil)).Elem()
}

func (o DSRecordsOutput) ToDSRecordsOutput() DSRecordsOutput {
	return o
}

func (o DSRecordsOutput) ToDSRecordsOutputWithContext(ctx context.Context) DSRecordsOutput {
	return o
}

// Domain name for which to manage DS records
func (o DSRecordsOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DSRecords) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Details about a DS record
func (o DSRecordsOutput) DsRecords() DSRecordsDsRecordArrayOutput {
	return o.ApplyT(func(v *DSRecords) DSRecordsDsRecordArrayOutput { return v.DsRecords }).(DSRecordsDsRecordArrayOutput)
}

type DSRecordsArrayOutput struct{ *pulumi.OutputState }

func (DSRecordsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DSRecords)(nil)).Elem()
}

func (o DSRecordsArrayOutput) ToDSRecordsArrayOutput() DSRecordsArrayOutput {
	return o
}

func (o DSRecordsArrayOutput) ToDSRecordsArrayOutputWithContext(ctx context.Context) DSRecordsArrayOutput {
	return o
}

func (o DSRecordsArrayOutput) Index(i pulumi.IntInput) DSRecordsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DSRecords {
		return vs[0].([]*DSRecords)[vs[1].(int)]
	}).(DSRecordsOutput)
}

type DSRecordsMapOutput struct{ *pulumi.OutputState }

func (DSRecordsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DSRecords)(nil)).Elem()
}

func (o DSRecordsMapOutput) ToDSRecordsMapOutput() DSRecordsMapOutput {
	return o
}

func (o DSRecordsMapOutput) ToDSRecordsMapOutputWithContext(ctx context.Context) DSRecordsMapOutput {
	return o
}

func (o DSRecordsMapOutput) MapIndex(k pulumi.StringInput) DSRecordsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DSRecords {
		return vs[0].(map[string]*DSRecords)[vs[1].(string)]
	}).(DSRecordsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DSRecordsInput)(nil)).Elem(), &DSRecords{})
	pulumi.RegisterInputType(reflect.TypeOf((*DSRecordsArrayInput)(nil)).Elem(), DSRecordsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DSRecordsMapInput)(nil)).Elem(), DSRecordsMap{})
	pulumi.RegisterOutputType(DSRecordsOutput{})
	pulumi.RegisterOutputType(DSRecordsArrayOutput{})
	pulumi.RegisterOutputType(DSRecordsMapOutput{})
}
