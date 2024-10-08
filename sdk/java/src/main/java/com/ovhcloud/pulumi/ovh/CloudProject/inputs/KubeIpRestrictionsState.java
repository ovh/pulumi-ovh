// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeIpRestrictionsState extends com.pulumi.resources.ResourceArgs {

    public static final KubeIpRestrictionsState Empty = new KubeIpRestrictionsState();

    /**
     * List of CIDR authorized to interact with the managed Kubernetes cluster.
     * 
     */
    @Import(name="ips")
    private @Nullable Output<List<String>> ips;

    /**
     * @return List of CIDR authorized to interact with the managed Kubernetes cluster.
     * 
     */
    public Optional<Output<List<String>>> ips() {
        return Optional.ofNullable(this.ips);
    }

    /**
     * The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="kubeId")
    private @Nullable Output<String> kubeId;

    /**
     * @return The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<String>> kubeId() {
        return Optional.ofNullable(this.kubeId);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private KubeIpRestrictionsState() {}

    private KubeIpRestrictionsState(KubeIpRestrictionsState $) {
        this.ips = $.ips;
        this.kubeId = $.kubeId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeIpRestrictionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeIpRestrictionsState $;

        public Builder() {
            $ = new KubeIpRestrictionsState();
        }

        public Builder(KubeIpRestrictionsState defaults) {
            $ = new KubeIpRestrictionsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ips List of CIDR authorized to interact with the managed Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder ips(@Nullable Output<List<String>> ips) {
            $.ips = ips;
            return this;
        }

        /**
         * @param ips List of CIDR authorized to interact with the managed Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder ips(List<String> ips) {
            return ips(Output.of(ips));
        }

        /**
         * @param ips List of CIDR authorized to interact with the managed Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder ips(String... ips) {
            return ips(List.of(ips));
        }

        /**
         * @param kubeId The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder kubeId(@Nullable Output<String> kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param kubeId The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            return kubeId(Output.of(kubeId));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public KubeIpRestrictionsState build() {
            return $;
        }
    }

}
