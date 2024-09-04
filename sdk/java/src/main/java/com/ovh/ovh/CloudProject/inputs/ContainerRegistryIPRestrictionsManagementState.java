// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerRegistryIPRestrictionsManagementState extends com.pulumi.resources.ResourceArgs {

    public static final ContainerRegistryIPRestrictionsManagementState Empty = new ContainerRegistryIPRestrictionsManagementState();

    /**
     * IP restrictions applied on Harbor UI and API.
     * 
     */
    @Import(name="ipRestrictions")
    private @Nullable Output<List<Map<String,String>>> ipRestrictions;

    /**
     * @return IP restrictions applied on Harbor UI and API.
     * 
     */
    public Optional<Output<List<Map<String,String>>>> ipRestrictions() {
        return Optional.ofNullable(this.ipRestrictions);
    }

    /**
     * The id of the Managed Private Registry.
     * 
     */
    @Import(name="registryId")
    private @Nullable Output<String> registryId;

    /**
     * @return The id of the Managed Private Registry.
     * 
     */
    public Optional<Output<String>> registryId() {
        return Optional.ofNullable(this.registryId);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private ContainerRegistryIPRestrictionsManagementState() {}

    private ContainerRegistryIPRestrictionsManagementState(ContainerRegistryIPRestrictionsManagementState $) {
        this.ipRestrictions = $.ipRestrictions;
        this.registryId = $.registryId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerRegistryIPRestrictionsManagementState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerRegistryIPRestrictionsManagementState $;

        public Builder() {
            $ = new ContainerRegistryIPRestrictionsManagementState();
        }

        public Builder(ContainerRegistryIPRestrictionsManagementState defaults) {
            $ = new ContainerRegistryIPRestrictionsManagementState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipRestrictions IP restrictions applied on Harbor UI and API.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(@Nullable Output<List<Map<String,String>>> ipRestrictions) {
            $.ipRestrictions = ipRestrictions;
            return this;
        }

        /**
         * @param ipRestrictions IP restrictions applied on Harbor UI and API.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(List<Map<String,String>> ipRestrictions) {
            return ipRestrictions(Output.of(ipRestrictions));
        }

        /**
         * @param ipRestrictions IP restrictions applied on Harbor UI and API.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(Map<String,String>... ipRestrictions) {
            return ipRestrictions(List.of(ipRestrictions));
        }

        /**
         * @param registryId The id of the Managed Private Registry.
         * 
         * @return builder
         * 
         */
        public Builder registryId(@Nullable Output<String> registryId) {
            $.registryId = registryId;
            return this;
        }

        /**
         * @param registryId The id of the Managed Private Registry.
         * 
         * @return builder
         * 
         */
        public Builder registryId(String registryId) {
            return registryId(Output.of(registryId));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public ContainerRegistryIPRestrictionsManagementState build() {
            return $;
        }
    }

}
