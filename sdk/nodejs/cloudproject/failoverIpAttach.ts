// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a failover IP address to a compute instance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myFailoverIp = new ovh.cloudproject.FailoverIpAttach("myFailoverIp", {
 *     ip: "XXXXXX",
 *     routedTo: "XXXXXX",
 *     serviceName: "XXXXXX",
 * });
 * ```
 */
export class FailoverIpAttach extends pulumi.CustomResource {
    /**
     * Get an existing FailoverIpAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FailoverIpAttachState, opts?: pulumi.CustomResourceOptions): FailoverIpAttach {
        return new FailoverIpAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/failoverIpAttach:FailoverIpAttach';

    /**
     * Returns true if the given object is an instance of FailoverIpAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FailoverIpAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FailoverIpAttach.__pulumiType;
    }

    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    public readonly block!: pulumi.Output<string>;
    /**
     * Ip continent
     */
    public readonly continentCode!: pulumi.Output<string>;
    /**
     * Ip location
     */
    public readonly geoLoc!: pulumi.Output<string>;
    /**
     * The failover ip address to attach
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Current operation progress in percent
     * * `routedTo` - Instance where ip is routed to
     */
    public /*out*/ readonly progress!: pulumi.Output<number>;
    /**
     * The GUID of an instance to which the failover IP address is be attached
     */
    public readonly routedTo!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Ip status, can be `ok` or `operationPending`
     * * `subType` - IP sub type, can be `cloud` or `ovh`
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * IP sub type
     */
    public /*out*/ readonly subType!: pulumi.Output<string>;

    /**
     * Create a FailoverIpAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FailoverIpAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FailoverIpAttachArgs | FailoverIpAttachState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FailoverIpAttachState | undefined;
            resourceInputs["block"] = state ? state.block : undefined;
            resourceInputs["continentCode"] = state ? state.continentCode : undefined;
            resourceInputs["geoLoc"] = state ? state.geoLoc : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["progress"] = state ? state.progress : undefined;
            resourceInputs["routedTo"] = state ? state.routedTo : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subType"] = state ? state.subType : undefined;
        } else {
            const args = argsOrState as FailoverIpAttachArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["block"] = args ? args.block : undefined;
            resourceInputs["continentCode"] = args ? args.continentCode : undefined;
            resourceInputs["geoLoc"] = args ? args.geoLoc : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["routedTo"] = args ? args.routedTo : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["progress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FailoverIpAttach.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FailoverIpAttach resources.
 */
export interface FailoverIpAttachState {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    block?: pulumi.Input<string>;
    /**
     * Ip continent
     */
    continentCode?: pulumi.Input<string>;
    /**
     * Ip location
     */
    geoLoc?: pulumi.Input<string>;
    /**
     * The failover ip address to attach
     */
    ip?: pulumi.Input<string>;
    /**
     * Current operation progress in percent
     * * `routedTo` - Instance where ip is routed to
     */
    progress?: pulumi.Input<number>;
    /**
     * The GUID of an instance to which the failover IP address is be attached
     */
    routedTo?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Ip status, can be `ok` or `operationPending`
     * * `subType` - IP sub type, can be `cloud` or `ovh`
     */
    status?: pulumi.Input<string>;
    /**
     * IP sub type
     */
    subType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FailoverIpAttach resource.
 */
export interface FailoverIpAttachArgs {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     */
    block?: pulumi.Input<string>;
    /**
     * Ip continent
     */
    continentCode?: pulumi.Input<string>;
    /**
     * Ip location
     */
    geoLoc?: pulumi.Input<string>;
    /**
     * The failover ip address to attach
     */
    ip?: pulumi.Input<string>;
    /**
     * The GUID of an instance to which the failover IP address is be attached
     */
    routedTo?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
