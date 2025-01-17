// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ResiliencyPolicyDataLocationConstraint = {
    AnyLocation: "AnyLocation",
    SameContinent: "SameContinent",
    SameCountry: "SameCountry",
} as const;

/**
 * Data Location Constraint of the Policy.
 */
export type ResiliencyPolicyDataLocationConstraint = (typeof ResiliencyPolicyDataLocationConstraint)[keyof typeof ResiliencyPolicyDataLocationConstraint];

export const ResiliencyPolicyTier = {
    MissionCritical: "MissionCritical",
    Critical: "Critical",
    Important: "Important",
    CoreServices: "CoreServices",
    NonCritical: "NonCritical",
} as const;

/**
 * Resiliency Policy Tier.
 */
export type ResiliencyPolicyTier = (typeof ResiliencyPolicyTier)[keyof typeof ResiliencyPolicyTier];
