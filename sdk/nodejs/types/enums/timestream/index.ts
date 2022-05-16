// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ScheduledQueryDimensionValueType = {
    Varchar: "VARCHAR",
} as const;

/**
 * Type for the dimension.
 */
export type ScheduledQueryDimensionValueType = (typeof ScheduledQueryDimensionValueType)[keyof typeof ScheduledQueryDimensionValueType];

export const ScheduledQueryEncryptionOption = {
    SseS3: "SSE_S3",
    SseKms: "SSE_KMS",
} as const;

/**
 * Encryption at rest options for the error reports. If no encryption option is specified, Timestream will choose SSE_S3 as default.
 */
export type ScheduledQueryEncryptionOption = (typeof ScheduledQueryEncryptionOption)[keyof typeof ScheduledQueryEncryptionOption];

export const ScheduledQueryMixedMeasureMappingMeasureValueType = {
    Bigint: "BIGINT",
    Boolean: "BOOLEAN",
    Double: "DOUBLE",
    Varchar: "VARCHAR",
    Multi: "MULTI",
} as const;

/**
 * Type of the value that is to be read from SourceColumn. If the mapping is for MULTI, use MeasureValueType.MULTI.
 */
export type ScheduledQueryMixedMeasureMappingMeasureValueType = (typeof ScheduledQueryMixedMeasureMappingMeasureValueType)[keyof typeof ScheduledQueryMixedMeasureMappingMeasureValueType];

export const ScheduledQueryMultiMeasureAttributeMappingMeasureValueType = {
    Bigint: "BIGINT",
    Boolean: "BOOLEAN",
    Double: "DOUBLE",
    Varchar: "VARCHAR",
} as const;

/**
 * Value type of the measure value column to be read from the query result.
 */
export type ScheduledQueryMultiMeasureAttributeMappingMeasureValueType = (typeof ScheduledQueryMultiMeasureAttributeMappingMeasureValueType)[keyof typeof ScheduledQueryMultiMeasureAttributeMappingMeasureValueType];
