// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of Management IP Restrictions of a container registry associated with a public cloud project.
func LookupContainerRegistryIPRestrictionsManagement(ctx *pulumi.Context, args *LookupContainerRegistryIPRestrictionsManagementArgs, opts ...pulumi.InvokeOption) (*LookupContainerRegistryIPRestrictionsManagementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerRegistryIPRestrictionsManagementResult
	err := ctx.Invoke("ovh:CloudProject/getContainerRegistryIPRestrictionsManagement:getContainerRegistryIPRestrictionsManagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistryIPRestrictionsManagement.
type LookupContainerRegistryIPRestrictionsManagementArgs struct {
	// The id of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getContainerRegistryIPRestrictionsManagement.
type LookupContainerRegistryIPRestrictionsManagementResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP restrictions applied on Harbor UI and API.
	IpRestrictions []map[string]string `pulumi:"ipRestrictions"`
	// The ID of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

func LookupContainerRegistryIPRestrictionsManagementOutput(ctx *pulumi.Context, args LookupContainerRegistryIPRestrictionsManagementOutputArgs, opts ...pulumi.InvokeOption) LookupContainerRegistryIPRestrictionsManagementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerRegistryIPRestrictionsManagementResult, error) {
			args := v.(LookupContainerRegistryIPRestrictionsManagementArgs)
			r, err := LookupContainerRegistryIPRestrictionsManagement(ctx, &args, opts...)
			var s LookupContainerRegistryIPRestrictionsManagementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerRegistryIPRestrictionsManagementResultOutput)
}

// A collection of arguments for invoking getContainerRegistryIPRestrictionsManagement.
type LookupContainerRegistryIPRestrictionsManagementOutputArgs struct {
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringInput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContainerRegistryIPRestrictionsManagementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryIPRestrictionsManagementArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistryIPRestrictionsManagement.
type LookupContainerRegistryIPRestrictionsManagementResultOutput struct{ *pulumi.OutputState }

func (LookupContainerRegistryIPRestrictionsManagementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryIPRestrictionsManagementResult)(nil)).Elem()
}

func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) ToLookupContainerRegistryIPRestrictionsManagementResultOutput() LookupContainerRegistryIPRestrictionsManagementResultOutput {
	return o
}

func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) ToLookupContainerRegistryIPRestrictionsManagementResultOutputWithContext(ctx context.Context) LookupContainerRegistryIPRestrictionsManagementResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsManagementResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP restrictions applied on Harbor UI and API.
func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) IpRestrictions() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsManagementResult) []map[string]string {
		return v.IpRestrictions
	}).(pulumi.StringMapArrayOutput)
}

// The ID of the Managed Private Registry.
func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsManagementResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o LookupContainerRegistryIPRestrictionsManagementResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryIPRestrictionsManagementResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerRegistryIPRestrictionsManagementResultOutput{})
}
