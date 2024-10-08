// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetLogsClustersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLogsClustersArgs Empty = new GetLogsClustersArgs();

    /**
     * The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetLogsClustersArgs() {}

    private GetLogsClustersArgs(GetLogsClustersArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogsClustersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogsClustersArgs $;

        public Builder() {
            $ = new GetLogsClustersArgs();
        }

        public Builder(GetLogsClustersArgs defaults) {
            $ = new GetLogsClustersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetLogsClustersArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetLogsClustersArgs", "serviceName");
            }
            return $;
        }
    }

}
