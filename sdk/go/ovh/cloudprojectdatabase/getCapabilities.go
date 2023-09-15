// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get information about capabilities of a public cloud project.
func GetCapabilities(ctx *pulumi.Context, args *GetCapabilitiesArgs, opts ...pulumi.InvokeOption) (*GetCapabilitiesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCapabilitiesResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getCapabilities:getCapabilities", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCapabilities.
type GetCapabilitiesArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCapabilities.
type GetCapabilitiesResult struct {
	// Database engines available.
	Engines []GetCapabilitiesEngine `pulumi:"engines"`
	// Flavors available.
	Flavors []GetCapabilitiesFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Options available.
	Options []GetCapabilitiesOption `pulumi:"options"`
	// Plans available.
	Plans []GetCapabilitiesPlan `pulumi:"plans"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetCapabilitiesOutput(ctx *pulumi.Context, args GetCapabilitiesOutputArgs, opts ...pulumi.InvokeOption) GetCapabilitiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCapabilitiesResult, error) {
			args := v.(GetCapabilitiesArgs)
			r, err := GetCapabilities(ctx, &args, opts...)
			var s GetCapabilitiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCapabilitiesResultOutput)
}

// A collection of arguments for invoking getCapabilities.
type GetCapabilitiesOutputArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCapabilitiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesArgs)(nil)).Elem()
}

// A collection of values returned by getCapabilities.
type GetCapabilitiesResultOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesResult)(nil)).Elem()
}

func (o GetCapabilitiesResultOutput) ToGetCapabilitiesResultOutput() GetCapabilitiesResultOutput {
	return o
}

func (o GetCapabilitiesResultOutput) ToGetCapabilitiesResultOutputWithContext(ctx context.Context) GetCapabilitiesResultOutput {
	return o
}

func (o GetCapabilitiesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetCapabilitiesResult] {
	return pulumix.Output[GetCapabilitiesResult]{
		OutputState: o.OutputState,
	}
}

// Database engines available.
func (o GetCapabilitiesResultOutput) Engines() GetCapabilitiesEngineArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) []GetCapabilitiesEngine { return v.Engines }).(GetCapabilitiesEngineArrayOutput)
}

// Flavors available.
func (o GetCapabilitiesResultOutput) Flavors() GetCapabilitiesFlavorArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) []GetCapabilitiesFlavor { return v.Flavors }).(GetCapabilitiesFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCapabilitiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Options available.
func (o GetCapabilitiesResultOutput) Options() GetCapabilitiesOptionArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) []GetCapabilitiesOption { return v.Options }).(GetCapabilitiesOptionArrayOutput)
}

// Plans available.
func (o GetCapabilitiesResultOutput) Plans() GetCapabilitiesPlanArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) []GetCapabilitiesPlan { return v.Plans }).(GetCapabilitiesPlanArrayOutput)
}

// See Argument Reference above.
func (o GetCapabilitiesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCapabilitiesResultOutput{})
}
