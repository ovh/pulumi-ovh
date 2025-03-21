// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.inputs;

import com.ovhcloud.pulumi.ovh.Domain.inputs.NamePlanConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamePlanArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamePlanArgs Empty = new NamePlanArgs();

    @Import(name="configurations")
    private @Nullable Output<List<NamePlanConfigurationArgs>> configurations;

    public Optional<Output<List<NamePlanConfigurationArgs>>> configurations() {
        return Optional.ofNullable(this.configurations);
    }

    /**
     * Duration selected for the purchase of the product (defaults to &#34;P1Y&#34;)
     * 
     */
    @Import(name="duration", required=true)
    private Output<String> duration;

    /**
     * @return Duration selected for the purchase of the product (defaults to &#34;P1Y&#34;)
     * 
     */
    public Output<String> duration() {
        return this.duration;
    }

    /**
     * Cart item to be linked
     * 
     */
    @Import(name="itemId")
    private @Nullable Output<Double> itemId;

    /**
     * @return Cart item to be linked
     * 
     */
    public Optional<Output<Double>> itemId() {
        return Optional.ofNullable(this.itemId);
    }

    /**
     * Identifier of the option offer
     * 
     */
    @Import(name="planCode", required=true)
    private Output<String> planCode;

    /**
     * @return Identifier of the option offer
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }

    /**
     * Pricing mode selected for the purchase of the product
     * 
     */
    @Import(name="pricingMode", required=true)
    private Output<String> pricingMode;

    /**
     * @return Pricing mode selected for the purchase of the product
     * 
     */
    public Output<String> pricingMode() {
        return this.pricingMode;
    }

    /**
     * Quantity of product desired
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<Double> quantity;

    /**
     * @return Quantity of product desired
     * 
     */
    public Optional<Output<Double>> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    private NamePlanArgs() {}

    private NamePlanArgs(NamePlanArgs $) {
        this.configurations = $.configurations;
        this.duration = $.duration;
        this.itemId = $.itemId;
        this.planCode = $.planCode;
        this.pricingMode = $.pricingMode;
        this.quantity = $.quantity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamePlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamePlanArgs $;

        public Builder() {
            $ = new NamePlanArgs();
        }

        public Builder(NamePlanArgs defaults) {
            $ = new NamePlanArgs(Objects.requireNonNull(defaults));
        }

        public Builder configurations(@Nullable Output<List<NamePlanConfigurationArgs>> configurations) {
            $.configurations = configurations;
            return this;
        }

        public Builder configurations(List<NamePlanConfigurationArgs> configurations) {
            return configurations(Output.of(configurations));
        }

        public Builder configurations(NamePlanConfigurationArgs... configurations) {
            return configurations(List.of(configurations));
        }

        /**
         * @param duration Duration selected for the purchase of the product (defaults to &#34;P1Y&#34;)
         * 
         * @return builder
         * 
         */
        public Builder duration(Output<String> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration Duration selected for the purchase of the product (defaults to &#34;P1Y&#34;)
         * 
         * @return builder
         * 
         */
        public Builder duration(String duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param itemId Cart item to be linked
         * 
         * @return builder
         * 
         */
        public Builder itemId(@Nullable Output<Double> itemId) {
            $.itemId = itemId;
            return this;
        }

        /**
         * @param itemId Cart item to be linked
         * 
         * @return builder
         * 
         */
        public Builder itemId(Double itemId) {
            return itemId(Output.of(itemId));
        }

        /**
         * @param planCode Identifier of the option offer
         * 
         * @return builder
         * 
         */
        public Builder planCode(Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Identifier of the option offer
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param pricingMode Pricing mode selected for the purchase of the product
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(Output<String> pricingMode) {
            $.pricingMode = pricingMode;
            return this;
        }

        /**
         * @param pricingMode Pricing mode selected for the purchase of the product
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(String pricingMode) {
            return pricingMode(Output.of(pricingMode));
        }

        /**
         * @param quantity Quantity of product desired
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<Double> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity Quantity of product desired
         * 
         * @return builder
         * 
         */
        public Builder quantity(Double quantity) {
            return quantity(Output.of(quantity));
        }

        public NamePlanArgs build() {
            if ($.duration == null) {
                throw new MissingRequiredPropertyException("NamePlanArgs", "duration");
            }
            if ($.planCode == null) {
                throw new MissingRequiredPropertyException("NamePlanArgs", "planCode");
            }
            if ($.pricingMode == null) {
                throw new MissingRequiredPropertyException("NamePlanArgs", "pricingMode");
            }
            return $;
        }
    }

}
