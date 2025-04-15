// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this resource to manage a domain's DS records.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const dsRecords = new ovh.domain.DSRecords("ds_records", {
 *     domain: "mydomain.ovh",
 *     dsRecords: [{
 *         algorithm: "RSASHA1_NSEC3_SHA1",
 *         flags: "KEY_SIGNING_KEY",
 *         publicKey: "my_base64_encoded_public_key",
 *         tag: 12345,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * DS records can be imported using their `domain`.
 *
 * Using the following configuration:
 *
 * terraform
 *
 * import {
 *
 *   to = ovh_domain_ds_records.ds_records
 *
 *   id = "<domain name>"
 *
 * }
 *
 * You can then run:
 *
 * bash
 *
 * $ pulumi preview -generate-config-out=ds_records.tf
 *
 * $ pulumi up
 *
 * The file `ds_records.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 */
export class DSRecords extends pulumi.CustomResource {
    /**
     * Get an existing DSRecords resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DSRecordsState, opts?: pulumi.CustomResourceOptions): DSRecords {
        return new DSRecords(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Domain/dSRecords:DSRecords';

    /**
     * Returns true if the given object is an instance of DSRecords.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DSRecords {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DSRecords.__pulumiType;
    }

    /**
     * Domain name for which to manage DS records
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Details about a DS record
     */
    public readonly dsRecords!: pulumi.Output<outputs.Domain.DSRecordsDsRecord[]>;

    /**
     * Create a DSRecords resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DSRecordsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DSRecordsArgs | DSRecordsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DSRecordsState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["dsRecords"] = state ? state.dsRecords : undefined;
        } else {
            const args = argsOrState as DSRecordsArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.dsRecords === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dsRecords'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["dsRecords"] = args ? args.dsRecords : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DSRecords.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DSRecords resources.
 */
export interface DSRecordsState {
    /**
     * Domain name for which to manage DS records
     */
    domain?: pulumi.Input<string>;
    /**
     * Details about a DS record
     */
    dsRecords?: pulumi.Input<pulumi.Input<inputs.Domain.DSRecordsDsRecord>[]>;
}

/**
 * The set of arguments for constructing a DSRecords resource.
 */
export interface DSRecordsArgs {
    /**
     * Domain name for which to manage DS records
     */
    domain: pulumi.Input<string>;
    /**
     * Details about a DS record
     */
    dsRecords: pulumi.Input<pulumi.Input<inputs.Domain.DSRecordsDsRecord>[]>;
}
