# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CrawlerArgs', 'Crawler']

@pulumi.input_type
class CrawlerArgs:
    def __init__(__self__, *,
                 role: pulumi.Input[str],
                 targets: pulumi.Input['CrawlerTargetsArgs'],
                 classifiers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 crawler_security_configuration: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recrawl_policy: Optional[pulumi.Input['CrawlerRecrawlPolicyArgs']] = None,
                 schedule: Optional[pulumi.Input['CrawlerScheduleArgs']] = None,
                 schema_change_policy: Optional[pulumi.Input['CrawlerSchemaChangePolicyArgs']] = None,
                 table_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Crawler resource.
        """
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "targets", targets)
        if classifiers is not None:
            pulumi.set(__self__, "classifiers", classifiers)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if crawler_security_configuration is not None:
            pulumi.set(__self__, "crawler_security_configuration", crawler_security_configuration)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if recrawl_policy is not None:
            pulumi.set(__self__, "recrawl_policy", recrawl_policy)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if schema_change_policy is not None:
            pulumi.set(__self__, "schema_change_policy", schema_change_policy)
        if table_prefix is not None:
            pulumi.set(__self__, "table_prefix", table_prefix)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Input['CrawlerTargetsArgs']:
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: pulumi.Input['CrawlerTargetsArgs']):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter
    def classifiers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "classifiers")

    @classifiers.setter
    def classifiers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "classifiers", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter(name="crawlerSecurityConfiguration")
    def crawler_security_configuration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "crawler_security_configuration")

    @crawler_security_configuration.setter
    def crawler_security_configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crawler_security_configuration", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recrawlPolicy")
    def recrawl_policy(self) -> Optional[pulumi.Input['CrawlerRecrawlPolicyArgs']]:
        return pulumi.get(self, "recrawl_policy")

    @recrawl_policy.setter
    def recrawl_policy(self, value: Optional[pulumi.Input['CrawlerRecrawlPolicyArgs']]):
        pulumi.set(self, "recrawl_policy", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['CrawlerScheduleArgs']]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['CrawlerScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="schemaChangePolicy")
    def schema_change_policy(self) -> Optional[pulumi.Input['CrawlerSchemaChangePolicyArgs']]:
        return pulumi.get(self, "schema_change_policy")

    @schema_change_policy.setter
    def schema_change_policy(self, value: Optional[pulumi.Input['CrawlerSchemaChangePolicyArgs']]):
        pulumi.set(self, "schema_change_policy", value)

    @property
    @pulumi.getter(name="tablePrefix")
    def table_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "table_prefix")

    @table_prefix.setter
    def table_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_prefix", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


warnings.warn("""Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Crawler(pulumi.CustomResource):
    warnings.warn("""Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 classifiers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 crawler_security_configuration: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recrawl_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerRecrawlPolicyArgs']]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['CrawlerScheduleArgs']]] = None,
                 schema_change_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerSchemaChangePolicyArgs']]] = None,
                 table_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 targets: Optional[pulumi.Input[pulumi.InputType['CrawlerTargetsArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Glue::Crawler

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CrawlerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Glue::Crawler

        :param str resource_name: The name of the resource.
        :param CrawlerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CrawlerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 classifiers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 crawler_security_configuration: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recrawl_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerRecrawlPolicyArgs']]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['CrawlerScheduleArgs']]] = None,
                 schema_change_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerSchemaChangePolicyArgs']]] = None,
                 table_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 targets: Optional[pulumi.Input[pulumi.InputType['CrawlerTargetsArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""Crawler is deprecated: Crawler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CrawlerArgs.__new__(CrawlerArgs)

            __props__.__dict__["classifiers"] = classifiers
            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["crawler_security_configuration"] = crawler_security_configuration
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["recrawl_policy"] = recrawl_policy
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["schema_change_policy"] = schema_change_policy
            __props__.__dict__["table_prefix"] = table_prefix
            __props__.__dict__["tags"] = tags
            if targets is None and not opts.urn:
                raise TypeError("Missing required property 'targets'")
            __props__.__dict__["targets"] = targets
        super(Crawler, __self__).__init__(
            'aws-native:glue:Crawler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Crawler':
        """
        Get an existing Crawler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CrawlerArgs.__new__(CrawlerArgs)

        __props__.__dict__["classifiers"] = None
        __props__.__dict__["configuration"] = None
        __props__.__dict__["crawler_security_configuration"] = None
        __props__.__dict__["database_name"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["recrawl_policy"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["schedule"] = None
        __props__.__dict__["schema_change_policy"] = None
        __props__.__dict__["table_prefix"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["targets"] = None
        return Crawler(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def classifiers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "classifiers")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="crawlerSecurityConfiguration")
    def crawler_security_configuration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "crawler_security_configuration")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recrawlPolicy")
    def recrawl_policy(self) -> pulumi.Output[Optional['outputs.CrawlerRecrawlPolicy']]:
        return pulumi.get(self, "recrawl_policy")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional['outputs.CrawlerSchedule']]:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="schemaChangePolicy")
    def schema_change_policy(self) -> pulumi.Output[Optional['outputs.CrawlerSchemaChangePolicy']]:
        return pulumi.get(self, "schema_change_policy")

    @property
    @pulumi.getter(name="tablePrefix")
    def table_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "table_prefix")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output['outputs.CrawlerTargets']:
        return pulumi.get(self, "targets")

