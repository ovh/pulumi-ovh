// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Apply IP restrictions container registry associated with a public cloud project on artifact manager component.
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
//			_, err := cloudproject.GetContainerRegistry(ctx, &cloudproject.GetContainerRegistryArgs{
//				ServiceName: "XXXXXX",
//				RegistryId:  "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudproject.NewContainerRegistryIPRestrictionsRegistry(ctx, "my_registry_iprestrictions", &cloudproject.ContainerRegistryIPRestrictionsRegistryArgs{
//				ServiceName: pulumi.Any(registryOvhCloudProjectContainerregistry.ServiceName),
//				RegistryId:  pulumi.Any(registryOvhCloudProjectContainerregistry.Id),
//				IpRestrictions: pulumi.StringMapArray{
//					pulumi.StringMap{
//						"ip_block":    pulumi.String("xxx.xxx.xxx.xxx/xx"),
//						"description": pulumi.String("xxxxxxx"),
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
type ContainerRegistryIPRestrictionsRegistry struct {
	pulumi.CustomResourceState

	// IP restrictions applied on artifact manager component.
	IpRestrictions pulumi.StringMapArrayOutput `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewContainerRegistryIPRestrictionsRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryIPRestrictionsRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryIPRestrictionsRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryIPRestrictionsRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpRestrictions == nil {
		return nil, errors.New("invalid value for required argument 'IpRestrictions'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistryIPRestrictionsRegistry
	err := ctx.RegisterResource("ovh:CloudProject/containerRegistryIPRestrictionsRegistry:ContainerRegistryIPRestrictionsRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryIPRestrictionsRegistry gets an existing ContainerRegistryIPRestrictionsRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryIPRestrictionsRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryIPRestrictionsRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistryIPRestrictionsRegistry, error) {
	var resource ContainerRegistryIPRestrictionsRegistry
	err := ctx.ReadResource("ovh:CloudProject/containerRegistryIPRestrictionsRegistry:ContainerRegistryIPRestrictionsRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryIPRestrictionsRegistry resources.
type containerRegistryIPRestrictionsRegistryState struct {
	// IP restrictions applied on artifact manager component.
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId *string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type ContainerRegistryIPRestrictionsRegistryState struct {
	// IP restrictions applied on artifact manager component.
	IpRestrictions pulumi.StringMapArrayInput
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (ContainerRegistryIPRestrictionsRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIPRestrictionsRegistryState)(nil)).Elem()
}

type containerRegistryIPRestrictionsRegistryArgs struct {
	// IP restrictions applied on artifact manager component.
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistryIPRestrictionsRegistry resource.
type ContainerRegistryIPRestrictionsRegistryArgs struct {
	// IP restrictions applied on artifact manager component.
	IpRestrictions pulumi.StringMapArrayInput
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (ContainerRegistryIPRestrictionsRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIPRestrictionsRegistryArgs)(nil)).Elem()
}

type ContainerRegistryIPRestrictionsRegistryInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsRegistryOutput() ContainerRegistryIPRestrictionsRegistryOutput
	ToContainerRegistryIPRestrictionsRegistryOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryOutput
}

func (*ContainerRegistryIPRestrictionsRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (i *ContainerRegistryIPRestrictionsRegistry) ToContainerRegistryIPRestrictionsRegistryOutput() ContainerRegistryIPRestrictionsRegistryOutput {
	return i.ToContainerRegistryIPRestrictionsRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistryIPRestrictionsRegistry) ToContainerRegistryIPRestrictionsRegistryOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsRegistryOutput)
}

// ContainerRegistryIPRestrictionsRegistryArrayInput is an input type that accepts ContainerRegistryIPRestrictionsRegistryArray and ContainerRegistryIPRestrictionsRegistryArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryIPRestrictionsRegistryArrayInput` via:
//
//	ContainerRegistryIPRestrictionsRegistryArray{ ContainerRegistryIPRestrictionsRegistryArgs{...} }
type ContainerRegistryIPRestrictionsRegistryArrayInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsRegistryArrayOutput() ContainerRegistryIPRestrictionsRegistryArrayOutput
	ToContainerRegistryIPRestrictionsRegistryArrayOutputWithContext(context.Context) ContainerRegistryIPRestrictionsRegistryArrayOutput
}

type ContainerRegistryIPRestrictionsRegistryArray []ContainerRegistryIPRestrictionsRegistryInput

func (ContainerRegistryIPRestrictionsRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (i ContainerRegistryIPRestrictionsRegistryArray) ToContainerRegistryIPRestrictionsRegistryArrayOutput() ContainerRegistryIPRestrictionsRegistryArrayOutput {
	return i.ToContainerRegistryIPRestrictionsRegistryArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryIPRestrictionsRegistryArray) ToContainerRegistryIPRestrictionsRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsRegistryArrayOutput)
}

// ContainerRegistryIPRestrictionsRegistryMapInput is an input type that accepts ContainerRegistryIPRestrictionsRegistryMap and ContainerRegistryIPRestrictionsRegistryMapOutput values.
// You can construct a concrete instance of `ContainerRegistryIPRestrictionsRegistryMapInput` via:
//
//	ContainerRegistryIPRestrictionsRegistryMap{ "key": ContainerRegistryIPRestrictionsRegistryArgs{...} }
type ContainerRegistryIPRestrictionsRegistryMapInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsRegistryMapOutput() ContainerRegistryIPRestrictionsRegistryMapOutput
	ToContainerRegistryIPRestrictionsRegistryMapOutputWithContext(context.Context) ContainerRegistryIPRestrictionsRegistryMapOutput
}

type ContainerRegistryIPRestrictionsRegistryMap map[string]ContainerRegistryIPRestrictionsRegistryInput

func (ContainerRegistryIPRestrictionsRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (i ContainerRegistryIPRestrictionsRegistryMap) ToContainerRegistryIPRestrictionsRegistryMapOutput() ContainerRegistryIPRestrictionsRegistryMapOutput {
	return i.ToContainerRegistryIPRestrictionsRegistryMapOutputWithContext(context.Background())
}

func (i ContainerRegistryIPRestrictionsRegistryMap) ToContainerRegistryIPRestrictionsRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsRegistryMapOutput)
}

type ContainerRegistryIPRestrictionsRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsRegistryOutput) ToContainerRegistryIPRestrictionsRegistryOutput() ContainerRegistryIPRestrictionsRegistryOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsRegistryOutput) ToContainerRegistryIPRestrictionsRegistryOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryOutput {
	return o
}

// IP restrictions applied on artifact manager component.
func (o ContainerRegistryIPRestrictionsRegistryOutput) IpRestrictions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsRegistry) pulumi.StringMapArrayOutput { return v.IpRestrictions }).(pulumi.StringMapArrayOutput)
}

// The id of the Managed Private Registry.
func (o ContainerRegistryIPRestrictionsRegistryOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsRegistry) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o ContainerRegistryIPRestrictionsRegistryOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsRegistry) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type ContainerRegistryIPRestrictionsRegistryArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsRegistryArrayOutput) ToContainerRegistryIPRestrictionsRegistryArrayOutput() ContainerRegistryIPRestrictionsRegistryArrayOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsRegistryArrayOutput) ToContainerRegistryIPRestrictionsRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryArrayOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsRegistryArrayOutput) Index(i pulumi.IntInput) ContainerRegistryIPRestrictionsRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistryIPRestrictionsRegistry {
		return vs[0].([]*ContainerRegistryIPRestrictionsRegistry)[vs[1].(int)]
	}).(ContainerRegistryIPRestrictionsRegistryOutput)
}

type ContainerRegistryIPRestrictionsRegistryMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIPRestrictionsRegistry)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsRegistryMapOutput) ToContainerRegistryIPRestrictionsRegistryMapOutput() ContainerRegistryIPRestrictionsRegistryMapOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsRegistryMapOutput) ToContainerRegistryIPRestrictionsRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsRegistryMapOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsRegistryMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryIPRestrictionsRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistryIPRestrictionsRegistry {
		return vs[0].(map[string]*ContainerRegistryIPRestrictionsRegistry)[vs[1].(string)]
	}).(ContainerRegistryIPRestrictionsRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsRegistryInput)(nil)).Elem(), &ContainerRegistryIPRestrictionsRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsRegistryArrayInput)(nil)).Elem(), ContainerRegistryIPRestrictionsRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsRegistryMapInput)(nil)).Elem(), ContainerRegistryIPRestrictionsRegistryMap{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsRegistryOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsRegistryArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsRegistryMapOutput{})
}
