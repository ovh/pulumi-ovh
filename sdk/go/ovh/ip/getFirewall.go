// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewall(ctx *pulumi.Context, args *LookupFirewallArgs, opts ...pulumi.InvokeOption) (*LookupFirewallResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallResult
	err := ctx.Invoke("ovh:Ip/getFirewall:getFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewall.
type LookupFirewallArgs struct {
	Ip           string `pulumi:"ip"`
	IpOnFirewall string `pulumi:"ipOnFirewall"`
}

// A collection of values returned by getFirewall.
type LookupFirewallResult struct {
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	Ip           string `pulumi:"ip"`
	IpOnFirewall string `pulumi:"ipOnFirewall"`
	State        string `pulumi:"state"`
}

func LookupFirewallOutput(ctx *pulumi.Context, args LookupFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFirewallResultOutput, error) {
			args := v.(LookupFirewallArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Ip/getFirewall:getFirewall", args, LookupFirewallResultOutput{}, options).(LookupFirewallResultOutput), nil
		}).(LookupFirewallResultOutput)
}

// A collection of arguments for invoking getFirewall.
type LookupFirewallOutputArgs struct {
	Ip           pulumi.StringInput `pulumi:"ip"`
	IpOnFirewall pulumi.StringInput `pulumi:"ipOnFirewall"`
}

func (LookupFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallArgs)(nil)).Elem()
}

// A collection of values returned by getFirewall.
type LookupFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallResult)(nil)).Elem()
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutput() LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutputWithContext(ctx context.Context) LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFirewallResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Ip }).(pulumi.StringOutput)
}

func (o LookupFirewallResultOutput) IpOnFirewall() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.IpOnFirewall }).(pulumi.StringOutput)
}

func (o LookupFirewallResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallResultOutput{})
}
