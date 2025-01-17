// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Account
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getAccount", {
        "id": args.id,
    }, opts);
}

export interface GetAccountArgs {
    /**
     * Primary identifier which is manually generated.
     */
    id: string;
}

export interface GetAccountResult {
    /**
     * The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.
     */
    readonly cloudWatchRoleArn?: string;
    /**
     * Primary identifier which is manually generated.
     */
    readonly id?: string;
}

export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(args).apply(a => getAccount(a, opts))
}

export interface GetAccountOutputArgs {
    /**
     * Primary identifier which is manually generated.
     */
    id: pulumi.Input<string>;
}
