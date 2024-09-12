// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Vps.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VpsOrderDetail {
    private @Nullable String description;
    /**
     * @return Product type of item in order
     * 
     */
    private @Nullable String detailType;
    private @Nullable String domain;
    private @Nullable Double orderDetailId;
    private @Nullable String quantity;

    private VpsOrderDetail() {}
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Product type of item in order
     * 
     */
    public Optional<String> detailType() {
        return Optional.ofNullable(this.detailType);
    }
    public Optional<String> domain() {
        return Optional.ofNullable(this.domain);
    }
    public Optional<Double> orderDetailId() {
        return Optional.ofNullable(this.orderDetailId);
    }
    public Optional<String> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VpsOrderDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String detailType;
        private @Nullable String domain;
        private @Nullable Double orderDetailId;
        private @Nullable String quantity;
        public Builder() {}
        public Builder(VpsOrderDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.detailType = defaults.detailType;
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
        public Builder detailType(@Nullable String detailType) {

            this.detailType = detailType;
            return this;
        }
        @CustomType.Setter
        public Builder domain(@Nullable String domain) {

            this.domain = domain;
            return this;
        }
        @CustomType.Setter
        public Builder orderDetailId(@Nullable Double orderDetailId) {

            this.orderDetailId = orderDetailId;
            return this;
        }
        @CustomType.Setter
        public Builder quantity(@Nullable String quantity) {

            this.quantity = quantity;
            return this;
        }
        public VpsOrderDetail build() {
            final var _resultValue = new VpsOrderDetail();
            _resultValue.description = description;
            _resultValue.detailType = detailType;
            _resultValue.domain = domain;
            _resultValue.orderDetailId = orderDetailId;
            _resultValue.quantity = quantity;
            return _resultValue;
        }
    }
}