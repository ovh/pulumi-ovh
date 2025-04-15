// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information about an image in the given public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const image = ovh.CloudProject.getImage({
 *     serviceName: "<public cloud project ID>",
 *     imageId: "<image ID>",
 * });
 * ```
 */
export function getImage(args: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getImage:getImage", {
        "imageId": args.imageId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * Image ID
     */
    imageId: string;
    /**
     * Public cloud project ID
     */
    serviceName: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * Image creation date
     */
    readonly creationDate: string;
    /**
     * Image usable only for this type of flavor if not null
     */
    readonly flavorType: string;
    /**
     * Image ID
     */
    readonly id: string;
    /**
     * Image ID
     */
    readonly imageId: string;
    /**
     * Minimum disks required to use image
     */
    readonly minDisk: number;
    /**
     * Minimum RAM required to use image
     */
    readonly minRam: number;
    /**
     * Image name
     */
    readonly name: string;
    /**
     * Order plan code
     */
    readonly planCode: string;
    /**
     * Image region
     */
    readonly region: string;
    /**
     * Public cloud project ID
     */
    readonly serviceName: string;
    /**
     * Image size (in GiB)
     */
    readonly size: number;
    /**
     * Image status
     */
    readonly status: string;
    /**
     * Tags about the image
     */
    readonly tags: string[];
    /**
     * Image type
     */
    readonly type: string;
    /**
     * User to connect with
     */
    readonly user: string;
    /**
     * Image visibility
     */
    readonly visibility: string;
}
/**
 * Get information about an image in the given public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const image = ovh.CloudProject.getImage({
 *     serviceName: "<public cloud project ID>",
 *     imageId: "<image ID>",
 * });
 * ```
 */
export function getImageOutput(args: GetImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetImageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getImage:getImage", {
        "imageId": args.imageId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageOutputArgs {
    /**
     * Image ID
     */
    imageId: pulumi.Input<string>;
    /**
     * Public cloud project ID
     */
    serviceName: pulumi.Input<string>;
}
