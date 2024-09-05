// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information about a subscription to a Managed Loadbalancer Logs Service in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const sub = ovh.CloudProject.getRegionLoadBalancerLogSubscription({
 *     loadbalancerId: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
 *     regionName: "gggg",
 *     serviceName: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     subscriptionId: "zzzzzzzz-yyyy-xxxx-wwww-vvvvvvvvvvvv",
 * });
 * ```
 */
export function getRegionLoadBalancerLogSubscription(args: GetRegionLoadBalancerLogSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionLoadBalancerLogSubscriptionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getRegionLoadBalancerLogSubscription:getRegionLoadBalancerLogSubscription", {
        "loadbalancerId": args.loadbalancerId,
        "regionName": args.regionName,
        "serviceName": args.serviceName,
        "subscriptionId": args.subscriptionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegionLoadBalancerLogSubscription.
 */
export interface GetRegionLoadBalancerLogSubscriptionArgs {
    /**
     * Loadbalancer id to get the logs
     */
    loadbalancerId: string;
    /**
     * A valid OVHcloud public cloud region name in which the loadbalancer is available. Ex.: "GRA11".
     */
    regionName: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
    /**
     * Subscription id
     */
    subscriptionId: string;
}

/**
 * A collection of values returned by getRegionLoadBalancerLogSubscription.
 */
export interface GetRegionLoadBalancerLogSubscriptionResult {
    /**
     * The date of the subscription creation
     */
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Router used for forwarding log
     */
    readonly kind: string;
    /**
     * LDP service name
     */
    readonly ldpServiceName: string;
    /**
     * Loadbalancer id to get the logs
     */
    readonly loadbalancerId: string;
    /**
     * A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: "GRA11".
     */
    readonly regionName: string;
    /**
     * The resource name
     */
    readonly resourceName: string;
    /**
     * The resource type
     */
    readonly resourceType: string;
    /**
     * The id of the public cloud project.
     */
    readonly serviceName: string;
    /**
     * Data stream id to use for the subscription
     */
    readonly streamId: string;
    /**
     * The subscription id
     */
    readonly subscriptionId: string;
    /**
     * The last update of the subscription
     */
    readonly updatedAt: string;
}
/**
 * Get information about a subscription to a Managed Loadbalancer Logs Service in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const sub = ovh.CloudProject.getRegionLoadBalancerLogSubscription({
 *     loadbalancerId: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
 *     regionName: "gggg",
 *     serviceName: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 *     subscriptionId: "zzzzzzzz-yyyy-xxxx-wwww-vvvvvvvvvvvv",
 * });
 * ```
 */
export function getRegionLoadBalancerLogSubscriptionOutput(args: GetRegionLoadBalancerLogSubscriptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionLoadBalancerLogSubscriptionResult> {
    return pulumi.output(args).apply((a: any) => getRegionLoadBalancerLogSubscription(a, opts))
}

/**
 * A collection of arguments for invoking getRegionLoadBalancerLogSubscription.
 */
export interface GetRegionLoadBalancerLogSubscriptionOutputArgs {
    /**
     * Loadbalancer id to get the logs
     */
    loadbalancerId: pulumi.Input<string>;
    /**
     * A valid OVHcloud public cloud region name in which the loadbalancer is available. Ex.: "GRA11".
     */
    regionName: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Subscription id
     */
    subscriptionId: pulumi.Input<string>;
}