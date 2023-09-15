// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a HTTP backend server group (farm) to be used by loadbalancing frontend(s)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovh-devrelteam/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farmname = new ovh.iploadbalancing.HttpFarm("farmname", {
 *     displayName: "ingress-8080-gra",
 *     serviceName: lb.then(lb => lb.serviceName),
 *     zone: "GRA",
 * });
 * ```
 */
export class HttpFarm extends pulumi.CustomResource {
    /**
     * Get an existing HttpFarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpFarmState, opts?: pulumi.CustomResourceOptions): HttpFarm {
        return new HttpFarm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/httpFarm:HttpFarm';

    /**
     * Returns true if the given object is an instance of HttpFarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HttpFarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HttpFarm.__pulumiType;
    }

    /**
     * Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     */
    public readonly balance!: pulumi.Output<string | undefined>;
    /**
     * Readable label for loadbalancer farm
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Port for backends to receive traffic on.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * define a backend healthcheck probe
     */
    public readonly probe!: pulumi.Output<outputs.IpLoadBalancing.HttpFarmProbe | undefined>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
     */
    public readonly stickiness!: pulumi.Output<string | undefined>;
    /**
     * Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     */
    public readonly vrackNetworkId!: pulumi.Output<number | undefined>;
    /**
     * Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a HttpFarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HttpFarmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpFarmArgs | HttpFarmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HttpFarmState | undefined;
            resourceInputs["balance"] = state ? state.balance : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["probe"] = state ? state.probe : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["stickiness"] = state ? state.stickiness : undefined;
            resourceInputs["vrackNetworkId"] = state ? state.vrackNetworkId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as HttpFarmArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["balance"] = args ? args.balance : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["probe"] = args ? args.probe : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["stickiness"] = args ? args.stickiness : undefined;
            resourceInputs["vrackNetworkId"] = args ? args.vrackNetworkId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HttpFarm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpFarm resources.
 */
export interface HttpFarmState {
    /**
     * Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     */
    balance?: pulumi.Input<string>;
    /**
     * Readable label for loadbalancer farm
     */
    displayName?: pulumi.Input<string>;
    /**
     * Port for backends to receive traffic on.
     */
    port?: pulumi.Input<number>;
    /**
     * define a backend healthcheck probe
     */
    probe?: pulumi.Input<inputs.IpLoadBalancing.HttpFarmProbe>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
     */
    stickiness?: pulumi.Input<string>;
    /**
     * Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     */
    vrackNetworkId?: pulumi.Input<number>;
    /**
     * Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HttpFarm resource.
 */
export interface HttpFarmArgs {
    /**
     * Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     */
    balance?: pulumi.Input<string>;
    /**
     * Readable label for loadbalancer farm
     */
    displayName?: pulumi.Input<string>;
    /**
     * Port for backends to receive traffic on.
     */
    port?: pulumi.Input<number>;
    /**
     * define a backend healthcheck probe
     */
    probe?: pulumi.Input<inputs.IpLoadBalancing.HttpFarmProbe>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
     */
    stickiness?: pulumi.Input<string>;
    /**
     * Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     */
    vrackNetworkId?: pulumi.Input<number>;
    /**
     * Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     */
    zone: pulumi.Input<string>;
}
