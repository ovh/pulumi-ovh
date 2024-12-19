// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the container registries of a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.GetContainerRegistries(ctx, &cloudproject.GetContainerRegistriesArgs{
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupContainerRegistries(ctx *pulumi.Context, args *LookupContainerRegistriesArgs, opts ...pulumi.InvokeOption) (*LookupContainerRegistriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerRegistriesResult
	err := ctx.Invoke("ovh:CloudProject/getContainerRegistries:getContainerRegistries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistries.
type LookupContainerRegistriesArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getContainerRegistries.
type LookupContainerRegistriesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of container registries associated with the project.
	Results     []GetContainerRegistriesResult `pulumi:"results"`
	ServiceName string                         `pulumi:"serviceName"`
}

func LookupContainerRegistriesOutput(ctx *pulumi.Context, args LookupContainerRegistriesOutputArgs, opts ...pulumi.InvokeOption) LookupContainerRegistriesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContainerRegistriesResultOutput, error) {
			args := v.(LookupContainerRegistriesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getContainerRegistries:getContainerRegistries", args, LookupContainerRegistriesResultOutput{}, options).(LookupContainerRegistriesResultOutput), nil
		}).(LookupContainerRegistriesResultOutput)
}

// A collection of arguments for invoking getContainerRegistries.
type LookupContainerRegistriesOutputArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContainerRegistriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistriesArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistries.
type LookupContainerRegistriesResultOutput struct{ *pulumi.OutputState }

func (LookupContainerRegistriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistriesResult)(nil)).Elem()
}

func (o LookupContainerRegistriesResultOutput) ToLookupContainerRegistriesResultOutput() LookupContainerRegistriesResultOutput {
	return o
}

func (o LookupContainerRegistriesResultOutput) ToLookupContainerRegistriesResultOutputWithContext(ctx context.Context) LookupContainerRegistriesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerRegistriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistriesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of container registries associated with the project.
func (o LookupContainerRegistriesResultOutput) Results() GetContainerRegistriesResultArrayOutput {
	return o.ApplyT(func(v LookupContainerRegistriesResult) []GetContainerRegistriesResult { return v.Results }).(GetContainerRegistriesResultArrayOutput)
}

func (o LookupContainerRegistriesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistriesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerRegistriesResultOutput{})
}
