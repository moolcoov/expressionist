"use client";
import { Spinner } from "@/components/Spinner";
import { OrchestraAgents, fetcher } from "@/orchestra";
import useSWR from "swr";
import { Agent } from "./Agent";

export function Agents(): JSX.Element {
    // Получение агентов с бэка
    const { data, error, isLoading } = useSWR<OrchestraAgents, Error>(
        `http://localhost:8080/agents`,
        (url: string) =>
            fetcher(url, {
                cache: "no-store",
            }),
        { refreshInterval: 1000 }
    );

    // При загрузке
    if (isLoading) {
        return (
            <div className="w-full h-32 relative">
                <Spinner />
            </div>
        );
    }

    // При ошибке
    if (error) {
        return (
            <div className="flex flex-col items-center">
                <h2 className="font-semibold text-xl mb-2">
                    Произошла ошибка при загрузке
                </h2>
                <p>{error.message}</p>
            </div>
        );
    }
    console.log(data);
    // Когда нет выражений
    if (!data || !data.agents?.at(0)) {
        return (
            <div className="flex flex-col items-center">
                Агентов пока не зарегистрировано
            </div>
        );
    }

    return (
        <div className="flex flex-col gap-3">
            {data.agents.map((agent) => (
                <Agent key={agent.id} {...agent} />
            ))}
        </div>
    );
}
