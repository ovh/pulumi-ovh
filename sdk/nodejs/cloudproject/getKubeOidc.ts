// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * export = async () => {
 *     const oidc = await ovh.CloudProject.getKubeOidc({
 *         serviceName: "XXXXXX",
 *         kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 *     });
 *     return {
 *         "oidc-val": oidc.clientId,
 *     };
 * }
 * ```
 */
export function getKubeOidc(args: GetKubeOidcArgs, opts?: pulumi.InvokeOptions): Promise<GetKubeOidcResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getKubeOidc:getKubeOidc", {
        "clientId": args.clientId,
        "issuerUrl": args.issuerUrl,
        "kubeId": args.kubeId,
        "oidcCaContent": args.oidcCaContent,
        "oidcGroupsClaims": args.oidcGroupsClaims,
        "oidcGroupsPrefix": args.oidcGroupsPrefix,
        "oidcRequiredClaims": args.oidcRequiredClaims,
        "oidcSigningAlgs": args.oidcSigningAlgs,
        "oidcUsernameClaim": args.oidcUsernameClaim,
        "oidcUsernamePrefix": args.oidcUsernamePrefix,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubeOidc.
 */
export interface GetKubeOidcArgs {
    /**
     * The OIDC client ID.
     */
    clientId?: string;
    /**
     * The OIDC issuer url.
     */
    issuerUrl?: string;
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: string;
    /**
     * Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
     */
    oidcCaContent?: string;
    /**
     * Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
     */
    oidcGroupsClaims?: string[];
    /**
     * Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
     */
    oidcGroupsPrefix?: string;
    /**
     * Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
     */
    oidcRequiredClaims?: string[];
    /**
     * Array of signing algorithms accepted. Default is \"RS256\".
     */
    oidcSigningAlgs?: string[];
    /**
     * JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
     */
    oidcUsernameClaim?: string;
    /**
     * Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
     */
    oidcUsernamePrefix?: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKubeOidc.
 */
export interface GetKubeOidcResult {
    /**
     * The OIDC client ID.
     */
    readonly clientId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The OIDC issuer url.
     */
    readonly issuerUrl?: string;
    /**
     * See Argument Reference above.
     */
    readonly kubeId: string;
    /**
     * Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
     */
    readonly oidcCaContent?: string;
    /**
     * Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
     */
    readonly oidcGroupsClaims?: string[];
    /**
     * Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
     */
    readonly oidcGroupsPrefix?: string;
    /**
     * Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
     */
    readonly oidcRequiredClaims?: string[];
    /**
     * Array of signing algorithms accepted. Default is \"RS256\".
     */
    readonly oidcSigningAlgs?: string[];
    /**
     * JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
     */
    readonly oidcUsernameClaim?: string;
    /**
     * Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
     */
    readonly oidcUsernamePrefix?: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * export = async () => {
 *     const oidc = await ovh.CloudProject.getKubeOidc({
 *         serviceName: "XXXXXX",
 *         kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 *     });
 *     return {
 *         "oidc-val": oidc.clientId,
 *     };
 * }
 * ```
 */
export function getKubeOidcOutput(args: GetKubeOidcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubeOidcResult> {
    return pulumi.output(args).apply((a: any) => getKubeOidc(a, opts))
}

/**
 * A collection of arguments for invoking getKubeOidc.
 */
export interface GetKubeOidcOutputArgs {
    /**
     * The OIDC client ID.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The OIDC issuer url.
     */
    issuerUrl?: pulumi.Input<string>;
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: pulumi.Input<string>;
    /**
     * Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
     */
    oidcCaContent?: pulumi.Input<string>;
    /**
     * Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
     */
    oidcGroupsClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
     */
    oidcGroupsPrefix?: pulumi.Input<string>;
    /**
     * Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
     */
    oidcRequiredClaims?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of signing algorithms accepted. Default is \"RS256\".
     */
    oidcSigningAlgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
     */
    oidcUsernameClaim?: pulumi.Input<string>;
    /**
     * Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
     */
    oidcUsernamePrefix?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
