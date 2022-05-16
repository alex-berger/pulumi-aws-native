// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodeDeploy::DeploymentGroup
 */
export function getDeploymentGroup(args: GetDeploymentGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:codedeploy:getDeploymentGroup", {
        "id": args.id,
    }, opts);
}

export interface GetDeploymentGroupArgs {
    id: string;
}

export interface GetDeploymentGroupResult {
    readonly alarmConfiguration?: outputs.codedeploy.DeploymentGroupAlarmConfiguration;
    readonly autoRollbackConfiguration?: outputs.codedeploy.DeploymentGroupAutoRollbackConfiguration;
    readonly autoScalingGroups?: string[];
    readonly blueGreenDeploymentConfiguration?: outputs.codedeploy.DeploymentGroupBlueGreenDeploymentConfiguration;
    readonly deployment?: outputs.codedeploy.DeploymentGroupDeployment;
    readonly deploymentConfigName?: string;
    readonly deploymentStyle?: outputs.codedeploy.DeploymentGroupDeploymentStyle;
    readonly eCSServices?: outputs.codedeploy.DeploymentGroupECSService[];
    readonly ec2TagFilters?: outputs.codedeploy.DeploymentGroupEC2TagFilter[];
    readonly ec2TagSet?: outputs.codedeploy.DeploymentGroupEC2TagSet;
    readonly id?: string;
    readonly loadBalancerInfo?: outputs.codedeploy.DeploymentGroupLoadBalancerInfo;
    readonly onPremisesInstanceTagFilters?: outputs.codedeploy.DeploymentGroupTagFilter[];
    readonly onPremisesTagSet?: outputs.codedeploy.DeploymentGroupOnPremisesTagSet;
    readonly outdatedInstancesStrategy?: string;
    readonly serviceRoleArn?: string;
    readonly tags?: outputs.codedeploy.DeploymentGroupTag[];
    readonly triggerConfigurations?: outputs.codedeploy.DeploymentGroupTriggerConfig[];
}

export function getDeploymentGroupOutput(args: GetDeploymentGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentGroupResult> {
    return pulumi.output(args).apply(a => getDeploymentGroup(a, opts))
}

export interface GetDeploymentGroupOutputArgs {
    id: pulumi.Input<string>;
}