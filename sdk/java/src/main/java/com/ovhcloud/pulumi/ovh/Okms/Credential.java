// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.ovhcloud.pulumi.ovh.Okms.CredentialArgs;
import com.ovhcloud.pulumi.ovh.Okms.inputs.CredentialState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates a credential for an OVHcloud KMS.
 * 
 */
@ResourceType(type="ovh:Okms/credential:Credential")
public class Credential extends com.pulumi.resources.CustomResource {
    /**
     * (String) Certificate PEM of the credential.
     * 
     */
    @Export(name="certificatePem", refs={String.class}, tree="[0]")
    private Output<String> certificatePem;

    /**
     * @return (String) Certificate PEM of the credential.
     * 
     */
    public Output<String> certificatePem() {
        return this.certificatePem;
    }
    /**
     * (String) Creation time of the credential
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return (String) Creation time of the credential
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Valid Certificate Signing Request
     * 
     */
    @Export(name="csr", refs={String.class}, tree="[0]")
    private Output<String> csr;

    /**
     * @return Valid Certificate Signing Request
     * 
     */
    public Output<String> csr() {
        return this.csr;
    }
    /**
     * Description of the credential (max 200)
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the credential (max 200)
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * (String) Expiration time of the credential
     * 
     */
    @Export(name="expiredAt", refs={String.class}, tree="[0]")
    private Output<String> expiredAt;

    /**
     * @return (String) Expiration time of the credential
     * 
     */
    public Output<String> expiredAt() {
        return this.expiredAt;
    }
    /**
     * (Boolean) Whether the credential was generated from a CSR
     * 
     */
    @Export(name="fromCsr", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> fromCsr;

    /**
     * @return (Boolean) Whether the credential was generated from a CSR
     * 
     */
    public Output<Boolean> fromCsr() {
        return this.fromCsr;
    }
    /**
     * List of identity URNs associated with the credential (max 25)
     * 
     */
    @Export(name="identityUrns", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> identityUrns;

    /**
     * @return List of identity URNs associated with the credential (max 25)
     * 
     */
    public Output<List<String>> identityUrns() {
        return this.identityUrns;
    }
    /**
     * Name of the credential (max 50)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the credential (max 50)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Okms ID
     * 
     */
    @Export(name="okmsId", refs={String.class}, tree="[0]")
    private Output<String> okmsId;

    /**
     * @return Okms ID
     * 
     */
    public Output<String> okmsId() {
        return this.okmsId;
    }
    /**
     * (String, Sensitive) Private Key PEM of the credential if no CSR is provided
     * 
     */
    @Export(name="privateKeyPem", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPem;

    /**
     * @return (String, Sensitive) Private Key PEM of the credential if no CSR is provided
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    /**
     * (String) Status of the credential
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return (String) Status of the credential
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Validity in days (default 365, max 365)
     * 
     */
    @Export(name="validity", refs={Double.class}, tree="[0]")
    private Output<Double> validity;

    /**
     * @return Validity in days (default 365, max 365)
     * 
     */
    public Output<Double> validity() {
        return this.validity;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Credential(java.lang.String name) {
        this(name, CredentialArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Credential(java.lang.String name, CredentialArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Credential(java.lang.String name, CredentialArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/credential:Credential", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Credential(java.lang.String name, Output<java.lang.String> id, @Nullable CredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/credential:Credential", name, state, makeResourceOptions(options, id), false);
    }

    private static CredentialArgs makeArgs(CredentialArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CredentialArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .pluginDownloadURL("github://api.github.com/ovh/pulumi-ovh")
            .additionalSecretOutputs(List.of(
                "privateKeyPem"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Credential get(java.lang.String name, Output<java.lang.String> id, @Nullable CredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Credential(name, id, state, options);
    }
}
