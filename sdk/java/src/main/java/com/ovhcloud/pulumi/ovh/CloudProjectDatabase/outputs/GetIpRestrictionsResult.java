// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIpRestrictionsResult {
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
     * @return The list of IP restriction of the database associated with the project.
     * 
     */
    private List<String> ips;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String serviceName;

    private GetIpRestrictionsResult() {}
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
     * @return The list of IP restriction of the database associated with the project.
     * 
     */
    public List<String> ips() {
        return this.ips;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpRestrictionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String engine;
        private String id;
        private List<String> ips;
        private @Nullable String serviceName;
        public Builder() {}
        public Builder(GetIpRestrictionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.ips = defaults.ips;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetIpRestrictionsResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetIpRestrictionsResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIpRestrictionsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ips(List<String> ips) {
            if (ips == null) {
              throw new MissingRequiredPropertyException("GetIpRestrictionsResult", "ips");
            }
            this.ips = ips;
            return this;
        }
        public Builder ips(String... ips) {
            return ips(List.of(ips));
        }
        @CustomType.Setter
        public Builder serviceName(@Nullable String serviceName) {

            this.serviceName = serviceName;
            return this;
        }
        public GetIpRestrictionsResult build() {
            final var _resultValue = new GetIpRestrictionsResult();
            _resultValue.clusterId = clusterId;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.ips = ips;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
