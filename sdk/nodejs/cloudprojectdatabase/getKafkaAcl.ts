// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an ACL of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const acl = ovh.CloudProjectDatabase.getKafkaAcl({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const aclPermission = acl.then(acl => acl.permission);
 * ```
 */
export function getKafkaAcl(args: GetKafkaAclArgs, opts?: pulumi.InvokeOptions): Promise<GetKafkaAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getKafkaAcl:getKafkaAcl", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaAcl.
 */
export interface GetKafkaAclArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * ACL ID
     */
    id: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKafkaAcl.
 */
export interface GetKafkaAclResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * See Argument Reference above.
     */
    readonly id: string;
    /**
     * Permission to give to this username on this topic.
     */
    readonly permission: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * Topic affected by this ACL.
     */
    readonly topic: string;
    /**
     * Username affected by this ACL.
     */
    readonly username: string;
}
/**
 * Use this data source to get information about an ACL of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const acl = ovh.CloudProjectDatabase.getKafkaAcl({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const aclPermission = acl.then(acl => acl.permission);
 * ```
 */
export function getKafkaAclOutput(args: GetKafkaAclOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKafkaAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getKafkaAcl:getKafkaAcl", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaAcl.
 */
export interface GetKafkaAclOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * ACL ID
     */
    id: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
