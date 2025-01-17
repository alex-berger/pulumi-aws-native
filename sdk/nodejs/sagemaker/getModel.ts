// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Model
 */
export function getModel(args: GetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetModelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getModel", {
        "id": args.id,
    }, opts);
}

export interface GetModelArgs {
    id: string;
}

export interface GetModelResult {
    readonly id?: string;
    readonly tags?: outputs.sagemaker.ModelTag[];
}

export function getModelOutput(args: GetModelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelResult> {
    return pulumi.output(args).apply(a => getModel(a, opts))
}

export interface GetModelOutputArgs {
    id: pulumi.Input<string>;
}
