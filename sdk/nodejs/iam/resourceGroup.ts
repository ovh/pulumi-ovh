// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an OVHcloud IAM resource group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myResourceGroup = new ovh.iam.ResourceGroup("my_resource_group", {
 *     name: "my_resource_group",
 *     resources: [
 *         "urn:v1:eu:resource:service1:service1-id",
 *         "urn:v1:eu:resource:service2:service2-id",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Resource groups can be imported by using their id.
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:Iam/resourceGroup:ResourceGroup my_resource_group resource_group_id
 * ```
 *
 * -> Read only resource groups cannot be imported
 */
export class ResourceGroup extends pulumi.CustomResource {
    /**
     * Get an existing ResourceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceGroupState, opts?: pulumi.CustomResourceOptions): ResourceGroup {
        return new ResourceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Iam/resourceGroup:ResourceGroup';

    /**
     * Returns true if the given object is an instance of ResourceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceGroup.__pulumiType;
    }

    /**
     * URN of the resource group, used when writing policies
     */
    public /*out*/ readonly GroupURN!: pulumi.Output<string>;
    /**
     * Date of the creation of the resource group
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the resource group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the account owning the resource group
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
     */
    public /*out*/ readonly readOnly!: pulumi.Output<boolean>;
    /**
     * Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     */
    public readonly resources!: pulumi.Output<string[] | undefined>;
    /**
     * Date of the last modification of the resource group
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ResourceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceGroupArgs | ResourceGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceGroupState | undefined;
            resourceInputs["GroupURN"] = state ? state.GroupURN : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ResourceGroupArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["GroupURN"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["readOnly"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourceGroup resources.
 */
export interface ResourceGroupState {
    /**
     * URN of the resource group, used when writing policies
     */
    GroupURN?: pulumi.Input<string>;
    /**
     * Date of the creation of the resource group
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the resource group
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the account owning the resource group
     */
    owner?: pulumi.Input<string>;
    /**
     * Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date of the last modification of the resource group
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceGroup resource.
 */
export interface ResourceGroupArgs {
    /**
     * Name of the resource group
     */
    name?: pulumi.Input<string>;
    /**
     * Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
}
