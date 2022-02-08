// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::IdentityPool
 */
export function getIdentityPool(args: GetIdentityPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityPoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cognito:getIdentityPool", {
        "id": args.id,
    }, opts);
}

export interface GetIdentityPoolArgs {
    id: string;
}

export interface GetIdentityPoolResult {
    readonly allowClassicFlow?: boolean;
    readonly allowUnauthenticatedIdentities?: boolean;
    readonly cognitoEvents?: any;
    readonly cognitoIdentityProviders?: outputs.cognito.IdentityPoolCognitoIdentityProvider[];
    readonly cognitoStreams?: outputs.cognito.IdentityPoolCognitoStreams;
    readonly developerProviderName?: string;
    readonly id?: string;
    readonly identityPoolName?: string;
    readonly name?: string;
    readonly openIdConnectProviderARNs?: string[];
    readonly pushSync?: outputs.cognito.IdentityPoolPushSync;
    readonly samlProviderARNs?: string[];
    readonly supportedLoginProviders?: any;
}

export function getIdentityPoolOutput(args: GetIdentityPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentityPoolResult> {
    return pulumi.output(args).apply(a => getIdentityPool(a, opts))
}

export interface GetIdentityPoolOutputArgs {
    id: pulumi.Input<string>;
}