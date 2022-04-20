// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ECS::Service
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecs:getService", {
        "cluster": args.cluster,
        "serviceArn": args.serviceArn,
    }, opts);
}

export interface GetServiceArgs {
    cluster: string;
    serviceArn: string;
}

export interface GetServiceResult {
    readonly capacityProviderStrategy?: outputs.ecs.ServiceCapacityProviderStrategyItem[];
    readonly deploymentConfiguration?: outputs.ecs.ServiceDeploymentConfiguration;
    readonly desiredCount?: number;
    readonly enableECSManagedTags?: boolean;
    readonly enableExecuteCommand?: boolean;
    readonly healthCheckGracePeriodSeconds?: number;
    readonly loadBalancers?: outputs.ecs.ServiceLoadBalancer[];
    readonly name?: string;
    readonly networkConfiguration?: outputs.ecs.ServiceNetworkConfiguration;
    readonly placementConstraints?: outputs.ecs.ServicePlacementConstraint[];
    readonly placementStrategies?: outputs.ecs.ServicePlacementStrategy[];
    readonly platformVersion?: string;
    readonly propagateTags?: enums.ecs.ServicePropagateTags;
    readonly serviceArn?: string;
    readonly serviceRegistries?: outputs.ecs.ServiceRegistry[];
    readonly tags?: outputs.ecs.ServiceTag[];
    readonly taskDefinition?: string;
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

export interface GetServiceOutputArgs {
    cluster: pulumi.Input<string>;
    serviceArn: pulumi.Input<string>;
}
