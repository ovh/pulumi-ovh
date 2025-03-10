// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getContainerRegistryIPRestrictionsManagement(args: GetContainerRegistryIPRestrictionsManagementArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerRegistryIPRestrictionsManagementResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getContainerRegistryIPRestrictionsManagement:getContainerRegistryIPRestrictionsManagement", {
        "registryId": args.registryId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistryIPRestrictionsManagement.
 */
export interface GetContainerRegistryIPRestrictionsManagementArgs {
    registryId: string;
    serviceName: string;
}

/**
 * A collection of values returned by getContainerRegistryIPRestrictionsManagement.
 */
export interface GetContainerRegistryIPRestrictionsManagementResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipRestrictions: {[key: string]: string}[];
    readonly registryId: string;
    readonly serviceName: string;
}
export function getContainerRegistryIPRestrictionsManagementOutput(args: GetContainerRegistryIPRestrictionsManagementOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerRegistryIPRestrictionsManagementResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getContainerRegistryIPRestrictionsManagement:getContainerRegistryIPRestrictionsManagement", {
        "registryId": args.registryId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistryIPRestrictionsManagement.
 */
export interface GetContainerRegistryIPRestrictionsManagementOutputArgs {
    registryId: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
