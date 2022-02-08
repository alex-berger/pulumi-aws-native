// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate
 */
export function getConfigurationTemplate(args: GetConfigurationTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:elasticbeanstalk:getConfigurationTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetConfigurationTemplateArgs {
    id: string;
}

export interface GetConfigurationTemplateResult {
    readonly description?: string;
    readonly id?: string;
    readonly optionSettings?: outputs.elasticbeanstalk.ConfigurationTemplateConfigurationOptionSetting[];
}

export function getConfigurationTemplateOutput(args: GetConfigurationTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationTemplateResult> {
    return pulumi.output(args).apply(a => getConfigurationTemplate(a, opts))
}

export interface GetConfigurationTemplateOutputArgs {
    id: pulumi.Input<string>;
}