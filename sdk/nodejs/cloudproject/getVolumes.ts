// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get all the volume from a region of a public cloud project
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const volume = ovh.CloudProject.getVolume({
 *     regionName: "xxx",
 *     serviceName: "yyy",
 * });
 * ```
 */
export function getVolumes(args: GetVolumesArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getVolumes:getVolumes", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolumes.
 */
export interface GetVolumesArgs {
    /**
     * A valid OVHcloud public cloud region name in which the volumes are available. Ex.: "GRA11".
     */
    regionName: string;
    /**
     * The id of the public cloud project.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getVolumes.
 */
export interface GetVolumesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The region name where volumes are available
     */
    readonly regionName: string;
    /**
     * The id of the public cloud project.
     */
    readonly serviceName: string;
    readonly volumes: outputs.CloudProject.GetVolumesVolume[];
}
/**
 * Get all the volume from a region of a public cloud project
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const volume = ovh.CloudProject.getVolume({
 *     regionName: "xxx",
 *     serviceName: "yyy",
 * });
 * ```
 */
export function getVolumesOutput(args: GetVolumesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getVolumes:getVolumes", {
        "regionName": args.regionName,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolumes.
 */
export interface GetVolumesOutputArgs {
    /**
     * A valid OVHcloud public cloud region name in which the volumes are available. Ex.: "GRA11".
     */
    regionName: pulumi.Input<string>;
    /**
     * The id of the public cloud project.
     */
    serviceName: pulumi.Input<string>;
}