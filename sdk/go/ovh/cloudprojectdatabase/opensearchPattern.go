// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/internal"
)

// Creates a pattern for a opensearch cluster associated with a public cloud project.
//
// ## Example Usage
//
// ## Import
//
// OVHcloud Managed opensearch clusters patterns can be imported using the `service_name`, `cluster_id` and `id` of the pattern, separated by "/" E.g., bash <break><break>```sh<break> $ pulumi import ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern my_pattern service_name/cluster_id/id <break>```<break><break>
type OpensearchPattern struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Maximum number of index for this pattern.
	MaxIndexCount pulumi.IntPtrOutput `pulumi:"maxIndexCount"`
	// Pattern format.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewOpensearchPattern registers a new resource with the given unique name, arguments, and options.
func NewOpensearchPattern(ctx *pulumi.Context,
	name string, args *OpensearchPatternArgs, opts ...pulumi.ResourceOption) (*OpensearchPattern, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OpensearchPattern
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpensearchPattern gets an existing OpensearchPattern resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpensearchPattern(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpensearchPatternState, opts ...pulumi.ResourceOption) (*OpensearchPattern, error) {
	var resource OpensearchPattern
	err := ctx.ReadResource("ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpensearchPattern resources.
type opensearchPatternState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Maximum number of index for this pattern.
	MaxIndexCount *int `pulumi:"maxIndexCount"`
	// Pattern format.
	Pattern *string `pulumi:"pattern"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type OpensearchPatternState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Maximum number of index for this pattern.
	MaxIndexCount pulumi.IntPtrInput
	// Pattern format.
	Pattern pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (OpensearchPatternState) ElementType() reflect.Type {
	return reflect.TypeOf((*opensearchPatternState)(nil)).Elem()
}

type opensearchPatternArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Maximum number of index for this pattern.
	MaxIndexCount *int `pulumi:"maxIndexCount"`
	// Pattern format.
	Pattern string `pulumi:"pattern"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a OpensearchPattern resource.
type OpensearchPatternArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Maximum number of index for this pattern.
	MaxIndexCount pulumi.IntPtrInput
	// Pattern format.
	Pattern pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (OpensearchPatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*opensearchPatternArgs)(nil)).Elem()
}

type OpensearchPatternInput interface {
	pulumi.Input

	ToOpensearchPatternOutput() OpensearchPatternOutput
	ToOpensearchPatternOutputWithContext(ctx context.Context) OpensearchPatternOutput
}

func (*OpensearchPattern) ElementType() reflect.Type {
	return reflect.TypeOf((**OpensearchPattern)(nil)).Elem()
}

func (i *OpensearchPattern) ToOpensearchPatternOutput() OpensearchPatternOutput {
	return i.ToOpensearchPatternOutputWithContext(context.Background())
}

func (i *OpensearchPattern) ToOpensearchPatternOutputWithContext(ctx context.Context) OpensearchPatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpensearchPatternOutput)
}

// OpensearchPatternArrayInput is an input type that accepts OpensearchPatternArray and OpensearchPatternArrayOutput values.
// You can construct a concrete instance of `OpensearchPatternArrayInput` via:
//
//	OpensearchPatternArray{ OpensearchPatternArgs{...} }
type OpensearchPatternArrayInput interface {
	pulumi.Input

	ToOpensearchPatternArrayOutput() OpensearchPatternArrayOutput
	ToOpensearchPatternArrayOutputWithContext(context.Context) OpensearchPatternArrayOutput
}

type OpensearchPatternArray []OpensearchPatternInput

func (OpensearchPatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpensearchPattern)(nil)).Elem()
}

func (i OpensearchPatternArray) ToOpensearchPatternArrayOutput() OpensearchPatternArrayOutput {
	return i.ToOpensearchPatternArrayOutputWithContext(context.Background())
}

func (i OpensearchPatternArray) ToOpensearchPatternArrayOutputWithContext(ctx context.Context) OpensearchPatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpensearchPatternArrayOutput)
}

// OpensearchPatternMapInput is an input type that accepts OpensearchPatternMap and OpensearchPatternMapOutput values.
// You can construct a concrete instance of `OpensearchPatternMapInput` via:
//
//	OpensearchPatternMap{ "key": OpensearchPatternArgs{...} }
type OpensearchPatternMapInput interface {
	pulumi.Input

	ToOpensearchPatternMapOutput() OpensearchPatternMapOutput
	ToOpensearchPatternMapOutputWithContext(context.Context) OpensearchPatternMapOutput
}

type OpensearchPatternMap map[string]OpensearchPatternInput

func (OpensearchPatternMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpensearchPattern)(nil)).Elem()
}

func (i OpensearchPatternMap) ToOpensearchPatternMapOutput() OpensearchPatternMapOutput {
	return i.ToOpensearchPatternMapOutputWithContext(context.Background())
}

func (i OpensearchPatternMap) ToOpensearchPatternMapOutputWithContext(ctx context.Context) OpensearchPatternMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpensearchPatternMapOutput)
}

type OpensearchPatternOutput struct{ *pulumi.OutputState }

func (OpensearchPatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpensearchPattern)(nil)).Elem()
}

func (o OpensearchPatternOutput) ToOpensearchPatternOutput() OpensearchPatternOutput {
	return o
}

func (o OpensearchPatternOutput) ToOpensearchPatternOutputWithContext(ctx context.Context) OpensearchPatternOutput {
	return o
}

// Cluster ID.
func (o OpensearchPatternOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *OpensearchPattern) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Maximum number of index for this pattern.
func (o OpensearchPatternOutput) MaxIndexCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OpensearchPattern) pulumi.IntPtrOutput { return v.MaxIndexCount }).(pulumi.IntPtrOutput)
}

// Pattern format.
func (o OpensearchPatternOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *OpensearchPattern) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o OpensearchPatternOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *OpensearchPattern) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type OpensearchPatternArrayOutput struct{ *pulumi.OutputState }

func (OpensearchPatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OpensearchPattern)(nil)).Elem()
}

func (o OpensearchPatternArrayOutput) ToOpensearchPatternArrayOutput() OpensearchPatternArrayOutput {
	return o
}

func (o OpensearchPatternArrayOutput) ToOpensearchPatternArrayOutputWithContext(ctx context.Context) OpensearchPatternArrayOutput {
	return o
}

func (o OpensearchPatternArrayOutput) Index(i pulumi.IntInput) OpensearchPatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OpensearchPattern {
		return vs[0].([]*OpensearchPattern)[vs[1].(int)]
	}).(OpensearchPatternOutput)
}

type OpensearchPatternMapOutput struct{ *pulumi.OutputState }

func (OpensearchPatternMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OpensearchPattern)(nil)).Elem()
}

func (o OpensearchPatternMapOutput) ToOpensearchPatternMapOutput() OpensearchPatternMapOutput {
	return o
}

func (o OpensearchPatternMapOutput) ToOpensearchPatternMapOutputWithContext(ctx context.Context) OpensearchPatternMapOutput {
	return o
}

func (o OpensearchPatternMapOutput) MapIndex(k pulumi.StringInput) OpensearchPatternOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OpensearchPattern {
		return vs[0].(map[string]*OpensearchPattern)[vs[1].(string)]
	}).(OpensearchPatternOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OpensearchPatternInput)(nil)).Elem(), &OpensearchPattern{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpensearchPatternArrayInput)(nil)).Elem(), OpensearchPatternArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OpensearchPatternMapInput)(nil)).Elem(), OpensearchPatternMap{})
	pulumi.RegisterOutputType(OpensearchPatternOutput{})
	pulumi.RegisterOutputType(OpensearchPatternArrayOutput{})
	pulumi.RegisterOutputType(OpensearchPatternMapOutput{})
}
