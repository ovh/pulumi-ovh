// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kafka, err := CloudProjectDatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXX",
//				Engine:      "kafka",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProjectDatabase.NewKafkaSchemaRegistryAcl(ctx, "schemaRegistryAcl", &CloudProjectDatabase.KafkaSchemaRegistryAclArgs{
//				ServiceName: pulumi.String(kafka.ServiceName),
//				ClusterId:   pulumi.String(kafka.Id),
//				Permission:  pulumi.String("schema_registry_read"),
//				Resource:    pulumi.String("Subject:myResource"),
//				Username:    pulumi.String("johndoe"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
// ```
type KafkaSchemaRegistryAcl struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Permission to give to this username on this resource.
	// Available permissions:
	Permission pulumi.StringOutput `pulumi:"permission"`
	// Resource affected by this schema registry ACL.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Username affected by this schema registry ACL.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewKafkaSchemaRegistryAcl registers a new resource with the given unique name, arguments, and options.
func NewKafkaSchemaRegistryAcl(ctx *pulumi.Context,
	name string, args *KafkaSchemaRegistryAclArgs, opts ...pulumi.ResourceOption) (*KafkaSchemaRegistryAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KafkaSchemaRegistryAcl
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaSchemaRegistryAcl gets an existing KafkaSchemaRegistryAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaSchemaRegistryAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaSchemaRegistryAclState, opts ...pulumi.ResourceOption) (*KafkaSchemaRegistryAcl, error) {
	var resource KafkaSchemaRegistryAcl
	err := ctx.ReadResource("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaSchemaRegistryAcl resources.
type kafkaSchemaRegistryAclState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Permission to give to this username on this resource.
	// Available permissions:
	Permission *string `pulumi:"permission"`
	// Resource affected by this schema registry ACL.
	Resource *string `pulumi:"resource"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Username affected by this schema registry ACL.
	Username *string `pulumi:"username"`
}

type KafkaSchemaRegistryAclState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Permission to give to this username on this resource.
	// Available permissions:
	Permission pulumi.StringPtrInput
	// Resource affected by this schema registry ACL.
	Resource pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Username affected by this schema registry ACL.
	Username pulumi.StringPtrInput
}

func (KafkaSchemaRegistryAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaSchemaRegistryAclState)(nil)).Elem()
}

type kafkaSchemaRegistryAclArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Permission to give to this username on this resource.
	// Available permissions:
	Permission string `pulumi:"permission"`
	// Resource affected by this schema registry ACL.
	Resource string `pulumi:"resource"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Username affected by this schema registry ACL.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a KafkaSchemaRegistryAcl resource.
type KafkaSchemaRegistryAclArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Permission to give to this username on this resource.
	// Available permissions:
	Permission pulumi.StringInput
	// Resource affected by this schema registry ACL.
	Resource pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// Username affected by this schema registry ACL.
	Username pulumi.StringInput
}

func (KafkaSchemaRegistryAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaSchemaRegistryAclArgs)(nil)).Elem()
}

type KafkaSchemaRegistryAclInput interface {
	pulumi.Input

	ToKafkaSchemaRegistryAclOutput() KafkaSchemaRegistryAclOutput
	ToKafkaSchemaRegistryAclOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclOutput
}

func (*KafkaSchemaRegistryAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (i *KafkaSchemaRegistryAcl) ToKafkaSchemaRegistryAclOutput() KafkaSchemaRegistryAclOutput {
	return i.ToKafkaSchemaRegistryAclOutputWithContext(context.Background())
}

func (i *KafkaSchemaRegistryAcl) ToKafkaSchemaRegistryAclOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaSchemaRegistryAclOutput)
}

// KafkaSchemaRegistryAclArrayInput is an input type that accepts KafkaSchemaRegistryAclArray and KafkaSchemaRegistryAclArrayOutput values.
// You can construct a concrete instance of `KafkaSchemaRegistryAclArrayInput` via:
//
//	KafkaSchemaRegistryAclArray{ KafkaSchemaRegistryAclArgs{...} }
type KafkaSchemaRegistryAclArrayInput interface {
	pulumi.Input

	ToKafkaSchemaRegistryAclArrayOutput() KafkaSchemaRegistryAclArrayOutput
	ToKafkaSchemaRegistryAclArrayOutputWithContext(context.Context) KafkaSchemaRegistryAclArrayOutput
}

type KafkaSchemaRegistryAclArray []KafkaSchemaRegistryAclInput

func (KafkaSchemaRegistryAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (i KafkaSchemaRegistryAclArray) ToKafkaSchemaRegistryAclArrayOutput() KafkaSchemaRegistryAclArrayOutput {
	return i.ToKafkaSchemaRegistryAclArrayOutputWithContext(context.Background())
}

func (i KafkaSchemaRegistryAclArray) ToKafkaSchemaRegistryAclArrayOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaSchemaRegistryAclArrayOutput)
}

// KafkaSchemaRegistryAclMapInput is an input type that accepts KafkaSchemaRegistryAclMap and KafkaSchemaRegistryAclMapOutput values.
// You can construct a concrete instance of `KafkaSchemaRegistryAclMapInput` via:
//
//	KafkaSchemaRegistryAclMap{ "key": KafkaSchemaRegistryAclArgs{...} }
type KafkaSchemaRegistryAclMapInput interface {
	pulumi.Input

	ToKafkaSchemaRegistryAclMapOutput() KafkaSchemaRegistryAclMapOutput
	ToKafkaSchemaRegistryAclMapOutputWithContext(context.Context) KafkaSchemaRegistryAclMapOutput
}

type KafkaSchemaRegistryAclMap map[string]KafkaSchemaRegistryAclInput

func (KafkaSchemaRegistryAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (i KafkaSchemaRegistryAclMap) ToKafkaSchemaRegistryAclMapOutput() KafkaSchemaRegistryAclMapOutput {
	return i.ToKafkaSchemaRegistryAclMapOutputWithContext(context.Background())
}

func (i KafkaSchemaRegistryAclMap) ToKafkaSchemaRegistryAclMapOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaSchemaRegistryAclMapOutput)
}

type KafkaSchemaRegistryAclOutput struct{ *pulumi.OutputState }

func (KafkaSchemaRegistryAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (o KafkaSchemaRegistryAclOutput) ToKafkaSchemaRegistryAclOutput() KafkaSchemaRegistryAclOutput {
	return o
}

func (o KafkaSchemaRegistryAclOutput) ToKafkaSchemaRegistryAclOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclOutput {
	return o
}

// Cluster ID.
func (o KafkaSchemaRegistryAclOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaSchemaRegistryAcl) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Permission to give to this username on this resource.
// Available permissions:
func (o KafkaSchemaRegistryAclOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaSchemaRegistryAcl) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// Resource affected by this schema registry ACL.
func (o KafkaSchemaRegistryAclOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaSchemaRegistryAcl) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o KafkaSchemaRegistryAclOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaSchemaRegistryAcl) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Username affected by this schema registry ACL.
func (o KafkaSchemaRegistryAclOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaSchemaRegistryAcl) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type KafkaSchemaRegistryAclArrayOutput struct{ *pulumi.OutputState }

func (KafkaSchemaRegistryAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (o KafkaSchemaRegistryAclArrayOutput) ToKafkaSchemaRegistryAclArrayOutput() KafkaSchemaRegistryAclArrayOutput {
	return o
}

func (o KafkaSchemaRegistryAclArrayOutput) ToKafkaSchemaRegistryAclArrayOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclArrayOutput {
	return o
}

func (o KafkaSchemaRegistryAclArrayOutput) Index(i pulumi.IntInput) KafkaSchemaRegistryAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaSchemaRegistryAcl {
		return vs[0].([]*KafkaSchemaRegistryAcl)[vs[1].(int)]
	}).(KafkaSchemaRegistryAclOutput)
}

type KafkaSchemaRegistryAclMapOutput struct{ *pulumi.OutputState }

func (KafkaSchemaRegistryAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaSchemaRegistryAcl)(nil)).Elem()
}

func (o KafkaSchemaRegistryAclMapOutput) ToKafkaSchemaRegistryAclMapOutput() KafkaSchemaRegistryAclMapOutput {
	return o
}

func (o KafkaSchemaRegistryAclMapOutput) ToKafkaSchemaRegistryAclMapOutputWithContext(ctx context.Context) KafkaSchemaRegistryAclMapOutput {
	return o
}

func (o KafkaSchemaRegistryAclMapOutput) MapIndex(k pulumi.StringInput) KafkaSchemaRegistryAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaSchemaRegistryAcl {
		return vs[0].(map[string]*KafkaSchemaRegistryAcl)[vs[1].(string)]
	}).(KafkaSchemaRegistryAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaSchemaRegistryAclInput)(nil)).Elem(), &KafkaSchemaRegistryAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaSchemaRegistryAclArrayInput)(nil)).Elem(), KafkaSchemaRegistryAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaSchemaRegistryAclMapInput)(nil)).Elem(), KafkaSchemaRegistryAclMap{})
	pulumi.RegisterOutputType(KafkaSchemaRegistryAclOutput{})
	pulumi.RegisterOutputType(KafkaSchemaRegistryAclArrayOutput{})
	pulumi.RegisterOutputType(KafkaSchemaRegistryAclMapOutput{})
}
