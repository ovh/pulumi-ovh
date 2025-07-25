// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Hosting;

import com.ovhcloud.pulumi.ovh.Hosting.PrivateDatabaseDbArgs;
import com.ovhcloud.pulumi.ovh.Hosting.inputs.PrivateDatabaseDbState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Create a new database on your private cloud database service.
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
 * import com.ovhcloud.pulumi.ovh.Hosting.PrivateDatabaseDb;
 * import com.ovhcloud.pulumi.ovh.Hosting.PrivateDatabaseDbArgs;
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
 *         var database = new PrivateDatabaseDb("database", PrivateDatabaseDbArgs.builder()
 *             .serviceName("XXXXXX")
 *             .databaseName("XXXXXX")
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
 * OVHcloud Webhosting database can be imported using the `service_name` and the `database_name`, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 * $ pulumi import ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb database service_name/database_name
 * ```
 * 
 */
@ResourceType(type="ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb")
public class PrivateDatabaseDb extends com.pulumi.resources.CustomResource {
    /**
     * Name of your new database
     * 
     */
    @Export(name="databaseName", refs={String.class}, tree="[0]")
    private Output<String> databaseName;

    /**
     * @return Name of your new database
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }
    /**
     * The internal name of your private database.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your private database.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateDatabaseDb(java.lang.String name) {
        this(name, PrivateDatabaseDbArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateDatabaseDb(java.lang.String name, PrivateDatabaseDbArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateDatabaseDb(java.lang.String name, PrivateDatabaseDbArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PrivateDatabaseDb(java.lang.String name, Output<java.lang.String> id, @Nullable PrivateDatabaseDbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, state, makeResourceOptions(options, id), false);
    }

    private static PrivateDatabaseDbArgs makeArgs(PrivateDatabaseDbArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PrivateDatabaseDbArgs.Empty : args;
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
    public static PrivateDatabaseDb get(java.lang.String name, Output<java.lang.String> id, @Nullable PrivateDatabaseDbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateDatabaseDb(name, id, state, options);
    }
}
