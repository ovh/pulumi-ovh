// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated;

import com.ovhcloud.pulumi.ovh.Dedicated.NasHAPartitionAccessArgs;
import com.ovhcloud.pulumi.ovh.Dedicated.inputs.NasHAPartitionAccessState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing access rights to partitions on HA-NAS services
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
 * import com.ovhcloud.pulumi.ovh.Dedicated.NasHAPartitionAccess;
 * import com.ovhcloud.pulumi.ovh.Dedicated.NasHAPartitionAccessArgs;
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
 *         var myPartition = new NasHAPartitionAccess("myPartition", NasHAPartitionAccessArgs.builder()
 *             .serviceName("zpool-12345")
 *             .partitionName("my-partition")
 *             .ip("123.123.123.123/32")
 *             .type("readwrite")
 *             .aclDescription("Description of the ACL")
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
 * HA-NAS partition access can be imported using the `{service_name}/{partition_name}/{ip}`, e.g.
 * 
 * ```sh
 * $ pulumi import ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess my-partition zpool-12345/my-partition/123.123.123.123%2F32`
 * ```
 * 
 */
@ResourceType(type="ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess")
public class NasHAPartitionAccess extends com.pulumi.resources.CustomResource {
    /**
     * A brief description of the acl
     * 
     */
    @Export(name="aclDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aclDescription;

    /**
     * @return A brief description of the acl
     * 
     */
    public Output<Optional<String>> aclDescription() {
        return Codegen.optional(this.aclDescription);
    }
    /**
     * IP block in x.x.x.x/x format
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return IP block in x.x.x.x/x format
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Name of the partition
     * 
     */
    @Export(name="partitionName", refs={String.class}, tree="[0]")
    private Output<String> partitionName;

    /**
     * @return Name of the partition
     * 
     */
    public Output<String> partitionName() {
        return this.partitionName;
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
     * One of &#34;readwrite&#34;, &#34;readonly&#34;
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return One of &#34;readwrite&#34;, &#34;readonly&#34;
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NasHAPartitionAccess(java.lang.String name) {
        this(name, NasHAPartitionAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NasHAPartitionAccess(java.lang.String name, NasHAPartitionAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NasHAPartitionAccess(java.lang.String name, NasHAPartitionAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NasHAPartitionAccess(java.lang.String name, Output<java.lang.String> id, @Nullable NasHAPartitionAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, state, makeResourceOptions(options, id), false);
    }

    private static NasHAPartitionAccessArgs makeArgs(NasHAPartitionAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NasHAPartitionAccessArgs.Empty : args;
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
    public static NasHAPartitionAccess get(java.lang.String name, Output<java.lang.String> id, @Nullable NasHAPartitionAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NasHAPartitionAccess(name, id, state, options);
    }
}
