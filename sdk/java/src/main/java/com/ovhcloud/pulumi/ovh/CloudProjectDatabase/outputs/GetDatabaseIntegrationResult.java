// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetDatabaseIntegrationResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return ID of the destination service.
     * 
     */
    private String destinationServiceId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String engine;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String id;
    /**
     * @return Parameters for the integration.
     * 
     */
    private Map<String,String> parameters;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return ID of the source service.
     * 
     */
    private String sourceServiceId;
    /**
     * @return Current status of the integration.
     * 
     */
    private String status;
    /**
     * @return Type of the integration.
     * 
     */
    private String type;

    private GetDatabaseIntegrationResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return ID of the destination service.
     * 
     */
    public String destinationServiceId() {
        return this.destinationServiceId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String engine() {
        return this.engine;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Parameters for the integration.
     * 
     */
    public Map<String,String> parameters() {
        return this.parameters;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return ID of the source service.
     * 
     */
    public String sourceServiceId() {
        return this.sourceServiceId;
    }
    /**
     * @return Current status of the integration.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Type of the integration.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseIntegrationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String destinationServiceId;
        private String engine;
        private String id;
        private Map<String,String> parameters;
        private String serviceName;
        private String sourceServiceId;
        private String status;
        private String type;
        public Builder() {}
        public Builder(GetDatabaseIntegrationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.destinationServiceId = defaults.destinationServiceId;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.parameters = defaults.parameters;
    	      this.serviceName = defaults.serviceName;
    	      this.sourceServiceId = defaults.sourceServiceId;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder destinationServiceId(String destinationServiceId) {
            if (destinationServiceId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "destinationServiceId");
            }
            this.destinationServiceId = destinationServiceId;
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder parameters(Map<String,String> parameters) {
            if (parameters == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "parameters");
            }
            this.parameters = parameters;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder sourceServiceId(String sourceServiceId) {
            if (sourceServiceId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "sourceServiceId");
            }
            this.sourceServiceId = sourceServiceId;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetDatabaseIntegrationResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetDatabaseIntegrationResult build() {
            final var _resultValue = new GetDatabaseIntegrationResult();
            _resultValue.clusterId = clusterId;
            _resultValue.destinationServiceId = destinationServiceId;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.parameters = parameters;
            _resultValue.serviceName = serviceName;
            _resultValue.sourceServiceId = sourceServiceId;
            _resultValue.status = status;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
