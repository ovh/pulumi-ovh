// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceAddress {
    /**
     * @return IP address
     * 
     */
    private String ip;
    /**
     * @return IP version
     * 
     */
    private Integer version;

    private GetInstanceAddress() {}
    /**
     * @return IP address
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return IP version
     * 
     */
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceAddress defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private Integer version;
        public Builder() {}
        public Builder(GetInstanceAddress defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            if (ip == null) {
              throw new MissingRequiredPropertyException("GetInstanceAddress", "ip");
            }
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetInstanceAddress", "version");
            }
            this.version = version;
            return this;
        }
        public GetInstanceAddress build() {
            final var _resultValue = new GetInstanceAddress();
            _resultValue.ip = ip;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
