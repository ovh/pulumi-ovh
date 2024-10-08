// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack.outputs;

import com.ovhcloud.pulumi.ovh.Vrack.outputs.VrackOrderDetail;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VrackOrder {
    /**
     * @return date
     * 
     */
    private @Nullable String date;
    /**
     * @return Information about a Bill entry
     * 
     */
    private @Nullable List<VrackOrderDetail> details;
    /**
     * @return expiration date
     * 
     */
    private @Nullable String expirationDate;
    /**
     * @return order id
     * 
     */
    private @Nullable Integer orderId;

    private VrackOrder() {}
    /**
     * @return date
     * 
     */
    public Optional<String> date() {
        return Optional.ofNullable(this.date);
    }
    /**
     * @return Information about a Bill entry
     * 
     */
    public List<VrackOrderDetail> details() {
        return this.details == null ? List.of() : this.details;
    }
    /**
     * @return expiration date
     * 
     */
    public Optional<String> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }
    /**
     * @return order id
     * 
     */
    public Optional<Integer> orderId() {
        return Optional.ofNullable(this.orderId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VrackOrder defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String date;
        private @Nullable List<VrackOrderDetail> details;
        private @Nullable String expirationDate;
        private @Nullable Integer orderId;
        public Builder() {}
        public Builder(VrackOrder defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.date = defaults.date;
    	      this.details = defaults.details;
    	      this.expirationDate = defaults.expirationDate;
    	      this.orderId = defaults.orderId;
        }

        @CustomType.Setter
        public Builder date(@Nullable String date) {

            this.date = date;
            return this;
        }
        @CustomType.Setter
        public Builder details(@Nullable List<VrackOrderDetail> details) {

            this.details = details;
            return this;
        }
        public Builder details(VrackOrderDetail... details) {
            return details(List.of(details));
        }
        @CustomType.Setter
        public Builder expirationDate(@Nullable String expirationDate) {

            this.expirationDate = expirationDate;
            return this;
        }
        @CustomType.Setter
        public Builder orderId(@Nullable Integer orderId) {

            this.orderId = orderId;
            return this;
        }
        public VrackOrder build() {
            final var _resultValue = new VrackOrder();
            _resultValue.date = date;
            _resultValue.details = details;
            _resultValue.expirationDate = expirationDate;
            _resultValue.orderId = orderId;
            return _resultValue;
        }
    }
}
