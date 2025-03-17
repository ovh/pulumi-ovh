// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for managing partitions on HA-NAS services
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/dedicated"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dedicated.NewNasHAPartition(ctx, "myPartition", &dedicated.NasHAPartitionArgs{
//				Protocol:    pulumi.String("NFS"),
//				ServiceName: pulumi.String("zpool-12345"),
//				Size:        pulumi.Int(20),
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
// HA-NAS can be imported using the `{service_name}/{name}`, e.g.
//
// ```sh
// $ pulumi import ovh:Dedicated/nasHAPartition:NasHAPartition my-partition zpool-12345/my-partition`
// ```
type NasHAPartition struct {
	pulumi.CustomResourceState

	// Percentage of partition space used in %
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// A brief description of the partition
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// name of the partition
	Name pulumi.StringOutput `pulumi:"name"`
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// size of the partition in GB
	Size pulumi.IntOutput `pulumi:"size"`
	// Percentage of partition space used by snapshots in %
	UsedBySnapshots pulumi.IntOutput `pulumi:"usedBySnapshots"`
}

// NewNasHAPartition registers a new resource with the given unique name, arguments, and options.
func NewNasHAPartition(ctx *pulumi.Context,
	name string, args *NasHAPartitionArgs, opts ...pulumi.ResourceOption) (*NasHAPartition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NasHAPartition
	err := ctx.RegisterResource("ovh:Dedicated/nasHAPartition:NasHAPartition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNasHAPartition gets an existing NasHAPartition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNasHAPartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NasHAPartitionState, opts ...pulumi.ResourceOption) (*NasHAPartition, error) {
	var resource NasHAPartition
	err := ctx.ReadResource("ovh:Dedicated/nasHAPartition:NasHAPartition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NasHAPartition resources.
type nasHAPartitionState struct {
	// Percentage of partition space used in %
	Capacity *int `pulumi:"capacity"`
	// A brief description of the partition
	Description *string `pulumi:"description"`
	// name of the partition
	Name *string `pulumi:"name"`
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol *string `pulumi:"protocol"`
	// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
	ServiceName *string `pulumi:"serviceName"`
	// size of the partition in GB
	Size *int `pulumi:"size"`
	// Percentage of partition space used by snapshots in %
	UsedBySnapshots *int `pulumi:"usedBySnapshots"`
}

type NasHAPartitionState struct {
	// Percentage of partition space used in %
	Capacity pulumi.IntPtrInput
	// A brief description of the partition
	Description pulumi.StringPtrInput
	// name of the partition
	Name pulumi.StringPtrInput
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol pulumi.StringPtrInput
	// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
	ServiceName pulumi.StringPtrInput
	// size of the partition in GB
	Size pulumi.IntPtrInput
	// Percentage of partition space used by snapshots in %
	UsedBySnapshots pulumi.IntPtrInput
}

func (NasHAPartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*nasHAPartitionState)(nil)).Elem()
}

type nasHAPartitionArgs struct {
	// A brief description of the partition
	Description *string `pulumi:"description"`
	// name of the partition
	Name *string `pulumi:"name"`
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol string `pulumi:"protocol"`
	// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
	ServiceName string `pulumi:"serviceName"`
	// size of the partition in GB
	Size int `pulumi:"size"`
}

// The set of arguments for constructing a NasHAPartition resource.
type NasHAPartitionArgs struct {
	// A brief description of the partition
	Description pulumi.StringPtrInput
	// name of the partition
	Name pulumi.StringPtrInput
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol pulumi.StringInput
	// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
	ServiceName pulumi.StringInput
	// size of the partition in GB
	Size pulumi.IntInput
}

func (NasHAPartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nasHAPartitionArgs)(nil)).Elem()
}

type NasHAPartitionInput interface {
	pulumi.Input

	ToNasHAPartitionOutput() NasHAPartitionOutput
	ToNasHAPartitionOutputWithContext(ctx context.Context) NasHAPartitionOutput
}

func (*NasHAPartition) ElementType() reflect.Type {
	return reflect.TypeOf((**NasHAPartition)(nil)).Elem()
}

func (i *NasHAPartition) ToNasHAPartitionOutput() NasHAPartitionOutput {
	return i.ToNasHAPartitionOutputWithContext(context.Background())
}

func (i *NasHAPartition) ToNasHAPartitionOutputWithContext(ctx context.Context) NasHAPartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasHAPartitionOutput)
}

// NasHAPartitionArrayInput is an input type that accepts NasHAPartitionArray and NasHAPartitionArrayOutput values.
// You can construct a concrete instance of `NasHAPartitionArrayInput` via:
//
//	NasHAPartitionArray{ NasHAPartitionArgs{...} }
type NasHAPartitionArrayInput interface {
	pulumi.Input

	ToNasHAPartitionArrayOutput() NasHAPartitionArrayOutput
	ToNasHAPartitionArrayOutputWithContext(context.Context) NasHAPartitionArrayOutput
}

type NasHAPartitionArray []NasHAPartitionInput

func (NasHAPartitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasHAPartition)(nil)).Elem()
}

func (i NasHAPartitionArray) ToNasHAPartitionArrayOutput() NasHAPartitionArrayOutput {
	return i.ToNasHAPartitionArrayOutputWithContext(context.Background())
}

func (i NasHAPartitionArray) ToNasHAPartitionArrayOutputWithContext(ctx context.Context) NasHAPartitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasHAPartitionArrayOutput)
}

// NasHAPartitionMapInput is an input type that accepts NasHAPartitionMap and NasHAPartitionMapOutput values.
// You can construct a concrete instance of `NasHAPartitionMapInput` via:
//
//	NasHAPartitionMap{ "key": NasHAPartitionArgs{...} }
type NasHAPartitionMapInput interface {
	pulumi.Input

	ToNasHAPartitionMapOutput() NasHAPartitionMapOutput
	ToNasHAPartitionMapOutputWithContext(context.Context) NasHAPartitionMapOutput
}

type NasHAPartitionMap map[string]NasHAPartitionInput

func (NasHAPartitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasHAPartition)(nil)).Elem()
}

func (i NasHAPartitionMap) ToNasHAPartitionMapOutput() NasHAPartitionMapOutput {
	return i.ToNasHAPartitionMapOutputWithContext(context.Background())
}

func (i NasHAPartitionMap) ToNasHAPartitionMapOutputWithContext(ctx context.Context) NasHAPartitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasHAPartitionMapOutput)
}

type NasHAPartitionOutput struct{ *pulumi.OutputState }

func (NasHAPartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NasHAPartition)(nil)).Elem()
}

func (o NasHAPartitionOutput) ToNasHAPartitionOutput() NasHAPartitionOutput {
	return o
}

func (o NasHAPartitionOutput) ToNasHAPartitionOutputWithContext(ctx context.Context) NasHAPartitionOutput {
	return o
}

// Percentage of partition space used in %
func (o NasHAPartitionOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

// A brief description of the partition
func (o NasHAPartitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// name of the partition
func (o NasHAPartitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// one of "NFS", "CIFS" or "NFS_CIFS"
func (o NasHAPartitionOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
func (o NasHAPartitionOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// size of the partition in GB
func (o NasHAPartitionOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Percentage of partition space used by snapshots in %
func (o NasHAPartitionOutput) UsedBySnapshots() pulumi.IntOutput {
	return o.ApplyT(func(v *NasHAPartition) pulumi.IntOutput { return v.UsedBySnapshots }).(pulumi.IntOutput)
}

type NasHAPartitionArrayOutput struct{ *pulumi.OutputState }

func (NasHAPartitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasHAPartition)(nil)).Elem()
}

func (o NasHAPartitionArrayOutput) ToNasHAPartitionArrayOutput() NasHAPartitionArrayOutput {
	return o
}

func (o NasHAPartitionArrayOutput) ToNasHAPartitionArrayOutputWithContext(ctx context.Context) NasHAPartitionArrayOutput {
	return o
}

func (o NasHAPartitionArrayOutput) Index(i pulumi.IntInput) NasHAPartitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NasHAPartition {
		return vs[0].([]*NasHAPartition)[vs[1].(int)]
	}).(NasHAPartitionOutput)
}

type NasHAPartitionMapOutput struct{ *pulumi.OutputState }

func (NasHAPartitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasHAPartition)(nil)).Elem()
}

func (o NasHAPartitionMapOutput) ToNasHAPartitionMapOutput() NasHAPartitionMapOutput {
	return o
}

func (o NasHAPartitionMapOutput) ToNasHAPartitionMapOutputWithContext(ctx context.Context) NasHAPartitionMapOutput {
	return o
}

func (o NasHAPartitionMapOutput) MapIndex(k pulumi.StringInput) NasHAPartitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NasHAPartition {
		return vs[0].(map[string]*NasHAPartition)[vs[1].(string)]
	}).(NasHAPartitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NasHAPartitionInput)(nil)).Elem(), &NasHAPartition{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasHAPartitionArrayInput)(nil)).Elem(), NasHAPartitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasHAPartitionMapInput)(nil)).Elem(), NasHAPartitionMap{})
	pulumi.RegisterOutputType(NasHAPartitionOutput{})
	pulumi.RegisterOutputType(NasHAPartitionArrayOutput{})
	pulumi.RegisterOutputType(NasHAPartitionMapOutput{})
}
