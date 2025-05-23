// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherTargetSpecIpRestrictionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RancherTargetSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final RancherTargetSpecArgs Empty = new RancherTargetSpecArgs();

    /**
     * List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
     * 
     */
    @Import(name="ipRestrictions")
    private @Nullable Output<List<RancherTargetSpecIpRestrictionArgs>> ipRestrictions;

    /**
     * @return List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
     * 
     */
    public Optional<Output<List<RancherTargetSpecIpRestrictionArgs>>> ipRestrictions() {
        return Optional.ofNullable(this.ipRestrictions);
    }

    /**
     * Name of the managed Rancher service
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the managed Rancher service
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan
     * 
     */
    @Import(name="plan", required=true)
    private Output<String> plan;

    /**
     * @return Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }

    /**
     * Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using ovh*cloud*project*rancher*version datasource. Default is the latest version.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using ovh*cloud*project*rancher*version datasource. Default is the latest version.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private RancherTargetSpecArgs() {}

    private RancherTargetSpecArgs(RancherTargetSpecArgs $) {
        this.ipRestrictions = $.ipRestrictions;
        this.name = $.name;
        this.plan = $.plan;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RancherTargetSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RancherTargetSpecArgs $;

        public Builder() {
            $ = new RancherTargetSpecArgs();
        }

        public Builder(RancherTargetSpecArgs defaults) {
            $ = new RancherTargetSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipRestrictions List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(@Nullable Output<List<RancherTargetSpecIpRestrictionArgs>> ipRestrictions) {
            $.ipRestrictions = ipRestrictions;
            return this;
        }

        /**
         * @param ipRestrictions List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(List<RancherTargetSpecIpRestrictionArgs> ipRestrictions) {
            return ipRestrictions(Output.of(ipRestrictions));
        }

        /**
         * @param ipRestrictions List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(RancherTargetSpecIpRestrictionArgs... ipRestrictions) {
            return ipRestrictions(List.of(ipRestrictions));
        }

        /**
         * @param name Name of the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param plan Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<String> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param version Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using ovh*cloud*project*rancher*version datasource. Default is the latest version.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using ovh*cloud*project*rancher*version datasource. Default is the latest version.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public RancherTargetSpecArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("RancherTargetSpecArgs", "name");
            }
            if ($.plan == null) {
                throw new MissingRequiredPropertyException("RancherTargetSpecArgs", "plan");
            }
            return $;
        }
    }

}
