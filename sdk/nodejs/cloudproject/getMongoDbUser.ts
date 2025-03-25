// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a user of a mongodb cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mongoUser = ovh.CloudProject.getMongoDbUser({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ@admin",
 * });
 * export const mongoUserRoles = mongoUser.then(mongoUser => mongoUser.roles);
 * ```
 */
export function getMongoDbUser(args: GetMongoDbUserArgs, opts?: pulumi.InvokeOptions): Promise<GetMongoDbUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getMongoDbUser:getMongoDbUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getMongoDbUser.
 */
export interface GetMongoDbUserArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Name of the user with the authentication database in the format name@authDB, for example: johndoe@admin
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getMongoDbUser.
 */
export interface GetMongoDbUserResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * Date of the creation of the user.
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * Roles the user belongs to
     */
    readonly roles: string[];
    /**
     * Current status of the user.
     */
    readonly serviceName: string;
    /**
     * Current status of the user.
     */
    readonly status: string;
}
/**
 * Use this data source to get information about a user of a mongodb cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mongoUser = ovh.CloudProject.getMongoDbUser({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ@admin",
 * });
 * export const mongoUserRoles = mongoUser.then(mongoUser => mongoUser.roles);
 * ```
 */
export function getMongoDbUserOutput(args: GetMongoDbUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMongoDbUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getMongoDbUser:getMongoDbUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getMongoDbUser.
 */
export interface GetMongoDbUserOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the user with the authentication database in the format name@authDB, for example: johndoe@admin
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
