// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get the list of installation templates available for dedicated servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const templates = ovh.getInstallationTemplates({});
 * ```
 */
export function getInstallationTemplates(opts?: pulumi.InvokeOptions): Promise<GetInstallationTemplatesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:index/getInstallationTemplates:getInstallationTemplates", {
    }, opts);
}

/**
 * A collection of values returned by getInstallationTemplates.
 */
export interface GetInstallationTemplatesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of installation templates IDs available for dedicated servers.
     */
    readonly results: string[];
}
/**
 * Use this data source to get the list of installation templates available for dedicated servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const templates = ovh.getInstallationTemplates({});
 * ```
 */
export function getInstallationTemplatesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstallationTemplatesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:index/getInstallationTemplates:getInstallationTemplates", {
    }, opts);
}
