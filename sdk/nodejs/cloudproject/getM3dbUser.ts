// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const m3dbUser = ovh.CloudProject.getM3dbUser({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const m3dbUserGroup = m3dbUser.then(m3dbUser => m3dbUser.group);
 * ```
 */
export function getM3dbUser(args: GetM3dbUserArgs, opts?: pulumi.InvokeOptions): Promise<GetM3dbUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getM3dbUser:getM3dbUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getM3dbUser.
 */
export interface GetM3dbUserArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Name of the user.
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getM3dbUser.
 */
export interface GetM3dbUserResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * Date of the creation of the user.
     */
    readonly createdAt: string;
    /**
     * See Argument Reference above.
     */
    readonly group: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
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
 * Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const m3dbUser = ovh.CloudProject.getM3dbUser({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     name: "ZZZ",
 * });
 * export const m3dbUserGroup = m3dbUser.then(m3dbUser => m3dbUser.group);
 * ```
 */
export function getM3dbUserOutput(args: GetM3dbUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetM3dbUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getM3dbUser:getM3dbUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getM3dbUser.
 */
export interface GetM3dbUserOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the user.
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
