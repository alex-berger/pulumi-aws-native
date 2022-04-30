// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::ApiKey
 */
export function getApiKey(args: GetApiKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetApiKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appsync:getApiKey", {
        "apiKeyId": args.apiKeyId,
    }, opts);
}

export interface GetApiKeyArgs {
    apiKeyId: string;
}

export interface GetApiKeyResult {
    readonly apiKey?: string;
    readonly apiKeyId?: string;
    readonly arn?: string;
    readonly description?: string;
    readonly expires?: number;
}

export function getApiKeyOutput(args: GetApiKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiKeyResult> {
    return pulumi.output(args).apply(a => getApiKey(a, opts))
}

export interface GetApiKeyOutputArgs {
    apiKeyId: pulumi.Input<string>;
}
