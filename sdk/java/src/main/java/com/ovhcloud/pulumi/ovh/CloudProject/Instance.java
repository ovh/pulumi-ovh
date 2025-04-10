// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.InstanceArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.InstanceState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceAddress;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceAttachedVolume;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceAutoBackup;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceBootFrom;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceFlavor;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceGroup;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetwork;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceSshKey;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceSshKeyCreate;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * **This resource uses a Beta API** Creates an instance associated with a public cloud project.
 * 
 * ## Example Usage
 * 
 * Create a instance.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.ovhcloud.pulumi.ovh.CloudProject.Instance;
 * import com.ovhcloud.pulumi.ovh.CloudProject.InstanceArgs;
 * import com.pulumi.ovh.CloudProject.inputs.InstanceBootFromArgs;
 * import com.pulumi.ovh.CloudProject.inputs.InstanceFlavorArgs;
 * import com.pulumi.ovh.CloudProject.inputs.InstanceNetworkArgs;
 * import com.pulumi.ovh.CloudProject.inputs.InstanceSshKeyArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .billingPeriod("hourly")
 *             .bootFrom(InstanceBootFromArgs.builder()
 *                 .imageId("UUID")
 *                 .build())
 *             .flavor(InstanceFlavorArgs.builder()
 *                 .flavorId("UUID")
 *                 .build())
 *             .network(InstanceNetworkArgs.builder()
 *                 .public_(true)
 *                 .build())
 *             .region("RRRR")
 *             .serviceName("XXX")
 *             .sshKey(InstanceSshKeyArgs.builder()
 *                 .name("sshname")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Instance IP addresses
     * 
     */
    @Export(name="addresses", refs={List.class,InstanceAddress.class}, tree="[0,1]")
    private Output<List<InstanceAddress>> addresses;

    /**
     * @return Instance IP addresses
     * 
     */
    public Output<List<InstanceAddress>> addresses() {
        return this.addresses;
    }
    /**
     * Volumes attached to the instance
     * 
     */
    @Export(name="attachedVolumes", refs={List.class,InstanceAttachedVolume.class}, tree="[0,1]")
    private Output<List<InstanceAttachedVolume>> attachedVolumes;

    /**
     * @return Volumes attached to the instance
     * 
     */
    public Output<List<InstanceAttachedVolume>> attachedVolumes() {
        return this.attachedVolumes;
    }
    /**
     * Create an autobackup workflow after instance start up.
     * 
     */
    @Export(name="autoBackup", refs={InstanceAutoBackup.class}, tree="[0]")
    private Output</* @Nullable */ InstanceAutoBackup> autoBackup;

    /**
     * @return Create an autobackup workflow after instance start up.
     * 
     */
    public Output<Optional<InstanceAutoBackup>> autoBackup() {
        return Codegen.optional(this.autoBackup);
    }
    /**
     * The availability zone where the instance will be created
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The availability zone where the instance will be created
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Billing period - hourly or monthly
     * 
     */
    @Export(name="billingPeriod", refs={String.class}, tree="[0]")
    private Output<String> billingPeriod;

    /**
     * @return Billing period - hourly or monthly
     * 
     */
    public Output<String> billingPeriod() {
        return this.billingPeriod;
    }
    /**
     * Boot the instance from an image or a volume
     * 
     */
    @Export(name="bootFrom", refs={InstanceBootFrom.class}, tree="[0]")
    private Output<InstanceBootFrom> bootFrom;

    /**
     * @return Boot the instance from an image or a volume
     * 
     */
    public Output<InstanceBootFrom> bootFrom() {
        return this.bootFrom;
    }
    /**
     * Create multiple instances
     * 
     */
    @Export(name="bulk", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> bulk;

    /**
     * @return Create multiple instances
     * 
     */
    public Output<Optional<Integer>> bulk() {
        return Codegen.optional(this.bulk);
    }
    /**
     * Flavor information
     * 
     */
    @Export(name="flavor", refs={InstanceFlavor.class}, tree="[0]")
    private Output<InstanceFlavor> flavor;

    /**
     * @return Flavor information
     * 
     */
    public Output<InstanceFlavor> flavor() {
        return this.flavor;
    }
    /**
     * Flavor id
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return Flavor id
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * Flavor name
     * 
     */
    @Export(name="flavorName", refs={String.class}, tree="[0]")
    private Output<String> flavorName;

    /**
     * @return Flavor name
     * 
     */
    public Output<String> flavorName() {
        return this.flavorName;
    }
    /**
     * Start instance in group
     * 
     */
    @Export(name="group", refs={InstanceGroup.class}, tree="[0]")
    private Output</* @Nullable */ InstanceGroup> group;

    /**
     * @return Start instance in group
     * 
     */
    public Output<Optional<InstanceGroup>> group() {
        return Codegen.optional(this.group);
    }
    /**
     * Image id
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return Image id
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * Instance name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Instance name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Create network interfaces
     * 
     */
    @Export(name="network", refs={InstanceNetwork.class}, tree="[0]")
    private Output<InstanceNetwork> network;

    /**
     * @return Create network interfaces
     * 
     */
    public Output<InstanceNetwork> network() {
        return this.network;
    }
    /**
     * Instance region
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Instance region
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Existing SSH Keypair
     * 
     */
    @Export(name="sshKey", refs={InstanceSshKey.class}, tree="[0]")
    private Output</* @Nullable */ InstanceSshKey> sshKey;

    /**
     * @return Existing SSH Keypair
     * 
     */
    public Output<Optional<InstanceSshKey>> sshKey() {
        return Codegen.optional(this.sshKey);
    }
    /**
     * Add existing SSH Key pair into your Public Cloud project and link it to the instance
     * 
     */
    @Export(name="sshKeyCreate", refs={InstanceSshKeyCreate.class}, tree="[0]")
    private Output</* @Nullable */ InstanceSshKeyCreate> sshKeyCreate;

    /**
     * @return Add existing SSH Key pair into your Public Cloud project and link it to the instance
     * 
     */
    public Output<Optional<InstanceSshKeyCreate>> sshKeyCreate() {
        return Codegen.optional(this.sshKeyCreate);
    }
    /**
     * Instance task state
     * 
     */
    @Export(name="taskState", refs={String.class}, tree="[0]")
    private Output<String> taskState;

    /**
     * @return Instance task state
     * 
     */
    public Output<String> taskState() {
        return this.taskState;
    }
    /**
     * Configuration information or scripts to use upon launch
     * 
     */
    @Export(name="userData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userData;

    /**
     * @return Configuration information or scripts to use upon launch
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(java.lang.String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(java.lang.String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(java.lang.String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/instance:Instance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Instance(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/instance:Instance", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceArgs makeArgs(InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceArgs.Empty : args;
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
    public static Instance get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
