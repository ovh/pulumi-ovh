// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class OVHcloudConnectArgs extends com.pulumi.resources.ResourceArgs {

    public static final OVHcloudConnectArgs Empty = new OVHcloudConnectArgs();

    /**
     * Your OVH Cloud Connect service name.
     * 
     */
    @Import(name="ovhCloudConnect", required=true)
    private Output<String> ovhCloudConnect;

    /**
     * @return Your OVH Cloud Connect service name.
     * 
     */
    public Output<String> ovhCloudConnect() {
        return this.ovhCloudConnect;
    }

    /**
     * The internal name of your vrack
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your vrack
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private OVHcloudConnectArgs() {}

    private OVHcloudConnectArgs(OVHcloudConnectArgs $) {
        this.ovhCloudConnect = $.ovhCloudConnect;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OVHcloudConnectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OVHcloudConnectArgs $;

        public Builder() {
            $ = new OVHcloudConnectArgs();
        }

        public Builder(OVHcloudConnectArgs defaults) {
            $ = new OVHcloudConnectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ovhCloudConnect Your OVH Cloud Connect service name.
         * 
         * @return builder
         * 
         */
        public Builder ovhCloudConnect(Output<String> ovhCloudConnect) {
            $.ovhCloudConnect = ovhCloudConnect;
            return this;
        }

        /**
         * @param ovhCloudConnect Your OVH Cloud Connect service name.
         * 
         * @return builder
         * 
         */
        public Builder ovhCloudConnect(String ovhCloudConnect) {
            return ovhCloudConnect(Output.of(ovhCloudConnect));
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public OVHcloudConnectArgs build() {
            if ($.ovhCloudConnect == null) {
                throw new MissingRequiredPropertyException("OVHcloudConnectArgs", "ovhCloudConnect");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("OVHcloudConnectArgs", "serviceName");
            }
            return $;
        }
    }

}
