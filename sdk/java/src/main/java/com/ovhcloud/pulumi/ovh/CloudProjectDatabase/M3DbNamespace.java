// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase;

import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.M3DbNamespaceArgs;
import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs.M3DbNamespaceState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a namespace for a M3DB cluster associated with a public cloud project.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase.CloudProjectDatabaseFunctions;
 * import com.pulumi.ovh.CloudProjectDatabase.inputs.GetDatabaseArgs;
 * import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.M3DbNamespace;
 * import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.M3DbNamespaceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var m3db = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXX")
 *             .engine("m3db")
 *             .id("ZZZ")
 *             .build());
 * 
 *         var namespace = new M3DbNamespace("namespace", M3DbNamespaceArgs.builder()
 *             .serviceName(m3db.serviceName())
 *             .clusterId(m3db.id())
 *             .name("mynamespace")
 *             .resolution("P2D")
 *             .retentionPeriodDuration("PT48H")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OVHcloud Managed M3DB clusters namespaces can be imported using the `service_name`, `cluster_id` and `id` of the namespace, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace my_namespace service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace")
public class M3DbNamespace extends com.pulumi.resources.CustomResource {
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
     * Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="resolution", refs={String.class}, tree="[0]")
    private Output<String> resolution;

    /**
     * @return Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<String> resolution() {
        return this.resolution;
    }
    /**
     * Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="retentionBlockDataExpirationDuration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> retentionBlockDataExpirationDuration;

    /**
     * @return Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<Optional<String>> retentionBlockDataExpirationDuration() {
        return Codegen.optional(this.retentionBlockDataExpirationDuration);
    }
    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="retentionBlockSizeDuration", refs={String.class}, tree="[0]")
    private Output<String> retentionBlockSizeDuration;

    /**
     * @return Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<String> retentionBlockSizeDuration() {
        return this.retentionBlockSizeDuration;
    }
    /**
     * Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="retentionBufferFutureDuration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> retentionBufferFutureDuration;

    /**
     * @return Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<Optional<String>> retentionBufferFutureDuration() {
        return Codegen.optional(this.retentionBufferFutureDuration);
    }
    /**
     * Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="retentionBufferPastDuration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> retentionBufferPastDuration;

    /**
     * @return Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<Optional<String>> retentionBufferPastDuration() {
        return Codegen.optional(this.retentionBufferPastDuration);
    }
    /**
     * Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Export(name="retentionPeriodDuration", refs={String.class}, tree="[0]")
    private Output<String> retentionPeriodDuration;

    /**
     * @return Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Output<String> retentionPeriodDuration() {
        return this.retentionPeriodDuration;
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
     * Defines whether M3DB will create snapshot files for this namespace.
     * 
     */
    @Export(name="snapshotEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> snapshotEnabled;

    /**
     * @return Defines whether M3DB will create snapshot files for this namespace.
     * 
     */
    public Output<Boolean> snapshotEnabled() {
        return this.snapshotEnabled;
    }
    /**
     * Type of namespace.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of namespace.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     * 
     */
    @Export(name="writesToCommitLogEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> writesToCommitLogEnabled;

    /**
     * @return Defines whether M3DB will include writes to this namespace in the commit log.
     * 
     */
    public Output<Boolean> writesToCommitLogEnabled() {
        return this.writesToCommitLogEnabled;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public M3DbNamespace(java.lang.String name) {
        this(name, M3DbNamespaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public M3DbNamespace(java.lang.String name, M3DbNamespaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public M3DbNamespace(java.lang.String name, M3DbNamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private M3DbNamespace(java.lang.String name, Output<java.lang.String> id, @Nullable M3DbNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace", name, state, makeResourceOptions(options, id), false);
    }

    private static M3DbNamespaceArgs makeArgs(M3DbNamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? M3DbNamespaceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static M3DbNamespace get(java.lang.String name, Output<java.lang.String> id, @Nullable M3DbNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new M3DbNamespace(name, id, state, options);
    }
}
