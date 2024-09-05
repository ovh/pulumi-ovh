// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetKafkaAclResult {
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
     * @return Permission to give to this username on this topic.
     * 
     */
    private String permission;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Topic affected by this ACL.
     * 
     */
    private String topic;
    /**
     * @return Username affected by this ACL.
     * 
     */
    private String username;

    private GetKafkaAclResult() {}
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
     * @return Permission to give to this username on this topic.
     * 
     */
    public String permission() {
        return this.permission;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Topic affected by this ACL.
     * 
     */
    public String topic() {
        return this.topic;
    }
    /**
     * @return Username affected by this ACL.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKafkaAclResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private String permission;
        private String serviceName;
        private String topic;
        private String username;
        public Builder() {}
        public Builder(GetKafkaAclResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.permission = defaults.permission;
    	      this.serviceName = defaults.serviceName;
    	      this.topic = defaults.topic;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            if (permission == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "permission");
            }
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder topic(String topic) {
            if (topic == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "topic");
            }
            this.topic = topic;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetKafkaAclResult build() {
            final var _resultValue = new GetKafkaAclResult();
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.permission = permission;
            _resultValue.serviceName = serviceName;
            _resultValue.topic = topic;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}