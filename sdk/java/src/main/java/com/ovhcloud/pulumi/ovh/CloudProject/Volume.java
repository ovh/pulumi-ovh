// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.VolumeArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.VolumeState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.VolumeSubOperation;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Create volume in a public cloud project.
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
 * import com.ovhcloud.pulumi.ovh.CloudProject.Volume;
 * import com.ovhcloud.pulumi.ovh.CloudProject.VolumeArgs;
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
 *         var vol = new Volume("vol", VolumeArgs.builder()
 *             .regionName("xxx")
 *             .serviceName("yyyyy")
 *             .description("Terraform volume")
 *             .name("terrformName")
 *             .size(15.0)
 *             .type("classic")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/volume:Volume")
public class Volume extends com.pulumi.resources.CustomResource {
    /**
     * The action of the operation
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return The action of the operation
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * The completed date of the operation
     * 
     */
    @Export(name="completedAt", refs={String.class}, tree="[0]")
    private Output<String> completedAt;

    /**
     * @return The completed date of the operation
     * 
     */
    public Output<String> completedAt() {
        return this.completedAt;
    }
    /**
     * The creation date of the operation
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The creation date of the operation
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * A description of the volume
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return A description of the volume
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Image ID
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return Image ID
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * Instance ID
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Instance ID
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Name of the volume
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the volume
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Volume status
     * 
     */
    @Export(name="progress", refs={Double.class}, tree="[0]")
    private Output<Double> progress;

    /**
     * @return Volume status
     * 
     */
    public Output<Double> progress() {
        return this.progress;
    }
    /**
     * Required. A valid OVHcloud public cloud region name in which the volume will be available. Ex.: &#34;GRA11&#34;. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return Required. A valid OVHcloud public cloud region name in which the volume will be available. Ex.: &#34;GRA11&#34;. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }
    /**
     * List of regions
     * 
     */
    @Export(name="regions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> regions;

    /**
     * @return List of regions
     * 
     */
    public Output<List<String>> regions() {
        return this.regions;
    }
    /**
     * Id of the resource
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return Id of the resource
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * Required. The id of the public cloud project. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Required. The id of the public cloud project. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Size (GB) of the volume
     * 
     */
    @Export(name="size", refs={Double.class}, tree="[0]")
    private Output<Double> size;

    /**
     * @return Size (GB) of the volume
     * 
     */
    public Output<Double> size() {
        return this.size;
    }
    /**
     * Snapshot ID
     * 
     */
    @Export(name="snapshotId", refs={String.class}, tree="[0]")
    private Output<String> snapshotId;

    /**
     * @return Snapshot ID
     * 
     */
    public Output<String> snapshotId() {
        return this.snapshotId;
    }
    /**
     * Datetime of the operation creation
     * 
     */
    @Export(name="startedAt", refs={String.class}, tree="[0]")
    private Output<String> startedAt;

    /**
     * @return Datetime of the operation creation
     * 
     */
    public Output<String> startedAt() {
        return this.startedAt;
    }
    /**
     * Volume status
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Volume status
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Sub-operations of the operation
     * 
     */
    @Export(name="subOperations", refs={List.class,VolumeSubOperation.class}, tree="[0,1]")
    private Output<List<VolumeSubOperation>> subOperations;

    /**
     * @return Sub-operations of the operation
     * 
     */
    public Output<List<VolumeSubOperation>> subOperations() {
        return this.subOperations;
    }
    /**
     * Type of the volume **Changing this value recreates the resource.**
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the volume **Changing this value recreates the resource.**
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Volume ID
     * 
     */
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    /**
     * @return Volume ID
     * 
     */
    public Output<String> volumeId() {
        return this.volumeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Volume(java.lang.String name) {
        this(name, VolumeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Volume(java.lang.String name, VolumeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Volume(java.lang.String name, VolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/volume:Volume", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Volume(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/volume:Volume", name, state, makeResourceOptions(options, id), false);
    }

    private static VolumeArgs makeArgs(VolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VolumeArgs.Empty : args;
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
    public static Volume get(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Volume(name, id, state, options);
    }
}
