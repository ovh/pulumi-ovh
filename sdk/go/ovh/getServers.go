// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/internal"
)

// Use this data source to get the list of dedicated servers associated with your OVHcloud Account.
//
// ## Example Usage
func GetServers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServersResult
	err := ctx.Invoke("ovh:index/getServers:getServers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getServers.
type GetServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of dedicated servers IDs associated with your OVHcloud Account.
	Results []string `pulumi:"results"`
}
