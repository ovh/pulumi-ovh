// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a Dedicated Cloud to the vrack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/vrack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vrack.NewVrackDedicatedCloud(ctx, "vrack-dedicatedCloud", &vrack.VrackDedicatedCloudArgs{
//				ServiceName:    pulumi.String("<vRack service name>"),
//				DedicatedCloud: pulumi.String("<Dedicated Cloud service name>"),
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
// Attachment of a Dedicated Cloud and a vRack can be imported using the `service_name` (vRack identifier) and the `dedicated_cloud` (Dedicated Cloud service name), separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Vrack/vrackDedicatedCloud:VrackDedicatedCloud myattach "<vRack service name>/<Dedicated Cloud service name>"
// ```
type VrackDedicatedCloud struct {
	pulumi.CustomResourceState

	// Your Dedicated Cloud service name
	DedicatedCloud pulumi.StringOutput `pulumi:"dedicatedCloud"`
	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewVrackDedicatedCloud registers a new resource with the given unique name, arguments, and options.
func NewVrackDedicatedCloud(ctx *pulumi.Context,
	name string, args *VrackDedicatedCloudArgs, opts ...pulumi.ResourceOption) (*VrackDedicatedCloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DedicatedCloud == nil {
		return nil, errors.New("invalid value for required argument 'DedicatedCloud'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VrackDedicatedCloud
	err := ctx.RegisterResource("ovh:Vrack/vrackDedicatedCloud:VrackDedicatedCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackDedicatedCloud gets an existing VrackDedicatedCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackDedicatedCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackDedicatedCloudState, opts ...pulumi.ResourceOption) (*VrackDedicatedCloud, error) {
	var resource VrackDedicatedCloud
	err := ctx.ReadResource("ovh:Vrack/vrackDedicatedCloud:VrackDedicatedCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackDedicatedCloud resources.
type vrackDedicatedCloudState struct {
	// Your Dedicated Cloud service name
	DedicatedCloud *string `pulumi:"dedicatedCloud"`
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
}

type VrackDedicatedCloudState struct {
	// Your Dedicated Cloud service name
	DedicatedCloud pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
}

func (VrackDedicatedCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedCloudState)(nil)).Elem()
}

type vrackDedicatedCloudArgs struct {
	// Your Dedicated Cloud service name
	DedicatedCloud string `pulumi:"dedicatedCloud"`
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VrackDedicatedCloud resource.
type VrackDedicatedCloudArgs struct {
	// Your Dedicated Cloud service name
	DedicatedCloud pulumi.StringInput
	// The internal name of your vrack
	ServiceName pulumi.StringInput
}

func (VrackDedicatedCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedCloudArgs)(nil)).Elem()
}

type VrackDedicatedCloudInput interface {
	pulumi.Input

	ToVrackDedicatedCloudOutput() VrackDedicatedCloudOutput
	ToVrackDedicatedCloudOutputWithContext(ctx context.Context) VrackDedicatedCloudOutput
}

func (*VrackDedicatedCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedCloud)(nil)).Elem()
}

func (i *VrackDedicatedCloud) ToVrackDedicatedCloudOutput() VrackDedicatedCloudOutput {
	return i.ToVrackDedicatedCloudOutputWithContext(context.Background())
}

func (i *VrackDedicatedCloud) ToVrackDedicatedCloudOutputWithContext(ctx context.Context) VrackDedicatedCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedCloudOutput)
}

// VrackDedicatedCloudArrayInput is an input type that accepts VrackDedicatedCloudArray and VrackDedicatedCloudArrayOutput values.
// You can construct a concrete instance of `VrackDedicatedCloudArrayInput` via:
//
//	VrackDedicatedCloudArray{ VrackDedicatedCloudArgs{...} }
type VrackDedicatedCloudArrayInput interface {
	pulumi.Input

	ToVrackDedicatedCloudArrayOutput() VrackDedicatedCloudArrayOutput
	ToVrackDedicatedCloudArrayOutputWithContext(context.Context) VrackDedicatedCloudArrayOutput
}

type VrackDedicatedCloudArray []VrackDedicatedCloudInput

func (VrackDedicatedCloudArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackDedicatedCloud)(nil)).Elem()
}

func (i VrackDedicatedCloudArray) ToVrackDedicatedCloudArrayOutput() VrackDedicatedCloudArrayOutput {
	return i.ToVrackDedicatedCloudArrayOutputWithContext(context.Background())
}

func (i VrackDedicatedCloudArray) ToVrackDedicatedCloudArrayOutputWithContext(ctx context.Context) VrackDedicatedCloudArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedCloudArrayOutput)
}

// VrackDedicatedCloudMapInput is an input type that accepts VrackDedicatedCloudMap and VrackDedicatedCloudMapOutput values.
// You can construct a concrete instance of `VrackDedicatedCloudMapInput` via:
//
//	VrackDedicatedCloudMap{ "key": VrackDedicatedCloudArgs{...} }
type VrackDedicatedCloudMapInput interface {
	pulumi.Input

	ToVrackDedicatedCloudMapOutput() VrackDedicatedCloudMapOutput
	ToVrackDedicatedCloudMapOutputWithContext(context.Context) VrackDedicatedCloudMapOutput
}

type VrackDedicatedCloudMap map[string]VrackDedicatedCloudInput

func (VrackDedicatedCloudMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackDedicatedCloud)(nil)).Elem()
}

func (i VrackDedicatedCloudMap) ToVrackDedicatedCloudMapOutput() VrackDedicatedCloudMapOutput {
	return i.ToVrackDedicatedCloudMapOutputWithContext(context.Background())
}

func (i VrackDedicatedCloudMap) ToVrackDedicatedCloudMapOutputWithContext(ctx context.Context) VrackDedicatedCloudMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedCloudMapOutput)
}

type VrackDedicatedCloudOutput struct{ *pulumi.OutputState }

func (VrackDedicatedCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedCloud)(nil)).Elem()
}

func (o VrackDedicatedCloudOutput) ToVrackDedicatedCloudOutput() VrackDedicatedCloudOutput {
	return o
}

func (o VrackDedicatedCloudOutput) ToVrackDedicatedCloudOutputWithContext(ctx context.Context) VrackDedicatedCloudOutput {
	return o
}

// Your Dedicated Cloud service name
func (o VrackDedicatedCloudOutput) DedicatedCloud() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackDedicatedCloud) pulumi.StringOutput { return v.DedicatedCloud }).(pulumi.StringOutput)
}

// The internal name of your vrack
func (o VrackDedicatedCloudOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackDedicatedCloud) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type VrackDedicatedCloudArrayOutput struct{ *pulumi.OutputState }

func (VrackDedicatedCloudArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackDedicatedCloud)(nil)).Elem()
}

func (o VrackDedicatedCloudArrayOutput) ToVrackDedicatedCloudArrayOutput() VrackDedicatedCloudArrayOutput {
	return o
}

func (o VrackDedicatedCloudArrayOutput) ToVrackDedicatedCloudArrayOutputWithContext(ctx context.Context) VrackDedicatedCloudArrayOutput {
	return o
}

func (o VrackDedicatedCloudArrayOutput) Index(i pulumi.IntInput) VrackDedicatedCloudOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VrackDedicatedCloud {
		return vs[0].([]*VrackDedicatedCloud)[vs[1].(int)]
	}).(VrackDedicatedCloudOutput)
}

type VrackDedicatedCloudMapOutput struct{ *pulumi.OutputState }

func (VrackDedicatedCloudMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackDedicatedCloud)(nil)).Elem()
}

func (o VrackDedicatedCloudMapOutput) ToVrackDedicatedCloudMapOutput() VrackDedicatedCloudMapOutput {
	return o
}

func (o VrackDedicatedCloudMapOutput) ToVrackDedicatedCloudMapOutputWithContext(ctx context.Context) VrackDedicatedCloudMapOutput {
	return o
}

func (o VrackDedicatedCloudMapOutput) MapIndex(k pulumi.StringInput) VrackDedicatedCloudOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VrackDedicatedCloud {
		return vs[0].(map[string]*VrackDedicatedCloud)[vs[1].(string)]
	}).(VrackDedicatedCloudOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrackDedicatedCloudInput)(nil)).Elem(), &VrackDedicatedCloud{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackDedicatedCloudArrayInput)(nil)).Elem(), VrackDedicatedCloudArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackDedicatedCloudMapInput)(nil)).Elem(), VrackDedicatedCloudMap{})
	pulumi.RegisterOutputType(VrackDedicatedCloudOutput{})
	pulumi.RegisterOutputType(VrackDedicatedCloudArrayOutput{})
	pulumi.RegisterOutputType(VrackDedicatedCloudMapOutput{})
}
