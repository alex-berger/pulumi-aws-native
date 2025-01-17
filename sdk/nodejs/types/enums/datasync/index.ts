// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AgentEndpointType = {
    Fips: "FIPS",
    Public: "PUBLIC",
    PrivateLink: "PRIVATE_LINK",
} as const;

/**
 * The service endpoints that the agent will connect to.
 */
export type AgentEndpointType = (typeof AgentEndpointType)[keyof typeof AgentEndpointType];

export const LocationFSxOpenZFSMountOptionsVersion = {
    Automatic: "AUTOMATIC",
    Nfs3: "NFS3",
    Nfs40: "NFS4_0",
    Nfs41: "NFS4_1",
} as const;

/**
 * The specific NFS version that you want DataSync to use to mount your NFS share.
 */
export type LocationFSxOpenZFSMountOptionsVersion = (typeof LocationFSxOpenZFSMountOptionsVersion)[keyof typeof LocationFSxOpenZFSMountOptionsVersion];

export const LocationHDFSAuthenticationType = {
    Simple: "SIMPLE",
    Kerberos: "KERBEROS",
} as const;

/**
 * The authentication mode used to determine identity of user.
 */
export type LocationHDFSAuthenticationType = (typeof LocationHDFSAuthenticationType)[keyof typeof LocationHDFSAuthenticationType];

export const LocationHDFSQopConfigurationDataTransferProtection = {
    Authentication: "AUTHENTICATION",
    Integrity: "INTEGRITY",
    Privacy: "PRIVACY",
    Disabled: "DISABLED",
} as const;

/**
 * Configuration for Data Transfer Protection.
 */
export type LocationHDFSQopConfigurationDataTransferProtection = (typeof LocationHDFSQopConfigurationDataTransferProtection)[keyof typeof LocationHDFSQopConfigurationDataTransferProtection];

export const LocationHDFSQopConfigurationRpcProtection = {
    Authentication: "AUTHENTICATION",
    Integrity: "INTEGRITY",
    Privacy: "PRIVACY",
    Disabled: "DISABLED",
} as const;

/**
 * Configuration for RPC Protection.
 */
export type LocationHDFSQopConfigurationRpcProtection = (typeof LocationHDFSQopConfigurationRpcProtection)[keyof typeof LocationHDFSQopConfigurationRpcProtection];

export const LocationNFSMountOptionsVersion = {
    Automatic: "AUTOMATIC",
    Nfs3: "NFS3",
    Nfs40: "NFS4_0",
    Nfs41: "NFS4_1",
} as const;

/**
 * The specific NFS version that you want DataSync to use to mount your NFS share.
 */
export type LocationNFSMountOptionsVersion = (typeof LocationNFSMountOptionsVersion)[keyof typeof LocationNFSMountOptionsVersion];

export const LocationObjectStorageServerProtocol = {
    Https: "HTTPS",
    Http: "HTTP",
} as const;

/**
 * The protocol that the object storage server uses to communicate.
 */
export type LocationObjectStorageServerProtocol = (typeof LocationObjectStorageServerProtocol)[keyof typeof LocationObjectStorageServerProtocol];

export const LocationS3S3StorageClass = {
    Standard: "STANDARD",
    StandardIa: "STANDARD_IA",
    OnezoneIa: "ONEZONE_IA",
    IntelligentTiering: "INTELLIGENT_TIERING",
    Glacier: "GLACIER",
    DeepArchive: "DEEP_ARCHIVE",
} as const;

/**
 * The Amazon S3 storage class you want to store your files in when this location is used as a task destination.
 */
export type LocationS3S3StorageClass = (typeof LocationS3S3StorageClass)[keyof typeof LocationS3S3StorageClass];

export const LocationSMBMountOptionsVersion = {
    Automatic: "AUTOMATIC",
    Smb2: "SMB2",
    Smb3: "SMB3",
} as const;

/**
 * The specific SMB version that you want DataSync to use to mount your SMB share.
 */
export type LocationSMBMountOptionsVersion = (typeof LocationSMBMountOptionsVersion)[keyof typeof LocationSMBMountOptionsVersion];

export const TaskFilterRuleFilterType = {
    SimplePattern: "SIMPLE_PATTERN",
} as const;

/**
 * The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.
 */
export type TaskFilterRuleFilterType = (typeof TaskFilterRuleFilterType)[keyof typeof TaskFilterRuleFilterType];

export const TaskOptionsAtime = {
    None: "NONE",
    BestEffort: "BEST_EFFORT",
} as const;

/**
 * A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).
 */
export type TaskOptionsAtime = (typeof TaskOptionsAtime)[keyof typeof TaskOptionsAtime];

export const TaskOptionsGid = {
    None: "NONE",
    IntValue: "INT_VALUE",
    Name: "NAME",
    Both: "BOTH",
} as const;

/**
 * The group ID (GID) of the file's owners.
 */
export type TaskOptionsGid = (typeof TaskOptionsGid)[keyof typeof TaskOptionsGid];

export const TaskOptionsLogLevel = {
    Off: "OFF",
    Basic: "BASIC",
    Transfer: "TRANSFER",
} as const;

/**
 * A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.
 */
export type TaskOptionsLogLevel = (typeof TaskOptionsLogLevel)[keyof typeof TaskOptionsLogLevel];

export const TaskOptionsMtime = {
    None: "NONE",
    Preserve: "PRESERVE",
} as const;

/**
 * A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.
 */
export type TaskOptionsMtime = (typeof TaskOptionsMtime)[keyof typeof TaskOptionsMtime];

export const TaskOptionsOverwriteMode = {
    Always: "ALWAYS",
    Never: "NEVER",
} as const;

/**
 * A value that determines whether files at the destination should be overwritten or preserved when copying files.
 */
export type TaskOptionsOverwriteMode = (typeof TaskOptionsOverwriteMode)[keyof typeof TaskOptionsOverwriteMode];

export const TaskOptionsPosixPermissions = {
    None: "NONE",
    Preserve: "PRESERVE",
} as const;

/**
 * A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.
 */
export type TaskOptionsPosixPermissions = (typeof TaskOptionsPosixPermissions)[keyof typeof TaskOptionsPosixPermissions];

export const TaskOptionsPreserveDeletedFiles = {
    Preserve: "PRESERVE",
    Remove: "REMOVE",
} as const;

/**
 * A value that specifies whether files in the destination that don't exist in the source file system should be preserved.
 */
export type TaskOptionsPreserveDeletedFiles = (typeof TaskOptionsPreserveDeletedFiles)[keyof typeof TaskOptionsPreserveDeletedFiles];

export const TaskOptionsPreserveDevices = {
    None: "NONE",
    Preserve: "PRESERVE",
} as const;

/**
 * A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.
 */
export type TaskOptionsPreserveDevices = (typeof TaskOptionsPreserveDevices)[keyof typeof TaskOptionsPreserveDevices];

export const TaskOptionsSecurityDescriptorCopyFlags = {
    None: "NONE",
    OwnerDacl: "OWNER_DACL",
    OwnerDaclSacl: "OWNER_DACL_SACL",
} as const;

/**
 * A value that determines which components of the SMB security descriptor are copied during transfer.
 */
export type TaskOptionsSecurityDescriptorCopyFlags = (typeof TaskOptionsSecurityDescriptorCopyFlags)[keyof typeof TaskOptionsSecurityDescriptorCopyFlags];

export const TaskOptionsTaskQueueing = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

/**
 * A value that determines whether tasks should be queued before executing the tasks.
 */
export type TaskOptionsTaskQueueing = (typeof TaskOptionsTaskQueueing)[keyof typeof TaskOptionsTaskQueueing];

export const TaskOptionsTransferMode = {
    Changed: "CHANGED",
    All: "ALL",
} as const;

/**
 * A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.
 */
export type TaskOptionsTransferMode = (typeof TaskOptionsTransferMode)[keyof typeof TaskOptionsTransferMode];

export const TaskOptionsUid = {
    None: "NONE",
    IntValue: "INT_VALUE",
    Name: "NAME",
    Both: "BOTH",
} as const;

/**
 * The user ID (UID) of the file's owner.
 */
export type TaskOptionsUid = (typeof TaskOptionsUid)[keyof typeof TaskOptionsUid];

export const TaskOptionsVerifyMode = {
    PointInTimeConsistent: "POINT_IN_TIME_CONSISTENT",
    OnlyFilesTransferred: "ONLY_FILES_TRANSFERRED",
    None: "NONE",
} as const;

/**
 * A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.
 */
export type TaskOptionsVerifyMode = (typeof TaskOptionsVerifyMode)[keyof typeof TaskOptionsVerifyMode];

export const TaskStatus = {
    Available: "AVAILABLE",
    Creating: "CREATING",
    Queued: "QUEUED",
    Running: "RUNNING",
    Unavailable: "UNAVAILABLE",
} as const;

/**
 * The status of the task that was described.
 */
export type TaskStatus = (typeof TaskStatus)[keyof typeof TaskStatus];
