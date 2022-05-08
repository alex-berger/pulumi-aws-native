// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:voiceid:getDomain", {
        "domainId": args.domainId,
    }, opts);
}

export interface GetDomainArgs {
    domainId: string;
}

export interface GetDomainResult {
    readonly domainId?: string;
    readonly tags?: outputs.voiceid.DomainTag[];
}

export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainResult> {
    return pulumi.output(args).apply(a => getDomain(a, opts))
}

export interface GetDomainOutputArgs {
    domainId: pulumi.Input<string>;
}
