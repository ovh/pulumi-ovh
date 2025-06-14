// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * **This resource uses a Beta API** Creates an instance associated with a public cloud project.
 *
 * ## Example Usage
 *
 * Create a instance.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const instance = new ovh.cloudproject.Instance("instance", {
 *     serviceName: "XXX",
 *     region: "RRRR",
 *     billingPeriod: "hourly",
 *     bootFrom: {
 *         imageId: "UUID",
 *     },
 *     flavor: {
 *         flavorId: "UUID",
 *     },
 *     name: "instance name",
 *     sshKey: {
 *         name: "sshname",
 *     },
 *     network: {
 *         "public": true,
 *     },
 * });
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Instance IP addresses
     */
    public /*out*/ readonly addresses!: pulumi.Output<outputs.CloudProject.InstanceAddress[]>;
    /**
     * Volumes attached to the instance
     */
    public /*out*/ readonly attachedVolumes!: pulumi.Output<outputs.CloudProject.InstanceAttachedVolume[]>;
    /**
     * Create an autobackup workflow after instance start up.
     */
    public readonly autoBackup!: pulumi.Output<outputs.CloudProject.InstanceAutoBackup | undefined>;
    /**
     * The availability zone where the instance will be created
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Billing period - hourly or monthly
     */
    public readonly billingPeriod!: pulumi.Output<string>;
    /**
     * Boot the instance from an image or a volume
     */
    public readonly bootFrom!: pulumi.Output<outputs.CloudProject.InstanceBootFrom>;
    /**
     * Create multiple instances
     */
    public readonly bulk!: pulumi.Output<number | undefined>;
    /**
     * Flavor information
     */
    public readonly flavor!: pulumi.Output<outputs.CloudProject.InstanceFlavor>;
    /**
     * Flavor id
     */
    public /*out*/ readonly flavorId!: pulumi.Output<string>;
    /**
     * Flavor name
     */
    public /*out*/ readonly flavorName!: pulumi.Output<string>;
    /**
     * Start instance in group
     */
    public readonly group!: pulumi.Output<outputs.CloudProject.InstanceGroup | undefined>;
    /**
     * Image id
     */
    public /*out*/ readonly imageId!: pulumi.Output<string>;
    /**
     * Instance name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Create network interfaces
     */
    public readonly network!: pulumi.Output<outputs.CloudProject.InstanceNetwork>;
    /**
     * Instance region
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Existing SSH Keypair
     */
    public readonly sshKey!: pulumi.Output<outputs.CloudProject.InstanceSshKey | undefined>;
    /**
     * Add existing SSH Key pair into your Public Cloud project and link it to the instance
     */
    public readonly sshKeyCreate!: pulumi.Output<outputs.CloudProject.InstanceSshKeyCreate | undefined>;
    /**
     * Instance status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Instance task state
     */
    public /*out*/ readonly taskState!: pulumi.Output<string>;
    /**
     * Configuration information or scripts to use upon launch
     */
    public readonly userData!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["attachedVolumes"] = state ? state.attachedVolumes : undefined;
            resourceInputs["autoBackup"] = state ? state.autoBackup : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["billingPeriod"] = state ? state.billingPeriod : undefined;
            resourceInputs["bootFrom"] = state ? state.bootFrom : undefined;
            resourceInputs["bulk"] = state ? state.bulk : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["flavorName"] = state ? state.flavorName : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sshKey"] = state ? state.sshKey : undefined;
            resourceInputs["sshKeyCreate"] = state ? state.sshKeyCreate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["taskState"] = state ? state.taskState : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.billingPeriod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingPeriod'");
            }
            if ((!args || args.bootFrom === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bootFrom'");
            }
            if ((!args || args.flavor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavor'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["autoBackup"] = args ? args.autoBackup : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["billingPeriod"] = args ? args.billingPeriod : undefined;
            resourceInputs["bootFrom"] = args ? args.bootFrom : undefined;
            resourceInputs["bulk"] = args ? args.bulk : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["sshKey"] = args ? args.sshKey : undefined;
            resourceInputs["sshKeyCreate"] = args ? args.sshKeyCreate : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["addresses"] = undefined /*out*/;
            resourceInputs["attachedVolumes"] = undefined /*out*/;
            resourceInputs["flavorId"] = undefined /*out*/;
            resourceInputs["flavorName"] = undefined /*out*/;
            resourceInputs["imageId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["taskState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Instance IP addresses
     */
    addresses?: pulumi.Input<pulumi.Input<inputs.CloudProject.InstanceAddress>[]>;
    /**
     * Volumes attached to the instance
     */
    attachedVolumes?: pulumi.Input<pulumi.Input<inputs.CloudProject.InstanceAttachedVolume>[]>;
    /**
     * Create an autobackup workflow after instance start up.
     */
    autoBackup?: pulumi.Input<inputs.CloudProject.InstanceAutoBackup>;
    /**
     * The availability zone where the instance will be created
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Billing period - hourly or monthly
     */
    billingPeriod?: pulumi.Input<string>;
    /**
     * Boot the instance from an image or a volume
     */
    bootFrom?: pulumi.Input<inputs.CloudProject.InstanceBootFrom>;
    /**
     * Create multiple instances
     */
    bulk?: pulumi.Input<number>;
    /**
     * Flavor information
     */
    flavor?: pulumi.Input<inputs.CloudProject.InstanceFlavor>;
    /**
     * Flavor id
     */
    flavorId?: pulumi.Input<string>;
    /**
     * Flavor name
     */
    flavorName?: pulumi.Input<string>;
    /**
     * Start instance in group
     */
    group?: pulumi.Input<inputs.CloudProject.InstanceGroup>;
    /**
     * Image id
     */
    imageId?: pulumi.Input<string>;
    /**
     * Instance name
     */
    name?: pulumi.Input<string>;
    /**
     * Create network interfaces
     */
    network?: pulumi.Input<inputs.CloudProject.InstanceNetwork>;
    /**
     * Instance region
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Existing SSH Keypair
     */
    sshKey?: pulumi.Input<inputs.CloudProject.InstanceSshKey>;
    /**
     * Add existing SSH Key pair into your Public Cloud project and link it to the instance
     */
    sshKeyCreate?: pulumi.Input<inputs.CloudProject.InstanceSshKeyCreate>;
    /**
     * Instance status
     */
    status?: pulumi.Input<string>;
    /**
     * Instance task state
     */
    taskState?: pulumi.Input<string>;
    /**
     * Configuration information or scripts to use upon launch
     */
    userData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Create an autobackup workflow after instance start up.
     */
    autoBackup?: pulumi.Input<inputs.CloudProject.InstanceAutoBackup>;
    /**
     * The availability zone where the instance will be created
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Billing period - hourly or monthly
     */
    billingPeriod: pulumi.Input<string>;
    /**
     * Boot the instance from an image or a volume
     */
    bootFrom: pulumi.Input<inputs.CloudProject.InstanceBootFrom>;
    /**
     * Create multiple instances
     */
    bulk?: pulumi.Input<number>;
    /**
     * Flavor information
     */
    flavor: pulumi.Input<inputs.CloudProject.InstanceFlavor>;
    /**
     * Start instance in group
     */
    group?: pulumi.Input<inputs.CloudProject.InstanceGroup>;
    /**
     * Instance name
     */
    name?: pulumi.Input<string>;
    /**
     * Create network interfaces
     */
    network: pulumi.Input<inputs.CloudProject.InstanceNetwork>;
    /**
     * Instance region
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Existing SSH Keypair
     */
    sshKey?: pulumi.Input<inputs.CloudProject.InstanceSshKey>;
    /**
     * Add existing SSH Key pair into your Public Cloud project and link it to the instance
     */
    sshKeyCreate?: pulumi.Input<inputs.CloudProject.InstanceSshKeyCreate>;
    /**
     * Configuration information or scripts to use upon launch
     */
    userData?: pulumi.Input<string>;
}
