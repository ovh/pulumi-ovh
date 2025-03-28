// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const object = ovh.CloudProject.getStorageObject({
 *     key: "<object name>",
 *     name: "<bucket name>",
 *     regionName: "GRA",
 *     serviceName: "<public cloud project ID>",
 * });
 * ```
 */
export function getStorageObject(args: GetStorageObjectArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageObjectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getStorageObject:getStorageObject", {
        "key": args.key,
        "name": args.name,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageObject.
 */
export interface GetStorageObjectArgs {
    /**
     * Key
     */
    key: string;
    /**
     * Name
     */
    name: string;
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
 * A collection of values returned by getStorageObject.
 */
export interface GetStorageObjectResult {
    /**
     * ETag
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Whether this object is a delete marker
     */
    readonly isDeleteMarker: boolean;
    /**
     * Whether this is the latest version of the object
     */
    readonly isLatest: boolean;
    /**
     * Key
     */
    readonly key: string;
    /**
     * Last modification date
     */
    readonly lastModified: string;
    /**
     * Name
     */
    readonly name: string;
    /**
     * Region name
     */
    readonly regionName: string;
    /**
     * Service name
     */
    readonly serviceName: string;
    /**
     * Size (bytes)
     */
    readonly size: number;
    /**
     * Storage class
     */
    readonly storageClass: string;
    /**
     * Version ID of the object
     */
    readonly versionId: string;
}
/**
 * Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const object = ovh.CloudProject.getStorageObject({
 *     key: "<object name>",
 *     name: "<bucket name>",
 *     regionName: "GRA",
 *     serviceName: "<public cloud project ID>",
 * });
 * ```
 */
export function getStorageObjectOutput(args: GetStorageObjectOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStorageObjectResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getStorageObject:getStorageObject", {
        "key": args.key,
        "name": args.name,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageObject.
 */
export interface GetStorageObjectOutputArgs {
    /**
     * Key
     */
    key: pulumi.Input<string>;
    /**
     * Name
     */
    name: pulumi.Input<string>;
    /**
     * Region name
     */
    regionName: pulumi.Input<string>;
    /**
     * Service name
     */
    serviceName: pulumi.Input<string>;
}
