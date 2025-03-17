// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class LogsRolePermissionStreamArgs extends com.pulumi.resources.ResourceArgs {

    public static final LogsRolePermissionStreamArgs Empty = new LogsRolePermissionStreamArgs();

    /**
     * The DBaaS Logs role id
     * 
     */
    @Import(name="roleId", required=true)
    private Output<String> roleId;

    /**
     * @return The DBaaS Logs role id
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     * The service name
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * The DBaaS Logs Graylog output stream id
     * 
     */
    @Import(name="streamId", required=true)
    private Output<String> streamId;

    /**
     * @return The DBaaS Logs Graylog output stream id
     * 
     */
    public Output<String> streamId() {
        return this.streamId;
    }

    private LogsRolePermissionStreamArgs() {}

    private LogsRolePermissionStreamArgs(LogsRolePermissionStreamArgs $) {
        this.roleId = $.roleId;
        this.serviceName = $.serviceName;
        this.streamId = $.streamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsRolePermissionStreamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsRolePermissionStreamArgs $;

        public Builder() {
            $ = new LogsRolePermissionStreamArgs();
        }

        public Builder(LogsRolePermissionStreamArgs defaults) {
            $ = new LogsRolePermissionStreamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param roleId The DBaaS Logs role id
         * 
         * @return builder
         * 
         */
        public Builder roleId(Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The DBaaS Logs role id
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param streamId The DBaaS Logs Graylog output stream id
         * 
         * @return builder
         * 
         */
        public Builder streamId(Output<String> streamId) {
            $.streamId = streamId;
            return this;
        }

        /**
         * @param streamId The DBaaS Logs Graylog output stream id
         * 
         * @return builder
         * 
         */
        public Builder streamId(String streamId) {
            return streamId(Output.of(streamId));
        }

        public LogsRolePermissionStreamArgs build() {
            if ($.roleId == null) {
                throw new MissingRequiredPropertyException("LogsRolePermissionStreamArgs", "roleId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("LogsRolePermissionStreamArgs", "serviceName");
            }
            if ($.streamId == null) {
                throw new MissingRequiredPropertyException("LogsRolePermissionStreamArgs", "streamId");
            }
            return $;
        }
    }

}
