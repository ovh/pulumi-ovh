// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceKeyJWKIam {
    /**
     * @return Resource display name
     * 
     */
    private @Nullable String displayName;
    /**
     * @return Unique identifier of the resource
     * 
     */
    private @Nullable String id;
    /**
     * @return Resource tags. Tags that were internally computed are prefixed with ovh:
     * 
     */
    private @Nullable Map<String,String> tags;
    /**
     * @return Unique resource name used in policies
     * 
     */
    private @Nullable String urn;

    private ServiceKeyJWKIam() {}
    /**
     * @return Resource display name
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    /**
     * @return Unique identifier of the resource
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return Resource tags. Tags that were internally computed are prefixed with ovh:
     * 
     */
    public Map<String,String> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return Unique resource name used in policies
     * 
     */
    public Optional<String> urn() {
        return Optional.ofNullable(this.urn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceKeyJWKIam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String displayName;
        private @Nullable String id;
        private @Nullable Map<String,String> tags;
        private @Nullable String urn;
        public Builder() {}
        public Builder(ServiceKeyJWKIam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.tags = defaults.tags;
    	      this.urn = defaults.urn;
        }

        @CustomType.Setter
        public Builder displayName(@Nullable String displayName) {

            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,String> tags) {

            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder urn(@Nullable String urn) {

            this.urn = urn;
            return this;
        }
        public ServiceKeyJWKIam build() {
            final var _resultValue = new ServiceKeyJWKIam();
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.tags = tags;
            _resultValue.urn = urn;
            return _resultValue;
        }
    }
}