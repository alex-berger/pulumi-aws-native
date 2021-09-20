// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuthorizerStatus = {
    Active: "ACTIVE",
    Inactive: "INACTIVE",
} as const;

export type AuthorizerStatus = (typeof AuthorizerStatus)[keyof typeof AuthorizerStatus];

export const CertificateCertificateMode = {
    Default: "DEFAULT",
    SniOnly: "SNI_ONLY",
} as const;

export type CertificateCertificateMode = (typeof CertificateCertificateMode)[keyof typeof CertificateCertificateMode];

export const CertificateStatus = {
    Active: "ACTIVE",
    Inactive: "INACTIVE",
    Revoked: "REVOKED",
    PendingTransfer: "PENDING_TRANSFER",
    PendingActivation: "PENDING_ACTIVATION",
} as const;

export type CertificateStatus = (typeof CertificateStatus)[keyof typeof CertificateStatus];

export const CustomMetricMetricType = {
    StringList: "string-list",
    IpAddressList: "ip-address-list",
    NumberList: "number-list",
    Number: "number",
} as const;

/**
 * The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
 */
export type CustomMetricMetricType = (typeof CustomMetricMetricType)[keyof typeof CustomMetricMetricType];

export const DimensionType = {
    TopicFilter: "TOPIC_FILTER",
} as const;

/**
 * Specifies the type of the dimension.
 */
export type DimensionType = (typeof DimensionType)[keyof typeof DimensionType];

export const DomainConfigurationDomainConfigurationStatus = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

export type DomainConfigurationDomainConfigurationStatus = (typeof DomainConfigurationDomainConfigurationStatus)[keyof typeof DomainConfigurationDomainConfigurationStatus];

export const DomainConfigurationDomainType = {
    Endpoint: "ENDPOINT",
    AwsManaged: "AWS_MANAGED",
    CustomerManaged: "CUSTOMER_MANAGED",
} as const;

export type DomainConfigurationDomainType = (typeof DomainConfigurationDomainType)[keyof typeof DomainConfigurationDomainType];

export const DomainConfigurationServerCertificateSummaryServerCertificateStatus = {
    Invalid: "INVALID",
    Valid: "VALID",
} as const;

export type DomainConfigurationServerCertificateSummaryServerCertificateStatus = (typeof DomainConfigurationServerCertificateSummaryServerCertificateStatus)[keyof typeof DomainConfigurationServerCertificateSummaryServerCertificateStatus];

export const DomainConfigurationServiceType = {
    Data: "DATA",
    CredentialProvider: "CREDENTIAL_PROVIDER",
    Jobs: "JOBS",
} as const;

export type DomainConfigurationServiceType = (typeof DomainConfigurationServiceType)[keyof typeof DomainConfigurationServiceType];

export const MitigationActionEnableIoTLoggingParamsLogLevel = {
    Debug: "DEBUG",
    Info: "INFO",
    Error: "ERROR",
    Warn: "WARN",
} as const;

/**
 *  Specifies which types of information are logged.
 */
export type MitigationActionEnableIoTLoggingParamsLogLevel = (typeof MitigationActionEnableIoTLoggingParamsLogLevel)[keyof typeof MitigationActionEnableIoTLoggingParamsLogLevel];

export const MitigationActionReplaceDefaultPolicyVersionParamsTemplateName = {
    BlankPolicy: "BLANK_POLICY",
} as const;

export type MitigationActionReplaceDefaultPolicyVersionParamsTemplateName = (typeof MitigationActionReplaceDefaultPolicyVersionParamsTemplateName)[keyof typeof MitigationActionReplaceDefaultPolicyVersionParamsTemplateName];

export const MitigationActionUpdateCACertificateParamsAction = {
    Deactivate: "DEACTIVATE",
} as const;

export type MitigationActionUpdateCACertificateParamsAction = (typeof MitigationActionUpdateCACertificateParamsAction)[keyof typeof MitigationActionUpdateCACertificateParamsAction];

export const MitigationActionUpdateDeviceCertificateParamsAction = {
    Deactivate: "DEACTIVATE",
} as const;

export type MitigationActionUpdateDeviceCertificateParamsAction = (typeof MitigationActionUpdateDeviceCertificateParamsAction)[keyof typeof MitigationActionUpdateDeviceCertificateParamsAction];

export const ScheduledAuditDayOfWeek = {
    Sun: "SUN",
    Mon: "MON",
    Tue: "TUE",
    Wed: "WED",
    Thu: "THU",
    Fri: "FRI",
    Sat: "SAT",
} as const;

/**
 * The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.
 */
export type ScheduledAuditDayOfWeek = (typeof ScheduledAuditDayOfWeek)[keyof typeof ScheduledAuditDayOfWeek];

export const ScheduledAuditFrequency = {
    Daily: "DAILY",
    Weekly: "WEEKLY",
    Biweekly: "BIWEEKLY",
    Monthly: "MONTHLY",
} as const;

/**
 * How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.
 */
export type ScheduledAuditFrequency = (typeof ScheduledAuditFrequency)[keyof typeof ScheduledAuditFrequency];

export const SecurityProfileBehaviorCriteriaComparisonOperator = {
    LessThan: "less-than",
    LessThanEquals: "less-than-equals",
    GreaterThan: "greater-than",
    GreaterThanEquals: "greater-than-equals",
    InCidrSet: "in-cidr-set",
    NotInCidrSet: "not-in-cidr-set",
    InPortSet: "in-port-set",
    NotInPortSet: "not-in-port-set",
    InSet: "in-set",
    NotInSet: "not-in-set",
} as const;

/**
 * The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).
 */
export type SecurityProfileBehaviorCriteriaComparisonOperator = (typeof SecurityProfileBehaviorCriteriaComparisonOperator)[keyof typeof SecurityProfileBehaviorCriteriaComparisonOperator];

export const SecurityProfileMachineLearningDetectionConfigConfidenceLevel = {
    Low: "LOW",
    Medium: "MEDIUM",
    High: "HIGH",
} as const;

/**
 * The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.
 */
export type SecurityProfileMachineLearningDetectionConfigConfidenceLevel = (typeof SecurityProfileMachineLearningDetectionConfigConfidenceLevel)[keyof typeof SecurityProfileMachineLearningDetectionConfigConfidenceLevel];

export const SecurityProfileMetricDimensionOperator = {
    In: "IN",
    NotIn: "NOT_IN",
} as const;

/**
 * Defines how the dimensionValues of a dimension are interpreted.
 */
export type SecurityProfileMetricDimensionOperator = (typeof SecurityProfileMetricDimensionOperator)[keyof typeof SecurityProfileMetricDimensionOperator];

export const SecurityProfileStatisticalThresholdStatistic = {
    Average: "Average",
    P0: "p0",
    P01: "p0.1",
    P001: "p0.01",
    P1: "p1",
    P10: "p10",
    P50: "p50",
    P90: "p90",
    P99: "p99",
    P999: "p99.9",
    P9999: "p99.99",
    P100: "p100",
} as const;

/**
 * The percentile which resolves to a threshold value by which compliance with a behavior is determined
 */
export type SecurityProfileStatisticalThresholdStatistic = (typeof SecurityProfileStatisticalThresholdStatistic)[keyof typeof SecurityProfileStatisticalThresholdStatistic];

export const TopicRuleCannedAccessControlList = {
    Private: "private",
    PublicRead: "public-read",
    PublicReadWrite: "public-read-write",
    AwsExecRead: "aws-exec-read",
    AuthenticatedRead: "authenticated-read",
    BucketOwnerRead: "bucket-owner-read",
    BucketOwnerFullControl: "bucket-owner-full-control",
    LogDeliveryWrite: "log-delivery-write",
} as const;

export type TopicRuleCannedAccessControlList = (typeof TopicRuleCannedAccessControlList)[keyof typeof TopicRuleCannedAccessControlList];

export const TopicRuleDestinationTopicRuleDestinationStatus = {
    Enabled: "ENABLED",
    InProgress: "IN_PROGRESS",
    Disabled: "DISABLED",
} as const;

export type TopicRuleDestinationTopicRuleDestinationStatus = (typeof TopicRuleDestinationTopicRuleDestinationStatus)[keyof typeof TopicRuleDestinationTopicRuleDestinationStatus];