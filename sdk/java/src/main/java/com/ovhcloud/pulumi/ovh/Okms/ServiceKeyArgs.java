// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceKeyArgs Empty = new ServiceKeyArgs();

    /**
     * Context of the key
     * 
     */
    @Import(name="context")
    private @Nullable Output<String> context;

    /**
     * @return Context of the key
     * 
     */
    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    /**
     * Curve type for Elliptic Curve (EC) keys
     * 
     */
    @Import(name="curve")
    private @Nullable Output<String> curve;

    /**
     * @return Curve type for Elliptic Curve (EC) keys
     * 
     */
    public Optional<Output<String>> curve() {
        return Optional.ofNullable(this.curve);
    }

    /**
     * Key name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Key name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Okms ID
     * 
     */
    @Import(name="okmsId", required=true)
    private Output<String> okmsId;

    /**
     * @return Okms ID
     * 
     */
    public Output<String> okmsId() {
        return this.okmsId;
    }

    /**
     * The operations for which the key is intended to be used
     * 
     */
    @Import(name="operations", required=true)
    private Output<List<String>> operations;

    /**
     * @return The operations for which the key is intended to be used
     * 
     */
    public Output<List<String>> operations() {
        return this.operations;
    }

    /**
     * Size of the key to be created
     * 
     */
    @Import(name="size")
    private @Nullable Output<Double> size;

    /**
     * @return Size of the key to be created
     * 
     */
    public Optional<Output<Double>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * Type of the key to be created
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of the key to be created
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ServiceKeyArgs() {}

    private ServiceKeyArgs(ServiceKeyArgs $) {
        this.context = $.context;
        this.curve = $.curve;
        this.name = $.name;
        this.okmsId = $.okmsId;
        this.operations = $.operations;
        this.size = $.size;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceKeyArgs $;

        public Builder() {
            $ = new ServiceKeyArgs();
        }

        public Builder(ServiceKeyArgs defaults) {
            $ = new ServiceKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param context Context of the key
         * 
         * @return builder
         * 
         */
        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context Context of the key
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param curve Curve type for Elliptic Curve (EC) keys
         * 
         * @return builder
         * 
         */
        public Builder curve(@Nullable Output<String> curve) {
            $.curve = curve;
            return this;
        }

        /**
         * @param curve Curve type for Elliptic Curve (EC) keys
         * 
         * @return builder
         * 
         */
        public Builder curve(String curve) {
            return curve(Output.of(curve));
        }

        /**
         * @param name Key name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Key name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param okmsId Okms ID
         * 
         * @return builder
         * 
         */
        public Builder okmsId(Output<String> okmsId) {
            $.okmsId = okmsId;
            return this;
        }

        /**
         * @param okmsId Okms ID
         * 
         * @return builder
         * 
         */
        public Builder okmsId(String okmsId) {
            return okmsId(Output.of(okmsId));
        }

        /**
         * @param operations The operations for which the key is intended to be used
         * 
         * @return builder
         * 
         */
        public Builder operations(Output<List<String>> operations) {
            $.operations = operations;
            return this;
        }

        /**
         * @param operations The operations for which the key is intended to be used
         * 
         * @return builder
         * 
         */
        public Builder operations(List<String> operations) {
            return operations(Output.of(operations));
        }

        /**
         * @param operations The operations for which the key is intended to be used
         * 
         * @return builder
         * 
         */
        public Builder operations(String... operations) {
            return operations(List.of(operations));
        }

        /**
         * @param size Size of the key to be created
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Double> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of the key to be created
         * 
         * @return builder
         * 
         */
        public Builder size(Double size) {
            return size(Output.of(size));
        }

        /**
         * @param type Type of the key to be created
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the key to be created
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ServiceKeyArgs build() {
            if ($.okmsId == null) {
                throw new MissingRequiredPropertyException("ServiceKeyArgs", "okmsId");
            }
            if ($.operations == null) {
                throw new MissingRequiredPropertyException("ServiceKeyArgs", "operations");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("ServiceKeyArgs", "type");
            }
            return $;
        }
    }

}