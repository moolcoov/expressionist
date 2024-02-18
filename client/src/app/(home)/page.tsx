import { Expressions } from "./Expressions";
import { ExpressionInput } from "./ExpressionInput";

export default function HomePage() {
    return (
        <main className="p-10 flex flex-col items-center">
            <div className="w-full max-w-xl mt-60 mb-40">
                <ExpressionInput />
            </div>
            <div className="w-full max-w-xl">
                <h1 className="font-bold text-3xl mb-6">Выражения</h1>
                <Expressions />
            </div>
        </main>
    );
}
