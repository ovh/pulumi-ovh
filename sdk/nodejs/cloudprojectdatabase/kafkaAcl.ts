// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an ACL for a kafka cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const kafka = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXX",
 *     engine: "kafka",
 *     id: "ZZZ",
 * });
 * const acl = new ovh.cloudprojectdatabase.KafkaAcl("acl", {
 *     serviceName: kafka.then(kafka => kafka.serviceName),
 *     clusterId: kafka.then(kafka => kafka.id),
 *     permission: "read",
 *     topic: "mytopic",
 *     username: "johndoe",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,
 *
 *  bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
 * ```
 */
export class KafkaAcl extends pulumi.CustomResource {
    /**
     * Get an existing KafkaAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaAclState, opts?: pulumi.CustomResourceOptions): KafkaAcl {
        return new KafkaAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl';

    /**
     * Returns true if the given object is an instance of KafkaAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaAcl.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Permission to give to this username on this topic.
     * Available permissions:
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Topic affected by this ACL.
     */
    public readonly topic!: pulumi.Output<string>;
    /**
     * Username affected by this ACL.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a KafkaAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaAclArgs | KafkaAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaAclState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as KafkaAclArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KafkaAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaAcl resources.
 */
export interface KafkaAclState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Permission to give to this username on this topic.
     * Available permissions:
     */
    permission?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Topic affected by this ACL.
     */
    topic?: pulumi.Input<string>;
    /**
     * Username affected by this ACL.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KafkaAcl resource.
 */
export interface KafkaAclArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Permission to give to this username on this topic.
     * Available permissions:
     */
    permission: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Topic affected by this ACL.
     */
    topic: pulumi.Input<string>;
    /**
     * Username affected by this ACL.
     */
    username: pulumi.Input<string>;
}
