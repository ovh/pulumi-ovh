// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetKubeNodes(ctx *pulumi.Context, args *GetKubeNodesArgs, opts ...pulumi.InvokeOption) (*GetKubeNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubeNodesResult
	err := ctx.Invoke("ovh:CloudProject/getKubeNodes:getKubeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeNodes.
type GetKubeNodesArgs struct {
	KubeId      string `pulumi:"kubeId"`
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeNodes.
type GetKubeNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string             `pulumi:"id"`
	KubeId      string             `pulumi:"kubeId"`
	Nodes       []GetKubeNodesNode `pulumi:"nodes"`
	ServiceName string             `pulumi:"serviceName"`
}

func GetKubeNodesOutput(ctx *pulumi.Context, args GetKubeNodesOutputArgs, opts ...pulumi.InvokeOption) GetKubeNodesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetKubeNodesResultOutput, error) {
			args := v.(GetKubeNodesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getKubeNodes:getKubeNodes", args, GetKubeNodesResultOutput{}, options).(GetKubeNodesResultOutput), nil
		}).(GetKubeNodesResultOutput)
}

// A collection of arguments for invoking getKubeNodes.
type GetKubeNodesOutputArgs struct {
	KubeId      pulumi.StringInput `pulumi:"kubeId"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetKubeNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodesArgs)(nil)).Elem()
}

// A collection of values returned by getKubeNodes.
type GetKubeNodesResultOutput struct{ *pulumi.OutputState }

func (GetKubeNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodesResult)(nil)).Elem()
}

func (o GetKubeNodesResultOutput) ToGetKubeNodesResultOutput() GetKubeNodesResultOutput {
	return o
}

func (o GetKubeNodesResultOutput) ToGetKubeNodesResultOutputWithContext(ctx context.Context) GetKubeNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubeNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetKubeNodesResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.KubeId }).(pulumi.StringOutput)
}

func (o GetKubeNodesResultOutput) Nodes() GetKubeNodesNodeArrayOutput {
	return o.ApplyT(func(v GetKubeNodesResult) []GetKubeNodesNode { return v.Nodes }).(GetKubeNodesNodeArrayOutput)
}

func (o GetKubeNodesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubeNodesResultOutput{})
}
