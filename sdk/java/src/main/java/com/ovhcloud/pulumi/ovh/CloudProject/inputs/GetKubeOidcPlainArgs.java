// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubeOidcPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeOidcPlainArgs Empty = new GetKubeOidcPlainArgs();

    /**
     * The OIDC client ID.
     * 
     */
    @Import(name="clientId")
    private @Nullable String clientId;

    /**
     * @return The OIDC client ID.
     * 
     */
    public Optional<String> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The OIDC issuer url.
     * 
     */
    @Import(name="issuerUrl")
    private @Nullable String issuerUrl;

    /**
     * @return The OIDC issuer url.
     * 
     */
    public Optional<String> issuerUrl() {
        return Optional.ofNullable(this.issuerUrl);
    }

    /**
     * The id of the managed kubernetes cluster.
     * 
     */
    @Import(name="kubeId", required=true)
    private String kubeId;

    /**
     * @return The id of the managed kubernetes cluster.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }

    /**
     * Content of the certificate for the CA, in base64 format, that signed your identity provider&#39;s web certificate. Defaults to the host&#39;s root CAs.
     * 
     */
    @Import(name="oidcCaContent")
    private @Nullable String oidcCaContent;

    /**
     * @return Content of the certificate for the CA, in base64 format, that signed your identity provider&#39;s web certificate. Defaults to the host&#39;s root CAs.
     * 
     */
    public Optional<String> oidcCaContent() {
        return Optional.ofNullable(this.oidcCaContent);
    }

    /**
     * Array of JWT claim to use as the user&#39;s group. If the claim is present it must be an array of strings.
     * 
     */
    @Import(name="oidcGroupsClaims")
    private @Nullable List<String> oidcGroupsClaims;

    /**
     * @return Array of JWT claim to use as the user&#39;s group. If the claim is present it must be an array of strings.
     * 
     */
    public Optional<List<String>> oidcGroupsClaims() {
        return Optional.ofNullable(this.oidcGroupsClaims);
    }

    /**
     * Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
     * 
     */
    @Import(name="oidcGroupsPrefix")
    private @Nullable String oidcGroupsPrefix;

    /**
     * @return Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
     * 
     */
    public Optional<String> oidcGroupsPrefix() {
        return Optional.ofNullable(this.oidcGroupsPrefix);
    }

    /**
     * Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.&#34;
     * 
     */
    @Import(name="oidcRequiredClaims")
    private @Nullable List<String> oidcRequiredClaims;

    /**
     * @return Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.&#34;
     * 
     */
    public Optional<List<String>> oidcRequiredClaims() {
        return Optional.ofNullable(this.oidcRequiredClaims);
    }

    /**
     * Array of signing algorithms accepted. Default is \&#34;RS256\&#34;.
     * 
     */
    @Import(name="oidcSigningAlgs")
    private @Nullable List<String> oidcSigningAlgs;

    /**
     * @return Array of signing algorithms accepted. Default is \&#34;RS256\&#34;.
     * 
     */
    public Optional<List<String>> oidcSigningAlgs() {
        return Optional.ofNullable(this.oidcSigningAlgs);
    }

    /**
     * JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
     * 
     */
    @Import(name="oidcUsernameClaim")
    private @Nullable String oidcUsernameClaim;

    /**
     * @return JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
     * 
     */
    public Optional<String> oidcUsernameClaim() {
        return Optional.ofNullable(this.oidcUsernameClaim);
    }

    /**
     * Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn&#39;t set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
     * 
     */
    @Import(name="oidcUsernamePrefix")
    private @Nullable String oidcUsernamePrefix;

    /**
     * @return Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn&#39;t set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
     * 
     */
    public Optional<String> oidcUsernamePrefix() {
        return Optional.ofNullable(this.oidcUsernamePrefix);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetKubeOidcPlainArgs() {}

    private GetKubeOidcPlainArgs(GetKubeOidcPlainArgs $) {
        this.clientId = $.clientId;
        this.issuerUrl = $.issuerUrl;
        this.kubeId = $.kubeId;
        this.oidcCaContent = $.oidcCaContent;
        this.oidcGroupsClaims = $.oidcGroupsClaims;
        this.oidcGroupsPrefix = $.oidcGroupsPrefix;
        this.oidcRequiredClaims = $.oidcRequiredClaims;
        this.oidcSigningAlgs = $.oidcSigningAlgs;
        this.oidcUsernameClaim = $.oidcUsernameClaim;
        this.oidcUsernamePrefix = $.oidcUsernamePrefix;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeOidcPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeOidcPlainArgs $;

        public Builder() {
            $ = new GetKubeOidcPlainArgs();
        }

        public Builder(GetKubeOidcPlainArgs defaults) {
            $ = new GetKubeOidcPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The OIDC client ID.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable String clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param issuerUrl The OIDC issuer url.
         * 
         * @return builder
         * 
         */
        public Builder issuerUrl(@Nullable String issuerUrl) {
            $.issuerUrl = issuerUrl;
            return this;
        }

        /**
         * @param kubeId The id of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param oidcCaContent Content of the certificate for the CA, in base64 format, that signed your identity provider&#39;s web certificate. Defaults to the host&#39;s root CAs.
         * 
         * @return builder
         * 
         */
        public Builder oidcCaContent(@Nullable String oidcCaContent) {
            $.oidcCaContent = oidcCaContent;
            return this;
        }

        /**
         * @param oidcGroupsClaims Array of JWT claim to use as the user&#39;s group. If the claim is present it must be an array of strings.
         * 
         * @return builder
         * 
         */
        public Builder oidcGroupsClaims(@Nullable List<String> oidcGroupsClaims) {
            $.oidcGroupsClaims = oidcGroupsClaims;
            return this;
        }

        /**
         * @param oidcGroupsClaims Array of JWT claim to use as the user&#39;s group. If the claim is present it must be an array of strings.
         * 
         * @return builder
         * 
         */
        public Builder oidcGroupsClaims(String... oidcGroupsClaims) {
            return oidcGroupsClaims(List.of(oidcGroupsClaims));
        }

        /**
         * @param oidcGroupsPrefix Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
         * 
         * @return builder
         * 
         */
        public Builder oidcGroupsPrefix(@Nullable String oidcGroupsPrefix) {
            $.oidcGroupsPrefix = oidcGroupsPrefix;
            return this;
        }

        /**
         * @param oidcRequiredClaims Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.&#34;
         * 
         * @return builder
         * 
         */
        public Builder oidcRequiredClaims(@Nullable List<String> oidcRequiredClaims) {
            $.oidcRequiredClaims = oidcRequiredClaims;
            return this;
        }

        /**
         * @param oidcRequiredClaims Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.&#34;
         * 
         * @return builder
         * 
         */
        public Builder oidcRequiredClaims(String... oidcRequiredClaims) {
            return oidcRequiredClaims(List.of(oidcRequiredClaims));
        }

        /**
         * @param oidcSigningAlgs Array of signing algorithms accepted. Default is \&#34;RS256\&#34;.
         * 
         * @return builder
         * 
         */
        public Builder oidcSigningAlgs(@Nullable List<String> oidcSigningAlgs) {
            $.oidcSigningAlgs = oidcSigningAlgs;
            return this;
        }

        /**
         * @param oidcSigningAlgs Array of signing algorithms accepted. Default is \&#34;RS256\&#34;.
         * 
         * @return builder
         * 
         */
        public Builder oidcSigningAlgs(String... oidcSigningAlgs) {
            return oidcSigningAlgs(List.of(oidcSigningAlgs));
        }

        /**
         * @param oidcUsernameClaim JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
         * 
         * @return builder
         * 
         */
        public Builder oidcUsernameClaim(@Nullable String oidcUsernameClaim) {
            $.oidcUsernameClaim = oidcUsernameClaim;
            return this;
        }

        /**
         * @param oidcUsernamePrefix Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn&#39;t set and `oidc_username_claim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
         * 
         * @return builder
         * 
         */
        public Builder oidcUsernamePrefix(@Nullable String oidcUsernamePrefix) {
            $.oidcUsernamePrefix = oidcUsernamePrefix;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetKubeOidcPlainArgs build() {
            if ($.kubeId == null) {
                throw new MissingRequiredPropertyException("GetKubeOidcPlainArgs", "kubeId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetKubeOidcPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
