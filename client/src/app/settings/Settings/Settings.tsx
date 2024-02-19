"use client";

import { OrchestraSettings, fetcher } from "@/orchestra";
import useSWR, { useSWRConfig } from "swr";
import { Setting } from "./Setting";
import { useRef, useState } from "react";
import cn from "classnames";
import { Spinner } from "@/components/Spinner";

export function Settings(): JSX.Element {
    // Для изменения состояния кнопки
    const [isChanged, setChanged] = useState(false);

    // Для получения данных из полей
    const additionRef = useRef<HTMLInputElement>(null);
    const subtractionRef = useRef<HTMLInputElement>(null);
    const multiplicationRef = useRef<HTMLInputElement>(null);
    const divisionRef = useRef<HTMLInputElement>(null);
    const inactiveAgentRef = useRef<HTMLInputElement>(null);

    // Получение настроек с бэка
    const { data, error, isLoading } = useSWR<OrchestraSettings>(
        `${process.env.BACKEND_ADDRESS_CLIENT}/settings`,
        (url: string) =>
            fetcher(url, {
                cache: "no-store",
            })
    );

    // Для ревалидации при изменении настроек
    const { mutate } = useSWRConfig();

    // При загрузке
    if (isLoading) {
        return (
            <div className="w-full h-32 relative">
                <Spinner />
            </div>
        );
    }

    // При ошибке
    if (error || !data) {
        return (
            <div className="flex flex-col items-center">
                <h2 className="font-semibold text-xl mb-3">
                    Произошла ошибка при загрузке
                </h2>
                <p>{error.message}</p>
            </div>
        );
    }

    function onChange() {
        setChanged(true);
    }

    async function onSubmit() {
        const additionValue = additionRef.current?.value;
        const subtractionValue = subtractionRef.current?.value;
        const multiplicationValue = multiplicationRef.current?.value;
        const divisionValue = divisionRef.current?.value;
        const inactiveAgentValue = inactiveAgentRef.current?.value;

        if (
            !additionValue ||
            !subtractionValue ||
            !multiplicationValue ||
            !divisionValue ||
            !inactiveAgentValue
        ) {
            return;
        }

        await fetch("/api/settings/update", {
            method: "POST",
            body: JSON.stringify({
                additionTime: Number(additionValue),
                subtractionTime: Number(subtractionValue),
                multiplicationTime: Number(multiplicationValue),
                divisionTime: Number(divisionValue),
                inactiveAgentTime: Number(inactiveAgentValue),
            }),
        });

        mutate("http://localhost:8080/settings");
    }

    return (
        <form className="flex flex-col gap-5" onSubmit={onSubmit}>
            <Setting
                slug="additionTime"
                label="Время сложения (в мс)"
                defaultValue={data.additionTime}
                symbol="+"
                ref={additionRef}
                onChange={onChange}
            />
            <Setting
                slug="subtractionTime"
                label="Время вычитания (в мс)"
                defaultValue={data.subtractionTime}
                symbol="-"
                ref={subtractionRef}
                onChange={onChange}
            />
            <Setting
                slug="multiplicationTime"
                label="Время умножения (в мс)"
                defaultValue={data.multiplicationTime}
                symbol="*"
                ref={multiplicationRef}
                onChange={onChange}
            />
            <Setting
                slug="divisionTime"
                label="Время деления (в мс)"
                defaultValue={data.divisionTime}
                symbol="/"
                ref={divisionRef}
                onChange={onChange}
            />
            <Setting
                slug="inactiveAgentTime"
                label="Время отображения неактивного агента (в с)"
                defaultValue={data.inactiveAgentTime}
                ref={inactiveAgentRef}
                onChange={onChange}
            />
            <button
                type="submit"
                disabled={!isChanged}
                className={cn(
                    "bg-neutral-950 text-slate-100 p-2 rounded-md transition-colors disabled:opacity-30",
                    {
                        ["hover:bg-neutral-700 active:bg-neutral-800"]:
                            isChanged,
                    }
                )}
            >
                Применить
            </button>
        </form>
    );
}
