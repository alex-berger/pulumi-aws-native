// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Contains detailed information about a report plan in AWS Backup Audit Manager.
 */
export function getReportPlan(args: GetReportPlanArgs, opts?: pulumi.InvokeOptions): Promise<GetReportPlanResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:backup:getReportPlan", {
        "reportPlanArn": args.reportPlanArn,
    }, opts);
}

export interface GetReportPlanArgs {
    /**
     * An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.
     */
    reportPlanArn: string;
}

export interface GetReportPlanResult {
    /**
     * A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
     */
    readonly reportDeliveryChannel?: outputs.backup.ReportDeliveryChannelProperties;
    /**
     * An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.
     */
    readonly reportPlanArn?: string;
    /**
     * An optional description of the report plan with a maximum of 1,024 characters.
     */
    readonly reportPlanDescription?: string;
    /**
     * Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.
     */
    readonly reportPlanTags?: outputs.backup.ReportPlanTag[];
    /**
     * Identifies the report template for the report. Reports are built using a report template.
     */
    readonly reportSetting?: outputs.backup.ReportSettingProperties;
}

export function getReportPlanOutput(args: GetReportPlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReportPlanResult> {
    return pulumi.output(args).apply(a => getReportPlan(a, opts))
}

export interface GetReportPlanOutputArgs {
    /**
     * An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.
     */
    reportPlanArn: pulumi.Input<string>;
}
