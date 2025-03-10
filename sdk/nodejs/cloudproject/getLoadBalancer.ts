// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLoadBalancer(args: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getLoadBalancer:getLoadBalancer", {
        "id": args.id,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerArgs {
    id: string;
    regionName: string;
    serviceName: string;
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    readonly createdAt: string;
    readonly flavorId: string;
    readonly floatingIp: outputs.CloudProject.GetLoadBalancerFloatingIp;
    readonly id: string;
    readonly name: string;
    readonly operatingStatus: string;
    readonly provisioningStatus: string;
    readonly regionName: string;
    readonly serviceName: string;
    readonly updatedAt: string;
    readonly vipAddress: string;
    readonly vipNetworkId: string;
    readonly vipSubnetId: string;
}
export function getLoadBalancerOutput(args: GetLoadBalancerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLoadBalancerResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getLoadBalancer:getLoadBalancer", {
        "id": args.id,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerOutputArgs {
    id: pulumi.Input<string>;
    regionName: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
