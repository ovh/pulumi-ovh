// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRegionService {
    /**
     * @return The name of the region associated with the public cloud project.
     * 
     */
    private String name;
    /**
     * @return the status of the service
     * 
     */
    private String status;

    private GetRegionService() {}
    /**
     * @return The name of the region associated with the public cloud project.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return the status of the service
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegionService defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String status;
        public Builder() {}
        public Builder(GetRegionService defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRegionService", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetRegionService", "status");
            }
            this.status = status;
            return this;
        }
        public GetRegionService build() {
            final var _resultValue = new GetRegionService();
            _resultValue.name = name;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
