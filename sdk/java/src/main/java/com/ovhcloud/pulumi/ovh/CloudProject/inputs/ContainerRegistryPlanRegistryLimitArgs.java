// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerRegistryPlanRegistryLimitArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerRegistryPlanRegistryLimitArgs Empty = new ContainerRegistryPlanRegistryLimitArgs();

    /**
     * Docker image storage limits in bytes
     * 
     */
    @Import(name="imageStorage")
    private @Nullable Output<Integer> imageStorage;

    /**
     * @return Docker image storage limits in bytes
     * 
     */
    public Optional<Output<Integer>> imageStorage() {
        return Optional.ofNullable(this.imageStorage);
    }

    /**
     * Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    @Import(name="parallelRequest")
    private @Nullable Output<Integer> parallelRequest;

    /**
     * @return Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    public Optional<Output<Integer>> parallelRequest() {
        return Optional.ofNullable(this.parallelRequest);
    }

    private ContainerRegistryPlanRegistryLimitArgs() {}

    private ContainerRegistryPlanRegistryLimitArgs(ContainerRegistryPlanRegistryLimitArgs $) {
        this.imageStorage = $.imageStorage;
        this.parallelRequest = $.parallelRequest;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerRegistryPlanRegistryLimitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerRegistryPlanRegistryLimitArgs $;

        public Builder() {
            $ = new ContainerRegistryPlanRegistryLimitArgs();
        }

        public Builder(ContainerRegistryPlanRegistryLimitArgs defaults) {
            $ = new ContainerRegistryPlanRegistryLimitArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param imageStorage Docker image storage limits in bytes
         * 
         * @return builder
         * 
         */
        public Builder imageStorage(@Nullable Output<Integer> imageStorage) {
            $.imageStorage = imageStorage;
            return this;
        }

        /**
         * @param imageStorage Docker image storage limits in bytes
         * 
         * @return builder
         * 
         */
        public Builder imageStorage(Integer imageStorage) {
            return imageStorage(Output.of(imageStorage));
        }

        /**
         * @param parallelRequest Parallel requests on Docker image API (/v2 Docker registry API)
         * 
         * @return builder
         * 
         */
        public Builder parallelRequest(@Nullable Output<Integer> parallelRequest) {
            $.parallelRequest = parallelRequest;
            return this;
        }

        /**
         * @param parallelRequest Parallel requests on Docker image API (/v2 Docker registry API)
         * 
         * @return builder
         * 
         */
        public Builder parallelRequest(Integer parallelRequest) {
            return parallelRequest(Output.of(parallelRequest));
        }

        public ContainerRegistryPlanRegistryLimitArgs build() {
            return $;
        }
    }

}
