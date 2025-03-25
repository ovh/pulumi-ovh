// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the list of users of a container registry associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myRegistry = ovh.CloudProject.getContainerRegistry({
 *     serviceName: "XXXXXX",
 *     registryId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 * });
 * const users = ovh.CloudProject.getContainerRegistryUsers({
 *     serviceName: ovh_cloud_project_containerregistry.my_registry.service_name,
 *     registryId: ovh_cloud_project_containerregistry.my_registry.id,
 * });
 * ```
 */
export function getContainerRegistryUsers(args: GetContainerRegistryUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerRegistryUsersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers", {
        "registryId": args.registryId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistryUsers.
 */
export interface GetContainerRegistryUsersArgs {
    /**
     * Registry ID
     */
    registryId: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getContainerRegistryUsers.
 */
export interface GetContainerRegistryUsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly registryId: string;
    /**
     * The list of users of the container registry associated with the project.
     */
    readonly results: outputs.CloudProject.GetContainerRegistryUsersResult[];
    readonly serviceName: string;
}
/**
 * Use this data source to get the list of users of a container registry associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myRegistry = ovh.CloudProject.getContainerRegistry({
 *     serviceName: "XXXXXX",
 *     registryId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 * });
 * const users = ovh.CloudProject.getContainerRegistryUsers({
 *     serviceName: ovh_cloud_project_containerregistry.my_registry.service_name,
 *     registryId: ovh_cloud_project_containerregistry.my_registry.id,
 * });
 * ```
 */
export function getContainerRegistryUsersOutput(args: GetContainerRegistryUsersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerRegistryUsersResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers", {
        "registryId": args.registryId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainerRegistryUsers.
 */
export interface GetContainerRegistryUsersOutputArgs {
    /**
     * Registry ID
     */
    registryId: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
