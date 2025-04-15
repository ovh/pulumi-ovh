// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a subscription to a Managed Loadbalancer Logs Service in a public cloud project.
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
//			_, err := cloudproject.GetRegionLoadBalancerLogSubscription(ctx, &cloudproject.GetRegionLoadBalancerLogSubscriptionArgs{
//				ServiceName:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//				RegionName:     "gggg",
//				LoadbalancerId: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//				SubscriptionId: "zzzzzzzz-yyyy-xxxx-wwww-vvvvvvvvvvvv",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRegionLoadBalancerLogSubscription(ctx *pulumi.Context, args *LookupRegionLoadBalancerLogSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupRegionLoadBalancerLogSubscriptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegionLoadBalancerLogSubscriptionResult
	err := ctx.Invoke("ovh:CloudProject/getRegionLoadBalancerLogSubscription:getRegionLoadBalancerLogSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionLoadBalancerLogSubscription.
type LookupRegionLoadBalancerLogSubscriptionArgs struct {
	// Loadbalancer id to get the logs
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer is available. Ex.: "GRA11".
	RegionName string `pulumi:"regionName"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Subscription id
	SubscriptionId string `pulumi:"subscriptionId"`
}

// A collection of values returned by getRegionLoadBalancerLogSubscription.
type LookupRegionLoadBalancerLogSubscriptionResult struct {
	// The date of the subscription creation
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Router used for forwarding log
	Kind string `pulumi:"kind"`
	// LDP service name
	LdpServiceName string `pulumi:"ldpServiceName"`
	// Loadbalancer id to get the logs
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11".
	RegionName string `pulumi:"regionName"`
	// The resource name
	ResourceName string `pulumi:"resourceName"`
	// The resource type
	ResourceType string `pulumi:"resourceType"`
	// The id of the public cloud project.
	ServiceName string `pulumi:"serviceName"`
	// Data stream id to use for the subscription
	StreamId string `pulumi:"streamId"`
	// The subscription id
	SubscriptionId string `pulumi:"subscriptionId"`
	// The last update of the subscription
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupRegionLoadBalancerLogSubscriptionOutput(ctx *pulumi.Context, args LookupRegionLoadBalancerLogSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupRegionLoadBalancerLogSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRegionLoadBalancerLogSubscriptionResultOutput, error) {
			args := v.(LookupRegionLoadBalancerLogSubscriptionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getRegionLoadBalancerLogSubscription:getRegionLoadBalancerLogSubscription", args, LookupRegionLoadBalancerLogSubscriptionResultOutput{}, options).(LookupRegionLoadBalancerLogSubscriptionResultOutput), nil
		}).(LookupRegionLoadBalancerLogSubscriptionResultOutput)
}

// A collection of arguments for invoking getRegionLoadBalancerLogSubscription.
type LookupRegionLoadBalancerLogSubscriptionOutputArgs struct {
	// Loadbalancer id to get the logs
	LoadbalancerId pulumi.StringInput `pulumi:"loadbalancerId"`
	// A valid OVHcloud public cloud region name in which the loadbalancer is available. Ex.: "GRA11".
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Subscription id
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
}

func (LookupRegionLoadBalancerLogSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionLoadBalancerLogSubscriptionArgs)(nil)).Elem()
}

// A collection of values returned by getRegionLoadBalancerLogSubscription.
type LookupRegionLoadBalancerLogSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupRegionLoadBalancerLogSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionLoadBalancerLogSubscriptionResult)(nil)).Elem()
}

func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) ToLookupRegionLoadBalancerLogSubscriptionResultOutput() LookupRegionLoadBalancerLogSubscriptionResultOutput {
	return o
}

func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) ToLookupRegionLoadBalancerLogSubscriptionResultOutputWithContext(ctx context.Context) LookupRegionLoadBalancerLogSubscriptionResultOutput {
	return o
}

// The date of the subscription creation
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Router used for forwarding log
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// LDP service name
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) LdpServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.LdpServiceName }).(pulumi.StringOutput)
}

// Loadbalancer id to get the logs
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11".
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// The resource name
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

// The resource type
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

// The id of the public cloud project.
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Data stream id to use for the subscription
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.StreamId }).(pulumi.StringOutput)
}

// The subscription id
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// The last update of the subscription
func (o LookupRegionLoadBalancerLogSubscriptionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionLoadBalancerLogSubscriptionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionLoadBalancerLogSubscriptionResultOutput{})
}
