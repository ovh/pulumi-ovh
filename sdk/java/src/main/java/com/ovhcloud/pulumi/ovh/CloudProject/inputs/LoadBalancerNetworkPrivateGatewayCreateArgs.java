// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerNetworkPrivateGatewayCreateArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerNetworkPrivateGatewayCreateArgs Empty = new LoadBalancerNetworkPrivateGatewayCreateArgs();

    /**
     * Model of the gateway
     * 
     */
    @Import(name="model")
    private @Nullable Output<String> model;

    /**
     * @return Model of the gateway
     * 
     */
    public Optional<Output<String>> model() {
        return Optional.ofNullable(this.model);
    }

    /**
     * Name of the gateway
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the gateway
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private LoadBalancerNetworkPrivateGatewayCreateArgs() {}

    private LoadBalancerNetworkPrivateGatewayCreateArgs(LoadBalancerNetworkPrivateGatewayCreateArgs $) {
        this.model = $.model;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerNetworkPrivateGatewayCreateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerNetworkPrivateGatewayCreateArgs $;

        public Builder() {
            $ = new LoadBalancerNetworkPrivateGatewayCreateArgs();
        }

        public Builder(LoadBalancerNetworkPrivateGatewayCreateArgs defaults) {
            $ = new LoadBalancerNetworkPrivateGatewayCreateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param model Model of the gateway
         * 
         * @return builder
         * 
         */
        public Builder model(@Nullable Output<String> model) {
            $.model = model;
            return this;
        }

        /**
         * @param model Model of the gateway
         * 
         * @return builder
         * 
         */
        public Builder model(String model) {
            return model(Output.of(model));
        }

        /**
         * @param name Name of the gateway
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the gateway
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public LoadBalancerNetworkPrivateGatewayCreateArgs build() {
            return $;
        }
    }

}
