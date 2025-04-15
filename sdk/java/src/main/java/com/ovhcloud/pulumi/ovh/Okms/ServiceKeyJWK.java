// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.ovhcloud.pulumi.ovh.Okms.ServiceKeyJWKArgs;
import com.ovhcloud.pulumi.ovh.Okms.inputs.ServiceKeyJWKState;
import com.ovhcloud.pulumi.ovh.Okms.outputs.ServiceKeyJWKIam;
import com.ovhcloud.pulumi.ovh.Okms.outputs.ServiceKeyJWKKey;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Import an existing key in the JWK format in an OVHcloud KMS.
 * 
 */
@ResourceType(type="ovh:Okms/serviceKeyJWK:ServiceKeyJWK")
public class ServiceKeyJWK extends com.pulumi.resources.CustomResource {
    /**
     * Context of the key
     * 
     */
    @Export(name="context", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> context;

    /**
     * @return Context of the key
     * 
     */
    public Output<Optional<String>> context() {
        return Codegen.optional(this.context);
    }
    /**
     * Creation time of the key
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Creation time of the key
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Key deactivation reason
     * 
     */
    @Export(name="deactivationReason", refs={String.class}, tree="[0]")
    private Output<String> deactivationReason;

    /**
     * @return Key deactivation reason
     * 
     */
    public Output<String> deactivationReason() {
        return this.deactivationReason;
    }
    /**
     * IAM resource metadata
     * 
     */
    @Export(name="iam", refs={ServiceKeyJWKIam.class}, tree="[0]")
    private Output<ServiceKeyJWKIam> iam;

    /**
     * @return IAM resource metadata
     * 
     */
    public Output<ServiceKeyJWKIam> iam() {
        return this.iam;
    }
    /**
     * Set of JSON Web Keys to import
     * 
     */
    @Export(name="keys", refs={List.class,ServiceKeyJWKKey.class}, tree="[0,1]")
    private Output<List<ServiceKeyJWKKey>> keys;

    /**
     * @return Set of JSON Web Keys to import
     * 
     */
    public Output<List<ServiceKeyJWKKey>> keys() {
        return this.keys;
    }
    /**
     * Key name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Key name
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
     * Size of the key to be created
     * 
     */
    @Export(name="size", refs={Double.class}, tree="[0]")
    private Output<Double> size;

    /**
     * @return Size of the key to be created
     * 
     */
    public Output<Double> size() {
        return this.size;
    }
    /**
     * State of the key
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the key
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Type of the key to be created
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the key to be created
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceKeyJWK(java.lang.String name) {
        this(name, ServiceKeyJWKArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceKeyJWK(java.lang.String name, ServiceKeyJWKArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceKeyJWK(java.lang.String name, ServiceKeyJWKArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceKeyJWK(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceKeyJWKState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceKeyJWKArgs makeArgs(ServiceKeyJWKArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceKeyJWKArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .pluginDownloadURL("github://api.github.com/ovh/pulumi-ovh")
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
    public static ServiceKeyJWK get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceKeyJWKState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceKeyJWK(name, id, state, options);
    }
}
