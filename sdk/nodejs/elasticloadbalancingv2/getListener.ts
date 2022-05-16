// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticLoadBalancingV2::Listener
 */
export function getListener(args: GetListenerArgs, opts?: pulumi.InvokeOptions): Promise<GetListenerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticloadbalancingv2:getListener", {
        "listenerArn": args.listenerArn,
    }, opts);
}

export interface GetListenerArgs {
    listenerArn: string;
}

export interface GetListenerResult {
    readonly alpnPolicy?: string[];
    readonly certificates?: outputs.elasticloadbalancingv2.ListenerCertificate[];
    readonly defaultActions?: outputs.elasticloadbalancingv2.ListenerAction[];
    readonly listenerArn?: string;
    readonly port?: number;
    readonly protocol?: string;
    readonly sslPolicy?: string;
}

export function getListenerOutput(args: GetListenerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetListenerResult> {
    return pulumi.output(args).apply(a => getListener(a, opts))
}

export interface GetListenerOutputArgs {
    listenerArn: pulumi.Input<string>;
}
