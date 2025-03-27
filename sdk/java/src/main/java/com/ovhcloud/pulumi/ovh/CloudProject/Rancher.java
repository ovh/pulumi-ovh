// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.RancherArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.RancherCurrentState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.RancherCurrentTask;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.RancherTargetSpec;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Manage a Rancher service in a public cloud project.
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
 * import com.pulumi.ovh.CloudProject.Rancher;
 * import com.pulumi.ovh.CloudProject.RancherArgs;
 * import com.pulumi.ovh.CloudProject.inputs.RancherTargetSpecArgs;
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
 *         var rancher = new Rancher("rancher", RancherArgs.builder()
 *             .projectId("<public cloud project ID>")
 *             .targetSpec(RancherTargetSpecArgs.builder()
 *                 .name("MyRancher")
 *                 .plan("STANDARD")
 *                 .build())
 *             .build());
 * 
 *         ctx.export("rancherUrl", rancher.currentState().applyValue(currentState -> currentState.url()));
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * A share in a public cloud project can be imported using the `project_id` and `id` attributes. Using the following configuration:
 * 
 * terraform
 * 
 * import {
 * 
 *   id = &#34;&lt;project_id&gt;/&lt;id&gt;&#34;
 * 
 *   to = ovh_cloud_project_rancher.rancher
 * 
 * }
 * 
 * You can then run:
 * 
 * bash
 * 
 * $ pulumi preview -generate-config-out=rancher.tf
 * 
 * $ pulumi up
 * 
 * The file `rancher.tf` will then contain the imported resource&#39;s configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 * 
 */
@ResourceType(type="ovh:CloudProject/rancher:Rancher")
public class Rancher extends com.pulumi.resources.CustomResource {
    /**
     * Date of the managed Rancher service creation
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of the managed Rancher service creation
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Current configuration applied to the managed Rancher service
     * 
     */
    @Export(name="currentState", refs={RancherCurrentState.class}, tree="[0]")
    private Output<RancherCurrentState> currentState;

    /**
     * @return Current configuration applied to the managed Rancher service
     * 
     */
    public Output<RancherCurrentState> currentState() {
        return this.currentState;
    }
    /**
     * Asynchronous operations ongoing on the managed Rancher service
     * 
     */
    @Export(name="currentTasks", refs={List.class,RancherCurrentTask.class}, tree="[0,1]")
    private Output<List<RancherCurrentTask>> currentTasks;

    /**
     * @return Asynchronous operations ongoing on the managed Rancher service
     * 
     */
    public Output<List<RancherCurrentTask>> currentTasks() {
        return this.currentTasks;
    }
    /**
     * Project ID
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return Project ID
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
     * 
     */
    @Export(name="resourceStatus", refs={String.class}, tree="[0]")
    private Output<String> resourceStatus;

    /**
     * @return Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
     * 
     */
    public Output<String> resourceStatus() {
        return this.resourceStatus;
    }
    /**
     * Target specification for the managed Rancher service
     * 
     */
    @Export(name="targetSpec", refs={RancherTargetSpec.class}, tree="[0]")
    private Output<RancherTargetSpec> targetSpec;

    /**
     * @return Target specification for the managed Rancher service
     * 
     */
    public Output<RancherTargetSpec> targetSpec() {
        return this.targetSpec;
    }
    /**
     * Date of the last managed Rancher service update
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Date of the last managed Rancher service update
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rancher(java.lang.String name) {
        this(name, RancherArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rancher(java.lang.String name, RancherArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rancher(java.lang.String name, RancherArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/rancher:Rancher", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Rancher(java.lang.String name, Output<java.lang.String> id, @Nullable RancherState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/rancher:Rancher", name, state, makeResourceOptions(options, id), false);
    }

    private static RancherArgs makeArgs(RancherArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RancherArgs.Empty : args;
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
    public static Rancher get(java.lang.String name, Output<java.lang.String> id, @Nullable RancherState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rancher(name, id, state, options);
    }
}
