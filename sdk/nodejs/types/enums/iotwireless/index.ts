// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DestinationExpressionType = {
    RuleName: "RuleName",
    MqttTopic: "MqttTopic",
} as const;

/**
 * Must be RuleName
 */
export type DestinationExpressionType = (typeof DestinationExpressionType)[keyof typeof DestinationExpressionType];

export const PartnerAccountPartnerType = {
    Sidewalk: "Sidewalk",
} as const;

/**
 * The partner type
 */
export type PartnerAccountPartnerType = (typeof PartnerAccountPartnerType)[keyof typeof PartnerAccountPartnerType];

export const TaskDefinitionType = {
    Update: "UPDATE",
} as const;

/**
 * A filter to list only the wireless gateway task definitions that use this task definition type
 */
export type TaskDefinitionType = (typeof TaskDefinitionType)[keyof typeof TaskDefinitionType];

export const WirelessDeviceType = {
    Sidewalk: "Sidewalk",
    LoRaWAN: "LoRaWAN",
} as const;

/**
 * Wireless device type, currently only Sidewalk and LoRa
 */
export type WirelessDeviceType = (typeof WirelessDeviceType)[keyof typeof WirelessDeviceType];
