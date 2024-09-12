// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.ovhcloud.pulumi.ovh.Dedicated.inputs.ServerNetworkingInterfaceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerNetworkingState extends com.pulumi.resources.ResourceArgs {

    public static final ServerNetworkingState Empty = new ServerNetworkingState();

    /**
     * Operation description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Operation description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Interface or interfaces aggregation.
     * 
     */
    @Import(name="interfaces")
    private @Nullable Output<List<ServerNetworkingInterfaceArgs>> interfaces;

    /**
     * @return Interface or interfaces aggregation.
     * 
     */
    public Optional<Output<List<ServerNetworkingInterfaceArgs>>> interfaces() {
        return Optional.ofNullable(this.interfaces);
    }

    /**
     * The internal name of your dedicated server.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your dedicated server.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Operation status
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Operation status
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private ServerNetworkingState() {}

    private ServerNetworkingState(ServerNetworkingState $) {
        this.description = $.description;
        this.interfaces = $.interfaces;
        this.serviceName = $.serviceName;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerNetworkingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerNetworkingState $;

        public Builder() {
            $ = new ServerNetworkingState();
        }

        public Builder(ServerNetworkingState defaults) {
            $ = new ServerNetworkingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Operation description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Operation description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(@Nullable Output<List<ServerNetworkingInterfaceArgs>> interfaces) {
            $.interfaces = interfaces;
            return this;
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(List<ServerNetworkingInterfaceArgs> interfaces) {
            return interfaces(Output.of(interfaces));
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(ServerNetworkingInterfaceArgs... interfaces) {
            return interfaces(List.of(interfaces));
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param status Operation status
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Operation status
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public ServerNetworkingState build() {
            return $;
        }
    }

}