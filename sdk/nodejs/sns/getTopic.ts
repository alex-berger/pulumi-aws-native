// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SNS::Topic
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sns:getTopic", {
        "id": args.id,
    }, opts);
}

export interface GetTopicArgs {
    id: string;
}

export interface GetTopicResult {
    readonly contentBasedDeduplication?: boolean;
    readonly displayName?: string;
    readonly id?: string;
    readonly kmsMasterKeyId?: string;
    readonly subscription?: outputs.sns.TopicSubscription[];
    readonly tags?: outputs.sns.TopicTag[];
}

export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply(a => getTopic(a, opts))
}

export interface GetTopicOutputArgs {
    id: pulumi.Input<string>;
}
