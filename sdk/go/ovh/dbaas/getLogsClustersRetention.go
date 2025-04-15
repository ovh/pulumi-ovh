// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a DBaas logs cluster retention.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.GetLogsClustersRetention(ctx, &dbaas.GetLogsClustersRetentionArgs{
//				ServiceName: "ldp-xx-xxxxx",
//				ClusterId:   "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//				RetentionId: pulumi.StringRef("yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// It is also possible to retrieve a retention using its duration:
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.GetLogsClustersRetention(ctx, &dbaas.GetLogsClustersRetentionArgs{
//				ServiceName: "ldp-xx-xxxxx",
//				ClusterId:   "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//				Duration:    pulumi.StringRef("P14D"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Additionnaly, you can filter retentions on their type:
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.GetLogsClustersRetention(ctx, &dbaas.GetLogsClustersRetentionArgs{
//				ServiceName:   "ldp-xx-xxxxx",
//				ClusterId:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//				Duration:      pulumi.StringRef("P14D"),
//				RetentionType: pulumi.StringRef("LOGS_INDEXING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLogsClustersRetention(ctx *pulumi.Context, args *GetLogsClustersRetentionArgs, opts ...pulumi.InvokeOption) (*GetLogsClustersRetentionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogsClustersRetentionResult
	err := ctx.Invoke("ovh:Dbaas/getLogsClustersRetention:getLogsClustersRetention", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogsClustersRetention.
type GetLogsClustersRetentionArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Indexed duration expressed in ISO-8601 format. Cannot be used if `retentionId` is defined.
	Duration *string `pulumi:"duration"`
	// ID of the retention object. Cannot be used if `duration` or `retentionType` is defined.
	RetentionId *string `pulumi:"retentionId"`
	// Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT). Cannot be used if `retentionId` is defined. Defaults to `LOGS_INDEXING` if not defined.
	RetentionType *string `pulumi:"retentionType"`
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getLogsClustersRetention.
type GetLogsClustersRetentionResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Indexed duration expressed in ISO-8601 format
	Duration string `pulumi:"duration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates if a new stream can use it
	IsSupported bool `pulumi:"isSupported"`
	// ID of the retention that can be used when creating a stream
	RetentionId string `pulumi:"retentionId"`
	// Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT)
	RetentionType string `pulumi:"retentionType"`
	ServiceName   string `pulumi:"serviceName"`
}

func GetLogsClustersRetentionOutput(ctx *pulumi.Context, args GetLogsClustersRetentionOutputArgs, opts ...pulumi.InvokeOption) GetLogsClustersRetentionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLogsClustersRetentionResultOutput, error) {
			args := v.(GetLogsClustersRetentionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dbaas/getLogsClustersRetention:getLogsClustersRetention", args, GetLogsClustersRetentionResultOutput{}, options).(GetLogsClustersRetentionResultOutput), nil
		}).(GetLogsClustersRetentionResultOutput)
}

// A collection of arguments for invoking getLogsClustersRetention.
type GetLogsClustersRetentionOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Indexed duration expressed in ISO-8601 format. Cannot be used if `retentionId` is defined.
	Duration pulumi.StringPtrInput `pulumi:"duration"`
	// ID of the retention object. Cannot be used if `duration` or `retentionType` is defined.
	RetentionId pulumi.StringPtrInput `pulumi:"retentionId"`
	// Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT). Cannot be used if `retentionId` is defined. Defaults to `LOGS_INDEXING` if not defined.
	RetentionType pulumi.StringPtrInput `pulumi:"retentionType"`
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetLogsClustersRetentionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsClustersRetentionArgs)(nil)).Elem()
}

// A collection of values returned by getLogsClustersRetention.
type GetLogsClustersRetentionResultOutput struct{ *pulumi.OutputState }

func (GetLogsClustersRetentionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsClustersRetentionResult)(nil)).Elem()
}

func (o GetLogsClustersRetentionResultOutput) ToGetLogsClustersRetentionResultOutput() GetLogsClustersRetentionResultOutput {
	return o
}

func (o GetLogsClustersRetentionResultOutput) ToGetLogsClustersRetentionResultOutputWithContext(ctx context.Context) GetLogsClustersRetentionResultOutput {
	return o
}

func (o GetLogsClustersRetentionResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Indexed duration expressed in ISO-8601 format
func (o GetLogsClustersRetentionResultOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.Duration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogsClustersRetentionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates if a new stream can use it
func (o GetLogsClustersRetentionResultOutput) IsSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) bool { return v.IsSupported }).(pulumi.BoolOutput)
}

// ID of the retention that can be used when creating a stream
func (o GetLogsClustersRetentionResultOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.RetentionId }).(pulumi.StringOutput)
}

// Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT)
func (o GetLogsClustersRetentionResultOutput) RetentionType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.RetentionType }).(pulumi.StringOutput)
}

func (o GetLogsClustersRetentionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogsClustersRetentionResultOutput{})
}
