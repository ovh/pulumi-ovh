// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Ip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class MoveRoutedToArgs extends com.pulumi.resources.ResourceArgs {

    public static final MoveRoutedToArgs Empty = new MoveRoutedToArgs();

    /**
     * Name of the service to route the IP to. IP will be parked if this value is an empty string
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return Name of the service to route the IP to. IP will be parked if this value is an empty string
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private MoveRoutedToArgs() {}

    private MoveRoutedToArgs(MoveRoutedToArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MoveRoutedToArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MoveRoutedToArgs $;

        public Builder() {
            $ = new MoveRoutedToArgs();
        }

        public Builder(MoveRoutedToArgs defaults) {
            $ = new MoveRoutedToArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName Name of the service to route the IP to. IP will be parked if this value is an empty string
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Name of the service to route the IP to. IP will be parked if this value is an empty string
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public MoveRoutedToArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("MoveRoutedToArgs", "serviceName");
            }
            return $;
        }
    }

}