// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Deprecated: Use ipRestriction field in cloudProjectDatabase resource instead. Continuing to use the ovh.CloudProjectDatabase.IpRestriction resource to add an IP restriction to a cloudProjectDatabase resource will cause the cloudProjectDatabase resource to be updated on every apply
 *
 * Apply IP restrictions to an OVHcloud Managed Database cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const db = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXXX",
 *     engine: "YYYY",
 *     id: "ZZZZ",
 * });
 * const ipRestriction = new ovh.cloudprojectdatabase.IpRestriction("ip_restriction", {
 *     serviceName: db.then(db => db.serviceName),
 *     engine: db.then(db => db.engine),
 *     clusterId: db.then(db => db.id),
 *     ip: "178.97.6.0/24",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/ipRestriction:IpRestriction my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24
 * ```
 */
export class IpRestriction extends pulumi.CustomResource {
    /**
     * Get an existing IpRestriction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpRestrictionState, opts?: pulumi.CustomResourceOptions): IpRestriction {
        return new IpRestriction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/ipRestriction:IpRestriction';

    /**
     * Returns true if the given object is an instance of IpRestriction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpRestriction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpRestriction.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Description of the IP restriction.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Authorized IP.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current status of the IP restriction.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a IpRestriction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpRestrictionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpRestrictionArgs | IpRestrictionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpRestrictionState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as IpRestrictionArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpRestriction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpRestriction resources.
 */
export interface IpRestrictionState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Description of the IP restriction.
     */
    description?: pulumi.Input<string>;
    /**
     * The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine?: pulumi.Input<string>;
    /**
     * Authorized IP.
     */
    ip?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current status of the IP restriction.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpRestriction resource.
 */
export interface IpRestrictionArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Description of the IP restriction.
     */
    description?: pulumi.Input<string>;
    /**
     * The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: pulumi.Input<string>;
    /**
     * Authorized IP.
     */
    ip: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}
