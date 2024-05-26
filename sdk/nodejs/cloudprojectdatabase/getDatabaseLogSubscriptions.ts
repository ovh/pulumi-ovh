// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of log subscrition for a cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const subscriptions = ovh.CloudProjectDatabase.getDatabaseLogSubscriptions({
 *     serviceName: "XXX",
 *     engine: "YYY",
 *     clusterId: "ZZZ",
 * });
 * export const subscriptionIds = subscriptions.then(subscriptions => subscriptions.subscriptionIds);
 * ```
 */
export function getDatabaseLogSubscriptions(args: GetDatabaseLogSubscriptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseLogSubscriptionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getDatabaseLogSubscriptions:getDatabaseLogSubscriptions", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseLogSubscriptions.
 */
export interface GetDatabaseLogSubscriptionsArgs {
    /**
     * Cluster ID.
     */
    clusterId: string;
    /**
     * The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getDatabaseLogSubscriptions.
 */
export interface GetDatabaseLogSubscriptionsResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * See Argument Reference above.
     */
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * The list of log subscription ids of the cluster associated with the project.
     */
    readonly subscriptionIds: string[];
}
/**
 * Use this data source to get the list of log subscrition for a cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const subscriptions = ovh.CloudProjectDatabase.getDatabaseLogSubscriptions({
 *     serviceName: "XXX",
 *     engine: "YYY",
 *     clusterId: "ZZZ",
 * });
 * export const subscriptionIds = subscriptions.then(subscriptions => subscriptions.subscriptionIds);
 * ```
 */
export function getDatabaseLogSubscriptionsOutput(args: GetDatabaseLogSubscriptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseLogSubscriptionsResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseLogSubscriptions(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseLogSubscriptions.
 */
export interface GetDatabaseLogSubscriptionsOutputArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
