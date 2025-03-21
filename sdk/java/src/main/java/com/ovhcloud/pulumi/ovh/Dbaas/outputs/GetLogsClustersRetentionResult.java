// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLogsClustersRetentionResult {
    private String clusterId;
    /**
     * @return Indexed duration expressed in ISO-8601 format
     * 
     */
    private String duration;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Indicates if a new stream can use it
     * 
     */
    private Boolean isSupported;
    /**
     * @return ID of the retention that can be used when creating a stream
     * 
     */
    private String retentionId;
    /**
     * @return Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT)
     * 
     */
    private String retentionType;
    private String serviceName;

    private GetLogsClustersRetentionResult() {}
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return Indexed duration expressed in ISO-8601 format
     * 
     */
    public String duration() {
        return this.duration;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Indicates if a new stream can use it
     * 
     */
    public Boolean isSupported() {
        return this.isSupported;
    }
    /**
     * @return ID of the retention that can be used when creating a stream
     * 
     */
    public String retentionId() {
        return this.retentionId;
    }
    /**
     * @return Type of the retention (LOGS_INDEXING | LOGS_COLD_STORAGE | METRICS_TENANT)
     * 
     */
    public String retentionType() {
        return this.retentionType;
    }
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogsClustersRetentionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String duration;
        private String id;
        private Boolean isSupported;
        private String retentionId;
        private String retentionType;
        private String serviceName;
        public Builder() {}
        public Builder(GetLogsClustersRetentionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.duration = defaults.duration;
    	      this.id = defaults.id;
    	      this.isSupported = defaults.isSupported;
    	      this.retentionId = defaults.retentionId;
    	      this.retentionType = defaults.retentionType;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder duration(String duration) {
            if (duration == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "duration");
            }
            this.duration = duration;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isSupported(Boolean isSupported) {
            if (isSupported == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "isSupported");
            }
            this.isSupported = isSupported;
            return this;
        }
        @CustomType.Setter
        public Builder retentionId(String retentionId) {
            if (retentionId == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "retentionId");
            }
            this.retentionId = retentionId;
            return this;
        }
        @CustomType.Setter
        public Builder retentionType(String retentionType) {
            if (retentionType == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "retentionType");
            }
            this.retentionType = retentionType;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersRetentionResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetLogsClustersRetentionResult build() {
            final var _resultValue = new GetLogsClustersRetentionResult();
            _resultValue.clusterId = clusterId;
            _resultValue.duration = duration;
            _resultValue.id = id;
            _resultValue.isSupported = isSupported;
            _resultValue.retentionId = retentionId;
            _resultValue.retentionType = retentionType;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
