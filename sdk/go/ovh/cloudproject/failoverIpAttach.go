// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a failover IP address to a compute instance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudproject.NewFailoverIpAttach(ctx, "myFailoverIp", &cloudproject.FailoverIpAttachArgs{
//				Ip:          pulumi.String("XXXXXX"),
//				RoutedTo:    pulumi.String("XXXXXX"),
//				ServiceName: pulumi.String("XXXXXX"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FailoverIpAttach struct {
	pulumi.CustomResourceState

	// The IP block
	// * `continentCode` - The Ip continent
	Block pulumi.StringOutput `pulumi:"block"`
	// Ip continent
	ContinentCode pulumi.StringOutput `pulumi:"continentCode"`
	// Ip location
	GeoLoc pulumi.StringOutput `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Current operation progress in percent
	// * `routedTo` - Instance where ip is routed to
	Progress pulumi.IntOutput `pulumi:"progress"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringOutput `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	// * `subType` - IP sub type, can be `cloud` or `ovh`
	Status pulumi.StringOutput `pulumi:"status"`
	// IP sub type
	SubType pulumi.StringOutput `pulumi:"subType"`
}

// NewFailoverIpAttach registers a new resource with the given unique name, arguments, and options.
func NewFailoverIpAttach(ctx *pulumi.Context,
	name string, args *FailoverIpAttachArgs, opts ...pulumi.ResourceOption) (*FailoverIpAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FailoverIpAttach
	err := ctx.RegisterResource("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverIpAttach gets an existing FailoverIpAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverIpAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverIpAttachState, opts ...pulumi.ResourceOption) (*FailoverIpAttach, error) {
	var resource FailoverIpAttach
	err := ctx.ReadResource("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverIpAttach resources.
type failoverIpAttachState struct {
	// The IP block
	// * `continentCode` - The Ip continent
	Block *string `pulumi:"block"`
	// Ip continent
	ContinentCode *string `pulumi:"continentCode"`
	// Ip location
	GeoLoc *string `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip *string `pulumi:"ip"`
	// Current operation progress in percent
	// * `routedTo` - Instance where ip is routed to
	Progress *int `pulumi:"progress"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	// * `subType` - IP sub type, can be `cloud` or `ovh`
	Status *string `pulumi:"status"`
	// IP sub type
	SubType *string `pulumi:"subType"`
}

type FailoverIpAttachState struct {
	// The IP block
	// * `continentCode` - The Ip continent
	Block pulumi.StringPtrInput
	// Ip continent
	ContinentCode pulumi.StringPtrInput
	// Ip location
	GeoLoc pulumi.StringPtrInput
	// The failover ip address to attach
	Ip pulumi.StringPtrInput
	// Current operation progress in percent
	// * `routedTo` - Instance where ip is routed to
	Progress pulumi.IntPtrInput
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Ip status, can be `ok` or `operationPending`
	// * `subType` - IP sub type, can be `cloud` or `ovh`
	Status pulumi.StringPtrInput
	// IP sub type
	SubType pulumi.StringPtrInput
}

func (FailoverIpAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverIpAttachState)(nil)).Elem()
}

type failoverIpAttachArgs struct {
	// The IP block
	// * `continentCode` - The Ip continent
	Block *string `pulumi:"block"`
	// Ip continent
	ContinentCode *string `pulumi:"continentCode"`
	// Ip location
	GeoLoc *string `pulumi:"geoLoc"`
	// The failover ip address to attach
	Ip *string `pulumi:"ip"`
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a FailoverIpAttach resource.
type FailoverIpAttachArgs struct {
	// The IP block
	// * `continentCode` - The Ip continent
	Block pulumi.StringPtrInput
	// Ip continent
	ContinentCode pulumi.StringPtrInput
	// Ip location
	GeoLoc pulumi.StringPtrInput
	// The failover ip address to attach
	Ip pulumi.StringPtrInput
	// The GUID of an instance to which the failover IP address is be attached
	RoutedTo pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (FailoverIpAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverIpAttachArgs)(nil)).Elem()
}

type FailoverIpAttachInput interface {
	pulumi.Input

	ToFailoverIpAttachOutput() FailoverIpAttachOutput
	ToFailoverIpAttachOutputWithContext(ctx context.Context) FailoverIpAttachOutput
}

func (*FailoverIpAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverIpAttach)(nil)).Elem()
}

func (i *FailoverIpAttach) ToFailoverIpAttachOutput() FailoverIpAttachOutput {
	return i.ToFailoverIpAttachOutputWithContext(context.Background())
}

func (i *FailoverIpAttach) ToFailoverIpAttachOutputWithContext(ctx context.Context) FailoverIpAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverIpAttachOutput)
}

// FailoverIpAttachArrayInput is an input type that accepts FailoverIpAttachArray and FailoverIpAttachArrayOutput values.
// You can construct a concrete instance of `FailoverIpAttachArrayInput` via:
//
//	FailoverIpAttachArray{ FailoverIpAttachArgs{...} }
type FailoverIpAttachArrayInput interface {
	pulumi.Input

	ToFailoverIpAttachArrayOutput() FailoverIpAttachArrayOutput
	ToFailoverIpAttachArrayOutputWithContext(context.Context) FailoverIpAttachArrayOutput
}

type FailoverIpAttachArray []FailoverIpAttachInput

func (FailoverIpAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FailoverIpAttach)(nil)).Elem()
}

func (i FailoverIpAttachArray) ToFailoverIpAttachArrayOutput() FailoverIpAttachArrayOutput {
	return i.ToFailoverIpAttachArrayOutputWithContext(context.Background())
}

func (i FailoverIpAttachArray) ToFailoverIpAttachArrayOutputWithContext(ctx context.Context) FailoverIpAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverIpAttachArrayOutput)
}

// FailoverIpAttachMapInput is an input type that accepts FailoverIpAttachMap and FailoverIpAttachMapOutput values.
// You can construct a concrete instance of `FailoverIpAttachMapInput` via:
//
//	FailoverIpAttachMap{ "key": FailoverIpAttachArgs{...} }
type FailoverIpAttachMapInput interface {
	pulumi.Input

	ToFailoverIpAttachMapOutput() FailoverIpAttachMapOutput
	ToFailoverIpAttachMapOutputWithContext(context.Context) FailoverIpAttachMapOutput
}

type FailoverIpAttachMap map[string]FailoverIpAttachInput

func (FailoverIpAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FailoverIpAttach)(nil)).Elem()
}

func (i FailoverIpAttachMap) ToFailoverIpAttachMapOutput() FailoverIpAttachMapOutput {
	return i.ToFailoverIpAttachMapOutputWithContext(context.Background())
}

func (i FailoverIpAttachMap) ToFailoverIpAttachMapOutputWithContext(ctx context.Context) FailoverIpAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverIpAttachMapOutput)
}

type FailoverIpAttachOutput struct{ *pulumi.OutputState }

func (FailoverIpAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverIpAttach)(nil)).Elem()
}

func (o FailoverIpAttachOutput) ToFailoverIpAttachOutput() FailoverIpAttachOutput {
	return o
}

func (o FailoverIpAttachOutput) ToFailoverIpAttachOutputWithContext(ctx context.Context) FailoverIpAttachOutput {
	return o
}

// The IP block
// * `continentCode` - The Ip continent
func (o FailoverIpAttachOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.Block }).(pulumi.StringOutput)
}

// Ip continent
func (o FailoverIpAttachOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.ContinentCode }).(pulumi.StringOutput)
}

// Ip location
func (o FailoverIpAttachOutput) GeoLoc() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.GeoLoc }).(pulumi.StringOutput)
}

// The failover ip address to attach
func (o FailoverIpAttachOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Current operation progress in percent
// * `routedTo` - Instance where ip is routed to
func (o FailoverIpAttachOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.IntOutput { return v.Progress }).(pulumi.IntOutput)
}

// The GUID of an instance to which the failover IP address is be attached
func (o FailoverIpAttachOutput) RoutedTo() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.RoutedTo }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o FailoverIpAttachOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Ip status, can be `ok` or `operationPending`
// * `subType` - IP sub type, can be `cloud` or `ovh`
func (o FailoverIpAttachOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// IP sub type
func (o FailoverIpAttachOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v *FailoverIpAttach) pulumi.StringOutput { return v.SubType }).(pulumi.StringOutput)
}

type FailoverIpAttachArrayOutput struct{ *pulumi.OutputState }

func (FailoverIpAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FailoverIpAttach)(nil)).Elem()
}

func (o FailoverIpAttachArrayOutput) ToFailoverIpAttachArrayOutput() FailoverIpAttachArrayOutput {
	return o
}

func (o FailoverIpAttachArrayOutput) ToFailoverIpAttachArrayOutputWithContext(ctx context.Context) FailoverIpAttachArrayOutput {
	return o
}

func (o FailoverIpAttachArrayOutput) Index(i pulumi.IntInput) FailoverIpAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FailoverIpAttach {
		return vs[0].([]*FailoverIpAttach)[vs[1].(int)]
	}).(FailoverIpAttachOutput)
}

type FailoverIpAttachMapOutput struct{ *pulumi.OutputState }

func (FailoverIpAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FailoverIpAttach)(nil)).Elem()
}

func (o FailoverIpAttachMapOutput) ToFailoverIpAttachMapOutput() FailoverIpAttachMapOutput {
	return o
}

func (o FailoverIpAttachMapOutput) ToFailoverIpAttachMapOutputWithContext(ctx context.Context) FailoverIpAttachMapOutput {
	return o
}

func (o FailoverIpAttachMapOutput) MapIndex(k pulumi.StringInput) FailoverIpAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FailoverIpAttach {
		return vs[0].(map[string]*FailoverIpAttach)[vs[1].(string)]
	}).(FailoverIpAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverIpAttachInput)(nil)).Elem(), &FailoverIpAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverIpAttachArrayInput)(nil)).Elem(), FailoverIpAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FailoverIpAttachMapInput)(nil)).Elem(), FailoverIpAttachMap{})
	pulumi.RegisterOutputType(FailoverIpAttachOutput{})
	pulumi.RegisterOutputType(FailoverIpAttachArrayOutput{})
	pulumi.RegisterOutputType(FailoverIpAttachMapOutput{})
}
