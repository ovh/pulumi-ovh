// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach an IP block to a VRack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/ip"
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/me"
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/order"
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/vrack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myAccount, err := me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			myCart, err := order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: myAccount.OvhSubsidiary,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vrack, err := order.GetCartProductPlan(ctx, &order.GetCartProductPlanArgs{
//				CartId:        myCart.Id,
//				PriceCapacity: "renew",
//				Product:       "vrack",
//				PlanCode:      "vrack",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vrackVrack, err := vrack.NewVrack(ctx, "vrack", &vrack.VrackArgs{
//				Description:   pulumi.String(myCart.Description),
//				Name:          pulumi.String(myCart.Description),
//				OvhSubsidiary: pulumi.String(myCart.OvhSubsidiary),
//				Plan: &vrack.VrackPlanArgs{
//					Duration:    pulumi.String(vrack.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(vrack.PlanCode),
//					PricingMode: pulumi.String(vrack.SelectedPrices[0].PricingMode),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ipblock, err := order.GetCartProductPlan(ctx, &order.GetCartProductPlanArgs{
//				CartId:        myCart.Id,
//				PriceCapacity: "renew",
//				Product:       "ip",
//				PlanCode:      "ip-v4-s30-ripe",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ipblockIpService, err := ip.NewIpService(ctx, "ipblock", &ip.IpServiceArgs{
//				OvhSubsidiary: pulumi.String(myCart.OvhSubsidiary),
//				Description:   pulumi.String(myCart.Description),
//				Plan: &ip.IpServicePlanArgs{
//					Duration:    pulumi.String(ipblock.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(ipblock.PlanCode),
//					PricingMode: pulumi.String(ipblock.SelectedPrices[0].PricingMode),
//					Configurations: ip.IpServicePlanConfigurationArray{
//						&ip.IpServicePlanConfigurationArgs{
//							Label: pulumi.String("country"),
//							Value: pulumi.String("FR"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vrack.NewIpAddress(ctx, "vrack_block", &vrack.IpAddressArgs{
//				ServiceName: vrackVrack.ServiceName,
//				Block:       ipblockIpService.Ip,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpAddress struct {
	pulumi.CustomResourceState

	// Your IP block.
	Block pulumi.StringOutput `pulumi:"block"`
	// Your gateway
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Your IP block
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The region (e.g: eu-west-gra) where want to route your block to.
	Region pulumi.StringOutput `pulumi:"region"`
	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIpAddress registers a new resource with the given unique name, arguments, and options.
func NewIpAddress(ctx *pulumi.Context,
	name string, args *IpAddressArgs, opts ...pulumi.ResourceOption) (*IpAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Block == nil {
		return nil, errors.New("invalid value for required argument 'Block'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpAddress
	err := ctx.RegisterResource("ovh:Vrack/ipAddress:IpAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAddress gets an existing IpAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAddressState, opts ...pulumi.ResourceOption) (*IpAddress, error) {
	var resource IpAddress
	err := ctx.ReadResource("ovh:Vrack/ipAddress:IpAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAddress resources.
type ipAddressState struct {
	// Your IP block.
	Block *string `pulumi:"block"`
	// Your gateway
	Gateway *string `pulumi:"gateway"`
	// Your IP block
	Ip *string `pulumi:"ip"`
	// The region (e.g: eu-west-gra) where want to route your block to.
	Region *string `pulumi:"region"`
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone *string `pulumi:"zone"`
}

type IpAddressState struct {
	// Your IP block.
	Block pulumi.StringPtrInput
	// Your gateway
	Gateway pulumi.StringPtrInput
	// Your IP block
	Ip pulumi.StringPtrInput
	// The region (e.g: eu-west-gra) where want to route your block to.
	Region pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
	// Where you want your block announced on the network
	Zone pulumi.StringPtrInput
}

func (IpAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAddressState)(nil)).Elem()
}

type ipAddressArgs struct {
	// Your IP block.
	Block string `pulumi:"block"`
	// The region (e.g: eu-west-gra) where want to route your block to.
	Region *string `pulumi:"region"`
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a IpAddress resource.
type IpAddressArgs struct {
	// Your IP block.
	Block pulumi.StringInput
	// The region (e.g: eu-west-gra) where want to route your block to.
	Region pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringInput
}

func (IpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAddressArgs)(nil)).Elem()
}

type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput
}

func (*IpAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddress)(nil)).Elem()
}

func (i *IpAddress) ToIpAddressOutput() IpAddressOutput {
	return i.ToIpAddressOutputWithContext(context.Background())
}

func (i *IpAddress) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput)
}

// IpAddressArrayInput is an input type that accepts IpAddressArray and IpAddressArrayOutput values.
// You can construct a concrete instance of `IpAddressArrayInput` via:
//
//	IpAddressArray{ IpAddressArgs{...} }
type IpAddressArrayInput interface {
	pulumi.Input

	ToIpAddressArrayOutput() IpAddressArrayOutput
	ToIpAddressArrayOutputWithContext(context.Context) IpAddressArrayOutput
}

type IpAddressArray []IpAddressInput

func (IpAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAddress)(nil)).Elem()
}

func (i IpAddressArray) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return i.ToIpAddressArrayOutputWithContext(context.Background())
}

func (i IpAddressArray) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressArrayOutput)
}

// IpAddressMapInput is an input type that accepts IpAddressMap and IpAddressMapOutput values.
// You can construct a concrete instance of `IpAddressMapInput` via:
//
//	IpAddressMap{ "key": IpAddressArgs{...} }
type IpAddressMapInput interface {
	pulumi.Input

	ToIpAddressMapOutput() IpAddressMapOutput
	ToIpAddressMapOutputWithContext(context.Context) IpAddressMapOutput
}

type IpAddressMap map[string]IpAddressInput

func (IpAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAddress)(nil)).Elem()
}

func (i IpAddressMap) ToIpAddressMapOutput() IpAddressMapOutput {
	return i.ToIpAddressMapOutputWithContext(context.Background())
}

func (i IpAddressMap) ToIpAddressMapOutputWithContext(ctx context.Context) IpAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressMapOutput)
}

type IpAddressOutput struct{ *pulumi.OutputState }

func (IpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddress)(nil)).Elem()
}

func (o IpAddressOutput) ToIpAddressOutput() IpAddressOutput {
	return o
}

func (o IpAddressOutput) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return o
}

// Your IP block.
func (o IpAddressOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.Block }).(pulumi.StringOutput)
}

// Your gateway
func (o IpAddressOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Your IP block
func (o IpAddressOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The region (e.g: eu-west-gra) where want to route your block to.
func (o IpAddressOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The internal name of your vrack
func (o IpAddressOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Where you want your block announced on the network
func (o IpAddressOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAddress) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IpAddressArrayOutput struct{ *pulumi.OutputState }

func (IpAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAddress)(nil)).Elem()
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) Index(i pulumi.IntInput) IpAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpAddress {
		return vs[0].([]*IpAddress)[vs[1].(int)]
	}).(IpAddressOutput)
}

type IpAddressMapOutput struct{ *pulumi.OutputState }

func (IpAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAddress)(nil)).Elem()
}

func (o IpAddressMapOutput) ToIpAddressMapOutput() IpAddressMapOutput {
	return o
}

func (o IpAddressMapOutput) ToIpAddressMapOutputWithContext(ctx context.Context) IpAddressMapOutput {
	return o
}

func (o IpAddressMapOutput) MapIndex(k pulumi.StringInput) IpAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpAddress {
		return vs[0].(map[string]*IpAddress)[vs[1].(string)]
	}).(IpAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressInput)(nil)).Elem(), &IpAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressArrayInput)(nil)).Elem(), IpAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressMapInput)(nil)).Elem(), IpAddressMap{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressArrayOutput{})
	pulumi.RegisterOutputType(IpAddressMapOutput{})
}
