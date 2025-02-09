/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./web/templates/**/*.html"],
    theme: {
        extend: {
            typography: {
                DEFAULT: {
                    css: {
                        maxWidth: '65ch',
                        color: '#333',
                        strong: {
                            color: '#333',
                        },
                    },
                },
            },
        },
    },
    plugins: [
        require('@tailwindcss/typography'),
    ],
}