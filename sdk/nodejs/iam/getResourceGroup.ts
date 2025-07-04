// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source get details about a resource group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myResourceGroup = ovh.Iam.getResourceGroup({
 *     id: "my_resource_group_id",
 * });
 * ```
 */
export function getResourceGroup(args: GetResourceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Iam/getResourceGroup:getResourceGroup", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceGroup.
 */
export interface GetResourceGroupArgs {
    /**
     * Id of the resource group
     */
    id: string;
}

/**
 * A collection of values returned by getResourceGroup.
 */
export interface GetResourceGroupResult {
    /**
     * URN of the resource group, used when writing policies
     */
    readonly GroupURN: string;
    /**
     * Date of the creation of the resource group
     */
    readonly createdAt: string;
    readonly id: string;
    /**
     * Name of the resource group
     */
    readonly name: string;
    /**
     * Name of the account owning the resource group
     */
    readonly owner: string;
    /**
     * Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
     */
    readonly readOnly: boolean;
    /**
     * Set of the URNs of the resources contained in the resource group
     */
    readonly resources: string[];
    /**
     * Date of the last modification of the resource group
     */
    readonly updatedAt: string;
}
/**
 * Use this data source get details about a resource group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myResourceGroup = ovh.Iam.getResourceGroup({
 *     id: "my_resource_group_id",
 * });
 * ```
 */
export function getResourceGroupOutput(args: GetResourceGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Iam/getResourceGroup:getResourceGroup", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceGroup.
 */
export interface GetResourceGroupOutputArgs {
    /**
     * Id of the resource group
     */
    id: pulumi.Input<string>;
}
