import { OrchestraSettings, updateSettings } from "@/orchestra";

export async function POST(request: Request) {
    const values: OrchestraSettings = await request.json();

    await updateSettings(values);

    return Response.json({});
}
