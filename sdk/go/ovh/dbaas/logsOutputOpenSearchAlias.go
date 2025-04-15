// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a DBaaS Logs Opensearch output alias.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.NewLogsOutputOpenSearchAlias(ctx, "alias", &dbaas.LogsOutputOpenSearchAliasArgs{
//				ServiceName: pulumi.String("...."),
//				Description: pulumi.String("my opensearch alias"),
//				Suffix:      pulumi.String("alias"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type LogsOutputOpenSearchAlias struct {
	pulumi.CustomResourceState

	// Alias Id
	AliasId pulumi.StringOutput `pulumi:"aliasId"`
	// Alias creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Current alias size (in bytes)
	CurrentSize pulumi.IntOutput `pulumi:"currentSize"`
	// Index description
	Description pulumi.StringOutput `pulumi:"description"`
	// List of attached indexes id
	Indexes pulumi.StringArrayOutput `pulumi:"indexes"`
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolOutput `pulumi:"isEditable"`
	// Alias name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of indices linked
	NbIndex pulumi.IntOutput `pulumi:"nbIndex"`
	// Number of streams linked
	NbStream pulumi.IntOutput `pulumi:"nbStream"`
	// The service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// List of attached streams id
	Streams pulumi.StringArrayOutput `pulumi:"streams"`
	// Index suffix
	Suffix pulumi.StringOutput `pulumi:"suffix"`
	// Input last update
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewLogsOutputOpenSearchAlias registers a new resource with the given unique name, arguments, and options.
func NewLogsOutputOpenSearchAlias(ctx *pulumi.Context,
	name string, args *LogsOutputOpenSearchAliasArgs, opts ...pulumi.ResourceOption) (*LogsOutputOpenSearchAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Suffix == nil {
		return nil, errors.New("invalid value for required argument 'Suffix'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogsOutputOpenSearchAlias
	err := ctx.RegisterResource("ovh:Dbaas/logsOutputOpenSearchAlias:LogsOutputOpenSearchAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsOutputOpenSearchAlias gets an existing LogsOutputOpenSearchAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsOutputOpenSearchAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsOutputOpenSearchAliasState, opts ...pulumi.ResourceOption) (*LogsOutputOpenSearchAlias, error) {
	var resource LogsOutputOpenSearchAlias
	err := ctx.ReadResource("ovh:Dbaas/logsOutputOpenSearchAlias:LogsOutputOpenSearchAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsOutputOpenSearchAlias resources.
type logsOutputOpenSearchAliasState struct {
	// Alias Id
	AliasId *string `pulumi:"aliasId"`
	// Alias creation
	CreatedAt *string `pulumi:"createdAt"`
	// Current alias size (in bytes)
	CurrentSize *int `pulumi:"currentSize"`
	// Index description
	Description *string `pulumi:"description"`
	// List of attached indexes id
	Indexes []string `pulumi:"indexes"`
	// Indicates if you are allowed to edit entry
	IsEditable *bool `pulumi:"isEditable"`
	// Alias name
	Name *string `pulumi:"name"`
	// Number of indices linked
	NbIndex *int `pulumi:"nbIndex"`
	// Number of streams linked
	NbStream *int `pulumi:"nbStream"`
	// The service name
	ServiceName *string `pulumi:"serviceName"`
	// List of attached streams id
	Streams []string `pulumi:"streams"`
	// Index suffix
	Suffix *string `pulumi:"suffix"`
	// Input last update
	UpdatedAt *string `pulumi:"updatedAt"`
}

type LogsOutputOpenSearchAliasState struct {
	// Alias Id
	AliasId pulumi.StringPtrInput
	// Alias creation
	CreatedAt pulumi.StringPtrInput
	// Current alias size (in bytes)
	CurrentSize pulumi.IntPtrInput
	// Index description
	Description pulumi.StringPtrInput
	// List of attached indexes id
	Indexes pulumi.StringArrayInput
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolPtrInput
	// Alias name
	Name pulumi.StringPtrInput
	// Number of indices linked
	NbIndex pulumi.IntPtrInput
	// Number of streams linked
	NbStream pulumi.IntPtrInput
	// The service name
	ServiceName pulumi.StringPtrInput
	// List of attached streams id
	Streams pulumi.StringArrayInput
	// Index suffix
	Suffix pulumi.StringPtrInput
	// Input last update
	UpdatedAt pulumi.StringPtrInput
}

func (LogsOutputOpenSearchAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputOpenSearchAliasState)(nil)).Elem()
}

type logsOutputOpenSearchAliasArgs struct {
	// Index description
	Description string `pulumi:"description"`
	// List of attached indexes id
	Indexes []string `pulumi:"indexes"`
	// Number of indices linked
	NbIndex *int `pulumi:"nbIndex"`
	// Number of streams linked
	NbStream *int `pulumi:"nbStream"`
	// The service name
	ServiceName string `pulumi:"serviceName"`
	// List of attached streams id
	Streams []string `pulumi:"streams"`
	// Index suffix
	Suffix string `pulumi:"suffix"`
}

// The set of arguments for constructing a LogsOutputOpenSearchAlias resource.
type LogsOutputOpenSearchAliasArgs struct {
	// Index description
	Description pulumi.StringInput
	// List of attached indexes id
	Indexes pulumi.StringArrayInput
	// Number of indices linked
	NbIndex pulumi.IntPtrInput
	// Number of streams linked
	NbStream pulumi.IntPtrInput
	// The service name
	ServiceName pulumi.StringInput
	// List of attached streams id
	Streams pulumi.StringArrayInput
	// Index suffix
	Suffix pulumi.StringInput
}

func (LogsOutputOpenSearchAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputOpenSearchAliasArgs)(nil)).Elem()
}

type LogsOutputOpenSearchAliasInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchAliasOutput() LogsOutputOpenSearchAliasOutput
	ToLogsOutputOpenSearchAliasOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasOutput
}

func (*LogsOutputOpenSearchAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (i *LogsOutputOpenSearchAlias) ToLogsOutputOpenSearchAliasOutput() LogsOutputOpenSearchAliasOutput {
	return i.ToLogsOutputOpenSearchAliasOutputWithContext(context.Background())
}

func (i *LogsOutputOpenSearchAlias) ToLogsOutputOpenSearchAliasOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchAliasOutput)
}

// LogsOutputOpenSearchAliasArrayInput is an input type that accepts LogsOutputOpenSearchAliasArray and LogsOutputOpenSearchAliasArrayOutput values.
// You can construct a concrete instance of `LogsOutputOpenSearchAliasArrayInput` via:
//
//	LogsOutputOpenSearchAliasArray{ LogsOutputOpenSearchAliasArgs{...} }
type LogsOutputOpenSearchAliasArrayInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchAliasArrayOutput() LogsOutputOpenSearchAliasArrayOutput
	ToLogsOutputOpenSearchAliasArrayOutputWithContext(context.Context) LogsOutputOpenSearchAliasArrayOutput
}

type LogsOutputOpenSearchAliasArray []LogsOutputOpenSearchAliasInput

func (LogsOutputOpenSearchAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (i LogsOutputOpenSearchAliasArray) ToLogsOutputOpenSearchAliasArrayOutput() LogsOutputOpenSearchAliasArrayOutput {
	return i.ToLogsOutputOpenSearchAliasArrayOutputWithContext(context.Background())
}

func (i LogsOutputOpenSearchAliasArray) ToLogsOutputOpenSearchAliasArrayOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchAliasArrayOutput)
}

// LogsOutputOpenSearchAliasMapInput is an input type that accepts LogsOutputOpenSearchAliasMap and LogsOutputOpenSearchAliasMapOutput values.
// You can construct a concrete instance of `LogsOutputOpenSearchAliasMapInput` via:
//
//	LogsOutputOpenSearchAliasMap{ "key": LogsOutputOpenSearchAliasArgs{...} }
type LogsOutputOpenSearchAliasMapInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchAliasMapOutput() LogsOutputOpenSearchAliasMapOutput
	ToLogsOutputOpenSearchAliasMapOutputWithContext(context.Context) LogsOutputOpenSearchAliasMapOutput
}

type LogsOutputOpenSearchAliasMap map[string]LogsOutputOpenSearchAliasInput

func (LogsOutputOpenSearchAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (i LogsOutputOpenSearchAliasMap) ToLogsOutputOpenSearchAliasMapOutput() LogsOutputOpenSearchAliasMapOutput {
	return i.ToLogsOutputOpenSearchAliasMapOutputWithContext(context.Background())
}

func (i LogsOutputOpenSearchAliasMap) ToLogsOutputOpenSearchAliasMapOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchAliasMapOutput)
}

type LogsOutputOpenSearchAliasOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (o LogsOutputOpenSearchAliasOutput) ToLogsOutputOpenSearchAliasOutput() LogsOutputOpenSearchAliasOutput {
	return o
}

func (o LogsOutputOpenSearchAliasOutput) ToLogsOutputOpenSearchAliasOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasOutput {
	return o
}

// Alias Id
func (o LogsOutputOpenSearchAliasOutput) AliasId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.AliasId }).(pulumi.StringOutput)
}

// Alias creation
func (o LogsOutputOpenSearchAliasOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Current alias size (in bytes)
func (o LogsOutputOpenSearchAliasOutput) CurrentSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.IntOutput { return v.CurrentSize }).(pulumi.IntOutput)
}

// Index description
func (o LogsOutputOpenSearchAliasOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// List of attached indexes id
func (o LogsOutputOpenSearchAliasOutput) Indexes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringArrayOutput { return v.Indexes }).(pulumi.StringArrayOutput)
}

// Indicates if you are allowed to edit entry
func (o LogsOutputOpenSearchAliasOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.BoolOutput { return v.IsEditable }).(pulumi.BoolOutput)
}

// Alias name
func (o LogsOutputOpenSearchAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of indices linked
func (o LogsOutputOpenSearchAliasOutput) NbIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.IntOutput { return v.NbIndex }).(pulumi.IntOutput)
}

// Number of streams linked
func (o LogsOutputOpenSearchAliasOutput) NbStream() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.IntOutput { return v.NbStream }).(pulumi.IntOutput)
}

// The service name
func (o LogsOutputOpenSearchAliasOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// List of attached streams id
func (o LogsOutputOpenSearchAliasOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringArrayOutput { return v.Streams }).(pulumi.StringArrayOutput)
}

// Index suffix
func (o LogsOutputOpenSearchAliasOutput) Suffix() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.Suffix }).(pulumi.StringOutput)
}

// Input last update
func (o LogsOutputOpenSearchAliasOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchAlias) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type LogsOutputOpenSearchAliasArrayOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (o LogsOutputOpenSearchAliasArrayOutput) ToLogsOutputOpenSearchAliasArrayOutput() LogsOutputOpenSearchAliasArrayOutput {
	return o
}

func (o LogsOutputOpenSearchAliasArrayOutput) ToLogsOutputOpenSearchAliasArrayOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasArrayOutput {
	return o
}

func (o LogsOutputOpenSearchAliasArrayOutput) Index(i pulumi.IntInput) LogsOutputOpenSearchAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogsOutputOpenSearchAlias {
		return vs[0].([]*LogsOutputOpenSearchAlias)[vs[1].(int)]
	}).(LogsOutputOpenSearchAliasOutput)
}

type LogsOutputOpenSearchAliasMapOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputOpenSearchAlias)(nil)).Elem()
}

func (o LogsOutputOpenSearchAliasMapOutput) ToLogsOutputOpenSearchAliasMapOutput() LogsOutputOpenSearchAliasMapOutput {
	return o
}

func (o LogsOutputOpenSearchAliasMapOutput) ToLogsOutputOpenSearchAliasMapOutputWithContext(ctx context.Context) LogsOutputOpenSearchAliasMapOutput {
	return o
}

func (o LogsOutputOpenSearchAliasMapOutput) MapIndex(k pulumi.StringInput) LogsOutputOpenSearchAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogsOutputOpenSearchAlias {
		return vs[0].(map[string]*LogsOutputOpenSearchAlias)[vs[1].(string)]
	}).(LogsOutputOpenSearchAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchAliasInput)(nil)).Elem(), &LogsOutputOpenSearchAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchAliasArrayInput)(nil)).Elem(), LogsOutputOpenSearchAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchAliasMapInput)(nil)).Elem(), LogsOutputOpenSearchAliasMap{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchAliasOutput{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchAliasArrayOutput{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchAliasMapOutput{})
}
