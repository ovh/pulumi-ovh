// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class DedicatedServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedServerArgs Empty = new DedicatedServerArgs();

    /**
     * The id of the dedicated server.
     * 
     */
    @Import(name="serverId", required=true)
    private Output<String> serverId;

    /**
     * @return The id of the dedicated server.
     * 
     */
    public Output<String> serverId() {
        return this.serverId;
    }

    /**
     * The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private DedicatedServerArgs() {}

    private DedicatedServerArgs(DedicatedServerArgs $) {
        this.serverId = $.serverId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedServerArgs $;

        public Builder() {
            $ = new DedicatedServerArgs();
        }

        public Builder(DedicatedServerArgs defaults) {
            $ = new DedicatedServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serverId The id of the dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serverId(Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        /**
         * @param serverId The id of the dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        /**
         * @param serviceName The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public DedicatedServerArgs build() {
            if ($.serverId == null) {
                throw new MissingRequiredPropertyException("DedicatedServerArgs", "serverId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("DedicatedServerArgs", "serviceName");
            }
            return $;
        }
    }

}
