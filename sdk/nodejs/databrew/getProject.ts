// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Project.
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:databrew:getProject", {
        "name": args.name,
    }, opts);
}

export interface GetProjectArgs {
    /**
     * Project name
     */
    name: string;
}

export interface GetProjectResult {
    /**
     * Dataset name
     */
    readonly datasetName?: string;
    /**
     * Recipe name
     */
    readonly recipeName?: string;
    /**
     * Role arn
     */
    readonly roleArn?: string;
    /**
     * Sample
     */
    readonly sample?: outputs.databrew.ProjectSample;
}

export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

export interface GetProjectOutputArgs {
    /**
     * Project name
     */
    name: pulumi.Input<string>;
}