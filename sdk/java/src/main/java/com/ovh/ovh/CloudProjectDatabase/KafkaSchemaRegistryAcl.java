// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase;

import com.ovh.ovh.CloudProjectDatabase.KafkaSchemaRegistryAclArgs;
import com.ovh.ovh.CloudProjectDatabase.inputs.KafkaSchemaRegistryAclState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.
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
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaSchemaRegistryAcl;
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaSchemaRegistryAclArgs;
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
 *         final var kafka = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXX")
 *             .engine("kafka")
 *             .id("ZZZ")
 *             .build());
 * 
 *         var schemaRegistryAcl = new KafkaSchemaRegistryAcl("schemaRegistryAcl", KafkaSchemaRegistryAclArgs.builder()
 *             .serviceName(kafka.applyValue(getDatabaseResult -> getDatabaseResult.serviceName()))
 *             .clusterId(kafka.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .permission("schema_registry_read")
 *             .resource("Subject:myResource")
 *             .username("johndoe")
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
 * OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl")
public class KafkaSchemaRegistryAcl extends com.pulumi.resources.CustomResource {
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
     * Permission to give to this username on this resource.
     * Available permissions:
     * 
     */
    @Export(name="permission", refs={String.class}, tree="[0]")
    private Output<String> permission;

    /**
     * @return Permission to give to this username on this resource.
     * Available permissions:
     * 
     */
    public Output<String> permission() {
        return this.permission;
    }
    /**
     * Resource affected by this schema registry ACL.
     * 
     */
    @Export(name="resource", refs={String.class}, tree="[0]")
    private Output<String> resource;

    /**
     * @return Resource affected by this schema registry ACL.
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Username affected by this schema registry ACL.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return Username affected by this schema registry ACL.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KafkaSchemaRegistryAcl(java.lang.String name) {
        this(name, KafkaSchemaRegistryAclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KafkaSchemaRegistryAcl(java.lang.String name, KafkaSchemaRegistryAclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KafkaSchemaRegistryAcl(java.lang.String name, KafkaSchemaRegistryAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private KafkaSchemaRegistryAcl(java.lang.String name, Output<java.lang.String> id, @Nullable KafkaSchemaRegistryAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, state, makeResourceOptions(options, id), false);
    }

    private static KafkaSchemaRegistryAclArgs makeArgs(KafkaSchemaRegistryAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? KafkaSchemaRegistryAclArgs.Empty : args;
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
    public static KafkaSchemaRegistryAcl get(java.lang.String name, Output<java.lang.String> id, @Nullable KafkaSchemaRegistryAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KafkaSchemaRegistryAcl(name, id, state, options);
    }
}
