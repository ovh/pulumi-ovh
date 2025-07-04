// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of pattern of a opensearch cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const patterns = ovh.CloudProject.getOpenSearchPatterns({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 * });
 * export const patternIds = patterns.then(patterns => patterns.patternIds);
 * ```
 */
export function getOpenSearchPatterns(args: GetOpenSearchPatternsArgs, opts?: pulumi.InvokeOptions): Promise<GetOpenSearchPatternsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getOpenSearchPatterns:getOpenSearchPatterns", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpenSearchPatterns.
 */
export interface GetOpenSearchPatternsArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getOpenSearchPatterns.
 */
export interface GetOpenSearchPatternsResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of patterns ids of the opensearch cluster associated with the project.
     */
    readonly patternIds: string[];
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get the list of pattern of a opensearch cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const patterns = ovh.CloudProject.getOpenSearchPatterns({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 * });
 * export const patternIds = patterns.then(patterns => patterns.patternIds);
 * ```
 */
export function getOpenSearchPatternsOutput(args: GetOpenSearchPatternsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOpenSearchPatternsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getOpenSearchPatterns:getOpenSearchPatterns", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getOpenSearchPatterns.
 */
export interface GetOpenSearchPatternsOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
