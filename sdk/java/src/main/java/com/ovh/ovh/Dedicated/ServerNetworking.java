// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dedicated;

import com.ovh.ovh.Dedicated.ServerNetworkingArgs;
import com.ovh.ovh.Dedicated.inputs.ServerNetworkingState;
import com.ovh.ovh.Dedicated.outputs.ServerNetworkingInterface;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

@ResourceType(type="ovh:Dedicated/serverNetworking:ServerNetworking")
public class ServerNetworking extends com.pulumi.resources.CustomResource {
    /**
     * Operation description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Operation description
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Interface or interfaces aggregation.
     * 
     */
    @Export(name="interfaces", refs={List.class,ServerNetworkingInterface.class}, tree="[0,1]")
    private Output<List<ServerNetworkingInterface>> interfaces;

    /**
     * @return Interface or interfaces aggregation.
     * 
     */
    public Output<List<ServerNetworkingInterface>> interfaces() {
        return this.interfaces;
    }
    /**
     * The internal name of your dedicated server.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your dedicated server.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Operation status
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Operation status
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerNetworking(java.lang.String name) {
        this(name, ServerNetworkingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerNetworking(java.lang.String name, ServerNetworkingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerNetworking(java.lang.String name, ServerNetworkingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/serverNetworking:ServerNetworking", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServerNetworking(java.lang.String name, Output<java.lang.String> id, @Nullable ServerNetworkingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/serverNetworking:ServerNetworking", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerNetworkingArgs makeArgs(ServerNetworkingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerNetworkingArgs.Empty : args;
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
    public static ServerNetworking get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerNetworkingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerNetworking(name, id, state, options);
    }
}
