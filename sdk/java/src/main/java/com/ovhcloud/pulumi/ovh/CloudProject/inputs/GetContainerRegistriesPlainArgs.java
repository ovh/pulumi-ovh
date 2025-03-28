// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetContainerRegistriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerRegistriesPlainArgs Empty = new GetContainerRegistriesPlainArgs();

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetContainerRegistriesPlainArgs() {}

    private GetContainerRegistriesPlainArgs(GetContainerRegistriesPlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerRegistriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerRegistriesPlainArgs $;

        public Builder() {
            $ = new GetContainerRegistriesPlainArgs();
        }

        public Builder(GetContainerRegistriesPlainArgs defaults) {
            $ = new GetContainerRegistriesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetContainerRegistriesPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetContainerRegistriesPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
