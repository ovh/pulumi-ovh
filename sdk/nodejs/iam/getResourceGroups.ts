// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getResourceGroups(opts?: pulumi.InvokeOptions): Promise<GetResourceGroupsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Iam/getResourceGroups:getResourceGroups", {
    }, opts);
}

/**
 * A collection of values returned by getResourceGroups.
 */
export interface GetResourceGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resourceGroups: string[];
}
export function getResourceGroupsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourceGroupsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Iam/getResourceGroups:getResourceGroups", {
    }, opts);
}
