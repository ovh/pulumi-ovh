// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an OIDC configuration in an OVHcloud Managed Private Registry.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myOidc = new ovh.cloudproject.ContainerRegistryOIDC("my_oidc", {
 *     serviceName: "XXXXXX",
 *     registryId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 *     oidcName: "my-oidc-provider",
 *     oidcEndpoint: "https://xxxx.yyy.com",
 *     oidcClientId: "xxx",
 *     oidcClientSecret: "xxx",
 *     oidcScope: "openid,profile,email,offline_access",
 *     oidcGroupsClaim: "groups",
 *     oidcAdminGroup: "harbor-admin",
 *     oidcVerifyCert: true,
 *     oidcAutoOnboard: true,
 *     oidcUserClaim: "preferred_username",
 *     deleteUsers: false,
 * });
 * export const oidcClientSecret = myOidc.oidcClientSecret;
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed Private Registry OIDC can be imported using the tenant `service_name` and registry id `registry_id` separated by "/" E.g.,
 *
 * bash
 *
 * ```sh
 * $ pulumi import ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC my-oidc service_name/registry_id
 * ```
 */
export class ContainerRegistryOIDC extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRegistryOIDC resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerRegistryOIDCState, opts?: pulumi.CustomResourceOptions): ContainerRegistryOIDC {
        return new ContainerRegistryOIDC(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC';

    /**
     * Returns true if the given object is an instance of ContainerRegistryOIDC.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRegistryOIDC {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRegistryOIDC.__pulumiType;
    }

    /**
     * Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
     */
    public readonly deleteUsers!: pulumi.Output<boolean | undefined>;
    /**
     * Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
     */
    public readonly oidcAdminGroup!: pulumi.Output<string | undefined>;
    /**
     * Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
     */
    public readonly oidcAutoOnboard!: pulumi.Output<boolean | undefined>;
    /**
     * The client ID with which Harbor is registered as client application with the OIDC provider.
     */
    public readonly oidcClientId!: pulumi.Output<string>;
    /**
     * The secret for the Harbor client application.
     */
    public readonly oidcClientSecret!: pulumi.Output<string>;
    /**
     * The URL of an OIDC-compliant server.
     */
    public readonly oidcEndpoint!: pulumi.Output<string>;
    /**
     * The name of Claim in the ID token whose value is the list of group names.
     */
    public readonly oidcGroupsClaim!: pulumi.Output<string | undefined>;
    /**
     * The name of the OIDC provider.
     */
    public readonly oidcName!: pulumi.Output<string>;
    /**
     * The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
     */
    public readonly oidcScope!: pulumi.Output<string>;
    /**
     * The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
     */
    public readonly oidcUserClaim!: pulumi.Output<string | undefined>;
    /**
     * Set it to `false` if your OIDC server is hosted via self-signed certificate.
     */
    public readonly oidcVerifyCert!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Managed Private Registry. **Changing this value recreates the resource.**
     */
    public readonly registryId!: pulumi.Output<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a ContainerRegistryOIDC resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerRegistryOIDCArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerRegistryOIDCArgs | ContainerRegistryOIDCState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerRegistryOIDCState | undefined;
            resourceInputs["deleteUsers"] = state ? state.deleteUsers : undefined;
            resourceInputs["oidcAdminGroup"] = state ? state.oidcAdminGroup : undefined;
            resourceInputs["oidcAutoOnboard"] = state ? state.oidcAutoOnboard : undefined;
            resourceInputs["oidcClientId"] = state ? state.oidcClientId : undefined;
            resourceInputs["oidcClientSecret"] = state ? state.oidcClientSecret : undefined;
            resourceInputs["oidcEndpoint"] = state ? state.oidcEndpoint : undefined;
            resourceInputs["oidcGroupsClaim"] = state ? state.oidcGroupsClaim : undefined;
            resourceInputs["oidcName"] = state ? state.oidcName : undefined;
            resourceInputs["oidcScope"] = state ? state.oidcScope : undefined;
            resourceInputs["oidcUserClaim"] = state ? state.oidcUserClaim : undefined;
            resourceInputs["oidcVerifyCert"] = state ? state.oidcVerifyCert : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as ContainerRegistryOIDCArgs | undefined;
            if ((!args || args.oidcClientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcClientId'");
            }
            if ((!args || args.oidcClientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcClientSecret'");
            }
            if ((!args || args.oidcEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcEndpoint'");
            }
            if ((!args || args.oidcName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcName'");
            }
            if ((!args || args.oidcScope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcScope'");
            }
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["deleteUsers"] = args ? args.deleteUsers : undefined;
            resourceInputs["oidcAdminGroup"] = args ? args.oidcAdminGroup : undefined;
            resourceInputs["oidcAutoOnboard"] = args ? args.oidcAutoOnboard : undefined;
            resourceInputs["oidcClientId"] = args ? args.oidcClientId : undefined;
            resourceInputs["oidcClientSecret"] = args?.oidcClientSecret ? pulumi.secret(args.oidcClientSecret) : undefined;
            resourceInputs["oidcEndpoint"] = args ? args.oidcEndpoint : undefined;
            resourceInputs["oidcGroupsClaim"] = args ? args.oidcGroupsClaim : undefined;
            resourceInputs["oidcName"] = args ? args.oidcName : undefined;
            resourceInputs["oidcScope"] = args ? args.oidcScope : undefined;
            resourceInputs["oidcUserClaim"] = args ? args.oidcUserClaim : undefined;
            resourceInputs["oidcVerifyCert"] = args ? args.oidcVerifyCert : undefined;
            resourceInputs["registryId"] = args ? args.registryId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["oidcClientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ContainerRegistryOIDC.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerRegistryOIDC resources.
 */
export interface ContainerRegistryOIDCState {
    /**
     * Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
     */
    deleteUsers?: pulumi.Input<boolean>;
    /**
     * Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
     */
    oidcAdminGroup?: pulumi.Input<string>;
    /**
     * Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
     */
    oidcAutoOnboard?: pulumi.Input<boolean>;
    /**
     * The client ID with which Harbor is registered as client application with the OIDC provider.
     */
    oidcClientId?: pulumi.Input<string>;
    /**
     * The secret for the Harbor client application.
     */
    oidcClientSecret?: pulumi.Input<string>;
    /**
     * The URL of an OIDC-compliant server.
     */
    oidcEndpoint?: pulumi.Input<string>;
    /**
     * The name of Claim in the ID token whose value is the list of group names.
     */
    oidcGroupsClaim?: pulumi.Input<string>;
    /**
     * The name of the OIDC provider.
     */
    oidcName?: pulumi.Input<string>;
    /**
     * The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
     */
    oidcScope?: pulumi.Input<string>;
    /**
     * The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
     */
    oidcUserClaim?: pulumi.Input<string>;
    /**
     * Set it to `false` if your OIDC server is hosted via self-signed certificate.
     */
    oidcVerifyCert?: pulumi.Input<boolean>;
    /**
     * The ID of the Managed Private Registry. **Changing this value recreates the resource.**
     */
    registryId?: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerRegistryOIDC resource.
 */
export interface ContainerRegistryOIDCArgs {
    /**
     * Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
     */
    deleteUsers?: pulumi.Input<boolean>;
    /**
     * Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
     */
    oidcAdminGroup?: pulumi.Input<string>;
    /**
     * Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
     */
    oidcAutoOnboard?: pulumi.Input<boolean>;
    /**
     * The client ID with which Harbor is registered as client application with the OIDC provider.
     */
    oidcClientId: pulumi.Input<string>;
    /**
     * The secret for the Harbor client application.
     */
    oidcClientSecret: pulumi.Input<string>;
    /**
     * The URL of an OIDC-compliant server.
     */
    oidcEndpoint: pulumi.Input<string>;
    /**
     * The name of Claim in the ID token whose value is the list of group names.
     */
    oidcGroupsClaim?: pulumi.Input<string>;
    /**
     * The name of the OIDC provider.
     */
    oidcName: pulumi.Input<string>;
    /**
     * The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
     */
    oidcScope: pulumi.Input<string>;
    /**
     * The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
     */
    oidcUserClaim?: pulumi.Input<string>;
    /**
     * Set it to `false` if your OIDC server is hosted via self-signed certificate.
     */
    oidcVerifyCert?: pulumi.Input<boolean>;
    /**
     * The ID of the Managed Private Registry. **Changing this value recreates the resource.**
     */
    registryId: pulumi.Input<string>;
    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     */
    serviceName: pulumi.Input<string>;
}
