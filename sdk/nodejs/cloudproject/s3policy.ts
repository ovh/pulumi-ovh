// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Set the S3 Policy of a public cloud project user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const user = new ovh.cloudproject.User("user", {
 *     serviceName: "XXX",
 *     description: "my user",
 *     roleNames: ["objectstore_operator"],
 * });
 * const myS3Credentials = new ovh.cloudproject.S3Credential("my_s3_credentials", {
 *     serviceName: user.serviceName,
 *     userId: user.id,
 * });
 * const policy = new ovh.cloudproject.S3Policy("policy", {
 *     serviceName: user.serviceName,
 *     userId: user.id,
 *     policy: JSON.stringify({
 *         Statement: [{
 *             Sid: "RWContainer",
 *             Effect: "Allow",
 *             Action: [
 *                 "s3:GetObject",
 *                 "s3:PutObject",
 *                 "s3:DeleteObject",
 *                 "s3:ListBucket",
 *                 "s3:ListMultipartUploadParts",
 *                 "s3:ListBucketMultipartUploads",
 *                 "s3:AbortMultipartUpload",
 *                 "s3:GetBucketLocation",
 *             ],
 *             Resource: [
 *                 "arn:aws:s3:::hp-bucket",
 *                 "arn:aws:s3:::hp-bucket/*",
 *             ],
 *         }],
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud User S3 Policy can be imported using the `service_name`, `user_id` of the policy, separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProject/s3Policy:S3Policy policy service_name/user_id
 * ```
 */
export class S3Policy extends pulumi.CustomResource {
    /**
     * Get an existing S3Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3PolicyState, opts?: pulumi.CustomResourceOptions): S3Policy {
        return new S3Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/s3Policy:S3Policy';

    /**
     * Returns true if the given object is an instance of S3Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3Policy.__pulumiType;
    }

    /**
     * The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The ID of a public cloud project's user.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a S3Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3PolicyArgs | S3PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as S3PolicyState | undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as S3PolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(S3Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3Policy resources.
 */
export interface S3PolicyState {
    /**
     * The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
     */
    policy?: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The ID of a public cloud project's user.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3Policy resource.
 */
export interface S3PolicyArgs {
    /**
     * The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
     */
    policy: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The ID of a public cloud project's user.
     */
    userId: pulumi.Input<string>;
}
