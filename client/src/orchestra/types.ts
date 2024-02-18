import { UUID } from "crypto";

export type OrchestraExpressions = OrchestraExpression[];

export interface OrchestraExpression {
    id: UUID;
    expression: string;
    dispatchTime: { Time: Date; Valid: boolean };
    calculationTime: { Time: Date; Valid: boolean };
    res: string;
    status: OrchestraExpressionStatus;
    isSent: boolean;
    agentId: UUID;
}

export type OrchestraExpressionStatus =
    | "created"
    | "calculated"
    | "in_progress"
    | "errored";

export interface OrchestraSettings {
    additionTime: number;
    subtractionTime: number;
    multiplicationTime: number;
    divisionTime: number;
    inactiveAgentTime: number;
}

export interface OrchestraAgents {
    agents?: OrchestraAgent[];
}

export interface OrchestraAgent {
    id: UUID;
    pingTime: Date;
    status: "active" | "inactive";
}
