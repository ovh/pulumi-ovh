// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the regions of a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const regions = ovh.CloudProject.getRegions({
 *     hasServicesUps: ["network"],
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getRegions(args: GetRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getRegions:getRegions", {
        "hasServicesUps": args.hasServicesUps,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsArgs {
    /**
     * List of services which has to be UP in regions.
     * Example: "image", "instance", "network", "storage", "volume", "workflow", ...
     * If left blank, returns all regions associated with the service_name.
     */
    hasServicesUps?: string[];
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getRegions.
 */
export interface GetRegionsResult {
    readonly hasServicesUps?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of regions associated with the project, filtered by services UP.
     */
    readonly names: string[];
    readonly serviceName: string;
}
/**
 * Use this data source to get the regions of a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const regions = ovh.CloudProject.getRegions({
 *     hasServicesUps: ["network"],
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getRegionsOutput(args: GetRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getRegions:getRegions", {
        "hasServicesUps": args.hasServicesUps,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsOutputArgs {
    /**
     * List of services which has to be UP in regions.
     * Example: "image", "instance", "network", "storage", "volume", "workflow", ...
     * If left blank, returns all regions associated with the service_name.
     */
    hasServicesUps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
