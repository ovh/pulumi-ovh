// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a share for an EFS service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const efs = ovh.getStorageEfs({
 *     serviceName: "XXX",
 * });
 * const share = new ovh.StorageEfsShare("share", {
 *     serviceName: efs.then(efs => efs.serviceName),
 *     name: "share",
 *     description: "My share",
 *     protocol: "NFS",
 *     size: 100,
 * });
 * ```
 */
export class StorageEfsShare extends pulumi.CustomResource {
    /**
     * Get an existing StorageEfsShare resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageEfsShareState, opts?: pulumi.CustomResourceOptions): StorageEfsShare {
        return new StorageEfsShare(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/storageEfsShare:StorageEfsShare';

    /**
     * Returns true if the given object is an instance of StorageEfsShare.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageEfsShare {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageEfsShare.__pulumiType;
    }

    /**
     * Share creation date
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Share description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User-defined name used to generate human readable access path for the share
     */
    public readonly mountPointName!: pulumi.Output<string>;
    /**
     * Share name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Share protocol
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Service name
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Share size in Gigabytes
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Snapshot ID used to create the share
     */
    public readonly snapshotId!: pulumi.Output<string>;
    /**
     * Share status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a StorageEfsShare resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageEfsShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageEfsShareArgs | StorageEfsShareState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StorageEfsShareState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["mountPointName"] = state ? state.mountPointName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as StorageEfsShareArgs | undefined;
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
            resourceInputs["mountPointName"] = args ? args.mountPointName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StorageEfsShare.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageEfsShare resources.
 */
export interface StorageEfsShareState {
    /**
     * Share creation date
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Share description
     */
    description?: pulumi.Input<string>;
    /**
     * User-defined name used to generate human readable access path for the share
     */
    mountPointName?: pulumi.Input<string>;
    /**
     * Share name
     */
    name?: pulumi.Input<string>;
    /**
     * Share protocol
     */
    protocol?: pulumi.Input<string>;
    /**
     * Service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Share size in Gigabytes
     */
    size?: pulumi.Input<number>;
    /**
     * Snapshot ID used to create the share
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Share status
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StorageEfsShare resource.
 */
export interface StorageEfsShareArgs {
    /**
     * Share description
     */
    description?: pulumi.Input<string>;
    /**
     * User-defined name used to generate human readable access path for the share
     */
    mountPointName?: pulumi.Input<string>;
    /**
     * Share name
     */
    name?: pulumi.Input<string>;
    /**
     * Share protocol
     */
    protocol: pulumi.Input<string>;
    /**
     * Service name
     */
    serviceName: pulumi.Input<string>;
    /**
     * Share size in Gigabytes
     */
    size: pulumi.Input<number>;
    /**
     * Snapshot ID used to create the share
     */
    snapshotId?: pulumi.Input<string>;
}
