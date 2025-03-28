// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get the details of a public cloud project loadbalancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lbLoadBalancer = ovh.CloudProject.getLoadBalancer({
 *     serviceName: "XXXXXX",
 *     regionName: "XXX",
 *     id: "XXX",
 * });
 * export const lb = lbLoadBalancer;
 * ```
 */
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
    /**
     * ID of the loadbalancer
     */
    id: string;
    /**
     * Region of the loadbalancer.
     */
    regionName: string;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    /**
     * Date of creation of the loadbalancer
     */
    readonly createdAt: string;
    /**
     * ID of the flavor
     */
    readonly flavorId: string;
    /**
     * Information about the floating IP
     */
    readonly floatingIp: outputs.CloudProject.GetLoadBalancerFloatingIp;
    /**
     * ID of the floating IP
     */
    readonly id: string;
    /**
     * Name of the loadbalancer
     */
    readonly name: string;
    /**
     * Operating status of the loadbalancer
     */
    readonly operatingStatus: string;
    /**
     * Provisioning status of the loadbalancer
     */
    readonly provisioningStatus: string;
    /**
     * Region of the loadbalancer
     */
    readonly regionName: string;
    /**
     * ID of the public cloud project
     */
    readonly serviceName: string;
    /**
     * Last update date of the loadbalancer
     */
    readonly updatedAt: string;
    /**
     * IP address of the Virtual IP
     */
    readonly vipAddress: string;
    /**
     * Openstack ID of the network for the Virtual IP
     */
    readonly vipNetworkId: string;
    /**
     * ID of the subnet for the Virtual IP
     */
    readonly vipSubnetId: string;
}
/**
 * Get the details of a public cloud project loadbalancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lbLoadBalancer = ovh.CloudProject.getLoadBalancer({
 *     serviceName: "XXXXXX",
 *     regionName: "XXX",
 *     id: "XXX",
 * });
 * export const lb = lbLoadBalancer;
 * ```
 */
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
    /**
     * ID of the loadbalancer
     */
    id: pulumi.Input<string>;
    /**
     * Region of the loadbalancer.
     */
    regionName: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
