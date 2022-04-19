// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getIPSet";
export * from "./getLoggingConfiguration";
export * from "./getRegexPatternSet";
export * from "./getRuleGroup";
export * from "./getWebACL";
export * from "./getWebACLAssociation";
export * from "./ipset";
export * from "./loggingConfiguration";
export * from "./regexPatternSet";
export * from "./ruleGroup";
export * from "./webACL";
export * from "./webACLAssociation";

// Export enums:
export * from "../types/enums/wafv2";

// Import resources to register:
import { IPSet } from "./ipset";
import { LoggingConfiguration } from "./loggingConfiguration";
import { RegexPatternSet } from "./regexPatternSet";
import { RuleGroup } from "./ruleGroup";
import { WebACL } from "./webACL";
import { WebACLAssociation } from "./webACLAssociation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:wafv2:IPSet":
                return new IPSet(name, <any>undefined, { urn })
            case "aws-native:wafv2:LoggingConfiguration":
                return new LoggingConfiguration(name, <any>undefined, { urn })
            case "aws-native:wafv2:RegexPatternSet":
                return new RegexPatternSet(name, <any>undefined, { urn })
            case "aws-native:wafv2:RuleGroup":
                return new RuleGroup(name, <any>undefined, { urn })
            case "aws-native:wafv2:WebACL":
                return new WebACL(name, <any>undefined, { urn })
            case "aws-native:wafv2:WebACLAssociation":
                return new WebACLAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "wafv2", _module)
