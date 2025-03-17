// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **This resource uses a Beta API**
// Creates an instance associated with a public cloud project.
//
// ## Example Usage
//
// Create a instance.
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
//			_, err := cloudproject.NewInstance(ctx, "instance", &cloudproject.InstanceArgs{
//				BillingPeriod: pulumi.String("hourly"),
//				BootFrom: &cloudproject.InstanceBootFromArgs{
//					ImageId: pulumi.String("UUID"),
//				},
//				Flavor: &cloudproject.InstanceFlavorArgs{
//					FlavorId: pulumi.String("UUID"),
//				},
//				Network: &cloudproject.InstanceNetworkArgs{
//					Public: pulumi.Bool(true),
//				},
//				Region:      pulumi.String("RRRR"),
//				ServiceName: pulumi.String("XXX"),
//				SshKey: &cloudproject.InstanceSshKeyArgs{
//					Name: pulumi.String("sshname"),
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
type Instance struct {
	pulumi.CustomResourceState

	// Instance IP addresses
	Addresses InstanceAddressArrayOutput `pulumi:"addresses"`
	// Volumes attached to the instance
	AttachedVolumes InstanceAttachedVolumeArrayOutput `pulumi:"attachedVolumes"`
	// Create an autobackup workflow after instance start up.
	AutoBackup InstanceAutoBackupPtrOutput `pulumi:"autoBackup"`
	// The availability zone where the instance will be created
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Billing period - hourly or monthly
	BillingPeriod pulumi.StringOutput `pulumi:"billingPeriod"`
	// Boot the instance from an image or a volume
	BootFrom InstanceBootFromOutput `pulumi:"bootFrom"`
	// Create multiple instances
	Bulk pulumi.IntPtrOutput `pulumi:"bulk"`
	// Flavor information
	Flavor InstanceFlavorOutput `pulumi:"flavor"`
	// Flavor id
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// Flavor name
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// Start instance in group
	Group InstanceGroupPtrOutput `pulumi:"group"`
	// Image id
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// Instance name
	Name pulumi.StringOutput `pulumi:"name"`
	// Create network interfaces
	Network InstanceNetworkOutput `pulumi:"network"`
	// Instance region
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Existing SSH Keypair
	SshKey InstanceSshKeyPtrOutput `pulumi:"sshKey"`
	// Add existing SSH Key pair into your Public Cloud project and link it to the instance
	SshKeyCreate InstanceSshKeyCreatePtrOutput `pulumi:"sshKeyCreate"`
	// Instance task state
	TaskState pulumi.StringOutput `pulumi:"taskState"`
	// Configuration information or scripts to use upon launch
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingPeriod == nil {
		return nil, errors.New("invalid value for required argument 'BillingPeriod'")
	}
	if args.BootFrom == nil {
		return nil, errors.New("invalid value for required argument 'BootFrom'")
	}
	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("ovh:CloudProject/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("ovh:CloudProject/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Instance IP addresses
	Addresses []InstanceAddress `pulumi:"addresses"`
	// Volumes attached to the instance
	AttachedVolumes []InstanceAttachedVolume `pulumi:"attachedVolumes"`
	// Create an autobackup workflow after instance start up.
	AutoBackup *InstanceAutoBackup `pulumi:"autoBackup"`
	// The availability zone where the instance will be created
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Billing period - hourly or monthly
	BillingPeriod *string `pulumi:"billingPeriod"`
	// Boot the instance from an image or a volume
	BootFrom *InstanceBootFrom `pulumi:"bootFrom"`
	// Create multiple instances
	Bulk *int `pulumi:"bulk"`
	// Flavor information
	Flavor *InstanceFlavor `pulumi:"flavor"`
	// Flavor id
	FlavorId *string `pulumi:"flavorId"`
	// Flavor name
	FlavorName *string `pulumi:"flavorName"`
	// Start instance in group
	Group *InstanceGroup `pulumi:"group"`
	// Image id
	ImageId *string `pulumi:"imageId"`
	// Instance name
	Name *string `pulumi:"name"`
	// Create network interfaces
	Network *InstanceNetwork `pulumi:"network"`
	// Instance region
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
	ServiceName *string `pulumi:"serviceName"`
	// Existing SSH Keypair
	SshKey *InstanceSshKey `pulumi:"sshKey"`
	// Add existing SSH Key pair into your Public Cloud project and link it to the instance
	SshKeyCreate *InstanceSshKeyCreate `pulumi:"sshKeyCreate"`
	// Instance task state
	TaskState *string `pulumi:"taskState"`
	// Configuration information or scripts to use upon launch
	UserData *string `pulumi:"userData"`
}

type InstanceState struct {
	// Instance IP addresses
	Addresses InstanceAddressArrayInput
	// Volumes attached to the instance
	AttachedVolumes InstanceAttachedVolumeArrayInput
	// Create an autobackup workflow after instance start up.
	AutoBackup InstanceAutoBackupPtrInput
	// The availability zone where the instance will be created
	AvailabilityZone pulumi.StringPtrInput
	// Billing period - hourly or monthly
	BillingPeriod pulumi.StringPtrInput
	// Boot the instance from an image or a volume
	BootFrom InstanceBootFromPtrInput
	// Create multiple instances
	Bulk pulumi.IntPtrInput
	// Flavor information
	Flavor InstanceFlavorPtrInput
	// Flavor id
	FlavorId pulumi.StringPtrInput
	// Flavor name
	FlavorName pulumi.StringPtrInput
	// Start instance in group
	Group InstanceGroupPtrInput
	// Image id
	ImageId pulumi.StringPtrInput
	// Instance name
	Name pulumi.StringPtrInput
	// Create network interfaces
	Network InstanceNetworkPtrInput
	// Instance region
	Region pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
	ServiceName pulumi.StringPtrInput
	// Existing SSH Keypair
	SshKey InstanceSshKeyPtrInput
	// Add existing SSH Key pair into your Public Cloud project and link it to the instance
	SshKeyCreate InstanceSshKeyCreatePtrInput
	// Instance task state
	TaskState pulumi.StringPtrInput
	// Configuration information or scripts to use upon launch
	UserData pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Create an autobackup workflow after instance start up.
	AutoBackup *InstanceAutoBackup `pulumi:"autoBackup"`
	// The availability zone where the instance will be created
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Billing period - hourly or monthly
	BillingPeriod string `pulumi:"billingPeriod"`
	// Boot the instance from an image or a volume
	BootFrom InstanceBootFrom `pulumi:"bootFrom"`
	// Create multiple instances
	Bulk *int `pulumi:"bulk"`
	// Flavor information
	Flavor InstanceFlavor `pulumi:"flavor"`
	// Start instance in group
	Group *InstanceGroup `pulumi:"group"`
	// Instance name
	Name *string `pulumi:"name"`
	// Create network interfaces
	Network InstanceNetwork `pulumi:"network"`
	// Instance region
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
	ServiceName string `pulumi:"serviceName"`
	// Existing SSH Keypair
	SshKey *InstanceSshKey `pulumi:"sshKey"`
	// Add existing SSH Key pair into your Public Cloud project and link it to the instance
	SshKeyCreate *InstanceSshKeyCreate `pulumi:"sshKeyCreate"`
	// Configuration information or scripts to use upon launch
	UserData *string `pulumi:"userData"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Create an autobackup workflow after instance start up.
	AutoBackup InstanceAutoBackupPtrInput
	// The availability zone where the instance will be created
	AvailabilityZone pulumi.StringPtrInput
	// Billing period - hourly or monthly
	BillingPeriod pulumi.StringInput
	// Boot the instance from an image or a volume
	BootFrom InstanceBootFromInput
	// Create multiple instances
	Bulk pulumi.IntPtrInput
	// Flavor information
	Flavor InstanceFlavorInput
	// Start instance in group
	Group InstanceGroupPtrInput
	// Instance name
	Name pulumi.StringPtrInput
	// Create network interfaces
	Network InstanceNetworkInput
	// Instance region
	Region pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
	ServiceName pulumi.StringInput
	// Existing SSH Keypair
	SshKey InstanceSshKeyPtrInput
	// Add existing SSH Key pair into your Public Cloud project and link it to the instance
	SshKeyCreate InstanceSshKeyCreatePtrInput
	// Configuration information or scripts to use upon launch
	UserData pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Instance IP addresses
func (o InstanceOutput) Addresses() InstanceAddressArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceAddressArrayOutput { return v.Addresses }).(InstanceAddressArrayOutput)
}

// Volumes attached to the instance
func (o InstanceOutput) AttachedVolumes() InstanceAttachedVolumeArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceAttachedVolumeArrayOutput { return v.AttachedVolumes }).(InstanceAttachedVolumeArrayOutput)
}

// Create an autobackup workflow after instance start up.
func (o InstanceOutput) AutoBackup() InstanceAutoBackupPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceAutoBackupPtrOutput { return v.AutoBackup }).(InstanceAutoBackupPtrOutput)
}

// The availability zone where the instance will be created
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Billing period - hourly or monthly
func (o InstanceOutput) BillingPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.BillingPeriod }).(pulumi.StringOutput)
}

// Boot the instance from an image or a volume
func (o InstanceOutput) BootFrom() InstanceBootFromOutput {
	return o.ApplyT(func(v *Instance) InstanceBootFromOutput { return v.BootFrom }).(InstanceBootFromOutput)
}

// Create multiple instances
func (o InstanceOutput) Bulk() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Bulk }).(pulumi.IntPtrOutput)
}

// Flavor information
func (o InstanceOutput) Flavor() InstanceFlavorOutput {
	return o.ApplyT(func(v *Instance) InstanceFlavorOutput { return v.Flavor }).(InstanceFlavorOutput)
}

// Flavor id
func (o InstanceOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.FlavorId }).(pulumi.StringOutput)
}

// Flavor name
func (o InstanceOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.FlavorName }).(pulumi.StringOutput)
}

// Start instance in group
func (o InstanceOutput) Group() InstanceGroupPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceGroupPtrOutput { return v.Group }).(InstanceGroupPtrOutput)
}

// Image id
func (o InstanceOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// Instance name
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Create network interfaces
func (o InstanceOutput) Network() InstanceNetworkOutput {
	return o.ApplyT(func(v *Instance) InstanceNetworkOutput { return v.Network }).(InstanceNetworkOutput)
}

// Instance region
func (o InstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
func (o InstanceOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Existing SSH Keypair
func (o InstanceOutput) SshKey() InstanceSshKeyPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceSshKeyPtrOutput { return v.SshKey }).(InstanceSshKeyPtrOutput)
}

// Add existing SSH Key pair into your Public Cloud project and link it to the instance
func (o InstanceOutput) SshKeyCreate() InstanceSshKeyCreatePtrOutput {
	return o.ApplyT(func(v *Instance) InstanceSshKeyCreatePtrOutput { return v.SshKeyCreate }).(InstanceSshKeyCreatePtrOutput)
}

// Instance task state
func (o InstanceOutput) TaskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TaskState }).(pulumi.StringOutput)
}

// Configuration information or scripts to use upon launch
func (o InstanceOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
