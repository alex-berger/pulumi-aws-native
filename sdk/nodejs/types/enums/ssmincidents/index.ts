// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ResponsePlanSsmAutomationTargetAccount = {
    ImpactedAccount: "IMPACTED_ACCOUNT",
    ResponsePlanOwnerAccount: "RESPONSE_PLAN_OWNER_ACCOUNT",
} as const;

/**
 * The account type to use when starting the SSM automation document.
 */
export type ResponsePlanSsmAutomationTargetAccount = (typeof ResponsePlanSsmAutomationTargetAccount)[keyof typeof ResponsePlanSsmAutomationTargetAccount];
