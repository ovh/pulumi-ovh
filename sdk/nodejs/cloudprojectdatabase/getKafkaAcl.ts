// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
    clusterId: string;
    id: string;
    serviceName: string;
}

/**
 * A collection of values returned by getKafkaAcl.
 */
export interface GetKafkaAclResult {
    readonly clusterId: string;
    readonly id: string;
    readonly permission: string;
    readonly serviceName: string;
    readonly topic: string;
    readonly username: string;
}
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
    clusterId: pulumi.Input<string>;
    id: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
