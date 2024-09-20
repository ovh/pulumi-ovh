// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a namespace of a M3DB cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const m3dbnamespace = ovh.CloudProject.getM3dbNamespace({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const m3dbnamespaceType = m3dbnamespace.then(m3dbnamespace => m3dbnamespace.type);
 * ```
 */
export function getM3dbNamespace(args: GetM3dbNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetM3dbNamespaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getM3dbNamespace:getM3dbNamespace", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getM3dbNamespace.
 */
export interface GetM3dbNamespaceArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Name of the namespace.
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getM3dbNamespace.
 */
export interface GetM3dbNamespaceResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * Resolution for an aggregated namespace.
     */
    readonly resolution: string;
    /**
     * Controls how long we wait before expiring stale data.
     */
    readonly retentionBlockDataExpirationDuration: string;
    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk.
     */
    readonly retentionBlockSizeDuration: string;
    /**
     * Controls how far into the future writes to the namespace will be accepted.
     */
    readonly retentionBufferFutureDuration: string;
    /**
     * Controls how far into the past writes to the namespace will be accepted.
     */
    readonly retentionBufferPastDuration: string;
    /**
     * Controls the duration of time that M3DB will retain data for the namespace.
     */
    readonly retentionPeriodDuration: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * SDefines whether M3db will create snapshot files for this namespace.
     */
    readonly snapshotEnabled: boolean;
    /**
     * Type of namespace.
     */
    readonly type: string;
    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     */
    readonly writesToCommitLogEnabled: boolean;
}
/**
 * Use this data source to get information about a namespace of a M3DB cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const m3dbnamespace = ovh.CloudProject.getM3dbNamespace({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const m3dbnamespaceType = m3dbnamespace.then(m3dbnamespace => m3dbnamespace.type);
 * ```
 */
export function getM3dbNamespaceOutput(args: GetM3dbNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetM3dbNamespaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getM3dbNamespace:getM3dbNamespace", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getM3dbNamespace.
 */
export interface GetM3dbNamespaceOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the namespace.
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
