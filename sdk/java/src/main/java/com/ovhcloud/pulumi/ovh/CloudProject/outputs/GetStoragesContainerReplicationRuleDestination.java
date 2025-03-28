// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetStoragesContainerReplicationRuleDestination {
    /**
     * @return Destination bucket name
     * 
     */
    private String name;
    /**
     * @return Destination region, can be null if destination bucket has been deleted
     * 
     */
    private String region;
    /**
     * @return Destination storage class
     * 
     */
    private String storageClass;

    private GetStoragesContainerReplicationRuleDestination() {}
    /**
     * @return Destination bucket name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Destination region, can be null if destination bucket has been deleted
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Destination storage class
     * 
     */
    public String storageClass() {
        return this.storageClass;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStoragesContainerReplicationRuleDestination defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String region;
        private String storageClass;
        public Builder() {}
        public Builder(GetStoragesContainerReplicationRuleDestination defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.storageClass = defaults.storageClass;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetStoragesContainerReplicationRuleDestination", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetStoragesContainerReplicationRuleDestination", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder storageClass(String storageClass) {
            if (storageClass == null) {
              throw new MissingRequiredPropertyException("GetStoragesContainerReplicationRuleDestination", "storageClass");
            }
            this.storageClass = storageClass;
            return this;
        }
        public GetStoragesContainerReplicationRuleDestination build() {
            final var _resultValue = new GetStoragesContainerReplicationRuleDestination();
            _resultValue.name = name;
            _resultValue.region = region;
            _resultValue.storageClass = storageClass;
            return _resultValue;
        }
    }
}
