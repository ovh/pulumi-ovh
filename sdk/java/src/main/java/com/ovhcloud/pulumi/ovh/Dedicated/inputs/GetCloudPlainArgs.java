// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetCloudPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCloudPlainArgs Empty = new GetCloudPlainArgs();

    /**
     * Domain of the service
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return Domain of the service
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetCloudPlainArgs() {}

    private GetCloudPlainArgs(GetCloudPlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCloudPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCloudPlainArgs $;

        public Builder() {
            $ = new GetCloudPlainArgs();
        }

        public Builder(GetCloudPlainArgs defaults) {
            $ = new GetCloudPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName Domain of the service
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetCloudPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetCloudPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
