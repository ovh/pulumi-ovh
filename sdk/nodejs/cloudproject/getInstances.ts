// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * **This datasource uses a Beta API**
 *
 * Use this data source to get the list of instances in a region of a public cloud project.
 *
 * ## Example Usage
 *
 * To list your instances:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const instance = ovh.CloudProject.getInstances({
 *     region: "XXXX",
 *     serviceName: "YYYY",
 * });
 * ```
 */
export function getInstances(args: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getInstances:getInstances", {
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * Instance region.
     */
    region: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of instances
     */
    readonly instances: outputs.CloudProject.GetInstancesInstance[];
    readonly region: string;
    readonly serviceName: string;
}
/**
 * **This datasource uses a Beta API**
 *
 * Use this data source to get the list of instances in a region of a public cloud project.
 *
 * ## Example Usage
 *
 * To list your instances:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const instance = ovh.CloudProject.getInstances({
 *     region: "XXXX",
 *     serviceName: "YYYY",
 * });
 * ```
 */
export function getInstancesOutput(args: GetInstancesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstancesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getInstances:getInstances", {
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * Instance region.
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
