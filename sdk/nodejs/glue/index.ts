// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./classifier";
export * from "./connection";
export * from "./crawler";
export * from "./dataCatalogEncryptionSettings";
export * from "./database";
export * from "./devEndpoint";
export * from "./getClassifier";
export * from "./getConnection";
export * from "./getCrawler";
export * from "./getDataCatalogEncryptionSettings";
export * from "./getDatabase";
export * from "./getDevEndpoint";
export * from "./getJob";
export * from "./getMLTransform";
export * from "./getPartition";
export * from "./getRegistry";
export * from "./getSchema";
export * from "./getSchemaVersion";
export * from "./getSecurityConfiguration";
export * from "./getTable";
export * from "./getTrigger";
export * from "./getWorkflow";
export * from "./job";
export * from "./mltransform";
export * from "./partition";
export * from "./registry";
export * from "./schema";
export * from "./schemaVersion";
export * from "./schemaVersionMetadata";
export * from "./securityConfiguration";
export * from "./table";
export * from "./trigger";
export * from "./workflow";

// Export enums:
export * from "../types/enums/glue";

// Import resources to register:
import { Classifier } from "./classifier";
import { Connection } from "./connection";
import { Crawler } from "./crawler";
import { DataCatalogEncryptionSettings } from "./dataCatalogEncryptionSettings";
import { Database } from "./database";
import { DevEndpoint } from "./devEndpoint";
import { Job } from "./job";
import { MLTransform } from "./mltransform";
import { Partition } from "./partition";
import { Registry } from "./registry";
import { Schema } from "./schema";
import { SchemaVersion } from "./schemaVersion";
import { SchemaVersionMetadata } from "./schemaVersionMetadata";
import { SecurityConfiguration } from "./securityConfiguration";
import { Table } from "./table";
import { Trigger } from "./trigger";
import { Workflow } from "./workflow";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:glue:Classifier":
                return new Classifier(name, <any>undefined, { urn })
            case "aws-native:glue:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "aws-native:glue:Crawler":
                return new Crawler(name, <any>undefined, { urn })
            case "aws-native:glue:DataCatalogEncryptionSettings":
                return new DataCatalogEncryptionSettings(name, <any>undefined, { urn })
            case "aws-native:glue:Database":
                return new Database(name, <any>undefined, { urn })
            case "aws-native:glue:DevEndpoint":
                return new DevEndpoint(name, <any>undefined, { urn })
            case "aws-native:glue:Job":
                return new Job(name, <any>undefined, { urn })
            case "aws-native:glue:MLTransform":
                return new MLTransform(name, <any>undefined, { urn })
            case "aws-native:glue:Partition":
                return new Partition(name, <any>undefined, { urn })
            case "aws-native:glue:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "aws-native:glue:Schema":
                return new Schema(name, <any>undefined, { urn })
            case "aws-native:glue:SchemaVersion":
                return new SchemaVersion(name, <any>undefined, { urn })
            case "aws-native:glue:SchemaVersionMetadata":
                return new SchemaVersionMetadata(name, <any>undefined, { urn })
            case "aws-native:glue:SecurityConfiguration":
                return new SecurityConfiguration(name, <any>undefined, { urn })
            case "aws-native:glue:Table":
                return new Table(name, <any>undefined, { urn })
            case "aws-native:glue:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            case "aws-native:glue:Workflow":
                return new Workflow(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "glue", _module)
