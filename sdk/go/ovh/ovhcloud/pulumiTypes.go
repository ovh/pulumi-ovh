// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovhcloud

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ConnectIam struct {
	// Resource display name
	DisplayName string `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id string `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags map[string]string `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn string `pulumi:"urn"`
}

// ConnectIamInput is an input type that accepts ConnectIamArgs and ConnectIamOutput values.
// You can construct a concrete instance of `ConnectIamInput` via:
//
//	ConnectIamArgs{...}
type ConnectIamInput interface {
	pulumi.Input

	ToConnectIamOutput() ConnectIamOutput
	ToConnectIamOutputWithContext(context.Context) ConnectIamOutput
}

type ConnectIamArgs struct {
	// Resource display name
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id pulumi.StringInput `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringInput `pulumi:"urn"`
}

func (ConnectIamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectIam)(nil)).Elem()
}

func (i ConnectIamArgs) ToConnectIamOutput() ConnectIamOutput {
	return i.ToConnectIamOutputWithContext(context.Background())
}

func (i ConnectIamArgs) ToConnectIamOutputWithContext(ctx context.Context) ConnectIamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectIamOutput)
}

type ConnectIamOutput struct{ *pulumi.OutputState }

func (ConnectIamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectIam)(nil)).Elem()
}

func (o ConnectIamOutput) ToConnectIamOutput() ConnectIamOutput {
	return o
}

func (o ConnectIamOutput) ToConnectIamOutputWithContext(ctx context.Context) ConnectIamOutput {
	return o
}

// Resource display name
func (o ConnectIamOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectIam) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Unique identifier of the resource in the IAM
func (o ConnectIamOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectIam) string { return v.Id }).(pulumi.StringOutput)
}

// Resource tags. Tags that were internally computed are prefixed with `ovh:`
func (o ConnectIamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectIam) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// URN of the private database, used when writing IAM policies
func (o ConnectIamOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectIam) string { return v.Urn }).(pulumi.StringOutput)
}

type ConnectsOcc struct {
	// Service bandwidth
	Bandwidth string `pulumi:"bandwidth"`
	// Service description
	Description string `pulumi:"description"`
	// IAM resource information
	Iam ConnectsOccIam `pulumi:"iam"`
	// List of interfaces linked to a service
	InterfaceLists []float64 `pulumi:"interfaceLists"`
	// Pop reference where the service is delivered
	Pop string `pulumi:"pop"`
	// Port quantity
	PortQuantity string `pulumi:"portQuantity"`
	// Product name of the service
	Product string `pulumi:"product"`
	// Service provider
	ProviderName string `pulumi:"providerName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
	// Service status
	Status string `pulumi:"status"`
	// uuid of the Ovhcloud Connect service
	Uuid string `pulumi:"uuid"`
	// vrack linked to the service
	Vrack string `pulumi:"vrack"`
}

// ConnectsOccInput is an input type that accepts ConnectsOccArgs and ConnectsOccOutput values.
// You can construct a concrete instance of `ConnectsOccInput` via:
//
//	ConnectsOccArgs{...}
type ConnectsOccInput interface {
	pulumi.Input

	ToConnectsOccOutput() ConnectsOccOutput
	ToConnectsOccOutputWithContext(context.Context) ConnectsOccOutput
}

type ConnectsOccArgs struct {
	// Service bandwidth
	Bandwidth pulumi.StringInput `pulumi:"bandwidth"`
	// Service description
	Description pulumi.StringInput `pulumi:"description"`
	// IAM resource information
	Iam ConnectsOccIamInput `pulumi:"iam"`
	// List of interfaces linked to a service
	InterfaceLists pulumi.Float64ArrayInput `pulumi:"interfaceLists"`
	// Pop reference where the service is delivered
	Pop pulumi.StringInput `pulumi:"pop"`
	// Port quantity
	PortQuantity pulumi.StringInput `pulumi:"portQuantity"`
	// Product name of the service
	Product pulumi.StringInput `pulumi:"product"`
	// Service provider
	ProviderName pulumi.StringInput `pulumi:"providerName"`
	// Service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Service status
	Status pulumi.StringInput `pulumi:"status"`
	// uuid of the Ovhcloud Connect service
	Uuid pulumi.StringInput `pulumi:"uuid"`
	// vrack linked to the service
	Vrack pulumi.StringInput `pulumi:"vrack"`
}

func (ConnectsOccArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectsOcc)(nil)).Elem()
}

func (i ConnectsOccArgs) ToConnectsOccOutput() ConnectsOccOutput {
	return i.ToConnectsOccOutputWithContext(context.Background())
}

func (i ConnectsOccArgs) ToConnectsOccOutputWithContext(ctx context.Context) ConnectsOccOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectsOccOutput)
}

// ConnectsOccArrayInput is an input type that accepts ConnectsOccArray and ConnectsOccArrayOutput values.
// You can construct a concrete instance of `ConnectsOccArrayInput` via:
//
//	ConnectsOccArray{ ConnectsOccArgs{...} }
type ConnectsOccArrayInput interface {
	pulumi.Input

	ToConnectsOccArrayOutput() ConnectsOccArrayOutput
	ToConnectsOccArrayOutputWithContext(context.Context) ConnectsOccArrayOutput
}

type ConnectsOccArray []ConnectsOccInput

func (ConnectsOccArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectsOcc)(nil)).Elem()
}

func (i ConnectsOccArray) ToConnectsOccArrayOutput() ConnectsOccArrayOutput {
	return i.ToConnectsOccArrayOutputWithContext(context.Background())
}

func (i ConnectsOccArray) ToConnectsOccArrayOutputWithContext(ctx context.Context) ConnectsOccArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectsOccArrayOutput)
}

type ConnectsOccOutput struct{ *pulumi.OutputState }

func (ConnectsOccOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectsOcc)(nil)).Elem()
}

func (o ConnectsOccOutput) ToConnectsOccOutput() ConnectsOccOutput {
	return o
}

func (o ConnectsOccOutput) ToConnectsOccOutputWithContext(ctx context.Context) ConnectsOccOutput {
	return o
}

// Service bandwidth
func (o ConnectsOccOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Bandwidth }).(pulumi.StringOutput)
}

// Service description
func (o ConnectsOccOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Description }).(pulumi.StringOutput)
}

// IAM resource information
func (o ConnectsOccOutput) Iam() ConnectsOccIamOutput {
	return o.ApplyT(func(v ConnectsOcc) ConnectsOccIam { return v.Iam }).(ConnectsOccIamOutput)
}

// List of interfaces linked to a service
func (o ConnectsOccOutput) InterfaceLists() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v ConnectsOcc) []float64 { return v.InterfaceLists }).(pulumi.Float64ArrayOutput)
}

// Pop reference where the service is delivered
func (o ConnectsOccOutput) Pop() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Pop }).(pulumi.StringOutput)
}

// Port quantity
func (o ConnectsOccOutput) PortQuantity() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.PortQuantity }).(pulumi.StringOutput)
}

// Product name of the service
func (o ConnectsOccOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Product }).(pulumi.StringOutput)
}

// Service provider
func (o ConnectsOccOutput) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.ProviderName }).(pulumi.StringOutput)
}

// Service name
func (o ConnectsOccOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Service status
func (o ConnectsOccOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Status }).(pulumi.StringOutput)
}

// uuid of the Ovhcloud Connect service
func (o ConnectsOccOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Uuid }).(pulumi.StringOutput)
}

// vrack linked to the service
func (o ConnectsOccOutput) Vrack() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOcc) string { return v.Vrack }).(pulumi.StringOutput)
}

type ConnectsOccArrayOutput struct{ *pulumi.OutputState }

func (ConnectsOccArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectsOcc)(nil)).Elem()
}

func (o ConnectsOccArrayOutput) ToConnectsOccArrayOutput() ConnectsOccArrayOutput {
	return o
}

func (o ConnectsOccArrayOutput) ToConnectsOccArrayOutputWithContext(ctx context.Context) ConnectsOccArrayOutput {
	return o
}

func (o ConnectsOccArrayOutput) Index(i pulumi.IntInput) ConnectsOccOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectsOcc {
		return vs[0].([]ConnectsOcc)[vs[1].(int)]
	}).(ConnectsOccOutput)
}

type ConnectsOccIam struct {
	// Resource display name
	DisplayName string `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id string `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags map[string]string `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn string `pulumi:"urn"`
}

// ConnectsOccIamInput is an input type that accepts ConnectsOccIamArgs and ConnectsOccIamOutput values.
// You can construct a concrete instance of `ConnectsOccIamInput` via:
//
//	ConnectsOccIamArgs{...}
type ConnectsOccIamInput interface {
	pulumi.Input

	ToConnectsOccIamOutput() ConnectsOccIamOutput
	ToConnectsOccIamOutputWithContext(context.Context) ConnectsOccIamOutput
}

type ConnectsOccIamArgs struct {
	// Resource display name
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id pulumi.StringInput `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringInput `pulumi:"urn"`
}

func (ConnectsOccIamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectsOccIam)(nil)).Elem()
}

func (i ConnectsOccIamArgs) ToConnectsOccIamOutput() ConnectsOccIamOutput {
	return i.ToConnectsOccIamOutputWithContext(context.Background())
}

func (i ConnectsOccIamArgs) ToConnectsOccIamOutputWithContext(ctx context.Context) ConnectsOccIamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectsOccIamOutput)
}

type ConnectsOccIamOutput struct{ *pulumi.OutputState }

func (ConnectsOccIamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectsOccIam)(nil)).Elem()
}

func (o ConnectsOccIamOutput) ToConnectsOccIamOutput() ConnectsOccIamOutput {
	return o
}

func (o ConnectsOccIamOutput) ToConnectsOccIamOutputWithContext(ctx context.Context) ConnectsOccIamOutput {
	return o
}

// Resource display name
func (o ConnectsOccIamOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOccIam) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Unique identifier of the resource in the IAM
func (o ConnectsOccIamOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOccIam) string { return v.Id }).(pulumi.StringOutput)
}

// Resource tags. Tags that were internally computed are prefixed with `ovh:`
func (o ConnectsOccIamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConnectsOccIam) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// URN of the private database, used when writing IAM policies
func (o ConnectsOccIamOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectsOccIam) string { return v.Urn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectIamInput)(nil)).Elem(), ConnectIamArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectsOccInput)(nil)).Elem(), ConnectsOccArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectsOccArrayInput)(nil)).Elem(), ConnectsOccArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectsOccIamInput)(nil)).Elem(), ConnectsOccIamArgs{})
	pulumi.RegisterOutputType(ConnectIamOutput{})
	pulumi.RegisterOutputType(ConnectsOccOutput{})
	pulumi.RegisterOutputType(ConnectsOccArrayOutput{})
	pulumi.RegisterOutputType(ConnectsOccIamOutput{})
}
