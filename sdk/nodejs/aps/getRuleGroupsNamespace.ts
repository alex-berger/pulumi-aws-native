// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * RuleGroupsNamespace schema for cloudformation.
 */
export function getRuleGroupsNamespace(args: GetRuleGroupsNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleGroupsNamespaceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:aps:getRuleGroupsNamespace", {
        "arn": args.arn,
    }, opts);
}

export interface GetRuleGroupsNamespaceArgs {
    /**
     * The RuleGroupsNamespace ARN.
     */
    arn: string;
}

export interface GetRuleGroupsNamespaceResult {
    /**
     * The RuleGroupsNamespace ARN.
     */
    readonly arn?: string;
    /**
     * The RuleGroupsNamespace data.
     */
    readonly data?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.aps.RuleGroupsNamespaceTag[];
    /**
     * Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
     */
    readonly workspace?: string;
}

export function getRuleGroupsNamespaceOutput(args: GetRuleGroupsNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleGroupsNamespaceResult> {
    return pulumi.output(args).apply(a => getRuleGroupsNamespace(a, opts))
}

export interface GetRuleGroupsNamespaceOutputArgs {
    /**
     * The RuleGroupsNamespace ARN.
     */
    arn: pulumi.Input<string>;
}
