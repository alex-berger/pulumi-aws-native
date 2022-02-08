// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerRule
 */
export function getListenerRule(args: GetListenerRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetListenerRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticloadbalancingv2:getListenerRule", {
        "ruleArn": args.ruleArn,
    }, opts);
}

export interface GetListenerRuleArgs {
    ruleArn: string;
}

export interface GetListenerRuleResult {
    readonly actions?: outputs.elasticloadbalancingv2.ListenerRuleAction[];
    readonly conditions?: outputs.elasticloadbalancingv2.ListenerRuleRuleCondition[];
    readonly isDefault?: boolean;
    readonly priority?: number;
    readonly ruleArn?: string;
}

export function getListenerRuleOutput(args: GetListenerRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetListenerRuleResult> {
    return pulumi.output(args).apply(a => getListenerRule(a, opts))
}

export interface GetListenerRuleOutputArgs {
    ruleArn: pulumi.Input<string>;
}