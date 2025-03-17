// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Subscribe to a Managed Loadbalance Logs Service in a public cloud project.
//
// ## Example Usage
//
// # Create a subscription
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudproject.NewRegionLoadBalancerLogSubscription(ctx, "subscription", &cloudproject.RegionLoadBalancerLogSubscriptionArgs{
//				Kind:           pulumi.String("haproxy"),
//				LoadbalancerId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
//				RegionName:     pulumi.String("yyyy"),
//				ServiceName:    pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				StreamId:       pulumi.String("ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RegionLoadBalancerLogSubscription struct {
	pulumi.CustomResourceState

	// The date of the subscription creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// haproxy  **Changing this value recreates the resource.**
	Kind pulumi.StringOutput `pulumi:"kind"`
	// LDP service name
	LdpServiceName pulumi.StringOutput `pulumi:"ldpServiceName"`
	// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// The operation ID
	OperationId pulumi.StringOutput `pulumi:"operationId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
	RegionName pulumi.StringOutput `pulumi:"regionName"`
	// The resource name
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// The resource type
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Data stream id to use for the subscription  **Changing this value recreates the resource.**
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// The subscription id
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// The last update of the subscription
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewRegionLoadBalancerLogSubscription registers a new resource with the given unique name, arguments, and options.
func NewRegionLoadBalancerLogSubscription(ctx *pulumi.Context,
	name string, args *RegionLoadBalancerLogSubscriptionArgs, opts ...pulumi.ResourceOption) (*RegionLoadBalancerLogSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.RegionName == nil {
		return nil, errors.New("invalid value for required argument 'RegionName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.StreamId == nil {
		return nil, errors.New("invalid value for required argument 'StreamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionLoadBalancerLogSubscription
	err := ctx.RegisterResource("ovh:CloudProject/regionLoadBalancerLogSubscription:RegionLoadBalancerLogSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionLoadBalancerLogSubscription gets an existing RegionLoadBalancerLogSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionLoadBalancerLogSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionLoadBalancerLogSubscriptionState, opts ...pulumi.ResourceOption) (*RegionLoadBalancerLogSubscription, error) {
	var resource RegionLoadBalancerLogSubscription
	err := ctx.ReadResource("ovh:CloudProject/regionLoadBalancerLogSubscription:RegionLoadBalancerLogSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionLoadBalancerLogSubscription resources.
type regionLoadBalancerLogSubscriptionState struct {
	// The date of the subscription creation
	CreatedAt *string `pulumi:"createdAt"`
	// haproxy  **Changing this value recreates the resource.**
	Kind *string `pulumi:"kind"`
	// LDP service name
	LdpServiceName *string `pulumi:"ldpServiceName"`
	// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// The operation ID
	OperationId *string `pulumi:"operationId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
	RegionName *string `pulumi:"regionName"`
	// The resource name
	ResourceName *string `pulumi:"resourceName"`
	// The resource type
	ResourceType *string `pulumi:"resourceType"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName *string `pulumi:"serviceName"`
	// Data stream id to use for the subscription  **Changing this value recreates the resource.**
	StreamId *string `pulumi:"streamId"`
	// The subscription id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The last update of the subscription
	UpdatedAt *string `pulumi:"updatedAt"`
}

type RegionLoadBalancerLogSubscriptionState struct {
	// The date of the subscription creation
	CreatedAt pulumi.StringPtrInput
	// haproxy  **Changing this value recreates the resource.**
	Kind pulumi.StringPtrInput
	// LDP service name
	LdpServiceName pulumi.StringPtrInput
	// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
	LoadbalancerId pulumi.StringPtrInput
	// The operation ID
	OperationId pulumi.StringPtrInput
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
	RegionName pulumi.StringPtrInput
	// The resource name
	ResourceName pulumi.StringPtrInput
	// The resource type
	ResourceType pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringPtrInput
	// Data stream id to use for the subscription  **Changing this value recreates the resource.**
	StreamId pulumi.StringPtrInput
	// The subscription id
	SubscriptionId pulumi.StringPtrInput
	// The last update of the subscription
	UpdatedAt pulumi.StringPtrInput
}

func (RegionLoadBalancerLogSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionLoadBalancerLogSubscriptionState)(nil)).Elem()
}

type regionLoadBalancerLogSubscriptionArgs struct {
	// haproxy  **Changing this value recreates the resource.**
	Kind string `pulumi:"kind"`
	// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
	RegionName string `pulumi:"regionName"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName string `pulumi:"serviceName"`
	// Data stream id to use for the subscription  **Changing this value recreates the resource.**
	StreamId string `pulumi:"streamId"`
}

// The set of arguments for constructing a RegionLoadBalancerLogSubscription resource.
type RegionLoadBalancerLogSubscriptionArgs struct {
	// haproxy  **Changing this value recreates the resource.**
	Kind pulumi.StringInput
	// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
	LoadbalancerId pulumi.StringInput
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
	RegionName pulumi.StringInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringInput
	// Data stream id to use for the subscription  **Changing this value recreates the resource.**
	StreamId pulumi.StringInput
}

func (RegionLoadBalancerLogSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionLoadBalancerLogSubscriptionArgs)(nil)).Elem()
}

type RegionLoadBalancerLogSubscriptionInput interface {
	pulumi.Input

	ToRegionLoadBalancerLogSubscriptionOutput() RegionLoadBalancerLogSubscriptionOutput
	ToRegionLoadBalancerLogSubscriptionOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionOutput
}

func (*RegionLoadBalancerLogSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (i *RegionLoadBalancerLogSubscription) ToRegionLoadBalancerLogSubscriptionOutput() RegionLoadBalancerLogSubscriptionOutput {
	return i.ToRegionLoadBalancerLogSubscriptionOutputWithContext(context.Background())
}

func (i *RegionLoadBalancerLogSubscription) ToRegionLoadBalancerLogSubscriptionOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionLoadBalancerLogSubscriptionOutput)
}

// RegionLoadBalancerLogSubscriptionArrayInput is an input type that accepts RegionLoadBalancerLogSubscriptionArray and RegionLoadBalancerLogSubscriptionArrayOutput values.
// You can construct a concrete instance of `RegionLoadBalancerLogSubscriptionArrayInput` via:
//
//	RegionLoadBalancerLogSubscriptionArray{ RegionLoadBalancerLogSubscriptionArgs{...} }
type RegionLoadBalancerLogSubscriptionArrayInput interface {
	pulumi.Input

	ToRegionLoadBalancerLogSubscriptionArrayOutput() RegionLoadBalancerLogSubscriptionArrayOutput
	ToRegionLoadBalancerLogSubscriptionArrayOutputWithContext(context.Context) RegionLoadBalancerLogSubscriptionArrayOutput
}

type RegionLoadBalancerLogSubscriptionArray []RegionLoadBalancerLogSubscriptionInput

func (RegionLoadBalancerLogSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (i RegionLoadBalancerLogSubscriptionArray) ToRegionLoadBalancerLogSubscriptionArrayOutput() RegionLoadBalancerLogSubscriptionArrayOutput {
	return i.ToRegionLoadBalancerLogSubscriptionArrayOutputWithContext(context.Background())
}

func (i RegionLoadBalancerLogSubscriptionArray) ToRegionLoadBalancerLogSubscriptionArrayOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionLoadBalancerLogSubscriptionArrayOutput)
}

// RegionLoadBalancerLogSubscriptionMapInput is an input type that accepts RegionLoadBalancerLogSubscriptionMap and RegionLoadBalancerLogSubscriptionMapOutput values.
// You can construct a concrete instance of `RegionLoadBalancerLogSubscriptionMapInput` via:
//
//	RegionLoadBalancerLogSubscriptionMap{ "key": RegionLoadBalancerLogSubscriptionArgs{...} }
type RegionLoadBalancerLogSubscriptionMapInput interface {
	pulumi.Input

	ToRegionLoadBalancerLogSubscriptionMapOutput() RegionLoadBalancerLogSubscriptionMapOutput
	ToRegionLoadBalancerLogSubscriptionMapOutputWithContext(context.Context) RegionLoadBalancerLogSubscriptionMapOutput
}

type RegionLoadBalancerLogSubscriptionMap map[string]RegionLoadBalancerLogSubscriptionInput

func (RegionLoadBalancerLogSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (i RegionLoadBalancerLogSubscriptionMap) ToRegionLoadBalancerLogSubscriptionMapOutput() RegionLoadBalancerLogSubscriptionMapOutput {
	return i.ToRegionLoadBalancerLogSubscriptionMapOutputWithContext(context.Background())
}

func (i RegionLoadBalancerLogSubscriptionMap) ToRegionLoadBalancerLogSubscriptionMapOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionLoadBalancerLogSubscriptionMapOutput)
}

type RegionLoadBalancerLogSubscriptionOutput struct{ *pulumi.OutputState }

func (RegionLoadBalancerLogSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (o RegionLoadBalancerLogSubscriptionOutput) ToRegionLoadBalancerLogSubscriptionOutput() RegionLoadBalancerLogSubscriptionOutput {
	return o
}

func (o RegionLoadBalancerLogSubscriptionOutput) ToRegionLoadBalancerLogSubscriptionOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionOutput {
	return o
}

// The date of the subscription creation
func (o RegionLoadBalancerLogSubscriptionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// haproxy  **Changing this value recreates the resource.**
func (o RegionLoadBalancerLogSubscriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// LDP service name
func (o RegionLoadBalancerLogSubscriptionOutput) LdpServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.LdpServiceName }).(pulumi.StringOutput)
}

// Loadbalancer id to get the logs  **Changing this value recreates the resource.**
func (o RegionLoadBalancerLogSubscriptionOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// The operation ID
func (o RegionLoadBalancerLogSubscriptionOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.OperationId }).(pulumi.StringOutput)
}

// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11". **Changing this value recreates the resource.**
func (o RegionLoadBalancerLogSubscriptionOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.RegionName }).(pulumi.StringOutput)
}

// The resource name
func (o RegionLoadBalancerLogSubscriptionOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// The resource type
func (o RegionLoadBalancerLogSubscriptionOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
func (o RegionLoadBalancerLogSubscriptionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Data stream id to use for the subscription  **Changing this value recreates the resource.**
func (o RegionLoadBalancerLogSubscriptionOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

// The subscription id
func (o RegionLoadBalancerLogSubscriptionOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The last update of the subscription
func (o RegionLoadBalancerLogSubscriptionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionLoadBalancerLogSubscription) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type RegionLoadBalancerLogSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (RegionLoadBalancerLogSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (o RegionLoadBalancerLogSubscriptionArrayOutput) ToRegionLoadBalancerLogSubscriptionArrayOutput() RegionLoadBalancerLogSubscriptionArrayOutput {
	return o
}

func (o RegionLoadBalancerLogSubscriptionArrayOutput) ToRegionLoadBalancerLogSubscriptionArrayOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionArrayOutput {
	return o
}

func (o RegionLoadBalancerLogSubscriptionArrayOutput) Index(i pulumi.IntInput) RegionLoadBalancerLogSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionLoadBalancerLogSubscription {
		return vs[0].([]*RegionLoadBalancerLogSubscription)[vs[1].(int)]
	}).(RegionLoadBalancerLogSubscriptionOutput)
}

type RegionLoadBalancerLogSubscriptionMapOutput struct{ *pulumi.OutputState }

func (RegionLoadBalancerLogSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionLoadBalancerLogSubscription)(nil)).Elem()
}

func (o RegionLoadBalancerLogSubscriptionMapOutput) ToRegionLoadBalancerLogSubscriptionMapOutput() RegionLoadBalancerLogSubscriptionMapOutput {
	return o
}

func (o RegionLoadBalancerLogSubscriptionMapOutput) ToRegionLoadBalancerLogSubscriptionMapOutputWithContext(ctx context.Context) RegionLoadBalancerLogSubscriptionMapOutput {
	return o
}

func (o RegionLoadBalancerLogSubscriptionMapOutput) MapIndex(k pulumi.StringInput) RegionLoadBalancerLogSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionLoadBalancerLogSubscription {
		return vs[0].(map[string]*RegionLoadBalancerLogSubscription)[vs[1].(string)]
	}).(RegionLoadBalancerLogSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionLoadBalancerLogSubscriptionInput)(nil)).Elem(), &RegionLoadBalancerLogSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionLoadBalancerLogSubscriptionArrayInput)(nil)).Elem(), RegionLoadBalancerLogSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionLoadBalancerLogSubscriptionMapInput)(nil)).Elem(), RegionLoadBalancerLogSubscriptionMap{})
	pulumi.RegisterOutputType(RegionLoadBalancerLogSubscriptionOutput{})
	pulumi.RegisterOutputType(RegionLoadBalancerLogSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(RegionLoadBalancerLogSubscriptionMapOutput{})
}
