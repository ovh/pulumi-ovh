// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of integrations of a database cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const integrations = ovh.CloudProjectDatabase.getDatabaseIntegrations({
 *     serviceName: "XXX",
 *     engine: "YYY",
 *     clusterId: "ZZZ",
 * });
 * export const integrationIds = integrations.then(integrations => integrations.integrationIds);
 * ```
 */
export function getDatabaseIntegrations(args: GetDatabaseIntegrationsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseIntegrationsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getDatabaseIntegrations:getDatabaseIntegrations", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseIntegrations.
 */
export interface GetDatabaseIntegrationsArgs {
    /**
     * Cluster ID
     */
    clusterId: string;
    /**
     * The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     */
    engine: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getDatabaseIntegrations.
 */
export interface GetDatabaseIntegrationsResult {
    /**
     * See Argument Reference above.
     */
    readonly clusterId: string;
    /**
     * See Argument Reference above.
     */
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of integrations ids of the database cluster associated with the project.
     */
    readonly integrationIds: string[];
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get the list of integrations of a database cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const integrations = ovh.CloudProjectDatabase.getDatabaseIntegrations({
 *     serviceName: "XXX",
 *     engine: "YYY",
 *     clusterId: "ZZZ",
 * });
 * export const integrationIds = integrations.then(integrations => integrations.integrationIds);
 * ```
 */
export function getDatabaseIntegrationsOutput(args: GetDatabaseIntegrationsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatabaseIntegrationsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProjectDatabase/getDatabaseIntegrations:getDatabaseIntegrations", {
        "clusterId": args.clusterId,
        "engine": args.engine,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseIntegrations.
 */
export interface GetDatabaseIntegrationsOutputArgs {
    /**
     * Cluster ID
     */
    clusterId: pulumi.Input<string>;
    /**
     * The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     */
    engine: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
