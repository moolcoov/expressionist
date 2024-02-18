export const fetcher = (
    input: string | URL | globalThis.Request,
    init?: RequestInit
) => fetch(input, init).then((res) => res.json());
