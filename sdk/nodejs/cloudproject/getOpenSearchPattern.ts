// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a pattern of a opensearch cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const pattern = ovh.CloudProject.getOpenSearchPattern({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const patternPattern = pattern.then(pattern => pattern.pattern);
 * ```
 */
export function getOpenSearchPattern(args: GetOpenSearchPatternArgs, opts?: pulumi.InvokeOptions): Promise<GetOpenSearchPatternResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getOpenSearchPattern:getOpenSearchPattern", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpenSearchPattern.
 */
export interface GetOpenSearchPatternArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Pattern ID.
     */
    id: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getOpenSearchPattern.
 */
export interface GetOpenSearchPatternResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * See Argument Reference above.
     */
    readonly id: string;
    /**
     * Maximum number of index for this pattern.
     */
    readonly maxIndexCount: number;
    /**
     * Pattern format.
     */
    readonly pattern: string;
    /**
     * Current status of the pattern.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get information about a pattern of a opensearch cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const pattern = ovh.CloudProject.getOpenSearchPattern({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const patternPattern = pattern.then(pattern => pattern.pattern);
 * ```
 */
export function getOpenSearchPatternOutput(args: GetOpenSearchPatternOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOpenSearchPatternResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getOpenSearchPattern:getOpenSearchPattern", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpenSearchPattern.
 */
export interface GetOpenSearchPatternOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Pattern ID.
     */
    id: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
