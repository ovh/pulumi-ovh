// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * List your S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const storage = ovh.CloudProject.getStorage({
 *     serviceName: "<public cloud project ID>",
 *     regionName: "GRA",
 * });
 * ```
 */
export function getStorages(args: GetStoragesArgs, opts?: pulumi.InvokeOptions): Promise<GetStoragesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getStorages:getStorages", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorages.
 */
export interface GetStoragesArgs {
    /**
     * Region name
     */
    regionName: string;
    /**
     * Service name
     */
    serviceName: string;
}

/**
 * A collection of values returned by getStorages.
 */
export interface GetStoragesResult {
    readonly containers: outputs.CloudProject.GetStoragesContainer[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Region name
     */
    readonly regionName: string;
    /**
     * Service name
     */
    readonly serviceName: string;
}
/**
 * List your S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const storage = ovh.CloudProject.getStorage({
 *     serviceName: "<public cloud project ID>",
 *     regionName: "GRA",
 * });
 * ```
 */
export function getStoragesOutput(args: GetStoragesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStoragesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getStorages:getStorages", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorages.
 */
export interface GetStoragesOutputArgs {
    /**
     * Region name
     */
    regionName: pulumi.Input<string>;
    /**
     * Service name
     */
    serviceName: pulumi.Input<string>;
}
