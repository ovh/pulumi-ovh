// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing partitions on HA-NAS services
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myPartition = new ovh.dedicated.NasHAPartition("my_partition", {
 *     serviceName: "zpool-12345",
 *     name: "my-partition",
 *     size: 20,
 *     protocol: "NFS",
 * });
 * ```
 *
 * ## Import
 *
 * HA-NAS can be imported using the `{service_name}/{name}`, e.g.
 *
 * ```sh
 * $ pulumi import ovh:Dedicated/nasHAPartition:NasHAPartition my-partition zpool-12345/my-partition`
 * ```
 */
export class NasHAPartition extends pulumi.CustomResource {
    /**
     * Get an existing NasHAPartition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasHAPartitionState, opts?: pulumi.CustomResourceOptions): NasHAPartition {
        return new NasHAPartition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/nasHAPartition:NasHAPartition';

    /**
     * Returns true if the given object is an instance of NasHAPartition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasHAPartition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasHAPartition.__pulumiType;
    }

    /**
     * Percentage of partition space used in %
     */
    public /*out*/ readonly capacity!: pulumi.Output<number>;
    /**
     * A brief description of the partition
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * name of the partition
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * one of "NFS", "CIFS" or "NFS_CIFS"
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * size of the partition in GB
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Percentage of partition space used by snapshots in %
     */
    public /*out*/ readonly usedBySnapshots!: pulumi.Output<number>;

    /**
     * Create a NasHAPartition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasHAPartitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasHAPartitionArgs | NasHAPartitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasHAPartitionState | undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["usedBySnapshots"] = state ? state.usedBySnapshots : undefined;
        } else {
            const args = argsOrState as NasHAPartitionArgs | undefined;
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["capacity"] = undefined /*out*/;
            resourceInputs["usedBySnapshots"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NasHAPartition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasHAPartition resources.
 */
export interface NasHAPartitionState {
    /**
     * Percentage of partition space used in %
     */
    capacity?: pulumi.Input<number>;
    /**
     * A brief description of the partition
     */
    description?: pulumi.Input<string>;
    /**
     * name of the partition
     */
    name?: pulumi.Input<string>;
    /**
     * one of "NFS", "CIFS" or "NFS_CIFS"
     */
    protocol?: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName?: pulumi.Input<string>;
    /**
     * size of the partition in GB
     */
    size?: pulumi.Input<number>;
    /**
     * Percentage of partition space used by snapshots in %
     */
    usedBySnapshots?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NasHAPartition resource.
 */
export interface NasHAPartitionArgs {
    /**
     * A brief description of the partition
     */
    description?: pulumi.Input<string>;
    /**
     * name of the partition
     */
    name?: pulumi.Input<string>;
    /**
     * one of "NFS", "CIFS" or "NFS_CIFS"
     */
    protocol: pulumi.Input<string>;
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     */
    serviceName: pulumi.Input<string>;
    /**
     * size of the partition in GB
     */
    size: pulumi.Input<number>;
}
