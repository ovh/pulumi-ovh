// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDatabasesResult {
    /**
     * @return The list of managed databases ids of the project.
     * 
     */
    private List<String> clusterIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String engine;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetDatabasesResult() {}
    /**
     * @return The list of managed databases ids of the project.
     * 
     */
    public List<String> clusterIds() {
        return this.clusterIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String engine() {
        return this.engine;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
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

    public static Builder builder(GetDatabasesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> clusterIds;
        private String engine;
        private String id;
        private String serviceName;
        public Builder() {}
        public Builder(GetDatabasesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterIds = defaults.clusterIds;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterIds(List<String> clusterIds) {
            if (clusterIds == null) {
              throw new MissingRequiredPropertyException("GetDatabasesResult", "clusterIds");
            }
            this.clusterIds = clusterIds;
            return this;
        }
        public Builder clusterIds(String... clusterIds) {
            return clusterIds(List.of(clusterIds));
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetDatabasesResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabasesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabasesResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetDatabasesResult build() {
            final var _resultValue = new GetDatabasesResult();
            _resultValue.clusterIds = clusterIds;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
