// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const gRA1 = ovh.CloudProject.getRegion({
 *     name: "GRA1",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getRegion(args: GetRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getRegion:getRegion", {
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionArgs {
    /**
     * The name of the region associated with the public cloud
     * project.
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getRegion.
 */
export interface GetRegionResult {
    /**
     * the code of the geographic continent the region is running.
     * E.g.: EU for Europe, US for America...
     */
    readonly continentCode: string;
    /**
     * The location code of the datacenter.
     * E.g.: "GRA", meaning Gravelines, for region "GRA1"
     */
    readonly datacenterLocation: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * the name of the public cloud service
     */
    readonly name: string;
    readonly serviceName: string;
    /**
     * The list of public cloud services running within the region
     */
    readonly services: outputs.CloudProject.GetRegionService[];
}
/**
 * Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const gRA1 = ovh.CloudProject.getRegion({
 *     name: "GRA1",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getRegionOutput(args: GetRegionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getRegion:getRegion", {
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionOutputArgs {
    /**
     * The name of the region associated with the public cloud
     * project.
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
