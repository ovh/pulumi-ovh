// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.LoadBalancerListenerPoolHealthMonitorHttpConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerListenerPoolHealthMonitorArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerListenerPoolHealthMonitorArgs Empty = new LoadBalancerListenerPoolHealthMonitorArgs();

    /**
     * Duration between sending probes to members, in seconds
     * 
     */
    @Import(name="delay")
    private @Nullable Output<Double> delay;

    /**
     * @return Duration between sending probes to members, in seconds
     * 
     */
    public Optional<Output<Double>> delay() {
        return Optional.ofNullable(this.delay);
    }

    /**
     * Monitor HTTP configuration
     * 
     */
    @Import(name="httpConfiguration")
    private @Nullable Output<LoadBalancerListenerPoolHealthMonitorHttpConfigurationArgs> httpConfiguration;

    /**
     * @return Monitor HTTP configuration
     * 
     */
    public Optional<Output<LoadBalancerListenerPoolHealthMonitorHttpConfigurationArgs>> httpConfiguration() {
        return Optional.ofNullable(this.httpConfiguration);
    }

    /**
     * Number of successful checks before changing the operating status of the member to ONLINE
     * 
     */
    @Import(name="maxRetries")
    private @Nullable Output<Double> maxRetries;

    /**
     * @return Number of successful checks before changing the operating status of the member to ONLINE
     * 
     */
    public Optional<Output<Double>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * Number of allowed check failures before changing the operating status of the member to ERROR
     * 
     */
    @Import(name="maxRetriesDown")
    private @Nullable Output<Double> maxRetriesDown;

    /**
     * @return Number of allowed check failures before changing the operating status of the member to ERROR
     * 
     */
    public Optional<Output<Double>> maxRetriesDown() {
        return Optional.ofNullable(this.maxRetriesDown);
    }

    /**
     * Type of the monitor
     * 
     */
    @Import(name="monitorType")
    private @Nullable Output<String> monitorType;

    /**
     * @return Type of the monitor
     * 
     */
    public Optional<Output<String>> monitorType() {
        return Optional.ofNullable(this.monitorType);
    }

    /**
     * The name of the resource
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the resource
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The operating status of the resource
     * 
     */
    @Import(name="operatingStatus")
    private @Nullable Output<String> operatingStatus;

    /**
     * @return The operating status of the resource
     * 
     */
    public Optional<Output<String>> operatingStatus() {
        return Optional.ofNullable(this.operatingStatus);
    }

    /**
     * The provisioning status of the resource
     * 
     */
    @Import(name="provisioningStatus")
    private @Nullable Output<String> provisioningStatus;

    /**
     * @return The provisioning status of the resource
     * 
     */
    public Optional<Output<String>> provisioningStatus() {
        return Optional.ofNullable(this.provisioningStatus);
    }

    /**
     * Maximum time, in seconds, that a monitor waits to connect before it times out. This value must be less than the delay value
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Double> timeout;

    /**
     * @return Maximum time, in seconds, that a monitor waits to connect before it times out. This value must be less than the delay value
     * 
     */
    public Optional<Output<Double>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private LoadBalancerListenerPoolHealthMonitorArgs() {}

    private LoadBalancerListenerPoolHealthMonitorArgs(LoadBalancerListenerPoolHealthMonitorArgs $) {
        this.delay = $.delay;
        this.httpConfiguration = $.httpConfiguration;
        this.maxRetries = $.maxRetries;
        this.maxRetriesDown = $.maxRetriesDown;
        this.monitorType = $.monitorType;
        this.name = $.name;
        this.operatingStatus = $.operatingStatus;
        this.provisioningStatus = $.provisioningStatus;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerListenerPoolHealthMonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerListenerPoolHealthMonitorArgs $;

        public Builder() {
            $ = new LoadBalancerListenerPoolHealthMonitorArgs();
        }

        public Builder(LoadBalancerListenerPoolHealthMonitorArgs defaults) {
            $ = new LoadBalancerListenerPoolHealthMonitorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param delay Duration between sending probes to members, in seconds
         * 
         * @return builder
         * 
         */
        public Builder delay(@Nullable Output<Double> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay Duration between sending probes to members, in seconds
         * 
         * @return builder
         * 
         */
        public Builder delay(Double delay) {
            return delay(Output.of(delay));
        }

        /**
         * @param httpConfiguration Monitor HTTP configuration
         * 
         * @return builder
         * 
         */
        public Builder httpConfiguration(@Nullable Output<LoadBalancerListenerPoolHealthMonitorHttpConfigurationArgs> httpConfiguration) {
            $.httpConfiguration = httpConfiguration;
            return this;
        }

        /**
         * @param httpConfiguration Monitor HTTP configuration
         * 
         * @return builder
         * 
         */
        public Builder httpConfiguration(LoadBalancerListenerPoolHealthMonitorHttpConfigurationArgs httpConfiguration) {
            return httpConfiguration(Output.of(httpConfiguration));
        }

        /**
         * @param maxRetries Number of successful checks before changing the operating status of the member to ONLINE
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Double> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries Number of successful checks before changing the operating status of the member to ONLINE
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Double maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param maxRetriesDown Number of allowed check failures before changing the operating status of the member to ERROR
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesDown(@Nullable Output<Double> maxRetriesDown) {
            $.maxRetriesDown = maxRetriesDown;
            return this;
        }

        /**
         * @param maxRetriesDown Number of allowed check failures before changing the operating status of the member to ERROR
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesDown(Double maxRetriesDown) {
            return maxRetriesDown(Output.of(maxRetriesDown));
        }

        /**
         * @param monitorType Type of the monitor
         * 
         * @return builder
         * 
         */
        public Builder monitorType(@Nullable Output<String> monitorType) {
            $.monitorType = monitorType;
            return this;
        }

        /**
         * @param monitorType Type of the monitor
         * 
         * @return builder
         * 
         */
        public Builder monitorType(String monitorType) {
            return monitorType(Output.of(monitorType));
        }

        /**
         * @param name The name of the resource
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the resource
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param operatingStatus The operating status of the resource
         * 
         * @return builder
         * 
         */
        public Builder operatingStatus(@Nullable Output<String> operatingStatus) {
            $.operatingStatus = operatingStatus;
            return this;
        }

        /**
         * @param operatingStatus The operating status of the resource
         * 
         * @return builder
         * 
         */
        public Builder operatingStatus(String operatingStatus) {
            return operatingStatus(Output.of(operatingStatus));
        }

        /**
         * @param provisioningStatus The provisioning status of the resource
         * 
         * @return builder
         * 
         */
        public Builder provisioningStatus(@Nullable Output<String> provisioningStatus) {
            $.provisioningStatus = provisioningStatus;
            return this;
        }

        /**
         * @param provisioningStatus The provisioning status of the resource
         * 
         * @return builder
         * 
         */
        public Builder provisioningStatus(String provisioningStatus) {
            return provisioningStatus(Output.of(provisioningStatus));
        }

        /**
         * @param timeout Maximum time, in seconds, that a monitor waits to connect before it times out. This value must be less than the delay value
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Double> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Maximum time, in seconds, that a monitor waits to connect before it times out. This value must be less than the delay value
         * 
         * @return builder
         * 
         */
        public Builder timeout(Double timeout) {
            return timeout(Output.of(timeout));
        }

        public LoadBalancerListenerPoolHealthMonitorArgs build() {
            return $;
        }
    }

}
