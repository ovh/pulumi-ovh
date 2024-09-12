// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.inputs;

import com.ovhcloud.ovh.CloudProject.inputs.KubeCustomizationApiserverArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeCustomizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeCustomizationArgs Empty = new KubeCustomizationArgs();

    /**
     * Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    @Import(name="apiservers")
    private @Nullable Output<List<KubeCustomizationApiserverArgs>> apiservers;

    /**
     * @return Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    public Optional<Output<List<KubeCustomizationApiserverArgs>>> apiservers() {
        return Optional.ofNullable(this.apiservers);
    }

    private KubeCustomizationArgs() {}

    private KubeCustomizationArgs(KubeCustomizationArgs $) {
        this.apiservers = $.apiservers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeCustomizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeCustomizationArgs $;

        public Builder() {
            $ = new KubeCustomizationArgs();
        }

        public Builder(KubeCustomizationArgs defaults) {
            $ = new KubeCustomizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiservers Kubernetes API server customization
         * 
         * @return builder
         * 
         * @deprecated
         * Use customization_apiserver instead
         * 
         */
        @Deprecated /* Use customization_apiserver instead */
        public Builder apiservers(@Nullable Output<List<KubeCustomizationApiserverArgs>> apiservers) {
            $.apiservers = apiservers;
            return this;
        }

        /**
         * @param apiservers Kubernetes API server customization
         * 
         * @return builder
         * 
         * @deprecated
         * Use customization_apiserver instead
         * 
         */
        @Deprecated /* Use customization_apiserver instead */
        public Builder apiservers(List<KubeCustomizationApiserverArgs> apiservers) {
            return apiservers(Output.of(apiservers));
        }

        /**
         * @param apiservers Kubernetes API server customization
         * 
         * @return builder
         * 
         * @deprecated
         * Use customization_apiserver instead
         * 
         */
        @Deprecated /* Use customization_apiserver instead */
        public Builder apiservers(KubeCustomizationApiserverArgs... apiservers) {
            return apiservers(List.of(apiservers));
        }

        public KubeCustomizationArgs build() {
            return $;
        }
    }

}