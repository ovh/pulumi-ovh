// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information of order cart product options plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * const mycart = myaccount.then(myaccount => ovh.Order.getCart({
 *     ovhSubsidiary: myaccount.ovhSubsidiary,
 * }));
 * const plan = mycart.then(mycart => ovh.Order.getCartProductOptionsPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "cloud",
 *     planCode: "project",
 *     optionsPlanCode: "vrack",
 * }));
 * ```
 */
export function getCartProductOptionsPlan(args: GetCartProductOptionsPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetCartProductOptionsPlanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "optionsPlanCode": args.optionsPlanCode,
        "planCode": args.planCode,
        "priceCapacity": args.priceCapacity,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getCartProductOptionsPlan.
 */
export interface GetCartProductOptionsPlanArgs {
    /**
     * Cart identifier
     */
    cartId: string;
    /**
     * Catalog name
     */
    catalogName?: string;
    /**
     * options plan code.
     */
    optionsPlanCode: string;
    /**
     * Product offer identifier
     */
    planCode: string;
    /**
     * Capacity of the pricing (type of pricing)
     */
    priceCapacity: string;
    /**
     * Product
     */
    product: string;
}

/**
 * A collection of values returned by getCartProductOptionsPlan.
 */
export interface GetCartProductOptionsPlanResult {
    readonly cartId: string;
    readonly catalogName?: string;
    /**
     * Define if options of this family are exclusive with each other
     */
    readonly exclusive: boolean;
    /**
     * Option family
     */
    readonly family: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Define if an option of this family is mandatory
     */
    readonly mandatory: boolean;
    readonly optionsPlanCode: string;
    /**
     * Product offer identifier
     */
    readonly planCode: string;
    readonly priceCapacity: string;
    /**
     * Prices of the product offer
     */
    readonly prices: outputs.Order.GetCartProductOptionsPlanPrice[];
    readonly product: string;
    /**
     * Name of the product
     */
    readonly productName: string;
    /**
     * Product type
     */
    readonly productType: string;
    /**
     * Selected Price according to capacity
     */
    readonly selectedPrices: outputs.Order.GetCartProductOptionsPlanSelectedPrice[];
}
/**
 * Use this data source to retrieve information of order cart product options plan.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * const mycart = myaccount.then(myaccount => ovh.Order.getCart({
 *     ovhSubsidiary: myaccount.ovhSubsidiary,
 * }));
 * const plan = mycart.then(mycart => ovh.Order.getCartProductOptionsPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "cloud",
 *     planCode: "project",
 *     optionsPlanCode: "vrack",
 * }));
 * ```
 */
export function getCartProductOptionsPlanOutput(args: GetCartProductOptionsPlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCartProductOptionsPlanResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", {
        "cartId": args.cartId,
        "catalogName": args.catalogName,
        "optionsPlanCode": args.optionsPlanCode,
        "planCode": args.planCode,
        "priceCapacity": args.priceCapacity,
        "product": args.product,
    }, opts);
}

/**
 * A collection of arguments for invoking getCartProductOptionsPlan.
 */
export interface GetCartProductOptionsPlanOutputArgs {
    /**
     * Cart identifier
     */
    cartId: pulumi.Input<string>;
    /**
     * Catalog name
     */
    catalogName?: pulumi.Input<string>;
    /**
     * options plan code.
     */
    optionsPlanCode: pulumi.Input<string>;
    /**
     * Product offer identifier
     */
    planCode: pulumi.Input<string>;
    /**
     * Capacity of the pricing (type of pricing)
     */
    priceCapacity: pulumi.Input<string>;
    /**
     * Product
     */
    product: pulumi.Input<string>;
}
