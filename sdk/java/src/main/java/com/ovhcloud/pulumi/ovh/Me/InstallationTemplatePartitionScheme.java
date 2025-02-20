// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me;

import com.ovhcloud.pulumi.ovh.Me.InstallationTemplatePartitionSchemeArgs;
import com.ovhcloud.pulumi.ovh.Me.inputs.InstallationTemplatePartitionSchemeState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Use this resource to create partition scheme for a custom installation template available for dedicated servers.
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
 * import com.pulumi.ovh.Me.InstallationTemplate;
 * import com.pulumi.ovh.Me.InstallationTemplateArgs;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionScheme;
 * import com.pulumi.ovh.Me.InstallationTemplatePartitionSchemeArgs;
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
 *         var myTemplate = new InstallationTemplate("myTemplate", InstallationTemplateArgs.builder()
 *             .baseTemplateName("debian12_64")
 *             .templateName("mytemplate")
 *             .build());
 * 
 *         var scheme = new InstallationTemplatePartitionScheme("scheme", InstallationTemplatePartitionSchemeArgs.builder()
 *             .templateName(myTemplate.templateName())
 *             .priority(1)
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
 * The resource can be imported using the `template_name`, `name` of the cluster, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Me/installationTemplatePartitionScheme:InstallationTemplatePartitionScheme scheme template_name/name
 * ```
 * 
 */
@ResourceType(type="ovh:Me/installationTemplatePartitionScheme:InstallationTemplatePartitionScheme")
public class InstallationTemplatePartitionScheme extends com.pulumi.resources.CustomResource {
    /**
     * (Required) This partition scheme name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return (Required) This partition scheme name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The template name of the partition scheme.
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output<String> templateName;

    /**
     * @return The template name of the partition scheme.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstallationTemplatePartitionScheme(java.lang.String name) {
        this(name, InstallationTemplatePartitionSchemeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstallationTemplatePartitionScheme(java.lang.String name, InstallationTemplatePartitionSchemeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstallationTemplatePartitionScheme(java.lang.String name, InstallationTemplatePartitionSchemeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/installationTemplatePartitionScheme:InstallationTemplatePartitionScheme", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InstallationTemplatePartitionScheme(java.lang.String name, Output<java.lang.String> id, @Nullable InstallationTemplatePartitionSchemeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/installationTemplatePartitionScheme:InstallationTemplatePartitionScheme", name, state, makeResourceOptions(options, id), false);
    }

    private static InstallationTemplatePartitionSchemeArgs makeArgs(InstallationTemplatePartitionSchemeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstallationTemplatePartitionSchemeArgs.Empty : args;
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
    public static InstallationTemplatePartitionScheme get(java.lang.String name, Output<java.lang.String> id, @Nullable InstallationTemplatePartitionSchemeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstallationTemplatePartitionScheme(name, id, state, options);
    }
}
