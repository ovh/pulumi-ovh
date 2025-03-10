// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogSubscription struct {
	pulumi.CustomResourceState

	// Id of the database cluster
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation date of the subscription
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the engine of the service
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Log kind name of this subscription
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the destination log service
	LdpServiceName pulumi.StringOutput `pulumi:"ldpServiceName"`
	// Identifier of the operation
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// Name of subscribed resource, where the logs come from
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// Type of subscribed resource, where the logs come from
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	ServiceName  pulumi.StringOutput `pulumi:"serviceName"`
	// Id of the target Log data platform stream
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// Last update date of the subscription
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
	// Id of the database cluster
	ClusterId *string `pulumi:"clusterId"`
	// Creation date of the subscription
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the engine of the service
	Engine *string `pulumi:"engine"`
	// Log kind name of this subscription
	Kind *string `pulumi:"kind"`
	// Name of the destination log service
	LdpServiceName *string `pulumi:"ldpServiceName"`
	// Identifier of the operation
	OperationId *string `pulumi:"operationId"`
	// Name of subscribed resource, where the logs come from
	ResourceName *string `pulumi:"resourceName"`
	// Type of subscribed resource, where the logs come from
	ResourceType *string `pulumi:"resourceType"`
	ServiceName  *string `pulumi:"serviceName"`
	// Id of the target Log data platform stream
	StreamId *string `pulumi:"streamId"`
	// Last update date of the subscription
	UpdatedAt *string `pulumi:"updatedAt"`
}

type LogSubscriptionState struct {
	// Id of the database cluster
	ClusterId pulumi.StringPtrInput
	// Creation date of the subscription
	CreatedAt pulumi.StringPtrInput
	// Name of the engine of the service
	Engine pulumi.StringPtrInput
	// Log kind name of this subscription
	Kind pulumi.StringPtrInput
	// Name of the destination log service
	LdpServiceName pulumi.StringPtrInput
	// Identifier of the operation
	OperationId pulumi.StringPtrInput
	// Name of subscribed resource, where the logs come from
	ResourceName pulumi.StringPtrInput
	// Type of subscribed resource, where the logs come from
	ResourceType pulumi.StringPtrInput
	ServiceName  pulumi.StringPtrInput
	// Id of the target Log data platform stream
	StreamId pulumi.StringPtrInput
	// Last update date of the subscription
	UpdatedAt pulumi.StringPtrInput
}

func (LogSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSubscriptionState)(nil)).Elem()
}

type logSubscriptionArgs struct {
	// Id of the database cluster
	ClusterId string `pulumi:"clusterId"`
	// Name of the engine of the service
	Engine      string `pulumi:"engine"`
	ServiceName string `pulumi:"serviceName"`
	// Id of the target Log data platform stream
	StreamId string `pulumi:"streamId"`
}

// The set of arguments for constructing a LogSubscription resource.
type LogSubscriptionArgs struct {
	// Id of the database cluster
	ClusterId pulumi.StringInput
	// Name of the engine of the service
	Engine      pulumi.StringInput
	ServiceName pulumi.StringInput
	// Id of the target Log data platform stream
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

// Id of the database cluster
func (o LogSubscriptionOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation date of the subscription
func (o LogSubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the engine of the service
func (o LogSubscriptionOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Log kind name of this subscription
func (o LogSubscriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the destination log service
func (o LogSubscriptionOutput) LdpServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.LdpServiceName }).(pulumi.StringOutput)
}

// Identifier of the operation
func (o LogSubscriptionOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.OperationId }).(pulumi.StringOutput)
}

// Name of subscribed resource, where the logs come from
func (o LogSubscriptionOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// Type of subscribed resource, where the logs come from
func (o LogSubscriptionOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LogSubscriptionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Id of the target Log data platform stream
func (o LogSubscriptionOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSubscription) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

// Last update date of the subscription
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
