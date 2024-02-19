import { OrchestraExpression, OrchestraSettings } from ".";

export async function addExpression({
    expressionValue,
}: {
    expressionValue: string;
}): Promise<OrchestraExpression> {
    const body = {
        expression: expressionValue,
        dispatchTime: { Time: new Date().toISOString() },
    };

    const res: Response = await fetch(`http://orchestra:8080/new`, {
        headers: {
            "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify(body),
    });

    if (!res.ok) {
        console.error(res.status);
        throw Error("ERR: Failed to fetch expression");
    }

    return await res.json();
}

export async function updateSettings(settings: OrchestraSettings) {
    const res = await fetch(`http://orchestra:8080/settings/update`, {
        method: "POST",
        body: JSON.stringify(settings),
    });

    if (!res.ok) {
        console.error(res.status);
        throw Error("ERR: Failed to update settings");
    }
}
