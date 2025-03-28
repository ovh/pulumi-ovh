// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Add grant on a database in your private cloud database instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const userGrant = new ovh.hosting.PrivateDatabaseUserGrant("userGrant", {
 *     databaseName: "ovhcloud",
 *     grant: "admin",
 *     serviceName: "XXXXXX",
 *     userName: "terraform",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud database user's grant can be imported using the `service_name`, the `user_name`, the `database_name` and the `grant`, separated by "/" E.g.,
 *
 * ```sh
 * $ pulumi import ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant user service_name/user_name/database_name/grant
 * ```
 */
export class PrivateDatabaseUserGrant extends pulumi.CustomResource {
    /**
     * Get an existing PrivateDatabaseUserGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateDatabaseUserGrantState, opts?: pulumi.CustomResourceOptions): PrivateDatabaseUserGrant {
        return new PrivateDatabaseUserGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant';

    /**
     * Returns true if the given object is an instance of PrivateDatabaseUserGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateDatabaseUserGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateDatabaseUserGrant.__pulumiType;
    }

    /**
     * Database name where add grant.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * Database name where add grant. Values can be:
     * - admin
     * - none
     * - ro
     * - rw
     */
    public readonly grant!: pulumi.Output<string>;
    /**
     * The internal name of your private database.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * User name used to connect on your databases.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a PrivateDatabaseUserGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateDatabaseUserGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateDatabaseUserGrantArgs | PrivateDatabaseUserGrantState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateDatabaseUserGrantState | undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["grant"] = state ? state.grant : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as PrivateDatabaseUserGrantArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.grant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'grant'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["grant"] = args ? args.grant : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateDatabaseUserGrant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateDatabaseUserGrant resources.
 */
export interface PrivateDatabaseUserGrantState {
    /**
     * Database name where add grant.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * Database name where add grant. Values can be:
     * - admin
     * - none
     * - ro
     * - rw
     */
    grant?: pulumi.Input<string>;
    /**
     * The internal name of your private database.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * User name used to connect on your databases.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateDatabaseUserGrant resource.
 */
export interface PrivateDatabaseUserGrantArgs {
    /**
     * Database name where add grant.
     */
    databaseName: pulumi.Input<string>;
    /**
     * Database name where add grant. Values can be:
     * - admin
     * - none
     * - ro
     * - rw
     */
    grant: pulumi.Input<string>;
    /**
     * The internal name of your private database.
     */
    serviceName: pulumi.Input<string>;
    /**
     * User name used to connect on your databases.
     */
    userName: pulumi.Input<string>;
}
