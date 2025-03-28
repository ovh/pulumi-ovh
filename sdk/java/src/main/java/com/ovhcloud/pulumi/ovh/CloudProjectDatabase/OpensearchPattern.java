// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase;

import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.OpensearchPatternArgs;
import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs.OpensearchPatternState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a pattern for a opensearch cluster associated with a public cloud project.
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
 * import com.pulumi.ovh.CloudProjectDatabase.OpensearchPattern;
 * import com.pulumi.ovh.CloudProjectDatabase.OpensearchPatternArgs;
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
 *         final var opensearch = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXX")
 *             .engine("opensearch")
 *             .id("ZZZ")
 *             .build());
 * 
 *         var pattern = new OpensearchPattern("pattern", OpensearchPatternArgs.builder()
 *             .serviceName(opensearch.applyValue(getDatabaseResult -> getDatabaseResult.serviceName()))
 *             .clusterId(opensearch.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .maxIndexCount(2)
 *             .pattern("logs_*")
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
 * OVHcloud Managed opensearch clusters patterns can be imported using the `service_name`, `cluster_id` and `id` of the pattern, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern my_pattern service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern")
public class OpensearchPattern extends com.pulumi.resources.CustomResource {
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
     * Maximum number of index for this pattern.
     * 
     */
    @Export(name="maxIndexCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxIndexCount;

    /**
     * @return Maximum number of index for this pattern.
     * 
     */
    public Output<Optional<Integer>> maxIndexCount() {
        return Codegen.optional(this.maxIndexCount);
    }
    /**
     * Pattern format.
     * 
     */
    @Export(name="pattern", refs={String.class}, tree="[0]")
    private Output<String> pattern;

    /**
     * @return Pattern format.
     * 
     */
    public Output<String> pattern() {
        return this.pattern;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OpensearchPattern(java.lang.String name) {
        this(name, OpensearchPatternArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OpensearchPattern(java.lang.String name, OpensearchPatternArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OpensearchPattern(java.lang.String name, OpensearchPatternArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OpensearchPattern(java.lang.String name, Output<java.lang.String> id, @Nullable OpensearchPatternState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/opensearchPattern:OpensearchPattern", name, state, makeResourceOptions(options, id), false);
    }

    private static OpensearchPatternArgs makeArgs(OpensearchPatternArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OpensearchPatternArgs.Empty : args;
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
    public static OpensearchPattern get(java.lang.String name, Output<java.lang.String> id, @Nullable OpensearchPatternState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OpensearchPattern(name, id, state, options);
    }
}
