// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class KubeNodePoolTemplateSpec {
    /**
     * @return taints
     * 
     */
    private List<Map<String,String>> taints;
    /**
     * @return unschedulable
     * 
     */
    private Boolean unschedulable;

    private KubeNodePoolTemplateSpec() {}
    /**
     * @return taints
     * 
     */
    public List<Map<String,String>> taints() {
        return this.taints;
    }
    /**
     * @return unschedulable
     * 
     */
    public Boolean unschedulable() {
        return this.unschedulable;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeNodePoolTemplateSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Map<String,String>> taints;
        private Boolean unschedulable;
        public Builder() {}
        public Builder(KubeNodePoolTemplateSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.taints = defaults.taints;
    	      this.unschedulable = defaults.unschedulable;
        }

        @CustomType.Setter
        public Builder taints(List<Map<String,String>> taints) {
            if (taints == null) {
              throw new MissingRequiredPropertyException("KubeNodePoolTemplateSpec", "taints");
            }
            this.taints = taints;
            return this;
        }
        @CustomType.Setter
        public Builder unschedulable(Boolean unschedulable) {
            if (unschedulable == null) {
              throw new MissingRequiredPropertyException("KubeNodePoolTemplateSpec", "unschedulable");
            }
            this.unschedulable = unschedulable;
            return this;
        }
        public KubeNodePoolTemplateSpec build() {
            final var _resultValue = new KubeNodePoolTemplateSpec();
            _resultValue.taints = taints;
            _resultValue.unschedulable = unschedulable;
            return _resultValue;
        }
    }
}
