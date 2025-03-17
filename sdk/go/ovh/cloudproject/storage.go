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

// Create S3™* compatible storage container
// (* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.)
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
//			_, err := cloudproject.NewStorage(ctx, "storage", &cloudproject.StorageArgs{
//				RegionName:  pulumi.String("GRA"),
//				ServiceName: pulumi.String("<public cloud project ID>"),
//				Versioning: &cloudproject.StorageVersioningArgs{
//					Status: pulumi.String("enabled"),
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
// A storage in a public cloud project can be imported using the `service_name`, `region_name` and `name` attributes.
//
// Using the following configuration:
//
// hcl
//
// import {
//
//	id = "<service_name>/<region_name>/<name>"
//
//	to = ovh_cloud_project_storage.storage
//
// }
//
// You can then run:
//
// bash
//
// $ pulumi preview -generate-config-out=storage.tf
//
// $ pulumi up
//
// The file `storage.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.
//
// See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type Storage struct {
	pulumi.CustomResourceState

	// The date and timestamp when the resource was created
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Encryption configuration
	Encryption StorageEncryptionOutput `pulumi:"encryption"`
	// Limit the number of objects returned (1000 maximum, defaults to 1000)
	Limit pulumi.Float64Output `pulumi:"limit"`
	// Key to start with when listing objects
	Marker pulumi.StringOutput `pulumi:"marker"`
	// Container name
	Name pulumi.StringOutput `pulumi:"name"`
	// Container objects
	Objects StorageObjectArrayOutput `pulumi:"objects"`
	// Container total objects count
	ObjectsCount pulumi.Float64Output `pulumi:"objectsCount"`
	// Container total objects size (bytes)
	ObjectsSize pulumi.Float64Output `pulumi:"objectsSize"`
	// Container owner user ID
	OwnerId pulumi.Float64Output `pulumi:"ownerId"`
	// List objects whose key begins with this prefix
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// Container region
	Region pulumi.StringOutput `pulumi:"region"`
	// Region name
	RegionName pulumi.StringOutput `pulumi:"regionName"`
	// Replication configuration
	Replication StorageReplicationOutput `pulumi:"replication"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Versioning configuration
	Versioning StorageVersioningOutput `pulumi:"versioning"`
	// Container virtual host
	VirtualHost pulumi.StringOutput `pulumi:"virtualHost"`
}

// NewStorage registers a new resource with the given unique name, arguments, and options.
func NewStorage(ctx *pulumi.Context,
	name string, args *StorageArgs, opts ...pulumi.ResourceOption) (*Storage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegionName == nil {
		return nil, errors.New("invalid value for required argument 'RegionName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Storage
	err := ctx.RegisterResource("ovh:CloudProject/storage:Storage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorage gets an existing Storage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageState, opts ...pulumi.ResourceOption) (*Storage, error) {
	var resource Storage
	err := ctx.ReadResource("ovh:CloudProject/storage:Storage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Storage resources.
type storageState struct {
	// The date and timestamp when the resource was created
	CreatedAt *string `pulumi:"createdAt"`
	// Encryption configuration
	Encryption *StorageEncryption `pulumi:"encryption"`
	// Limit the number of objects returned (1000 maximum, defaults to 1000)
	Limit *float64 `pulumi:"limit"`
	// Key to start with when listing objects
	Marker *string `pulumi:"marker"`
	// Container name
	Name *string `pulumi:"name"`
	// Container objects
	Objects []StorageObject `pulumi:"objects"`
	// Container total objects count
	ObjectsCount *float64 `pulumi:"objectsCount"`
	// Container total objects size (bytes)
	ObjectsSize *float64 `pulumi:"objectsSize"`
	// Container owner user ID
	OwnerId *float64 `pulumi:"ownerId"`
	// List objects whose key begins with this prefix
	Prefix *string `pulumi:"prefix"`
	// Container region
	Region *string `pulumi:"region"`
	// Region name
	RegionName *string `pulumi:"regionName"`
	// Replication configuration
	Replication *StorageReplication `pulumi:"replication"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Versioning configuration
	Versioning *StorageVersioning `pulumi:"versioning"`
	// Container virtual host
	VirtualHost *string `pulumi:"virtualHost"`
}

type StorageState struct {
	// The date and timestamp when the resource was created
	CreatedAt pulumi.StringPtrInput
	// Encryption configuration
	Encryption StorageEncryptionPtrInput
	// Limit the number of objects returned (1000 maximum, defaults to 1000)
	Limit pulumi.Float64PtrInput
	// Key to start with when listing objects
	Marker pulumi.StringPtrInput
	// Container name
	Name pulumi.StringPtrInput
	// Container objects
	Objects StorageObjectArrayInput
	// Container total objects count
	ObjectsCount pulumi.Float64PtrInput
	// Container total objects size (bytes)
	ObjectsSize pulumi.Float64PtrInput
	// Container owner user ID
	OwnerId pulumi.Float64PtrInput
	// List objects whose key begins with this prefix
	Prefix pulumi.StringPtrInput
	// Container region
	Region pulumi.StringPtrInput
	// Region name
	RegionName pulumi.StringPtrInput
	// Replication configuration
	Replication StorageReplicationPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Versioning configuration
	Versioning StorageVersioningPtrInput
	// Container virtual host
	VirtualHost pulumi.StringPtrInput
}

func (StorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageState)(nil)).Elem()
}

type storageArgs struct {
	// Encryption configuration
	Encryption *StorageEncryption `pulumi:"encryption"`
	// Limit the number of objects returned (1000 maximum, defaults to 1000)
	Limit *float64 `pulumi:"limit"`
	// Key to start with when listing objects
	Marker *string `pulumi:"marker"`
	// Container name
	Name *string `pulumi:"name"`
	// Container owner user ID
	OwnerId *float64 `pulumi:"ownerId"`
	// List objects whose key begins with this prefix
	Prefix *string `pulumi:"prefix"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Replication configuration
	Replication *StorageReplication `pulumi:"replication"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
	// Versioning configuration
	Versioning *StorageVersioning `pulumi:"versioning"`
}

// The set of arguments for constructing a Storage resource.
type StorageArgs struct {
	// Encryption configuration
	Encryption StorageEncryptionPtrInput
	// Limit the number of objects returned (1000 maximum, defaults to 1000)
	Limit pulumi.Float64PtrInput
	// Key to start with when listing objects
	Marker pulumi.StringPtrInput
	// Container name
	Name pulumi.StringPtrInput
	// Container owner user ID
	OwnerId pulumi.Float64PtrInput
	// List objects whose key begins with this prefix
	Prefix pulumi.StringPtrInput
	// Region name
	RegionName pulumi.StringInput
	// Replication configuration
	Replication StorageReplicationPtrInput
	// Service name
	ServiceName pulumi.StringInput
	// Versioning configuration
	Versioning StorageVersioningPtrInput
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageArgs)(nil)).Elem()
}

type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(ctx context.Context) StorageOutput
}

func (*Storage) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *Storage) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i *Storage) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

// StorageArrayInput is an input type that accepts StorageArray and StorageArrayOutput values.
// You can construct a concrete instance of `StorageArrayInput` via:
//
//	StorageArray{ StorageArgs{...} }
type StorageArrayInput interface {
	pulumi.Input

	ToStorageArrayOutput() StorageArrayOutput
	ToStorageArrayOutputWithContext(context.Context) StorageArrayOutput
}

type StorageArray []StorageInput

func (StorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Storage)(nil)).Elem()
}

func (i StorageArray) ToStorageArrayOutput() StorageArrayOutput {
	return i.ToStorageArrayOutputWithContext(context.Background())
}

func (i StorageArray) ToStorageArrayOutputWithContext(ctx context.Context) StorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageArrayOutput)
}

// StorageMapInput is an input type that accepts StorageMap and StorageMapOutput values.
// You can construct a concrete instance of `StorageMapInput` via:
//
//	StorageMap{ "key": StorageArgs{...} }
type StorageMapInput interface {
	pulumi.Input

	ToStorageMapOutput() StorageMapOutput
	ToStorageMapOutputWithContext(context.Context) StorageMapOutput
}

type StorageMap map[string]StorageInput

func (StorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Storage)(nil)).Elem()
}

func (i StorageMap) ToStorageMapOutput() StorageMapOutput {
	return i.ToStorageMapOutputWithContext(context.Background())
}

func (i StorageMap) ToStorageMapOutputWithContext(ctx context.Context) StorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMapOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

// The date and timestamp when the resource was created
func (o StorageOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Encryption configuration
func (o StorageOutput) Encryption() StorageEncryptionOutput {
	return o.ApplyT(func(v *Storage) StorageEncryptionOutput { return v.Encryption }).(StorageEncryptionOutput)
}

// Limit the number of objects returned (1000 maximum, defaults to 1000)
func (o StorageOutput) Limit() pulumi.Float64Output {
	return o.ApplyT(func(v *Storage) pulumi.Float64Output { return v.Limit }).(pulumi.Float64Output)
}

// Key to start with when listing objects
func (o StorageOutput) Marker() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Marker }).(pulumi.StringOutput)
}

// Container name
func (o StorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Container objects
func (o StorageOutput) Objects() StorageObjectArrayOutput {
	return o.ApplyT(func(v *Storage) StorageObjectArrayOutput { return v.Objects }).(StorageObjectArrayOutput)
}

// Container total objects count
func (o StorageOutput) ObjectsCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Storage) pulumi.Float64Output { return v.ObjectsCount }).(pulumi.Float64Output)
}

// Container total objects size (bytes)
func (o StorageOutput) ObjectsSize() pulumi.Float64Output {
	return o.ApplyT(func(v *Storage) pulumi.Float64Output { return v.ObjectsSize }).(pulumi.Float64Output)
}

// Container owner user ID
func (o StorageOutput) OwnerId() pulumi.Float64Output {
	return o.ApplyT(func(v *Storage) pulumi.Float64Output { return v.OwnerId }).(pulumi.Float64Output)
}

// List objects whose key begins with this prefix
func (o StorageOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// Container region
func (o StorageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Region name
func (o StorageOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.RegionName }).(pulumi.StringOutput)
}

// Replication configuration
func (o StorageOutput) Replication() StorageReplicationOutput {
	return o.ApplyT(func(v *Storage) StorageReplicationOutput { return v.Replication }).(StorageReplicationOutput)
}

// Service name
func (o StorageOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Versioning configuration
func (o StorageOutput) Versioning() StorageVersioningOutput {
	return o.ApplyT(func(v *Storage) StorageVersioningOutput { return v.Versioning }).(StorageVersioningOutput)
}

// Container virtual host
func (o StorageOutput) VirtualHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.VirtualHost }).(pulumi.StringOutput)
}

type StorageArrayOutput struct{ *pulumi.OutputState }

func (StorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Storage)(nil)).Elem()
}

func (o StorageArrayOutput) ToStorageArrayOutput() StorageArrayOutput {
	return o
}

func (o StorageArrayOutput) ToStorageArrayOutputWithContext(ctx context.Context) StorageArrayOutput {
	return o
}

func (o StorageArrayOutput) Index(i pulumi.IntInput) StorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Storage {
		return vs[0].([]*Storage)[vs[1].(int)]
	}).(StorageOutput)
}

type StorageMapOutput struct{ *pulumi.OutputState }

func (StorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Storage)(nil)).Elem()
}

func (o StorageMapOutput) ToStorageMapOutput() StorageMapOutput {
	return o
}

func (o StorageMapOutput) ToStorageMapOutputWithContext(ctx context.Context) StorageMapOutput {
	return o
}

func (o StorageMapOutput) MapIndex(k pulumi.StringInput) StorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Storage {
		return vs[0].(map[string]*Storage)[vs[1].(string)]
	}).(StorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageInput)(nil)).Elem(), &Storage{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageArrayInput)(nil)).Elem(), StorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageMapInput)(nil)).Elem(), StorageMap{})
	pulumi.RegisterOutputType(StorageOutput{})
	pulumi.RegisterOutputType(StorageArrayOutput{})
	pulumi.RegisterOutputType(StorageMapOutput{})
}
