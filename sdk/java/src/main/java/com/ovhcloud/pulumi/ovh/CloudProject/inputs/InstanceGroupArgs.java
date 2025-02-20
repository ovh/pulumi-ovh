// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceGroupArgs Empty = new InstanceGroupArgs();

    /**
     * Group id
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return Group id
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    private InstanceGroupArgs() {}

    private InstanceGroupArgs(InstanceGroupArgs $) {
        this.groupId = $.groupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceGroupArgs $;

        public Builder() {
            $ = new InstanceGroupArgs();
        }

        public Builder(InstanceGroupArgs defaults) {
            $ = new InstanceGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId Group id
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId Group id
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        public InstanceGroupArgs build() {
            return $;
        }
    }

}
