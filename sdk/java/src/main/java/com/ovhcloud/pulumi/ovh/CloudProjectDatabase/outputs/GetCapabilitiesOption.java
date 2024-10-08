// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesOption {
    /**
     * @return Name of the plan.
     * 
     */
    private String name;
    /**
     * @return Type of the option.
     * 
     */
    private String type;

    private GetCapabilitiesOption() {}
    /**
     * @return Name of the plan.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Type of the option.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String type;
        public Builder() {}
        public Builder(GetCapabilitiesOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesOption", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesOption", "type");
            }
            this.type = type;
            return this;
        }
        public GetCapabilitiesOption build() {
            final var _resultValue = new GetCapabilitiesOption();
            _resultValue.name = name;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
