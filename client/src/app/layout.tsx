import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Header } from "@/components/Header";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: {
        default: "Expressionist - Проект для Лицея Академии Яндекс",
        template: "%s | Expressionist",
    },
    description: "Калькулятор.",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="ru">
            <body className={inter.className}>
                <Header />
                {children}
            </body>
        </html>
    );
}
