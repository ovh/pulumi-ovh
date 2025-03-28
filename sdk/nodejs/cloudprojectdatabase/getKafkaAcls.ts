// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const acls = ovh.CloudProjectDatabase.getKafkaAcls({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 * });
 * export const aclIds = acls.then(acls => acls.aclIds);
 * ```
 */
export function getKafkaAcls(args: GetKafkaAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetKafkaAclsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getKafkaAcls:getKafkaAcls", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaAcls.
 */
export interface GetKafkaAclsArgs {
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
 * A collection of values returned by getKafkaAcls.
 */
export interface GetKafkaAclsResult {
    /**
     * The list of ACLs ids of the kafka cluster associated with the project.
     */
    readonly aclIds: string[];
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
    readonly serviceName: string;
}
/**
 * Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const acls = ovh.CloudProjectDatabase.getKafkaAcls({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 * });
 * export const aclIds = acls.then(acls => acls.aclIds);
 * ```
 */
export function getKafkaAclsOutput(args: GetKafkaAclsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKafkaAclsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getKafkaAcls:getKafkaAcls", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaAcls.
 */
export interface GetKafkaAclsOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
