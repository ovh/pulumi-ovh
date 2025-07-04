// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about an IP firewall.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myFirewall = ovh.Ip.getFirewall({
 *     ip: "XXXXXX",
 *     ipOnFirewall: "XXXXXX",
 * });
 * ```
 */
export function getFirewall(args: GetFirewallArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Ip/getFirewall:getFirewall", {
        "ip": args.ip,
        "ipOnFirewall": args.ipOnFirewall,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewall.
 */
export interface GetFirewallArgs {
    /**
     * The IP or the CIDR
     */
    ip: string;
    /**
     * IPv4 address
     */
    ipOnFirewall: string;
}

/**
 * A collection of values returned by getFirewall.
 */
export interface GetFirewallResult {
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IP or the CIDR
     */
    readonly ip: string;
    /**
     * IPv4 address
     * * `enabled ` - Whether firewall is enabled
     */
    readonly ipOnFirewall: string;
    /**
     * Current state of your ip on firewall
     */
    readonly state: string;
}
/**
 * Use this data source to retrieve information about an IP firewall.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myFirewall = ovh.Ip.getFirewall({
 *     ip: "XXXXXX",
 *     ipOnFirewall: "XXXXXX",
 * });
 * ```
 */
export function getFirewallOutput(args: GetFirewallOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFirewallResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Ip/getFirewall:getFirewall", {
        "ip": args.ip,
        "ipOnFirewall": args.ipOnFirewall,
    }, opts);
}

/**
 * A collection of arguments for invoking getFirewall.
 */
export interface GetFirewallOutputArgs {
    /**
     * The IP or the CIDR
     */
    ip: pulumi.Input<string>;
    /**
     * IPv4 address
     */
    ipOnFirewall: pulumi.Input<string>;
}
