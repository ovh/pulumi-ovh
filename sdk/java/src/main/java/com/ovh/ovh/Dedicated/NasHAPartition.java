// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dedicated;

import com.ovh.ovh.Dedicated.NasHAPartitionArgs;
import com.ovh.ovh.Dedicated.inputs.NasHAPartitionState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing partitions on HA-NAS services
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
 * import com.pulumi.ovh.Dedicated.NasHAPartition;
 * import com.pulumi.ovh.Dedicated.NasHAPartitionArgs;
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
 *         var my_partition = new NasHAPartition("my-partition", NasHAPartitionArgs.builder()
 *             .protocol("NFS")
 *             .serviceName("zpool-12345")
 *             .size(20)
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
 * HA-NAS can be imported using the `{service_name}/{name}`, e.g.
 * 
 * ```sh
 * $ pulumi import ovh:Dedicated/nasHAPartition:NasHAPartition my-partition zpool-12345/my-partition`
 * ```
 * 
 */
@ResourceType(type="ovh:Dedicated/nasHAPartition:NasHAPartition")
public class NasHAPartition extends com.pulumi.resources.CustomResource {
    /**
     * Percentage of partition space used in %
     * 
     */
    @Export(name="capacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> capacity;

    /**
     * @return Percentage of partition space used in %
     * 
     */
    public Output<Integer> capacity() {
        return this.capacity;
    }
    /**
     * A brief description of the partition
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of the partition
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * name of the partition
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return name of the partition
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * size of the partition in GB
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return size of the partition in GB
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * Percentage of partition space used by snapshots in %
     * 
     */
    @Export(name="usedBySnapshots", refs={Integer.class}, tree="[0]")
    private Output<Integer> usedBySnapshots;

    /**
     * @return Percentage of partition space used by snapshots in %
     * 
     */
    public Output<Integer> usedBySnapshots() {
        return this.usedBySnapshots;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NasHAPartition(java.lang.String name) {
        this(name, NasHAPartitionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NasHAPartition(java.lang.String name, NasHAPartitionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NasHAPartition(java.lang.String name, NasHAPartitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartition:NasHAPartition", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NasHAPartition(java.lang.String name, Output<java.lang.String> id, @Nullable NasHAPartitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartition:NasHAPartition", name, state, makeResourceOptions(options, id), false);
    }

    private static NasHAPartitionArgs makeArgs(NasHAPartitionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NasHAPartitionArgs.Empty : args;
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
    public static NasHAPartition get(java.lang.String name, Output<java.lang.String> id, @Nullable NasHAPartitionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NasHAPartition(name, id, state, options);
    }
}
