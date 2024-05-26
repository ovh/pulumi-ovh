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

// Creates a log subscrition for a cluster associated with a public cloud project.
//
// ## Example Usage
//
// Create a log subscription for a database.
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			stream, err := Dbaas.GetLogsOutputGraylogStream(ctx, &dbaas.GetLogsOutputGraylogStreamArgs{
//				ServiceName: "ldp-xx-xxxxx",
//				Title:       "my stream",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			db, err := CloudProjectDatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProjectDatabase.NewLogSubscription(ctx, "subscription", &CloudProjectDatabase.LogSubscriptionArgs{
//				ServiceName: pulumi.String(db.ServiceName),
//				Engine:      pulumi.String(db.Engine),
//				ClusterId:   pulumi.String(db.Id),
//				StreamId:    pulumi.String(stream.Id),
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
// OVHcloud Managed clusters logs subscription can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the subscription, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/logSubscription:LogSubscription sub service_name/engine/cluster_id/id
// ```
type LogSubscription struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation date of the subscription.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Log kind name of this subscription.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the destination log service.
	LdpServiceName pulumi.StringOutput `pulumi:"ldpServiceName"`
	// Identifier of the operation.
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// Name of subscribed resource, where the logs come from.
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// Type of subscribed resource, where the logs come from.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Id of the target Log data platform stream.
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// Last update date of the subscription.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewLogSubscription registers a new resource with the given unique name, arguments, and options.
func NewLogSubscription(ctx *pulumi.Context,
	name string, args *LogSubscriptionArgs, opts ...pulumi.ResourceOption) (*LogSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.StreamId == nil {
		return nil, errors.New("invalid value for required argument 'StreamId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ldpServiceName",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogSubscription
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/logSubscription:LogSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSubscription gets an existing LogSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSubscriptionState, opts ...pulumi.ResourceOption) (*LogSubscription, error) {
	var resource LogSubscription
	err := ctx.ReadResource("ovh:CloudProjectDatabase/logSubscription:LogSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSubscription resources.
type logSubscriptionState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Creation date of the subscription.
	CreatedAt *string `pulumi:"createdAt"`
	// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine *string `pulumi:"engine"`
	// Log kind name of this subscription.
	Kind *string `pulumi:"kind"`
	// Name of the destination log service.
	LdpServiceName *string `pulumi:"ldpServiceName"`
	// Identifier of the operation.
	OperationId *string `pulumi:"operationId"`
	// Name of subscribed resource, where the logs come from.
	ResourceName *string `pulumi:"resourceName"`
	// Type of subscribed resource, where the logs come from.
	ResourceType *string `pulumi:"resourceType"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Id of the target Log data platform stream.
	StreamId *string `pulumi:"streamId"`
	// Last update date of the subscription.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type LogSubscriptionState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Creation date of the subscription.
	CreatedAt pulumi.StringPtrInput
	// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringPtrInput
	// Log kind name of this subscription.
	Kind pulumi.StringPtrInput
	// Name of the destination log service.
	LdpServiceName pulumi.StringPtrInput
	// Identifier of the operation.
	OperationId pulumi.StringPtrInput
	// Name of subscribed resource, where the logs come from.
	ResourceName pulumi.StringPtrInput
	// Type of subscribed resource, where the logs come from.
	ResourceType pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Id of the target Log data platform stream.
	StreamId pulumi.StringPtrInput
	// Last update date of the subscription.
	UpdatedAt pulumi.StringPtrInput
}

func (LogSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSubscriptionState)(nil)).Elem()
}

type logSubscriptionArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Id of the target Log data platform stream.
	StreamId string `pulumi:"streamId"`
}

// The set of arguments for constructing a LogSubscription resource.
type LogSubscriptionArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// Id of the target Log data platform stream.
	StreamId pulumi.StringInput
}

func (LogSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSubscriptionArgs)(nil)).Elem()
}

type LogSubscriptionInput interface {
	pulumi.Input

	ToLogSubscriptionOutput() LogSubscriptionOutput
	ToLogSubscriptionOutputWithContext(ctx context.Context) LogSubscriptionOutput
}

func (*LogSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSubscription)(nil)).Elem()
}

func (i *LogSubscription) ToLogSubscriptionOutput() LogSubscriptionOutput {
	return i.ToLogSubscriptionOutputWithContext(context.Background())
}

func (i *LogSubscription) ToLogSubscriptionOutputWithContext(ctx context.Context) LogSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSubscriptionOutput)
}

// LogSubscriptionArrayInput is an input type that accepts LogSubscriptionArray and LogSubscriptionArrayOutput values.
// You can construct a concrete instance of `LogSubscriptionArrayInput` via:
//
//	LogSubscriptionArray{ LogSubscriptionArgs{...} }
type LogSubscriptionArrayInput interface {
	pulumi.Input

	ToLogSubscriptionArrayOutput() LogSubscriptionArrayOutput
	ToLogSubscriptionArrayOutputWithContext(context.Context) LogSubscriptionArrayOutput
}

type LogSubscriptionArray []LogSubscriptionInput

func (LogSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSubscription)(nil)).Elem()
}

func (i LogSubscriptionArray) ToLogSubscriptionArrayOutput() LogSubscriptionArrayOutput {
	return i.ToLogSubscriptionArrayOutputWithContext(context.Background())
}

func (i LogSubscriptionArray) ToLogSubscriptionArrayOutputWithContext(ctx context.Context) LogSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSubscriptionArrayOutput)
}

// LogSubscriptionMapInput is an input type that accepts LogSubscriptionMap and LogSubscriptionMapOutput values.
// You can construct a concrete instance of `LogSubscriptionMapInput` via:
//
//	LogSubscriptionMap{ "key": LogSubscriptionArgs{...} }
type LogSubscriptionMapInput interface {
	pulumi.Input

	ToLogSubscriptionMapOutput() LogSubscriptionMapOutput
	ToLogSubscriptionMapOutputWithContext(context.Context) LogSubscriptionMapOutput
}

type LogSubscriptionMap map[string]LogSubscriptionInput

func (LogSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSubscription)(nil)).Elem()
}

func (i LogSubscriptionMap) ToLogSubscriptionMapOutput() LogSubscriptionMapOutput {
	return i.ToLogSubscriptionMapOutputWithContext(context.Background())
}

func (i LogSubscriptionMap) ToLogSubscriptionMapOutputWithContext(ctx context.Context) LogSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSubscriptionMapOutput)
}

type LogSubscriptionOutput struct{ *pulumi.OutputState }

func (LogSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSubscription)(nil)).Elem()
}

func (o LogSubscriptionOutput) ToLogSubscriptionOutput() LogSubscriptionOutput {
	return o
}

func (o LogSubscriptionOutput) ToLogSubscriptionOutputWithContext(ctx context.Context) LogSubscriptionOutput {
	return o
}

// Cluster ID.
func (o LogSubscriptionOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation date of the subscription.
func (o LogSubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The database engine for which you want to manage a subscription. To get a full list of available engine visit.
// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
func (o LogSubscriptionOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Log kind name of this subscription.
func (o LogSubscriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the destination log service.
func (o LogSubscriptionOutput) LdpServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.LdpServiceName }).(pulumi.StringOutput)
}

// Identifier of the operation.
func (o LogSubscriptionOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.OperationId }).(pulumi.StringOutput)
}

// Name of subscribed resource, where the logs come from.
func (o LogSubscriptionOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// Type of subscribed resource, where the logs come from.
func (o LogSubscriptionOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o LogSubscriptionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Id of the target Log data platform stream.
func (o LogSubscriptionOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

// Last update date of the subscription.
func (o LogSubscriptionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type LogSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (LogSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSubscription)(nil)).Elem()
}

func (o LogSubscriptionArrayOutput) ToLogSubscriptionArrayOutput() LogSubscriptionArrayOutput {
	return o
}

func (o LogSubscriptionArrayOutput) ToLogSubscriptionArrayOutputWithContext(ctx context.Context) LogSubscriptionArrayOutput {
	return o
}

func (o LogSubscriptionArrayOutput) Index(i pulumi.IntInput) LogSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSubscription {
		return vs[0].([]*LogSubscription)[vs[1].(int)]
	}).(LogSubscriptionOutput)
}

type LogSubscriptionMapOutput struct{ *pulumi.OutputState }

func (LogSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSubscription)(nil)).Elem()
}

func (o LogSubscriptionMapOutput) ToLogSubscriptionMapOutput() LogSubscriptionMapOutput {
	return o
}

func (o LogSubscriptionMapOutput) ToLogSubscriptionMapOutputWithContext(ctx context.Context) LogSubscriptionMapOutput {
	return o
}

func (o LogSubscriptionMapOutput) MapIndex(k pulumi.StringInput) LogSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSubscription {
		return vs[0].(map[string]*LogSubscription)[vs[1].(string)]
	}).(LogSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSubscriptionInput)(nil)).Elem(), &LogSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSubscriptionArrayInput)(nil)).Elem(), LogSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSubscriptionMapInput)(nil)).Elem(), LogSubscriptionMap{})
	pulumi.RegisterOutputType(LogSubscriptionOutput{})
	pulumi.RegisterOutputType(LogSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(LogSubscriptionMapOutput{})
}
