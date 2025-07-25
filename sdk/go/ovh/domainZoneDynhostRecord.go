// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a dynhost record for a given domain zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.NewDomainZoneDynhostRecord(ctx, "dynhost_record", &ovh.DomainZoneDynhostRecordArgs{
//				ZoneName:  pulumi.String("mydomain.ovh"),
//				SubDomain: pulumi.String("dynhost"),
//				Ip:        pulumi.String("1.2.3.4"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DomainZoneDynhostRecord struct {
	pulumi.CustomResourceState

	// Record IP
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Record sub-domain
	SubDomain pulumi.StringOutput `pulumi:"subDomain"`
	// Record TTL (Time to Live)
	Ttl pulumi.Float64Output `pulumi:"ttl"`
	// Record zone
	Zone pulumi.StringOutput `pulumi:"zone"`
	// Zone name
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewDomainZoneDynhostRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainZoneDynhostRecord(ctx *pulumi.Context,
	name string, args *DomainZoneDynhostRecordArgs, opts ...pulumi.ResourceOption) (*DomainZoneDynhostRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainZoneDynhostRecord
	err := ctx.RegisterResource("ovh:index/domainZoneDynhostRecord:DomainZoneDynhostRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainZoneDynhostRecord gets an existing DomainZoneDynhostRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainZoneDynhostRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainZoneDynhostRecordState, opts ...pulumi.ResourceOption) (*DomainZoneDynhostRecord, error) {
	var resource DomainZoneDynhostRecord
	err := ctx.ReadResource("ovh:index/domainZoneDynhostRecord:DomainZoneDynhostRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainZoneDynhostRecord resources.
type domainZoneDynhostRecordState struct {
	// Record IP
	Ip *string `pulumi:"ip"`
	// Record sub-domain
	SubDomain *string `pulumi:"subDomain"`
	// Record TTL (Time to Live)
	Ttl *float64 `pulumi:"ttl"`
	// Record zone
	Zone *string `pulumi:"zone"`
	// Zone name
	ZoneName *string `pulumi:"zoneName"`
}

type DomainZoneDynhostRecordState struct {
	// Record IP
	Ip pulumi.StringPtrInput
	// Record sub-domain
	SubDomain pulumi.StringPtrInput
	// Record TTL (Time to Live)
	Ttl pulumi.Float64PtrInput
	// Record zone
	Zone pulumi.StringPtrInput
	// Zone name
	ZoneName pulumi.StringPtrInput
}

func (DomainZoneDynhostRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneDynhostRecordState)(nil)).Elem()
}

type domainZoneDynhostRecordArgs struct {
	// Record IP
	Ip *string `pulumi:"ip"`
	// Record sub-domain
	SubDomain *string `pulumi:"subDomain"`
	// Zone name
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a DomainZoneDynhostRecord resource.
type DomainZoneDynhostRecordArgs struct {
	// Record IP
	Ip pulumi.StringPtrInput
	// Record sub-domain
	SubDomain pulumi.StringPtrInput
	// Zone name
	ZoneName pulumi.StringInput
}

func (DomainZoneDynhostRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneDynhostRecordArgs)(nil)).Elem()
}

type DomainZoneDynhostRecordInput interface {
	pulumi.Input

	ToDomainZoneDynhostRecordOutput() DomainZoneDynhostRecordOutput
	ToDomainZoneDynhostRecordOutputWithContext(ctx context.Context) DomainZoneDynhostRecordOutput
}

func (*DomainZoneDynhostRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneDynhostRecord)(nil)).Elem()
}

func (i *DomainZoneDynhostRecord) ToDomainZoneDynhostRecordOutput() DomainZoneDynhostRecordOutput {
	return i.ToDomainZoneDynhostRecordOutputWithContext(context.Background())
}

func (i *DomainZoneDynhostRecord) ToDomainZoneDynhostRecordOutputWithContext(ctx context.Context) DomainZoneDynhostRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneDynhostRecordOutput)
}

// DomainZoneDynhostRecordArrayInput is an input type that accepts DomainZoneDynhostRecordArray and DomainZoneDynhostRecordArrayOutput values.
// You can construct a concrete instance of `DomainZoneDynhostRecordArrayInput` via:
//
//	DomainZoneDynhostRecordArray{ DomainZoneDynhostRecordArgs{...} }
type DomainZoneDynhostRecordArrayInput interface {
	pulumi.Input

	ToDomainZoneDynhostRecordArrayOutput() DomainZoneDynhostRecordArrayOutput
	ToDomainZoneDynhostRecordArrayOutputWithContext(context.Context) DomainZoneDynhostRecordArrayOutput
}

type DomainZoneDynhostRecordArray []DomainZoneDynhostRecordInput

func (DomainZoneDynhostRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZoneDynhostRecord)(nil)).Elem()
}

func (i DomainZoneDynhostRecordArray) ToDomainZoneDynhostRecordArrayOutput() DomainZoneDynhostRecordArrayOutput {
	return i.ToDomainZoneDynhostRecordArrayOutputWithContext(context.Background())
}

func (i DomainZoneDynhostRecordArray) ToDomainZoneDynhostRecordArrayOutputWithContext(ctx context.Context) DomainZoneDynhostRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneDynhostRecordArrayOutput)
}

// DomainZoneDynhostRecordMapInput is an input type that accepts DomainZoneDynhostRecordMap and DomainZoneDynhostRecordMapOutput values.
// You can construct a concrete instance of `DomainZoneDynhostRecordMapInput` via:
//
//	DomainZoneDynhostRecordMap{ "key": DomainZoneDynhostRecordArgs{...} }
type DomainZoneDynhostRecordMapInput interface {
	pulumi.Input

	ToDomainZoneDynhostRecordMapOutput() DomainZoneDynhostRecordMapOutput
	ToDomainZoneDynhostRecordMapOutputWithContext(context.Context) DomainZoneDynhostRecordMapOutput
}

type DomainZoneDynhostRecordMap map[string]DomainZoneDynhostRecordInput

func (DomainZoneDynhostRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZoneDynhostRecord)(nil)).Elem()
}

func (i DomainZoneDynhostRecordMap) ToDomainZoneDynhostRecordMapOutput() DomainZoneDynhostRecordMapOutput {
	return i.ToDomainZoneDynhostRecordMapOutputWithContext(context.Background())
}

func (i DomainZoneDynhostRecordMap) ToDomainZoneDynhostRecordMapOutputWithContext(ctx context.Context) DomainZoneDynhostRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneDynhostRecordMapOutput)
}

type DomainZoneDynhostRecordOutput struct{ *pulumi.OutputState }

func (DomainZoneDynhostRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneDynhostRecord)(nil)).Elem()
}

func (o DomainZoneDynhostRecordOutput) ToDomainZoneDynhostRecordOutput() DomainZoneDynhostRecordOutput {
	return o
}

func (o DomainZoneDynhostRecordOutput) ToDomainZoneDynhostRecordOutputWithContext(ctx context.Context) DomainZoneDynhostRecordOutput {
	return o
}

// Record IP
func (o DomainZoneDynhostRecordOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneDynhostRecord) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Record sub-domain
func (o DomainZoneDynhostRecordOutput) SubDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneDynhostRecord) pulumi.StringOutput { return v.SubDomain }).(pulumi.StringOutput)
}

// Record TTL (Time to Live)
func (o DomainZoneDynhostRecordOutput) Ttl() pulumi.Float64Output {
	return o.ApplyT(func(v *DomainZoneDynhostRecord) pulumi.Float64Output { return v.Ttl }).(pulumi.Float64Output)
}

// Record zone
func (o DomainZoneDynhostRecordOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneDynhostRecord) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

// Zone name
func (o DomainZoneDynhostRecordOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneDynhostRecord) pulumi.StringOutput { return v.ZoneName }).(pulumi.StringOutput)
}

type DomainZoneDynhostRecordArrayOutput struct{ *pulumi.OutputState }

func (DomainZoneDynhostRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZoneDynhostRecord)(nil)).Elem()
}

func (o DomainZoneDynhostRecordArrayOutput) ToDomainZoneDynhostRecordArrayOutput() DomainZoneDynhostRecordArrayOutput {
	return o
}

func (o DomainZoneDynhostRecordArrayOutput) ToDomainZoneDynhostRecordArrayOutputWithContext(ctx context.Context) DomainZoneDynhostRecordArrayOutput {
	return o
}

func (o DomainZoneDynhostRecordArrayOutput) Index(i pulumi.IntInput) DomainZoneDynhostRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainZoneDynhostRecord {
		return vs[0].([]*DomainZoneDynhostRecord)[vs[1].(int)]
	}).(DomainZoneDynhostRecordOutput)
}

type DomainZoneDynhostRecordMapOutput struct{ *pulumi.OutputState }

func (DomainZoneDynhostRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZoneDynhostRecord)(nil)).Elem()
}

func (o DomainZoneDynhostRecordMapOutput) ToDomainZoneDynhostRecordMapOutput() DomainZoneDynhostRecordMapOutput {
	return o
}

func (o DomainZoneDynhostRecordMapOutput) ToDomainZoneDynhostRecordMapOutputWithContext(ctx context.Context) DomainZoneDynhostRecordMapOutput {
	return o
}

func (o DomainZoneDynhostRecordMapOutput) MapIndex(k pulumi.StringInput) DomainZoneDynhostRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainZoneDynhostRecord {
		return vs[0].(map[string]*DomainZoneDynhostRecord)[vs[1].(string)]
	}).(DomainZoneDynhostRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneDynhostRecordInput)(nil)).Elem(), &DomainZoneDynhostRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneDynhostRecordArrayInput)(nil)).Elem(), DomainZoneDynhostRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneDynhostRecordMapInput)(nil)).Elem(), DomainZoneDynhostRecordMap{})
	pulumi.RegisterOutputType(DomainZoneDynhostRecordOutput{})
	pulumi.RegisterOutputType(DomainZoneDynhostRecordArrayOutput{})
	pulumi.RegisterOutputType(DomainZoneDynhostRecordMapOutput{})
}
