// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetKafkaTopicResult {
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
     * @return Minimum insync replica accepted for this topic.
     * 
     */
    private Integer minInsyncReplicas;
    /**
     * @return Name of the topic.
     * 
     */
    private String name;
    /**
     * @return Number of partitions for this topic.
     * 
     */
    private Integer partitions;
    /**
     * @return Number of replication for this topic.
     * 
     */
    private Integer replication;
    /**
     * @return Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
     * 
     */
    private Integer retentionBytes;
    /**
     * @return Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
     * 
     */
    private Integer retentionHours;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetKafkaTopicResult() {}
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
     * @return Minimum insync replica accepted for this topic.
     * 
     */
    public Integer minInsyncReplicas() {
        return this.minInsyncReplicas;
    }
    /**
     * @return Name of the topic.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Number of partitions for this topic.
     * 
     */
    public Integer partitions() {
        return this.partitions;
    }
    /**
     * @return Number of replication for this topic.
     * 
     */
    public Integer replication() {
        return this.replication;
    }
    /**
     * @return Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited
     * 
     */
    public Integer retentionBytes() {
        return this.retentionBytes;
    }
    /**
     * @return Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited
     * 
     */
    public Integer retentionHours() {
        return this.retentionHours;
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

    public static Builder builder(GetKafkaTopicResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private Integer minInsyncReplicas;
        private String name;
        private Integer partitions;
        private Integer replication;
        private Integer retentionBytes;
        private Integer retentionHours;
        private String serviceName;
        public Builder() {}
        public Builder(GetKafkaTopicResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.minInsyncReplicas = defaults.minInsyncReplicas;
    	      this.name = defaults.name;
    	      this.partitions = defaults.partitions;
    	      this.replication = defaults.replication;
    	      this.retentionBytes = defaults.retentionBytes;
    	      this.retentionHours = defaults.retentionHours;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder minInsyncReplicas(Integer minInsyncReplicas) {
            if (minInsyncReplicas == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "minInsyncReplicas");
            }
            this.minInsyncReplicas = minInsyncReplicas;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder partitions(Integer partitions) {
            if (partitions == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "partitions");
            }
            this.partitions = partitions;
            return this;
        }
        @CustomType.Setter
        public Builder replication(Integer replication) {
            if (replication == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "replication");
            }
            this.replication = replication;
            return this;
        }
        @CustomType.Setter
        public Builder retentionBytes(Integer retentionBytes) {
            if (retentionBytes == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "retentionBytes");
            }
            this.retentionBytes = retentionBytes;
            return this;
        }
        @CustomType.Setter
        public Builder retentionHours(Integer retentionHours) {
            if (retentionHours == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "retentionHours");
            }
            this.retentionHours = retentionHours;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKafkaTopicResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetKafkaTopicResult build() {
            final var _resultValue = new GetKafkaTopicResult();
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.minInsyncReplicas = minInsyncReplicas;
            _resultValue.name = name;
            _resultValue.partitions = partitions;
            _resultValue.replication = replication;
            _resultValue.retentionBytes = retentionBytes;
            _resultValue.retentionHours = retentionHours;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
