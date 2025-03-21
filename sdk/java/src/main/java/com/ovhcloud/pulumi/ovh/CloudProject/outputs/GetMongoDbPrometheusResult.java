// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMongoDbPrometheusResult {
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
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Name of the srv domain endpoint.
     * 
     */
    private String srvDomain;
    /**
     * @return name of the prometheus user.
     * 
     */
    private String username;

    private GetMongoDbPrometheusResult() {}
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
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Name of the srv domain endpoint.
     * 
     */
    public String srvDomain() {
        return this.srvDomain;
    }
    /**
     * @return name of the prometheus user.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMongoDbPrometheusResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private String serviceName;
        private String srvDomain;
        private String username;
        public Builder() {}
        public Builder(GetMongoDbPrometheusResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
    	      this.srvDomain = defaults.srvDomain;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetMongoDbPrometheusResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMongoDbPrometheusResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetMongoDbPrometheusResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder srvDomain(String srvDomain) {
            if (srvDomain == null) {
              throw new MissingRequiredPropertyException("GetMongoDbPrometheusResult", "srvDomain");
            }
            this.srvDomain = srvDomain;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetMongoDbPrometheusResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetMongoDbPrometheusResult build() {
            final var _resultValue = new GetMongoDbPrometheusResult();
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            _resultValue.srvDomain = srvDomain;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
