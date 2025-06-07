import { useEffect, useState } from 'react';

const CurrentPrice = () => {
    const [price, setPrice] = useState<number | null>(null);

    useEffect(() => {
        const ws = new WebSocket(`${window.location.origin.replace(/^http/, 'ws')}/ws/price`);

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setPrice(data.price);
        };

        return () => ws.close();
    }, []);

    return (
        <div style={{ fontSize: '2rem', fontWeight: 'bold' }}>
            {price !== null ? `USD/JPY: ${price.toFixed(2)}` : 'Loading...'}
        </div>
    );
};

export default CurrentPrice;

