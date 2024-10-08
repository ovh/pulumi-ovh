// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Iam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceGroupState extends com.pulumi.resources.ResourceArgs {

    public static final ResourceGroupState Empty = new ResourceGroupState();

    /**
     * URN of the resource group, used when writing policies
     * 
     */
    @Import(name="GroupURN")
    private @Nullable Output<String> GroupURN;

    /**
     * @return URN of the resource group, used when writing policies
     * 
     */
    public Optional<Output<String>> GroupURN() {
        return Optional.ofNullable(this.GroupURN);
    }

    /**
     * Date of the creation of the resource group
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Date of the creation of the resource group
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Name of the resource group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name of the account owning the resource group
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Name of the account owning the resource group
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
     * 
     */
    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    /**
     * @return Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
     * 
     */
    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    /**
     * Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     * 
     */
    @Import(name="resources")
    private @Nullable Output<List<String>> resources;

    /**
     * @return Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
     * 
     */
    public Optional<Output<List<String>>> resources() {
        return Optional.ofNullable(this.resources);
    }

    /**
     * Date of the last modification of the resource group
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Date of the last modification of the resource group
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private ResourceGroupState() {}

    private ResourceGroupState(ResourceGroupState $) {
        this.GroupURN = $.GroupURN;
        this.createdAt = $.createdAt;
        this.name = $.name;
        this.owner = $.owner;
        this.readOnly = $.readOnly;
        this.resources = $.resources;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceGroupState $;

        public Builder() {
            $ = new ResourceGroupState();
        }

        public Builder(ResourceGroupState defaults) {
            $ = new ResourceGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param GroupURN URN of the resource group, used when writing policies
         * 
         * @return builder
         * 
         */
        public Builder GroupURN(@Nullable Output<String> GroupURN) {
            $.GroupURN = GroupURN;
            return this;
        }

        /**
         * @param GroupURN URN of the resource group, used when writing policies
         * 
         * @return builder
         * 
         */
        public Builder GroupURN(String GroupURN) {
            return GroupURN(Output.of(GroupURN));
        }

        /**
         * @param createdAt Date of the creation of the resource group
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Date of the creation of the resource group
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param name Name of the resource group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param owner Name of the account owning the resource group
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Name of the account owning the resource group
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param readOnly Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
         * 
         * @return builder
         * 
         */
        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        /**
         * @param readOnly Marks that the resource group is not editable. Usually means that this is a default resource group created by OVHcloud
         * 
         * @return builder
         * 
         */
        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(@Nullable Output<List<String>> resources) {
            $.resources = resources;
            return this;
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(List<String> resources) {
            return resources(Output.of(resources));
        }

        /**
         * @param resources Set of the URNs of the resources contained in the resource group. All urns must be ones of valid resources
         * 
         * @return builder
         * 
         */
        public Builder resources(String... resources) {
            return resources(List.of(resources));
        }

        /**
         * @param updatedAt Date of the last modification of the resource group
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Date of the last modification of the resource group
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public ResourceGroupState build() {
            return $;
        }
    }

}
