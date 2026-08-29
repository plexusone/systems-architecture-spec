/**
 * Auto-generated Zod schemas from JSON Schema.
 * DO NOT EDIT - regenerate with: npm run generate
 * Source: architecture.schema.json
 */

import { z } from 'zod';

export const AgentExtensionSchema = z.object({ "runtime": z.boolean().optional() }).strict();
export type AgentExtension = z.infer<typeof AgentExtensionSchema>;

export const AssuranceSchema = z.object({ "tests": z.array(z.string()).optional(), "metrics": z.array(z.string()).optional(), "detections": z.array(z.string()).optional(), "deployment": z.array(z.string()).optional() }).strict();
export type Assurance = z.infer<typeof AssuranceSchema>;

export const EntitlementSchema = z.object({ "subject": z.string(), "action": z.string(), "resource": z.string(), "conditions": z.array(z.string()).optional() }).strict();
export type Entitlement = z.infer<typeof EntitlementSchema>;

export const AuthorizationSchema = z.object({ "entitlements": z.array(z.object({ "subject": z.string(), "action": z.string(), "resource": z.string(), "conditions": z.array(z.string()).optional() }).strict()).optional() }).strict();
export type Authorization = z.infer<typeof AuthorizationSchema>;

export const ComplianceBoundarySchema = z.object({ "fedrampBoundary": z.boolean().optional(), "controls": z.array(z.string()).optional() }).strict();
export type ComplianceBoundary = z.infer<typeof ComplianceBoundarySchema>;

export const SecurityExtensionSchema = z.object({ "trustZone": z.string().optional() }).strict();
export type SecurityExtension = z.infer<typeof SecurityExtensionSchema>;

export const SREExtensionSchema = z.object({ "availabilityTarget": z.string().optional() }).strict();
export type SREExtension = z.infer<typeof SREExtensionSchema>;

export const ComplianceExtensionSchema = z.object({ "fedrampBoundary": z.boolean().optional() }).strict();
export type ComplianceExtension = z.infer<typeof ComplianceExtensionSchema>;

export const ExtensionsSchema = z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict();
export type Extensions = z.infer<typeof ExtensionsSchema>;

export const BoundarySchema = z.object({ "id": z.string(), "kind": z.string(), "name": z.string(), "attributes": z.record(z.string(), z.string()).optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional(), "controls": z.array(z.string()).optional() }).strict().optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict();
export type Boundary = z.infer<typeof BoundarySchema>;

export const DataFlowSchema = z.object({ "classifications": z.array(z.string()).optional() }).strict();
export type DataFlow = z.infer<typeof DataFlowSchema>;

export const ExternalRefSchema = z.object({ "type": z.string(), "ref": z.string() }).strict();
export type ExternalRef = z.infer<typeof ExternalRefSchema>;

export const IdentitySchema = z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict();
export type Identity = z.infer<typeof IdentitySchema>;

export const IdentityRefSchema = z.object({ "nodeId": z.string().optional(), "identity": z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict().optional() }).strict();
export type IdentityRef = z.infer<typeof IdentityRefSchema>;

export const MetadataSchema = z.object({ "name": z.string().optional(), "description": z.string().optional(), "owner": z.string().optional() }).strict();
export type Metadata = z.infer<typeof MetadataSchema>;

export const TechnologySchema = z.object({ "provider": z.string(), "service": z.string().optional() }).strict();
export type Technology = z.infer<typeof TechnologySchema>;

export const NodeSchema = z.object({ "id": z.string(), "kind": z.string(), "name": z.string(), "owner": z.string().optional(), "technology": z.object({ "provider": z.string(), "service": z.string().optional() }).strict().optional(), "identity": z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict().optional(), "boundaries": z.array(z.string()).optional(), "criticality": z.string().optional(), "assurance": z.object({ "tests": z.array(z.string()).optional(), "metrics": z.array(z.string()).optional(), "detections": z.array(z.string()).optional(), "deployment": z.array(z.string()).optional() }).strict().optional(), "refs": z.array(z.object({ "type": z.string(), "ref": z.string() }).strict()).optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict();
export type Node = z.infer<typeof NodeSchema>;

export const TransportSchema = z.object({ "protocol": z.string().optional(), "port": z.number().int().optional(), "applicationProtocol": z.string().optional(), "encryption": z.string().optional() }).strict();
export type Transport = z.infer<typeof TransportSchema>;

export const RelationshipSchema = z.object({ "id": z.string(), "from": z.string(), "to": z.string(), "kind": z.string(), "transport": z.object({ "protocol": z.string().optional(), "port": z.number().int().optional(), "applicationProtocol": z.string().optional(), "encryption": z.string().optional() }).strict().optional(), "operations": z.array(z.string()).optional(), "identity": z.object({ "nodeId": z.string().optional(), "identity": z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict().optional() }).strict().optional(), "authorization": z.object({ "entitlements": z.array(z.object({ "subject": z.string(), "action": z.string(), "resource": z.string(), "conditions": z.array(z.string()).optional() }).strict()).optional() }).strict().optional(), "data": z.object({ "classifications": z.array(z.string()).optional() }).strict().optional(), "crossesBoundaries": z.array(z.string()).optional(), "criticalPath": z.boolean().optional(), "sync": z.string().optional(), "assurance": z.object({ "tests": z.array(z.string()).optional(), "metrics": z.array(z.string()).optional(), "detections": z.array(z.string()).optional(), "deployment": z.array(z.string()).optional() }).strict().optional(), "protocolRef": z.string().optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict();
export type Relationship = z.infer<typeof RelationshipSchema>;

export const ArchitectureSchema = z.object({ "version": z.string(), "metadata": z.object({ "name": z.string().optional(), "description": z.string().optional(), "owner": z.string().optional() }).strict().optional(), "nodes": z.array(z.object({ "id": z.string(), "kind": z.string(), "name": z.string(), "owner": z.string().optional(), "technology": z.object({ "provider": z.string(), "service": z.string().optional() }).strict().optional(), "identity": z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict().optional(), "boundaries": z.array(z.string()).optional(), "criticality": z.string().optional(), "assurance": z.object({ "tests": z.array(z.string()).optional(), "metrics": z.array(z.string()).optional(), "detections": z.array(z.string()).optional(), "deployment": z.array(z.string()).optional() }).strict().optional(), "refs": z.array(z.object({ "type": z.string(), "ref": z.string() }).strict()).optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict()), "relationships": z.array(z.object({ "id": z.string(), "from": z.string(), "to": z.string(), "kind": z.string(), "transport": z.object({ "protocol": z.string().optional(), "port": z.number().int().optional(), "applicationProtocol": z.string().optional(), "encryption": z.string().optional() }).strict().optional(), "operations": z.array(z.string()).optional(), "identity": z.object({ "nodeId": z.string().optional(), "identity": z.object({ "type": z.string(), "mechanism": z.string().optional(), "ref": z.string().optional() }).strict().optional() }).strict().optional(), "authorization": z.object({ "entitlements": z.array(z.object({ "subject": z.string(), "action": z.string(), "resource": z.string(), "conditions": z.array(z.string()).optional() }).strict()).optional() }).strict().optional(), "data": z.object({ "classifications": z.array(z.string()).optional() }).strict().optional(), "crossesBoundaries": z.array(z.string()).optional(), "criticalPath": z.boolean().optional(), "sync": z.string().optional(), "assurance": z.object({ "tests": z.array(z.string()).optional(), "metrics": z.array(z.string()).optional(), "detections": z.array(z.string()).optional(), "deployment": z.array(z.string()).optional() }).strict().optional(), "protocolRef": z.string().optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict()).optional(), "boundaries": z.array(z.object({ "id": z.string(), "kind": z.string(), "name": z.string(), "attributes": z.record(z.string(), z.string()).optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional(), "controls": z.array(z.string()).optional() }).strict().optional(), "extensions": z.object({ "security": z.object({ "trustZone": z.string().optional() }).strict().optional(), "sre": z.object({ "availabilityTarget": z.string().optional() }).strict().optional(), "compliance": z.object({ "fedrampBoundary": z.boolean().optional() }).strict().optional(), "agent": z.object({ "runtime": z.boolean().optional() }).strict().optional() }).strict().optional() }).strict()).optional() }).strict();
export type Architecture = z.infer<typeof ArchitectureSchema>;

