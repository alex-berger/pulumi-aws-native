// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::AppImageConfig
 */
export function getAppImageConfig(args: GetAppImageConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetAppImageConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getAppImageConfig", {
        "appImageConfigName": args.appImageConfigName,
    }, opts);
}

export interface GetAppImageConfigArgs {
    /**
     * The Name of the AppImageConfig.
     */
    appImageConfigName: string;
}

export interface GetAppImageConfigResult {
    /**
     * The Amazon Resource Name (ARN) of the AppImageConfig.
     */
    readonly appImageConfigArn?: string;
    /**
     * The KernelGatewayImageConfig.
     */
    readonly kernelGatewayImageConfig?: outputs.sagemaker.AppImageConfigKernelGatewayImageConfig;
}

export function getAppImageConfigOutput(args: GetAppImageConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppImageConfigResult> {
    return pulumi.output(args).apply(a => getAppImageConfig(a, opts))
}

export interface GetAppImageConfigOutputArgs {
    /**
     * The Name of the AppImageConfig.
     */
    appImageConfigName: pulumi.Input<string>;
}
