// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getUserS3Credentials(args: GetUserS3CredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserS3CredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getUserS3Credentials:getUserS3Credentials", {
        "serviceName": args.serviceName,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserS3Credentials.
 */
export interface GetUserS3CredentialsArgs {
    serviceName: string;
    userId: string;
}

/**
 * A collection of values returned by getUserS3Credentials.
 */
export interface GetUserS3CredentialsResult {
    readonly accessKeyIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly serviceName: string;
    readonly userId: string;
}
export function getUserS3CredentialsOutput(args: GetUserS3CredentialsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUserS3CredentialsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:CloudProject/getUserS3Credentials:getUserS3Credentials", {
        "serviceName": args.serviceName,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserS3Credentials.
 */
export interface GetUserS3CredentialsOutputArgs {
    serviceName: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}
