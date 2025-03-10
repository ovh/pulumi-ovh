// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getDatabaseInstance(args: GetDatabaseInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getDatabaseInstance:getDatabaseInstance", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseInstance.
 */
export interface GetDatabaseInstanceArgs {
    clusterId: string;
    engine: string;
    name: string;
    serviceName: string;
}

/**
 * A collection of values returned by getDatabaseInstance.
 */
export interface GetDatabaseInstanceResult {
    readonly clusterId: string;
    readonly default: boolean;
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly serviceName: string;
}
export function getDatabaseInstanceOutput(args: GetDatabaseInstanceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatabaseInstanceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getDatabaseInstance:getDatabaseInstance", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseInstance.
 */
export interface GetDatabaseInstanceOutputArgs {
    clusterId: pulumi.Input<string>;
    engine: pulumi.Input<string>;
    name: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
