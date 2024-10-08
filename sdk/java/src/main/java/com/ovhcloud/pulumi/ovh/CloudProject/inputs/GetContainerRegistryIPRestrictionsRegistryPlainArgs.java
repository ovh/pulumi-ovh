// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetContainerRegistryIPRestrictionsRegistryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerRegistryIPRestrictionsRegistryPlainArgs Empty = new GetContainerRegistryIPRestrictionsRegistryPlainArgs();

    /**
     * The id of the Managed Private Registry.
     * 
     */
    @Import(name="registryId", required=true)
    private String registryId;

    /**
     * @return The id of the Managed Private Registry.
     * 
     */
    public String registryId() {
        return this.registryId;
    }

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

    private GetContainerRegistryIPRestrictionsRegistryPlainArgs() {}

    private GetContainerRegistryIPRestrictionsRegistryPlainArgs(GetContainerRegistryIPRestrictionsRegistryPlainArgs $) {
        this.registryId = $.registryId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerRegistryIPRestrictionsRegistryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerRegistryIPRestrictionsRegistryPlainArgs $;

        public Builder() {
            $ = new GetContainerRegistryIPRestrictionsRegistryPlainArgs();
        }

        public Builder(GetContainerRegistryIPRestrictionsRegistryPlainArgs defaults) {
            $ = new GetContainerRegistryIPRestrictionsRegistryPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param registryId The id of the Managed Private Registry.
         * 
         * @return builder
         * 
         */
        public Builder registryId(String registryId) {
            $.registryId = registryId;
            return this;
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

        public GetContainerRegistryIPRestrictionsRegistryPlainArgs build() {
            if ($.registryId == null) {
                throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryPlainArgs", "registryId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
