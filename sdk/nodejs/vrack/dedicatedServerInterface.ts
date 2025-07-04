// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attach a Dedicated Server Network Interface to a vRack.
 *
 * > **NOTE:** The resource `ovh.Vrack.DedicatedServerInterface` is intended to be used for dedicated servers that have configurable network interfaces.<br /> Legacy Dedicated servers that do not have configurable network interfaces MUST use the resource `ovh.Vrack.DedicatedServer` instead.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const server = ovh.getServer({
 *     serviceName: "nsxxxxxxx.ip-xx-xx-xx.eu",
 * });
 * const vdsi = new ovh.vrack.DedicatedServerInterface("vdsi", {
 *     serviceName: "pn-xxxxxxx",
 *     interfaceId: server.then(server => server.enabledVrackVnis?.[0]),
 * });
 * ```
 */
export class DedicatedServerInterface extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedServerInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedServerInterfaceState, opts?: pulumi.CustomResourceOptions): DedicatedServerInterface {
        return new DedicatedServerInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface';

    /**
     * Returns true if the given object is an instance of DedicatedServerInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedServerInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedServerInterface.__pulumiType;
    }

    /**
     * The id of dedicated server network interface.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a DedicatedServerInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedServerInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedServerInterfaceArgs | DedicatedServerInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedServerInterfaceState | undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as DedicatedServerInterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DedicatedServerInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedServerInterface resources.
 */
export interface DedicatedServerInterfaceState {
    /**
     * The id of dedicated server network interface.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedServerInterface resource.
 */
export interface DedicatedServerInterfaceArgs {
    /**
     * The id of dedicated server network interface.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}
