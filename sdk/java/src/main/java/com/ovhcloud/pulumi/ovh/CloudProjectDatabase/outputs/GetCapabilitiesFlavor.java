// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesFlavor {
    /**
     * @return Flavor core number.
     * 
     */
    private Integer core;
    /**
     * @return Flavor ram size in GB.
     * 
     */
    private Integer memory;
    /**
     * @return Name of the plan.
     * 
     */
    private String name;
    /**
     * @return Flavor disk size in GB.
     * 
     */
    private Integer storage;

    private GetCapabilitiesFlavor() {}
    /**
     * @return Flavor core number.
     * 
     */
    public Integer core() {
        return this.core;
    }
    /**
     * @return Flavor ram size in GB.
     * 
     */
    public Integer memory() {
        return this.memory;
    }
    /**
     * @return Name of the plan.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Flavor disk size in GB.
     * 
     */
    public Integer storage() {
        return this.storage;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesFlavor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer core;
        private Integer memory;
        private String name;
        private Integer storage;
        public Builder() {}
        public Builder(GetCapabilitiesFlavor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.core = defaults.core;
    	      this.memory = defaults.memory;
    	      this.name = defaults.name;
    	      this.storage = defaults.storage;
        }

        @CustomType.Setter
        public Builder core(Integer core) {
            if (core == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesFlavor", "core");
            }
            this.core = core;
            return this;
        }
        @CustomType.Setter
        public Builder memory(Integer memory) {
            if (memory == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesFlavor", "memory");
            }
            this.memory = memory;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesFlavor", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder storage(Integer storage) {
            if (storage == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesFlavor", "storage");
            }
            this.storage = storage;
            return this;
        }
        public GetCapabilitiesFlavor build() {
            final var _resultValue = new GetCapabilitiesFlavor();
            _resultValue.core = core;
            _resultValue.memory = memory;
            _resultValue.name = name;
            _resultValue.storage = storage;
            return _resultValue;
        }
    }
}
