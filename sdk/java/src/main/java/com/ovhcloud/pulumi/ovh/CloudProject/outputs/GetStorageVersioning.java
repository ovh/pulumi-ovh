// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetStorageVersioning {
    /**
     * @return Versioning status
     * 
     */
    private String status;

    private GetStorageVersioning() {}
    /**
     * @return Versioning status
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStorageVersioning defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String status;
        public Builder() {}
        public Builder(GetStorageVersioning defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetStorageVersioning", "status");
            }
            this.status = status;
            return this;
        }
        public GetStorageVersioning build() {
            final var _resultValue = new GetStorageVersioning();
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
