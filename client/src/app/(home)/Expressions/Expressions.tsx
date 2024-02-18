"use client";
import { OrchestraExpressions, fetcher } from "@/orchestra";
import useSWR from "swr";
import { ExpressionCard } from "./ExpressionCard";
import { Spinner } from "@/components/Spinner";

export function Expressions(): JSX.Element {
    // Получение выражений с бэка
    const { data, error, isLoading } = useSWR<OrchestraExpressions, Error>(
        "http://localhost:8080/list",
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

    // Когда нет выражений
    if (!data) {
        return (
            <div className="flex flex-col items-center">Выражений пока нет</div>
        );
    }

    return (
        <div className="flex flex-col gap-3">
            {[...data].map((expression) => (
                <ExpressionCard key={expression.id} expression={expression} />
            ))}
        </div>
    );
}
