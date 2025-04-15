// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Moves a given IP to a different service, or inversely, parks it if empty service is given
 *
 * ## Move IP `1.2.3.4` to service loadbalancer-XXXXX
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const moveIpToLoadBalancerXxxxx = new ovh.ip.Move("move_ip_to_load_balancer_xxxxx", {
 *     ip: "1.2.3.4",
 *     routedTo: {
 *         serviceName: "loadbalancer-XXXXX",
 *     },
 * });
 * ```
 *
 * ## Park IP/Detach IP `1.2.3.4` from any service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const parkIp = new ovh.ip.Move("park_ip", {
 *     ip: "1.2.3.4",
 *     routedTo: {
 *         serviceName: "",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the `ip` field, e.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:Ip/move:Move mv '1.2.3.4/32'
 * ```
 */
export class Move extends pulumi.CustomResource {
    /**
     * Get an existing Move resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MoveState, opts?: pulumi.CustomResourceOptions): Move {
        return new Move(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Ip/move:Move';

    /**
     * Returns true if the given object is an instance of Move.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Move {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Move.__pulumiType;
    }

    /**
     * Whether IP service can be terminated
     */
    public /*out*/ readonly canBeTerminated!: pulumi.Output<boolean>;
    /**
     * Country
     */
    public /*out*/ readonly country!: pulumi.Output<string>;
    /**
     * Description attached to the IP
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * IP block that we want to attach to a different service
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * IP block organisation Id
     */
    public /*out*/ readonly organisationId!: pulumi.Output<string>;
    /**
     * Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST) instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
     */
    public readonly routedTo!: pulumi.Output<outputs.Ip.MoveRoutedTo>;
    /**
     * Service name in the form of `ip-<part-1>.<part-2>.<part-3>.<part-4>`
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
     */
    public /*out*/ readonly taskStartDate!: pulumi.Output<string>;
    /**
     * Status field of the current IP task that is in charge of changing the service the IP is attached to
     */
    public /*out*/ readonly taskStatus!: pulumi.Output<string>;
    /**
     * Possible values for ip type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Move resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MoveArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MoveArgs | MoveState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MoveState | undefined;
            resourceInputs["canBeTerminated"] = state ? state.canBeTerminated : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["organisationId"] = state ? state.organisationId : undefined;
            resourceInputs["routedTo"] = state ? state.routedTo : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["taskStartDate"] = state ? state.taskStartDate : undefined;
            resourceInputs["taskStatus"] = state ? state.taskStatus : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MoveArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.routedTo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routedTo'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["routedTo"] = args ? args.routedTo : undefined;
            resourceInputs["canBeTerminated"] = undefined /*out*/;
            resourceInputs["country"] = undefined /*out*/;
            resourceInputs["organisationId"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["taskStartDate"] = undefined /*out*/;
            resourceInputs["taskStatus"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Move.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Move resources.
 */
export interface MoveState {
    /**
     * Whether IP service can be terminated
     */
    canBeTerminated?: pulumi.Input<boolean>;
    /**
     * Country
     */
    country?: pulumi.Input<string>;
    /**
     * Description attached to the IP
     */
    description?: pulumi.Input<string>;
    /**
     * IP block that we want to attach to a different service
     */
    ip?: pulumi.Input<string>;
    /**
     * IP block organisation Id
     */
    organisationId?: pulumi.Input<string>;
    /**
     * Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST) instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
     */
    routedTo?: pulumi.Input<inputs.Ip.MoveRoutedTo>;
    /**
     * Service name in the form of `ip-<part-1>.<part-2>.<part-3>.<part-4>`
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
     */
    taskStartDate?: pulumi.Input<string>;
    /**
     * Status field of the current IP task that is in charge of changing the service the IP is attached to
     */
    taskStatus?: pulumi.Input<string>;
    /**
     * Possible values for ip type
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Move resource.
 */
export interface MoveArgs {
    /**
     * Description attached to the IP
     */
    description?: pulumi.Input<string>;
    /**
     * IP block that we want to attach to a different service
     */
    ip: pulumi.Input<string>;
    /**
     * Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST) instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
     */
    routedTo: pulumi.Input<inputs.Ip.MoveRoutedTo>;
}
