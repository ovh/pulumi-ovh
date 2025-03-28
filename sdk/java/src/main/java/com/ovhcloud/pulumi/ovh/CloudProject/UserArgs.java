// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * A description associated with the user.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description associated with the user.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.
     * 
     */
    @Import(name="openstackRc")
    private @Nullable Output<Map<String,String>> openstackRc;

    /**
     * @return a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.
     * 
     */
    public Optional<Output<Map<String,String>>> openstackRc() {
        return Optional.ofNullable(this.openstackRc);
    }

    /**
     * The name of a role. See `role_names`.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return The name of a role. See `role_names`.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * A list of role names. Values can be:
     * - administrator,
     * - ai_training_operator
     * - ai_training_read
     * - authentication
     * - backup_operator
     * - compute_operator
     * - image_operator
     * - infrastructure_supervisor
     * - network_operator
     * - network_security_operator
     * - objectstore_operator
     * - volume_operator
     * 
     */
    @Import(name="roleNames")
    private @Nullable Output<List<String>> roleNames;

    /**
     * @return A list of role names. Values can be:
     * - administrator,
     * - ai_training_operator
     * - ai_training_read
     * - authentication
     * - backup_operator
     * - compute_operator
     * - image_operator
     * - infrastructure_supervisor
     * - network_operator
     * - network_security_operator
     * - objectstore_operator
     * - volume_operator
     * 
     */
    public Optional<Output<List<String>>> roleNames() {
        return Optional.ofNullable(this.roleNames);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.description = $.description;
        this.openstackRc = $.openstackRc;
        this.roleName = $.roleName;
        this.roleNames = $.roleNames;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description associated with the user.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description associated with the user.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param openstackRc a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.
         * 
         * @return builder
         * 
         */
        public Builder openstackRc(@Nullable Output<Map<String,String>> openstackRc) {
            $.openstackRc = openstackRc;
            return this;
        }

        /**
         * @param openstackRc a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.
         * 
         * @return builder
         * 
         */
        public Builder openstackRc(Map<String,String> openstackRc) {
            return openstackRc(Output.of(openstackRc));
        }

        /**
         * @param roleName The name of a role. See `role_names`.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of a role. See `role_names`.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param roleNames A list of role names. Values can be:
         * - administrator,
         * - ai_training_operator
         * - ai_training_read
         * - authentication
         * - backup_operator
         * - compute_operator
         * - image_operator
         * - infrastructure_supervisor
         * - network_operator
         * - network_security_operator
         * - objectstore_operator
         * - volume_operator
         * 
         * @return builder
         * 
         */
        public Builder roleNames(@Nullable Output<List<String>> roleNames) {
            $.roleNames = roleNames;
            return this;
        }

        /**
         * @param roleNames A list of role names. Values can be:
         * - administrator,
         * - ai_training_operator
         * - ai_training_read
         * - authentication
         * - backup_operator
         * - compute_operator
         * - image_operator
         * - infrastructure_supervisor
         * - network_operator
         * - network_security_operator
         * - objectstore_operator
         * - volume_operator
         * 
         * @return builder
         * 
         */
        public Builder roleNames(List<String> roleNames) {
            return roleNames(Output.of(roleNames));
        }

        /**
         * @param roleNames A list of role names. Values can be:
         * - administrator,
         * - ai_training_operator
         * - ai_training_read
         * - authentication
         * - backup_operator
         * - compute_operator
         * - image_operator
         * - infrastructure_supervisor
         * - network_operator
         * - network_security_operator
         * - objectstore_operator
         * - volume_operator
         * 
         * @return builder
         * 
         */
        public Builder roleNames(String... roleNames) {
            return roleNames(List.of(roleNames));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public UserArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("UserArgs", "serviceName");
            }
            return $;
        }
    }

}
