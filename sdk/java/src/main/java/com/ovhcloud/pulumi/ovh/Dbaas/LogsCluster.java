// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas;

import com.ovhcloud.pulumi.ovh.Dbaas.LogsClusterArgs;
import com.ovhcloud.pulumi.ovh.Dbaas.inputs.LogsClusterState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsCluster;
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsClusterArgs;
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
 *         var ldp = new LogsCluster("ldp", LogsClusterArgs.builder()
 *             .serviceName("ldp-xx-xxxxx")
 *             .clusterId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .archiveAllowedNetworks("10.0.0.0/16")
 *             .directInputAllowedNetworks("10.0.0.0/16")
 *             .queryAllowedNetworks("10.0.0.0/16")
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
 * OVHcloud DBaaS Log Data Platform clusters can be imported using the `service_name` and `cluster_id` of the cluster, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Dbaas/logsCluster:LogsCluster ldp service_name/cluster_id
 * ```
 * 
 */
@ResourceType(type="ovh:Dbaas/logsCluster:LogsCluster")
public class LogsCluster extends com.pulumi.resources.CustomResource {
    /**
     * List of IP blocks
     * 
     */
    @Export(name="archiveAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> archiveAllowedNetworks;

    /**
     * @return List of IP blocks
     * 
     */
    public Output<Optional<List<String>>> archiveAllowedNetworks() {
        return Codegen.optional(this.archiveAllowedNetworks);
    }
    /**
     * Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterId;

    /**
     * @return Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    public Output<Optional<String>> clusterId() {
        return Codegen.optional(this.clusterId);
    }
    /**
     * type of cluster (DEDICATED, PRO or TRIAL)
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output<String> clusterType;

    /**
     * @return type of cluster (DEDICATED, PRO or TRIAL)
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * PEM for dedicated inputs
     * 
     */
    @Export(name="dedicatedInputPem", refs={String.class}, tree="[0]")
    private Output<String> dedicatedInputPem;

    /**
     * @return PEM for dedicated inputs
     * 
     */
    public Output<String> dedicatedInputPem() {
        return this.dedicatedInputPem;
    }
    /**
     * List of IP blocks
     * 
     */
    @Export(name="directInputAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> directInputAllowedNetworks;

    /**
     * @return List of IP blocks
     * 
     */
    public Output<Optional<List<String>>> directInputAllowedNetworks() {
        return Codegen.optional(this.directInputAllowedNetworks);
    }
    /**
     * PEM for direct inputs
     * 
     */
    @Export(name="directInputPem", refs={String.class}, tree="[0]")
    private Output<String> directInputPem;

    /**
     * @return PEM for direct inputs
     * 
     */
    public Output<String> directInputPem() {
        return this.directInputPem;
    }
    /**
     * cluster hostname hosting tenant
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output<String> hostname;

    /**
     * @return cluster hostname hosting tenant
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * Initial allowed networks for ARCHIVE flow type
     * 
     */
    @Export(name="initialArchiveAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> initialArchiveAllowedNetworks;

    /**
     * @return Initial allowed networks for ARCHIVE flow type
     * 
     */
    public Output<List<String>> initialArchiveAllowedNetworks() {
        return this.initialArchiveAllowedNetworks;
    }
    /**
     * Initial allowed networks for DIRECT_INPUT flow type
     * 
     */
    @Export(name="initialDirectInputAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> initialDirectInputAllowedNetworks;

    /**
     * @return Initial allowed networks for DIRECT_INPUT flow type
     * 
     */
    public Output<List<String>> initialDirectInputAllowedNetworks() {
        return this.initialDirectInputAllowedNetworks;
    }
    /**
     * Initial allowed networks for QUERY flow type
     * 
     */
    @Export(name="initialQueryAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> initialQueryAllowedNetworks;

    /**
     * @return Initial allowed networks for QUERY flow type
     * 
     */
    public Output<List<String>> initialQueryAllowedNetworks() {
        return this.initialQueryAllowedNetworks;
    }
    /**
     * true if all content generated by given service will be placed on this cluster
     * 
     */
    @Export(name="isDefault", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isDefault;

    /**
     * @return true if all content generated by given service will be placed on this cluster
     * 
     */
    public Output<Boolean> isDefault() {
        return this.isDefault;
    }
    /**
     * true if given service can perform advanced operations on cluster
     * 
     */
    @Export(name="isUnlocked", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isUnlocked;

    /**
     * @return true if given service can perform advanced operations on cluster
     * 
     */
    public Output<Boolean> isUnlocked() {
        return this.isUnlocked;
    }
    /**
     * List of IP blocks
     * 
     */
    @Export(name="queryAllowedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> queryAllowedNetworks;

    /**
     * @return List of IP blocks
     * 
     */
    public Output<Optional<List<String>>> queryAllowedNetworks() {
        return Codegen.optional(this.queryAllowedNetworks);
    }
    /**
     * datacenter localization
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return datacenter localization
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogsCluster(java.lang.String name) {
        this(name, LogsClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogsCluster(java.lang.String name, LogsClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogsCluster(java.lang.String name, LogsClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsCluster:LogsCluster", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LogsCluster(java.lang.String name, Output<java.lang.String> id, @Nullable LogsClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsCluster:LogsCluster", name, state, makeResourceOptions(options, id), false);
    }

    private static LogsClusterArgs makeArgs(LogsClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LogsClusterArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "dedicatedInputPem",
                "directInputPem",
                "initialArchiveAllowedNetworks",
                "initialDirectInputAllowedNetworks",
                "initialQueryAllowedNetworks"
            ))
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
    public static LogsCluster get(java.lang.String name, Output<java.lang.String> id, @Nullable LogsClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogsCluster(name, id, state, options);
    }
}
