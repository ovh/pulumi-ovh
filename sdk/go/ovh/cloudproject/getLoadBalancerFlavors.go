// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List loadbalancer flavors in the given public cloud region.
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
//			_, err := cloudproject.GetLoadBalancerFlavors(ctx, &cloudproject.GetLoadBalancerFlavorsArgs{
//				ServiceName: "<public cloud project ID>",
//				RegionName:  "GRA9",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLoadBalancerFlavors(ctx *pulumi.Context, args *GetLoadBalancerFlavorsArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancerFlavorsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLoadBalancerFlavorsResult
	err := ctx.Invoke("ovh:CloudProject/getLoadBalancerFlavors:getLoadBalancerFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancerFlavors.
type GetLoadBalancerFlavorsArgs struct {
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getLoadBalancerFlavors.
type GetLoadBalancerFlavorsResult struct {
	Flavors []GetLoadBalancerFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

func GetLoadBalancerFlavorsOutput(ctx *pulumi.Context, args GetLoadBalancerFlavorsOutputArgs, opts ...pulumi.InvokeOption) GetLoadBalancerFlavorsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLoadBalancerFlavorsResultOutput, error) {
			args := v.(GetLoadBalancerFlavorsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getLoadBalancerFlavors:getLoadBalancerFlavors", args, GetLoadBalancerFlavorsResultOutput{}, options).(GetLoadBalancerFlavorsResultOutput), nil
		}).(GetLoadBalancerFlavorsResultOutput)
}

// A collection of arguments for invoking getLoadBalancerFlavors.
type GetLoadBalancerFlavorsOutputArgs struct {
	// Region name
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// Service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetLoadBalancerFlavorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancerFlavorsArgs)(nil)).Elem()
}

// A collection of values returned by getLoadBalancerFlavors.
type GetLoadBalancerFlavorsResultOutput struct{ *pulumi.OutputState }

func (GetLoadBalancerFlavorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancerFlavorsResult)(nil)).Elem()
}

func (o GetLoadBalancerFlavorsResultOutput) ToGetLoadBalancerFlavorsResultOutput() GetLoadBalancerFlavorsResultOutput {
	return o
}

func (o GetLoadBalancerFlavorsResultOutput) ToGetLoadBalancerFlavorsResultOutputWithContext(ctx context.Context) GetLoadBalancerFlavorsResultOutput {
	return o
}

func (o GetLoadBalancerFlavorsResultOutput) Flavors() GetLoadBalancerFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetLoadBalancerFlavorsResult) []GetLoadBalancerFlavorsFlavor { return v.Flavors }).(GetLoadBalancerFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadBalancerFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Region name
func (o GetLoadBalancerFlavorsResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerFlavorsResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// Service name
func (o GetLoadBalancerFlavorsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerFlavorsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadBalancerFlavorsResultOutput{})
}
