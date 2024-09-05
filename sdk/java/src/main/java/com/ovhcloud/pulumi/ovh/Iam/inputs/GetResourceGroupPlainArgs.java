// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Iam.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetResourceGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResourceGroupPlainArgs Empty = new GetResourceGroupPlainArgs();

    /**
     * Id of the resource group
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return Id of the resource group
     * 
     */
    public String id() {
        return this.id;
    }

    private GetResourceGroupPlainArgs() {}

    private GetResourceGroupPlainArgs(GetResourceGroupPlainArgs $) {
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourceGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourceGroupPlainArgs $;

        public Builder() {
            $ = new GetResourceGroupPlainArgs();
        }

        public Builder(GetResourceGroupPlainArgs defaults) {
            $ = new GetResourceGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Id of the resource group
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public GetResourceGroupPlainArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetResourceGroupPlainArgs", "id");
            }
            return $;
        }
    }

}