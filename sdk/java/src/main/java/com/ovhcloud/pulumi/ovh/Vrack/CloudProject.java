// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack;

import com.ovhcloud.pulumi.ovh.Utilities;
import com.ovhcloud.pulumi.ovh.Vrack.CloudProjectArgs;
import com.ovhcloud.pulumi.ovh.Vrack.inputs.CloudProjectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Attach a Public Cloud Project to a VRack.
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
 * import com.ovhcloud.pulumi.ovh.Vrack.CloudProject;
 * import com.ovhcloud.pulumi.ovh.Vrack.CloudProjectArgs;
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
 *         var vcp = new CloudProject("vcp", CloudProjectArgs.builder()
 *             .serviceName("12345")
 *             .projectId("67890")
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
 * Attachment of a public cloud project and a VRack can be imported using the `service_name` (vRack identifier) and the `project_id` (Cloud Project identifier), separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Vrack/cloudProject:CloudProject myattach service_name/project_id
 * ```
 * 
 */
@ResourceType(type="ovh:Vrack/cloudProject:CloudProject")
public class CloudProject extends com.pulumi.resources.CustomResource {
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudProject(java.lang.String name) {
        this(name, CloudProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudProject(java.lang.String name, CloudProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudProject(java.lang.String name, CloudProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/cloudProject:CloudProject", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CloudProject(java.lang.String name, Output<java.lang.String> id, @Nullable CloudProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/cloudProject:CloudProject", name, state, makeResourceOptions(options, id), false);
    }

    private static CloudProjectArgs makeArgs(CloudProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CloudProjectArgs.Empty : args;
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
    public static CloudProject get(java.lang.String name, Output<java.lang.String> id, @Nullable CloudProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudProject(name, id, state, options);
    }
}
