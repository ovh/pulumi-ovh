// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows to manipulate LDP tokens.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const token = new ovh.dbaas.LogsToken("token", {
 *     serviceName: "ldp-xx-xxxxx",
 *     name: "ExampleToken",
 * });
 * ```
 */
export class LogsToken extends pulumi.CustomResource {
    /**
     * Get an existing LogsToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogsTokenState, opts?: pulumi.CustomResourceOptions): LogsToken {
        return new LogsToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dbaas/logsToken:LogsToken';

    /**
     * Returns true if the given object is an instance of LogsToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogsToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogsToken.__pulumiType;
    }

    /**
     * Cluster ID. If not provided, the default clusterId is used
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Token creation date
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the token
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The LDP service name
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * ID of the token
     */
    public /*out*/ readonly tokenId!: pulumi.Output<string>;
    /**
     * Token last update date
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Token value
     */
    public /*out*/ readonly value!: pulumi.Output<string>;

    /**
     * Create a LogsToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogsTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogsTokenArgs | LogsTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogsTokenState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["tokenId"] = state ? state.tokenId : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as LogsTokenArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["tokenId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LogsToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogsToken resources.
 */
export interface LogsTokenState {
    /**
     * Cluster ID. If not provided, the default clusterId is used
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Token creation date
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the token
     */
    name?: pulumi.Input<string>;
    /**
     * The LDP service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * ID of the token
     */
    tokenId?: pulumi.Input<string>;
    /**
     * Token last update date
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Token value
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogsToken resource.
 */
export interface LogsTokenArgs {
    /**
     * Cluster ID. If not provided, the default clusterId is used
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Name of the token
     */
    name?: pulumi.Input<string>;
    /**
     * The LDP service name
     */
    serviceName: pulumi.Input<string>;
}
