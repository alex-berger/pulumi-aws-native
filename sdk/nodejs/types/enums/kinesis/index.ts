// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const StreamEncryptionEncryptionType = {
    Kms: "KMS",
} as const;

/**
 * The encryption type to use. The only valid value is KMS. 
 */
export type StreamEncryptionEncryptionType = (typeof StreamEncryptionEncryptionType)[keyof typeof StreamEncryptionEncryptionType];

export const StreamModeDetailsStreamMode = {
    OnDemand: "ON_DEMAND",
    Provisioned: "PROVISIONED",
} as const;

/**
 * The mode of the stream
 */
export type StreamModeDetailsStreamMode = (typeof StreamModeDetailsStreamMode)[keyof typeof StreamModeDetailsStreamMode];
