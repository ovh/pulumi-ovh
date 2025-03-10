// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPrivateDatabaseUser(args: GetPrivateDatabaseUserArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateDatabaseUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser", {
        "serviceName": args.serviceName,
        "userName": args.userName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateDatabaseUser.
 */
export interface GetPrivateDatabaseUserArgs {
    serviceName: string;
    userName: string;
}

/**
 * A collection of values returned by getPrivateDatabaseUser.
 */
export interface GetPrivateDatabaseUserResult {
    readonly creationDate: string;
    readonly databases: outputs.Hosting.GetPrivateDatabaseUserDatabase[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly serviceName: string;
    readonly userName: string;
}
export function getPrivateDatabaseUserOutput(args: GetPrivateDatabaseUserOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrivateDatabaseUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser", {
        "serviceName": args.serviceName,
        "userName": args.userName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateDatabaseUser.
 */
export interface GetPrivateDatabaseUserOutputArgs {
    serviceName: pulumi.Input<string>;
    userName: pulumi.Input<string>;
}
