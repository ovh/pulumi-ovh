// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage backups for the given volume in a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudproject.NewVolumeBackup(ctx, "backup", &cloudproject.VolumeBackupArgs{
//				ServiceName: pulumi.String("<public cloud project ID>"),
//				RegionName:  pulumi.String("GRA9"),
//				VolumeId:    pulumi.String("<volume ID>"),
//				Name:        pulumi.String("ExampleBackup"),
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
// A volume backup in a public cloud project can be imported using the `service_name`, `region_name` and `id` attributes. Using the following configuration:
//
// terraform
//
// import {
//
//	id = "<service_name>/<region_name>/<id>"
//
//	to = ovh_cloud_project_volume_backup.backup
//
// }
//
// You can then run:
//
// bash
//
// $ pulumi preview -generate-config-out=backup.tf
//
// $ pulumi up
//
// The file `backup.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type VolumeBackup struct {
	pulumi.CustomResourceState

	// Creation date of the backup
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// name of the backup
	Name pulumi.StringOutput `pulumi:"name"`
	// Volume backup region
	Region pulumi.StringOutput `pulumi:"region"`
	// Region name
	RegionName pulumi.StringOutput `pulumi:"regionName"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Size of the backup in GiB
	Size pulumi.Float64Output `pulumi:"size"`
	// Staus of the backup
	Status pulumi.StringOutput `pulumi:"status"`
	// ID of the volume to backup
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeBackup registers a new resource with the given unique name, arguments, and options.
func NewVolumeBackup(ctx *pulumi.Context,
	name string, args *VolumeBackupArgs, opts ...pulumi.ResourceOption) (*VolumeBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegionName == nil {
		return nil, errors.New("invalid value for required argument 'RegionName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VolumeBackup
	err := ctx.RegisterResource("ovh:CloudProject/volumeBackup:VolumeBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeBackup gets an existing VolumeBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeBackupState, opts ...pulumi.ResourceOption) (*VolumeBackup, error) {
	var resource VolumeBackup
	err := ctx.ReadResource("ovh:CloudProject/volumeBackup:VolumeBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeBackup resources.
type volumeBackupState struct {
	// Creation date of the backup
	CreationDate *string `pulumi:"creationDate"`
	// name of the backup
	Name *string `pulumi:"name"`
	// Volume backup region
	Region *string `pulumi:"region"`
	// Region name
	RegionName *string `pulumi:"regionName"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Size of the backup in GiB
	Size *float64 `pulumi:"size"`
	// Staus of the backup
	Status *string `pulumi:"status"`
	// ID of the volume to backup
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeBackupState struct {
	// Creation date of the backup
	CreationDate pulumi.StringPtrInput
	// name of the backup
	Name pulumi.StringPtrInput
	// Volume backup region
	Region pulumi.StringPtrInput
	// Region name
	RegionName pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Size of the backup in GiB
	Size pulumi.Float64PtrInput
	// Staus of the backup
	Status pulumi.StringPtrInput
	// ID of the volume to backup
	VolumeId pulumi.StringPtrInput
}

func (VolumeBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeBackupState)(nil)).Elem()
}

type volumeBackupArgs struct {
	// name of the backup
	Name *string `pulumi:"name"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
	// ID of the volume to backup
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeBackup resource.
type VolumeBackupArgs struct {
	// name of the backup
	Name pulumi.StringPtrInput
	// Region name
	RegionName pulumi.StringInput
	// Service name
	ServiceName pulumi.StringInput
	// ID of the volume to backup
	VolumeId pulumi.StringInput
}

func (VolumeBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeBackupArgs)(nil)).Elem()
}

type VolumeBackupInput interface {
	pulumi.Input

	ToVolumeBackupOutput() VolumeBackupOutput
	ToVolumeBackupOutputWithContext(ctx context.Context) VolumeBackupOutput
}

func (*VolumeBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackup)(nil)).Elem()
}

func (i *VolumeBackup) ToVolumeBackupOutput() VolumeBackupOutput {
	return i.ToVolumeBackupOutputWithContext(context.Background())
}

func (i *VolumeBackup) ToVolumeBackupOutputWithContext(ctx context.Context) VolumeBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupOutput)
}

// VolumeBackupArrayInput is an input type that accepts VolumeBackupArray and VolumeBackupArrayOutput values.
// You can construct a concrete instance of `VolumeBackupArrayInput` via:
//
//	VolumeBackupArray{ VolumeBackupArgs{...} }
type VolumeBackupArrayInput interface {
	pulumi.Input

	ToVolumeBackupArrayOutput() VolumeBackupArrayOutput
	ToVolumeBackupArrayOutputWithContext(context.Context) VolumeBackupArrayOutput
}

type VolumeBackupArray []VolumeBackupInput

func (VolumeBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeBackup)(nil)).Elem()
}

func (i VolumeBackupArray) ToVolumeBackupArrayOutput() VolumeBackupArrayOutput {
	return i.ToVolumeBackupArrayOutputWithContext(context.Background())
}

func (i VolumeBackupArray) ToVolumeBackupArrayOutputWithContext(ctx context.Context) VolumeBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupArrayOutput)
}

// VolumeBackupMapInput is an input type that accepts VolumeBackupMap and VolumeBackupMapOutput values.
// You can construct a concrete instance of `VolumeBackupMapInput` via:
//
//	VolumeBackupMap{ "key": VolumeBackupArgs{...} }
type VolumeBackupMapInput interface {
	pulumi.Input

	ToVolumeBackupMapOutput() VolumeBackupMapOutput
	ToVolumeBackupMapOutputWithContext(context.Context) VolumeBackupMapOutput
}

type VolumeBackupMap map[string]VolumeBackupInput

func (VolumeBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeBackup)(nil)).Elem()
}

func (i VolumeBackupMap) ToVolumeBackupMapOutput() VolumeBackupMapOutput {
	return i.ToVolumeBackupMapOutputWithContext(context.Background())
}

func (i VolumeBackupMap) ToVolumeBackupMapOutputWithContext(ctx context.Context) VolumeBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeBackupMapOutput)
}

type VolumeBackupOutput struct{ *pulumi.OutputState }

func (VolumeBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeBackup)(nil)).Elem()
}

func (o VolumeBackupOutput) ToVolumeBackupOutput() VolumeBackupOutput {
	return o
}

func (o VolumeBackupOutput) ToVolumeBackupOutputWithContext(ctx context.Context) VolumeBackupOutput {
	return o
}

// Creation date of the backup
func (o VolumeBackupOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// name of the backup
func (o VolumeBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Volume backup region
func (o VolumeBackupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Region name
func (o VolumeBackupOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.RegionName }).(pulumi.StringOutput)
}

// Service name
func (o VolumeBackupOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Size of the backup in GiB
func (o VolumeBackupOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *VolumeBackup) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Staus of the backup
func (o VolumeBackupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// ID of the volume to backup
func (o VolumeBackupOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeBackup) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

type VolumeBackupArrayOutput struct{ *pulumi.OutputState }

func (VolumeBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeBackup)(nil)).Elem()
}

func (o VolumeBackupArrayOutput) ToVolumeBackupArrayOutput() VolumeBackupArrayOutput {
	return o
}

func (o VolumeBackupArrayOutput) ToVolumeBackupArrayOutputWithContext(ctx context.Context) VolumeBackupArrayOutput {
	return o
}

func (o VolumeBackupArrayOutput) Index(i pulumi.IntInput) VolumeBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeBackup {
		return vs[0].([]*VolumeBackup)[vs[1].(int)]
	}).(VolumeBackupOutput)
}

type VolumeBackupMapOutput struct{ *pulumi.OutputState }

func (VolumeBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeBackup)(nil)).Elem()
}

func (o VolumeBackupMapOutput) ToVolumeBackupMapOutput() VolumeBackupMapOutput {
	return o
}

func (o VolumeBackupMapOutput) ToVolumeBackupMapOutputWithContext(ctx context.Context) VolumeBackupMapOutput {
	return o
}

func (o VolumeBackupMapOutput) MapIndex(k pulumi.StringInput) VolumeBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeBackup {
		return vs[0].(map[string]*VolumeBackup)[vs[1].(string)]
	}).(VolumeBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeBackupInput)(nil)).Elem(), &VolumeBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeBackupArrayInput)(nil)).Elem(), VolumeBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeBackupMapInput)(nil)).Elem(), VolumeBackupMap{})
	pulumi.RegisterOutputType(VolumeBackupOutput{})
	pulumi.RegisterOutputType(VolumeBackupArrayOutput{})
	pulumi.RegisterOutputType(VolumeBackupMapOutput{})
}
