// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a user in a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const user1 = new ovh.cloudproject.User("user1", {serviceName: "XXX"});
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * the date the user was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * A description associated with the user.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * a convenient map representing an openstackRc file.
     * Note: no password nor sensitive token is set in this map.
     */
    public readonly openstackRc!: pulumi.Output<{[key: string]: string}>;
    /**
     * (Sensitive) the password generated for the user. The password can
     * be used with the Openstack API. This attribute is sensitive and will only be
     * retrieve once during creation.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * The name of a role. See `roleNames`.
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * A list of role names. Values can be: 
     * - administrator,
     * - aiTrainingOperator
     * - aiTrainingRead
     * - authentication
     * - backupOperator
     * - computeOperator
     * - imageOperator
     * - infrastructureSupervisor
     * - networkOperator
     * - networkSecurityOperator
     * - objectstoreOperator
     * - volume_operator
     */
    public readonly roleNames!: pulumi.Output<string[] | undefined>;
    /**
     * A list of roles associated with the user.
     */
    public /*out*/ readonly roles!: pulumi.Output<outputs.CloudProject.UserRole[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * the status of the user. should be normally set to 'ok'.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * the username generated for the user. This username can be used with
     * the Openstack API.
     */
    public /*out*/ readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["openstackRc"] = state ? state.openstackRc : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["roleNames"] = state ? state.roleNames : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["openstackRc"] = args ? args.openstackRc : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["roleNames"] = args ? args.roleNames : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["roles"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["username"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * the date the user was created.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * A description associated with the user.
     */
    description?: pulumi.Input<string>;
    /**
     * a convenient map representing an openstackRc file.
     * Note: no password nor sensitive token is set in this map.
     */
    openstackRc?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * (Sensitive) the password generated for the user. The password can
     * be used with the Openstack API. This attribute is sensitive and will only be
     * retrieve once during creation.
     */
    password?: pulumi.Input<string>;
    /**
     * The name of a role. See `roleNames`.
     */
    roleName?: pulumi.Input<string>;
    /**
     * A list of role names. Values can be: 
     * - administrator,
     * - aiTrainingOperator
     * - aiTrainingRead
     * - authentication
     * - backupOperator
     * - computeOperator
     * - imageOperator
     * - infrastructureSupervisor
     * - networkOperator
     * - networkSecurityOperator
     * - objectstoreOperator
     * - volume_operator
     */
    roleNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of roles associated with the user.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.CloudProject.UserRole>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * the status of the user. should be normally set to 'ok'.
     */
    status?: pulumi.Input<string>;
    /**
     * the username generated for the user. This username can be used with
     * the Openstack API.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * A description associated with the user.
     */
    description?: pulumi.Input<string>;
    /**
     * a convenient map representing an openstackRc file.
     * Note: no password nor sensitive token is set in this map.
     */
    openstackRc?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of a role. See `roleNames`.
     */
    roleName?: pulumi.Input<string>;
    /**
     * A list of role names. Values can be: 
     * - administrator,
     * - aiTrainingOperator
     * - aiTrainingRead
     * - authentication
     * - backupOperator
     * - computeOperator
     * - imageOperator
     * - infrastructureSupervisor
     * - networkOperator
     * - networkSecurityOperator
     * - objectstoreOperator
     * - volume_operator
     */
    roleNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
