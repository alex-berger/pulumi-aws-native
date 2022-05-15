// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DevicePoolRuleAttribute = {
    Arn: "ARN",
    Platform: "PLATFORM",
    FormFactor: "FORM_FACTOR",
    Manufacturer: "MANUFACTURER",
    RemoteAccessEnabled: "REMOTE_ACCESS_ENABLED",
    RemoteDebugEnabled: "REMOTE_DEBUG_ENABLED",
    AppiumVersion: "APPIUM_VERSION",
    InstanceArn: "INSTANCE_ARN",
    InstanceLabels: "INSTANCE_LABELS",
    FleetType: "FLEET_TYPE",
    OsVersion: "OS_VERSION",
    Model: "MODEL",
    Availability: "AVAILABILITY",
} as const;

/**
 * The rule's stringified attribute.
 */
export type DevicePoolRuleAttribute = (typeof DevicePoolRuleAttribute)[keyof typeof DevicePoolRuleAttribute];

export const DevicePoolRuleOperator = {
    Equals: "EQUALS",
    LessThan: "LESS_THAN",
    LessThanOrEquals: "LESS_THAN_OR_EQUALS",
    GreaterThan: "GREATER_THAN",
    GreaterThanOrEquals: "GREATER_THAN_OR_EQUALS",
    In: "IN",
    NotIn: "NOT_IN",
    Contains: "CONTAINS",
} as const;

/**
 * Specifies how Device Farm compares the rule's attribute to the value.
 */
export type DevicePoolRuleOperator = (typeof DevicePoolRuleOperator)[keyof typeof DevicePoolRuleOperator];
