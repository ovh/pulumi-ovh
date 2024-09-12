// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerRegistryPlanRegistryLimit {
    /**
     * @return Docker image storage limits in bytes
     * 
     */
    private @Nullable Integer imageStorage;
    /**
     * @return Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    private @Nullable Integer parallelRequest;

    private ContainerRegistryPlanRegistryLimit() {}
    /**
     * @return Docker image storage limits in bytes
     * 
     */
    public Optional<Integer> imageStorage() {
        return Optional.ofNullable(this.imageStorage);
    }
    /**
     * @return Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    public Optional<Integer> parallelRequest() {
        return Optional.ofNullable(this.parallelRequest);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerRegistryPlanRegistryLimit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer imageStorage;
        private @Nullable Integer parallelRequest;
        public Builder() {}
        public Builder(ContainerRegistryPlanRegistryLimit defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.imageStorage = defaults.imageStorage;
    	      this.parallelRequest = defaults.parallelRequest;
        }

        @CustomType.Setter
        public Builder imageStorage(@Nullable Integer imageStorage) {

            this.imageStorage = imageStorage;
            return this;
        }
        @CustomType.Setter
        public Builder parallelRequest(@Nullable Integer parallelRequest) {

            this.parallelRequest = parallelRequest;
            return this;
        }
        public ContainerRegistryPlanRegistryLimit build() {
            final var _resultValue = new ContainerRegistryPlanRegistryLimit();
            _resultValue.imageStorage = imageStorage;
            _resultValue.parallelRequest = parallelRequest;
            return _resultValue;
        }
    }
}