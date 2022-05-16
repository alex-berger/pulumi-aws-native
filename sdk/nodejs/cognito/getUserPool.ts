// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPool
 */
export function getUserPool(args: GetUserPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cognito:getUserPool", {
        "id": args.id,
    }, opts);
}

export interface GetUserPoolArgs {
    id: string;
}

export interface GetUserPoolResult {
    readonly accountRecoverySetting?: outputs.cognito.UserPoolAccountRecoverySetting;
    readonly adminCreateUserConfig?: outputs.cognito.UserPoolAdminCreateUserConfig;
    readonly aliasAttributes?: string[];
    readonly arn?: string;
    readonly autoVerifiedAttributes?: string[];
    readonly deviceConfiguration?: outputs.cognito.UserPoolDeviceConfiguration;
    readonly emailConfiguration?: outputs.cognito.UserPoolEmailConfiguration;
    readonly emailVerificationMessage?: string;
    readonly emailVerificationSubject?: string;
    readonly enabledMfas?: string[];
    readonly id?: string;
    readonly lambdaConfig?: outputs.cognito.UserPoolLambdaConfig;
    readonly mfaConfiguration?: string;
    readonly policies?: outputs.cognito.UserPoolPolicies;
    readonly providerName?: string;
    readonly providerURL?: string;
    readonly schema?: outputs.cognito.UserPoolSchemaAttribute[];
    readonly smsAuthenticationMessage?: string;
    readonly smsConfiguration?: outputs.cognito.UserPoolSmsConfiguration;
    readonly smsVerificationMessage?: string;
    readonly userPoolAddOns?: outputs.cognito.UserPoolAddOns;
    readonly userPoolName?: string;
    readonly userPoolTags?: any;
    readonly usernameAttributes?: string[];
    readonly usernameConfiguration?: outputs.cognito.UserPoolUsernameConfiguration;
    readonly verificationMessageTemplate?: outputs.cognito.UserPoolVerificationMessageTemplate;
}

export function getUserPoolOutput(args: GetUserPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPoolResult> {
    return pulumi.output(args).apply(a => getUserPool(a, opts))
}

export interface GetUserPoolOutputArgs {
    id: pulumi.Input<string>;
}
