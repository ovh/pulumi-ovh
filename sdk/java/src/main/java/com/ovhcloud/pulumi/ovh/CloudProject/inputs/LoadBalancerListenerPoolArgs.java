// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.LoadBalancerListenerPoolHealthMonitorArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.LoadBalancerListenerPoolMemberArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.LoadBalancerListenerPoolSessionPersistenceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerListenerPoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerListenerPoolArgs Empty = new LoadBalancerListenerPoolArgs();

    /**
     * Pool algorithm to split traffic between members
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Pool algorithm to split traffic between members
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * Pool health monitor
     * 
     */
    @Import(name="healthMonitor")
    private @Nullable Output<LoadBalancerListenerPoolHealthMonitorArgs> healthMonitor;

    /**
     * @return Pool health monitor
     * 
     */
    public Optional<Output<LoadBalancerListenerPoolHealthMonitorArgs>> healthMonitor() {
        return Optional.ofNullable(this.healthMonitor);
    }

    /**
     * Pool members
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<LoadBalancerListenerPoolMemberArgs>> members;

    /**
     * @return Pool members
     * 
     */
    public Optional<Output<List<LoadBalancerListenerPoolMemberArgs>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * Name of the pool
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the pool
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Protocol for the pool
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Protocol for the pool
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Pool session persistence
     * 
     */
    @Import(name="sessionPersistence")
    private @Nullable Output<LoadBalancerListenerPoolSessionPersistenceArgs> sessionPersistence;

    /**
     * @return Pool session persistence
     * 
     */
    public Optional<Output<LoadBalancerListenerPoolSessionPersistenceArgs>> sessionPersistence() {
        return Optional.ofNullable(this.sessionPersistence);
    }

    private LoadBalancerListenerPoolArgs() {}

    private LoadBalancerListenerPoolArgs(LoadBalancerListenerPoolArgs $) {
        this.algorithm = $.algorithm;
        this.healthMonitor = $.healthMonitor;
        this.members = $.members;
        this.name = $.name;
        this.protocol = $.protocol;
        this.sessionPersistence = $.sessionPersistence;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerListenerPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerListenerPoolArgs $;

        public Builder() {
            $ = new LoadBalancerListenerPoolArgs();
        }

        public Builder(LoadBalancerListenerPoolArgs defaults) {
            $ = new LoadBalancerListenerPoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Pool algorithm to split traffic between members
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Pool algorithm to split traffic between members
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param healthMonitor Pool health monitor
         * 
         * @return builder
         * 
         */
        public Builder healthMonitor(@Nullable Output<LoadBalancerListenerPoolHealthMonitorArgs> healthMonitor) {
            $.healthMonitor = healthMonitor;
            return this;
        }

        /**
         * @param healthMonitor Pool health monitor
         * 
         * @return builder
         * 
         */
        public Builder healthMonitor(LoadBalancerListenerPoolHealthMonitorArgs healthMonitor) {
            return healthMonitor(Output.of(healthMonitor));
        }

        /**
         * @param members Pool members
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<LoadBalancerListenerPoolMemberArgs>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members Pool members
         * 
         * @return builder
         * 
         */
        public Builder members(List<LoadBalancerListenerPoolMemberArgs> members) {
            return members(Output.of(members));
        }

        /**
         * @param members Pool members
         * 
         * @return builder
         * 
         */
        public Builder members(LoadBalancerListenerPoolMemberArgs... members) {
            return members(List.of(members));
        }

        /**
         * @param name Name of the pool
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the pool
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param protocol Protocol for the pool
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Protocol for the pool
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param sessionPersistence Pool session persistence
         * 
         * @return builder
         * 
         */
        public Builder sessionPersistence(@Nullable Output<LoadBalancerListenerPoolSessionPersistenceArgs> sessionPersistence) {
            $.sessionPersistence = sessionPersistence;
            return this;
        }

        /**
         * @param sessionPersistence Pool session persistence
         * 
         * @return builder
         * 
         */
        public Builder sessionPersistence(LoadBalancerListenerPoolSessionPersistenceArgs sessionPersistence) {
            return sessionPersistence(Output.of(sessionPersistence));
        }

        public LoadBalancerListenerPoolArgs build() {
            return $;
        }
    }

}
