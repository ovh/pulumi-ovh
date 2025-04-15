// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Reference a DBaaS logs role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const ro = new ovh.dbaas.LogsRole("ro", {
 *     serviceName: "ldp-xx-xxxxx",
 *     name: "Devops - RO",
 *     description: "Devops - RO",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud DBaaS Log Role can be imported using the `service_name` and `role_id` of the role, separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:Dbaas/logsRole:LogsRole  ovh_dbaas_logs_role.ro ldp-ra-XX/dc145bc2-eb01-4efe-a802-XXXXXX
 * ```
 */
export class LogsRole extends pulumi.CustomResource {
    /**
     * Get an existing LogsRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogsRoleState, opts?: pulumi.CustomResourceOptions): LogsRole {
        return new LogsRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dbaas/logsRole:LogsRole';

    /**
     * Returns true if the given object is an instance of LogsRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogsRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogsRole.__pulumiType;
    }

    /**
     * Role creation date
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The role description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The role name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * number of member for the role
     */
    public /*out*/ readonly nbMember!: pulumi.Output<number>;
    /**
     * number of configured permission for the role
     */
    public /*out*/ readonly nbPermission!: pulumi.Output<number>;
    /**
     * Role identifier
     */
    public /*out*/ readonly roleId!: pulumi.Output<string>;
    /**
     * The service name
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Role last update date
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a LogsRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogsRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogsRoleArgs | LogsRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogsRoleState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nbMember"] = state ? state.nbMember : undefined;
            resourceInputs["nbPermission"] = state ? state.nbPermission : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as LogsRoleArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["nbMember"] = undefined /*out*/;
            resourceInputs["nbPermission"] = undefined /*out*/;
            resourceInputs["roleId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogsRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogsRole resources.
 */
export interface LogsRoleState {
    /**
     * Role creation date
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The role description
     */
    description?: pulumi.Input<string>;
    /**
     * The role name
     */
    name?: pulumi.Input<string>;
    /**
     * number of member for the role
     */
    nbMember?: pulumi.Input<number>;
    /**
     * number of configured permission for the role
     */
    nbPermission?: pulumi.Input<number>;
    /**
     * Role identifier
     */
    roleId?: pulumi.Input<string>;
    /**
     * The service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Role last update date
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogsRole resource.
 */
export interface LogsRoleArgs {
    /**
     * The role description
     */
    description: pulumi.Input<string>;
    /**
     * The role name
     */
    name?: pulumi.Input<string>;
    /**
     * The service name
     */
    serviceName: pulumi.Input<string>;
}
