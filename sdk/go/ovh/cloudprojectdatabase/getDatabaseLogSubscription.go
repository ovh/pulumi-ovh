// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a log subscription for a cluster associated with a public cloud project.
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
//			subscription, err := CloudProjectDatabase.GetDatabaseLogSubscription(ctx, &cloudprojectdatabase.GetDatabaseLogSubscriptionArgs{
//				ServiceName: "VVV",
//				Engine:      "XXX",
//				ClusterId:   "YYY",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("subscriptionLdpName", subscription.LdpServiceName)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseLogSubscription(ctx *pulumi.Context, args *GetDatabaseLogSubscriptionArgs, opts ...pulumi.InvokeOption) (*GetDatabaseLogSubscriptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabaseLogSubscriptionResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseLogSubscription.
type GetDatabaseLogSubscriptionArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// Id of the log subscription.
	Id string `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabaseLogSubscription.
type GetDatabaseLogSubscriptionResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Creation date of the subscription.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// ID of the log subscription.
	Id string `pulumi:"id"`
	// Log kind name of this subscription.
	Kind string `pulumi:"kind"`
	// Name of the destination log service.
	LdpServiceName string `pulumi:"ldpServiceName"`
	// Name of subscribed resource, where the logs come from.
	ResourceName string `pulumi:"resourceName"`
	// Type of subscribed resource, where the logs come from.
	ResourceType string `pulumi:"resourceType"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// See Argument Reference above.
	StreamId string `pulumi:"streamId"`
	// Last update date of the subscription.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetDatabaseLogSubscriptionOutput(ctx *pulumi.Context, args GetDatabaseLogSubscriptionOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseLogSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseLogSubscriptionResultOutput, error) {
			args := v.(GetDatabaseLogSubscriptionArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDatabaseLogSubscriptionResult
			secret, err := ctx.InvokePackageRaw("ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription", args, &rv, "", opts...)
			if err != nil {
				return GetDatabaseLogSubscriptionResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDatabaseLogSubscriptionResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDatabaseLogSubscriptionResultOutput), nil
			}
			return output, nil
		}).(GetDatabaseLogSubscriptionResultOutput)
}

// A collection of arguments for invoking getDatabaseLogSubscription.
type GetDatabaseLogSubscriptionOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// Id of the log subscription.
	Id pulumi.StringInput `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabaseLogSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseLogSubscriptionArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseLogSubscription.
type GetDatabaseLogSubscriptionResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseLogSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseLogSubscriptionResult)(nil)).Elem()
}

func (o GetDatabaseLogSubscriptionResultOutput) ToGetDatabaseLogSubscriptionResultOutput() GetDatabaseLogSubscriptionResultOutput {
	return o
}

func (o GetDatabaseLogSubscriptionResultOutput) ToGetDatabaseLogSubscriptionResultOutputWithContext(ctx context.Context) GetDatabaseLogSubscriptionResultOutput {
	return o
}

// See Argument Reference above.
func (o GetDatabaseLogSubscriptionResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation date of the subscription.
func (o GetDatabaseLogSubscriptionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseLogSubscriptionResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.Engine }).(pulumi.StringOutput)
}

// ID of the log subscription.
func (o GetDatabaseLogSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Log kind name of this subscription.
func (o GetDatabaseLogSubscriptionResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the destination log service.
func (o GetDatabaseLogSubscriptionResultOutput) LdpServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.LdpServiceName }).(pulumi.StringOutput)
}

// Name of subscribed resource, where the logs come from.
func (o GetDatabaseLogSubscriptionResultOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.ResourceName }).(pulumi.StringOutput)
}

// Type of subscribed resource, where the logs come from.
func (o GetDatabaseLogSubscriptionResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseLogSubscriptionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseLogSubscriptionResultOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.StreamId }).(pulumi.StringOutput)
}

// Last update date of the subscription.
func (o GetDatabaseLogSubscriptionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseLogSubscriptionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseLogSubscriptionResultOutput{})
}
