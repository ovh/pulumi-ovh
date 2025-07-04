// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of users of a database cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const users = ovh.CloudProjectDatabase.getUsers({
 *     serviceName: "XXXX",
 *     engine: "YYYY",
 *     clusterId: "ZZZ",
 * });
 * export const userIds = users.then(users => users.userIds);
 * ```
 */
export function getUsers(args: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getUsers:getUsers", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * The engine of the database cluster you want to list users. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
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
     * The list of users ids of the database cluster associated with the project.
     */
    readonly userIds: string[];
}
/**
 * Use this data source to get the list of users of a database cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const users = ovh.CloudProjectDatabase.getUsers({
 *     serviceName: "XXXX",
 *     engine: "YYYY",
 *     clusterId: "ZZZ",
 * });
 * export const userIds = users.then(users => users.userIds);
 * ```
 */
export function getUsersOutput(args: GetUsersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUsersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getUsers:getUsers", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * The engine of the database cluster you want to list users. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
