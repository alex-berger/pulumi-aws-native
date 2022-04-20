// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./application";
export * from "./deploymentConfig";
export * from "./deploymentGroup";
export * from "./getApplication";
export * from "./getDeploymentConfig";
export * from "./getDeploymentGroup";

// Import resources to register:
import { Application } from "./application";
import { DeploymentConfig } from "./deploymentConfig";
import { DeploymentGroup } from "./deploymentGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:codedeploy:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:codedeploy:DeploymentConfig":
                return new DeploymentConfig(name, <any>undefined, { urn })
            case "aws-native:codedeploy:DeploymentGroup":
                return new DeploymentGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "codedeploy", _module)
