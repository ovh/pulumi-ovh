// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about an IP service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myIp = ovh.Ip.getService({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Ip/getService:getService", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * The service name
     */
    serviceName: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    /**
     * can be terminated
     */
    readonly canBeTerminated: boolean;
    /**
     * country
     */
    readonly country: string;
    /**
     * Custom description on your ip
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ip block
     */
    readonly ip: string;
    /**
     * IP block organisation Id
     */
    readonly organisationId: string;
    /**
     * Routage information
     */
    readonly routedTos: outputs.Ip.GetServiceRoutedTo[];
    /**
     * Service where ip is routed to
     */
    readonly serviceName: string;
    /**
     * Possible values for ip type ( "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
     */
    readonly type: string;
}
/**
 * Use this data source to retrieve information about an IP service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myIp = ovh.Ip.getService({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Ip/getService:getService", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceOutputArgs {
    /**
     * The service name
     */
    serviceName: pulumi.Input<string>;
}
