// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::LaunchTemplateConstraint
 */
export function getLaunchTemplateConstraint(args: GetLaunchTemplateConstraintArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchTemplateConstraintResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:servicecatalog:getLaunchTemplateConstraint", {
        "id": args.id,
    }, opts);
}

export interface GetLaunchTemplateConstraintArgs {
    id: string;
}

export interface GetLaunchTemplateConstraintResult {
    readonly acceptLanguage?: string;
    readonly description?: string;
    readonly id?: string;
    readonly rules?: string;
}

export function getLaunchTemplateConstraintOutput(args: GetLaunchTemplateConstraintOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLaunchTemplateConstraintResult> {
    return pulumi.output(args).apply(a => getLaunchTemplateConstraint(a, opts))
}

export interface GetLaunchTemplateConstraintOutputArgs {
    id: pulumi.Input<string>;
}