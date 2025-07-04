// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the details of a failover IP address of a service in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myFailoverIp = ovh.CloudProject.getFailoverIpAttach({
 *     serviceName: "XXXXXX",
 *     ip: "XXXXXX",
 * });
 * ```
 */
export function getFailoverIpAttach(args: GetFailoverIpAttachArgs, opts?: pulumi.InvokeOptions): Promise<GetFailoverIpAttachResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", {
        "block": args.block,
        "continentCode": args.continentCode,
        "geoLoc": args.geoLoc,
        "ip": args.ip,
        "routedTo": args.routedTo,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getFailoverIpAttach.
 */
export interface GetFailoverIpAttachArgs {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    block?: string;
    continentCode?: string;
    geoLoc?: string;
    /**
     * The failover ip address to query
     */
    ip?: string;
    routedTo?: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getFailoverIpAttach.
 */
export interface GetFailoverIpAttachResult {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    readonly block: string;
    readonly continentCode: string;
    readonly geoLoc: string;
    /**
     * The Ip id
     */
    readonly id: string;
    /**
     * The Ip Address
     */
    readonly ip: string;
    /**
     * Current operation progress in percent
     * * `routedTo` - Instance where ip is routed to
     */
    readonly progress: number;
    readonly routedTo: string;
    readonly serviceName: string;
    /**
     * Ip status, can be `ok` or `operationPending`
     * * `subType` - IP sub type, can be `cloud` or `ovh`
     */
    readonly status: string;
    readonly subType: string;
}
/**
 * Use this data source to get the details of a failover IP address of a service in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myFailoverIp = ovh.CloudProject.getFailoverIpAttach({
 *     serviceName: "XXXXXX",
 *     ip: "XXXXXX",
 * });
 * ```
 */
export function getFailoverIpAttachOutput(args: GetFailoverIpAttachOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFailoverIpAttachResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", {
        "block": args.block,
        "continentCode": args.continentCode,
        "geoLoc": args.geoLoc,
        "ip": args.ip,
        "routedTo": args.routedTo,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getFailoverIpAttach.
 */
export interface GetFailoverIpAttachOutputArgs {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    block?: pulumi.Input<string>;
    continentCode?: pulumi.Input<string>;
    geoLoc?: pulumi.Input<string>;
    /**
     * The failover ip address to query
     */
    ip?: pulumi.Input<string>;
    routedTo?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
