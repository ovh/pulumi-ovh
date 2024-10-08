// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCertificatesResult {
    /**
     * @return CA certificate used for the service.
     * 
     */
    private String ca;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
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

    private GetCertificatesResult() {}
    /**
     * @return CA certificate used for the service.
     * 
     */
    public String ca() {
        return this.ca;
    }
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

    public static Builder builder(GetCertificatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ca;
        private String clusterId;
        private String engine;
        private String id;
        private String serviceName;
        public Builder() {}
        public Builder(GetCertificatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ca = defaults.ca;
    	      this.clusterId = defaults.clusterId;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder ca(String ca) {
            if (ca == null) {
              throw new MissingRequiredPropertyException("GetCertificatesResult", "ca");
            }
            this.ca = ca;
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetCertificatesResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetCertificatesResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCertificatesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetCertificatesResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetCertificatesResult build() {
            final var _resultValue = new GetCertificatesResult();
            _resultValue.ca = ca;
            _resultValue.clusterId = clusterId;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
