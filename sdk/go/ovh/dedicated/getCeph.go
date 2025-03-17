// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a dedicated CEPH.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dedicated"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dedicated.GetCeph(ctx, &dedicated.GetCephArgs{
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCeph(ctx *pulumi.Context, args *GetCephArgs, opts ...pulumi.InvokeOption) (*GetCephResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCephResult
	err := ctx.Invoke("ovh:Dedicated/getCeph:getCeph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCeph.
type GetCephArgs struct {
	// CEPH cluster version
	CephVersion *string `pulumi:"cephVersion"`
	// The service name of the dedicated CEPH cluster.
	ServiceName string `pulumi:"serviceName"`
	// the status of the service
	Status *string `pulumi:"status"`
}

// A collection of values returned by getCeph.
type GetCephResult struct {
	// URN of the CEPH instance
	CephURN string `pulumi:"CephURN"`
	// list of CEPH monitors IPs
	CephMons []string `pulumi:"cephMons"`
	// CEPH cluster version
	CephVersion string `pulumi:"cephVersion"`
	// CRUSH algorithm settings. Possible values
	// * OPTIMAL
	// * DEFAULT
	// * LEGACY
	// * BOBTAIL
	// * ARGONAUT
	// * FIREFLY
	// * HAMMER
	// * JEWEL
	CrushTunables string `pulumi:"crushTunables"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// CEPH cluster label
	Label string `pulumi:"label"`
	// cluster region
	Region      string `pulumi:"region"`
	ServiceName string `pulumi:"serviceName"`
	// Cluster size in TB
	Size float64 `pulumi:"size"`
	// the state of the cluster
	State string `pulumi:"state"`
	// the status of the service
	Status string `pulumi:"status"`
}

func GetCephOutput(ctx *pulumi.Context, args GetCephOutputArgs, opts ...pulumi.InvokeOption) GetCephResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCephResultOutput, error) {
			args := v.(GetCephArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dedicated/getCeph:getCeph", args, GetCephResultOutput{}, options).(GetCephResultOutput), nil
		}).(GetCephResultOutput)
}

// A collection of arguments for invoking getCeph.
type GetCephOutputArgs struct {
	// CEPH cluster version
	CephVersion pulumi.StringPtrInput `pulumi:"cephVersion"`
	// The service name of the dedicated CEPH cluster.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// the status of the service
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetCephOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCephArgs)(nil)).Elem()
}

// A collection of values returned by getCeph.
type GetCephResultOutput struct{ *pulumi.OutputState }

func (GetCephResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCephResult)(nil)).Elem()
}

func (o GetCephResultOutput) ToGetCephResultOutput() GetCephResultOutput {
	return o
}

func (o GetCephResultOutput) ToGetCephResultOutputWithContext(ctx context.Context) GetCephResultOutput {
	return o
}

// URN of the CEPH instance
func (o GetCephResultOutput) CephURN() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.CephURN }).(pulumi.StringOutput)
}

// list of CEPH monitors IPs
func (o GetCephResultOutput) CephMons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCephResult) []string { return v.CephMons }).(pulumi.StringArrayOutput)
}

// CEPH cluster version
func (o GetCephResultOutput) CephVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.CephVersion }).(pulumi.StringOutput)
}

// CRUSH algorithm settings. Possible values
// * OPTIMAL
// * DEFAULT
// * LEGACY
// * BOBTAIL
// * ARGONAUT
// * FIREFLY
// * HAMMER
// * JEWEL
func (o GetCephResultOutput) CrushTunables() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.CrushTunables }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCephResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.Id }).(pulumi.StringOutput)
}

// CEPH cluster label
func (o GetCephResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.Label }).(pulumi.StringOutput)
}

// cluster region
func (o GetCephResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetCephResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Cluster size in TB
func (o GetCephResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v GetCephResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// the state of the cluster
func (o GetCephResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.State }).(pulumi.StringOutput)
}

// the status of the service
func (o GetCephResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetCephResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCephResultOutput{})
}
