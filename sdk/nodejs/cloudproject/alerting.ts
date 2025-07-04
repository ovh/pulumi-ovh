// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates an alert on a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myAlert = new ovh.cloudproject.Alerting("my_alert", {
 *     serviceName: "XXX",
 *     delay: 3600,
 *     email: "aaa.bbb@domain.com",
 *     monthlyThreshold: 1000,
 * });
 * ```
 */
export class Alerting extends pulumi.CustomResource {
    /**
     * Get an existing Alerting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertingState, opts?: pulumi.CustomResourceOptions): Alerting {
        return new Alerting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/alerting:Alerting';

    /**
     * Returns true if the given object is an instance of Alerting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alerting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alerting.__pulumiType;
    }

    /**
     * Alerting creation date
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Delay between two alerts in seconds
     */
    public readonly delay!: pulumi.Output<number>;
    /**
     * Email to contact
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Formatted monthly threshold for this alerting
     */
    public /*out*/ readonly formattedMonthlyThreshold!: pulumi.Output<outputs.CloudProject.AlertingFormattedMonthlyThreshold>;
    /**
     * Monthly threshold for this alerting in currency
     */
    public readonly monthlyThreshold!: pulumi.Output<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a Alerting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertingArgs | AlertingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertingState | undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["delay"] = state ? state.delay : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["formattedMonthlyThreshold"] = state ? state.formattedMonthlyThreshold : undefined;
            resourceInputs["monthlyThreshold"] = state ? state.monthlyThreshold : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as AlertingArgs | undefined;
            if ((!args || args.delay === undefined) && !opts.urn) {
                throw new Error("Missing required property 'delay'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.monthlyThreshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monthlyThreshold'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["delay"] = args ? args.delay : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["monthlyThreshold"] = args ? args.monthlyThreshold : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["formattedMonthlyThreshold"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alerting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alerting resources.
 */
export interface AlertingState {
    /**
     * Alerting creation date
     */
    creationDate?: pulumi.Input<string>;
    /**
     * Delay between two alerts in seconds
     */
    delay?: pulumi.Input<number>;
    /**
     * Email to contact
     */
    email?: pulumi.Input<string>;
    /**
     * Formatted monthly threshold for this alerting
     */
    formattedMonthlyThreshold?: pulumi.Input<inputs.CloudProject.AlertingFormattedMonthlyThreshold>;
    /**
     * Monthly threshold for this alerting in currency
     */
    monthlyThreshold?: pulumi.Input<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alerting resource.
 */
export interface AlertingArgs {
    /**
     * Delay between two alerts in seconds
     */
    delay: pulumi.Input<number>;
    /**
     * Email to contact
     */
    email: pulumi.Input<string>;
    /**
     * Monthly threshold for this alerting in currency
     */
    monthlyThreshold: pulumi.Input<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
