// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Order.outputs;

import com.ovhcloud.pulumi.ovh.Order.outputs.GetCartProductPlanSelectedPricePrice;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCartProductPlanSelectedPrice {
    /**
     * @return Capacities of the pricing (type of pricing)
     * 
     */
    private List<Object> capacities;
    /**
     * @return Description of the pricing
     * 
     */
    private String description;
    /**
     * @return Duration for ordering the product
     * 
     */
    private String duration;
    /**
     * @return Interval of renewal
     * 
     */
    private Integer interval;
    /**
     * @return Maximum quantity that can be ordered
     * 
     */
    private Integer maximumQuantity;
    /**
     * @return Maximum repeat for renewal
     * 
     */
    private Integer maximumRepeat;
    /**
     * @return Minimum quantity that can be ordered
     * 
     */
    private Integer minimumQuantity;
    /**
     * @return Minimum repeat for renewal
     * 
     */
    private Integer minimumRepeat;
    /**
     * @return Price of the product in micro-centims
     * 
     */
    private Integer priceInUcents;
    /**
     * @return Price of the product (Price with its currency and textual representation)
     * 
     */
    private List<GetCartProductPlanSelectedPricePrice> prices;
    /**
     * @return Pricing model identifier
     * 
     */
    private String pricingMode;
    /**
     * @return Pricing type
     * 
     */
    private String pricingType;

    private GetCartProductPlanSelectedPrice() {}
    /**
     * @return Capacities of the pricing (type of pricing)
     * 
     */
    public List<Object> capacities() {
        return this.capacities;
    }
    /**
     * @return Description of the pricing
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Duration for ordering the product
     * 
     */
    public String duration() {
        return this.duration;
    }
    /**
     * @return Interval of renewal
     * 
     */
    public Integer interval() {
        return this.interval;
    }
    /**
     * @return Maximum quantity that can be ordered
     * 
     */
    public Integer maximumQuantity() {
        return this.maximumQuantity;
    }
    /**
     * @return Maximum repeat for renewal
     * 
     */
    public Integer maximumRepeat() {
        return this.maximumRepeat;
    }
    /**
     * @return Minimum quantity that can be ordered
     * 
     */
    public Integer minimumQuantity() {
        return this.minimumQuantity;
    }
    /**
     * @return Minimum repeat for renewal
     * 
     */
    public Integer minimumRepeat() {
        return this.minimumRepeat;
    }
    /**
     * @return Price of the product in micro-centims
     * 
     */
    public Integer priceInUcents() {
        return this.priceInUcents;
    }
    /**
     * @return Price of the product (Price with its currency and textual representation)
     * 
     */
    public List<GetCartProductPlanSelectedPricePrice> prices() {
        return this.prices;
    }
    /**
     * @return Pricing model identifier
     * 
     */
    public String pricingMode() {
        return this.pricingMode;
    }
    /**
     * @return Pricing type
     * 
     */
    public String pricingType() {
        return this.pricingType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartProductPlanSelectedPrice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Object> capacities;
        private String description;
        private String duration;
        private Integer interval;
        private Integer maximumQuantity;
        private Integer maximumRepeat;
        private Integer minimumQuantity;
        private Integer minimumRepeat;
        private Integer priceInUcents;
        private List<GetCartProductPlanSelectedPricePrice> prices;
        private String pricingMode;
        private String pricingType;
        public Builder() {}
        public Builder(GetCartProductPlanSelectedPrice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.capacities = defaults.capacities;
    	      this.description = defaults.description;
    	      this.duration = defaults.duration;
    	      this.interval = defaults.interval;
    	      this.maximumQuantity = defaults.maximumQuantity;
    	      this.maximumRepeat = defaults.maximumRepeat;
    	      this.minimumQuantity = defaults.minimumQuantity;
    	      this.minimumRepeat = defaults.minimumRepeat;
    	      this.priceInUcents = defaults.priceInUcents;
    	      this.prices = defaults.prices;
    	      this.pricingMode = defaults.pricingMode;
    	      this.pricingType = defaults.pricingType;
        }

        @CustomType.Setter
        public Builder capacities(List<Object> capacities) {
            if (capacities == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "capacities");
            }
            this.capacities = capacities;
            return this;
        }
        public Builder capacities(Object... capacities) {
            return capacities(List.of(capacities));
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder duration(String duration) {
            if (duration == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "duration");
            }
            this.duration = duration;
            return this;
        }
        @CustomType.Setter
        public Builder interval(Integer interval) {
            if (interval == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "interval");
            }
            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder maximumQuantity(Integer maximumQuantity) {
            if (maximumQuantity == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "maximumQuantity");
            }
            this.maximumQuantity = maximumQuantity;
            return this;
        }
        @CustomType.Setter
        public Builder maximumRepeat(Integer maximumRepeat) {
            if (maximumRepeat == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "maximumRepeat");
            }
            this.maximumRepeat = maximumRepeat;
            return this;
        }
        @CustomType.Setter
        public Builder minimumQuantity(Integer minimumQuantity) {
            if (minimumQuantity == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "minimumQuantity");
            }
            this.minimumQuantity = minimumQuantity;
            return this;
        }
        @CustomType.Setter
        public Builder minimumRepeat(Integer minimumRepeat) {
            if (minimumRepeat == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "minimumRepeat");
            }
            this.minimumRepeat = minimumRepeat;
            return this;
        }
        @CustomType.Setter
        public Builder priceInUcents(Integer priceInUcents) {
            if (priceInUcents == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "priceInUcents");
            }
            this.priceInUcents = priceInUcents;
            return this;
        }
        @CustomType.Setter
        public Builder prices(List<GetCartProductPlanSelectedPricePrice> prices) {
            if (prices == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "prices");
            }
            this.prices = prices;
            return this;
        }
        public Builder prices(GetCartProductPlanSelectedPricePrice... prices) {
            return prices(List.of(prices));
        }
        @CustomType.Setter
        public Builder pricingMode(String pricingMode) {
            if (pricingMode == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "pricingMode");
            }
            this.pricingMode = pricingMode;
            return this;
        }
        @CustomType.Setter
        public Builder pricingType(String pricingType) {
            if (pricingType == null) {
              throw new MissingRequiredPropertyException("GetCartProductPlanSelectedPrice", "pricingType");
            }
            this.pricingType = pricingType;
            return this;
        }
        public GetCartProductPlanSelectedPrice build() {
            final var _resultValue = new GetCartProductPlanSelectedPrice();
            _resultValue.capacities = capacities;
            _resultValue.description = description;
            _resultValue.duration = duration;
            _resultValue.interval = interval;
            _resultValue.maximumQuantity = maximumQuantity;
            _resultValue.maximumRepeat = maximumRepeat;
            _resultValue.minimumQuantity = minimumQuantity;
            _resultValue.minimumRepeat = minimumRepeat;
            _resultValue.priceInUcents = priceInUcents;
            _resultValue.prices = prices;
            _resultValue.pricingMode = pricingMode;
            _resultValue.pricingType = pricingType;
            return _resultValue;
        }
    }
}