/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type NewCluster struct {
	NumWorkers int32 `json:"num_workers,omitempty"`

	ClusterName string `json:"cluster_name,omitempty"`

	SparkVersion string `json:"spark_version"`

	NodeTypeId string `json:"node_type_id"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`
}
