import { useEffect, useState } from 'react';

const CurrentPrice = () => {
    const [price, setPrice] = useState<number | null>(null);

    useEffect(() => {
        const interval = setInterval(() => {
            fetch("http://localhost:8080/price") // 仮URL
                .then(res => res.json())
                .then(data => setPrice(data.price))
                .catch(err => console.error("Price fetch failed", err));
        }, 1000); // 1秒ごとに価格取得（雑にポーリング）

        return () => clearInterval(interval);
    }, []);

    return (
        <div style={{ fontSize: '2rem', fontWeight: 'bold' }}>
            {price !== null ? `USD/JPY: ${price.toFixed(2)}` : 'Loading...'}
        </div>
    );
};

export default CurrentPrice;

