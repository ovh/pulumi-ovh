// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attach a Dedicated Cloud to the vrack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const vrack_dedicatedCloud = new ovh.vrack.VrackDedicatedCloud("vrack-dedicatedCloud", {
 *     serviceName: "<vRack service name>",
 *     dedicatedCloud: "<Dedicated Cloud service name>",
 * });
 * ```
 *
 * ## Import
 *
 * Attachment of a Dedicated Cloud and a vRack can be imported using the `service_name` (vRack identifier) and the `dedicated_cloud` (Dedicated Cloud service name), separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:Vrack/vrackDedicatedCloud:VrackDedicatedCloud myattach "<vRack service name>/<Dedicated Cloud service name>"
 * ```
 */
export class VrackDedicatedCloud extends pulumi.CustomResource {
    /**
     * Get an existing VrackDedicatedCloud resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VrackDedicatedCloudState, opts?: pulumi.CustomResourceOptions): VrackDedicatedCloud {
        return new VrackDedicatedCloud(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Vrack/vrackDedicatedCloud:VrackDedicatedCloud';

    /**
     * Returns true if the given object is an instance of VrackDedicatedCloud.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VrackDedicatedCloud {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VrackDedicatedCloud.__pulumiType;
    }

    /**
     * Your Dedicated Cloud service name
     */
    public readonly dedicatedCloud!: pulumi.Output<string>;
    /**
     * The internal name of your vrack
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a VrackDedicatedCloud resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VrackDedicatedCloudArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VrackDedicatedCloudArgs | VrackDedicatedCloudState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VrackDedicatedCloudState | undefined;
            resourceInputs["dedicatedCloud"] = state ? state.dedicatedCloud : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as VrackDedicatedCloudArgs | undefined;
            if ((!args || args.dedicatedCloud === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dedicatedCloud'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["dedicatedCloud"] = args ? args.dedicatedCloud : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VrackDedicatedCloud.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VrackDedicatedCloud resources.
 */
export interface VrackDedicatedCloudState {
    /**
     * Your Dedicated Cloud service name
     */
    dedicatedCloud?: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VrackDedicatedCloud resource.
 */
export interface VrackDedicatedCloudArgs {
    /**
     * Your Dedicated Cloud service name
     */
    dedicatedCloud: pulumi.Input<string>;
    /**
     * The internal name of your vrack
     */
    serviceName: pulumi.Input<string>;
}
