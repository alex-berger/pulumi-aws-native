// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::NetworkFirewall::FirewallPolicy
 */
export function getFirewallPolicy(args: GetFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:networkfirewall:getFirewallPolicy", {
        "firewallPolicyArn": args.firewallPolicyArn,
    }, opts);
}

export interface GetFirewallPolicyArgs {
    firewallPolicyArn: string;
}

export interface GetFirewallPolicyResult {
    readonly description?: string;
    readonly firewallPolicy?: outputs.networkfirewall.FirewallPolicy;
    readonly firewallPolicyArn?: string;
    readonly firewallPolicyId?: string;
    readonly tags?: outputs.networkfirewall.FirewallPolicyTag[];
}

export function getFirewallPolicyOutput(args: GetFirewallPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallPolicyResult> {
    return pulumi.output(args).apply(a => getFirewallPolicy(a, opts))
}

export interface GetFirewallPolicyOutputArgs {
    firewallPolicyArn: pulumi.Input<string>;
}
