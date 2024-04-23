# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FirewallRuleArgs', 'FirewallRule']

@pulumi.input_type
class FirewallRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 ip: pulumi.Input[str],
                 ip_on_firewall: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 sequence: pulumi.Input[float],
                 destination_port: Optional[pulumi.Input[float]] = None,
                 fragments: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[float]] = None,
                 tcp_option: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallRule resource.
        :param pulumi.Input[str] action: Possible values for action (deny|permit)
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_firewall: IPv4 address
        :param pulumi.Input[str] protocol: Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        :param pulumi.Input[float] sequence: Rule position in the rules array
        :param pulumi.Input[float] destination_port: Destination port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[bool] fragments: Fragments option
        :param pulumi.Input[str] source: IPv4 CIDR notation (e.g., 192.0.2.0/24)
        :param pulumi.Input[float] source_port: Source port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] tcp_option: TCP option on your rule (syn|established)
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "ip_on_firewall", ip_on_firewall)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "sequence", sequence)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if fragments is not None:
            pulumi.set(__self__, "fragments", fragments)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)
        if tcp_option is not None:
            pulumi.set(__self__, "tcp_option", tcp_option)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Possible values for action (deny|permit)
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipOnFirewall")
    def ip_on_firewall(self) -> pulumi.Input[str]:
        """
        IPv4 address
        """
        return pulumi.get(self, "ip_on_firewall")

    @ip_on_firewall.setter
    def ip_on_firewall(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_on_firewall", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def sequence(self) -> pulumi.Input[float]:
        """
        Rule position in the rules array
        """
        return pulumi.get(self, "sequence")

    @sequence.setter
    def sequence(self, value: pulumi.Input[float]):
        pulumi.set(self, "sequence", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[float]]:
        """
        Destination port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter
    def fragments(self) -> Optional[pulumi.Input[bool]]:
        """
        Fragments option
        """
        return pulumi.get(self, "fragments")

    @fragments.setter
    def fragments(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fragments", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 CIDR notation (e.g., 192.0.2.0/24)
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[float]]:
        """
        Source port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "source_port", value)

    @property
    @pulumi.getter(name="tcpOption")
    def tcp_option(self) -> Optional[pulumi.Input[str]]:
        """
        TCP option on your rule (syn|established)
        """
        return pulumi.get(self, "tcp_option")

    @tcp_option.setter
    def tcp_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tcp_option", value)


@pulumi.input_type
class _FirewallRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[float]] = None,
                 destination_port_desc: Optional[pulumi.Input[str]] = None,
                 fragments: Optional[pulumi.Input[bool]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_firewall: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[str]] = None,
                 sequence: Optional[pulumi.Input[float]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[float]] = None,
                 source_port_desc: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tcp_option: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallRule resources.
        :param pulumi.Input[str] action: Possible values for action (deny|permit)
        :param pulumi.Input[str] creation_date: Creation date of the rule
        :param pulumi.Input[str] destination: Destination IP for your rule
        :param pulumi.Input[float] destination_port: Destination port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] destination_port_desc: String description of field `destination_port`
        :param pulumi.Input[bool] fragments: Fragments option
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_firewall: IPv4 address
        :param pulumi.Input[str] protocol: Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        :param pulumi.Input[str] rule: Description of the rule
        :param pulumi.Input[float] sequence: Rule position in the rules array
        :param pulumi.Input[str] source: IPv4 CIDR notation (e.g., 192.0.2.0/24)
        :param pulumi.Input[float] source_port: Source port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] source_port_desc: String description of field `source_port`
        :param pulumi.Input[str] state: Current state of your rule
        :param pulumi.Input[str] tcp_option: TCP option on your rule (syn|established)
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if destination_port_desc is not None:
            pulumi.set(__self__, "destination_port_desc", destination_port_desc)
        if fragments is not None:
            pulumi.set(__self__, "fragments", fragments)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_on_firewall is not None:
            pulumi.set(__self__, "ip_on_firewall", ip_on_firewall)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if rule is not None:
            pulumi.set(__self__, "rule", rule)
        if sequence is not None:
            pulumi.set(__self__, "sequence", sequence)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)
        if source_port_desc is not None:
            pulumi.set(__self__, "source_port_desc", source_port_desc)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tcp_option is not None:
            pulumi.set(__self__, "tcp_option", tcp_option)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Possible values for action (deny|permit)
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[pulumi.Input[str]]:
        """
        Creation date of the rule
        """
        return pulumi.get(self, "creation_date")

    @creation_date.setter
    def creation_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_date", value)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        Destination IP for your rule
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[float]]:
        """
        Destination port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter(name="destinationPortDesc")
    def destination_port_desc(self) -> Optional[pulumi.Input[str]]:
        """
        String description of field `destination_port`
        """
        return pulumi.get(self, "destination_port_desc")

    @destination_port_desc.setter
    def destination_port_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port_desc", value)

    @property
    @pulumi.getter
    def fragments(self) -> Optional[pulumi.Input[bool]]:
        """
        Fragments option
        """
        return pulumi.get(self, "fragments")

    @fragments.setter
    def fragments(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fragments", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipOnFirewall")
    def ip_on_firewall(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address
        """
        return pulumi.get(self, "ip_on_firewall")

    @ip_on_firewall.setter
    def ip_on_firewall(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_on_firewall", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def rule(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the rule
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule", value)

    @property
    @pulumi.getter
    def sequence(self) -> Optional[pulumi.Input[float]]:
        """
        Rule position in the rules array
        """
        return pulumi.get(self, "sequence")

    @sequence.setter
    def sequence(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "sequence", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 CIDR notation (e.g., 192.0.2.0/24)
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[float]]:
        """
        Source port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "source_port", value)

    @property
    @pulumi.getter(name="sourcePortDesc")
    def source_port_desc(self) -> Optional[pulumi.Input[str]]:
        """
        String description of field `source_port`
        """
        return pulumi.get(self, "source_port_desc")

    @source_port_desc.setter
    def source_port_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port_desc", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Current state of your rule
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="tcpOption")
    def tcp_option(self) -> Optional[pulumi.Input[str]]:
        """
        TCP option on your rule (syn|established)
        """
        return pulumi.get(self, "tcp_option")

    @tcp_option.setter
    def tcp_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tcp_option", value)


class FirewallRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[float]] = None,
                 fragments: Optional[pulumi.Input[bool]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_firewall: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 sequence: Optional[pulumi.Input[float]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[float]] = None,
                 tcp_option: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to manage a rule on an IP firewall.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        myfirewallrule = ovh.ip.FirewallRule("myfirewallrule",
            action="deny",
            ip="XXXXXX",
            ip_on_firewall="XXXXXX",
            protocol="tcp",
            sequence=0)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Possible values for action (deny|permit)
        :param pulumi.Input[float] destination_port: Destination port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[bool] fragments: Fragments option
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_firewall: IPv4 address
        :param pulumi.Input[str] protocol: Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        :param pulumi.Input[float] sequence: Rule position in the rules array
        :param pulumi.Input[str] source: IPv4 CIDR notation (e.g., 192.0.2.0/24)
        :param pulumi.Input[float] source_port: Source port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] tcp_option: TCP option on your rule (syn|established)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to manage a rule on an IP firewall.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ovh as ovh

        myfirewallrule = ovh.ip.FirewallRule("myfirewallrule",
            action="deny",
            ip="XXXXXX",
            ip_on_firewall="XXXXXX",
            protocol="tcp",
            sequence=0)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param FirewallRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[float]] = None,
                 fragments: Optional[pulumi.Input[bool]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_firewall: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 sequence: Optional[pulumi.Input[float]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[float]] = None,
                 tcp_option: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallRuleArgs.__new__(FirewallRuleArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["destination_port"] = destination_port
            __props__.__dict__["fragments"] = fragments
            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            if ip_on_firewall is None and not opts.urn:
                raise TypeError("Missing required property 'ip_on_firewall'")
            __props__.__dict__["ip_on_firewall"] = ip_on_firewall
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if sequence is None and not opts.urn:
                raise TypeError("Missing required property 'sequence'")
            __props__.__dict__["sequence"] = sequence
            __props__.__dict__["source"] = source
            __props__.__dict__["source_port"] = source_port
            __props__.__dict__["tcp_option"] = tcp_option
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["destination"] = None
            __props__.__dict__["destination_port_desc"] = None
            __props__.__dict__["rule"] = None
            __props__.__dict__["source_port_desc"] = None
            __props__.__dict__["state"] = None
        super(FirewallRule, __self__).__init__(
            'ovh:Ip/firewallRule:FirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[str]] = None,
            destination_port: Optional[pulumi.Input[float]] = None,
            destination_port_desc: Optional[pulumi.Input[str]] = None,
            fragments: Optional[pulumi.Input[bool]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip_on_firewall: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            rule: Optional[pulumi.Input[str]] = None,
            sequence: Optional[pulumi.Input[float]] = None,
            source: Optional[pulumi.Input[str]] = None,
            source_port: Optional[pulumi.Input[float]] = None,
            source_port_desc: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tcp_option: Optional[pulumi.Input[str]] = None) -> 'FirewallRule':
        """
        Get an existing FirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Possible values for action (deny|permit)
        :param pulumi.Input[str] creation_date: Creation date of the rule
        :param pulumi.Input[str] destination: Destination IP for your rule
        :param pulumi.Input[float] destination_port: Destination port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] destination_port_desc: String description of field `destination_port`
        :param pulumi.Input[bool] fragments: Fragments option
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_firewall: IPv4 address
        :param pulumi.Input[str] protocol: Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        :param pulumi.Input[str] rule: Description of the rule
        :param pulumi.Input[float] sequence: Rule position in the rules array
        :param pulumi.Input[str] source: IPv4 CIDR notation (e.g., 192.0.2.0/24)
        :param pulumi.Input[float] source_port: Source port for your rule. Only with TCP/UDP protocol
        :param pulumi.Input[str] source_port_desc: String description of field `source_port`
        :param pulumi.Input[str] state: Current state of your rule
        :param pulumi.Input[str] tcp_option: TCP option on your rule (syn|established)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallRuleState.__new__(_FirewallRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["creation_date"] = creation_date
        __props__.__dict__["destination"] = destination
        __props__.__dict__["destination_port"] = destination_port
        __props__.__dict__["destination_port_desc"] = destination_port_desc
        __props__.__dict__["fragments"] = fragments
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip_on_firewall"] = ip_on_firewall
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["rule"] = rule
        __props__.__dict__["sequence"] = sequence
        __props__.__dict__["source"] = source
        __props__.__dict__["source_port"] = source_port
        __props__.__dict__["source_port_desc"] = source_port_desc
        __props__.__dict__["state"] = state
        __props__.__dict__["tcp_option"] = tcp_option
        return FirewallRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Possible values for action (deny|permit)
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        Creation date of the rule
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        Destination IP for your rule
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> pulumi.Output[float]:
        """
        Destination port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "destination_port")

    @property
    @pulumi.getter(name="destinationPortDesc")
    def destination_port_desc(self) -> pulumi.Output[str]:
        """
        String description of field `destination_port`
        """
        return pulumi.get(self, "destination_port_desc")

    @property
    @pulumi.getter
    def fragments(self) -> pulumi.Output[bool]:
        """
        Fragments option
        """
        return pulumi.get(self, "fragments")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipOnFirewall")
    def ip_on_firewall(self) -> pulumi.Output[str]:
        """
        IPv4 address
        """
        return pulumi.get(self, "ip_on_firewall")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Possible values for protocol (ah|esp|gre|icmp|ipv4|tcp|udp)
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Output[str]:
        """
        Description of the rule
        """
        return pulumi.get(self, "rule")

    @property
    @pulumi.getter
    def sequence(self) -> pulumi.Output[float]:
        """
        Rule position in the rules array
        """
        return pulumi.get(self, "sequence")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        IPv4 CIDR notation (e.g., 192.0.2.0/24)
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> pulumi.Output[float]:
        """
        Source port for your rule. Only with TCP/UDP protocol
        """
        return pulumi.get(self, "source_port")

    @property
    @pulumi.getter(name="sourcePortDesc")
    def source_port_desc(self) -> pulumi.Output[str]:
        """
        String description of field `source_port`
        """
        return pulumi.get(self, "source_port_desc")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of your rule
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="tcpOption")
    def tcp_option(self) -> pulumi.Output[str]:
        """
        TCP option on your rule (syn|established)
        """
        return pulumi.get(self, "tcp_option")
