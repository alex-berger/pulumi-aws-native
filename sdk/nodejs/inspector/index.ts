// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./assessmentTarget";
export * from "./assessmentTemplate";
export * from "./getAssessmentTarget";
export * from "./getAssessmentTemplate";
export * from "./getResourceGroup";
export * from "./resourceGroup";

// Import resources to register:
import { AssessmentTarget } from "./assessmentTarget";
import { AssessmentTemplate } from "./assessmentTemplate";
import { ResourceGroup } from "./resourceGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:inspector:AssessmentTarget":
                return new AssessmentTarget(name, <any>undefined, { urn })
            case "aws-native:inspector:AssessmentTemplate":
                return new AssessmentTemplate(name, <any>undefined, { urn })
            case "aws-native:inspector:ResourceGroup":
                return new ResourceGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "inspector", _module)