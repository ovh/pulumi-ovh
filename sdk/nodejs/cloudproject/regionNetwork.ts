// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a network in a public cloud project.
 */
export class RegionNetwork extends pulumi.CustomResource {
    /**
     * Get an existing RegionNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionNetworkState, opts?: pulumi.CustomResourceOptions): RegionNetwork {
        return new RegionNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/regionNetwork:RegionNetwork';

    /**
     * Returns true if the given object is an instance of RegionNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionNetwork.__pulumiType;
    }

    /**
     * Name of the network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network region returned by the API
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Network region
     */
    public readonly regionName!: pulumi.Output<string>;
    /**
     * The id of the public cloud project
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Parameters to create a subnet
     */
    public readonly subnet!: pulumi.Output<outputs.CloudProject.RegionNetworkSubnet>;
    /**
     * Network visibility
     */
    public /*out*/ readonly visibility!: pulumi.Output<string>;
    /**
     * VLAN ID, between 1 and 4000
     */
    public readonly vlanId!: pulumi.Output<number>;

    /**
     * Create a RegionNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionNetworkArgs | RegionNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionNetworkState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["regionName"] = state ? state.regionName : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
            resourceInputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as RegionNetworkArgs | undefined;
            if ((!args || args.regionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionName'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionName"] = args ? args.regionName : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionNetwork resources.
 */
export interface RegionNetworkState {
    /**
     * Name of the network
     */
    name?: pulumi.Input<string>;
    /**
     * Network region returned by the API
     */
    region?: pulumi.Input<string>;
    /**
     * Network region
     */
    regionName?: pulumi.Input<string>;
    /**
     * The id of the public cloud project
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Parameters to create a subnet
     */
    subnet?: pulumi.Input<inputs.CloudProject.RegionNetworkSubnet>;
    /**
     * Network visibility
     */
    visibility?: pulumi.Input<string>;
    /**
     * VLAN ID, between 1 and 4000
     */
    vlanId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RegionNetwork resource.
 */
export interface RegionNetworkArgs {
    /**
     * Name of the network
     */
    name?: pulumi.Input<string>;
    /**
     * Network region
     */
    regionName: pulumi.Input<string>;
    /**
     * The id of the public cloud project
     */
    serviceName: pulumi.Input<string>;
    /**
     * Parameters to create a subnet
     */
    subnet: pulumi.Input<inputs.CloudProject.RegionNetworkSubnet>;
    /**
     * VLAN ID, between 1 and 4000
     */
    vlanId?: pulumi.Input<number>;
}
