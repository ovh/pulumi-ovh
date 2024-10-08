// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Iam;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceGroupArgs Empty = new ResourceGroupArgs();

    /**
     * Name of the resource group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     * 
     */
    @Import(name="resources")
    private @Nullable Output<List<String>> resources;

    /**
     * @return Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     * 
     */
    public Optional<Output<List<String>>> resources() {
        return Optional.ofNullable(this.resources);
    }

    private ResourceGroupArgs() {}

    private ResourceGroupArgs(ResourceGroupArgs $) {
        this.name = $.name;
        this.resources = $.resources;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceGroupArgs $;

        public Builder() {
            $ = new ResourceGroupArgs();
        }

        public Builder(ResourceGroupArgs defaults) {
            $ = new ResourceGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the resource group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(@Nullable Output<List<String>> resources) {
            $.resources = resources;
            return this;
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(List<String> resources) {
            return resources(Output.of(resources));
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(String... resources) {
            return resources(List.of(resources));
        }

        public ResourceGroupArgs build() {
            return $;
        }
    }

}
