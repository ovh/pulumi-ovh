// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.GetKubeCustomizationApiserverArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;


public final class GetKubeCustomizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetKubeCustomizationArgs Empty = new GetKubeCustomizationArgs();

    /**
     * Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    @Import(name="apiservers", required=true)
    private Output<List<GetKubeCustomizationApiserverArgs>> apiservers;

    /**
     * @return Kubernetes API server customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    public Output<List<GetKubeCustomizationApiserverArgs>> apiservers() {
        return this.apiservers;
    }

    private GetKubeCustomizationArgs() {}

    private GetKubeCustomizationArgs(GetKubeCustomizationArgs $) {
        this.apiservers = $.apiservers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeCustomizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeCustomizationArgs $;

        public Builder() {
            $ = new GetKubeCustomizationArgs();
        }

        public Builder(GetKubeCustomizationArgs defaults) {
            $ = new GetKubeCustomizationArgs(Objects.requireNonNull(defaults));
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
        public Builder apiservers(Output<List<GetKubeCustomizationApiserverArgs>> apiservers) {
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
        public Builder apiservers(List<GetKubeCustomizationApiserverArgs> apiservers) {
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
        public Builder apiservers(GetKubeCustomizationApiserverArgs... apiservers) {
            return apiservers(List.of(apiservers));
        }

        public GetKubeCustomizationArgs build() {
            if ($.apiservers == null) {
                throw new MissingRequiredPropertyException("GetKubeCustomizationArgs", "apiservers");
            }
            return $;
        }
    }

}
