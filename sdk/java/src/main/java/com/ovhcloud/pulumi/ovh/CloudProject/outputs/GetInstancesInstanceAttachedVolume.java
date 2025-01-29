// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstancesInstanceAttachedVolume {
    /**
     * @return Instance id
     * 
     */
    private String id;

    private GetInstancesInstanceAttachedVolume() {}
    /**
     * @return Instance id
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstanceAttachedVolume defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        public Builder() {}
        public Builder(GetInstancesInstanceAttachedVolume defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstancesInstanceAttachedVolume", "id");
            }
            this.id = id;
            return this;
        }
        public GetInstancesInstanceAttachedVolume build() {
            final var _resultValue = new GetInstancesInstanceAttachedVolume();
            _resultValue.id = id;
            return _resultValue;
        }
    }
}
