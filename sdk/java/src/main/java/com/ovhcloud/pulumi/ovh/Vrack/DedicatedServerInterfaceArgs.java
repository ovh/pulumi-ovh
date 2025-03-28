// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class DedicatedServerInterfaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedServerInterfaceArgs Empty = new DedicatedServerInterfaceArgs();

    /**
     * The id of dedicated server network interface.
     * 
     */
    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    /**
     * @return The id of dedicated server network interface.
     * 
     */
    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    /**
     * The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private DedicatedServerInterfaceArgs() {}

    private DedicatedServerInterfaceArgs(DedicatedServerInterfaceArgs $) {
        this.interfaceId = $.interfaceId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedServerInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedServerInterfaceArgs $;

        public Builder() {
            $ = new DedicatedServerInterfaceArgs();
        }

        public Builder(DedicatedServerInterfaceArgs defaults) {
            $ = new DedicatedServerInterfaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param interfaceId The id of dedicated server network interface.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        /**
         * @param interfaceId The id of dedicated server network interface.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        /**
         * @param serviceName The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public DedicatedServerInterfaceArgs build() {
            if ($.interfaceId == null) {
                throw new MissingRequiredPropertyException("DedicatedServerInterfaceArgs", "interfaceId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("DedicatedServerInterfaceArgs", "serviceName");
            }
            return $;
        }
    }

}
