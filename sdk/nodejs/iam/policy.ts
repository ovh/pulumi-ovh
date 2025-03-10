// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Iam/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    public readonly allows!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly denies!: pulumi.Output<string[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly excepts!: pulumi.Output<string[] | undefined>;
    public readonly identities!: pulumi.Output<string[]>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly owner!: pulumi.Output<string>;
    public readonly permissionsGroups!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly readOnly!: pulumi.Output<boolean>;
    public readonly resources!: pulumi.Output<string[]>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["allows"] = state ? state.allows : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["denies"] = state ? state.denies : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excepts"] = state ? state.excepts : undefined;
            resourceInputs["identities"] = state ? state.identities : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["permissionsGroups"] = state ? state.permissionsGroups : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.identities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identities'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            resourceInputs["allows"] = args ? args.allows : undefined;
            resourceInputs["denies"] = args ? args.denies : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excepts"] = args ? args.excepts : undefined;
            resourceInputs["identities"] = args ? args.identities : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissionsGroups"] = args ? args.permissionsGroups : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["readOnly"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    allows?: pulumi.Input<pulumi.Input<string>[]>;
    createdAt?: pulumi.Input<string>;
    denies?: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    excepts?: pulumi.Input<pulumi.Input<string>[]>;
    identities?: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    owner?: pulumi.Input<string>;
    permissionsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readOnly?: pulumi.Input<boolean>;
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    allows?: pulumi.Input<pulumi.Input<string>[]>;
    denies?: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    excepts?: pulumi.Input<pulumi.Input<string>[]>;
    identities: pulumi.Input<pulumi.Input<string>[]>;
    name?: pulumi.Input<string>;
    permissionsGroups?: pulumi.Input<pulumi.Input<string>[]>;
    resources: pulumi.Input<pulumi.Input<string>[]>;
}
