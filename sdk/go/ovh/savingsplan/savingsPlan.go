// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package savingsplan

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SavingsPlan struct {
	pulumi.CustomResourceState

	// Whether Savings Plan should be renewed at the end of the period (defaults to false)
	AutoRenewal pulumi.BoolOutput `pulumi:"autoRenewal"`
	// Custom display name, used in invoices
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// End date of the Savings Plan
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// Savings Plan flavor
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Periodicity of the Savings Plan
	Period pulumi.StringOutput `pulumi:"period"`
	// Action performed when reaching the end of the period
	PeriodEndAction pulumi.StringOutput `pulumi:"periodEndAction"`
	// End date of the current period
	PeriodEndDate pulumi.StringOutput `pulumi:"periodEndDate"`
	// Start date of the current period
	PeriodStartDate pulumi.StringOutput `pulumi:"periodStartDate"`
	// ID of the service
	ServiceId pulumi.IntOutput `pulumi:"serviceId"`
	// ID of the public cloud project
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Size of the Savings Plan
	Size pulumi.IntOutput `pulumi:"size"`
	// Start date of the Savings Plan
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Status of the Savings Plan
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSavingsPlan registers a new resource with the given unique name, arguments, and options.
func NewSavingsPlan(ctx *pulumi.Context,
	name string, args *SavingsPlanArgs, opts ...pulumi.ResourceOption) (*SavingsPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SavingsPlan
	err := ctx.RegisterResource("ovh:SavingsPlan/savingsPlan:SavingsPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSavingsPlan gets an existing SavingsPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSavingsPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavingsPlanState, opts ...pulumi.ResourceOption) (*SavingsPlan, error) {
	var resource SavingsPlan
	err := ctx.ReadResource("ovh:SavingsPlan/savingsPlan:SavingsPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SavingsPlan resources.
type savingsPlanState struct {
	// Whether Savings Plan should be renewed at the end of the period (defaults to false)
	AutoRenewal *bool `pulumi:"autoRenewal"`
	// Custom display name, used in invoices
	DisplayName *string `pulumi:"displayName"`
	// End date of the Savings Plan
	EndDate *string `pulumi:"endDate"`
	// Savings Plan flavor
	Flavor *string `pulumi:"flavor"`
	// Periodicity of the Savings Plan
	Period *string `pulumi:"period"`
	// Action performed when reaching the end of the period
	PeriodEndAction *string `pulumi:"periodEndAction"`
	// End date of the current period
	PeriodEndDate *string `pulumi:"periodEndDate"`
	// Start date of the current period
	PeriodStartDate *string `pulumi:"periodStartDate"`
	// ID of the service
	ServiceId *int `pulumi:"serviceId"`
	// ID of the public cloud project
	ServiceName *string `pulumi:"serviceName"`
	// Size of the Savings Plan
	Size *int `pulumi:"size"`
	// Start date of the Savings Plan
	StartDate *string `pulumi:"startDate"`
	// Status of the Savings Plan
	Status *string `pulumi:"status"`
}

type SavingsPlanState struct {
	// Whether Savings Plan should be renewed at the end of the period (defaults to false)
	AutoRenewal pulumi.BoolPtrInput
	// Custom display name, used in invoices
	DisplayName pulumi.StringPtrInput
	// End date of the Savings Plan
	EndDate pulumi.StringPtrInput
	// Savings Plan flavor
	Flavor pulumi.StringPtrInput
	// Periodicity of the Savings Plan
	Period pulumi.StringPtrInput
	// Action performed when reaching the end of the period
	PeriodEndAction pulumi.StringPtrInput
	// End date of the current period
	PeriodEndDate pulumi.StringPtrInput
	// Start date of the current period
	PeriodStartDate pulumi.StringPtrInput
	// ID of the service
	ServiceId pulumi.IntPtrInput
	// ID of the public cloud project
	ServiceName pulumi.StringPtrInput
	// Size of the Savings Plan
	Size pulumi.IntPtrInput
	// Start date of the Savings Plan
	StartDate pulumi.StringPtrInput
	// Status of the Savings Plan
	Status pulumi.StringPtrInput
}

func (SavingsPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*savingsPlanState)(nil)).Elem()
}

type savingsPlanArgs struct {
	// Whether Savings Plan should be renewed at the end of the period (defaults to false)
	AutoRenewal *bool `pulumi:"autoRenewal"`
	// Custom display name, used in invoices
	DisplayName string `pulumi:"displayName"`
	// Savings Plan flavor
	Flavor string `pulumi:"flavor"`
	// Periodicity of the Savings Plan
	Period string `pulumi:"period"`
	// ID of the public cloud project
	ServiceName string `pulumi:"serviceName"`
	// Size of the Savings Plan
	Size int `pulumi:"size"`
}

// The set of arguments for constructing a SavingsPlan resource.
type SavingsPlanArgs struct {
	// Whether Savings Plan should be renewed at the end of the period (defaults to false)
	AutoRenewal pulumi.BoolPtrInput
	// Custom display name, used in invoices
	DisplayName pulumi.StringInput
	// Savings Plan flavor
	Flavor pulumi.StringInput
	// Periodicity of the Savings Plan
	Period pulumi.StringInput
	// ID of the public cloud project
	ServiceName pulumi.StringInput
	// Size of the Savings Plan
	Size pulumi.IntInput
}

func (SavingsPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savingsPlanArgs)(nil)).Elem()
}

type SavingsPlanInput interface {
	pulumi.Input

	ToSavingsPlanOutput() SavingsPlanOutput
	ToSavingsPlanOutputWithContext(ctx context.Context) SavingsPlanOutput
}

func (*SavingsPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**SavingsPlan)(nil)).Elem()
}

func (i *SavingsPlan) ToSavingsPlanOutput() SavingsPlanOutput {
	return i.ToSavingsPlanOutputWithContext(context.Background())
}

func (i *SavingsPlan) ToSavingsPlanOutputWithContext(ctx context.Context) SavingsPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanOutput)
}

// SavingsPlanArrayInput is an input type that accepts SavingsPlanArray and SavingsPlanArrayOutput values.
// You can construct a concrete instance of `SavingsPlanArrayInput` via:
//
//	SavingsPlanArray{ SavingsPlanArgs{...} }
type SavingsPlanArrayInput interface {
	pulumi.Input

	ToSavingsPlanArrayOutput() SavingsPlanArrayOutput
	ToSavingsPlanArrayOutputWithContext(context.Context) SavingsPlanArrayOutput
}

type SavingsPlanArray []SavingsPlanInput

func (SavingsPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SavingsPlan)(nil)).Elem()
}

func (i SavingsPlanArray) ToSavingsPlanArrayOutput() SavingsPlanArrayOutput {
	return i.ToSavingsPlanArrayOutputWithContext(context.Background())
}

func (i SavingsPlanArray) ToSavingsPlanArrayOutputWithContext(ctx context.Context) SavingsPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanArrayOutput)
}

// SavingsPlanMapInput is an input type that accepts SavingsPlanMap and SavingsPlanMapOutput values.
// You can construct a concrete instance of `SavingsPlanMapInput` via:
//
//	SavingsPlanMap{ "key": SavingsPlanArgs{...} }
type SavingsPlanMapInput interface {
	pulumi.Input

	ToSavingsPlanMapOutput() SavingsPlanMapOutput
	ToSavingsPlanMapOutputWithContext(context.Context) SavingsPlanMapOutput
}

type SavingsPlanMap map[string]SavingsPlanInput

func (SavingsPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SavingsPlan)(nil)).Elem()
}

func (i SavingsPlanMap) ToSavingsPlanMapOutput() SavingsPlanMapOutput {
	return i.ToSavingsPlanMapOutputWithContext(context.Background())
}

func (i SavingsPlanMap) ToSavingsPlanMapOutputWithContext(ctx context.Context) SavingsPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavingsPlanMapOutput)
}

type SavingsPlanOutput struct{ *pulumi.OutputState }

func (SavingsPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SavingsPlan)(nil)).Elem()
}

func (o SavingsPlanOutput) ToSavingsPlanOutput() SavingsPlanOutput {
	return o
}

func (o SavingsPlanOutput) ToSavingsPlanOutputWithContext(ctx context.Context) SavingsPlanOutput {
	return o
}

// Whether Savings Plan should be renewed at the end of the period (defaults to false)
func (o SavingsPlanOutput) AutoRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.BoolOutput { return v.AutoRenewal }).(pulumi.BoolOutput)
}

// Custom display name, used in invoices
func (o SavingsPlanOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// End date of the Savings Plan
func (o SavingsPlanOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

// Savings Plan flavor
func (o SavingsPlanOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// Periodicity of the Savings Plan
func (o SavingsPlanOutput) Period() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.Period }).(pulumi.StringOutput)
}

// Action performed when reaching the end of the period
func (o SavingsPlanOutput) PeriodEndAction() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.PeriodEndAction }).(pulumi.StringOutput)
}

// End date of the current period
func (o SavingsPlanOutput) PeriodEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.PeriodEndDate }).(pulumi.StringOutput)
}

// Start date of the current period
func (o SavingsPlanOutput) PeriodStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.PeriodStartDate }).(pulumi.StringOutput)
}

// ID of the service
func (o SavingsPlanOutput) ServiceId() pulumi.IntOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.IntOutput { return v.ServiceId }).(pulumi.IntOutput)
}

// ID of the public cloud project
func (o SavingsPlanOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Size of the Savings Plan
func (o SavingsPlanOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Start date of the Savings Plan
func (o SavingsPlanOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Status of the Savings Plan
func (o SavingsPlanOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SavingsPlan) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SavingsPlanArrayOutput struct{ *pulumi.OutputState }

func (SavingsPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SavingsPlan)(nil)).Elem()
}

func (o SavingsPlanArrayOutput) ToSavingsPlanArrayOutput() SavingsPlanArrayOutput {
	return o
}

func (o SavingsPlanArrayOutput) ToSavingsPlanArrayOutputWithContext(ctx context.Context) SavingsPlanArrayOutput {
	return o
}

func (o SavingsPlanArrayOutput) Index(i pulumi.IntInput) SavingsPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SavingsPlan {
		return vs[0].([]*SavingsPlan)[vs[1].(int)]
	}).(SavingsPlanOutput)
}

type SavingsPlanMapOutput struct{ *pulumi.OutputState }

func (SavingsPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SavingsPlan)(nil)).Elem()
}

func (o SavingsPlanMapOutput) ToSavingsPlanMapOutput() SavingsPlanMapOutput {
	return o
}

func (o SavingsPlanMapOutput) ToSavingsPlanMapOutputWithContext(ctx context.Context) SavingsPlanMapOutput {
	return o
}

func (o SavingsPlanMapOutput) MapIndex(k pulumi.StringInput) SavingsPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SavingsPlan {
		return vs[0].(map[string]*SavingsPlan)[vs[1].(string)]
	}).(SavingsPlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SavingsPlanInput)(nil)).Elem(), &SavingsPlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*SavingsPlanArrayInput)(nil)).Elem(), SavingsPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SavingsPlanMapInput)(nil)).Elem(), SavingsPlanMap{})
	pulumi.RegisterOutputType(SavingsPlanOutput{})
	pulumi.RegisterOutputType(SavingsPlanArrayOutput{})
	pulumi.RegisterOutputType(SavingsPlanMapOutput{})
}
