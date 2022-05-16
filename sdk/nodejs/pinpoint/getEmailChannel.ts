// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::EmailChannel
 */
export function getEmailChannel(args: GetEmailChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetEmailChannelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpoint:getEmailChannel", {
        "id": args.id,
    }, opts);
}

export interface GetEmailChannelArgs {
    id: string;
}

export interface GetEmailChannelResult {
    readonly configurationSet?: string;
    readonly enabled?: boolean;
    readonly fromAddress?: string;
    readonly id?: string;
    readonly identity?: string;
    readonly roleArn?: string;
}

export function getEmailChannelOutput(args: GetEmailChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEmailChannelResult> {
    return pulumi.output(args).apply(a => getEmailChannel(a, opts))
}

export interface GetEmailChannelOutputArgs {
    id: pulumi.Input<string>;
}