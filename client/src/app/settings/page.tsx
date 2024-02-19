import { Settings } from "./Settings";
import { Metadata } from "next";

export const metadata: Metadata = {
    title: "Настройки",
};

export default function SettingsPage() {
    return (
        <main className="p-10 flex flex-col items-center">
            <div className="w-full max-w-xl flex flex-col gap-6">
                <h1 className="font-bold text-3xl mb-0">Настройки</h1>
                <Settings />
            </div>
        </main>
    );
}
