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

// Creates an ACL for a kafka cluster associated with a public cloud project.
//
// ## Example Usage
//
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
//			_, err = CloudProjectDatabase.NewKafkaAcl(ctx, "acl", &CloudProjectDatabase.KafkaAclArgs{
//				ServiceName: pulumi.String(kafka.ServiceName),
//				ClusterId:   pulumi.String(kafka.Id),
//				Permission:  pulumi.String("read"),
//				Topic:       pulumi.String("mytopic"),
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
//
// ## Import
//
// OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
// ```
type KafkaAcl struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Permission to give to this username on this topic.
	// Available permissions:
	Permission pulumi.StringOutput `pulumi:"permission"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Topic affected by this ACL.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// Username affected by this ACL.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewKafkaAcl registers a new resource with the given unique name, arguments, and options.
func NewKafkaAcl(ctx *pulumi.Context,
	name string, args *KafkaAclArgs, opts ...pulumi.ResourceOption) (*KafkaAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KafkaAcl
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaAcl gets an existing KafkaAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaAclState, opts ...pulumi.ResourceOption) (*KafkaAcl, error) {
	var resource KafkaAcl
	err := ctx.ReadResource("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaAcl resources.
type kafkaAclState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Permission to give to this username on this topic.
	// Available permissions:
	Permission *string `pulumi:"permission"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Topic affected by this ACL.
	Topic *string `pulumi:"topic"`
	// Username affected by this ACL.
	Username *string `pulumi:"username"`
}

type KafkaAclState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Permission to give to this username on this topic.
	// Available permissions:
	Permission pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Topic affected by this ACL.
	Topic pulumi.StringPtrInput
	// Username affected by this ACL.
	Username pulumi.StringPtrInput
}

func (KafkaAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaAclState)(nil)).Elem()
}

type kafkaAclArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Permission to give to this username on this topic.
	// Available permissions:
	Permission string `pulumi:"permission"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Topic affected by this ACL.
	Topic string `pulumi:"topic"`
	// Username affected by this ACL.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a KafkaAcl resource.
type KafkaAclArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Permission to give to this username on this topic.
	// Available permissions:
	Permission pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// Topic affected by this ACL.
	Topic pulumi.StringInput
	// Username affected by this ACL.
	Username pulumi.StringInput
}

func (KafkaAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaAclArgs)(nil)).Elem()
}

type KafkaAclInput interface {
	pulumi.Input

	ToKafkaAclOutput() KafkaAclOutput
	ToKafkaAclOutputWithContext(ctx context.Context) KafkaAclOutput
}

func (*KafkaAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaAcl)(nil)).Elem()
}

func (i *KafkaAcl) ToKafkaAclOutput() KafkaAclOutput {
	return i.ToKafkaAclOutputWithContext(context.Background())
}

func (i *KafkaAcl) ToKafkaAclOutputWithContext(ctx context.Context) KafkaAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaAclOutput)
}

// KafkaAclArrayInput is an input type that accepts KafkaAclArray and KafkaAclArrayOutput values.
// You can construct a concrete instance of `KafkaAclArrayInput` via:
//
//	KafkaAclArray{ KafkaAclArgs{...} }
type KafkaAclArrayInput interface {
	pulumi.Input

	ToKafkaAclArrayOutput() KafkaAclArrayOutput
	ToKafkaAclArrayOutputWithContext(context.Context) KafkaAclArrayOutput
}

type KafkaAclArray []KafkaAclInput

func (KafkaAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaAcl)(nil)).Elem()
}

func (i KafkaAclArray) ToKafkaAclArrayOutput() KafkaAclArrayOutput {
	return i.ToKafkaAclArrayOutputWithContext(context.Background())
}

func (i KafkaAclArray) ToKafkaAclArrayOutputWithContext(ctx context.Context) KafkaAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaAclArrayOutput)
}

// KafkaAclMapInput is an input type that accepts KafkaAclMap and KafkaAclMapOutput values.
// You can construct a concrete instance of `KafkaAclMapInput` via:
//
//	KafkaAclMap{ "key": KafkaAclArgs{...} }
type KafkaAclMapInput interface {
	pulumi.Input

	ToKafkaAclMapOutput() KafkaAclMapOutput
	ToKafkaAclMapOutputWithContext(context.Context) KafkaAclMapOutput
}

type KafkaAclMap map[string]KafkaAclInput

func (KafkaAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaAcl)(nil)).Elem()
}

func (i KafkaAclMap) ToKafkaAclMapOutput() KafkaAclMapOutput {
	return i.ToKafkaAclMapOutputWithContext(context.Background())
}

func (i KafkaAclMap) ToKafkaAclMapOutputWithContext(ctx context.Context) KafkaAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaAclMapOutput)
}

type KafkaAclOutput struct{ *pulumi.OutputState }

func (KafkaAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaAcl)(nil)).Elem()
}

func (o KafkaAclOutput) ToKafkaAclOutput() KafkaAclOutput {
	return o
}

func (o KafkaAclOutput) ToKafkaAclOutputWithContext(ctx context.Context) KafkaAclOutput {
	return o
}

// Cluster ID.
func (o KafkaAclOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaAcl) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Permission to give to this username on this topic.
// Available permissions:
func (o KafkaAclOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaAcl) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o KafkaAclOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaAcl) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Topic affected by this ACL.
func (o KafkaAclOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaAcl) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// Username affected by this ACL.
func (o KafkaAclOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaAcl) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type KafkaAclArrayOutput struct{ *pulumi.OutputState }

func (KafkaAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaAcl)(nil)).Elem()
}

func (o KafkaAclArrayOutput) ToKafkaAclArrayOutput() KafkaAclArrayOutput {
	return o
}

func (o KafkaAclArrayOutput) ToKafkaAclArrayOutputWithContext(ctx context.Context) KafkaAclArrayOutput {
	return o
}

func (o KafkaAclArrayOutput) Index(i pulumi.IntInput) KafkaAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaAcl {
		return vs[0].([]*KafkaAcl)[vs[1].(int)]
	}).(KafkaAclOutput)
}

type KafkaAclMapOutput struct{ *pulumi.OutputState }

func (KafkaAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaAcl)(nil)).Elem()
}

func (o KafkaAclMapOutput) ToKafkaAclMapOutput() KafkaAclMapOutput {
	return o
}

func (o KafkaAclMapOutput) ToKafkaAclMapOutputWithContext(ctx context.Context) KafkaAclMapOutput {
	return o
}

func (o KafkaAclMapOutput) MapIndex(k pulumi.StringInput) KafkaAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaAcl {
		return vs[0].(map[string]*KafkaAcl)[vs[1].(string)]
	}).(KafkaAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaAclInput)(nil)).Elem(), &KafkaAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaAclArrayInput)(nil)).Elem(), KafkaAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaAclMapInput)(nil)).Elem(), KafkaAclMap{})
	pulumi.RegisterOutputType(KafkaAclOutput{})
	pulumi.RegisterOutputType(KafkaAclArrayOutput{})
	pulumi.RegisterOutputType(KafkaAclMapOutput{})
}
