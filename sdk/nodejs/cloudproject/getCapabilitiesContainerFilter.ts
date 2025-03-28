// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const capability = ovh.CloudProject.getCapabilitiesContainerFilter({
 *     planName: "SMALL",
 *     region: "GRA",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getCapabilitiesContainerFilter(args: GetCapabilitiesContainerFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetCapabilitiesContainerFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", {
        "planName": args.planName,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCapabilitiesContainerFilter.
 */
export interface GetCapabilitiesContainerFilterArgs {
    /**
     * The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
     */
    planName: string;
    /**
     * The region name
     */
    region: string;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getCapabilitiesContainerFilter.
 */
export interface GetCapabilitiesContainerFilterResult {
    /**
     * Plan code from the catalog
     */
    readonly code: string;
    /**
     * Plan creation date
     */
    readonly createdAt: string;
    /**
     * Features of the plan
     */
    readonly features: outputs.CloudProject.GetCapabilitiesContainerFilterFeature[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Plan name
     */
    readonly name: string;
    readonly planName: string;
    readonly region: string;
    /**
     * Container registry limits
     */
    readonly registryLimits: outputs.CloudProject.GetCapabilitiesContainerFilterRegistryLimit[];
    readonly serviceName: string;
    /**
     * Plan last update date
     */
    readonly updatedAt: string;
}
/**
 * Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const capability = ovh.CloudProject.getCapabilitiesContainerFilter({
 *     planName: "SMALL",
 *     region: "GRA",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export function getCapabilitiesContainerFilterOutput(args: GetCapabilitiesContainerFilterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCapabilitiesContainerFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", {
        "planName": args.planName,
        "region": args.region,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getCapabilitiesContainerFilter.
 */
export interface GetCapabilitiesContainerFilterOutputArgs {
    /**
     * The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
     */
    planName: pulumi.Input<string>;
    /**
     * The region name
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
