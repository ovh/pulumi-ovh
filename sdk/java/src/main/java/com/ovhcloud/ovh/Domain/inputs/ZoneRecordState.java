// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Domain.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneRecordState extends com.pulumi.resources.ResourceArgs {

    public static final ZoneRecordState Empty = new ZoneRecordState();

    /**
     * The type of the record
     * 
     */
    @Import(name="fieldtype")
    private @Nullable Output<String> fieldtype;

    /**
     * @return The type of the record
     * 
     */
    public Optional<Output<String>> fieldtype() {
        return Optional.ofNullable(this.fieldtype);
    }

    /**
     * The name of the record. It can be an empty string.
     * 
     */
    @Import(name="subdomain")
    private @Nullable Output<String> subdomain;

    /**
     * @return The name of the record. It can be an empty string.
     * 
     */
    public Optional<Output<String>> subdomain() {
        return Optional.ofNullable(this.subdomain);
    }

    /**
     * The value of the record
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return The value of the record
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * The TTL of the record, it shall be &gt;= to 60.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The TTL of the record, it shall be &gt;= to 60.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * The domain to add the record to
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return The domain to add the record to
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private ZoneRecordState() {}

    private ZoneRecordState(ZoneRecordState $) {
        this.fieldtype = $.fieldtype;
        this.subdomain = $.subdomain;
        this.target = $.target;
        this.ttl = $.ttl;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneRecordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneRecordState $;

        public Builder() {
            $ = new ZoneRecordState();
        }

        public Builder(ZoneRecordState defaults) {
            $ = new ZoneRecordState(Objects.requireNonNull(defaults));
        }

        /**
         * @param fieldtype The type of the record
         * 
         * @return builder
         * 
         */
        public Builder fieldtype(@Nullable Output<String> fieldtype) {
            $.fieldtype = fieldtype;
            return this;
        }

        /**
         * @param fieldtype The type of the record
         * 
         * @return builder
         * 
         */
        public Builder fieldtype(String fieldtype) {
            return fieldtype(Output.of(fieldtype));
        }

        /**
         * @param subdomain The name of the record. It can be an empty string.
         * 
         * @return builder
         * 
         */
        public Builder subdomain(@Nullable Output<String> subdomain) {
            $.subdomain = subdomain;
            return this;
        }

        /**
         * @param subdomain The name of the record. It can be an empty string.
         * 
         * @return builder
         * 
         */
        public Builder subdomain(String subdomain) {
            return subdomain(Output.of(subdomain));
        }

        /**
         * @param target The value of the record
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target The value of the record
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        /**
         * @param ttl The TTL of the record, it shall be &gt;= to 60.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The TTL of the record, it shall be &gt;= to 60.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param zone The domain to add the record to
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The domain to add the record to
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public ZoneRecordState build() {
            return $;
        }
    }

}