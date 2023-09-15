// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to list the IAM action associated with a resource type.
 */
export function getReferenceActions(args: GetReferenceActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetReferenceActionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Iam/getReferenceActions:getReferenceActions", {
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getReferenceActions.
 */
export interface GetReferenceActionsArgs {
    /**
     * Kind of resource we want the actions for
     */
    type: string;
}

/**
 * A collection of values returned by getReferenceActions.
 */
export interface GetReferenceActionsResult {
    /**
     * List of actions
     */
    readonly actions: outputs.Iam.GetReferenceActionsAction[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly type: string;
}
/**
 * Use this data source to list the IAM action associated with a resource type.
 */
export function getReferenceActionsOutput(args: GetReferenceActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReferenceActionsResult> {
    return pulumi.output(args).apply((a: any) => getReferenceActions(a, opts))
}

/**
 * A collection of arguments for invoking getReferenceActions.
 */
export interface GetReferenceActionsOutputArgs {
    /**
     * Kind of resource we want the actions for
     */
    type: pulumi.Input<string>;
}
