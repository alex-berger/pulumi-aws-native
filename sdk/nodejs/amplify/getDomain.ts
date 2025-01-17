// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:amplify:getDomain", {
        "arn": args.arn,
    }, opts);
}

export interface GetDomainArgs {
    arn: string;
}

export interface GetDomainResult {
    readonly arn?: string;
    readonly autoSubDomainCreationPatterns?: string[];
    readonly autoSubDomainIAMRole?: string;
    readonly certificateRecord?: string;
    readonly domainStatus?: string;
    readonly enableAutoSubDomain?: boolean;
    readonly statusReason?: string;
    readonly subDomainSettings?: outputs.amplify.DomainSubDomainSetting[];
}

export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainResult> {
    return pulumi.output(args).apply(a => getDomain(a, opts))
}

export interface GetDomainOutputArgs {
    arn: pulumi.Input<string>;
}
