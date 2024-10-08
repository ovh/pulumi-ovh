// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLogsClusterPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLogsClusterPlainArgs Empty = new GetLogsClusterPlainArgs();

    /**
     * Cluster ID. If not provided, the default cluster_id is returned
     * 
     */
    @Import(name="clusterId")
    private @Nullable String clusterId;

    /**
     * @return Cluster ID. If not provided, the default cluster_id is returned
     * 
     */
    public Optional<String> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetLogsClusterPlainArgs() {}

    private GetLogsClusterPlainArgs(GetLogsClusterPlainArgs $) {
        this.clusterId = $.clusterId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogsClusterPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogsClusterPlainArgs $;

        public Builder() {
            $ = new GetLogsClusterPlainArgs();
        }

        public Builder(GetLogsClusterPlainArgs defaults) {
            $ = new GetLogsClusterPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID. If not provided, the default cluster_id is returned
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetLogsClusterPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetLogsClusterPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
