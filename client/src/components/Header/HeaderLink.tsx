"use client";

import { Route } from "next";
import Link from "next/link";
import { usePathname } from "next/navigation";
import cn from "classnames";

interface HeaderLinkParams {
    title: string;
    href: string;
}

export function HeaderLink({ title, href }: HeaderLinkParams): JSX.Element {
    const pathname = usePathname();

    return (
        <Link
            href={href}
            className={cn("px-2 py-0.5 rounded-md", {
                ["text-neutral-600"]: pathname !== href,
                ["bg-neutral-950 text-neutral-100"]: pathname === href,
            })}
        >
            <span className="text-sm font-medium ">{title}</span>
        </Link>
    );
}
