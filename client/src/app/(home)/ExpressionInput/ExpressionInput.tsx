"use client";
import { useState } from "react";
import { useSWRConfig } from "swr";

export function ExpressionInput(): JSX.Element {
    // Для ревалидации при добавлении нового выражения
    const { mutate } = useSWRConfig();

    // Для ввода выражения
    const [query, setQuery] = useState("");

    // При отправке выражения
    async function onSubmit() {
        if (query === "") {
            return;
        }

        await fetch("/api/new", {
            headers: {
                ExpressionValue: query,
            },
        });

        // Ревалидация выражений
        mutate("http://localhost:8080/list");
    }

    // При изменении введенного выражения
    function onChange(e: React.ChangeEvent<HTMLInputElement>) {
        setQuery(e.target.value);
    }

    return (
        <form className="flex gap-3" onSubmit={onSubmit}>
            <input
                name="expressionValue"
                className="border border-neutral-300 rounded-md w-full p-3 transition-colors hover:border-neutral-400"
                placeholder="Введите выражение"
                value={query}
                onChange={onChange}
            />
            <button
                type="submit"
                className="bg-neutral-950 text-slate-100 p-3 rounded-md transition-colors hover:bg-neutral-700 active:bg-neutral-800"
            >
                Вычислить
            </button>
        </form>
    );
}
