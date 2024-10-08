// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceKeyState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceKeyState Empty = new ServiceKeyState();

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
     * Creation time of the key
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Creation time of the key
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
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
     * Key deactivation reason
     * 
     */
    @Import(name="deactivationReason")
    private @Nullable Output<String> deactivationReason;

    /**
     * @return Key deactivation reason
     * 
     */
    public Optional<Output<String>> deactivationReason() {
        return Optional.ofNullable(this.deactivationReason);
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
    @Import(name="okmsId")
    private @Nullable Output<String> okmsId;

    /**
     * @return Okms ID
     * 
     */
    public Optional<Output<String>> okmsId() {
        return Optional.ofNullable(this.okmsId);
    }

    /**
     * The operations for which the key is intended to be used
     * 
     */
    @Import(name="operations")
    private @Nullable Output<List<String>> operations;

    /**
     * @return The operations for which the key is intended to be used
     * 
     */
    public Optional<Output<List<String>>> operations() {
        return Optional.ofNullable(this.operations);
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
     * State of the key
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return State of the key
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * Type of the key to be created
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of the key to be created
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ServiceKeyState() {}

    private ServiceKeyState(ServiceKeyState $) {
        this.context = $.context;
        this.createdAt = $.createdAt;
        this.curve = $.curve;
        this.deactivationReason = $.deactivationReason;
        this.name = $.name;
        this.okmsId = $.okmsId;
        this.operations = $.operations;
        this.size = $.size;
        this.state = $.state;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceKeyState $;

        public Builder() {
            $ = new ServiceKeyState();
        }

        public Builder(ServiceKeyState defaults) {
            $ = new ServiceKeyState(Objects.requireNonNull(defaults));
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
         * @param createdAt Creation time of the key
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Creation time of the key
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
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
         * @param deactivationReason Key deactivation reason
         * 
         * @return builder
         * 
         */
        public Builder deactivationReason(@Nullable Output<String> deactivationReason) {
            $.deactivationReason = deactivationReason;
            return this;
        }

        /**
         * @param deactivationReason Key deactivation reason
         * 
         * @return builder
         * 
         */
        public Builder deactivationReason(String deactivationReason) {
            return deactivationReason(Output.of(deactivationReason));
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
        public Builder okmsId(@Nullable Output<String> okmsId) {
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
        public Builder operations(@Nullable Output<List<String>> operations) {
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
         * @param state State of the key
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state State of the key
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param type Type of the key to be created
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
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

        public ServiceKeyState build() {
            return $;
        }
    }

}
