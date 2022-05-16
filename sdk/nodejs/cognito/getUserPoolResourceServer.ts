// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPoolResourceServer
 */
export function getUserPoolResourceServer(args: GetUserPoolResourceServerArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolResourceServerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cognito:getUserPoolResourceServer", {
        "id": args.id,
    }, opts);
}

export interface GetUserPoolResourceServerArgs {
    id: string;
}

export interface GetUserPoolResourceServerResult {
    readonly id?: string;
    readonly name?: string;
    readonly scopes?: outputs.cognito.UserPoolResourceServerResourceServerScopeType[];
}

export function getUserPoolResourceServerOutput(args: GetUserPoolResourceServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPoolResourceServerResult> {
    return pulumi.output(args).apply(a => getUserPoolResourceServer(a, opts))
}

export interface GetUserPoolResourceServerOutputArgs {
    id: pulumi.Input<string>;
}