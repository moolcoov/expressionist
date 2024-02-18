import { OrchestraAgent } from "@/orchestra";
import moment from "moment";
import cn from "classnames";

export function Agent({ id, pingTime, status }: OrchestraAgent): JSX.Element {
    return (
        <div
            className={cn(
                "border border-gray-300 rounded-md p-5 flex flex-col gap-3",
                {
                    ["bg-green-100 border-green-600"]: status === "active",
                    ["bg-gray-100 border-gray-600"]: status === "inactive",
                }
            )}
        >
            <div
                className={cn("flex justify-between font-semibold text-sm", {
                    ["text-green-700"]: status === "active",
                    ["text-gray-700"]: status === "inactive",
                })}
            >
                <span>{getStatus(status)}</span>
                <span>{moment(pingTime).format("hh:mm:ss DD.MM.YY")}</span>
            </div>
            <h2
                className={cn("font-bold text-xl flex justify-between", {
                    ["text-green-900"]: status === "active",
                })}
            >
                Агент
                <br />
                {id}
            </h2>
        </div>
    );
}

function getStatus(status: "active" | "inactive"): string {
    switch (status) {
        case "active":
            return "Активен";
        case "inactive":
            return "Неактивен";
    }
}
