// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EnvironmentFederationMode = {
    Local: "LOCAL",
    Federated: "FEDERATED",
} as const;

/**
 * Federation mode used with the Environment
 */
export type EnvironmentFederationMode = (typeof EnvironmentFederationMode)[keyof typeof EnvironmentFederationMode];

export const EnvironmentStatus = {
    CreateRequested: "CREATE_REQUESTED",
    Creating: "CREATING",
    Created: "CREATED",
    DeleteRequested: "DELETE_REQUESTED",
    Deleting: "DELETING",
    Deleted: "DELETED",
    FailedCreation: "FAILED_CREATION",
    FailedDeletion: "FAILED_DELETION",
    RetryDeletion: "RETRY_DELETION",
    Suspended: "SUSPENDED",
} as const;

/**
 * State of the Environment
 */
export type EnvironmentStatus = (typeof EnvironmentStatus)[keyof typeof EnvironmentStatus];
