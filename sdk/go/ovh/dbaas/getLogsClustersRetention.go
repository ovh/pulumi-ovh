// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	ClusterId     string  `pulumi:"clusterId"`
	Duration      *string `pulumi:"duration"`
	RetentionId   *string `pulumi:"retentionId"`
	RetentionType *string `pulumi:"retentionType"`
	ServiceName   string  `pulumi:"serviceName"`
}

// A collection of values returned by getLogsClustersRetention.
type GetLogsClustersRetentionResult struct {
	ClusterId string `pulumi:"clusterId"`
	Duration  string `pulumi:"duration"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IsSupported   bool   `pulumi:"isSupported"`
	RetentionId   string `pulumi:"retentionId"`
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
	ClusterId     pulumi.StringInput    `pulumi:"clusterId"`
	Duration      pulumi.StringPtrInput `pulumi:"duration"`
	RetentionId   pulumi.StringPtrInput `pulumi:"retentionId"`
	RetentionType pulumi.StringPtrInput `pulumi:"retentionType"`
	ServiceName   pulumi.StringInput    `pulumi:"serviceName"`
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

func (o GetLogsClustersRetentionResultOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.Duration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogsClustersRetentionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogsClustersRetentionResultOutput) IsSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) bool { return v.IsSupported }).(pulumi.BoolOutput)
}

func (o GetLogsClustersRetentionResultOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.RetentionId }).(pulumi.StringOutput)
}

func (o GetLogsClustersRetentionResultOutput) RetentionType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.RetentionType }).(pulumi.StringOutput)
}

func (o GetLogsClustersRetentionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsClustersRetentionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogsClustersRetentionResultOutput{})
}
