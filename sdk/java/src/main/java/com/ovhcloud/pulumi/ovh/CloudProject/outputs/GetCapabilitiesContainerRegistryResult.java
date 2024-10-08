// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetCapabilitiesContainerRegistryResultPlan;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesContainerRegistryResult {
    /**
     * @return Available plans in the region
     * 
     */
    private List<GetCapabilitiesContainerRegistryResultPlan> plans;
    /**
     * @return The region name
     * 
     */
    private String regionName;

    private GetCapabilitiesContainerRegistryResult() {}
    /**
     * @return Available plans in the region
     * 
     */
    public List<GetCapabilitiesContainerRegistryResultPlan> plans() {
        return this.plans;
    }
    /**
     * @return The region name
     * 
     */
    public String regionName() {
        return this.regionName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesContainerRegistryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetCapabilitiesContainerRegistryResultPlan> plans;
        private String regionName;
        public Builder() {}
        public Builder(GetCapabilitiesContainerRegistryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.plans = defaults.plans;
    	      this.regionName = defaults.regionName;
        }

        @CustomType.Setter
        public Builder plans(List<GetCapabilitiesContainerRegistryResultPlan> plans) {
            if (plans == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesContainerRegistryResult", "plans");
            }
            this.plans = plans;
            return this;
        }
        public Builder plans(GetCapabilitiesContainerRegistryResultPlan... plans) {
            return plans(List.of(plans));
        }
        @CustomType.Setter
        public Builder regionName(String regionName) {
            if (regionName == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesContainerRegistryResult", "regionName");
            }
            this.regionName = regionName;
            return this;
        }
        public GetCapabilitiesContainerRegistryResult build() {
            final var _resultValue = new GetCapabilitiesContainerRegistryResult();
            _resultValue.plans = plans;
            _resultValue.regionName = regionName;
            return _resultValue;
        }
    }
}
