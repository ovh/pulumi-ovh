// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List your public cloud loadbalancers.
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
//			lbsLoadBalancers, err := CloudProject.GetLoadBalancers(ctx, &cloudproject.GetLoadBalancersArgs{
//				ServiceName: "XXXXXX",
//				RegionName:  "XXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("lbs", lbsLoadBalancers)
//			return nil
//		})
//	}
//
// ```
func GetLoadBalancers(ctx *pulumi.Context, args *GetLoadBalancersArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLoadBalancersResult
	err := ctx.Invoke("ovh:CloudProject/getLoadBalancers:getLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersArgs struct {
	// Region of the loadbalancers.
	RegionName string `pulumi:"regionName"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of loadbalancer
	Loadbalancers []GetLoadBalancersLoadbalancer `pulumi:"loadbalancers"`
	// Region of the loadbalancers
	RegionName string `pulumi:"regionName"`
	// ID of the public cloud project
	ServiceName string `pulumi:"serviceName"`
}

func GetLoadBalancersOutput(ctx *pulumi.Context, args GetLoadBalancersOutputArgs, opts ...pulumi.InvokeOption) GetLoadBalancersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLoadBalancersResultOutput, error) {
			args := v.(GetLoadBalancersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getLoadBalancers:getLoadBalancers", args, GetLoadBalancersResultOutput{}, options).(GetLoadBalancersResultOutput), nil
		}).(GetLoadBalancersResultOutput)
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersOutputArgs struct {
	// Region of the loadbalancers.
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetLoadBalancersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancersArgs)(nil)).Elem()
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResultOutput struct{ *pulumi.OutputState }

func (GetLoadBalancersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancersResult)(nil)).Elem()
}

func (o GetLoadBalancersResultOutput) ToGetLoadBalancersResultOutput() GetLoadBalancersResultOutput {
	return o
}

func (o GetLoadBalancersResultOutput) ToGetLoadBalancersResultOutputWithContext(ctx context.Context) GetLoadBalancersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadBalancersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of loadbalancer
func (o GetLoadBalancersResultOutput) Loadbalancers() GetLoadBalancersLoadbalancerArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []GetLoadBalancersLoadbalancer { return v.Loadbalancers }).(GetLoadBalancersLoadbalancerArrayOutput)
}

// Region of the loadbalancers
func (o GetLoadBalancersResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// ID of the public cloud project
func (o GetLoadBalancersResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadBalancersResultOutput{})
}
