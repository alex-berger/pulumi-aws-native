// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::ApiKey
 */
export function getApiKey(args: GetApiKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetApiKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getApiKey", {
        "aPIKeyId": args.aPIKeyId,
    }, opts);
}

export interface GetApiKeyArgs {
    /**
     * A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
     */
    aPIKeyId: string;
}

export interface GetApiKeyResult {
    /**
     * A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
     */
    readonly aPIKeyId?: string;
    /**
     * An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
     */
    readonly customerId?: string;
    /**
     * A description of the purpose of the API key.
     */
    readonly description?: string;
    /**
     * Indicates whether the API key can be used by clients.
     */
    readonly enabled?: boolean;
    /**
     * A list of stages to associate with this API key.
     */
    readonly stageKeys?: outputs.apigateway.ApiKeyStageKey[];
    /**
     * An array of arbitrary tags (key-value pairs) to associate with the API key.
     */
    readonly tags?: outputs.apigateway.ApiKeyTag[];
}

export function getApiKeyOutput(args: GetApiKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiKeyResult> {
    return pulumi.output(args).apply(a => getApiKey(a, opts))
}

export interface GetApiKeyOutputArgs {
    /**
     * A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs 
     */
    aPIKeyId: pulumi.Input<string>;
}
