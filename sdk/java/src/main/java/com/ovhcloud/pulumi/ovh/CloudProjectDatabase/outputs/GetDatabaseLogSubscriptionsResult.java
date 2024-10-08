// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDatabaseLogSubscriptionsResult {
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
    /**
     * @return The list of log subscription ids of the cluster associated with the project.
     * 
     */
    private List<String> subscriptionIds;

    private GetDatabaseLogSubscriptionsResult() {}
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
    /**
     * @return The list of log subscription ids of the cluster associated with the project.
     * 
     */
    public List<String> subscriptionIds() {
        return this.subscriptionIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseLogSubscriptionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String engine;
        private String id;
        private String serviceName;
        private List<String> subscriptionIds;
        public Builder() {}
        public Builder(GetDatabaseLogSubscriptionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
    	      this.subscriptionIds = defaults.subscriptionIds;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionsResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionsResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionIds(List<String> subscriptionIds) {
            if (subscriptionIds == null) {
              throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionsResult", "subscriptionIds");
            }
            this.subscriptionIds = subscriptionIds;
            return this;
        }
        public Builder subscriptionIds(String... subscriptionIds) {
            return subscriptionIds(List.of(subscriptionIds));
        }
        public GetDatabaseLogSubscriptionsResult build() {
            final var _resultValue = new GetDatabaseLogSubscriptionsResult();
            _resultValue.clusterId = clusterId;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            _resultValue.subscriptionIds = subscriptionIds;
            return _resultValue;
        }
    }
}
