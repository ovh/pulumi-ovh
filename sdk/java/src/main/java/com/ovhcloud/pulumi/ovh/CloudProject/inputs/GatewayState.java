// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.GatewayExternalInformationArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.GatewayInterfaceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GatewayState extends com.pulumi.resources.ResourceArgs {

    public static final GatewayState Empty = new GatewayState();

    /**
     * List of External Information of the gateway.
     * 
     */
    @Import(name="externalInformations")
    private @Nullable Output<List<GatewayExternalInformationArgs>> externalInformations;

    /**
     * @return List of External Information of the gateway.
     * 
     */
    public Optional<Output<List<GatewayExternalInformationArgs>>> externalInformations() {
        return Optional.ofNullable(this.externalInformations);
    }

    /**
     * Interfaces list of the gateway.
     * 
     */
    @Import(name="interfaces")
    private @Nullable Output<List<GatewayInterfaceArgs>> interfaces;

    /**
     * @return Interfaces list of the gateway.
     * 
     */
    public Optional<Output<List<GatewayInterfaceArgs>>> interfaces() {
        return Optional.ofNullable(this.interfaces);
    }

    /**
     * Model of the gateway.
     * 
     */
    @Import(name="model")
    private @Nullable Output<String> model;

    /**
     * @return Model of the gateway.
     * 
     */
    public Optional<Output<String>> model() {
        return Optional.ofNullable(this.model);
    }

    /**
     * Name of the gateway.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the gateway.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * ID of the private network.
     * 
     */
    @Import(name="networkId")
    private @Nullable Output<String> networkId;

    /**
     * @return ID of the private network.
     * 
     */
    public Optional<Output<String>> networkId() {
        return Optional.ofNullable(this.networkId);
    }

    /**
     * Region of the gateway.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Region of the gateway.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * ID of the private network.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return ID of the private network.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Status of the gateway.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the gateway.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * ID of the subnet.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return ID of the subnet.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    private GatewayState() {}

    private GatewayState(GatewayState $) {
        this.externalInformations = $.externalInformations;
        this.interfaces = $.interfaces;
        this.model = $.model;
        this.name = $.name;
        this.networkId = $.networkId;
        this.region = $.region;
        this.serviceName = $.serviceName;
        this.status = $.status;
        this.subnetId = $.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GatewayState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GatewayState $;

        public Builder() {
            $ = new GatewayState();
        }

        public Builder(GatewayState defaults) {
            $ = new GatewayState(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalInformations List of External Information of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder externalInformations(@Nullable Output<List<GatewayExternalInformationArgs>> externalInformations) {
            $.externalInformations = externalInformations;
            return this;
        }

        /**
         * @param externalInformations List of External Information of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder externalInformations(List<GatewayExternalInformationArgs> externalInformations) {
            return externalInformations(Output.of(externalInformations));
        }

        /**
         * @param externalInformations List of External Information of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder externalInformations(GatewayExternalInformationArgs... externalInformations) {
            return externalInformations(List.of(externalInformations));
        }

        /**
         * @param interfaces Interfaces list of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(@Nullable Output<List<GatewayInterfaceArgs>> interfaces) {
            $.interfaces = interfaces;
            return this;
        }

        /**
         * @param interfaces Interfaces list of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(List<GatewayInterfaceArgs> interfaces) {
            return interfaces(Output.of(interfaces));
        }

        /**
         * @param interfaces Interfaces list of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(GatewayInterfaceArgs... interfaces) {
            return interfaces(List.of(interfaces));
        }

        /**
         * @param model Model of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder model(@Nullable Output<String> model) {
            $.model = model;
            return this;
        }

        /**
         * @param model Model of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder model(String model) {
            return model(Output.of(model));
        }

        /**
         * @param name Name of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkId ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder networkId(@Nullable Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param region Region of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Region of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceName ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName ID of the private network.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param status Status of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param subnetId ID of the subnet.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId ID of the subnet.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        public GatewayState build() {
            return $;
        }
    }

}
