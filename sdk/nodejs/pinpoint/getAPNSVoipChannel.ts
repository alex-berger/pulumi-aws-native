// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::APNSVoipChannel
 */
export function getAPNSVoipChannel(args: GetAPNSVoipChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetAPNSVoipChannelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpoint:getAPNSVoipChannel", {
        "id": args.id,
    }, opts);
}

export interface GetAPNSVoipChannelArgs {
    id: string;
}

export interface GetAPNSVoipChannelResult {
    readonly bundleId?: string;
    readonly certificate?: string;
    readonly defaultAuthenticationMethod?: string;
    readonly enabled?: boolean;
    readonly id?: string;
    readonly privateKey?: string;
    readonly teamId?: string;
    readonly tokenKey?: string;
    readonly tokenKeyId?: string;
}

export function getAPNSVoipChannelOutput(args: GetAPNSVoipChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAPNSVoipChannelResult> {
    return pulumi.output(args).apply(a => getAPNSVoipChannel(a, opts))
}

export interface GetAPNSVoipChannelOutputArgs {
    id: pulumi.Input<string>;
}
