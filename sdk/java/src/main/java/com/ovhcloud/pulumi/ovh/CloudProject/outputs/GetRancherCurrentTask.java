// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRancherCurrentTask {
    /**
     * @return Identifier of the current task
     * 
     */
    private String id;
    /**
     * @return Link to the task details
     * 
     */
    private String link;
    /**
     * @return Current global status of the current task
     * 
     */
    private String status;
    /**
     * @return Type of the current task
     * 
     */
    private String type;

    private GetRancherCurrentTask() {}
    /**
     * @return Identifier of the current task
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Link to the task details
     * 
     */
    public String link() {
        return this.link;
    }
    /**
     * @return Current global status of the current task
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Type of the current task
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRancherCurrentTask defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String link;
        private String status;
        private String type;
        public Builder() {}
        public Builder(GetRancherCurrentTask defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.link = defaults.link;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentTask", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder link(String link) {
            if (link == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentTask", "link");
            }
            this.link = link;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentTask", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentTask", "type");
            }
            this.type = type;
            return this;
        }
        public GetRancherCurrentTask build() {
            final var _resultValue = new GetRancherCurrentTask();
            _resultValue.id = id;
            _resultValue.link = link;
            _resultValue.status = status;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
