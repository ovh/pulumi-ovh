// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubeNodePoolTemplateSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetKubeNodePoolTemplateSpecArgs Empty = new GetKubeNodePoolTemplateSpecArgs();

    /**
     * taints
     * 
     */
    @Import(name="taints")
    private @Nullable Output<List<Map<String,String>>> taints;

    /**
     * @return taints
     * 
     */
    public Optional<Output<List<Map<String,String>>>> taints() {
        return Optional.ofNullable(this.taints);
    }

    /**
     * unschedulable
     * 
     */
    @Import(name="unschedulable")
    private @Nullable Output<Boolean> unschedulable;

    /**
     * @return unschedulable
     * 
     */
    public Optional<Output<Boolean>> unschedulable() {
        return Optional.ofNullable(this.unschedulable);
    }

    private GetKubeNodePoolTemplateSpecArgs() {}

    private GetKubeNodePoolTemplateSpecArgs(GetKubeNodePoolTemplateSpecArgs $) {
        this.taints = $.taints;
        this.unschedulable = $.unschedulable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeNodePoolTemplateSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeNodePoolTemplateSpecArgs $;

        public Builder() {
            $ = new GetKubeNodePoolTemplateSpecArgs();
        }

        public Builder(GetKubeNodePoolTemplateSpecArgs defaults) {
            $ = new GetKubeNodePoolTemplateSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param taints taints
         * 
         * @return builder
         * 
         */
        public Builder taints(@Nullable Output<List<Map<String,String>>> taints) {
            $.taints = taints;
            return this;
        }

        /**
         * @param taints taints
         * 
         * @return builder
         * 
         */
        public Builder taints(List<Map<String,String>> taints) {
            return taints(Output.of(taints));
        }

        /**
         * @param taints taints
         * 
         * @return builder
         * 
         */
        public Builder taints(Map<String,String>... taints) {
            return taints(List.of(taints));
        }

        /**
         * @param unschedulable unschedulable
         * 
         * @return builder
         * 
         */
        public Builder unschedulable(@Nullable Output<Boolean> unschedulable) {
            $.unschedulable = unschedulable;
            return this;
        }

        /**
         * @param unschedulable unschedulable
         * 
         * @return builder
         * 
         */
        public Builder unschedulable(Boolean unschedulable) {
            return unschedulable(Output.of(unschedulable));
        }

        public GetKubeNodePoolTemplateSpecArgs build() {
            return $;
        }
    }

}
