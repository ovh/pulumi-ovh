// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of Registry IP Restrictions of a container registry associated with a public cloud project.
func LookupContainerRegistryIPRestrictionsRegistry(ctx *pulumi.Context, args *LookupContainerRegistryIPRestrictionsRegistryArgs, opts ...pulumi.InvokeOption) (*LookupContainerRegistryIPRestrictionsRegistryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerRegistryIPRestrictionsRegistryResult
	err := ctx.Invoke("ovh:CloudProject/getContainerRegistryIPRestrictionsRegistry:getContainerRegistryIPRestrictionsRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistryIPRestrictionsRegistry.
type LookupContainerRegistryIPRestrictionsRegistryArgs struct {
	// The id of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getContainerRegistryIPRestrictionsRegistry.
type LookupContainerRegistryIPRestrictionsRegistryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP restrictions applied on artifact manager component.
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// The ID of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

func LookupContainerRegistryIPRestrictionsRegistryOutput(ctx *pulumi.Context, args LookupContainerRegistryIPRestrictionsRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupContainerRegistryIPRestrictionsRegistryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupContainerRegistryIPRestrictionsRegistryResultOutput, error) {
			args := v.(LookupContainerRegistryIPRestrictionsRegistryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getContainerRegistryIPRestrictionsRegistry:getContainerRegistryIPRestrictionsRegistry", args, LookupContainerRegistryIPRestrictionsRegistryResultOutput{}, options).(LookupContainerRegistryIPRestrictionsRegistryResultOutput), nil
		}).(LookupContainerRegistryIPRestrictionsRegistryResultOutput)
}

// A collection of arguments for invoking getContainerRegistryIPRestrictionsRegistry.
type LookupContainerRegistryIPRestrictionsRegistryOutputArgs struct {
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringInput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContainerRegistryIPRestrictionsRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryIPRestrictionsRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistryIPRestrictionsRegistry.
type LookupContainerRegistryIPRestrictionsRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupContainerRegistryIPRestrictionsRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryIPRestrictionsRegistryResult)(nil)).Elem()
}

func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) ToLookupContainerRegistryIPRestrictionsRegistryResultOutput() LookupContainerRegistryIPRestrictionsRegistryResultOutput {
	return o
}

func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) ToLookupContainerRegistryIPRestrictionsRegistryResultOutputWithContext(ctx context.Context) LookupContainerRegistryIPRestrictionsRegistryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP restrictions applied on artifact manager component.
func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) IpRestrictions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsRegistryResult) []map[string]string {
		return v.IpRestrictions
	}).(pulumi.StringMapArrayOutput)
}

// The ID of the Managed Private Registry.
func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsRegistryResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o LookupContainerRegistryIPRestrictionsRegistryResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsRegistryResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerRegistryIPRestrictionsRegistryResultOutput{})
}
