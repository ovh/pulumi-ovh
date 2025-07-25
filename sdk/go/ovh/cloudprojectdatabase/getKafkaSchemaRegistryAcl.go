// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a schema registry ACL of a kafka cluster associated with a public cloud project.
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
//			schemaRegistryAcl, err := cloudprojectdatabase.GetKafkaSchemaRegistryAcl(ctx, &cloudprojectdatabase.GetKafkaSchemaRegistryAclArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("aclPermission", schemaRegistryAcl.Permission)
//			return nil
//		})
//	}
//
// ```
func LookupKafkaSchemaRegistryAcl(ctx *pulumi.Context, args *LookupKafkaSchemaRegistryAclArgs, opts ...pulumi.InvokeOption) (*LookupKafkaSchemaRegistryAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKafkaSchemaRegistryAclResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKafkaSchemaRegistryAcl.
type LookupKafkaSchemaRegistryAclArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Schema registry ACL ID
	Id string `pulumi:"id"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKafkaSchemaRegistryAcl.
type LookupKafkaSchemaRegistryAclResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Id string `pulumi:"id"`
	// Permission to give to this username on this topic.
	Permission string `pulumi:"permission"`
	// Resource affected by this ACL.
	Resource string `pulumi:"resource"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Username affected by this ACL.
	Username string `pulumi:"username"`
}

func LookupKafkaSchemaRegistryAclOutput(ctx *pulumi.Context, args LookupKafkaSchemaRegistryAclOutputArgs, opts ...pulumi.InvokeOption) LookupKafkaSchemaRegistryAclResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKafkaSchemaRegistryAclResultOutput, error) {
			args := v.(LookupKafkaSchemaRegistryAclArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", args, LookupKafkaSchemaRegistryAclResultOutput{}, options).(LookupKafkaSchemaRegistryAclResultOutput), nil
		}).(LookupKafkaSchemaRegistryAclResultOutput)
}

// A collection of arguments for invoking getKafkaSchemaRegistryAcl.
type LookupKafkaSchemaRegistryAclOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Schema registry ACL ID
	Id pulumi.StringInput `pulumi:"id"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupKafkaSchemaRegistryAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaSchemaRegistryAclArgs)(nil)).Elem()
}

// A collection of values returned by getKafkaSchemaRegistryAcl.
type LookupKafkaSchemaRegistryAclResultOutput struct{ *pulumi.OutputState }

func (LookupKafkaSchemaRegistryAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaSchemaRegistryAclResult)(nil)).Elem()
}

func (o LookupKafkaSchemaRegistryAclResultOutput) ToLookupKafkaSchemaRegistryAclResultOutput() LookupKafkaSchemaRegistryAclResultOutput {
	return o
}

func (o LookupKafkaSchemaRegistryAclResultOutput) ToLookupKafkaSchemaRegistryAclResultOutputWithContext(ctx context.Context) LookupKafkaSchemaRegistryAclResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupKafkaSchemaRegistryAclResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKafkaSchemaRegistryAclResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.Id }).(pulumi.StringOutput)
}

// Permission to give to this username on this topic.
func (o LookupKafkaSchemaRegistryAclResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.Permission }).(pulumi.StringOutput)
}

// Resource affected by this ACL.
func (o LookupKafkaSchemaRegistryAclResultOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.Resource }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKafkaSchemaRegistryAclResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Username affected by this ACL.
func (o LookupKafkaSchemaRegistryAclResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaSchemaRegistryAclResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKafkaSchemaRegistryAclResultOutput{})
}
