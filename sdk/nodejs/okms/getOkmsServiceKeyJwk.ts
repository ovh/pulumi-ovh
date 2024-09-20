// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a KMS service key, in the JWK format.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const keyInfo = ovh.Okms.getOkmsServiceKey({
 *     id: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
 *     okmsId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
 * });
 * ```
 */
export function getOkmsServiceKeyJwk(args: GetOkmsServiceKeyJwkArgs, opts?: pulumi.InvokeOptions): Promise<GetOkmsServiceKeyJwkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk", {
        "id": args.id,
        "okmsId": args.okmsId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOkmsServiceKeyJwk.
 */
export interface GetOkmsServiceKeyJwkArgs {
    /**
     * ID of the service key
     */
    id: string;
    /**
     * ID of the KMS
     */
    okmsId: string;
}

/**
 * A collection of values returned by getOkmsServiceKeyJwk.
 */
export interface GetOkmsServiceKeyJwkResult {
    readonly createdAt: string;
    readonly id: string;
    readonly keys: outputs.Okms.GetOkmsServiceKeyJwkKey[];
    readonly name: string;
    readonly okmsId: string;
    readonly size: number;
    readonly state: string;
    readonly type: string;
}
/**
 * Use this data source to retrieve information about a KMS service key, in the JWK format.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const keyInfo = ovh.Okms.getOkmsServiceKey({
 *     id: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
 *     okmsId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
 * });
 * ```
 */
export function getOkmsServiceKeyJwkOutput(args: GetOkmsServiceKeyJwkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOkmsServiceKeyJwkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk", {
        "id": args.id,
        "okmsId": args.okmsId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOkmsServiceKeyJwk.
 */
export interface GetOkmsServiceKeyJwkOutputArgs {
    /**
     * ID of the service key
     */
    id: pulumi.Input<string>;
    /**
     * ID of the KMS
     */
    okmsId: pulumi.Input<string>;
}
