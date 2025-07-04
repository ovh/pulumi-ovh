// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a topic of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const topic = ovh.CloudProjectDatabase.getKafkaTopic({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const topicName = topic.then(topic => topic.name);
 * ```
 */
export function getKafkaTopic(args: GetKafkaTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetKafkaTopicResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getKafkaTopic:getKafkaTopic", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaTopic.
 */
export interface GetKafkaTopicArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * Topic ID
     */
    id: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKafkaTopic.
 */
export interface GetKafkaTopicResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * See Argument Reference above.
     */
    readonly id: string;
    /**
     * Minimum insync replica accepted for this topic.
     */
    readonly minInsyncReplicas: number;
    /**
     * Name of the topic.
     */
    readonly name: string;
    /**
     * Number of partitions for this topic.
     */
    readonly partitions: number;
    /**
     * Number of replication for this topic.
     */
    readonly replication: number;
    /**
     * Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
     */
    readonly retentionBytes: number;
    /**
     * Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
     */
    readonly retentionHours: number;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get information about a topic of a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const topic = ovh.CloudProjectDatabase.getKafkaTopic({
 *     serviceName: "XXX",
 *     clusterId: "YYY",
 *     id: "ZZZ",
 * });
 * export const topicName = topic.then(topic => topic.name);
 * ```
 */
export function getKafkaTopicOutput(args: GetKafkaTopicOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKafkaTopicResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getKafkaTopic:getKafkaTopic", {
        "clusterId": args.clusterId,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKafkaTopic.
 */
export interface GetKafkaTopicOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * Topic ID
     */
    id: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
