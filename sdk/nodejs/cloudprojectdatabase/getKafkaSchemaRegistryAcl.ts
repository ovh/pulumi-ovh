// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a schema registry ACL of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const schemaRegistryAcl = ovh.CloudProjectDatabase.getKafkaSchemaRegistryAcl({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const aclPermission = schemaRegistryAcl.then(schemaRegistryAcl => schemaRegistryAcl.permission);
 * ```
 */
export function getKafkaSchemaRegistryAcl(args: GetKafkaSchemaRegistryAclArgs, opts?: pulumi.InvokeOptions): Promise<GetKafkaSchemaRegistryAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaSchemaRegistryAcl.
 */
export interface GetKafkaSchemaRegistryAclArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Schema registry ACL ID
     */
    id: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKafkaSchemaRegistryAcl.
 */
export interface GetKafkaSchemaRegistryAclResult {
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
     * Resource affected by this ACL.
     */
    readonly resource: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * Username affected by this ACL.
     */
    readonly username: string;
}
/**
 * Use this data source to get information about a schema registry ACL of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const schemaRegistryAcl = ovh.CloudProjectDatabase.getKafkaSchemaRegistryAcl({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const aclPermission = schemaRegistryAcl.then(schemaRegistryAcl => schemaRegistryAcl.permission);
 * ```
 */
export function getKafkaSchemaRegistryAclOutput(args: GetKafkaSchemaRegistryAclOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKafkaSchemaRegistryAclResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaSchemaRegistryAcl.
 */
export interface GetKafkaSchemaRegistryAclOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Schema registry ACL ID
     */
    id: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
