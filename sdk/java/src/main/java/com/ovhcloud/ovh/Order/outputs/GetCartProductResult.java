// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Order.outputs;

import com.ovhcloud.ovh.Order.outputs.GetCartProductResultPrice;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCartProductResult {
    /**
     * @return Product offer identifier
     * 
     */
    private String planCode;
    /**
     * @return Prices of the product offer
     * 
     */
    private List<GetCartProductResultPrice> prices;
    /**
     * @return Name of the product
     * 
     */
    private String productName;
    /**
     * @return Product type
     * 
     */
    private String productType;

    private GetCartProductResult() {}
    /**
     * @return Product offer identifier
     * 
     */
    public String planCode() {
        return this.planCode;
    }
    /**
     * @return Prices of the product offer
     * 
     */
    public List<GetCartProductResultPrice> prices() {
        return this.prices;
    }
    /**
     * @return Name of the product
     * 
     */
    public String productName() {
        return this.productName;
    }
    /**
     * @return Product type
     * 
     */
    public String productType() {
        return this.productType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartProductResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String planCode;
        private List<GetCartProductResultPrice> prices;
        private String productName;
        private String productType;
        public Builder() {}
        public Builder(GetCartProductResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.planCode = defaults.planCode;
    	      this.prices = defaults.prices;
    	      this.productName = defaults.productName;
    	      this.productType = defaults.productType;
        }

        @CustomType.Setter
        public Builder planCode(String planCode) {
            if (planCode == null) {
              throw new MissingRequiredPropertyException("GetCartProductResult", "planCode");
            }
            this.planCode = planCode;
            return this;
        }
        @CustomType.Setter
        public Builder prices(List<GetCartProductResultPrice> prices) {
            if (prices == null) {
              throw new MissingRequiredPropertyException("GetCartProductResult", "prices");
            }
            this.prices = prices;
            return this;
        }
        public Builder prices(GetCartProductResultPrice... prices) {
            return prices(List.of(prices));
        }
        @CustomType.Setter
        public Builder productName(String productName) {
            if (productName == null) {
              throw new MissingRequiredPropertyException("GetCartProductResult", "productName");
            }
            this.productName = productName;
            return this;
        }
        @CustomType.Setter
        public Builder productType(String productType) {
            if (productType == null) {
              throw new MissingRequiredPropertyException("GetCartProductResult", "productType");
            }
            this.productType = productType;
            return this;
        }
        public GetCartProductResult build() {
            final var _resultValue = new GetCartProductResult();
            _resultValue.planCode = planCode;
            _resultValue.prices = prices;
            _resultValue.productName = productName;
            _resultValue.productType = productType;
            return _resultValue;
        }
    }
}