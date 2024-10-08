// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOpenSearchPatternsResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of patterns ids of the opensearch cluster associated with the project.
     * 
     */
    private List<String> patternIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetOpenSearchPatternsResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of patterns ids of the opensearch cluster associated with the project.
     * 
     */
    public List<String> patternIds() {
        return this.patternIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOpenSearchPatternsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private List<String> patternIds;
        private String serviceName;
        public Builder() {}
        public Builder(GetOpenSearchPatternsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.patternIds = defaults.patternIds;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternsResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder patternIds(List<String> patternIds) {
            if (patternIds == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternsResult", "patternIds");
            }
            this.patternIds = patternIds;
            return this;
        }
        public Builder patternIds(String... patternIds) {
            return patternIds(List.of(patternIds));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetOpenSearchPatternsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetOpenSearchPatternsResult build() {
            final var _resultValue = new GetOpenSearchPatternsResult();
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.patternIds = patternIds;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
