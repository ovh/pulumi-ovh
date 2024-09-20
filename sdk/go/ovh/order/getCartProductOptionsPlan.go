// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package order

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product options plan.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myaccount, err := Me.GetMe(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			mycart, err := Order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: myaccount.OvhSubsidiary,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Order.GetCartProductOptionsPlan(ctx, &order.GetCartProductOptionsPlanArgs{
//				CartId:          mycart.Id,
//				PriceCapacity:   "renew",
//				Product:         "cloud",
//				PlanCode:        "project",
//				OptionsPlanCode: "vrack",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCartProductOptionsPlan(ctx *pulumi.Context, args *GetCartProductOptionsPlanArgs, opts ...pulumi.InvokeOption) (*GetCartProductOptionsPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCartProductOptionsPlanResult
	err := ctx.Invoke("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCartProductOptionsPlan.
type GetCartProductOptionsPlanArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// Catalog name
	CatalogName *string `pulumi:"catalogName"`
	// options plan code.
	OptionsPlanCode string `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode string `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity string `pulumi:"priceCapacity"`
	// Product
	Product string `pulumi:"product"`
}

// A collection of values returned by getCartProductOptionsPlan.
type GetCartProductOptionsPlanResult struct {
	CartId      string  `pulumi:"cartId"`
	CatalogName *string `pulumi:"catalogName"`
	// Define if options of this family are exclusive with each other
	Exclusive bool `pulumi:"exclusive"`
	// Option family
	Family string `pulumi:"family"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Define if an option of this family is mandatory
	Mandatory       bool   `pulumi:"mandatory"`
	OptionsPlanCode string `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode      string `pulumi:"planCode"`
	PriceCapacity string `pulumi:"priceCapacity"`
	// Prices of the product offer
	Prices  []GetCartProductOptionsPlanPrice `pulumi:"prices"`
	Product string                           `pulumi:"product"`
	// Name of the product
	ProductName string `pulumi:"productName"`
	// Product type
	ProductType string `pulumi:"productType"`
	// Selected Price according to capacity
	SelectedPrices []GetCartProductOptionsPlanSelectedPrice `pulumi:"selectedPrices"`
}

func GetCartProductOptionsPlanOutput(ctx *pulumi.Context, args GetCartProductOptionsPlanOutputArgs, opts ...pulumi.InvokeOption) GetCartProductOptionsPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCartProductOptionsPlanResultOutput, error) {
			args := v.(GetCartProductOptionsPlanArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetCartProductOptionsPlanResult
			secret, err := ctx.InvokePackageRaw("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", args, &rv, "", opts...)
			if err != nil {
				return GetCartProductOptionsPlanResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetCartProductOptionsPlanResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetCartProductOptionsPlanResultOutput), nil
			}
			return output, nil
		}).(GetCartProductOptionsPlanResultOutput)
}

// A collection of arguments for invoking getCartProductOptionsPlan.
type GetCartProductOptionsPlanOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// Catalog name
	CatalogName pulumi.StringPtrInput `pulumi:"catalogName"`
	// options plan code.
	OptionsPlanCode pulumi.StringInput `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode pulumi.StringInput `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity pulumi.StringInput `pulumi:"priceCapacity"`
	// Product
	Product pulumi.StringInput `pulumi:"product"`
}

func (GetCartProductOptionsPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCartProductOptionsPlanArgs)(nil)).Elem()
}

// A collection of values returned by getCartProductOptionsPlan.
type GetCartProductOptionsPlanResultOutput struct{ *pulumi.OutputState }

func (GetCartProductOptionsPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCartProductOptionsPlanResult)(nil)).Elem()
}

func (o GetCartProductOptionsPlanResultOutput) ToGetCartProductOptionsPlanResultOutput() GetCartProductOptionsPlanResultOutput {
	return o
}

func (o GetCartProductOptionsPlanResultOutput) ToGetCartProductOptionsPlanResultOutputWithContext(ctx context.Context) GetCartProductOptionsPlanResultOutput {
	return o
}

func (o GetCartProductOptionsPlanResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetCartProductOptionsPlanResultOutput) CatalogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) *string { return v.CatalogName }).(pulumi.StringPtrOutput)
}

// Define if options of this family are exclusive with each other
func (o GetCartProductOptionsPlanResultOutput) Exclusive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) bool { return v.Exclusive }).(pulumi.BoolOutput)
}

// Option family
func (o GetCartProductOptionsPlanResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.Family }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCartProductOptionsPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Define if an option of this family is mandatory
func (o GetCartProductOptionsPlanResultOutput) Mandatory() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) bool { return v.Mandatory }).(pulumi.BoolOutput)
}

func (o GetCartProductOptionsPlanResultOutput) OptionsPlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.OptionsPlanCode }).(pulumi.StringOutput)
}

// Product offer identifier
func (o GetCartProductOptionsPlanResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

func (o GetCartProductOptionsPlanResultOutput) PriceCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.PriceCapacity }).(pulumi.StringOutput)
}

// Prices of the product offer
func (o GetCartProductOptionsPlanResultOutput) Prices() GetCartProductOptionsPlanPriceArrayOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) []GetCartProductOptionsPlanPrice { return v.Prices }).(GetCartProductOptionsPlanPriceArrayOutput)
}

func (o GetCartProductOptionsPlanResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.Product }).(pulumi.StringOutput)
}

// Name of the product
func (o GetCartProductOptionsPlanResultOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.ProductName }).(pulumi.StringOutput)
}

// Product type
func (o GetCartProductOptionsPlanResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) string { return v.ProductType }).(pulumi.StringOutput)
}

// Selected Price according to capacity
func (o GetCartProductOptionsPlanResultOutput) SelectedPrices() GetCartProductOptionsPlanSelectedPriceArrayOutput {
	return o.ApplyT(func(v GetCartProductOptionsPlanResult) []GetCartProductOptionsPlanSelectedPrice {
		return v.SelectedPrices
	}).(GetCartProductOptionsPlanSelectedPriceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCartProductOptionsPlanResultOutput{})
}
