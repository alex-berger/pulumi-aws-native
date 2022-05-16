// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::NotebookInstance
 */
export function getNotebookInstance(args: GetNotebookInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetNotebookInstanceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getNotebookInstance", {
        "id": args.id,
    }, opts);
}

export interface GetNotebookInstanceArgs {
    id: string;
}

export interface GetNotebookInstanceResult {
    readonly acceleratorTypes?: string[];
    readonly additionalCodeRepositories?: string[];
    readonly defaultCodeRepository?: string;
    readonly id?: string;
    readonly instanceType?: string;
    readonly lifecycleConfigName?: string;
    readonly roleArn?: string;
    readonly rootAccess?: string;
    readonly tags?: outputs.sagemaker.NotebookInstanceTag[];
    readonly volumeSizeInGB?: number;
}

export function getNotebookInstanceOutput(args: GetNotebookInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNotebookInstanceResult> {
    return pulumi.output(args).apply(a => getNotebookInstance(a, opts))
}

export interface GetNotebookInstanceOutputArgs {
    id: pulumi.Input<string>;
}