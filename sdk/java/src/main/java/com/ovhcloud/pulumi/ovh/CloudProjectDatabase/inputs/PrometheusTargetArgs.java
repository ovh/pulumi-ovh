// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrometheusTargetArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrometheusTargetArgs Empty = new PrometheusTargetArgs();

    /**
     * Host of the endpoint
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return Host of the endpoint
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Connection port for the endpoint
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Connection port for the endpoint
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private PrometheusTargetArgs() {}

    private PrometheusTargetArgs(PrometheusTargetArgs $) {
        this.host = $.host;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrometheusTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrometheusTargetArgs $;

        public Builder() {
            $ = new PrometheusTargetArgs();
        }

        public Builder(PrometheusTargetArgs defaults) {
            $ = new PrometheusTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param host Host of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Host of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param port Connection port for the endpoint
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Connection port for the endpoint
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public PrometheusTargetArgs build() {
            return $;
        }
    }

}
