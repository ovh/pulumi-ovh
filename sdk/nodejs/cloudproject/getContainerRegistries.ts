// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the container registries of a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const registries = ovh.CloudProject.getContainerRegistries({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getContainerRegistries(args: GetContainerRegistriesArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerRegistriesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getContainerRegistries:getContainerRegistries", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistries.
 */
export interface GetContainerRegistriesArgs {
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getContainerRegistries.
 */
export interface GetContainerRegistriesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of container registries associated with the project.
     */
    readonly results: outputs.CloudProject.GetContainerRegistriesResult[];
    readonly serviceName: string;
}
/**
 * Use this data source to get the container registries of a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const registries = ovh.CloudProject.getContainerRegistries({
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getContainerRegistriesOutput(args: GetContainerRegistriesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerRegistriesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getContainerRegistries:getContainerRegistries", {
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistries.
 */
export interface GetContainerRegistriesOutputArgs {
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
