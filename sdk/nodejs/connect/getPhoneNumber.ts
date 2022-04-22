// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Connect::PhoneNumber
 */
export function getPhoneNumber(args: GetPhoneNumberArgs, opts?: pulumi.InvokeOptions): Promise<GetPhoneNumberResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:connect:getPhoneNumber", {
        "phoneNumberArn": args.phoneNumberArn,
    }, opts);
}

export interface GetPhoneNumberArgs {
    /**
     * The phone number ARN
     */
    phoneNumberArn: string;
}

export interface GetPhoneNumberResult {
    /**
     * The phone number e164 address.
     */
    readonly address?: string;
    /**
     * The phone number ARN
     */
    readonly phoneNumberArn?: string;
    /**
     * One or more tags.
     */
    readonly tags?: outputs.connect.PhoneNumberTag[];
    /**
     * The ARN of the Amazon Connect instance the phone number is claimed to.
     */
    readonly targetArn?: string;
}

export function getPhoneNumberOutput(args: GetPhoneNumberOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPhoneNumberResult> {
    return pulumi.output(args).apply(a => getPhoneNumber(a, opts))
}

export interface GetPhoneNumberOutputArgs {
    /**
     * The phone number ARN
     */
    phoneNumberArn: pulumi.Input<string>;
}