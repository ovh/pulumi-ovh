// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Iam/getPolicy:getPolicy", {
        "allows": args.allows,
        "denies": args.denies,
        "description": args.description,
        "excepts": args.excepts,
        "id": args.id,
        "permissionsGroups": args.permissionsGroups,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    allows?: string[];
    denies?: string[];
    description?: string;
    excepts?: string[];
    id: string;
    permissionsGroups?: string[];
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    readonly allows?: string[];
    readonly createdAt: string;
    readonly denies?: string[];
    readonly description?: string;
    readonly excepts?: string[];
    readonly id: string;
    readonly identities: string[];
    readonly name: string;
    readonly owner: string;
    readonly permissionsGroups?: string[];
    readonly readOnly: boolean;
    readonly resources: string[];
    readonly updatedAt: string;
}
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Iam/getPolicy:getPolicy", {
        "allows": args.allows,
        "denies": args.denies,
        "description": args.description,
        "excepts": args.excepts,
        "id": args.id,
        "permissionsGroups": args.permissionsGroups,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyOutputArgs {
    allows?: pulumi.Input<pulumi.Input<string>[]>;
    denies?: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    excepts?: pulumi.Input<pulumi.Input<string>[]>;
    id: pulumi.Input<string>;
    permissionsGroups?: pulumi.Input<pulumi.Input<string>[]>;
}
