// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/internal"
)

// Use this data source to get a OVHcloud Managed Kubernetes Service cluster IP restrictions.
//
// ## Example Usage
func LookupKubeIpRestrictions(ctx *pulumi.Context, args *LookupKubeIpRestrictionsArgs, opts ...pulumi.InvokeOption) (*LookupKubeIpRestrictionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubeIpRestrictionsResult
	err := ctx.Invoke("ovh:CloudProject/getKubeIpRestrictions:getKubeIpRestrictions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeIpRestrictions.
type LookupKubeIpRestrictionsArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeIpRestrictions.
type LookupKubeIpRestrictionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of CIDRs that restricts the access to the API server.
	Ips []string `pulumi:"ips"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func LookupKubeIpRestrictionsOutput(ctx *pulumi.Context, args LookupKubeIpRestrictionsOutputArgs, opts ...pulumi.InvokeOption) LookupKubeIpRestrictionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubeIpRestrictionsResult, error) {
			args := v.(LookupKubeIpRestrictionsArgs)
			r, err := LookupKubeIpRestrictions(ctx, &args, opts...)
			var s LookupKubeIpRestrictionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubeIpRestrictionsResultOutput)
}

// A collection of arguments for invoking getKubeIpRestrictions.
type LookupKubeIpRestrictionsOutputArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupKubeIpRestrictionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeIpRestrictionsArgs)(nil)).Elem()
}

// A collection of values returned by getKubeIpRestrictions.
type LookupKubeIpRestrictionsResultOutput struct{ *pulumi.OutputState }

func (LookupKubeIpRestrictionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeIpRestrictionsResult)(nil)).Elem()
}

func (o LookupKubeIpRestrictionsResultOutput) ToLookupKubeIpRestrictionsResultOutput() LookupKubeIpRestrictionsResultOutput {
	return o
}

func (o LookupKubeIpRestrictionsResultOutput) ToLookupKubeIpRestrictionsResultOutputWithContext(ctx context.Context) LookupKubeIpRestrictionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKubeIpRestrictionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeIpRestrictionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of CIDRs that restricts the access to the API server.
func (o LookupKubeIpRestrictionsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubeIpRestrictionsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupKubeIpRestrictionsResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeIpRestrictionsResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKubeIpRestrictionsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeIpRestrictionsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubeIpRestrictionsResultOutput{})
}
