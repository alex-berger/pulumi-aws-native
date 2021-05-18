# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AccessPointAccessPointTag',
    'AccessPointCreationInfo',
    'AccessPointPosixUser',
    'AccessPointRootDirectory',
    'FileSystemBackupPolicy',
    'FileSystemElasticFileSystemTag',
    'FileSystemLifecyclePolicy',
]

@pulumi.output_type
class AccessPointAccessPointTag(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html
    """
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html#cfn-efs-accesspoint-accesspointtag-key
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html#cfn-efs-accesspoint-accesspointtag-value
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html#cfn-efs-accesspoint-accesspointtag-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-accesspointtag.html#cfn-efs-accesspoint-accesspointtag-value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class AccessPointCreationInfo(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ownerGid":
            suggest = "owner_gid"
        elif key == "ownerUid":
            suggest = "owner_uid"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointCreationInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointCreationInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointCreationInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 owner_gid: str,
                 owner_uid: str,
                 permissions: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html
        :param str owner_gid: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-ownergid
        :param str owner_uid: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-owneruid
        :param str permissions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-permissions
        """
        pulumi.set(__self__, "owner_gid", owner_gid)
        pulumi.set(__self__, "owner_uid", owner_uid)
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="ownerGid")
    def owner_gid(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-ownergid
        """
        return pulumi.get(self, "owner_gid")

    @property
    @pulumi.getter(name="ownerUid")
    def owner_uid(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-owneruid
        """
        return pulumi.get(self, "owner_uid")

    @property
    @pulumi.getter
    def permissions(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-creationinfo.html#cfn-efs-accesspoint-creationinfo-permissions
        """
        return pulumi.get(self, "permissions")


@pulumi.output_type
class AccessPointPosixUser(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secondaryGids":
            suggest = "secondary_gids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointPosixUser. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointPosixUser.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointPosixUser.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 gid: str,
                 uid: str,
                 secondary_gids: Optional[Sequence[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html
        :param str gid: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-gid
        :param str uid: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-uid
        :param Sequence[str] secondary_gids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-secondarygids
        """
        pulumi.set(__self__, "gid", gid)
        pulumi.set(__self__, "uid", uid)
        if secondary_gids is not None:
            pulumi.set(__self__, "secondary_gids", secondary_gids)

    @property
    @pulumi.getter
    def gid(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-gid
        """
        return pulumi.get(self, "gid")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-uid
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="secondaryGids")
    def secondary_gids(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-posixuser.html#cfn-efs-accesspoint-posixuser-secondarygids
        """
        return pulumi.get(self, "secondary_gids")


@pulumi.output_type
class AccessPointRootDirectory(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "creationInfo":
            suggest = "creation_info"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessPointRootDirectory. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessPointRootDirectory.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessPointRootDirectory.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 creation_info: Optional['outputs.AccessPointCreationInfo'] = None,
                 path: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html
        :param 'AccessPointCreationInfo' creation_info: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-creationinfo
        :param str path: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-path
        """
        if creation_info is not None:
            pulumi.set(__self__, "creation_info", creation_info)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="creationInfo")
    def creation_info(self) -> Optional['outputs.AccessPointCreationInfo']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-creationinfo
        """
        return pulumi.get(self, "creation_info")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-path
        """
        return pulumi.get(self, "path")


@pulumi.output_type
class FileSystemBackupPolicy(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html
    """
    def __init__(__self__, *,
                 status: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html
        :param str status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html#cfn-efs-filesystem-backuppolicy-status
        """
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html#cfn-efs-filesystem-backuppolicy-status
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class FileSystemElasticFileSystemTag(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html
    """
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html
        :param str key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html#cfn-efs-filesystem-elasticfilesystemtag-key
        :param str value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html#cfn-efs-filesystem-elasticfilesystemtag-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html#cfn-efs-filesystem-elasticfilesystemtag-key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-elasticfilesystemtag.html#cfn-efs-filesystem-elasticfilesystemtag-value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class FileSystemLifecyclePolicy(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "transitionToIA":
            suggest = "transition_to_ia"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FileSystemLifecyclePolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FileSystemLifecyclePolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FileSystemLifecyclePolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 transition_to_ia: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html
        :param str transition_to_ia: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html#cfn-efs-filesystem-lifecyclepolicy-transitiontoia
        """
        pulumi.set(__self__, "transition_to_ia", transition_to_ia)

    @property
    @pulumi.getter(name="transitionToIA")
    def transition_to_ia(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-lifecyclepolicy.html#cfn-efs-filesystem-lifecyclepolicy-transitiontoia
        """
        return pulumi.get(self, "transition_to_ia")


