// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertingFormattedMonthlyThresholdArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertingFormattedMonthlyThresholdArgs Empty = new AlertingFormattedMonthlyThresholdArgs();

    /**
     * Currency of the monthly threshold
     * 
     */
    @Import(name="currencyCode")
    private @Nullable Output<String> currencyCode;

    /**
     * @return Currency of the monthly threshold
     * 
     */
    public Optional<Output<String>> currencyCode() {
        return Optional.ofNullable(this.currencyCode);
    }

    /**
     * Text representation of the monthly threshold
     * 
     */
    @Import(name="text")
    private @Nullable Output<String> text;

    /**
     * @return Text representation of the monthly threshold
     * 
     */
    public Optional<Output<String>> text() {
        return Optional.ofNullable(this.text);
    }

    /**
     * Value of the monthly threshold
     * 
     */
    @Import(name="value")
    private @Nullable Output<Double> value;

    /**
     * @return Value of the monthly threshold
     * 
     */
    public Optional<Output<Double>> value() {
        return Optional.ofNullable(this.value);
    }

    private AlertingFormattedMonthlyThresholdArgs() {}

    private AlertingFormattedMonthlyThresholdArgs(AlertingFormattedMonthlyThresholdArgs $) {
        this.currencyCode = $.currencyCode;
        this.text = $.text;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertingFormattedMonthlyThresholdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertingFormattedMonthlyThresholdArgs $;

        public Builder() {
            $ = new AlertingFormattedMonthlyThresholdArgs();
        }

        public Builder(AlertingFormattedMonthlyThresholdArgs defaults) {
            $ = new AlertingFormattedMonthlyThresholdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param currencyCode Currency of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder currencyCode(@Nullable Output<String> currencyCode) {
            $.currencyCode = currencyCode;
            return this;
        }

        /**
         * @param currencyCode Currency of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder currencyCode(String currencyCode) {
            return currencyCode(Output.of(currencyCode));
        }

        /**
         * @param text Text representation of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder text(@Nullable Output<String> text) {
            $.text = text;
            return this;
        }

        /**
         * @param text Text representation of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder text(String text) {
            return text(Output.of(text));
        }

        /**
         * @param value Value of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<Double> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value of the monthly threshold
         * 
         * @return builder
         * 
         */
        public Builder value(Double value) {
            return value(Output.of(value));
        }

        public AlertingFormattedMonthlyThresholdArgs build() {
            return $;
        }
    }

}