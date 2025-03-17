// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class LoadBalancerNetworkPrivateNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerNetworkPrivateNetworkArgs Empty = new LoadBalancerNetworkPrivateNetworkArgs();

    /**
     * Private network ID
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return Private network ID
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    /**
     * Subnet ID
     * 
     */
    @Import(name="subnetId", required=true)
    private Output<String> subnetId;

    /**
     * @return Subnet ID
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    private LoadBalancerNetworkPrivateNetworkArgs() {}

    private LoadBalancerNetworkPrivateNetworkArgs(LoadBalancerNetworkPrivateNetworkArgs $) {
        this.id = $.id;
        this.subnetId = $.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerNetworkPrivateNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerNetworkPrivateNetworkArgs $;

        public Builder() {
            $ = new LoadBalancerNetworkPrivateNetworkArgs();
        }

        public Builder(LoadBalancerNetworkPrivateNetworkArgs defaults) {
            $ = new LoadBalancerNetworkPrivateNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Private network ID
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Private network ID
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param subnetId Subnet ID
         * 
         * @return builder
         * 
         */
        public Builder subnetId(Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId Subnet ID
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        public LoadBalancerNetworkPrivateNetworkArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("LoadBalancerNetworkPrivateNetworkArgs", "id");
            }
            if ($.subnetId == null) {
                throw new MissingRequiredPropertyException("LoadBalancerNetworkPrivateNetworkArgs", "subnetId");
            }
            return $;
        }
    }

}
