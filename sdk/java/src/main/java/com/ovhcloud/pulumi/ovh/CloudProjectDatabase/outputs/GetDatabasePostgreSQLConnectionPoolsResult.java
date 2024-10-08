// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDatabasePostgreSQLConnectionPoolsResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return The list of patterns ids of the opensearch cluster associated with the project.
     * 
     */
    private List<String> connectionPoolIds;
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

    private GetDatabasePostgreSQLConnectionPoolsResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The list of patterns ids of the opensearch cluster associated with the project.
     * 
     */
    public List<String> connectionPoolIds() {
        return this.connectionPoolIds;
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

    public static Builder builder(GetDatabasePostgreSQLConnectionPoolsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private List<String> connectionPoolIds;
        private String id;
        private String serviceName;
        public Builder() {}
        public Builder(GetDatabasePostgreSQLConnectionPoolsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.connectionPoolIds = defaults.connectionPoolIds;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgreSQLConnectionPoolsResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder connectionPoolIds(List<String> connectionPoolIds) {
            if (connectionPoolIds == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgreSQLConnectionPoolsResult", "connectionPoolIds");
            }
            this.connectionPoolIds = connectionPoolIds;
            return this;
        }
        public Builder connectionPoolIds(String... connectionPoolIds) {
            return connectionPoolIds(List.of(connectionPoolIds));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgreSQLConnectionPoolsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgreSQLConnectionPoolsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetDatabasePostgreSQLConnectionPoolsResult build() {
            final var _resultValue = new GetDatabasePostgreSQLConnectionPoolsResult();
            _resultValue.clusterId = clusterId;
            _resultValue.connectionPoolIds = connectionPoolIds;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
