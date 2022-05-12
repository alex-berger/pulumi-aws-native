# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 access_key: Optional[pulumi.Input[str]] = None,
                 allowed_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 assume_role: Optional[pulumi.Input['ProviderAssumeRoleArgs']] = None,
                 default_tags: Optional[pulumi.Input['ProviderDefaultTagsArgs']] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]] = None,
                 forbidden_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ignore_tags: Optional[pulumi.Input['ProviderIgnoreTagsArgs']] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 s3_force_path_style: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_get_ec2_platforms: Optional[pulumi.Input[bool]] = None,
                 skip_metadata_api_check: Optional[pulumi.Input[bool]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 skip_requesting_account_id: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] region: The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
        :param pulumi.Input[str] access_key: The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_account_ids: List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
        :param pulumi.Input['ProviderAssumeRoleArgs'] assume_role: Configuration for retrieving temporary credentials from the STS service.
        :param pulumi.Input['ProviderDefaultTagsArgs'] default_tags: Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]] endpoints: Configuration block for customizing service endpoints.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] forbidden_account_ids: List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
        :param pulumi.Input['ProviderIgnoreTagsArgs'] ignore_tags: Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
        :param pulumi.Input[bool] insecure: Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
        :param pulumi.Input[int] max_retries: The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
        :param pulumi.Input[str] profile: The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        :param pulumi.Input[bool] s3_force_path_style: Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
        :param pulumi.Input[str] secret_key: The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
        :param pulumi.Input[str] shared_credentials_file: The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
        :param pulumi.Input[bool] skip_credentials_validation: Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
        :param pulumi.Input[bool] skip_get_ec2_platforms: Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
        :param pulumi.Input[bool] skip_metadata_api_check: Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
        :param pulumi.Input[bool] skip_region_validation: Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
        :param pulumi.Input[bool] skip_requesting_account_id: Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
        :param pulumi.Input[str] token: Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
        """
        if region is None:
            region = _utilities.get_env('AWS_REGION', 'AWS_DEFAULT_REGION')
        pulumi.set(__self__, "region", region)
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if allowed_account_ids is not None:
            pulumi.set(__self__, "allowed_account_ids", allowed_account_ids)
        if assume_role is not None:
            pulumi.set(__self__, "assume_role", assume_role)
        if default_tags is not None:
            pulumi.set(__self__, "default_tags", default_tags)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if forbidden_account_ids is not None:
            pulumi.set(__self__, "forbidden_account_ids", forbidden_account_ids)
        if ignore_tags is not None:
            pulumi.set(__self__, "ignore_tags", ignore_tags)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if profile is None:
            profile = _utilities.get_env('AWS_PROFILE')
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if s3_force_path_style is not None:
            pulumi.set(__self__, "s3_force_path_style", s3_force_path_style)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if shared_credentials_file is None:
            shared_credentials_file = _utilities.get_env('AWS_SHARED_CREDENTIALS_FILE')
        if shared_credentials_file is not None:
            pulumi.set(__self__, "shared_credentials_file", shared_credentials_file)
        if skip_credentials_validation is None:
            skip_credentials_validation = True
        if skip_credentials_validation is not None:
            pulumi.set(__self__, "skip_credentials_validation", skip_credentials_validation)
        if skip_get_ec2_platforms is None:
            skip_get_ec2_platforms = True
        if skip_get_ec2_platforms is not None:
            pulumi.set(__self__, "skip_get_ec2_platforms", skip_get_ec2_platforms)
        if skip_metadata_api_check is None:
            skip_metadata_api_check = True
        if skip_metadata_api_check is not None:
            pulumi.set(__self__, "skip_metadata_api_check", skip_metadata_api_check)
        if skip_region_validation is None:
            skip_region_validation = True
        if skip_region_validation is not None:
            pulumi.set(__self__, "skip_region_validation", skip_region_validation)
        if skip_requesting_account_id is not None:
            pulumi.set(__self__, "skip_requesting_account_id", skip_requesting_account_id)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="allowedAccountIds")
    def allowed_account_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
        """
        return pulumi.get(self, "allowed_account_ids")

    @allowed_account_ids.setter
    def allowed_account_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_account_ids", value)

    @property
    @pulumi.getter(name="assumeRole")
    def assume_role(self) -> Optional[pulumi.Input['ProviderAssumeRoleArgs']]:
        """
        Configuration for retrieving temporary credentials from the STS service.
        """
        return pulumi.get(self, "assume_role")

    @assume_role.setter
    def assume_role(self, value: Optional[pulumi.Input['ProviderAssumeRoleArgs']]):
        pulumi.set(self, "assume_role", value)

    @property
    @pulumi.getter(name="defaultTags")
    def default_tags(self) -> Optional[pulumi.Input['ProviderDefaultTagsArgs']]:
        """
        Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
        """
        return pulumi.get(self, "default_tags")

    @default_tags.setter
    def default_tags(self, value: Optional[pulumi.Input['ProviderDefaultTagsArgs']]):
        pulumi.set(self, "default_tags", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]]:
        """
        Configuration block for customizing service endpoints.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter(name="forbiddenAccountIds")
    def forbidden_account_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
        """
        return pulumi.get(self, "forbidden_account_ids")

    @forbidden_account_ids.setter
    def forbidden_account_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "forbidden_account_ids", value)

    @property
    @pulumi.getter(name="ignoreTags")
    def ignore_tags(self) -> Optional[pulumi.Input['ProviderIgnoreTagsArgs']]:
        """
        Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
        """
        return pulumi.get(self, "ignore_tags")

    @ignore_tags.setter
    def ignore_tags(self, value: Optional[pulumi.Input['ProviderIgnoreTagsArgs']]):
        pulumi.set(self, "ignore_tags", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="s3ForcePathStyle")
    def s3_force_path_style(self) -> Optional[pulumi.Input[bool]]:
        """
        Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
        """
        return pulumi.get(self, "s3_force_path_style")

    @s3_force_path_style.setter
    def s3_force_path_style(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "s3_force_path_style", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
        """
        return pulumi.get(self, "shared_credentials_file")

    @shared_credentials_file.setter
    def shared_credentials_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_credentials_file", value)

    @property
    @pulumi.getter(name="skipCredentialsValidation")
    def skip_credentials_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
        """
        return pulumi.get(self, "skip_credentials_validation")

    @skip_credentials_validation.setter
    def skip_credentials_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_credentials_validation", value)

    @property
    @pulumi.getter(name="skipGetEc2Platforms")
    def skip_get_ec2_platforms(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
        """
        return pulumi.get(self, "skip_get_ec2_platforms")

    @skip_get_ec2_platforms.setter
    def skip_get_ec2_platforms(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_get_ec2_platforms", value)

    @property
    @pulumi.getter(name="skipMetadataApiCheck")
    def skip_metadata_api_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
        """
        return pulumi.get(self, "skip_metadata_api_check")

    @skip_metadata_api_check.setter
    def skip_metadata_api_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_metadata_api_check", value)

    @property
    @pulumi.getter(name="skipRegionValidation")
    def skip_region_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
        """
        return pulumi.get(self, "skip_region_validation")

    @skip_region_validation.setter
    def skip_region_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_region_validation", value)

    @property
    @pulumi.getter(name="skipRequestingAccountId")
    def skip_requesting_account_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
        """
        return pulumi.get(self, "skip_requesting_account_id")

    @skip_requesting_account_id.setter
    def skip_requesting_account_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_requesting_account_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 allowed_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 assume_role: Optional[pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']]] = None,
                 default_tags: Optional[pulumi.Input[pulumi.InputType['ProviderDefaultTagsArgs']]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]]] = None,
                 forbidden_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ignore_tags: Optional[pulumi.Input[pulumi.InputType['ProviderIgnoreTagsArgs']]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 s3_force_path_style: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_get_ec2_platforms: Optional[pulumi.Input[bool]] = None,
                 skip_metadata_api_check: Optional[pulumi.Input[bool]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 skip_requesting_account_id: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the AWS native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_account_ids: List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.
        :param pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']] assume_role: Configuration for retrieving temporary credentials from the STS service.
        :param pulumi.Input[pulumi.InputType['ProviderDefaultTagsArgs']] default_tags: Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]] endpoints: Configuration block for customizing service endpoints.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] forbidden_account_ids: List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.
        :param pulumi.Input[pulumi.InputType['ProviderIgnoreTagsArgs']] ignore_tags: Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.
        :param pulumi.Input[bool] insecure: Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`.
        :param pulumi.Input[int] max_retries: The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
        :param pulumi.Input[str] profile: The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        :param pulumi.Input[str] region: The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
        :param pulumi.Input[bool] s3_force_path_style: Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.
        :param pulumi.Input[str] secret_key: The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
        :param pulumi.Input[str] shared_credentials_file: The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
        :param pulumi.Input[bool] skip_credentials_validation: Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
        :param pulumi.Input[bool] skip_get_ec2_platforms: Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.
        :param pulumi.Input[bool] skip_metadata_api_check: Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
        :param pulumi.Input[bool] skip_region_validation: Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.
        :param pulumi.Input[bool] skip_requesting_account_id: Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
        :param pulumi.Input[str] token: Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the AWS native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 allowed_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 assume_role: Optional[pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']]] = None,
                 default_tags: Optional[pulumi.Input[pulumi.InputType['ProviderDefaultTagsArgs']]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]]] = None,
                 forbidden_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ignore_tags: Optional[pulumi.Input[pulumi.InputType['ProviderIgnoreTagsArgs']]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 s3_force_path_style: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_credentials_validation: Optional[pulumi.Input[bool]] = None,
                 skip_get_ec2_platforms: Optional[pulumi.Input[bool]] = None,
                 skip_metadata_api_check: Optional[pulumi.Input[bool]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 skip_requesting_account_id: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["access_key"] = None if access_key is None else pulumi.Output.secret(access_key)
            __props__.__dict__["allowed_account_ids"] = pulumi.Output.from_input(allowed_account_ids).apply(pulumi.runtime.to_json) if allowed_account_ids is not None else None
            __props__.__dict__["assume_role"] = pulumi.Output.from_input(assume_role).apply(pulumi.runtime.to_json) if assume_role is not None else None
            __props__.__dict__["default_tags"] = pulumi.Output.from_input(default_tags).apply(pulumi.runtime.to_json) if default_tags is not None else None
            __props__.__dict__["endpoints"] = pulumi.Output.from_input(endpoints).apply(pulumi.runtime.to_json) if endpoints is not None else None
            __props__.__dict__["forbidden_account_ids"] = pulumi.Output.from_input(forbidden_account_ids).apply(pulumi.runtime.to_json) if forbidden_account_ids is not None else None
            __props__.__dict__["ignore_tags"] = pulumi.Output.from_input(ignore_tags).apply(pulumi.runtime.to_json) if ignore_tags is not None else None
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            __props__.__dict__["max_retries"] = pulumi.Output.from_input(max_retries).apply(pulumi.runtime.to_json) if max_retries is not None else None
            if profile is None:
                profile = _utilities.get_env('AWS_PROFILE')
            __props__.__dict__["profile"] = profile
            if region is None:
                region = _utilities.get_env('AWS_REGION', 'AWS_DEFAULT_REGION')
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["s3_force_path_style"] = pulumi.Output.from_input(s3_force_path_style).apply(pulumi.runtime.to_json) if s3_force_path_style is not None else None
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            if shared_credentials_file is None:
                shared_credentials_file = _utilities.get_env('AWS_SHARED_CREDENTIALS_FILE')
            __props__.__dict__["shared_credentials_file"] = shared_credentials_file
            if skip_credentials_validation is None:
                skip_credentials_validation = True
            __props__.__dict__["skip_credentials_validation"] = pulumi.Output.from_input(skip_credentials_validation).apply(pulumi.runtime.to_json) if skip_credentials_validation is not None else None
            if skip_get_ec2_platforms is None:
                skip_get_ec2_platforms = True
            __props__.__dict__["skip_get_ec2_platforms"] = pulumi.Output.from_input(skip_get_ec2_platforms).apply(pulumi.runtime.to_json) if skip_get_ec2_platforms is not None else None
            if skip_metadata_api_check is None:
                skip_metadata_api_check = True
            __props__.__dict__["skip_metadata_api_check"] = pulumi.Output.from_input(skip_metadata_api_check).apply(pulumi.runtime.to_json) if skip_metadata_api_check is not None else None
            if skip_region_validation is None:
                skip_region_validation = True
            __props__.__dict__["skip_region_validation"] = pulumi.Output.from_input(skip_region_validation).apply(pulumi.runtime.to_json) if skip_region_validation is not None else None
            __props__.__dict__["skip_requesting_account_id"] = pulumi.Output.from_input(skip_requesting_account_id).apply(pulumi.runtime.to_json) if skip_requesting_account_id is not None else None
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
        super(Provider, __self__).__init__(
            'aws-native',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[Optional[str]]:
        """
        The profile for API operations. If not set, the default profile created with `aws configure` will be used.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.
        """
        return pulumi.get(self, "shared_credentials_file")

