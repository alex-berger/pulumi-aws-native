// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.
 */
export function getAssociation(args: GetAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ssm:getAssociation", {
        "associationId": args.associationId,
    }, opts);
}

export interface GetAssociationArgs {
    /**
     * Unique identifier of the association.
     */
    associationId: string;
}

export interface GetAssociationResult {
    readonly applyOnlyAtCronInterval?: boolean;
    /**
     * Unique identifier of the association.
     */
    readonly associationId?: string;
    /**
     * The name of the association.
     */
    readonly associationName?: string;
    readonly automationTargetParameterName?: string;
    readonly calendarNames?: string[];
    readonly complianceSeverity?: enums.ssm.AssociationComplianceSeverity;
    /**
     * The version of the SSM document to associate with the target.
     */
    readonly documentVersion?: string;
    /**
     * The ID of the instance that the SSM document is associated with.
     */
    readonly instanceId?: string;
    readonly maxConcurrency?: string;
    readonly maxErrors?: string;
    /**
     * The name of the SSM document.
     */
    readonly name?: string;
    readonly outputLocation?: outputs.ssm.AssociationInstanceAssociationOutputLocation;
    /**
     * Parameter values that the SSM document uses at runtime.
     */
    readonly parameters?: any;
    /**
     * A Cron or Rate expression that specifies when the association is applied to the target.
     */
    readonly scheduleExpression?: string;
    readonly scheduleOffset?: number;
    readonly syncCompliance?: enums.ssm.AssociationSyncCompliance;
    /**
     * The targets that the SSM document sends commands to.
     */
    readonly targets?: outputs.ssm.AssociationTarget[];
    readonly waitForSuccessTimeoutSeconds?: number;
}

export function getAssociationOutput(args: GetAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssociationResult> {
    return pulumi.output(args).apply(a => getAssociation(a, opts))
}

export interface GetAssociationOutputArgs {
    /**
     * Unique identifier of the association.
     */
    associationId: pulumi.Input<string>;
}
