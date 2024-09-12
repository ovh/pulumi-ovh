// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOpenSearchPatternResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String id;
    /**
     * @return Maximum number of index for this pattern.
     * 
     */
    private Integer maxIndexCount;
    /**
     * @return Pattern format.
     * 
     */
    private String pattern;
    /**
     * @return Current status of the pattern.
     * 
     */
    private String serviceName;

    private GetOpenSearchPatternResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Maximum number of index for this pattern.
     * 
     */
    public Integer maxIndexCount() {
        return this.maxIndexCount;
    }
    /**
     * @return Pattern format.
     * 
     */
    public String pattern() {
        return this.pattern;
    }
    /**
     * @return Current status of the pattern.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOpenSearchPatternResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private Integer maxIndexCount;
        private String pattern;
        private String serviceName;
        public Builder() {}
        public Builder(GetOpenSearchPatternResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.maxIndexCount = defaults.maxIndexCount;
    	      this.pattern = defaults.pattern;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder maxIndexCount(Integer maxIndexCount) {
            if (maxIndexCount == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternResult", "maxIndexCount");
            }
            this.maxIndexCount = maxIndexCount;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(String pattern) {
            if (pattern == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternResult", "pattern");
            }
            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetOpenSearchPatternResult build() {
            final var _resultValue = new GetOpenSearchPatternResult();
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.maxIndexCount = maxIndexCount;
            _resultValue.pattern = pattern;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}