// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const HealthCheckConfigPropertiesInsufficientDataHealthStatus = {
    Healthy: "Healthy",
    LastKnownStatus: "LastKnownStatus",
    Unhealthy: "Unhealthy",
} as const;

export type HealthCheckConfigPropertiesInsufficientDataHealthStatus = (typeof HealthCheckConfigPropertiesInsufficientDataHealthStatus)[keyof typeof HealthCheckConfigPropertiesInsufficientDataHealthStatus];

export const HealthCheckConfigPropertiesType = {
    Calculated: "CALCULATED",
    CloudwatchMetric: "CLOUDWATCH_METRIC",
    Http: "HTTP",
    HttpStrMatch: "HTTP_STR_MATCH",
    Https: "HTTPS",
    HttpsStrMatch: "HTTPS_STR_MATCH",
    Tcp: "TCP",
    RecoveryControl: "RECOVERY_CONTROL",
} as const;

export type HealthCheckConfigPropertiesType = (typeof HealthCheckConfigPropertiesType)[keyof typeof HealthCheckConfigPropertiesType];

export const KeySigningKeyStatus = {
    Active: "ACTIVE",
    Inactive: "INACTIVE",
} as const;

/**
 * A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
 */
export type KeySigningKeyStatus = (typeof KeySigningKeyStatus)[keyof typeof KeySigningKeyStatus];
