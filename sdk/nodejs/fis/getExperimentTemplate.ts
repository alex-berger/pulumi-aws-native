// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::FIS::ExperimentTemplate
 */
export function getExperimentTemplate(args: GetExperimentTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetExperimentTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:fis:getExperimentTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetExperimentTemplateArgs {
    id: string;
}

export interface GetExperimentTemplateResult {
    readonly actions?: outputs.fis.ExperimentTemplateActionMap;
    readonly description?: string;
    readonly id?: string;
    readonly logConfiguration?: outputs.fis.ExperimentTemplateLogConfiguration;
    readonly roleArn?: string;
    readonly stopConditions?: outputs.fis.ExperimentTemplateStopCondition[];
    readonly targets?: outputs.fis.ExperimentTemplateTargetMap;
}

export function getExperimentTemplateOutput(args: GetExperimentTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExperimentTemplateResult> {
    return pulumi.output(args).apply(a => getExperimentTemplate(a, opts))
}

export interface GetExperimentTemplateOutputArgs {
    id: pulumi.Input<string>;
}
