// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of topics of a kafka cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloudprojectdatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			topics, err := cloudprojectdatabase.GetKafkaTopics(ctx, &cloudprojectdatabase.GetKafkaTopicsArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("topicIds", topics.TopicIds)
//			return nil
//		})
//	}
//
// ```
func GetKafkaTopics(ctx *pulumi.Context, args *GetKafkaTopicsArgs, opts ...pulumi.InvokeOption) (*GetKafkaTopicsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKafkaTopicsResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getKafkaTopics:getKafkaTopics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKafkaTopics.
type GetKafkaTopicsArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKafkaTopics.
type GetKafkaTopicsResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// The list of topics ids of the kafka cluster associated with the project.
	TopicIds []string `pulumi:"topicIds"`
}

func GetKafkaTopicsOutput(ctx *pulumi.Context, args GetKafkaTopicsOutputArgs, opts ...pulumi.InvokeOption) GetKafkaTopicsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKafkaTopicsResultOutput, error) {
			args := v.(GetKafkaTopicsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getKafkaTopics:getKafkaTopics", args, GetKafkaTopicsResultOutput{}, options).(GetKafkaTopicsResultOutput), nil
		}).(GetKafkaTopicsResultOutput)
}

// A collection of arguments for invoking getKafkaTopics.
type GetKafkaTopicsOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetKafkaTopicsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKafkaTopicsArgs)(nil)).Elem()
}

// A collection of values returned by getKafkaTopics.
type GetKafkaTopicsResultOutput struct{ *pulumi.OutputState }

func (GetKafkaTopicsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKafkaTopicsResult)(nil)).Elem()
}

func (o GetKafkaTopicsResultOutput) ToGetKafkaTopicsResultOutput() GetKafkaTopicsResultOutput {
	return o
}

func (o GetKafkaTopicsResultOutput) ToGetKafkaTopicsResultOutputWithContext(ctx context.Context) GetKafkaTopicsResultOutput {
	return o
}

// See Argument Reference above.
func (o GetKafkaTopicsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaTopicsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKafkaTopicsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaTopicsResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKafkaTopicsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaTopicsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The list of topics ids of the kafka cluster associated with the project.
func (o GetKafkaTopicsResultOutput) TopicIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKafkaTopicsResult) []string { return v.TopicIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKafkaTopicsResultOutput{})
}
