// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadBalancerOrderDetail {
    /**
     * @return description
     * 
     */
    private @Nullable String description;
    /**
     * @return expiration date
     * 
     */
    private @Nullable String domain;
    /**
     * @return order detail id
     * 
     */
    private @Nullable Integer orderDetailId;
    /**
     * @return quantity
     * 
     */
    private @Nullable String quantity;

    private LoadBalancerOrderDetail() {}
    /**
     * @return description
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return expiration date
     * 
     */
    public Optional<String> domain() {
        return Optional.ofNullable(this.domain);
    }
    /**
     * @return order detail id
     * 
     */
    public Optional<Integer> orderDetailId() {
        return Optional.ofNullable(this.orderDetailId);
    }
    /**
     * @return quantity
     * 
     */
    public Optional<String> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerOrderDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String domain;
        private @Nullable Integer orderDetailId;
        private @Nullable String quantity;
        public Builder() {}
        public Builder(LoadBalancerOrderDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.domain = defaults.domain;
    	      this.orderDetailId = defaults.orderDetailId;
    	      this.quantity = defaults.quantity;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder domain(@Nullable String domain) {

            this.domain = domain;
            return this;
        }
        @CustomType.Setter
        public Builder orderDetailId(@Nullable Integer orderDetailId) {

            this.orderDetailId = orderDetailId;
            return this;
        }
        @CustomType.Setter
        public Builder quantity(@Nullable String quantity) {

            this.quantity = quantity;
            return this;
        }
        public LoadBalancerOrderDetail build() {
            final var _resultValue = new LoadBalancerOrderDetail();
            _resultValue.description = description;
            _resultValue.domain = domain;
            _resultValue.orderDetailId = orderDetailId;
            _resultValue.quantity = quantity;
            return _resultValue;
        }
    }
}
