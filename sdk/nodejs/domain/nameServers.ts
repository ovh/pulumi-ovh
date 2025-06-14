// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this resource to manage a domain's name servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const nameServers = new ovh.domain.NameServers("name_servers", {
 *     domain: "mydomain.ovh",
 *     servers: [
 *         {
 *             host: "dns105.ovh.net",
 *             ip: "213.251.188.144",
 *         },
 *         {
 *             host: "ns105.ovh.net",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Name servers can be imported using their `domain`.
 *
 * Using the following configuration:
 *
 * terraform
 *
 * import {
 *
 *   to = ovh_domain_name_servers.name_servers
 *
 *   id = "<domain name>"
 *
 * }
 *
 * You can then run:
 *
 * bash
 *
 * $ pulumi preview -generate-config-out=name_servers.tf
 *
 * $ pulumi up
 *
 * The file `name_servers.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 */
export class NameServers extends pulumi.CustomResource {
    /**
     * Get an existing NameServers resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NameServersState, opts?: pulumi.CustomResourceOptions): NameServers {
        return new NameServers(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Domain/nameServers:NameServers';

    /**
     * Returns true if the given object is an instance of NameServers.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NameServers {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NameServers.__pulumiType;
    }

    /**
     * Domain name for which to manage name servers
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Details about a name server
     */
    public readonly servers!: pulumi.Output<outputs.Domain.NameServersServer[]>;

    /**
     * Create a NameServers resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NameServersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NameServersArgs | NameServersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NameServersState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as NameServersArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.servers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servers'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NameServers.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NameServers resources.
 */
export interface NameServersState {
    /**
     * Domain name for which to manage name servers
     */
    domain?: pulumi.Input<string>;
    /**
     * Details about a name server
     */
    servers?: pulumi.Input<pulumi.Input<inputs.Domain.NameServersServer>[]>;
}

/**
 * The set of arguments for constructing a NameServers resource.
 */
export interface NameServersArgs {
    /**
     * Domain name for which to manage name servers
     */
    domain: pulumi.Input<string>;
    /**
     * Details about a name server
     */
    servers: pulumi.Input<pulumi.Input<inputs.Domain.NameServersServer>[]>;
}
