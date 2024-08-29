// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject;

import com.ovh.ovh.CloudProject.ContainerRegistryArgs;
import com.ovh.ovh.CloudProject.inputs.ContainerRegistryState;
import com.ovh.ovh.CloudProject.outputs.ContainerRegistryPlan;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates a container registry associated with a public cloud project.
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
 * import com.pulumi.ovh.CloudProject.CloudProjectFunctions;
 * import com.pulumi.ovh.CloudProject.inputs.GetCapabilitiesContainerFilterArgs;
 * import com.pulumi.ovh.CloudProject.ContainerRegistry;
 * import com.pulumi.ovh.CloudProject.ContainerRegistryArgs;
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
 *         final var regcap = CloudProjectFunctions.getCapabilitiesContainerFilter(GetCapabilitiesContainerFilterArgs.builder()
 *             .serviceName("XXXXXX")
 *             .planName("SMALL")
 *             .region("GRA")
 *             .build());
 * 
 *         var my_registry = new ContainerRegistry("my-registry", ContainerRegistryArgs.builder()
 *             .serviceName(regcap.applyValue(getCapabilitiesContainerFilterResult -> getCapabilitiesContainerFilterResult.serviceName()))
 *             .planId(regcap.applyValue(getCapabilitiesContainerFilterResult -> getCapabilitiesContainerFilterResult.id()))
 *             .region(regcap.applyValue(getCapabilitiesContainerFilterResult -> getCapabilitiesContainerFilterResult.region()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; __WARNING__ You can update and migrate to a higher plan at any time but not the contrary.
 * 
 */
@ResourceType(type="ovh:CloudProject/containerRegistry:ContainerRegistry")
public class ContainerRegistry extends com.pulumi.resources.CustomResource {
    /**
     * Plan creation date
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Plan creation date
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Registry name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Registry name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Plan ID of the registry
     * 
     */
    @Export(name="planId", refs={String.class}, tree="[0]")
    private Output<String> planId;

    /**
     * @return Plan ID of the registry
     * 
     */
    public Output<String> planId() {
        return this.planId;
    }
    /**
     * Plan of the registry
     * 
     */
    @Export(name="plans", refs={List.class,ContainerRegistryPlan.class}, tree="[0,1]")
    private Output<List<ContainerRegistryPlan>> plans;

    /**
     * @return Plan of the registry
     * 
     */
    public Output<List<ContainerRegistryPlan>> plans() {
        return this.plans;
    }
    /**
     * Project ID of your registry
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return Project ID of your registry
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Region of the registry
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Region of the registry
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * Current size of the registry (bytes)
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return Current size of the registry (bytes)
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * Registry status
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Registry status
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Registry last update date
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Registry last update date
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Access url of the registry
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Access url of the registry
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Version of your registry
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Version of your registry
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerRegistry(java.lang.String name) {
        this(name, ContainerRegistryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerRegistry(java.lang.String name, ContainerRegistryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerRegistry(java.lang.String name, ContainerRegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistry:ContainerRegistry", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ContainerRegistry(java.lang.String name, Output<java.lang.String> id, @Nullable ContainerRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistry:ContainerRegistry", name, state, makeResourceOptions(options, id), false);
    }

    private static ContainerRegistryArgs makeArgs(ContainerRegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ContainerRegistryArgs.Empty : args;
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
    public static ContainerRegistry get(java.lang.String name, Output<java.lang.String> id, @Nullable ContainerRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerRegistry(name, id, state, options);
    }
}
