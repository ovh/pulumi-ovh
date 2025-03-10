// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFailoverIpAttach(ctx *pulumi.Context, args *LookupFailoverIpAttachArgs, opts ...pulumi.InvokeOption) (*LookupFailoverIpAttachResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFailoverIpAttachResult
	err := ctx.Invoke("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFailoverIpAttach.
type LookupFailoverIpAttachArgs struct {
	Block         *string `pulumi:"block"`
	ContinentCode *string `pulumi:"continentCode"`
	GeoLoc        *string `pulumi:"geoLoc"`
	Ip            *string `pulumi:"ip"`
	RoutedTo      *string `pulumi:"routedTo"`
	ServiceName   string  `pulumi:"serviceName"`
}

// A collection of values returned by getFailoverIpAttach.
type LookupFailoverIpAttachResult struct {
	Block         string `pulumi:"block"`
	ContinentCode string `pulumi:"continentCode"`
	GeoLoc        string `pulumi:"geoLoc"`
	Id            string `pulumi:"id"`
	Ip            string `pulumi:"ip"`
	Progress      int    `pulumi:"progress"`
	RoutedTo      string `pulumi:"routedTo"`
	ServiceName   string `pulumi:"serviceName"`
	Status        string `pulumi:"status"`
	SubType       string `pulumi:"subType"`
}

func LookupFailoverIpAttachOutput(ctx *pulumi.Context, args LookupFailoverIpAttachOutputArgs, opts ...pulumi.InvokeOption) LookupFailoverIpAttachResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFailoverIpAttachResultOutput, error) {
			args := v.(LookupFailoverIpAttachArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args, LookupFailoverIpAttachResultOutput{}, options).(LookupFailoverIpAttachResultOutput), nil
		}).(LookupFailoverIpAttachResultOutput)
}

// A collection of arguments for invoking getFailoverIpAttach.
type LookupFailoverIpAttachOutputArgs struct {
	Block         pulumi.StringPtrInput `pulumi:"block"`
	ContinentCode pulumi.StringPtrInput `pulumi:"continentCode"`
	GeoLoc        pulumi.StringPtrInput `pulumi:"geoLoc"`
	Ip            pulumi.StringPtrInput `pulumi:"ip"`
	RoutedTo      pulumi.StringPtrInput `pulumi:"routedTo"`
	ServiceName   pulumi.StringInput    `pulumi:"serviceName"`
}

func (LookupFailoverIpAttachOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverIpAttachArgs)(nil)).Elem()
}

// A collection of values returned by getFailoverIpAttach.
type LookupFailoverIpAttachResultOutput struct{ *pulumi.OutputState }

func (LookupFailoverIpAttachResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverIpAttachResult)(nil)).Elem()
}

func (o LookupFailoverIpAttachResultOutput) ToLookupFailoverIpAttachResultOutput() LookupFailoverIpAttachResultOutput {
	return o
}

func (o LookupFailoverIpAttachResultOutput) ToLookupFailoverIpAttachResultOutputWithContext(ctx context.Context) LookupFailoverIpAttachResultOutput {
	return o
}

func (o LookupFailoverIpAttachResultOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Block }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.ContinentCode }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) GeoLoc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.GeoLoc }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Ip }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) int { return v.Progress }).(pulumi.IntOutput)
}

func (o LookupFailoverIpAttachResultOutput) RoutedTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.RoutedTo }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.SubType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFailoverIpAttachResultOutput{})
}
