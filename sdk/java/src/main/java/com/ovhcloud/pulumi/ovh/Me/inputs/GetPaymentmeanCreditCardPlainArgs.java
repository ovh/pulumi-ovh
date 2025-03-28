// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPaymentmeanCreditCardPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPaymentmeanCreditCardPlainArgs Empty = new GetPaymentmeanCreditCardPlainArgs();

    /**
     * a regexp used to filter credit cards on their `description` attributes.
     * 
     */
    @Import(name="descriptionRegexp")
    private @Nullable String descriptionRegexp;

    /**
     * @return a regexp used to filter credit cards on their `description` attributes.
     * 
     */
    public Optional<String> descriptionRegexp() {
        return Optional.ofNullable(this.descriptionRegexp);
    }

    /**
     * Filter credit cards on their `state` attribute. Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
     * 
     */
    @Import(name="states")
    private @Nullable List<String> states;

    /**
     * @return Filter credit cards on their `state` attribute. Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
     * 
     */
    public Optional<List<String>> states() {
        return Optional.ofNullable(this.states);
    }

    /**
     * Retrieve credit card marked as default payment mean.
     * 
     */
    @Import(name="useDefault")
    private @Nullable Boolean useDefault;

    /**
     * @return Retrieve credit card marked as default payment mean.
     * 
     */
    public Optional<Boolean> useDefault() {
        return Optional.ofNullable(this.useDefault);
    }

    /**
     * Retrieve the credit card that will be the last to expire according to its expiration date.
     * 
     */
    @Import(name="useLastToExpire")
    private @Nullable Boolean useLastToExpire;

    /**
     * @return Retrieve the credit card that will be the last to expire according to its expiration date.
     * 
     */
    public Optional<Boolean> useLastToExpire() {
        return Optional.ofNullable(this.useLastToExpire);
    }

    private GetPaymentmeanCreditCardPlainArgs() {}

    private GetPaymentmeanCreditCardPlainArgs(GetPaymentmeanCreditCardPlainArgs $) {
        this.descriptionRegexp = $.descriptionRegexp;
        this.states = $.states;
        this.useDefault = $.useDefault;
        this.useLastToExpire = $.useLastToExpire;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPaymentmeanCreditCardPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPaymentmeanCreditCardPlainArgs $;

        public Builder() {
            $ = new GetPaymentmeanCreditCardPlainArgs();
        }

        public Builder(GetPaymentmeanCreditCardPlainArgs defaults) {
            $ = new GetPaymentmeanCreditCardPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptionRegexp a regexp used to filter credit cards on their `description` attributes.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegexp(@Nullable String descriptionRegexp) {
            $.descriptionRegexp = descriptionRegexp;
            return this;
        }

        /**
         * @param states Filter credit cards on their `state` attribute. Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
         * 
         * @return builder
         * 
         */
        public Builder states(@Nullable List<String> states) {
            $.states = states;
            return this;
        }

        /**
         * @param states Filter credit cards on their `state` attribute. Can be &#34;expired&#34;, &#34;valid&#34;, &#34;tooManyFailures&#34;
         * 
         * @return builder
         * 
         */
        public Builder states(String... states) {
            return states(List.of(states));
        }

        /**
         * @param useDefault Retrieve credit card marked as default payment mean.
         * 
         * @return builder
         * 
         */
        public Builder useDefault(@Nullable Boolean useDefault) {
            $.useDefault = useDefault;
            return this;
        }

        /**
         * @param useLastToExpire Retrieve the credit card that will be the last to expire according to its expiration date.
         * 
         * @return builder
         * 
         */
        public Builder useLastToExpire(@Nullable Boolean useLastToExpire) {
            $.useLastToExpire = useLastToExpire;
            return this;
        }

        public GetPaymentmeanCreditCardPlainArgs build() {
            return $;
        }
    }

}
