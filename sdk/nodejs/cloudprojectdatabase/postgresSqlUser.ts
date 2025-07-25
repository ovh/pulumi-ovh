// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * OVHcloud Managed PostgreSQL clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser my_user service_name/cluster_id/id
 * ```
 */
export class PostgresSqlUser extends pulumi.CustomResource {
    /**
     * Get an existing PostgresSqlUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PostgresSqlUserState, opts?: pulumi.CustomResourceOptions): PostgresSqlUser {
        return new PostgresSqlUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser';

    /**
     * Returns true if the given object is an instance of PostgresSqlUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PostgresSqlUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PostgresSqlUser.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Date of the creation of the user.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the user. A user named "avnadmin" is mapped with already created admin user and reset his password instead of creating a new user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Sensitive) Password of the user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    public readonly passwordReset!: pulumi.Output<string | undefined>;
    /**
     * Roles the user belongs to. Available roles:
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current status of the user.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a PostgresSqlUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PostgresSqlUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PostgresSqlUserArgs | PostgresSqlUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PostgresSqlUserState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordReset"] = state ? state.passwordReset : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as PostgresSqlUserArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passwordReset"] = args ? args.passwordReset : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(PostgresSqlUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PostgresSqlUser resources.
 */
export interface PostgresSqlUserState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Date of the creation of the user.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the user. A user named "avnadmin" is mapped with already created admin user and reset his password instead of creating a new user.
     */
    name?: pulumi.Input<string>;
    /**
     * (Sensitive) Password of the user.
     */
    password?: pulumi.Input<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    passwordReset?: pulumi.Input<string>;
    /**
     * Roles the user belongs to. Available roles:
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current status of the user.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PostgresSqlUser resource.
 */
export interface PostgresSqlUserArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the user. A user named "avnadmin" is mapped with already created admin user and reset his password instead of creating a new user.
     */
    name?: pulumi.Input<string>;
    /**
     * Arbitrary string to change to trigger a password update
     */
    passwordReset?: pulumi.Input<string>;
    /**
     * Roles the user belongs to. Available roles:
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}
