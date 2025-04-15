// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an identity group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myGroup = new ovh.me.IdentityGroup("my_group", {
 *     description: "Some custom description",
 *     name: "my_group_name",
 *     role: "NONE",
 * });
 * ```
 */
export class IdentityGroup extends pulumi.CustomResource {
    /**
     * Get an existing IdentityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityGroupState, opts?: pulumi.CustomResourceOptions): IdentityGroup {
        return new IdentityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Me/identityGroup:IdentityGroup';

    /**
     * Returns true if the given object is an instance of IdentityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityGroup.__pulumiType;
    }

    /**
     * URN of the user group, used when writing IAM policies
     */
    public /*out*/ readonly GroupURN!: pulumi.Output<string>;
    /**
     * Creation date of this group.
     */
    public /*out*/ readonly creation!: pulumi.Output<string>;
    /**
     * Is the group a default and immutable one.
     */
    public /*out*/ readonly defaultGroup!: pulumi.Output<boolean>;
    /**
     * Group description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Date of the last update of this group.
     */
    public /*out*/ readonly lastUpdate!: pulumi.Output<string>;
    /**
     * Group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
     */
    public readonly role!: pulumi.Output<string | undefined>;

    /**
     * Create a IdentityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IdentityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityGroupArgs | IdentityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityGroupState | undefined;
            resourceInputs["GroupURN"] = state ? state.GroupURN : undefined;
            resourceInputs["creation"] = state ? state.creation : undefined;
            resourceInputs["defaultGroup"] = state ? state.defaultGroup : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lastUpdate"] = state ? state.lastUpdate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IdentityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["GroupURN"] = undefined /*out*/;
            resourceInputs["creation"] = undefined /*out*/;
            resourceInputs["defaultGroup"] = undefined /*out*/;
            resourceInputs["lastUpdate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IdentityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityGroup resources.
 */
export interface IdentityGroupState {
    /**
     * URN of the user group, used when writing IAM policies
     */
    GroupURN?: pulumi.Input<string>;
    /**
     * Creation date of this group.
     */
    creation?: pulumi.Input<string>;
    /**
     * Is the group a default and immutable one.
     */
    defaultGroup?: pulumi.Input<boolean>;
    /**
     * Group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Date of the last update of this group.
     */
    lastUpdate?: pulumi.Input<string>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityGroup resource.
 */
export interface IdentityGroupArgs {
    /**
     * Group description.
     */
    description?: pulumi.Input<string>;
    /**
     * Group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
     */
    role?: pulumi.Input<string>;
}
