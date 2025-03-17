// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about user acces of a kafka cluster associated with a public cloud project.
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
//			access, err := cloudprojectdatabase.GetKafkaUserAccess(ctx, &cloudprojectdatabase.GetKafkaUserAccessArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				UserId:      "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("accessCert", access.Cert)
//			return nil
//		})
//	}
//
// ```
func GetKafkaUserAccess(ctx *pulumi.Context, args *GetKafkaUserAccessArgs, opts ...pulumi.InvokeOption) (*GetKafkaUserAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKafkaUserAccessResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getKafkaUserAccess:getKafkaUserAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKafkaUserAccess.
type GetKafkaUserAccessArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// User ID
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getKafkaUserAccess.
type GetKafkaUserAccessResult struct {
	// User cert.
	Cert string `pulumi:"cert"`
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Sensitive) User key for the cert.
	Key string `pulumi:"key"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// See Argument Reference above.
	UserId string `pulumi:"userId"`
}

func GetKafkaUserAccessOutput(ctx *pulumi.Context, args GetKafkaUserAccessOutputArgs, opts ...pulumi.InvokeOption) GetKafkaUserAccessResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKafkaUserAccessResultOutput, error) {
			args := v.(GetKafkaUserAccessArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getKafkaUserAccess:getKafkaUserAccess", args, GetKafkaUserAccessResultOutput{}, options).(GetKafkaUserAccessResultOutput), nil
		}).(GetKafkaUserAccessResultOutput)
}

// A collection of arguments for invoking getKafkaUserAccess.
type GetKafkaUserAccessOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// User ID
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (GetKafkaUserAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKafkaUserAccessArgs)(nil)).Elem()
}

// A collection of values returned by getKafkaUserAccess.
type GetKafkaUserAccessResultOutput struct{ *pulumi.OutputState }

func (GetKafkaUserAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKafkaUserAccessResult)(nil)).Elem()
}

func (o GetKafkaUserAccessResultOutput) ToGetKafkaUserAccessResultOutput() GetKafkaUserAccessResultOutput {
	return o
}

func (o GetKafkaUserAccessResultOutput) ToGetKafkaUserAccessResultOutputWithContext(ctx context.Context) GetKafkaUserAccessResultOutput {
	return o
}

// User cert.
func (o GetKafkaUserAccessResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.Cert }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKafkaUserAccessResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKafkaUserAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Sensitive) User key for the cert.
func (o GetKafkaUserAccessResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.Key }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKafkaUserAccessResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKafkaUserAccessResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKafkaUserAccessResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKafkaUserAccessResultOutput{})
}
