// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getReferenceResourceType(opts?: pulumi.InvokeOptions): Promise<GetReferenceResourceTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Iam/getReferenceResourceType:getReferenceResourceType", {
    }, opts);
}

/**
 * A collection of values returned by getReferenceResourceType.
 */
export interface GetReferenceResourceTypeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly types: string[];
}
export function getReferenceResourceTypeOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetReferenceResourceTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Iam/getReferenceResourceType:getReferenceResourceType", {
    }, opts);
}
