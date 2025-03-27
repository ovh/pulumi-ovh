// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase;

import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.MongoDbPrometheusArgs;
import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs.MongoDbPrometheusState;
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
 * ## Import
 * 
 * OVHcloud Managed MongoDB clusters prometheus can be imported using the `service_name` and `cluster_id`, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/mongoDbPrometheus:MongoDbPrometheus my_prometheus service_name/engine/cluster_id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/mongoDbPrometheus:MongoDbPrometheus")
public class MongoDbPrometheus extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * (Sensitive) Password of the user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Sensitive) Password of the user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Arbitrary string to change to trigger a password update
     * 
     */
    @Export(name="passwordReset", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> passwordReset;

    /**
     * @return Arbitrary string to change to trigger a password update
     * 
     */
    public Output<Optional<String>> passwordReset() {
        return Codegen.optional(this.passwordReset);
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Name of the srv domain endpoint.
     * 
     */
    @Export(name="srvDomain", refs={String.class}, tree="[0]")
    private Output<String> srvDomain;

    /**
     * @return Name of the srv domain endpoint.
     * 
     */
    public Output<String> srvDomain() {
        return this.srvDomain;
    }
    /**
     * name of the prometheus user.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return name of the prometheus user.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MongoDbPrometheus(java.lang.String name) {
        this(name, MongoDbPrometheusArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MongoDbPrometheus(java.lang.String name, MongoDbPrometheusArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MongoDbPrometheus(java.lang.String name, MongoDbPrometheusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/mongoDbPrometheus:MongoDbPrometheus", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MongoDbPrometheus(java.lang.String name, Output<java.lang.String> id, @Nullable MongoDbPrometheusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/mongoDbPrometheus:MongoDbPrometheus", name, state, makeResourceOptions(options, id), false);
    }

    private static MongoDbPrometheusArgs makeArgs(MongoDbPrometheusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MongoDbPrometheusArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static MongoDbPrometheus get(java.lang.String name, Output<java.lang.String> id, @Nullable MongoDbPrometheusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MongoDbPrometheus(name, id, state, options);
    }
}
