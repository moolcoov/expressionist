import { OrchestraExpression, OrchestraExpressionStatus } from "@/orchestra";
import moment from "moment";
import cn from "classnames";

export function ExpressionCard({
    expression,
}: {
    expression: OrchestraExpression;
}): JSX.Element {
    return (
        <div
            className={cn(
                "border border-gray-300 rounded-md p-5 flex flex-col gap-3",
                {
                    ["bg-yellow-100 border-yellow-600"]:
                        expression.status === "in_progress",
                    ["bg-green-100 border-green-600"]:
                        expression.status === "calculated",
                    ["bg-red-100 border-red-600"]:
                        expression.status === "errored",
                }
            )}
        >
            <div
                className={cn("flex justify-between font-semibold text-sm", {
                    ["text-neutral-700"]: expression.status === "created",
                    ["text-yellow-700"]: expression.status === "in_progress",
                    ["text-green-700"]: expression.status === "calculated",
                    ["text-red-700"]: expression.status === "errored",
                })}
            >
                <span>{getStatus(expression.status)}</span>
                {(expression.status === "calculated" ||
                    expression.status === "errored") && (
                    <span>
                        {moment(expression.calculationTime.Time).format(
                            "hh:mm:ss DD.MM.YY"
                        )}
                    </span>
                )}
                {expression.status === "in_progress" && (
                    <span>{expression.agentId}</span>
                )}
            </div>
            <div
                className={cn("font-bold text-2xl flex justify-between", {
                    ["text-yellow-900"]: expression.status === "in_progress",
                    ["text-green-900"]: expression.status === "calculated",
                    ["text-red-900"]: expression.status === "errored",
                })}
            >
                <h3>{expression.expression}</h3>
                {expression.status === "calculated" && (
                    <h3>Ответ: {expression.res}</h3>
                )}
            </div>
        </div>
    );
}

function getStatus(status: OrchestraExpressionStatus): string {
    switch (status) {
        case "created":
            return "Создано";
        case "calculated":
            return "Вычислено";
        case "in_progress":
            return "Вычисляется";
        case "errored":
            return "Ошибка";
    }
}
