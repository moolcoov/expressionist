import { usePathname } from "next/navigation";
import { HeaderLink } from "./HeaderLink";

export function Header(): JSX.Element {
    return (
        <header className="w-full flex justify-center p-2">
            <nav className="flex gap-3">
                <HeaderLink title="Главная" href="/" />
                <HeaderLink title="Агенты" href="/agents" />
                <HeaderLink title="Настройки" href="/settings" />
            </nav>
        </header>
    );
}
