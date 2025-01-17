// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::CertificateManager::Account.
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:certificatemanager:getAccount", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetAccountArgs {
    accountId: string;
}

export interface GetAccountResult {
    readonly accountId?: string;
    readonly expiryEventsConfiguration?: outputs.certificatemanager.AccountExpiryEventsConfiguration;
}

export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(args).apply(a => getAccount(a, opts))
}

export interface GetAccountOutputArgs {
    accountId: pulumi.Input<string>;
}
