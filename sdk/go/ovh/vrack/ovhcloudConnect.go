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

// Attach an OVH Cloud Connect to the vrack.
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
//			_, err := vrack.NewOVHcloudConnect(ctx, "vrack_ovhcloudconnect", &vrack.OVHcloudConnectArgs{
//				ServiceName:     pulumi.String("<vRack service name>"),
//				OvhCloudConnect: pulumi.String("<OVH Cloud Connect service name>"),
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
// Attachment of an OVH Cloud Connect and a vRack can be imported using the `service_name` (vRack identifier) and the `ovh_cloud_connect` (OVH Cloud Connect service name), separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Vrack/oVHcloudConnect:OVHcloudConnect myattach "<service_name>/<OVH Cloud Connect service name>"
// ```
type OVHcloudConnect struct {
	pulumi.CustomResourceState

	// Your OVH Cloud Connect service name.
	OvhCloudConnect pulumi.StringOutput `pulumi:"ovhCloudConnect"`
	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewOVHcloudConnect registers a new resource with the given unique name, arguments, and options.
func NewOVHcloudConnect(ctx *pulumi.Context,
	name string, args *OVHcloudConnectArgs, opts ...pulumi.ResourceOption) (*OVHcloudConnect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OvhCloudConnect == nil {
		return nil, errors.New("invalid value for required argument 'OvhCloudConnect'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OVHcloudConnect
	err := ctx.RegisterResource("ovh:Vrack/oVHcloudConnect:OVHcloudConnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOVHcloudConnect gets an existing OVHcloudConnect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOVHcloudConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OVHcloudConnectState, opts ...pulumi.ResourceOption) (*OVHcloudConnect, error) {
	var resource OVHcloudConnect
	err := ctx.ReadResource("ovh:Vrack/oVHcloudConnect:OVHcloudConnect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OVHcloudConnect resources.
type ovhcloudConnectState struct {
	// Your OVH Cloud Connect service name.
	OvhCloudConnect *string `pulumi:"ovhCloudConnect"`
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
}

type OVHcloudConnectState struct {
	// Your OVH Cloud Connect service name.
	OvhCloudConnect pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
}

func (OVHcloudConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovhcloudConnectState)(nil)).Elem()
}

type ovhcloudConnectArgs struct {
	// Your OVH Cloud Connect service name.
	OvhCloudConnect string `pulumi:"ovhCloudConnect"`
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a OVHcloudConnect resource.
type OVHcloudConnectArgs struct {
	// Your OVH Cloud Connect service name.
	OvhCloudConnect pulumi.StringInput
	// The internal name of your vrack
	ServiceName pulumi.StringInput
}

func (OVHcloudConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovhcloudConnectArgs)(nil)).Elem()
}

type OVHcloudConnectInput interface {
	pulumi.Input

	ToOVHcloudConnectOutput() OVHcloudConnectOutput
	ToOVHcloudConnectOutputWithContext(ctx context.Context) OVHcloudConnectOutput
}

func (*OVHcloudConnect) ElementType() reflect.Type {
	return reflect.TypeOf((**OVHcloudConnect)(nil)).Elem()
}

func (i *OVHcloudConnect) ToOVHcloudConnectOutput() OVHcloudConnectOutput {
	return i.ToOVHcloudConnectOutputWithContext(context.Background())
}

func (i *OVHcloudConnect) ToOVHcloudConnectOutputWithContext(ctx context.Context) OVHcloudConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OVHcloudConnectOutput)
}

// OVHcloudConnectArrayInput is an input type that accepts OVHcloudConnectArray and OVHcloudConnectArrayOutput values.
// You can construct a concrete instance of `OVHcloudConnectArrayInput` via:
//
//	OVHcloudConnectArray{ OVHcloudConnectArgs{...} }
type OVHcloudConnectArrayInput interface {
	pulumi.Input

	ToOVHcloudConnectArrayOutput() OVHcloudConnectArrayOutput
	ToOVHcloudConnectArrayOutputWithContext(context.Context) OVHcloudConnectArrayOutput
}

type OVHcloudConnectArray []OVHcloudConnectInput

func (OVHcloudConnectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OVHcloudConnect)(nil)).Elem()
}

func (i OVHcloudConnectArray) ToOVHcloudConnectArrayOutput() OVHcloudConnectArrayOutput {
	return i.ToOVHcloudConnectArrayOutputWithContext(context.Background())
}

func (i OVHcloudConnectArray) ToOVHcloudConnectArrayOutputWithContext(ctx context.Context) OVHcloudConnectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OVHcloudConnectArrayOutput)
}

// OVHcloudConnectMapInput is an input type that accepts OVHcloudConnectMap and OVHcloudConnectMapOutput values.
// You can construct a concrete instance of `OVHcloudConnectMapInput` via:
//
//	OVHcloudConnectMap{ "key": OVHcloudConnectArgs{...} }
type OVHcloudConnectMapInput interface {
	pulumi.Input

	ToOVHcloudConnectMapOutput() OVHcloudConnectMapOutput
	ToOVHcloudConnectMapOutputWithContext(context.Context) OVHcloudConnectMapOutput
}

type OVHcloudConnectMap map[string]OVHcloudConnectInput

func (OVHcloudConnectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OVHcloudConnect)(nil)).Elem()
}

func (i OVHcloudConnectMap) ToOVHcloudConnectMapOutput() OVHcloudConnectMapOutput {
	return i.ToOVHcloudConnectMapOutputWithContext(context.Background())
}

func (i OVHcloudConnectMap) ToOVHcloudConnectMapOutputWithContext(ctx context.Context) OVHcloudConnectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OVHcloudConnectMapOutput)
}

type OVHcloudConnectOutput struct{ *pulumi.OutputState }

func (OVHcloudConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OVHcloudConnect)(nil)).Elem()
}

func (o OVHcloudConnectOutput) ToOVHcloudConnectOutput() OVHcloudConnectOutput {
	return o
}

func (o OVHcloudConnectOutput) ToOVHcloudConnectOutputWithContext(ctx context.Context) OVHcloudConnectOutput {
	return o
}

// Your OVH Cloud Connect service name.
func (o OVHcloudConnectOutput) OvhCloudConnect() pulumi.StringOutput {
	return o.ApplyT(func(v *OVHcloudConnect) pulumi.StringOutput { return v.OvhCloudConnect }).(pulumi.StringOutput)
}

// The internal name of your vrack
func (o OVHcloudConnectOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *OVHcloudConnect) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type OVHcloudConnectArrayOutput struct{ *pulumi.OutputState }

func (OVHcloudConnectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OVHcloudConnect)(nil)).Elem()
}

func (o OVHcloudConnectArrayOutput) ToOVHcloudConnectArrayOutput() OVHcloudConnectArrayOutput {
	return o
}

func (o OVHcloudConnectArrayOutput) ToOVHcloudConnectArrayOutputWithContext(ctx context.Context) OVHcloudConnectArrayOutput {
	return o
}

func (o OVHcloudConnectArrayOutput) Index(i pulumi.IntInput) OVHcloudConnectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OVHcloudConnect {
		return vs[0].([]*OVHcloudConnect)[vs[1].(int)]
	}).(OVHcloudConnectOutput)
}

type OVHcloudConnectMapOutput struct{ *pulumi.OutputState }

func (OVHcloudConnectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OVHcloudConnect)(nil)).Elem()
}

func (o OVHcloudConnectMapOutput) ToOVHcloudConnectMapOutput() OVHcloudConnectMapOutput {
	return o
}

func (o OVHcloudConnectMapOutput) ToOVHcloudConnectMapOutputWithContext(ctx context.Context) OVHcloudConnectMapOutput {
	return o
}

func (o OVHcloudConnectMapOutput) MapIndex(k pulumi.StringInput) OVHcloudConnectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OVHcloudConnect {
		return vs[0].(map[string]*OVHcloudConnect)[vs[1].(string)]
	}).(OVHcloudConnectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OVHcloudConnectInput)(nil)).Elem(), &OVHcloudConnect{})
	pulumi.RegisterInputType(reflect.TypeOf((*OVHcloudConnectArrayInput)(nil)).Elem(), OVHcloudConnectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OVHcloudConnectMapInput)(nil)).Elem(), OVHcloudConnectMap{})
	pulumi.RegisterOutputType(OVHcloudConnectOutput{})
	pulumi.RegisterOutputType(OVHcloudConnectArrayOutput{})
	pulumi.RegisterOutputType(OVHcloudConnectMapOutput{})
}
