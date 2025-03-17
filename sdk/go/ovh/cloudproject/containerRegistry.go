// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a container registry associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			regcap, err := cloudproject.GetCapabilitiesContainerFilter(ctx, &cloudproject.GetCapabilitiesContainerFilterArgs{
//				ServiceName: "XXXXXX",
//				PlanName:    "SMALL",
//				Region:      "GRA",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudproject.NewContainerRegistry(ctx, "myRegistry", &cloudproject.ContainerRegistryArgs{
//				ServiceName: pulumi.String(regcap.ServiceName),
//				PlanId:      pulumi.String(regcap.Id),
//				Region:      pulumi.String(regcap.Region),
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
// > __WARNING__ You can update and migrate to a higher plan at any time but not the contrary.
type ContainerRegistry struct {
	pulumi.CustomResourceState

	// Plan creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Registry name
	Name pulumi.StringOutput `pulumi:"name"`
	// Plan ID of the registry
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// Plan of the registry
	Plans ContainerRegistryPlanArrayOutput `pulumi:"plans"`
	// Project ID of your registry
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Region of the registry
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size pulumi.IntOutput `pulumi:"size"`
	// Registry status
	Status pulumi.StringOutput `pulumi:"status"`
	// Registry last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Access url of the registry
	Url pulumi.StringOutput `pulumi:"url"`
	// Version of your registry
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistry
	err := ctx.RegisterResource("ovh:CloudProject/containerRegistry:ContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	err := ctx.ReadResource("ovh:CloudProject/containerRegistry:ContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
	// Plan creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry
	PlanId *string `pulumi:"planId"`
	// Plan of the registry
	Plans []ContainerRegistryPlan `pulumi:"plans"`
	// Project ID of your registry
	ProjectId *string `pulumi:"projectId"`
	// Region of the registry
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size *int `pulumi:"size"`
	// Registry status
	Status *string `pulumi:"status"`
	// Registry last update date
	UpdatedAt *string `pulumi:"updatedAt"`
	// Access url of the registry
	Url *string `pulumi:"url"`
	// Version of your registry
	Version *string `pulumi:"version"`
}

type ContainerRegistryState struct {
	// Plan creation date
	CreatedAt pulumi.StringPtrInput
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry
	PlanId pulumi.StringPtrInput
	// Plan of the registry
	Plans ContainerRegistryPlanArrayInput
	// Project ID of your registry
	ProjectId pulumi.StringPtrInput
	// Region of the registry
	Region pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current size of the registry (bytes)
	Size pulumi.IntPtrInput
	// Registry status
	Status pulumi.StringPtrInput
	// Registry last update date
	UpdatedAt pulumi.StringPtrInput
	// Access url of the registry
	Url pulumi.StringPtrInput
	// Version of your registry
	Version pulumi.StringPtrInput
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry
	PlanId *string `pulumi:"planId"`
	// Region of the registry
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry
	PlanId pulumi.StringPtrInput
	// Region of the registry
	Region pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}

type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput
}

func (*ContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *ContainerRegistry) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistry) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

// ContainerRegistryArrayInput is an input type that accepts ContainerRegistryArray and ContainerRegistryArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryArrayInput` via:
//
//	ContainerRegistryArray{ ContainerRegistryArgs{...} }
type ContainerRegistryArrayInput interface {
	pulumi.Input

	ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput
	ToContainerRegistryArrayOutputWithContext(context.Context) ContainerRegistryArrayOutput
}

type ContainerRegistryArray []ContainerRegistryInput

func (ContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return i.ToContainerRegistryArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryArrayOutput)
}

// ContainerRegistryMapInput is an input type that accepts ContainerRegistryMap and ContainerRegistryMapOutput values.
// You can construct a concrete instance of `ContainerRegistryMapInput` via:
//
//	ContainerRegistryMap{ "key": ContainerRegistryArgs{...} }
type ContainerRegistryMapInput interface {
	pulumi.Input

	ToContainerRegistryMapOutput() ContainerRegistryMapOutput
	ToContainerRegistryMapOutputWithContext(context.Context) ContainerRegistryMapOutput
}

type ContainerRegistryMap map[string]ContainerRegistryInput

func (ContainerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return i.ToContainerRegistryMapOutputWithContext(context.Background())
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryMapOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

// Plan creation date
func (o ContainerRegistryOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Registry name
func (o ContainerRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Plan ID of the registry
func (o ContainerRegistryOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// Plan of the registry
func (o ContainerRegistryOutput) Plans() ContainerRegistryPlanArrayOutput {
	return o.ApplyT(func(v *ContainerRegistry) ContainerRegistryPlanArrayOutput { return v.Plans }).(ContainerRegistryPlanArrayOutput)
}

// Project ID of your registry
func (o ContainerRegistryOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Region of the registry
func (o ContainerRegistryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o ContainerRegistryOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current size of the registry (bytes)
func (o ContainerRegistryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Registry status
func (o ContainerRegistryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Registry last update date
func (o ContainerRegistryOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Access url of the registry
func (o ContainerRegistryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Version of your registry
func (o ContainerRegistryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) Index(i pulumi.IntInput) ContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].([]*ContainerRegistry)[vs[1].(int)]
	}).(ContainerRegistryOutput)
}

type ContainerRegistryMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].(map[string]*ContainerRegistry)[vs[1].(string)]
	}).(ContainerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryInput)(nil)).Elem(), &ContainerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryArrayInput)(nil)).Elem(), ContainerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryMapInput)(nil)).Elem(), ContainerRegistryMap{})
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
	pulumi.RegisterOutputType(ContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryMapOutput{})
}
