/** @type {import('next').NextConfig} */
const nextConfig = {
    basePath: '/nextjs',
    images: {
        remotePatterns: [
            {
                protocol: 'https',
                hostname: 'images.unsplash.com'
            }
        ]
    },
    experimental: {
        serverActions: {
            allowedOrigins: ['localhost:8000']
        }
    }
};

export default nextConfig;

// http://localhost:8000/nextjs/events/11111/spots-layout

