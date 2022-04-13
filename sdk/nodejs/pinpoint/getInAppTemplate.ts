// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::InAppTemplate
 */
export function getInAppTemplate(args: GetInAppTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetInAppTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpoint:getInAppTemplate", {
        "templateName": args.templateName,
    }, opts);
}

export interface GetInAppTemplateArgs {
    templateName: string;
}

export interface GetInAppTemplateResult {
    readonly arn?: string;
    readonly content?: outputs.pinpoint.InAppTemplateInAppMessageContent[];
    readonly customConfig?: any;
    readonly layout?: enums.pinpoint.InAppTemplateLayout;
    readonly tags?: any;
    readonly templateDescription?: string;
}

export function getInAppTemplateOutput(args: GetInAppTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInAppTemplateResult> {
    return pulumi.output(args).apply(a => getInAppTemplate(a, opts))
}

export interface GetInAppTemplateOutputArgs {
    templateName: pulumi.Input<string>;
}
