// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetKubeCustomizationApiserver;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetKubeCustomization {
    /**
     * @return Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    private List<GetKubeCustomizationApiserver> apiservers;

    private GetKubeCustomization() {}
    /**
     * @return Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    public List<GetKubeCustomizationApiserver> apiservers() {
        return this.apiservers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeCustomization defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetKubeCustomizationApiserver> apiservers;
        public Builder() {}
        public Builder(GetKubeCustomization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiservers = defaults.apiservers;
        }

        @CustomType.Setter
        public Builder apiservers(List<GetKubeCustomizationApiserver> apiservers) {
            if (apiservers == null) {
              throw new MissingRequiredPropertyException("GetKubeCustomization", "apiservers");
            }
            this.apiservers = apiservers;
            return this;
        }
        public Builder apiservers(GetKubeCustomizationApiserver... apiservers) {
            return apiservers(List.of(apiservers));
        }
        public GetKubeCustomization build() {
            final var _resultValue = new GetKubeCustomization();
            _resultValue.apiservers = apiservers;
            return _resultValue;
        }
    }
}
