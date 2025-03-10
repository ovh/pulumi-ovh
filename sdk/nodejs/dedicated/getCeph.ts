// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getCeph(args: GetCephArgs, opts?: pulumi.InvokeOptions): Promise<GetCephResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dedicated/getCeph:getCeph", {
        "cephVersion": args.cephVersion,
        "serviceName": args.serviceName,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getCeph.
 */
export interface GetCephArgs {
    cephVersion?: string;
    serviceName: string;
    status?: string;
}

/**
 * A collection of values returned by getCeph.
 */
export interface GetCephResult {
    readonly CephURN: string;
    readonly cephMons: string[];
    readonly cephVersion: string;
    readonly crushTunables: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly label: string;
    readonly region: string;
    readonly serviceName: string;
    readonly size: number;
    readonly state: string;
    readonly status: string;
}
export function getCephOutput(args: GetCephOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCephResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Dedicated/getCeph:getCeph", {
        "cephVersion": args.cephVersion,
        "serviceName": args.serviceName,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getCeph.
 */
export interface GetCephOutputArgs {
    cephVersion?: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
