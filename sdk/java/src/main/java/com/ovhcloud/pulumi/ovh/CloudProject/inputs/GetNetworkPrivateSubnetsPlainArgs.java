// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetNetworkPrivateSubnetsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkPrivateSubnetsPlainArgs Empty = new GetNetworkPrivateSubnetsPlainArgs();

    /**
     * ID of the network
     * 
     */
    @Import(name="networkId", required=true)
    private String networkId;

    /**
     * @return ID of the network
     * 
     */
    public String networkId() {
        return this.networkId;
    }

    /**
     * The ID of the public cloud project.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The ID of the public cloud project.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetNetworkPrivateSubnetsPlainArgs() {}

    private GetNetworkPrivateSubnetsPlainArgs(GetNetworkPrivateSubnetsPlainArgs $) {
        this.networkId = $.networkId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkPrivateSubnetsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkPrivateSubnetsPlainArgs $;

        public Builder() {
            $ = new GetNetworkPrivateSubnetsPlainArgs();
        }

        public Builder(GetNetworkPrivateSubnetsPlainArgs defaults) {
            $ = new GetNetworkPrivateSubnetsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param networkId ID of the network
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetNetworkPrivateSubnetsPlainArgs build() {
            if ($.networkId == null) {
                throw new MissingRequiredPropertyException("GetNetworkPrivateSubnetsPlainArgs", "networkId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetNetworkPrivateSubnetsPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
