"use client";
export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <main>
            <div className="md:p-20">{children}</div>
        </main>
    );
}
