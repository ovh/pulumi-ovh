// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List your S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
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
//			_, err := cloudproject.GetStorage(ctx, &cloudproject.GetStorageArgs{
//				ServiceName: "<public cloud project ID>",
//				RegionName:  "GRA",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetStorages(ctx *pulumi.Context, args *GetStoragesArgs, opts ...pulumi.InvokeOption) (*GetStoragesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStoragesResult
	err := ctx.Invoke("ovh:CloudProject/getStorages:getStorages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorages.
type GetStoragesArgs struct {
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getStorages.
type GetStoragesResult struct {
	Containers []GetStoragesContainer `pulumi:"containers"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

func GetStoragesOutput(ctx *pulumi.Context, args GetStoragesOutputArgs, opts ...pulumi.InvokeOption) GetStoragesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetStoragesResultOutput, error) {
			args := v.(GetStoragesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getStorages:getStorages", args, GetStoragesResultOutput{}, options).(GetStoragesResultOutput), nil
		}).(GetStoragesResultOutput)
}

// A collection of arguments for invoking getStorages.
type GetStoragesOutputArgs struct {
	// Region name
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// Service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetStoragesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesArgs)(nil)).Elem()
}

// A collection of values returned by getStorages.
type GetStoragesResultOutput struct{ *pulumi.OutputState }

func (GetStoragesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesResult)(nil)).Elem()
}

func (o GetStoragesResultOutput) ToGetStoragesResultOutput() GetStoragesResultOutput {
	return o
}

func (o GetStoragesResultOutput) ToGetStoragesResultOutputWithContext(ctx context.Context) GetStoragesResultOutput {
	return o
}

func (o GetStoragesResultOutput) Containers() GetStoragesContainerArrayOutput {
	return o.ApplyT(func(v GetStoragesResult) []GetStoragesContainer { return v.Containers }).(GetStoragesContainerArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStoragesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Region name
func (o GetStoragesResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// Service name
func (o GetStoragesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStoragesResultOutput{})
}
