// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ServerNetworking extends pulumi.CustomResource {
    /**
     * Get an existing ServerNetworking resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerNetworkingState, opts?: pulumi.CustomResourceOptions): ServerNetworking {
        return new ServerNetworking(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/serverNetworking:ServerNetworking';

    /**
     * Returns true if the given object is an instance of ServerNetworking.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerNetworking {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerNetworking.__pulumiType;
    }

    /**
     * Operation description
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Interface or interfaces aggregation.
     */
    public readonly interfaces!: pulumi.Output<outputs.Dedicated.ServerNetworkingInterface[]>;
    /**
     * The internal name of your dedicated server.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Operation status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ServerNetworking resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerNetworkingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerNetworkingArgs | ServerNetworkingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerNetworkingState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ServerNetworkingArgs | undefined;
            if ((!args || args.interfaces === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaces'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerNetworking.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerNetworking resources.
 */
export interface ServerNetworkingState {
    /**
     * Operation description
     */
    description?: pulumi.Input<string>;
    /**
     * Interface or interfaces aggregation.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.Dedicated.ServerNetworkingInterface>[]>;
    /**
     * The internal name of your dedicated server.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Operation status
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerNetworking resource.
 */
export interface ServerNetworkingArgs {
    /**
     * Interface or interfaces aggregation.
     */
    interfaces: pulumi.Input<pulumi.Input<inputs.Dedicated.ServerNetworkingInterface>[]>;
    /**
     * The internal name of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
}
