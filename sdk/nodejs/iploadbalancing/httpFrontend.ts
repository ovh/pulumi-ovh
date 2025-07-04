// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a backend HTTP server group (frontend) to be used by loadbalancing frontend(s)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farm80 = new ovh.iploadbalancing.HttpFarm("farm80", {
 *     serviceName: lb.then(lb => lb.serviceName),
 *     displayName: "ingress-8080-gra",
 *     zone: "all",
 *     port: 80,
 * });
 * const testFrontend = new ovh.iploadbalancing.HttpFrontend("test_frontend", {
 *     serviceName: lb.then(lb => lb.serviceName),
 *     displayName: "ingress-8080-gra",
 *     zone: "all",
 *     port: "80,443",
 *     defaultFarmId: farm80.id,
 * });
 * ```
 *
 * ### With HTTP Header
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farm80 = new ovh.iploadbalancing.HttpFarm("farm80", {
 *     serviceName: lb.then(lb => lb.serviceName),
 *     displayName: "ingress-8080-gra",
 *     zone: "all",
 *     port: 80,
 * });
 * const testFrontend = new ovh.iploadbalancing.HttpFrontend("test_frontend", {
 *     serviceName: lb.then(lb => lb.serviceName),
 *     displayName: "ingress-8080-gra",
 *     zone: "all",
 *     port: "80,443",
 *     defaultFarmId: farm80.id,
 *     httpHeaders: [
 *         "X-Ip-Header %%ci",
 *         "X-Port-Header %%cp",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * HTTP frontend can be imported using the following format `service_name` and the `id` of the frontend separated by "/" e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:IpLoadBalancing/httpFrontend:HttpFrontend testfrontend service_name/http_frontend_id
 * ```
 */
export class HttpFrontend extends pulumi.CustomResource {
    /**
     * Get an existing HttpFrontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpFrontendState, opts?: pulumi.CustomResourceOptions): HttpFrontend {
        return new HttpFrontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/httpFrontend:HttpFrontend';

    /**
     * Returns true if the given object is an instance of HttpFrontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HttpFrontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HttpFrontend.__pulumiType;
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
     * Disable your frontend. Default: 'false'
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * HTTP Strict Transport Security. Default: 'false'
     */
    public readonly hsts!: pulumi.Output<boolean | undefined>;
    /**
     * HTTP headers to add to the frontend. List of string.
     */
    public readonly httpHeaders!: pulumi.Output<string[] | undefined>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Redirection HTTP'
     */
    public readonly redirectLocation!: pulumi.Output<string | undefined>;
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
     * Create a HttpFrontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HttpFrontendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpFrontendArgs | HttpFrontendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HttpFrontendState | undefined;
            resourceInputs["allowedSources"] = state ? state.allowedSources : undefined;
            resourceInputs["dedicatedIpfos"] = state ? state.dedicatedIpfos : undefined;
            resourceInputs["defaultFarmId"] = state ? state.defaultFarmId : undefined;
            resourceInputs["defaultSslId"] = state ? state.defaultSslId : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["hsts"] = state ? state.hsts : undefined;
            resourceInputs["httpHeaders"] = state ? state.httpHeaders : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["redirectLocation"] = state ? state.redirectLocation : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as HttpFrontendArgs | undefined;
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
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["hsts"] = args ? args.hsts : undefined;
            resourceInputs["httpHeaders"] = args ? args.httpHeaders : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["redirectLocation"] = args ? args.redirectLocation : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HttpFrontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpFrontend resources.
 */
export interface HttpFrontendState {
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
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * HTTP Strict Transport Security. Default: 'false'
     */
    hsts?: pulumi.Input<boolean>;
    /**
     * HTTP headers to add to the frontend. List of string.
     */
    httpHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
     */
    port?: pulumi.Input<string>;
    /**
     * Redirection HTTP'
     */
    redirectLocation?: pulumi.Input<string>;
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
 * The set of arguments for constructing a HttpFrontend resource.
 */
export interface HttpFrontendArgs {
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
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * HTTP Strict Transport Security. Default: 'false'
     */
    hsts?: pulumi.Input<boolean>;
    /**
     * HTTP headers to add to the frontend. List of string.
     */
    httpHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range
     */
    port: pulumi.Input<string>;
    /**
     * Redirection HTTP'
     */
    redirectLocation?: pulumi.Input<string>;
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
