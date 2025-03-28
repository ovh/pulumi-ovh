// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.outputs;

import com.ovhcloud.pulumi.ovh.outputs.GetServerVni;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetServerResult {
    /**
     * @return URN of the dedicated server instance
     * 
     */
    private String ServerURN;
    /**
     * @return Dedicated AZ localisation
     * 
     */
    private String availabilityZone;
    /**
     * @return Boot id of the server
     * 
     */
    private Integer bootId;
    /**
     * @return Boot script of the server
     * 
     */
    private String bootScript;
    /**
     * @return Dedicated server commercial range
     * 
     */
    private String commercialRange;
    /**
     * @return Dedicated datacenter localisation (bhs1,bhs2,...)
     * 
     */
    private String datacenter;
    /**
     * @return Dedicated server display name
     * 
     */
    private String displayName;
    /**
     * @return Path of the EFI bootloader of the dedicated server
     * 
     */
    private String efiBootloaderPath;
    /**
     * @return List of enabled public VNI uuids
     * 
     */
    private List<String> enabledPublicVnis;
    /**
     * @return List of enabled vrack_aggregation VNI uuids
     * 
     */
    private List<String> enabledVrackAggregationVnis;
    /**
     * @return List of enabled vrack VNI uuids
     * 
     */
    private List<String> enabledVrackVnis;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Dedicated server ip (IPv4)
     * 
     */
    private String ip;
    /**
     * @return Dedicated server ip blocks
     * 
     */
    private List<String> ips;
    /**
     * @return Link speed of the server
     * 
     */
    private Integer linkSpeed;
    /**
     * @return Icmp monitoring state
     * 
     */
    private Boolean monitoring;
    /**
     * @return User defined VirtualNetworkInterface name
     * 
     */
    private String name;
    private Boolean newUpgradeSystem;
    /**
     * @return Prevent datacenter intervention
     * 
     */
    private Boolean noIntervention;
    /**
     * @return Operating system
     * 
     */
    private String os;
    /**
     * @return Power state of the server (poweroff, poweron)
     * 
     */
    private String powerState;
    /**
     * @return Does this server have professional use option
     * 
     */
    private Boolean professionalUse;
    /**
     * @return Rack id of the server
     * 
     */
    private String rack;
    /**
     * @return Dedicated region localisation
     * 
     */
    private String region;
    /**
     * @return Rescue mail of the server
     * 
     */
    private String rescueMail;
    /**
     * @return Public SSH Key used in the rescue mode
     * 
     */
    private String rescueSshKey;
    /**
     * @return Dedicated server reverse
     * 
     */
    private String reverse;
    /**
     * @return Root device of the server
     * 
     */
    private String rootDevice;
    /**
     * @return Server id
     * 
     */
    private Integer serverId;
    private String serviceName;
    /**
     * @return Error, hacked, hackedBlocked, ok
     * 
     */
    private String state;
    /**
     * @return Dedicated server support level (critical, fastpath, gs, pro)
     * 
     */
    private String supportLevel;
    /**
     * @return The list of Virtualnetworkinterface associated with this server
     * 
     */
    private List<GetServerVni> vnis;

    private GetServerResult() {}
    /**
     * @return URN of the dedicated server instance
     * 
     */
    public String ServerURN() {
        return this.ServerURN;
    }
    /**
     * @return Dedicated AZ localisation
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * @return Boot id of the server
     * 
     */
    public Integer bootId() {
        return this.bootId;
    }
    /**
     * @return Boot script of the server
     * 
     */
    public String bootScript() {
        return this.bootScript;
    }
    /**
     * @return Dedicated server commercial range
     * 
     */
    public String commercialRange() {
        return this.commercialRange;
    }
    /**
     * @return Dedicated datacenter localisation (bhs1,bhs2,...)
     * 
     */
    public String datacenter() {
        return this.datacenter;
    }
    /**
     * @return Dedicated server display name
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return Path of the EFI bootloader of the dedicated server
     * 
     */
    public String efiBootloaderPath() {
        return this.efiBootloaderPath;
    }
    /**
     * @return List of enabled public VNI uuids
     * 
     */
    public List<String> enabledPublicVnis() {
        return this.enabledPublicVnis;
    }
    /**
     * @return List of enabled vrack_aggregation VNI uuids
     * 
     */
    public List<String> enabledVrackAggregationVnis() {
        return this.enabledVrackAggregationVnis;
    }
    /**
     * @return List of enabled vrack VNI uuids
     * 
     */
    public List<String> enabledVrackVnis() {
        return this.enabledVrackVnis;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Dedicated server ip (IPv4)
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return Dedicated server ip blocks
     * 
     */
    public List<String> ips() {
        return this.ips;
    }
    /**
     * @return Link speed of the server
     * 
     */
    public Integer linkSpeed() {
        return this.linkSpeed;
    }
    /**
     * @return Icmp monitoring state
     * 
     */
    public Boolean monitoring() {
        return this.monitoring;
    }
    /**
     * @return User defined VirtualNetworkInterface name
     * 
     */
    public String name() {
        return this.name;
    }
    public Boolean newUpgradeSystem() {
        return this.newUpgradeSystem;
    }
    /**
     * @return Prevent datacenter intervention
     * 
     */
    public Boolean noIntervention() {
        return this.noIntervention;
    }
    /**
     * @return Operating system
     * 
     */
    public String os() {
        return this.os;
    }
    /**
     * @return Power state of the server (poweroff, poweron)
     * 
     */
    public String powerState() {
        return this.powerState;
    }
    /**
     * @return Does this server have professional use option
     * 
     */
    public Boolean professionalUse() {
        return this.professionalUse;
    }
    /**
     * @return Rack id of the server
     * 
     */
    public String rack() {
        return this.rack;
    }
    /**
     * @return Dedicated region localisation
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Rescue mail of the server
     * 
     */
    public String rescueMail() {
        return this.rescueMail;
    }
    /**
     * @return Public SSH Key used in the rescue mode
     * 
     */
    public String rescueSshKey() {
        return this.rescueSshKey;
    }
    /**
     * @return Dedicated server reverse
     * 
     */
    public String reverse() {
        return this.reverse;
    }
    /**
     * @return Root device of the server
     * 
     */
    public String rootDevice() {
        return this.rootDevice;
    }
    /**
     * @return Server id
     * 
     */
    public Integer serverId() {
        return this.serverId;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Error, hacked, hackedBlocked, ok
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return Dedicated server support level (critical, fastpath, gs, pro)
     * 
     */
    public String supportLevel() {
        return this.supportLevel;
    }
    /**
     * @return The list of Virtualnetworkinterface associated with this server
     * 
     */
    public List<GetServerVni> vnis() {
        return this.vnis;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ServerURN;
        private String availabilityZone;
        private Integer bootId;
        private String bootScript;
        private String commercialRange;
        private String datacenter;
        private String displayName;
        private String efiBootloaderPath;
        private List<String> enabledPublicVnis;
        private List<String> enabledVrackAggregationVnis;
        private List<String> enabledVrackVnis;
        private String id;
        private String ip;
        private List<String> ips;
        private Integer linkSpeed;
        private Boolean monitoring;
        private String name;
        private Boolean newUpgradeSystem;
        private Boolean noIntervention;
        private String os;
        private String powerState;
        private Boolean professionalUse;
        private String rack;
        private String region;
        private String rescueMail;
        private String rescueSshKey;
        private String reverse;
        private String rootDevice;
        private Integer serverId;
        private String serviceName;
        private String state;
        private String supportLevel;
        private List<GetServerVni> vnis;
        public Builder() {}
        public Builder(GetServerResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ServerURN = defaults.ServerURN;
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.bootId = defaults.bootId;
    	      this.bootScript = defaults.bootScript;
    	      this.commercialRange = defaults.commercialRange;
    	      this.datacenter = defaults.datacenter;
    	      this.displayName = defaults.displayName;
    	      this.efiBootloaderPath = defaults.efiBootloaderPath;
    	      this.enabledPublicVnis = defaults.enabledPublicVnis;
    	      this.enabledVrackAggregationVnis = defaults.enabledVrackAggregationVnis;
    	      this.enabledVrackVnis = defaults.enabledVrackVnis;
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
    	      this.ips = defaults.ips;
    	      this.linkSpeed = defaults.linkSpeed;
    	      this.monitoring = defaults.monitoring;
    	      this.name = defaults.name;
    	      this.newUpgradeSystem = defaults.newUpgradeSystem;
    	      this.noIntervention = defaults.noIntervention;
    	      this.os = defaults.os;
    	      this.powerState = defaults.powerState;
    	      this.professionalUse = defaults.professionalUse;
    	      this.rack = defaults.rack;
    	      this.region = defaults.region;
    	      this.rescueMail = defaults.rescueMail;
    	      this.rescueSshKey = defaults.rescueSshKey;
    	      this.reverse = defaults.reverse;
    	      this.rootDevice = defaults.rootDevice;
    	      this.serverId = defaults.serverId;
    	      this.serviceName = defaults.serviceName;
    	      this.state = defaults.state;
    	      this.supportLevel = defaults.supportLevel;
    	      this.vnis = defaults.vnis;
        }

        @CustomType.Setter
        public Builder ServerURN(String ServerURN) {
            if (ServerURN == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "ServerURN");
            }
            this.ServerURN = ServerURN;
            return this;
        }
        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            if (availabilityZone == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "availabilityZone");
            }
            this.availabilityZone = availabilityZone;
            return this;
        }
        @CustomType.Setter
        public Builder bootId(Integer bootId) {
            if (bootId == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "bootId");
            }
            this.bootId = bootId;
            return this;
        }
        @CustomType.Setter
        public Builder bootScript(String bootScript) {
            if (bootScript == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "bootScript");
            }
            this.bootScript = bootScript;
            return this;
        }
        @CustomType.Setter
        public Builder commercialRange(String commercialRange) {
            if (commercialRange == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "commercialRange");
            }
            this.commercialRange = commercialRange;
            return this;
        }
        @CustomType.Setter
        public Builder datacenter(String datacenter) {
            if (datacenter == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "datacenter");
            }
            this.datacenter = datacenter;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder efiBootloaderPath(String efiBootloaderPath) {
            if (efiBootloaderPath == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "efiBootloaderPath");
            }
            this.efiBootloaderPath = efiBootloaderPath;
            return this;
        }
        @CustomType.Setter
        public Builder enabledPublicVnis(List<String> enabledPublicVnis) {
            if (enabledPublicVnis == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "enabledPublicVnis");
            }
            this.enabledPublicVnis = enabledPublicVnis;
            return this;
        }
        public Builder enabledPublicVnis(String... enabledPublicVnis) {
            return enabledPublicVnis(List.of(enabledPublicVnis));
        }
        @CustomType.Setter
        public Builder enabledVrackAggregationVnis(List<String> enabledVrackAggregationVnis) {
            if (enabledVrackAggregationVnis == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "enabledVrackAggregationVnis");
            }
            this.enabledVrackAggregationVnis = enabledVrackAggregationVnis;
            return this;
        }
        public Builder enabledVrackAggregationVnis(String... enabledVrackAggregationVnis) {
            return enabledVrackAggregationVnis(List.of(enabledVrackAggregationVnis));
        }
        @CustomType.Setter
        public Builder enabledVrackVnis(List<String> enabledVrackVnis) {
            if (enabledVrackVnis == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "enabledVrackVnis");
            }
            this.enabledVrackVnis = enabledVrackVnis;
            return this;
        }
        public Builder enabledVrackVnis(String... enabledVrackVnis) {
            return enabledVrackVnis(List.of(enabledVrackVnis));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            if (ip == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "ip");
            }
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder ips(List<String> ips) {
            if (ips == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "ips");
            }
            this.ips = ips;
            return this;
        }
        public Builder ips(String... ips) {
            return ips(List.of(ips));
        }
        @CustomType.Setter
        public Builder linkSpeed(Integer linkSpeed) {
            if (linkSpeed == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "linkSpeed");
            }
            this.linkSpeed = linkSpeed;
            return this;
        }
        @CustomType.Setter
        public Builder monitoring(Boolean monitoring) {
            if (monitoring == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "monitoring");
            }
            this.monitoring = monitoring;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder newUpgradeSystem(Boolean newUpgradeSystem) {
            if (newUpgradeSystem == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "newUpgradeSystem");
            }
            this.newUpgradeSystem = newUpgradeSystem;
            return this;
        }
        @CustomType.Setter
        public Builder noIntervention(Boolean noIntervention) {
            if (noIntervention == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "noIntervention");
            }
            this.noIntervention = noIntervention;
            return this;
        }
        @CustomType.Setter
        public Builder os(String os) {
            if (os == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "os");
            }
            this.os = os;
            return this;
        }
        @CustomType.Setter
        public Builder powerState(String powerState) {
            if (powerState == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "powerState");
            }
            this.powerState = powerState;
            return this;
        }
        @CustomType.Setter
        public Builder professionalUse(Boolean professionalUse) {
            if (professionalUse == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "professionalUse");
            }
            this.professionalUse = professionalUse;
            return this;
        }
        @CustomType.Setter
        public Builder rack(String rack) {
            if (rack == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "rack");
            }
            this.rack = rack;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder rescueMail(String rescueMail) {
            if (rescueMail == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "rescueMail");
            }
            this.rescueMail = rescueMail;
            return this;
        }
        @CustomType.Setter
        public Builder rescueSshKey(String rescueSshKey) {
            if (rescueSshKey == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "rescueSshKey");
            }
            this.rescueSshKey = rescueSshKey;
            return this;
        }
        @CustomType.Setter
        public Builder reverse(String reverse) {
            if (reverse == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "reverse");
            }
            this.reverse = reverse;
            return this;
        }
        @CustomType.Setter
        public Builder rootDevice(String rootDevice) {
            if (rootDevice == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "rootDevice");
            }
            this.rootDevice = rootDevice;
            return this;
        }
        @CustomType.Setter
        public Builder serverId(Integer serverId) {
            if (serverId == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "serverId");
            }
            this.serverId = serverId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder supportLevel(String supportLevel) {
            if (supportLevel == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "supportLevel");
            }
            this.supportLevel = supportLevel;
            return this;
        }
        @CustomType.Setter
        public Builder vnis(List<GetServerVni> vnis) {
            if (vnis == null) {
              throw new MissingRequiredPropertyException("GetServerResult", "vnis");
            }
            this.vnis = vnis;
            return this;
        }
        public Builder vnis(GetServerVni... vnis) {
            return vnis(List.of(vnis));
        }
        public GetServerResult build() {
            final var _resultValue = new GetServerResult();
            _resultValue.ServerURN = ServerURN;
            _resultValue.availabilityZone = availabilityZone;
            _resultValue.bootId = bootId;
            _resultValue.bootScript = bootScript;
            _resultValue.commercialRange = commercialRange;
            _resultValue.datacenter = datacenter;
            _resultValue.displayName = displayName;
            _resultValue.efiBootloaderPath = efiBootloaderPath;
            _resultValue.enabledPublicVnis = enabledPublicVnis;
            _resultValue.enabledVrackAggregationVnis = enabledVrackAggregationVnis;
            _resultValue.enabledVrackVnis = enabledVrackVnis;
            _resultValue.id = id;
            _resultValue.ip = ip;
            _resultValue.ips = ips;
            _resultValue.linkSpeed = linkSpeed;
            _resultValue.monitoring = monitoring;
            _resultValue.name = name;
            _resultValue.newUpgradeSystem = newUpgradeSystem;
            _resultValue.noIntervention = noIntervention;
            _resultValue.os = os;
            _resultValue.powerState = powerState;
            _resultValue.professionalUse = professionalUse;
            _resultValue.rack = rack;
            _resultValue.region = region;
            _resultValue.rescueMail = rescueMail;
            _resultValue.rescueSshKey = rescueSshKey;
            _resultValue.reverse = reverse;
            _resultValue.rootDevice = rootDevice;
            _resultValue.serverId = serverId;
            _resultValue.serviceName = serviceName;
            _resultValue.state = state;
            _resultValue.supportLevel = supportLevel;
            _resultValue.vnis = vnis;
            return _resultValue;
        }
    }
}
