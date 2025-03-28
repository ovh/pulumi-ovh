// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get the S3 Policy of a public cloud project user. The policy can be set by using the `ovh.CloudProject.S3Policy` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const projectUsers = ovh.CloudProject.getUsers({
 *     serviceName: "XXX",
 * });
 * // Get the user ID of a previously created user with the description "S3-User"
 * const users = projectUsers.then(projectUsers => .filter(user => user.description == "S3-User").map(user => (user.userId)));
 * const s3UserId = users[0];
 * const policy = Promise.all([projectUsers, s3UserId]).then(([projectUsers, s3UserId]) => ovh.CloudProject.getUserS3Policy({
 *     serviceName: projectUsers.serviceName,
 *     userId: s3UserId,
 * }));
 * ```
 */
export function getUserS3Policy(args: GetUserS3PolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetUserS3PolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getUserS3Policy:getUserS3Policy", {
        "serviceName": args.serviceName,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserS3Policy.
 */
export interface GetUserS3PolicyArgs {
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
    /**
     * The ID of a public cloud project's user.
     */
    userId: string;
}

/**
 * A collection of values returned by getUserS3Policy.
 */
export interface GetUserS3PolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required) The policy document. This is a JSON formatted string.
     */
    readonly policy: string;
    readonly serviceName: string;
    readonly userId: string;
}
/**
 * Get the S3 Policy of a public cloud project user. The policy can be set by using the `ovh.CloudProject.S3Policy` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const projectUsers = ovh.CloudProject.getUsers({
 *     serviceName: "XXX",
 * });
 * // Get the user ID of a previously created user with the description "S3-User"
 * const users = projectUsers.then(projectUsers => .filter(user => user.description == "S3-User").map(user => (user.userId)));
 * const s3UserId = users[0];
 * const policy = Promise.all([projectUsers, s3UserId]).then(([projectUsers, s3UserId]) => ovh.CloudProject.getUserS3Policy({
 *     serviceName: projectUsers.serviceName,
 *     userId: s3UserId,
 * }));
 * ```
 */
export function getUserS3PolicyOutput(args: GetUserS3PolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserS3PolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getUserS3Policy:getUserS3Policy", {
        "serviceName": args.serviceName,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserS3Policy.
 */
export interface GetUserS3PolicyOutputArgs {
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The ID of a public cloud project's user.
     */
    userId: pulumi.Input<string>;
}
