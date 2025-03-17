// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get the details of your Ovhcloud Connect products.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const occs = ovh.OVHcloud.Connect({});
 * ```
 */
export function connects(opts?: pulumi.InvokeOptions): Promise<ConnectsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:OVHcloud/connects:Connects", {
    }, opts);
}

/**
 * A collection of values returned by Connects.
 */
export interface ConnectsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly occs: outputs.OVHcloud.ConnectsOcc[];
}
/**
 * Get the details of your Ovhcloud Connect products.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const occs = ovh.OVHcloud.Connect({});
 * ```
 */
export function connectsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<ConnectsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:OVHcloud/connects:Connects", {
    }, opts);
}
