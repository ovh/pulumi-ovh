// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getVrackNetwork(args: GetVrackNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVrackNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", {
        "serviceName": args.serviceName,
        "vrackNetworkId": args.vrackNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrackNetwork.
 */
export interface GetVrackNetworkArgs {
    serviceName: string;
    vrackNetworkId: number;
}

/**
 * A collection of values returned by getVrackNetwork.
 */
export interface GetVrackNetworkResult {
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly natIp: string;
    readonly serviceName: string;
    readonly subnet: string;
    readonly vlan: number;
    readonly vrackNetworkId: number;
}
export function getVrackNetworkOutput(args: GetVrackNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVrackNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", {
        "serviceName": args.serviceName,
        "vrackNetworkId": args.vrackNetworkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrackNetwork.
 */
export interface GetVrackNetworkOutputArgs {
    serviceName: pulumi.Input<string>;
    vrackNetworkId: pulumi.Input<number>;
}
