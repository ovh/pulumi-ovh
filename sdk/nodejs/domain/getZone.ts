// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getZone(args: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Domain/getZone:getZone", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    name: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    readonly ZoneURN: string;
    readonly dnssecSupported: boolean;
    readonly hasDnsAnycast: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastUpdate: string;
    readonly name: string;
    readonly nameServers: string[];
}
export function getZoneOutput(args: GetZoneOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Domain/getZone:getZone", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    name: pulumi.Input<string>;
}
