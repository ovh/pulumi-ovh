// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetKubeNodePoolTemplateMetadata {
    /**
     * @return annotations
     * 
     */
    private @Nullable Map<String,String> annotations;
    /**
     * @return finalizers
     * 
     */
    private @Nullable List<String> finalizers;
    /**
     * @return labels
     * 
     */
    private @Nullable Map<String,String> labels;

    private GetKubeNodePoolTemplateMetadata() {}
    /**
     * @return annotations
     * 
     */
    public Map<String,String> annotations() {
        return this.annotations == null ? Map.of() : this.annotations;
    }
    /**
     * @return finalizers
     * 
     */
    public List<String> finalizers() {
        return this.finalizers == null ? List.of() : this.finalizers;
    }
    /**
     * @return labels
     * 
     */
    public Map<String,String> labels() {
        return this.labels == null ? Map.of() : this.labels;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeNodePoolTemplateMetadata defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable List<String> finalizers;
        private @Nullable Map<String,String> labels;
        public Builder() {}
        public Builder(GetKubeNodePoolTemplateMetadata defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.finalizers = defaults.finalizers;
    	      this.labels = defaults.labels;
        }

        @CustomType.Setter
        public Builder annotations(@Nullable Map<String,String> annotations) {

            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder finalizers(@Nullable List<String> finalizers) {

            this.finalizers = finalizers;
            return this;
        }
        public Builder finalizers(String... finalizers) {
            return finalizers(List.of(finalizers));
        }
        @CustomType.Setter
        public Builder labels(@Nullable Map<String,String> labels) {

            this.labels = labels;
            return this;
        }
        public GetKubeNodePoolTemplateMetadata build() {
            final var _resultValue = new GetKubeNodePoolTemplateMetadata();
            _resultValue.annotations = annotations;
            _resultValue.finalizers = finalizers;
            _resultValue.labels = labels;
            return _resultValue;
        }
    }
}
