// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a backend server group (frontend) to be used by loadbalancing frontend(s)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farm80 = new ovh.iploadbalancing.TcpFarm("farm80", {
 *     displayName: "ingress-8080-gra",
 *     port: 80,
 *     serviceName: lb.then(lb => lb.serviceName),
 *     zone: "all",
 * });
 * const testFrontend = new ovh.iploadbalancing.TcpFrontend("testFrontend", {
 *     defaultFarmId: farm80.id,
 *     displayName: "ingress-8080-gra",
 *     port: "80,443",
 *     serviceName: lb.then(lb => lb.serviceName),
 *     zone: "all",
 * });
 * ```
 *
 * ## Import
 *
 * TCP frontend can be imported using the following format `serviceName` and the `id` of the frontend separated by "/" e.g.
 */
export class TcpFrontend extends pulumi.CustomResource {
    /**
     * Get an existing TcpFrontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TcpFrontendState, opts?: pulumi.CustomResourceOptions): TcpFrontend {
        return new TcpFrontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/tcpFrontend:TcpFrontend';

    /**
     * Returns true if the given object is an instance of TcpFrontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TcpFrontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TcpFrontend.__pulumiType;
    }

    /**
     * Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     */
    public readonly allowedSources!: pulumi.Output<string[] | undefined>;
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    public readonly dedicatedIpfos!: pulumi.Output<string[] | undefined>;
    /**
     * Default TCP Farm of your frontend
     */
    public readonly defaultFarmId!: pulumi.Output<number>;
    /**
     * Default ssl served to your customer
     */
    public readonly defaultSslId!: pulumi.Output<number>;
    /**
     * Deny IP Load Balancing access to these ip block. No restriction if null. You cannot specify both `allowedSource` and `deniedSource` at the same time. List of IP blocks.
     */
    public readonly deniedSources!: pulumi.Output<string[] | undefined>;
    /**
     * Disable your frontend. Default: 'false'
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * SSL deciphering. Default: 'false'
     */
    public readonly ssl!: pulumi.Output<boolean | undefined>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a TcpFrontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TcpFrontendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TcpFrontendArgs | TcpFrontendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TcpFrontendState | undefined;
            resourceInputs["allowedSources"] = state ? state.allowedSources : undefined;
            resourceInputs["dedicatedIpfos"] = state ? state.dedicatedIpfos : undefined;
            resourceInputs["defaultFarmId"] = state ? state.defaultFarmId : undefined;
            resourceInputs["defaultSslId"] = state ? state.defaultSslId : undefined;
            resourceInputs["deniedSources"] = state ? state.deniedSources : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as TcpFrontendArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["allowedSources"] = args ? args.allowedSources : undefined;
            resourceInputs["dedicatedIpfos"] = args ? args.dedicatedIpfos : undefined;
            resourceInputs["defaultFarmId"] = args ? args.defaultFarmId : undefined;
            resourceInputs["defaultSslId"] = args ? args.defaultSslId : undefined;
            resourceInputs["deniedSources"] = args ? args.deniedSources : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TcpFrontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TcpFrontend resources.
 */
export interface TcpFrontendState {
    /**
     * Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     */
    allowedSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default TCP Farm of your frontend
     */
    defaultFarmId?: pulumi.Input<number>;
    /**
     * Default ssl served to your customer
     */
    defaultSslId?: pulumi.Input<number>;
    /**
     * Deny IP Load Balancing access to these ip block. No restriction if null. You cannot specify both `allowedSource` and `deniedSource` at the same time. List of IP blocks.
     */
    deniedSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    port?: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * SSL deciphering. Default: 'false'
     */
    ssl?: pulumi.Input<boolean>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TcpFrontend resource.
 */
export interface TcpFrontendArgs {
    /**
     * Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     */
    allowedSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default TCP Farm of your frontend
     */
    defaultFarmId?: pulumi.Input<number>;
    /**
     * Default ssl served to your customer
     */
    defaultSslId?: pulumi.Input<number>;
    /**
     * Deny IP Load Balancing access to these ip block. No restriction if null. You cannot specify both `allowedSource` and `deniedSource` at the same time. List of IP blocks.
     */
    deniedSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    port: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * SSL deciphering. Default: 'false'
     */
    ssl?: pulumi.Input<boolean>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    zone: pulumi.Input<string>;
}
