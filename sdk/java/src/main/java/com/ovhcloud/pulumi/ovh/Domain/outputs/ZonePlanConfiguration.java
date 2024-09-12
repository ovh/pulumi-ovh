// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ZonePlanConfiguration {
    /**
     * @return Identifier of the resource : `zone` or `template`
     * 
     */
    private String label;
    /**
     * @return For `zone`, the value is the zone name `myzone.example.com`. For `template`, the value can be `basic`, `minimized` or  `redirect` which is the same as `minimized` with additional entries for a redirect configuration.
     * 
     */
    private String value;

    private ZonePlanConfiguration() {}
    /**
     * @return Identifier of the resource : `zone` or `template`
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return For `zone`, the value is the zone name `myzone.example.com`. For `template`, the value can be `basic`, `minimized` or  `redirect` which is the same as `minimized` with additional entries for a redirect configuration.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ZonePlanConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(ZonePlanConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("ZonePlanConfiguration", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("ZonePlanConfiguration", "value");
            }
            this.value = value;
            return this;
        }
        public ZonePlanConfiguration build() {
            final var _resultValue = new ZonePlanConfiguration();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}