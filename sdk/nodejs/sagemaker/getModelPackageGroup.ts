// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelPackageGroup
 */
export function getModelPackageGroup(args: GetModelPackageGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetModelPackageGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getModelPackageGroup", {
        "modelPackageGroupArn": args.modelPackageGroupArn,
    }, opts);
}

export interface GetModelPackageGroupArgs {
    modelPackageGroupArn: string;
}

export interface GetModelPackageGroupResult {
    /**
     * The time at which the model package group was created.
     */
    readonly creationTime?: string;
    readonly modelPackageGroupArn?: string;
    readonly modelPackageGroupPolicy?: any;
    /**
     * The status of a modelpackage group job.
     */
    readonly modelPackageGroupStatus?: enums.sagemaker.ModelPackageGroupStatus;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.sagemaker.ModelPackageGroupTag[];
}

export function getModelPackageGroupOutput(args: GetModelPackageGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelPackageGroupResult> {
    return pulumi.output(args).apply(a => getModelPackageGroup(a, opts))
}

export interface GetModelPackageGroupOutputArgs {
    modelPackageGroupArn: pulumi.Input<string>;
}