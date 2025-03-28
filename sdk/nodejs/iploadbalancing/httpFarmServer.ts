// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a backend server entry linked to HTTP loadbalancing group (farm)
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
 * const farmname = new ovh.iploadbalancing.HttpFarm("farmname", {
 *     port: 8080,
 *     serviceName: lb.then(lb => lb.serviceName),
 *     zone: "all",
 * });
 * const backend = new ovh.iploadbalancing.HttpFarmServer("backend", {
 *     address: "4.5.6.7",
 *     backup: true,
 *     displayName: "mybackend",
 *     farmId: farmname.id,
 *     port: 80,
 *     probe: true,
 *     proxyProtocolVersion: "v2",
 *     serviceName: lb.then(lb => lb.serviceName),
 *     ssl: false,
 *     status: "active",
 *     weight: 2,
 * });
 * ```
 *
 * ## Import
 *
 * HTTP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer backend service_name/farm_id/server_id
 * ```
 */
export class HttpFarmServer extends pulumi.CustomResource {
    /**
     * Get an existing HttpFarmServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpFarmServerState, opts?: pulumi.CustomResourceOptions): HttpFarmServer {
        return new HttpFarmServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer';

    /**
     * Returns true if the given object is an instance of HttpFarmServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HttpFarmServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HttpFarmServer.__pulumiType;
    }

    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * is it a backup server used in case of failure of all the non-backup backends
     */
    public readonly backup!: pulumi.Output<boolean | undefined>;
    public readonly chain!: pulumi.Output<string | undefined>;
    /**
     * Value of the stickiness cookie used for this backend.
     */
    public /*out*/ readonly cookie!: pulumi.Output<string>;
    /**
     * Label for the server
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * ID of the farm this server is attached to
     */
    public readonly farmId!: pulumi.Output<number>;
    /**
     * enable action when backend marked down. (`shutdown-sessions`)
     */
    public readonly onMarkedDown!: pulumi.Output<string | undefined>;
    /**
     * Port that backend will respond on
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * defines if backend will be probed to determine health and keep as active in farm if healthy
     */
    public readonly probe!: pulumi.Output<boolean | undefined>;
    /**
     * version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
     */
    public readonly proxyProtocolVersion!: pulumi.Output<string | undefined>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * is the connection ciphered with SSL (TLS)
     */
    public readonly ssl!: pulumi.Output<boolean | undefined>;
    /**
     * backend status - `active` or `inactive`
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * used in loadbalancing algorithm
     */
    public readonly weight!: pulumi.Output<number | undefined>;

    /**
     * Create a HttpFarmServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HttpFarmServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpFarmServerArgs | HttpFarmServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HttpFarmServerState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["backup"] = state ? state.backup : undefined;
            resourceInputs["chain"] = state ? state.chain : undefined;
            resourceInputs["cookie"] = state ? state.cookie : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["farmId"] = state ? state.farmId : undefined;
            resourceInputs["onMarkedDown"] = state ? state.onMarkedDown : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["probe"] = state ? state.probe : undefined;
            resourceInputs["proxyProtocolVersion"] = state ? state.proxyProtocolVersion : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as HttpFarmServerArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.farmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'farmId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["backup"] = args ? args.backup : undefined;
            resourceInputs["chain"] = args ? args.chain : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["farmId"] = args ? args.farmId : undefined;
            resourceInputs["onMarkedDown"] = args ? args.onMarkedDown : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["probe"] = args ? args.probe : undefined;
            resourceInputs["proxyProtocolVersion"] = args ? args.proxyProtocolVersion : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["cookie"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HttpFarmServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpFarmServer resources.
 */
export interface HttpFarmServerState {
    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    address?: pulumi.Input<string>;
    /**
     * is it a backup server used in case of failure of all the non-backup backends
     */
    backup?: pulumi.Input<boolean>;
    chain?: pulumi.Input<string>;
    /**
     * Value of the stickiness cookie used for this backend.
     */
    cookie?: pulumi.Input<string>;
    /**
     * Label for the server
     */
    displayName?: pulumi.Input<string>;
    /**
     * ID of the farm this server is attached to
     */
    farmId?: pulumi.Input<number>;
    /**
     * enable action when backend marked down. (`shutdown-sessions`)
     */
    onMarkedDown?: pulumi.Input<string>;
    /**
     * Port that backend will respond on
     */
    port?: pulumi.Input<number>;
    /**
     * defines if backend will be probed to determine health and keep as active in farm if healthy
     */
    probe?: pulumi.Input<boolean>;
    /**
     * version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
     */
    proxyProtocolVersion?: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * is the connection ciphered with SSL (TLS)
     */
    ssl?: pulumi.Input<boolean>;
    /**
     * backend status - `active` or `inactive`
     */
    status?: pulumi.Input<string>;
    /**
     * used in loadbalancing algorithm
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HttpFarmServer resource.
 */
export interface HttpFarmServerArgs {
    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    address: pulumi.Input<string>;
    /**
     * is it a backup server used in case of failure of all the non-backup backends
     */
    backup?: pulumi.Input<boolean>;
    chain?: pulumi.Input<string>;
    /**
     * Label for the server
     */
    displayName?: pulumi.Input<string>;
    /**
     * ID of the farm this server is attached to
     */
    farmId: pulumi.Input<number>;
    /**
     * enable action when backend marked down. (`shutdown-sessions`)
     */
    onMarkedDown?: pulumi.Input<string>;
    /**
     * Port that backend will respond on
     */
    port?: pulumi.Input<number>;
    /**
     * defines if backend will be probed to determine health and keep as active in farm if healthy
     */
    probe?: pulumi.Input<boolean>;
    /**
     * version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
     */
    proxyProtocolVersion?: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * is the connection ciphered with SSL (TLS)
     */
    ssl?: pulumi.Input<boolean>;
    /**
     * backend status - `active` or `inactive`
     */
    status: pulumi.Input<string>;
    /**
     * used in loadbalancing algorithm
     */
    weight?: pulumi.Input<number>;
}
