import { Agents } from "./Agents";

export default function AgentsPage() {
    return (
        <main className="p-10 flex flex-col items-center">
            <div className="w-full max-w-xl">
                <h1 className="font-bold text-3xl mb-6">Агенты</h1>
                <Agents />
            </div>
        </main>
    );
}
