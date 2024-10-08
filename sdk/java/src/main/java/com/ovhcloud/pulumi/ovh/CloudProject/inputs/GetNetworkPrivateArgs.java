// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetNetworkPrivateArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkPrivateArgs Empty = new GetNetworkPrivateArgs();

    /**
     * ID of the network
     * 
     */
    @Import(name="networkId", required=true)
    private Output<String> networkId;

    /**
     * @return ID of the network
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }

    /**
     * The ID of the public cloud project.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The ID of the public cloud project.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetNetworkPrivateArgs() {}

    private GetNetworkPrivateArgs(GetNetworkPrivateArgs $) {
        this.networkId = $.networkId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkPrivateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkPrivateArgs $;

        public Builder() {
            $ = new GetNetworkPrivateArgs();
        }

        public Builder(GetNetworkPrivateArgs defaults) {
            $ = new GetNetworkPrivateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param networkId ID of the network
         * 
         * @return builder
         * 
         */
        public Builder networkId(Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId ID of the network
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param serviceName The ID of the public cloud project.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetNetworkPrivateArgs build() {
            if ($.networkId == null) {
                throw new MissingRequiredPropertyException("GetNetworkPrivateArgs", "networkId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetNetworkPrivateArgs", "serviceName");
            }
            return $;
        }
    }

}
