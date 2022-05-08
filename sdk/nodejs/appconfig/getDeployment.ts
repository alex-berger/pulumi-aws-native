// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppConfig::Deployment
 */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appconfig:getDeployment", {
        "id": args.id,
    }, opts);
}

export interface GetDeploymentArgs {
    id: string;
}

export interface GetDeploymentResult {
    readonly id?: string;
    readonly tags?: outputs.appconfig.DeploymentTags[];
}

export function getDeploymentOutput(args: GetDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentResult> {
    return pulumi.output(args).apply(a => getDeployment(a, opts))
}

export interface GetDeploymentOutputArgs {
    id: pulumi.Input<string>;
}
