// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./application";
export * from "./applicationVersion";
export * from "./configurationTemplate";
export * from "./environment";
export * from "./getApplication";
export * from "./getApplicationVersion";
export * from "./getConfigurationTemplate";
export * from "./getEnvironment";

// Import resources to register:
import { Application } from "./application";
import { ApplicationVersion } from "./applicationVersion";
import { ConfigurationTemplate } from "./configurationTemplate";
import { Environment } from "./environment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:elasticbeanstalk:Application":
                return new Application(name, <any>undefined, { urn })
            case "aws-native:elasticbeanstalk:ApplicationVersion":
                return new ApplicationVersion(name, <any>undefined, { urn })
            case "aws-native:elasticbeanstalk:ConfigurationTemplate":
                return new ConfigurationTemplate(name, <any>undefined, { urn })
            case "aws-native:elasticbeanstalk:Environment":
                return new Environment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "elasticbeanstalk", _module)
