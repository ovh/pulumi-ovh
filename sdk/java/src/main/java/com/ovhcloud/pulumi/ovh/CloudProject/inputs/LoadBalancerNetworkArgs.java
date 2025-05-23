// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.LoadBalancerNetworkPrivateArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.Objects;


public final class LoadBalancerNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerNetworkArgs Empty = new LoadBalancerNetworkArgs();

    /**
     * Information to private network
     * 
     */
    @Import(name="private", required=true)
    private Output<LoadBalancerNetworkPrivateArgs> private_;

    /**
     * @return Information to private network
     * 
     */
    public Output<LoadBalancerNetworkPrivateArgs> private_() {
        return this.private_;
    }

    private LoadBalancerNetworkArgs() {}

    private LoadBalancerNetworkArgs(LoadBalancerNetworkArgs $) {
        this.private_ = $.private_;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerNetworkArgs $;

        public Builder() {
            $ = new LoadBalancerNetworkArgs();
        }

        public Builder(LoadBalancerNetworkArgs defaults) {
            $ = new LoadBalancerNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param private_ Information to private network
         * 
         * @return builder
         * 
         */
        public Builder private_(Output<LoadBalancerNetworkPrivateArgs> private_) {
            $.private_ = private_;
            return this;
        }

        /**
         * @param private_ Information to private network
         * 
         * @return builder
         * 
         */
        public Builder private_(LoadBalancerNetworkPrivateArgs private_) {
            return private_(Output.of(private_));
        }

        public LoadBalancerNetworkArgs build() {
            if ($.private_ == null) {
                throw new MissingRequiredPropertyException("LoadBalancerNetworkArgs", "private_");
            }
            return $;
        }
    }

}
