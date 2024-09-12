// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetKafkaUserAccessResult {
    /**
     * @return User cert.
     * 
     */
    private String cert;
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
     * @return (Sensitive) User key for the cert.
     * 
     */
    private String key;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String userId;

    private GetKafkaUserAccessResult() {}
    /**
     * @return User cert.
     * 
     */
    public String cert() {
        return this.cert;
    }
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
     * @return (Sensitive) User key for the cert.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKafkaUserAccessResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cert;
        private String clusterId;
        private String id;
        private String key;
        private String serviceName;
        private String userId;
        public Builder() {}
        public Builder(GetKafkaUserAccessResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cert = defaults.cert;
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.serviceName = defaults.serviceName;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder cert(String cert) {
            if (cert == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "cert");
            }
            this.cert = cert;
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetKafkaUserAccessResult", "userId");
            }
            this.userId = userId;
            return this;
        }
        public GetKafkaUserAccessResult build() {
            final var _resultValue = new GetKafkaUserAccessResult();
            _resultValue.cert = cert;
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.key = key;
            _resultValue.serviceName = serviceName;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}