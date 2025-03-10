// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getDatabasePostgreSQLConnectionPools(args: GetDatabasePostgreSQLConnectionPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasePostgreSQLConnectionPoolsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabasePostgreSQLConnectionPools.
 */
export interface GetDatabasePostgreSQLConnectionPoolsArgs {
    clusterId: string;
    serviceName: string;
}

/**
 * A collection of values returned by getDatabasePostgreSQLConnectionPools.
 */
export interface GetDatabasePostgreSQLConnectionPoolsResult {
    readonly clusterId: string;
    readonly connectionPoolIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly serviceName: string;
}
export function getDatabasePostgreSQLConnectionPoolsOutput(args: GetDatabasePostgreSQLConnectionPoolsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatabasePostgreSQLConnectionPoolsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", {
        "clusterId": args.clusterId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabasePostgreSQLConnectionPools.
 */
export interface GetDatabasePostgreSQLConnectionPoolsOutputArgs {
    clusterId: pulumi.Input<string>;
    serviceName: pulumi.Input<string>;
}
