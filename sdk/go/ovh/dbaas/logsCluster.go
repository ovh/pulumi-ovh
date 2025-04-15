// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.NewLogsCluster(ctx, "ldp", &dbaas.LogsClusterArgs{
//				ServiceName: pulumi.String("ldp-xx-xxxxx"),
//				ClusterId:   pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
//				ArchiveAllowedNetworks: pulumi.StringArray{
//					pulumi.String("10.0.0.0/16"),
//				},
//				DirectInputAllowedNetworks: pulumi.StringArray{
//					pulumi.String("10.0.0.0/16"),
//				},
//				QueryAllowedNetworks: pulumi.StringArray{
//					pulumi.String("10.0.0.0/16"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OVHcloud DBaaS Log Data Platform clusters can be imported using the `service_name` and `cluster_id` of the cluster, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Dbaas/logsCluster:LogsCluster ldp service_name/cluster_id
// ```
type LogsCluster struct {
	pulumi.CustomResourceState

	// List of IP blocks
	ArchiveAllowedNetworks pulumi.StringArrayOutput `pulumi:"archiveAllowedNetworks"`
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringPtrOutput `pulumi:"clusterId"`
	// type of cluster (DEDICATED, PRO or TRIAL)
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// PEM for dedicated inputs
	DedicatedInputPem pulumi.StringOutput `pulumi:"dedicatedInputPem"`
	// List of IP blocks
	DirectInputAllowedNetworks pulumi.StringArrayOutput `pulumi:"directInputAllowedNetworks"`
	// PEM for direct inputs
	DirectInputPem pulumi.StringOutput `pulumi:"directInputPem"`
	// cluster hostname hosting tenant
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Initial allowed networks for ARCHIVE flow type
	InitialArchiveAllowedNetworks pulumi.StringArrayOutput `pulumi:"initialArchiveAllowedNetworks"`
	// Initial allowed networks for DIRECT_INPUT flow type
	InitialDirectInputAllowedNetworks pulumi.StringArrayOutput `pulumi:"initialDirectInputAllowedNetworks"`
	// Initial allowed networks for QUERY flow type
	InitialQueryAllowedNetworks pulumi.StringArrayOutput `pulumi:"initialQueryAllowedNetworks"`
	// true if all content generated by given service will be placed on this cluster
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// true if given service can perform advanced operations on cluster
	IsUnlocked pulumi.BoolOutput `pulumi:"isUnlocked"`
	// List of IP blocks
	QueryAllowedNetworks pulumi.StringArrayOutput `pulumi:"queryAllowedNetworks"`
	// datacenter localization
	Region pulumi.StringOutput `pulumi:"region"`
	// The service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewLogsCluster registers a new resource with the given unique name, arguments, and options.
func NewLogsCluster(ctx *pulumi.Context,
	name string, args *LogsClusterArgs, opts ...pulumi.ResourceOption) (*LogsCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"dedicatedInputPem",
		"directInputPem",
		"initialArchiveAllowedNetworks",
		"initialDirectInputAllowedNetworks",
		"initialQueryAllowedNetworks",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogsCluster
	err := ctx.RegisterResource("ovh:Dbaas/logsCluster:LogsCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsCluster gets an existing LogsCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsClusterState, opts ...pulumi.ResourceOption) (*LogsCluster, error) {
	var resource LogsCluster
	err := ctx.ReadResource("ovh:Dbaas/logsCluster:LogsCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsCluster resources.
type logsClusterState struct {
	// List of IP blocks
	ArchiveAllowedNetworks []string `pulumi:"archiveAllowedNetworks"`
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId *string `pulumi:"clusterId"`
	// type of cluster (DEDICATED, PRO or TRIAL)
	ClusterType *string `pulumi:"clusterType"`
	// PEM for dedicated inputs
	DedicatedInputPem *string `pulumi:"dedicatedInputPem"`
	// List of IP blocks
	DirectInputAllowedNetworks []string `pulumi:"directInputAllowedNetworks"`
	// PEM for direct inputs
	DirectInputPem *string `pulumi:"directInputPem"`
	// cluster hostname hosting tenant
	Hostname *string `pulumi:"hostname"`
	// Initial allowed networks for ARCHIVE flow type
	InitialArchiveAllowedNetworks []string `pulumi:"initialArchiveAllowedNetworks"`
	// Initial allowed networks for DIRECT_INPUT flow type
	InitialDirectInputAllowedNetworks []string `pulumi:"initialDirectInputAllowedNetworks"`
	// Initial allowed networks for QUERY flow type
	InitialQueryAllowedNetworks []string `pulumi:"initialQueryAllowedNetworks"`
	// true if all content generated by given service will be placed on this cluster
	IsDefault *bool `pulumi:"isDefault"`
	// true if given service can perform advanced operations on cluster
	IsUnlocked *bool `pulumi:"isUnlocked"`
	// List of IP blocks
	QueryAllowedNetworks []string `pulumi:"queryAllowedNetworks"`
	// datacenter localization
	Region *string `pulumi:"region"`
	// The service name
	ServiceName *string `pulumi:"serviceName"`
}

type LogsClusterState struct {
	// List of IP blocks
	ArchiveAllowedNetworks pulumi.StringArrayInput
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringPtrInput
	// type of cluster (DEDICATED, PRO or TRIAL)
	ClusterType pulumi.StringPtrInput
	// PEM for dedicated inputs
	DedicatedInputPem pulumi.StringPtrInput
	// List of IP blocks
	DirectInputAllowedNetworks pulumi.StringArrayInput
	// PEM for direct inputs
	DirectInputPem pulumi.StringPtrInput
	// cluster hostname hosting tenant
	Hostname pulumi.StringPtrInput
	// Initial allowed networks for ARCHIVE flow type
	InitialArchiveAllowedNetworks pulumi.StringArrayInput
	// Initial allowed networks for DIRECT_INPUT flow type
	InitialDirectInputAllowedNetworks pulumi.StringArrayInput
	// Initial allowed networks for QUERY flow type
	InitialQueryAllowedNetworks pulumi.StringArrayInput
	// true if all content generated by given service will be placed on this cluster
	IsDefault pulumi.BoolPtrInput
	// true if given service can perform advanced operations on cluster
	IsUnlocked pulumi.BoolPtrInput
	// List of IP blocks
	QueryAllowedNetworks pulumi.StringArrayInput
	// datacenter localization
	Region pulumi.StringPtrInput
	// The service name
	ServiceName pulumi.StringPtrInput
}

func (LogsClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsClusterState)(nil)).Elem()
}

type logsClusterArgs struct {
	// List of IP blocks
	ArchiveAllowedNetworks []string `pulumi:"archiveAllowedNetworks"`
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId *string `pulumi:"clusterId"`
	// List of IP blocks
	DirectInputAllowedNetworks []string `pulumi:"directInputAllowedNetworks"`
	// List of IP blocks
	QueryAllowedNetworks []string `pulumi:"queryAllowedNetworks"`
	// The service name
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a LogsCluster resource.
type LogsClusterArgs struct {
	// List of IP blocks
	ArchiveAllowedNetworks pulumi.StringArrayInput
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringPtrInput
	// List of IP blocks
	DirectInputAllowedNetworks pulumi.StringArrayInput
	// List of IP blocks
	QueryAllowedNetworks pulumi.StringArrayInput
	// The service name
	ServiceName pulumi.StringInput
}

func (LogsClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsClusterArgs)(nil)).Elem()
}

type LogsClusterInput interface {
	pulumi.Input

	ToLogsClusterOutput() LogsClusterOutput
	ToLogsClusterOutputWithContext(ctx context.Context) LogsClusterOutput
}

func (*LogsCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsCluster)(nil)).Elem()
}

func (i *LogsCluster) ToLogsClusterOutput() LogsClusterOutput {
	return i.ToLogsClusterOutputWithContext(context.Background())
}

func (i *LogsCluster) ToLogsClusterOutputWithContext(ctx context.Context) LogsClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsClusterOutput)
}

// LogsClusterArrayInput is an input type that accepts LogsClusterArray and LogsClusterArrayOutput values.
// You can construct a concrete instance of `LogsClusterArrayInput` via:
//
//	LogsClusterArray{ LogsClusterArgs{...} }
type LogsClusterArrayInput interface {
	pulumi.Input

	ToLogsClusterArrayOutput() LogsClusterArrayOutput
	ToLogsClusterArrayOutputWithContext(context.Context) LogsClusterArrayOutput
}

type LogsClusterArray []LogsClusterInput

func (LogsClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsCluster)(nil)).Elem()
}

func (i LogsClusterArray) ToLogsClusterArrayOutput() LogsClusterArrayOutput {
	return i.ToLogsClusterArrayOutputWithContext(context.Background())
}

func (i LogsClusterArray) ToLogsClusterArrayOutputWithContext(ctx context.Context) LogsClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsClusterArrayOutput)
}

// LogsClusterMapInput is an input type that accepts LogsClusterMap and LogsClusterMapOutput values.
// You can construct a concrete instance of `LogsClusterMapInput` via:
//
//	LogsClusterMap{ "key": LogsClusterArgs{...} }
type LogsClusterMapInput interface {
	pulumi.Input

	ToLogsClusterMapOutput() LogsClusterMapOutput
	ToLogsClusterMapOutputWithContext(context.Context) LogsClusterMapOutput
}

type LogsClusterMap map[string]LogsClusterInput

func (LogsClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsCluster)(nil)).Elem()
}

func (i LogsClusterMap) ToLogsClusterMapOutput() LogsClusterMapOutput {
	return i.ToLogsClusterMapOutputWithContext(context.Background())
}

func (i LogsClusterMap) ToLogsClusterMapOutputWithContext(ctx context.Context) LogsClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsClusterMapOutput)
}

type LogsClusterOutput struct{ *pulumi.OutputState }

func (LogsClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsCluster)(nil)).Elem()
}

func (o LogsClusterOutput) ToLogsClusterOutput() LogsClusterOutput {
	return o
}

func (o LogsClusterOutput) ToLogsClusterOutputWithContext(ctx context.Context) LogsClusterOutput {
	return o
}

// List of IP blocks
func (o LogsClusterOutput) ArchiveAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.ArchiveAllowedNetworks }).(pulumi.StringArrayOutput)
}

// Cluster ID. If not provided, the default clusterId is used
func (o LogsClusterOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringPtrOutput { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// type of cluster (DEDICATED, PRO or TRIAL)
func (o LogsClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// PEM for dedicated inputs
func (o LogsClusterOutput) DedicatedInputPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.DedicatedInputPem }).(pulumi.StringOutput)
}

// List of IP blocks
func (o LogsClusterOutput) DirectInputAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.DirectInputAllowedNetworks }).(pulumi.StringArrayOutput)
}

// PEM for direct inputs
func (o LogsClusterOutput) DirectInputPem() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.DirectInputPem }).(pulumi.StringOutput)
}

// cluster hostname hosting tenant
func (o LogsClusterOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Initial allowed networks for ARCHIVE flow type
func (o LogsClusterOutput) InitialArchiveAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.InitialArchiveAllowedNetworks }).(pulumi.StringArrayOutput)
}

// Initial allowed networks for DIRECT_INPUT flow type
func (o LogsClusterOutput) InitialDirectInputAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.InitialDirectInputAllowedNetworks }).(pulumi.StringArrayOutput)
}

// Initial allowed networks for QUERY flow type
func (o LogsClusterOutput) InitialQueryAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.InitialQueryAllowedNetworks }).(pulumi.StringArrayOutput)
}

// true if all content generated by given service will be placed on this cluster
func (o LogsClusterOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// true if given service can perform advanced operations on cluster
func (o LogsClusterOutput) IsUnlocked() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.BoolOutput { return v.IsUnlocked }).(pulumi.BoolOutput)
}

// List of IP blocks
func (o LogsClusterOutput) QueryAllowedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringArrayOutput { return v.QueryAllowedNetworks }).(pulumi.StringArrayOutput)
}

// datacenter localization
func (o LogsClusterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The service name
func (o LogsClusterOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsCluster) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type LogsClusterArrayOutput struct{ *pulumi.OutputState }

func (LogsClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsCluster)(nil)).Elem()
}

func (o LogsClusterArrayOutput) ToLogsClusterArrayOutput() LogsClusterArrayOutput {
	return o
}

func (o LogsClusterArrayOutput) ToLogsClusterArrayOutputWithContext(ctx context.Context) LogsClusterArrayOutput {
	return o
}

func (o LogsClusterArrayOutput) Index(i pulumi.IntInput) LogsClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogsCluster {
		return vs[0].([]*LogsCluster)[vs[1].(int)]
	}).(LogsClusterOutput)
}

type LogsClusterMapOutput struct{ *pulumi.OutputState }

func (LogsClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsCluster)(nil)).Elem()
}

func (o LogsClusterMapOutput) ToLogsClusterMapOutput() LogsClusterMapOutput {
	return o
}

func (o LogsClusterMapOutput) ToLogsClusterMapOutputWithContext(ctx context.Context) LogsClusterMapOutput {
	return o
}

func (o LogsClusterMapOutput) MapIndex(k pulumi.StringInput) LogsClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogsCluster {
		return vs[0].(map[string]*LogsCluster)[vs[1].(string)]
	}).(LogsClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogsClusterInput)(nil)).Elem(), &LogsCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsClusterArrayInput)(nil)).Elem(), LogsClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsClusterMapInput)(nil)).Elem(), LogsClusterMap{})
	pulumi.RegisterOutputType(LogsClusterOutput{})
	pulumi.RegisterOutputType(LogsClusterArrayOutput{})
	pulumi.RegisterOutputType(LogsClusterMapOutput{})
}
