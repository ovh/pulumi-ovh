// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.InstanceSnapshotArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.InstanceSnapshotState;
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
 * Create and manage snapshots for an instance in a public cloud project.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.ovhcloud.pulumi.ovh.CloudProject.InstanceSnapshot;
 * import com.ovhcloud.pulumi.ovh.CloudProject.InstanceSnapshotArgs;
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
 *         var snapshot = new InstanceSnapshot("snapshot", InstanceSnapshotArgs.builder()
 *             .serviceName("<public cloud project ID>")
 *             .instanceId("<instance ID>")
 *             .name("SnapshotExample")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/instanceSnapshot:InstanceSnapshot")
public class InstanceSnapshot extends com.pulumi.resources.CustomResource {
    /**
     * Image creation date
     * 
     */
    @Export(name="creationDate", refs={String.class}, tree="[0]")
    private Output<String> creationDate;

    /**
     * @return Image creation date
     * 
     */
    public Output<String> creationDate() {
        return this.creationDate;
    }
    /**
     * Image usable only for this type of flavor if not null
     * 
     */
    @Export(name="flavorType", refs={String.class}, tree="[0]")
    private Output<String> flavorType;

    /**
     * @return Image usable only for this type of flavor if not null
     * 
     */
    public Output<String> flavorType() {
        return this.flavorType;
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
     * Minimum disks required to use image
     * 
     */
    @Export(name="minDisk", refs={Double.class}, tree="[0]")
    private Output<Double> minDisk;

    /**
     * @return Minimum disks required to use image
     * 
     */
    public Output<Double> minDisk() {
        return this.minDisk;
    }
    /**
     * Minimum RAM required to use image
     * 
     */
    @Export(name="minRam", refs={Double.class}, tree="[0]")
    private Output<Double> minRam;

    /**
     * @return Minimum RAM required to use image
     * 
     */
    public Output<Double> minRam() {
        return this.minRam;
    }
    /**
     * Snapshot name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Snapshot name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Order plan code
     * 
     */
    @Export(name="planCode", refs={String.class}, tree="[0]")
    private Output<String> planCode;

    /**
     * @return Order plan code
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }
    /**
     * Image region
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Image region
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Image size (in GiB)
     * 
     */
    @Export(name="size", refs={Double.class}, tree="[0]")
    private Output<Double> size;

    /**
     * @return Image size (in GiB)
     * 
     */
    public Output<Double> size() {
        return this.size;
    }
    /**
     * Image status
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Image status
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Tags about the image
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return Tags about the image
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * Image type
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Image type
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * User to connect with
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output<String> user;

    /**
     * @return User to connect with
     * 
     */
    public Output<String> user() {
        return this.user;
    }
    /**
     * Image visibility
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Image visibility
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceSnapshot(java.lang.String name) {
        this(name, InstanceSnapshotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceSnapshot(java.lang.String name, InstanceSnapshotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceSnapshot(java.lang.String name, InstanceSnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/instanceSnapshot:InstanceSnapshot", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InstanceSnapshot(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/instanceSnapshot:InstanceSnapshot", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceSnapshotArgs makeArgs(InstanceSnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceSnapshotArgs.Empty : args;
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
    public static InstanceSnapshot get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceSnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceSnapshot(name, id, state, options);
    }
}
