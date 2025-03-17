// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.VolumeBackupArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.VolumeBackupState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manage backups for the given volume in a public cloud project.
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
 * import com.pulumi.ovh.CloudProject.VolumeBackup;
 * import com.pulumi.ovh.CloudProject.VolumeBackupArgs;
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
 *         var backup = new VolumeBackup("backup", VolumeBackupArgs.builder()
 *             .regionName("GRA9")
 *             .serviceName("<public cloud project ID>")
 *             .volumeId("<volume ID>")
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
 * A volume backup in a public cloud project can be imported using the `service_name`, `region_name` and `id` attributes.
 * 
 * Using the following configuration:
 * 
 * hcl
 * 
 * import {
 * 
 *   id = &#34;&lt;service_name&gt;/&lt;region_name&gt;/&lt;id&gt;&#34;
 * 
 *   to = ovh_cloud_project_volume_backup.backup
 * 
 * }
 * 
 * You can then run:
 * 
 * bash
 * 
 * $ pulumi preview -generate-config-out=backup.tf
 * 
 * $ pulumi up
 * 
 * The file `backup.tf` will then contain the imported resource&#39;s configuration, that can be copied next to the `import` block above.
 * 
 * See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 * 
 */
@ResourceType(type="ovh:CloudProject/volumeBackup:VolumeBackup")
public class VolumeBackup extends com.pulumi.resources.CustomResource {
    /**
     * Creation date of the backup
     * 
     */
    @Export(name="creationDate", refs={String.class}, tree="[0]")
    private Output<String> creationDate;

    /**
     * @return Creation date of the backup
     * 
     */
    public Output<String> creationDate() {
        return this.creationDate;
    }
    /**
     * name of the backup
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return name of the backup
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Volume backup region
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Volume backup region
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Region name
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return Region name
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
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
     * Size of the backup in GiB
     * 
     */
    @Export(name="size", refs={Double.class}, tree="[0]")
    private Output<Double> size;

    /**
     * @return Size of the backup in GiB
     * 
     */
    public Output<Double> size() {
        return this.size;
    }
    /**
     * Staus of the backup
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Staus of the backup
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * ID of the volume to backup
     * 
     */
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    /**
     * @return ID of the volume to backup
     * 
     */
    public Output<String> volumeId() {
        return this.volumeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeBackup(java.lang.String name) {
        this(name, VolumeBackupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeBackup(java.lang.String name, VolumeBackupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeBackup(java.lang.String name, VolumeBackupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/volumeBackup:VolumeBackup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VolumeBackup(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeBackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/volumeBackup:VolumeBackup", name, state, makeResourceOptions(options, id), false);
    }

    private static VolumeBackupArgs makeArgs(VolumeBackupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VolumeBackupArgs.Empty : args;
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
    public static VolumeBackup get(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeBackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeBackup(name, id, state, options);
    }
}
