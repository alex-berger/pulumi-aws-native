// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Logs::LogStream
 */
export function getLogStream(args: GetLogStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetLogStreamResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:logs:getLogStream", {
        "id": args.id,
    }, opts);
}

export interface GetLogStreamArgs {
    id: string;
}

export interface GetLogStreamResult {
    readonly id?: string;
}

export function getLogStreamOutput(args: GetLogStreamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogStreamResult> {
    return pulumi.output(args).apply(a => getLogStream(a, opts))
}

export interface GetLogStreamOutputArgs {
    id: pulumi.Input<string>;
}
