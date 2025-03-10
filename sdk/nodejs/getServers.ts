// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getServers(opts?: pulumi.InvokeOptions): Promise<GetServersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:index/getServers:getServers", {
    }, opts);
}

/**
 * A collection of values returned by getServers.
 */
export interface GetServersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly results: string[];
}
export function getServersOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:index/getServers:getServers", {
    }, opts);
}
