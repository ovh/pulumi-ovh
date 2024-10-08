// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerOrderableZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerOrderableZoneArgs Empty = new LoadBalancerOrderableZoneArgs();

    /**
     * The zone three letter code
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The zone three letter code
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The billing planCode for this zone
     * 
     */
    @Import(name="planCode")
    private @Nullable Output<String> planCode;

    /**
     * @return The billing planCode for this zone
     * 
     */
    public Optional<Output<String>> planCode() {
        return Optional.ofNullable(this.planCode);
    }

    private LoadBalancerOrderableZoneArgs() {}

    private LoadBalancerOrderableZoneArgs(LoadBalancerOrderableZoneArgs $) {
        this.name = $.name;
        this.planCode = $.planCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerOrderableZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerOrderableZoneArgs $;

        public Builder() {
            $ = new LoadBalancerOrderableZoneArgs();
        }

        public Builder(LoadBalancerOrderableZoneArgs defaults) {
            $ = new LoadBalancerOrderableZoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The zone three letter code
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The zone three letter code
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param planCode The billing planCode for this zone
         * 
         * @return builder
         * 
         */
        public Builder planCode(@Nullable Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode The billing planCode for this zone
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        public LoadBalancerOrderableZoneArgs build() {
            return $;
        }
    }

}
