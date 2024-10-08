// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.outputs;

import com.ovhcloud.pulumi.ovh.Dbaas.outputs.LogsInputConfigurationFlowgger;
import com.ovhcloud.pulumi.ovh.Dbaas.outputs.LogsInputConfigurationLogstash;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LogsInputConfiguration {
    /**
     * @return Flowgger configuration
     * 
     */
    private @Nullable LogsInputConfigurationFlowgger flowgger;
    /**
     * @return Logstash configuration
     * 
     */
    private @Nullable LogsInputConfigurationLogstash logstash;

    private LogsInputConfiguration() {}
    /**
     * @return Flowgger configuration
     * 
     */
    public Optional<LogsInputConfigurationFlowgger> flowgger() {
        return Optional.ofNullable(this.flowgger);
    }
    /**
     * @return Logstash configuration
     * 
     */
    public Optional<LogsInputConfigurationLogstash> logstash() {
        return Optional.ofNullable(this.logstash);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LogsInputConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable LogsInputConfigurationFlowgger flowgger;
        private @Nullable LogsInputConfigurationLogstash logstash;
        public Builder() {}
        public Builder(LogsInputConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.flowgger = defaults.flowgger;
    	      this.logstash = defaults.logstash;
        }

        @CustomType.Setter
        public Builder flowgger(@Nullable LogsInputConfigurationFlowgger flowgger) {

            this.flowgger = flowgger;
            return this;
        }
        @CustomType.Setter
        public Builder logstash(@Nullable LogsInputConfigurationLogstash logstash) {

            this.logstash = logstash;
            return this;
        }
        public LogsInputConfiguration build() {
            final var _resultValue = new LogsInputConfiguration();
            _resultValue.flowgger = flowgger;
            _resultValue.logstash = logstash;
            return _resultValue;
        }
    }
}
