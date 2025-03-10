// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRedisUser(args: GetRedisUserArgs, opts?: pulumi.InvokeOptions): Promise<GetRedisUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getRedisUser:getRedisUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisUser.
 */
export interface GetRedisUserArgs {
    clusterId: string;
    name: string;
    serviceName: string;
}

/**
 * A collection of values returned by getRedisUser.
 */
export interface GetRedisUserResult {
    readonly categories: string[];
    readonly channels: string[];
    readonly clusterId: string;
    readonly commands: string[];
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: string[];
    readonly name: string;
    readonly serviceName: string;
    readonly status: string;
}
export function getRedisUserOutput(args: GetRedisUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRedisUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getRedisUser:getRedisUser", {
        "clusterId": args.clusterId,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisUser.
 */
export interface GetRedisUserOutputArgs {
    clusterId: pulumi.Input<string>;
    name: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
