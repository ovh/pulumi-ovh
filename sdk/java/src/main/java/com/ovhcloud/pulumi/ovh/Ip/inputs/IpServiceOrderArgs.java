// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Ip.inputs;

import com.ovhcloud.pulumi.ovh.Ip.inputs.IpServiceOrderDetailArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpServiceOrderArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpServiceOrderArgs Empty = new IpServiceOrderArgs();

    /**
     * date
     * 
     */
    @Import(name="date")
    private @Nullable Output<String> date;

    /**
     * @return date
     * 
     */
    public Optional<Output<String>> date() {
        return Optional.ofNullable(this.date);
    }

    /**
     * Information about a Bill entry
     * 
     */
    @Import(name="details")
    private @Nullable Output<List<IpServiceOrderDetailArgs>> details;

    /**
     * @return Information about a Bill entry
     * 
     */
    public Optional<Output<List<IpServiceOrderDetailArgs>>> details() {
        return Optional.ofNullable(this.details);
    }

    /**
     * expiration date
     * 
     */
    @Import(name="expirationDate")
    private @Nullable Output<String> expirationDate;

    /**
     * @return expiration date
     * 
     */
    public Optional<Output<String>> expirationDate() {
        return Optional.ofNullable(this.expirationDate);
    }

    /**
     * order id
     * 
     */
    @Import(name="orderId")
    private @Nullable Output<Integer> orderId;

    /**
     * @return order id
     * 
     */
    public Optional<Output<Integer>> orderId() {
        return Optional.ofNullable(this.orderId);
    }

    private IpServiceOrderArgs() {}

    private IpServiceOrderArgs(IpServiceOrderArgs $) {
        this.date = $.date;
        this.details = $.details;
        this.expirationDate = $.expirationDate;
        this.orderId = $.orderId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpServiceOrderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpServiceOrderArgs $;

        public Builder() {
            $ = new IpServiceOrderArgs();
        }

        public Builder(IpServiceOrderArgs defaults) {
            $ = new IpServiceOrderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param date date
         * 
         * @return builder
         * 
         */
        public Builder date(@Nullable Output<String> date) {
            $.date = date;
            return this;
        }

        /**
         * @param date date
         * 
         * @return builder
         * 
         */
        public Builder date(String date) {
            return date(Output.of(date));
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(@Nullable Output<List<IpServiceOrderDetailArgs>> details) {
            $.details = details;
            return this;
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(List<IpServiceOrderDetailArgs> details) {
            return details(Output.of(details));
        }

        /**
         * @param details Information about a Bill entry
         * 
         * @return builder
         * 
         */
        public Builder details(IpServiceOrderDetailArgs... details) {
            return details(List.of(details));
        }

        /**
         * @param expirationDate expiration date
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(@Nullable Output<String> expirationDate) {
            $.expirationDate = expirationDate;
            return this;
        }

        /**
         * @param expirationDate expiration date
         * 
         * @return builder
         * 
         */
        public Builder expirationDate(String expirationDate) {
            return expirationDate(Output.of(expirationDate));
        }

        /**
         * @param orderId order id
         * 
         * @return builder
         * 
         */
        public Builder orderId(@Nullable Output<Integer> orderId) {
            $.orderId = orderId;
            return this;
        }

        /**
         * @param orderId order id
         * 
         * @return builder
         * 
         */
        public Builder orderId(Integer orderId) {
            return orderId(Output.of(orderId));
        }

        public IpServiceOrderArgs build() {
            return $;
        }
    }

}
