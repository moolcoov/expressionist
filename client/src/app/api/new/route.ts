import { addExpression } from "@/orchestra";

export async function GET(request: Request) {
    const expressionValue = request.headers.get("ExpressionValue");

    if (!expressionValue) {
        return;
    }

    const expression = await addExpression({ expressionValue });

    return Response.json({ expression });
}
