import { ForwardedRef, forwardRef } from "react";

interface SettingProps {
    slug: string;
    label: string;
    defaultValue: number;
    symbol?: string;
    onChange: () => void;
}

function setting(
    { slug, label, defaultValue, symbol, onChange }: SettingProps,
    ref: ForwardedRef<HTMLInputElement>
): JSX.Element {
    return (
        <div className="flex flex-col gap-2">
            <div className="flex justify-between items-center">
                <label htmlFor={slug} className="font-semibold">
                    {label}
                </label>
                {symbol && <span className="font-bold">{symbol}</span>}
            </div>
            <input
                type="number"
                name={slug}
                id={slug}
                defaultValue={defaultValue}
                className="border border-neutral-300 rounded-md transition-colors hover:border-neutral-400 p-2"
                ref={ref}
                onChange={onChange}
                min={0}
                max={999_999_999_999}
                required
            />
        </div>
    );
}

export const Setting = forwardRef<HTMLInputElement, SettingProps>(setting);
