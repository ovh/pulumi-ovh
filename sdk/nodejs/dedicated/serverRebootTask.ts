// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const rescue = ovh.Dedicated.getServerBoots({
 *     serviceName: "nsxxxxxxx.ip-xx-xx-xx.eu",
 *     bootType: "rescue",
 *     kernel: "rescue64-pro",
 * });
 * const serverOnRescue = new ovh.dedicated.ServerUpdate("server_on_rescue", {
 *     serviceName: "nsxxxxxxx.ip-xx-xx-xx.eu",
 *     bootId: rescue.then(rescue => rescue.results?.[0]),
 *     monitoring: true,
 *     state: "ok",
 * });
 * const serverReboot = new ovh.dedicated.ServerRebootTask("server_reboot", {
 *     serviceName: rescue.then(rescue => rescue.serviceName),
 *     keepers: [serverOnRescue.bootId],
 * });
 * ```
 */
export class ServerRebootTask extends pulumi.CustomResource {
    /**
     * Get an existing ServerRebootTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerRebootTaskState, opts?: pulumi.CustomResourceOptions): ServerRebootTask {
        return new ServerRebootTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/serverRebootTask:ServerRebootTask';

    /**
     * Returns true if the given object is an instance of ServerRebootTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerRebootTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerRebootTask.__pulumiType;
    }

    /**
     * Details of this task. (should be `Reboot asked`)
     */
    public /*out*/ readonly comment!: pulumi.Output<string>;
    /**
     * Completion date in RFC3339 format.
     */
    public /*out*/ readonly doneDate!: pulumi.Output<string>;
    /**
     * Function name (should be `hardReboot`).
     */
    public /*out*/ readonly function!: pulumi.Output<string>;
    /**
     * List of values tracked to trigger reboot, used also to form implicit dependencies.
     */
    public readonly keepers!: pulumi.Output<string[]>;
    /**
     * Last update in RFC3339 format.
     */
    public /*out*/ readonly lastUpdate!: pulumi.Output<string>;
    /**
     * The serviceName of your dedicated server.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Task creation date in RFC3339 format.
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * Task status (should be `done`)
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ServerRebootTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerRebootTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerRebootTaskArgs | ServerRebootTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerRebootTaskState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["doneDate"] = state ? state.doneDate : undefined;
            resourceInputs["function"] = state ? state.function : undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["lastUpdate"] = state ? state.lastUpdate : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ServerRebootTaskArgs | undefined;
            if ((!args || args.keepers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keepers'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["comment"] = undefined /*out*/;
            resourceInputs["doneDate"] = undefined /*out*/;
            resourceInputs["function"] = undefined /*out*/;
            resourceInputs["lastUpdate"] = undefined /*out*/;
            resourceInputs["startDate"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerRebootTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerRebootTask resources.
 */
export interface ServerRebootTaskState {
    /**
     * Details of this task. (should be `Reboot asked`)
     */
    comment?: pulumi.Input<string>;
    /**
     * Completion date in RFC3339 format.
     */
    doneDate?: pulumi.Input<string>;
    /**
     * Function name (should be `hardReboot`).
     */
    function?: pulumi.Input<string>;
    /**
     * List of values tracked to trigger reboot, used also to form implicit dependencies.
     */
    keepers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Last update in RFC3339 format.
     */
    lastUpdate?: pulumi.Input<string>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Task creation date in RFC3339 format.
     */
    startDate?: pulumi.Input<string>;
    /**
     * Task status (should be `done`)
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerRebootTask resource.
 */
export interface ServerRebootTaskArgs {
    /**
     * List of values tracked to trigger reboot, used also to form implicit dependencies.
     */
    keepers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
}
