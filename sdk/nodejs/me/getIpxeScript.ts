// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about an IPXE Script.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const script = ovh.Me.getIpxeScript({
 *     name: "myscript",
 * });
 * ```
 */
export function getIpxeScript(args: GetIpxeScriptArgs, opts?: pulumi.InvokeOptions): Promise<GetIpxeScriptResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Me/getIpxeScript:getIpxeScript", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpxeScript.
 */
export interface GetIpxeScriptArgs {
    /**
     * The name of the IPXE Script.
     */
    name: string;
}

/**
 * A collection of values returned by getIpxeScript.
 */
export interface GetIpxeScriptResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * The content of the script.
     */
    readonly script: string;
}
/**
 * Use this data source to retrieve information about an IPXE Script.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const script = ovh.Me.getIpxeScript({
 *     name: "myscript",
 * });
 * ```
 */
export function getIpxeScriptOutput(args: GetIpxeScriptOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpxeScriptResult> {
    return pulumi.output(args).apply((a: any) => getIpxeScript(a, opts))
}

/**
 * A collection of arguments for invoking getIpxeScript.
 */
export interface GetIpxeScriptOutputArgs {
    /**
     * The name of the IPXE Script.
     */
    name: pulumi.Input<string>;
}
