// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing;

import com.ovhcloud.pulumi.ovh.IpLoadBalancing.SslArgs;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs.SslState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a new custom SSL certificate on your IP Load Balancing
 * 
 * ## Import
 * 
 * SSL can be imported using the following format `service_name` and the `id` of the ssl, separated by &#34;/&#34; e.g.
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:IpLoadBalancing/ssl:Ssl sslname service_name/ssl_id
 * ```
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/ssl:Ssl")
public class Ssl extends com.pulumi.resources.CustomResource {
    /**
     * Certificate
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return Certificate
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * Certificate chain
     * 
     */
    @Export(name="chain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> chain;

    /**
     * @return Certificate chain
     * 
     */
    public Output<Optional<String>> chain() {
        return Codegen.optional(this.chain);
    }
    /**
     * Readable label for loadbalancer ssl
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Readable label for loadbalancer ssl
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Expire date of your SSL certificate.
     * 
     */
    @Export(name="expireDate", refs={String.class}, tree="[0]")
    private Output<String> expireDate;

    /**
     * @return Expire date of your SSL certificate.
     * 
     */
    public Output<String> expireDate() {
        return this.expireDate;
    }
    /**
     * Fingerprint of your SSL certificate.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return Fingerprint of your SSL certificate.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * Certificate key
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return Certificate key
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Subject Alternative Name of your SSL certificate.
     * 
     */
    @Export(name="sans", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> sans;

    /**
     * @return Subject Alternative Name of your SSL certificate.
     * 
     */
    public Output<List<String>> sans() {
        return this.sans;
    }
    /**
     * Serial of your SSL certificate (Deprecated, use fingerprint instead !)
     * 
     */
    @Export(name="serial", refs={String.class}, tree="[0]")
    private Output<String> serial;

    /**
     * @return Serial of your SSL certificate (Deprecated, use fingerprint instead !)
     * 
     */
    public Output<String> serial() {
        return this.serial;
    }
    /**
     * The internal name of your IP load balancing
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Subject of your SSL certificate.
     * 
     */
    @Export(name="subject", refs={String.class}, tree="[0]")
    private Output<String> subject;

    /**
     * @return Subject of your SSL certificate.
     * 
     */
    public Output<String> subject() {
        return this.subject;
    }
    /**
     * Type of your SSL certificate. &#39;built&#39; for SSL certificates managed by the IP Load Balancing. &#39;custom&#39; for user manager certificates.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of your SSL certificate. &#39;built&#39; for SSL certificates managed by the IP Load Balancing. &#39;custom&#39; for user manager certificates.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ssl(java.lang.String name) {
        this(name, SslArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ssl(java.lang.String name, SslArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ssl(java.lang.String name, SslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/ssl:Ssl", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Ssl(java.lang.String name, Output<java.lang.String> id, @Nullable SslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/ssl:Ssl", name, state, makeResourceOptions(options, id), false);
    }

    private static SslArgs makeArgs(SslArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SslArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .pluginDownloadURL("github://api.github.com/ovh/pulumi-ovh")
            .additionalSecretOutputs(List.of(
                "key"
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
    public static Ssl get(java.lang.String name, Output<java.lang.String> id, @Nullable SslState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ssl(name, id, state, options);
    }
}
