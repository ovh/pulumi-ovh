// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a dedicated HA-NAS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myNasHa = ovh.Dedicated.getNasHA({
 *     serviceName: "zpool-12345",
 * });
 * ```
 */
export function getNasHA(args: GetNasHAArgs, opts?: pulumi.InvokeOptions): Promise<GetNasHAResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dedicated/getNasHA:getNasHA", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNasHA.
 */
export interface GetNasHAArgs {
    /**
     * The serviceName of your dedicated HA-NAS.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getNasHA.
 */
export interface GetNasHAResult {
    /**
     * the URN of the HA-NAS instance
     */
    readonly NasHAURN: string;
    /**
     * True, if partition creation is allowed on this HA-NAS
     */
    readonly canCreatePartition: boolean;
    /**
     * The name you give to the HA-NAS
     */
    readonly customName: string;
    /**
     * area of HA-NAS
     */
    readonly datacenter: string;
    /**
     * the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
     */
    readonly diskType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Access IP of HA-NAS
     */
    readonly ip: string;
    /**
     * Send an email to customer if any issue is detected
     */
    readonly monitored: boolean;
    /**
     * The storage service name
     */
    readonly serviceName: string;
    /**
     * percentage of HA-NAS space used in %
     */
    readonly zpoolCapacity: number;
    /**
     * the size of the HA-NAS in GB
     */
    readonly zpoolSize: number;
}
/**
 * Use this data source to retrieve information about a dedicated HA-NAS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myNasHa = ovh.Dedicated.getNasHA({
 *     serviceName: "zpool-12345",
 * });
 * ```
 */
export function getNasHAOutput(args: GetNasHAOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNasHAResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Dedicated/getNasHA:getNasHA", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNasHA.
 */
export interface GetNasHAOutputArgs {
    /**
     * The serviceName of your dedicated HA-NAS.
     */
    serviceName: pulumi.Input<string>;
}
