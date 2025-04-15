// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get available images in the given public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const images = ovh.CloudProject.getImages({
 *     serviceName: "<public cloud project ID>",
 *     region: "WAW1",
 *     osType: "linux",
 * });
 * ```
 */
export function getImages(args: GetImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getImages:getImages", {
        "flavorType": args.flavorType,
        "osType": args.osType,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesArgs {
    /**
     * Get compatible images with flavor type
     */
    flavorType?: string;
    /**
     * Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     */
    osType?: string;
    /**
     * Image region
     */
    region?: string;
    /**
     * Public cloud project ID
     */
    serviceName: string;
}

/**
 * A collection of values returned by getImages.
 */
export interface GetImagesResult {
    /**
     * Get compatible images with flavor type
     */
    readonly flavorType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly images: outputs.CloudProject.GetImagesImage[];
    /**
     * Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     */
    readonly osType: string;
    /**
     * Image region
     */
    readonly region: string;
    /**
     * Public cloud project ID
     */
    readonly serviceName: string;
}
/**
 * Get available images in the given public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const images = ovh.CloudProject.getImages({
 *     serviceName: "<public cloud project ID>",
 *     region: "WAW1",
 *     osType: "linux",
 * });
 * ```
 */
export function getImagesOutput(args: GetImagesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetImagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getImages:getImages", {
        "flavorType": args.flavorType,
        "osType": args.osType,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesOutputArgs {
    /**
     * Get compatible images with flavor type
     */
    flavorType?: pulumi.Input<string>;
    /**
     * Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     */
    osType?: pulumi.Input<string>;
    /**
     * Image region
     */
    region?: pulumi.Input<string>;
    /**
     * Public cloud project ID
     */
    serviceName: pulumi.Input<string>;
}
