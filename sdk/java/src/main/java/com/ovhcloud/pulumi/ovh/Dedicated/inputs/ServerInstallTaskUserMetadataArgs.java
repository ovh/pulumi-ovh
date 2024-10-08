// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServerInstallTaskUserMetadataArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerInstallTaskUserMetadataArgs Empty = new ServerInstallTaskUserMetadataArgs();

    /**
     * The key for the user_metadata
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The key for the user_metadata
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The value for the user_metadata
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return The value for the user_metadata
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ServerInstallTaskUserMetadataArgs() {}

    private ServerInstallTaskUserMetadataArgs(ServerInstallTaskUserMetadataArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerInstallTaskUserMetadataArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerInstallTaskUserMetadataArgs $;

        public Builder() {
            $ = new ServerInstallTaskUserMetadataArgs();
        }

        public Builder(ServerInstallTaskUserMetadataArgs defaults) {
            $ = new ServerInstallTaskUserMetadataArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key for the user_metadata
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The key for the user_metadata
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value for the user_metadata
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value for the user_metadata
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServerInstallTaskUserMetadataArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("ServerInstallTaskUserMetadataArgs", "key");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ServerInstallTaskUserMetadataArgs", "value");
            }
            return $;
        }
    }

}
