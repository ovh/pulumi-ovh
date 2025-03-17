// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NameCurrentTaskArgs extends com.pulumi.resources.ResourceArgs {

    public static final NameCurrentTaskArgs Empty = new NameCurrentTaskArgs();

    /**
     * Identifier of the current task
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return Identifier of the current task
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Link to the task details
     * 
     */
    @Import(name="link")
    private @Nullable Output<String> link;

    /**
     * @return Link to the task details
     * 
     */
    public Optional<Output<String>> link() {
        return Optional.ofNullable(this.link);
    }

    /**
     * Current global status of the current task
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Current global status of the current task
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Type of the current task
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of the current task
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private NameCurrentTaskArgs() {}

    private NameCurrentTaskArgs(NameCurrentTaskArgs $) {
        this.id = $.id;
        this.link = $.link;
        this.status = $.status;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NameCurrentTaskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NameCurrentTaskArgs $;

        public Builder() {
            $ = new NameCurrentTaskArgs();
        }

        public Builder(NameCurrentTaskArgs defaults) {
            $ = new NameCurrentTaskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Identifier of the current task
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Identifier of the current task
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param link Link to the task details
         * 
         * @return builder
         * 
         */
        public Builder link(@Nullable Output<String> link) {
            $.link = link;
            return this;
        }

        /**
         * @param link Link to the task details
         * 
         * @return builder
         * 
         */
        public Builder link(String link) {
            return link(Output.of(link));
        }

        /**
         * @param status Current global status of the current task
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Current global status of the current task
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param type Type of the current task
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the current task
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NameCurrentTaskArgs build() {
            return $;
        }
    }

}
