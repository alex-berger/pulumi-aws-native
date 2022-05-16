# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamedQueryArgs', 'NamedQuery']

@pulumi.input_type
class NamedQueryArgs:
    def __init__(__self__, *,
                 database: pulumi.Input[str],
                 query_string: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 work_group: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NamedQuery resource.
        :param pulumi.Input[str] database: The database to which the query belongs.
        :param pulumi.Input[str] query_string: The contents of the query with all query statements.
        :param pulumi.Input[str] description: The query description.
        :param pulumi.Input[str] name: The query name.
        :param pulumi.Input[str] work_group: The name of the workgroup that contains the named query.
        """
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "query_string", query_string)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if work_group is not None:
            pulumi.set(__self__, "work_group", work_group)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Input[str]:
        """
        The database to which the query belongs.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: pulumi.Input[str]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> pulumi.Input[str]:
        """
        The contents of the query with all query statements.
        """
        return pulumi.get(self, "query_string")

    @query_string.setter
    def query_string(self, value: pulumi.Input[str]):
        pulumi.set(self, "query_string", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The query description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The query name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="workGroup")
    def work_group(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workgroup that contains the named query.
        """
        return pulumi.get(self, "work_group")

    @work_group.setter
    def work_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_group", value)


class NamedQuery(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 work_group: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::Athena::NamedQuery

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to which the query belongs.
        :param pulumi.Input[str] description: The query description.
        :param pulumi.Input[str] name: The query name.
        :param pulumi.Input[str] query_string: The contents of the query with all query statements.
        :param pulumi.Input[str] work_group: The name of the workgroup that contains the named query.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamedQueryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::Athena::NamedQuery

        :param str resource_name: The name of the resource.
        :param NamedQueryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamedQueryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_string: Optional[pulumi.Input[str]] = None,
                 work_group: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamedQueryArgs.__new__(NamedQueryArgs)

            if database is None and not opts.urn:
                raise TypeError("Missing required property 'database'")
            __props__.__dict__["database"] = database
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if query_string is None and not opts.urn:
                raise TypeError("Missing required property 'query_string'")
            __props__.__dict__["query_string"] = query_string
            __props__.__dict__["work_group"] = work_group
            __props__.__dict__["named_query_id"] = None
        super(NamedQuery, __self__).__init__(
            'aws-native:athena:NamedQuery',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NamedQuery':
        """
        Get an existing NamedQuery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NamedQueryArgs.__new__(NamedQueryArgs)

        __props__.__dict__["database"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["named_query_id"] = None
        __props__.__dict__["query_string"] = None
        __props__.__dict__["work_group"] = None
        return NamedQuery(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        The database to which the query belongs.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The query description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The query name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedQueryId")
    def named_query_id(self) -> pulumi.Output[str]:
        """
        The unique ID of the query.
        """
        return pulumi.get(self, "named_query_id")

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> pulumi.Output[str]:
        """
        The contents of the query with all query statements.
        """
        return pulumi.get(self, "query_string")

    @property
    @pulumi.getter(name="workGroup")
    def work_group(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the workgroup that contains the named query.
        """
        return pulumi.get(self, "work_group")
