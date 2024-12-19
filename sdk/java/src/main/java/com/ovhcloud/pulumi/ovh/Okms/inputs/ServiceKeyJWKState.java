// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.inputs;

import com.ovhcloud.pulumi.ovh.Okms.inputs.ServiceKeyJWKIamArgs;
import com.ovhcloud.pulumi.ovh.Okms.inputs.ServiceKeyJWKKeyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceKeyJWKState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceKeyJWKState Empty = new ServiceKeyJWKState();

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
     * IAM resource metadata
     * 
     */
    @Import(name="iam")
    private @Nullable Output<ServiceKeyJWKIamArgs> iam;

    /**
     * @return IAM resource metadata
     * 
     */
    public Optional<Output<ServiceKeyJWKIamArgs>> iam() {
        return Optional.ofNullable(this.iam);
    }

    /**
     * Set of JSON Web Keys to import
     * 
     */
    @Import(name="keys")
    private @Nullable Output<List<ServiceKeyJWKKeyArgs>> keys;

    /**
     * @return Set of JSON Web Keys to import
     * 
     */
    public Optional<Output<List<ServiceKeyJWKKeyArgs>>> keys() {
        return Optional.ofNullable(this.keys);
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

    private ServiceKeyJWKState() {}

    private ServiceKeyJWKState(ServiceKeyJWKState $) {
        this.context = $.context;
        this.createdAt = $.createdAt;
        this.deactivationReason = $.deactivationReason;
        this.iam = $.iam;
        this.keys = $.keys;
        this.name = $.name;
        this.okmsId = $.okmsId;
        this.size = $.size;
        this.state = $.state;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceKeyJWKState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceKeyJWKState $;

        public Builder() {
            $ = new ServiceKeyJWKState();
        }

        public Builder(ServiceKeyJWKState defaults) {
            $ = new ServiceKeyJWKState(Objects.requireNonNull(defaults));
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
         * @param iam IAM resource metadata
         * 
         * @return builder
         * 
         */
        public Builder iam(@Nullable Output<ServiceKeyJWKIamArgs> iam) {
            $.iam = iam;
            return this;
        }

        /**
         * @param iam IAM resource metadata
         * 
         * @return builder
         * 
         */
        public Builder iam(ServiceKeyJWKIamArgs iam) {
            return iam(Output.of(iam));
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(@Nullable Output<List<ServiceKeyJWKKeyArgs>> keys) {
            $.keys = keys;
            return this;
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(List<ServiceKeyJWKKeyArgs> keys) {
            return keys(Output.of(keys));
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(ServiceKeyJWKKeyArgs... keys) {
            return keys(List.of(keys));
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

        public ServiceKeyJWKState build() {
            return $;
        }
    }

}
