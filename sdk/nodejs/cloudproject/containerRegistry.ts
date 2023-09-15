// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a container registry associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovh-devrelteam/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const regcap = ovh.CloudProject.getCapabilitiesContainerFilter({
 *     serviceName: "XXXXXX",
 *     planName: "SMALL",
 *     region: "GRA",
 * });
 * const my_registry = new ovh.cloudproject.ContainerRegistry("my-registry", {
 *     serviceName: regcap.then(regcap => regcap.serviceName),
 *     planId: regcap.then(regcap => regcap.id),
 *     region: regcap.then(regcap => regcap.region),
 * });
 * ```
 *
 * > __WARNING__ You can update and migrate to a higher plan at any time but not the contrary.
 */
export class ContainerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerRegistryState, opts?: pulumi.CustomResourceOptions): ContainerRegistry {
        return new ContainerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/containerRegistry:ContainerRegistry';

    /**
     * Returns true if the given object is an instance of ContainerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRegistry.__pulumiType;
    }

    /**
     * Plan creation date
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Registry name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Plan ID of the registry
     */
    public readonly planId!: pulumi.Output<string>;
    /**
     * Plan of the registry
     */
    public /*out*/ readonly plans!: pulumi.Output<outputs.CloudProject.ContainerRegistryPlan[]>;
    /**
     * Project ID of your registry
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * Region of the registry
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current size of the registry (bytes)
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Registry status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Registry last update date
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Access url of the registry
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Version of your registry
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a ContainerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerRegistryArgs | ContainerRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerRegistryState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["planId"] = state ? state.planId : undefined;
            resourceInputs["plans"] = state ? state.plans : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ContainerRegistryArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["planId"] = args ? args.planId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["plans"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerRegistry resources.
 */
export interface ContainerRegistryState {
    /**
     * Plan creation date
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Registry name
     */
    name?: pulumi.Input<string>;
    /**
     * Plan ID of the registry
     */
    planId?: pulumi.Input<string>;
    /**
     * Plan of the registry
     */
    plans?: pulumi.Input<pulumi.Input<inputs.CloudProject.ContainerRegistryPlan>[]>;
    /**
     * Project ID of your registry
     */
    projectId?: pulumi.Input<string>;
    /**
     * Region of the registry
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current size of the registry (bytes)
     */
    size?: pulumi.Input<number>;
    /**
     * Registry status
     */
    status?: pulumi.Input<string>;
    /**
     * Registry last update date
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Access url of the registry
     */
    url?: pulumi.Input<string>;
    /**
     * Version of your registry
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerRegistry resource.
 */
export interface ContainerRegistryArgs {
    /**
     * Registry name
     */
    name?: pulumi.Input<string>;
    /**
     * Plan ID of the registry
     */
    planId?: pulumi.Input<string>;
    /**
     * Region of the registry
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
