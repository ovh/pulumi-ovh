// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const storage = ovh.CloudProject.getStorage({
 *     name: "my-storage-name",
 *     regionName: "GRA",
 *     serviceName: "<public cloud project ID>",
 * });
 * ```
 */
export function getStorage(args: GetStorageArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getStorage:getStorage", {
        "limit": args.limit,
        "marker": args.marker,
        "name": args.name,
        "prefix": args.prefix,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorage.
 */
export interface GetStorageArgs {
    /**
     * Limit the number of objects returned (1000 maximum, defaults to 1000)
     */
    limit?: number;
    /**
     * Key to start with when listing objects
     */
    marker?: string;
    /**
     * Name
     */
    name: string;
    /**
     * List objects whose key begins with this prefix
     */
    prefix?: string;
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
 * A collection of values returned by getStorage.
 */
export interface GetStorageResult {
    /**
     * The date and timestamp when the resource was created
     */
    readonly createdAt: string;
    /**
     * Encryption configuration
     */
    readonly encryption: outputs.CloudProject.GetStorageEncryption;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Limit the number of objects returned (1000 maximum, defaults to 1000)
     */
    readonly limit: number;
    /**
     * Key to start with when listing objects
     */
    readonly marker: string;
    /**
     * Name
     */
    readonly name: string;
    /**
     * Container objects
     */
    readonly objects: outputs.CloudProject.GetStorageObject[];
    /**
     * Container total objects count
     */
    readonly objectsCount: number;
    /**
     * Container total objects size (bytes)
     */
    readonly objectsSize: number;
    /**
     * Container owner user ID
     */
    readonly ownerId: number;
    /**
     * List objects whose key begins with this prefix
     */
    readonly prefix: string;
    /**
     * Container region
     */
    readonly region: string;
    /**
     * Region name
     */
    readonly regionName: string;
    /**
     * Replication configuration
     */
    readonly replication: outputs.CloudProject.GetStorageReplication;
    /**
     * Service name
     */
    readonly serviceName: string;
    /**
     * Container tags
     */
    readonly tags: {[key: string]: string};
    /**
     * Versioning configuration
     */
    readonly versioning: outputs.CloudProject.GetStorageVersioning;
    /**
     * Container virtual host
     */
    readonly virtualHost: string;
}
/**
 * Get S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const storage = ovh.CloudProject.getStorage({
 *     name: "my-storage-name",
 *     regionName: "GRA",
 *     serviceName: "<public cloud project ID>",
 * });
 * ```
 */
export function getStorageOutput(args: GetStorageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStorageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getStorage:getStorage", {
        "limit": args.limit,
        "marker": args.marker,
        "name": args.name,
        "prefix": args.prefix,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorage.
 */
export interface GetStorageOutputArgs {
    /**
     * Limit the number of objects returned (1000 maximum, defaults to 1000)
     */
    limit?: pulumi.Input<number>;
    /**
     * Key to start with when listing objects
     */
    marker?: pulumi.Input<string>;
    /**
     * Name
     */
    name: pulumi.Input<string>;
    /**
     * List objects whose key begins with this prefix
     */
    prefix?: pulumi.Input<string>;
    /**
     * Region name
     */
    regionName: pulumi.Input<string>;
    /**
     * Service name
     */
    serviceName: pulumi.Input<string>;
}
